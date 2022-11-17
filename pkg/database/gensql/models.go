// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package gensql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

type AccessRequestStatusType string

const (
	AccessRequestStatusTypePending  AccessRequestStatusType = "pending"
	AccessRequestStatusTypeApproved AccessRequestStatusType = "approved"
	AccessRequestStatusTypeDenied   AccessRequestStatusType = "denied"
)

func (e *AccessRequestStatusType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccessRequestStatusType(s)
	case string:
		*e = AccessRequestStatusType(s)
	default:
		return fmt.Errorf("unsupported scan type for AccessRequestStatusType: %T", src)
	}
	return nil
}

type NullAccessRequestStatusType struct {
	AccessRequestStatusType AccessRequestStatusType
	Valid                   bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccessRequestStatusType) Scan(value interface{}) error {
	if value == nil {
		ns.AccessRequestStatusType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AccessRequestStatusType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccessRequestStatusType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.AccessRequestStatusType, nil
}

type DatasourceType string

const (
	DatasourceTypeBigquery DatasourceType = "bigquery"
)

func (e *DatasourceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DatasourceType(s)
	case string:
		*e = DatasourceType(s)
	default:
		return fmt.Errorf("unsupported scan type for DatasourceType: %T", src)
	}
	return nil
}

type NullDatasourceType struct {
	DatasourceType DatasourceType
	Valid          bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullDatasourceType) Scan(value interface{}) error {
	if value == nil {
		ns.DatasourceType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.DatasourceType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullDatasourceType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.DatasourceType, nil
}

type PiiLevel string

const (
	PiiLevelSensitive  PiiLevel = "sensitive"
	PiiLevelAnonymised PiiLevel = "anonymised"
	PiiLevelNone       PiiLevel = "none"
)

func (e *PiiLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PiiLevel(s)
	case string:
		*e = PiiLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for PiiLevel: %T", src)
	}
	return nil
}

type NullPiiLevel struct {
	PiiLevel PiiLevel
	Valid    bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPiiLevel) Scan(value interface{}) error {
	if value == nil {
		ns.PiiLevel, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PiiLevel.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPiiLevel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.PiiLevel, nil
}

type StoryViewType string

const (
	StoryViewTypeMarkdown StoryViewType = "markdown"
	StoryViewTypeHeader   StoryViewType = "header"
	StoryViewTypePlotly   StoryViewType = "plotly"
	StoryViewTypeVega     StoryViewType = "vega"
)

func (e *StoryViewType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = StoryViewType(s)
	case string:
		*e = StoryViewType(s)
	default:
		return fmt.Errorf("unsupported scan type for StoryViewType: %T", src)
	}
	return nil
}

type NullStoryViewType struct {
	StoryViewType StoryViewType
	Valid         bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStoryViewType) Scan(value interface{}) error {
	if value == nil {
		ns.StoryViewType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.StoryViewType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStoryViewType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.StoryViewType, nil
}

type Dashboard struct {
	ID  string
	Url string
}

type Dataproduct struct {
	ID               uuid.UUID
	Name             string
	Description      sql.NullString
	Group            string
	Created          time.Time
	LastModified     time.Time
	TsvDocument      interface{}
	Slug             string
	TeamkatalogenUrl sql.NullString
	TeamContact      sql.NullString
	ProductAreaID    sql.NullString
	TeamID           sql.NullString
}

type Dataset struct {
	ID                       uuid.UUID
	Name                     string
	Description              sql.NullString
	Pii                      PiiLevel
	Created                  time.Time
	LastModified             time.Time
	Type                     DatasourceType
	TsvDocument              interface{}
	Slug                     string
	Repo                     sql.NullString
	Keywords                 []string
	DataproductID            uuid.UUID
	AnonymisationDescription sql.NullString
}

type DatasetAccess struct {
	ID              uuid.UUID
	DatasetID       uuid.UUID
	Subject         string
	Granter         string
	Expires         sql.NullTime
	Created         time.Time
	Revoked         sql.NullTime
	AccessRequestID uuid.NullUUID
}

type DatasetAccessRequest struct {
	ID                   uuid.UUID
	DatasetID            uuid.UUID
	Subject              string
	Owner                string
	PollyDocumentationID uuid.NullUUID
	LastModified         time.Time
	Created              time.Time
	Expires              sql.NullTime
	Status               AccessRequestStatusType
	Closed               sql.NullTime
	Granter              sql.NullString
	Reason               sql.NullString
}

type DatasetRequester struct {
	Subject   string
	DatasetID uuid.UUID
}

type DatasourceBigquery struct {
	DatasetID    uuid.UUID
	ProjectID    string
	Dataset      string
	TableName    string
	Schema       pqtype.NullRawMessage
	LastModified time.Time
	Created      time.Time
	Expires      sql.NullTime
	TableType    string
	Description  sql.NullString
	PiiTags      pqtype.NullRawMessage
}

type MetabaseMetadatum struct {
	DatabaseID           int32
	PermissionGroupID    sql.NullInt32
	SaEmail              string
	CollectionID         sql.NullInt32
	DeletedAt            sql.NullTime
	DatasetID            uuid.UUID
	AadPremissionGroupID sql.NullInt32
}

type PollyDocumentation struct {
	ID         uuid.UUID
	ExternalID string
	Name       string
	Url        string
}

type Quarto struct {
	ID           uuid.UUID
	Owner        string
	Created      time.Time
	LastModified time.Time
	Keywords     []string
	Content      string
}

type Search struct {
	ElementID    uuid.UUID
	ElementType  interface{}
	Description  string
	Keywords     interface{}
	Group        string
	Created      time.Time
	LastModified time.Time
	TsvDocument  interface{}
	Services     interface{}
}

type Session struct {
	Token       string
	AccessToken string
	Email       string
	Name        string
	Created     time.Time
	Expires     time.Time
}

type Story struct {
	ID               uuid.UUID
	Name             string
	Created          time.Time
	LastModified     time.Time
	Group            string
	Description      sql.NullString
	Keywords         []string
	TeamkatalogenUrl sql.NullString
	ProductAreaID    sql.NullString
	TeamID           sql.NullString
}

type StoryDraft struct {
	ID      uuid.UUID
	Name    string
	Created time.Time
}

type StoryToken struct {
	ID      uuid.UUID
	StoryID uuid.UUID
	Token   uuid.UUID
}

type StoryView struct {
	ID      uuid.UUID
	StoryID uuid.UUID
	Sort    int32
	Type    StoryViewType
	Spec    json.RawMessage
}

type StoryViewDraft struct {
	ID      uuid.UUID
	StoryID uuid.UUID
	Sort    int32
	Type    StoryViewType
	Spec    json.RawMessage
}

type Tag struct {
	ID     uuid.UUID
	Phrase string
}

type TeamProject struct {
	Team    string
	Project string
}

type ThirdPartyMapping struct {
	Services  []string
	DatasetID uuid.UUID
}
