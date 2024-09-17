package workstations

import (
	"context"
	"fmt"
	"regexp"

	workstations "cloud.google.com/go/workstations/apiv1"
	"cloud.google.com/go/workstations/apiv1/workstationspb"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	LabelCreatedBy    = "created-by"
	LabelSubjectEmail = "subject-email"

	DefaultIdleTimeoutInSec    = 7200  // 2 hours
	DefaultRunningTimeoutInSec = 43200 // 12 hours
	DefaultBootDiskSizeInGB    = 120
	DefaultHomeDiskSizeInGB    = 100
	DefaultHomeDiskType        = "pd-ssd"
	DefaultHomeDiskFsType      = "ext4"

	MachineTypeN2DStandard2  = "n2d-standard-2"
	MachineTypeN2DStandard4  = "n2d-standard-4"
	MachineTypeN2DStandard8  = "n2d-standard-8"
	MachineTypeN2DStandard16 = "n2d-standard-16"
	MachineTypeN2DStandard32 = "n2d-standard-32"

	ContainerImageVSCode           = "us-central1-docker.pkg.dev/cloud-workstations-images/predefined/code-oss:latest"
	ContainerImageIntellijUltimate = "us-central1-docker.pkg.dev/cloud-workstations-images/predefined/intellij-ultimate:latest"
	ContainerImagePosit            = "us-central1-docker.pkg.dev/posit-images/cloud-workstations/workbench:latest"
)

var _ Operations = &Client{}

type Operations interface {
	CreateWorkstationConfig(ctx context.Context, opts *WorkstationConfigOpts) (*WorkstationConfig, error)
	UpdateWorkstationConfig(ctx context.Context, opts *WorkstationConfigUpdateOpts) (*WorkstationConfig, error)
	DeleteWorkstationConfig(ctx context.Context, opts *WorkstationConfigDeleteOpts) error
	CreateWorkstation(ctx context.Context, opts *WorkstationOpts) (*Workstation, error)
}

type WorkstationCluster struct {
	Name string
}

type WorkstationConfigOpts struct {
	// Slug is the unique identifier of the workstation
	Slug string

	// DisplayName is the human-readable name of the workstation
	DisplayName string

	// MachineType is the type of machine that will be used for the workstation, e.g.:
	// - n2d-standard-2
	// - n2d-standard-4
	// - n2d-standard-8
	// - n2d-standard-16
	// - n2d-standard-32
	MachineType string

	// ServiceAccountEmail is the email address of the service account that will be associated with the workstation,
	// which we can use to grant permissions to the workstation, e.g.:
	// - Secure Web Proxy rules
	// - VPC Service controls
	// - Login
	ServiceAccountEmail string

	// CreatedBy is the entity that created the workstation
	CreatedBy string

	// SubjectEmail is the email address of the subject that will be using the workstation
	SubjectEmail string

	// ContainerImage is the image that will be used to run the workstation
	ContainerImage string
}

type WorkstationConfigUpdateOpts struct {
	// Slug is the unique identifier of the workstation
	Slug string

	// MachineType is the type of machine that will be used for the workstation, e.g.:
	// - n2d-standard-2
	// - n2d-standard-4
	// - n2d-standard-8
	// - n2d-standard-16
	// - n2d-standard-32
	MachineType string

	// ContainerImage is the image that will be used to run the workstation
	ContainerImage string
}

type WorkstationConfigDeleteOpts struct {
	// Slug is the unique identifier of the workstation
	Slug string
}

type WorkstationOpts struct {
	// Slug is the unique identifier of the workstation
	Slug string

	// DisplayName is the human-readable name of the workstation
	DisplayName string

	// Labels applied to the resource and propagated to the underlying Compute Engine resources.
	Labels map[string]string

	// Workstation configuration
	ConfigName string
}

func (o WorkstationConfigOpts) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Slug, validation.Required, validation.Length(3, 63), validation.Match(regexp.MustCompile(`^[a-z][a-z0-9-]+[a-z0-9]$`))),
		validation.Field(&o.DisplayName, validation.Required),
		validation.Field(&o.MachineType, validation.Required, validation.In(
			MachineTypeN2DStandard2,
			MachineTypeN2DStandard4,
			MachineTypeN2DStandard8,
			MachineTypeN2DStandard16,
			MachineTypeN2DStandard32,
		)),
		validation.Field(&o.ServiceAccountEmail, validation.Required, is.EmailFormat),
		validation.Field(&o.CreatedBy, validation.Required),
		validation.Field(&o.SubjectEmail, validation.Required, is.EmailFormat),
		validation.Field(&o.ContainerImage, validation.Required, validation.In(
			ContainerImageVSCode,
			ContainerImageIntellijUltimate,
			ContainerImagePosit,
		)),
	)
}

type WorkstationConfig struct {
	Name        string
	DisplayName string
}

type Workstation struct {
	Name string
}

type Client struct {
	project              string
	location             string
	workstationClusterID string

	apiEndpoint string
	disableAuth bool
}

func (c *Client) CreateWorkstationConfig(ctx context.Context, opts *WorkstationConfigOpts) (*WorkstationConfig, error) {
	err := opts.Validate()
	if err != nil {
		return nil, err
	}

	client, err := c.newClient(ctx)
	if err != nil {
		return nil, err
	}

	op, err := client.CreateWorkstationConfig(ctx, &workstationspb.CreateWorkstationConfigRequest{
		Parent:              c.WorkstationConfigParent(),
		WorkstationConfigId: opts.Slug,
		WorkstationConfig: &workstationspb.WorkstationConfig{
			Name:        opts.Slug,
			DisplayName: opts.DisplayName,
			Annotations: nil,
			Labels: map[string]string{
				LabelCreatedBy:    opts.CreatedBy,
				LabelSubjectEmail: opts.SubjectEmail,
			},
			IdleTimeout: &durationpb.Duration{
				Seconds: DefaultIdleTimeoutInSec,
			},
			RunningTimeout: &durationpb.Duration{
				Seconds: DefaultRunningTimeoutInSec,
			},
			Host: &workstationspb.WorkstationConfig_Host{
				Config: &workstationspb.WorkstationConfig_Host_GceInstance_{
					GceInstance: &workstationspb.WorkstationConfig_Host_GceInstance{
						MachineType:                opts.MachineType,
						ServiceAccount:             opts.ServiceAccountEmail,
						Tags:                       nil, // FIXME:  lets try to avoid using this, but we might need it for some default rules
						PoolSize:                   0,
						DisablePublicIpAddresses:   true,
						EnableNestedVirtualization: false,
						ShieldedInstanceConfig: &workstationspb.WorkstationConfig_Host_GceInstance_GceShieldedInstanceConfig{
							EnableSecureBoot:          true,
							EnableVtpm:                true,
							EnableIntegrityMonitoring: true,
						},
						ConfidentialInstanceConfig: &workstationspb.WorkstationConfig_Host_GceInstance_GceConfidentialInstanceConfig{
							EnableConfidentialCompute: true,
						},
						BootDiskSizeGb: DefaultBootDiskSizeInGB,
					},
				},
			},
			PersistentDirectories: []*workstationspb.WorkstationConfig_PersistentDirectory{
				{
					DirectoryType: &workstationspb.WorkstationConfig_PersistentDirectory_GcePd{
						GcePd: &workstationspb.WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk{
							SizeGb:        DefaultHomeDiskSizeInGB,
							FsType:        DefaultHomeDiskFsType,
							DiskType:      DefaultHomeDiskType,
							ReclaimPolicy: workstationspb.WorkstationConfig_PersistentDirectory_GceRegionalPersistentDisk_DELETE,
						},
					},
					MountPath: fmt.Sprintf("/home/%s", opts.Slug), // FIXME: is this the correct path?
				},
			},
			Container: &workstationspb.WorkstationConfig_Container{
				Image:   ContainerImagePosit,
				Command: nil,
				Args:    nil,
				// FIXME: we need to set PIP_INDEX_URL=..., HTTP_PROXY=..., NO_PROXY=.adeo.no, ...
				Env: map[string]string{
					"WORKSTATION_NAME": opts.Slug,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	workstationConfig, err := op.Wait(ctx)
	if err != nil {
		return nil, err
	}

	return &WorkstationConfig{
		Name:        workstationConfig.GetName(),
		DisplayName: workstationConfig.GetDisplayName(),
	}, nil
}

func (c *Client) CreateWorkstation(ctx context.Context, opts *WorkstationOpts) (*Workstation, error) {
	client, err := c.newClient(ctx)
	if err != nil {
		return nil, err
	}

	op, err := client.CreateWorkstation(ctx, &workstationspb.CreateWorkstationRequest{
		Parent:        c.WorkstationParent(opts.ConfigName),
		WorkstationId: opts.Slug,
		Workstation: &workstationspb.Workstation{
			Name:        opts.Slug,
			DisplayName: opts.DisplayName,
			Labels:      opts.Labels,
		},
	})
	if err != nil {
		return nil, err
	}

	workstation, err := op.Wait(ctx)
	if err != nil {
		return nil, err
	}

	return &Workstation{
		Name: workstation.Name,
	}, nil
}

func (c *Client) UpdateWorkstationConfig(ctx context.Context, opts *WorkstationConfigUpdateOpts) (*WorkstationConfig, error) {
	client, err := c.newClient(ctx)
	if err != nil {
		return nil, err
	}

	op, err := client.UpdateWorkstationConfig(ctx, &workstationspb.UpdateWorkstationConfigRequest{
		WorkstationConfig: &workstationspb.WorkstationConfig{
			Name: c.WorkstationParent(opts.Slug),
			Host: &workstationspb.WorkstationConfig_Host{
				Config: &workstationspb.WorkstationConfig_Host_GceInstance_{
					GceInstance: &workstationspb.WorkstationConfig_Host_GceInstance{
						MachineType: opts.MachineType,
					},
				},
			},
			Container: &workstationspb.WorkstationConfig_Container{
				Image: opts.ContainerImage,
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{
				"host.config.gce_instance.machine_type",
				"container.image",
			},
		},
		ValidateOnly: false,
		AllowMissing: false,
	})
	if err != nil {
		return nil, err
	}

	workstationConfigUpdated, err := op.Wait(ctx)
	if err != nil {
		return nil, err
	}

	return &WorkstationConfig{
		Name: workstationConfigUpdated.Name,
	}, nil
}

func (c *Client) DeleteWorkstationConfig(ctx context.Context, opts *WorkstationConfigDeleteOpts) error {
	client, err := c.newClient(ctx)
	if err != nil {
		return err
	}

	op, err := client.DeleteWorkstationConfig(ctx, &workstationspb.DeleteWorkstationConfigRequest{
		Name: c.WorkstationParent(opts.Slug),
	})
	if err != nil {
		return err
	}

	_, err = op.Wait(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) newClient(ctx context.Context) (*workstations.Client, error) {
	var options []option.ClientOption

	if c.apiEndpoint != "" {
		options = append(options, option.WithEndpoint(c.apiEndpoint))
	}

	if c.disableAuth {
		options = append(options,
			option.WithoutAuthentication(),
		)
	}

	client, err := workstations.NewRESTClient(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("creating workstations client: %w", err)
	}

	return client, nil
}

func (c *Client) WorkstationConfigParent() string {
	return fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s", c.project, c.location, c.workstationClusterID)
}

func (c *Client) WorkstationParent(configName string) string {
	return fmt.Sprintf("projects/%s/locations/%s/workstationClusters/%s/workstationConfigs/%s", c.project, c.location, c.workstationClusterID, configName)
}

func New(project, location, workstationClusterID, apiEndpoint string, disableAuth bool) *Client {
	return &Client{
		project:              project,
		location:             location,
		workstationClusterID: workstationClusterID,
		apiEndpoint:          apiEndpoint,
		disableAuth:          disableAuth,
	}
}
