package sa

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/rs/zerolog"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
)

const (
	DeletedPrefix = "deleted:"
)

var ErrNotFound = errors.New("not found")

type Operations interface {
	GetServiceAccount(ctx context.Context, name string) (*ServiceAccount, error)
	CreateServiceAccount(ctx context.Context, sa *ServiceAccountRequest) (*ServiceAccount, error)
	DeleteServiceAccount(ctx context.Context, name string) error
	ListServiceAccounts(ctx context.Context, project string) ([]*ServiceAccount, error)
	AddProjectServiceAccountPolicyBinding(ctx context.Context, project string, binding *Binding) error
	RemoveProjectServiceAccountPolicyBinding(ctx context.Context, project string, email string) error
	ListProjectServiceAccountPolicyBindings(ctx context.Context, project, email string) ([]*Binding, error)
	UpdateProjectPolicyBindingsMembers(ctx context.Context, project string, fn UpdateProjectPolicyBindingsMembersFn) error
	CreateServiceAccountKey(ctx context.Context, name string) (*ServiceAccountKeyWithPrivateKeyData, error)
	DeleteServiceAccountKey(ctx context.Context, name string) error
	ListServiceAccountKeys(ctx context.Context, name string) ([]*ServiceAccountKey, error)
}

type ServiceAccountKey struct {
	Name         string
	KeyAlgorithm string
	KeyOrigin    string
	KeyType      string
}

type ServiceAccountKeyWithPrivateKeyData struct {
	*ServiceAccountKey
	PrivateKeyData []byte
}

type Binding struct {
	Role    string
	Members []string
}

type ServiceAccountRequest struct {
	ProjectID   string
	AccountID   string
	DisplayName string
	Description string
}

func (s ServiceAccountRequest) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.ProjectID, validation.Required),
		validation.Field(&s.AccountID, validation.Required),
		validation.Field(&s.DisplayName, validation.Required),
		validation.Field(&s.Description, validation.Required),
	)
}

type ServiceAccount = iam.ServiceAccount

var _ Operations = &Client{}

type Client struct {
	endpoint    string
	disableAuth bool
}

type UpdateProjectPolicyBindingsMembersFn func(role string, members []string) []string

func (c *Client) UpdateProjectPolicyBindingsMembers(ctx context.Context, project string, fn UpdateProjectPolicyBindingsMembersFn) error {
	service, err := c.crmService(ctx)
	if err != nil {
		return err
	}

	policy, err := service.Projects.GetIamPolicy(project, &cloudresourcemanager.GetIamPolicyRequest{}).Do()
	if err != nil {
		return fmt.Errorf("getting project %s policy: %w", project, err)
	}

	for _, binding := range policy.Bindings {
		binding.Members = fn(binding.Role, binding.Members)
	}

	_, err = service.Projects.SetIamPolicy(project, &cloudresourcemanager.SetIamPolicyRequest{
		Policy: policy,
	}).Do()
	if err != nil {
		return fmt.Errorf("setting project %s policy: %w", project, err)
	}

	return nil
}

func (c *Client) RemoveProjectServiceAccountPolicyBinding(ctx context.Context, project string, email string) error {
	service, err := c.crmService(ctx)
	if err != nil {
		return err
	}

	policy, err := service.Projects.GetIamPolicy(project, &cloudresourcemanager.GetIamPolicyRequest{}).Do()
	if err != nil {
		var gerr *googleapi.Error
		if errors.As(err, &gerr) && gerr.Code == http.StatusNotFound {
			return fmt.Errorf("project %s: %w", project, ErrNotFound)
		}

		return fmt.Errorf("getting project %s policy: %w", project, err)
	}

	var bindings []*cloudresourcemanager.Binding

	for _, binding := range policy.Bindings {
		var members []string

		for _, member := range binding.Members {
			if member != "serviceAccount:"+email {
				members = append(members, member)
			}
		}

		if len(members) > 0 {
			bindings = append(bindings, &cloudresourcemanager.Binding{
				Role:    binding.Role,
				Members: members,
			})
		}
	}

	policy.Bindings = bindings

	_, err = service.Projects.SetIamPolicy(project, &cloudresourcemanager.SetIamPolicyRequest{
		Policy: policy,
	}).Do()
	if err != nil {
		return fmt.Errorf("setting project %s policy: %w", project, err)
	}

	return nil
}

func (c *Client) ListProjectServiceAccountPolicyBindings(ctx context.Context, project, email string) ([]*Binding, error) {
	service, err := c.crmService(ctx)
	if err != nil {
		return nil, err
	}

	policy, err := service.Projects.GetIamPolicy(project, &cloudresourcemanager.GetIamPolicyRequest{}).Do()
	if err != nil {
		var gerr *googleapi.Error
		if errors.As(err, &gerr) && gerr.Code == http.StatusNotFound {
			return nil, fmt.Errorf("project %s: %w", project, ErrNotFound)
		}

		return nil, fmt.Errorf("getting project %s policy: %w", project, err)
	}

	var bindings []*Binding

	for _, binding := range policy.Bindings {
		for _, member := range binding.Members {
			if member == "serviceAccount:"+email {
				bindings = append(bindings, &Binding{
					Role:    binding.Role,
					Members: binding.Members,
				})

				break
			}
		}
	}

	return bindings, nil
}

func (c *Client) CreateServiceAccountKey(ctx context.Context, name string) (*ServiceAccountKeyWithPrivateKeyData, error) {
	service, err := c.iamService(ctx)
	if err != nil {
		return nil, err
	}

	key, err := service.Projects.ServiceAccounts.Keys.Create(name, &iam.CreateServiceAccountKeyRequest{}).Do()
	if err != nil {
		var gerr *googleapi.Error
		if errors.As(err, &gerr) && gerr.Code == http.StatusNotFound {
			return nil, fmt.Errorf("service account %s: %w", name, ErrNotFound)
		}

		return nil, fmt.Errorf("creating service account key %s: %w", name, err)
	}

	keyMatter, err := base64.StdEncoding.DecodeString(key.PrivateKeyData)
	if err != nil {
		return nil, fmt.Errorf("decoding private key data: %w", err)
	}

	return &ServiceAccountKeyWithPrivateKeyData{
		ServiceAccountKey: &ServiceAccountKey{
			Name:         key.Name,
			KeyAlgorithm: key.KeyAlgorithm,
			KeyOrigin:    key.KeyOrigin,
			KeyType:      key.KeyType,
		},
		PrivateKeyData: keyMatter,
	}, nil
}

func (c *Client) DeleteServiceAccountKey(ctx context.Context, name string) error {
	service, err := c.iamService(ctx)
	if err != nil {
		return err
	}

	_, err = service.Projects.ServiceAccounts.Keys.Delete(name).Do()
	if err != nil {
		var gerr *googleapi.Error
		if errors.As(err, &gerr) && gerr.Code == http.StatusNotFound {
			return fmt.Errorf("service account key %s: %w", name, ErrNotFound)
		}

		return fmt.Errorf("deleting service account key %s: %w", name, err)
	}

	return nil
}

func (c *Client) ListServiceAccountKeys(ctx context.Context, name string) ([]*ServiceAccountKey, error) {
	service, err := c.iamService(ctx)
	if err != nil {
		return nil, err
	}

	keys, err := service.Projects.ServiceAccounts.Keys.List(name).Do()
	if err != nil {
		return nil, fmt.Errorf("listing service account keys %s: %w", name, err)
	}

	result := make([]*ServiceAccountKey, len(keys.Keys))
	for i, key := range keys.Keys {
		result[i] = &ServiceAccountKey{
			Name:         key.Name,
			KeyAlgorithm: key.KeyAlgorithm,
			KeyOrigin:    key.KeyOrigin,
			KeyType:      key.KeyType,
		}
	}

	return result, nil
}

func (c *Client) AddProjectServiceAccountPolicyBinding(ctx context.Context, project string, binding *Binding) error {
	service, err := c.crmService(ctx)
	if err != nil {
		return err
	}

	policy, err := service.Projects.GetIamPolicy(project, &cloudresourcemanager.GetIamPolicyRequest{}).Do()
	if err != nil {
		return fmt.Errorf("getting project %s policy: %w", project, err)
	}

	uniqueMembers := make(map[string]struct{})
	for _, member := range binding.Members {
		uniqueMembers[member] = struct{}{}
	}

	found := false

	for _, b := range policy.Bindings {
		if b.Role == binding.Role {
			for _, member := range b.Members {
				uniqueMembers[member] = struct{}{}
			}

			b.Members = maps.Keys(uniqueMembers)
			found = true
			break
		}
	}

	if !found {
		policy.Bindings = append(policy.Bindings, &cloudresourcemanager.Binding{
			Role:    binding.Role,
			Members: binding.Members,
		})
	}

	_, err = service.Projects.SetIamPolicy(project, &cloudresourcemanager.SetIamPolicyRequest{
		Policy: policy,
	}).Do()
	if err != nil {
		return fmt.Errorf("setting project %s policy: %w", project, err)
	}

	return nil
}

func (c *Client) ListServiceAccounts(ctx context.Context, project string) ([]*ServiceAccount, error) {
	service, err := c.iamService(ctx)
	if err != nil {
		return nil, err
	}

	raw, err := service.Projects.ServiceAccounts.List("projects/" + project).Do()
	if err != nil {
		return nil, fmt.Errorf("listing service accounts: %w", err)
	}

	return raw.Accounts, nil
}

func (c *Client) DeleteServiceAccount(ctx context.Context, name string) error {
	service, err := c.iamService(ctx)
	if err != nil {
		return err
	}

	_, err = service.Projects.ServiceAccounts.Delete(name).Do()
	if err != nil {
		var gerr *googleapi.Error
		if errors.As(err, &gerr) && gerr.Code == http.StatusNotFound {
			return fmt.Errorf("service account %s: %w", name, ErrNotFound)
		}

		return fmt.Errorf("deleting service account: %w", err)
	}

	return nil
}

func (c *Client) GetServiceAccount(ctx context.Context, name string) (*ServiceAccount, error) {
	service, err := c.iamService(ctx)
	if err != nil {
		return nil, err
	}

	account, err := service.Projects.ServiceAccounts.Get(name).Do()
	if err != nil {
		var gerr *googleapi.Error
		if errors.As(err, &gerr) && gerr.Code == http.StatusNotFound {
			return nil, fmt.Errorf("service account %s: %w", name, ErrNotFound)
		}

		return nil, fmt.Errorf("getting service account: %w", err)
	}

	return account, nil
}

func (c *Client) CreateServiceAccount(ctx context.Context, sa *ServiceAccountRequest) (*ServiceAccount, error) {
	if err := sa.Validate(); err != nil {
		return nil, fmt.Errorf("validating service account request: %w", err)
	}

	service, err := c.iamService(ctx)
	if err != nil {
		return nil, err
	}

	request := &iam.CreateServiceAccountRequest{
		AccountId: sa.AccountID,
		ServiceAccount: &iam.ServiceAccount{
			Description: sa.Description,
			DisplayName: sa.DisplayName,
		},
	}

	account, err := service.Projects.ServiceAccounts.Create("projects/"+sa.ProjectID, request).Do()
	if err != nil {
		return nil, fmt.Errorf("creating service account: %w", err)
	}

	return account, nil
}

func (c *Client) iamService(ctx context.Context) (*iam.Service, error) {
	var opts []option.ClientOption

	if c.disableAuth {
		opts = append(opts, option.WithoutAuthentication())
	}

	if len(c.endpoint) > 0 {
		opts = append(opts, option.WithEndpoint(c.endpoint))
	}

	service, err := iam.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("creating iam service: %w", err)
	}

	return service, nil
}

func (c *Client) crmService(ctx context.Context) (*cloudresourcemanager.Service, error) {
	var opts []option.ClientOption

	if c.disableAuth {
		opts = append(opts, option.WithoutAuthentication())
	}

	if len(c.endpoint) > 0 {
		opts = append(opts, option.WithEndpoint(c.endpoint))
	}

	service, err := cloudresourcemanager.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("creating cloudresourcemanager service: %w", err)
	}

	return service, nil
}

func NewClient(endpoint string, disableAuth bool) *Client {
	return &Client{
		endpoint:    endpoint,
		disableAuth: disableAuth,
	}
}

func ServiceAccountNameFromAccountID(project, accountID string) string {
	return "projects/" + project + "/serviceAccounts/" + accountID + "@" + project + ".iam.gserviceaccount.com"
}

func ServiceAccountNameFromEmail(project, email string) string {
	return "projects/" + project + "/serviceAccounts/" + email
}

func ServiceAccountKeyName(project, accountID, keyID string) string {
	return "projects/" + project + "/serviceAccounts/" + accountID + "@" + project + ".iam.gserviceaccount.com/keys/" + keyID
}

func RemoveDeletedMembersWithRole(roles []string, log zerolog.Logger) UpdateProjectPolicyBindingsMembersFn {
	return func(role string, members []string) []string {
		if !slices.Contains(roles, role) {
			log.Info().Str("role", role).Msg("Skipping role")

			return members
		}

		var keep, remove []string

		for _, member := range members {
			if strings.HasPrefix(member, DeletedPrefix) {
				remove = append(remove, member)

				continue
			}

			keep = append(keep, member)
		}

		log.Info().Str("role", role).Fields(map[string]interface{}{
			"removed_members": remove,
			"kept_members":    keep,
			"removed_count":   len(members) - len(keep),
		}).Msg("Removed deleted members")

		return keep
	}
}
