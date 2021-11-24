import { gql } from '@apollo/client'
import * as Apollo from '@apollo/client'
export type Maybe<T> = T | null
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K]
}
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>
}
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>
}
const defaultOptions = {}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string
  String: string
  Boolean: boolean
  Int: number
  Float: number
  /** Time is a string in [RFC 3339](https://rfc-editor.org/rfc/rfc3339.html) format, with sub-second precision added if present. */
  Time: any
}

/** Access contains metadata on an access entry. */
export type Access = {
  __typename?: 'Access'
  /** created is timestamp for when access was created */
  created: Scalars['Time']
  /** expires is timestamp for when access expires */
  expires?: Maybe<Scalars['Time']>
  /** name of the granter */
  granter: Scalars['String']
  /** id for the access entry */
  id: Scalars['ID']
  /** revoked is timestamp for when access was revoked */
  revoked?: Maybe<Scalars['Time']>
  /** subject to grant access */
  subject: Scalars['String']
}

/** BigQuery contains metadata on a BigQuery table. */
export type BigQuery = {
  __typename?: 'BigQuery'
  /** created is when the table was created */
  created: Scalars['Time']
  /** dataset is the dataset that contains the BigQuery table */
  dataset: Scalars['String']
  /** expires, if set, is when the table expires */
  expires?: Maybe<Scalars['Time']>
  /** lastModified is the time when the table was last modified */
  lastModified: Scalars['Time']
  /** projectID is the GCP project ID that contains the BigQuery table */
  projectID: Scalars['String']
  /** schema for the BigQuery table */
  schema: Array<TableColumn>
  /** table name for BigQuery table */
  table: Scalars['String']
  /** tableType is what type the table is */
  tableType: BigQueryType
}

/** BigQueryTable contains information about a BigQuery table. */
export type BigQueryTable = {
  __typename?: 'BigQueryTable'
  /** description defined on the bigquery table. */
  description: Scalars['String']
  /** lastModified defines the last modified time of the BigQuery metadata. */
  lastModified: Scalars['Time']
  /** name of the BigQuery table. */
  name: Scalars['String']
  /** type of the BigQuery table. */
  type: BigQueryType
}

/** BigQueryType defines supported table types in a BigQuery set. */
export enum BigQueryType {
  /** materialized_view is when the table is a BigQuery materialized view. */
  MaterializedView = 'materialized_view',
  /** table is when the table is a regular BigQuery table. */
  Table = 'table',
  /** view is when the table is a BigQuery view. */
  View = 'view',
}

/** Collection contains metadata about a collection of elements. */
export type Collection = {
  __typename?: 'Collection'
  /** created contains the timestamp when the collection was created. */
  created: Scalars['Time']
  /** description of the collection. */
  description?: Maybe<Scalars['String']>
  /** elements of the collection. */
  elements: Array<CollectionElement>
  /** id is the identifier for the collection. */
  id: Scalars['ID']
  /** keywords for the collection used as tags. */
  keywords: Array<Scalars['String']>
  /** lastModified contains the timestamp when the collection was last modified. */
  lastModified: Scalars['Time']
  /** name of the collection. */
  name: Scalars['String']
  /** owner of the collection. Changes to the collection can only be done by a member of the owner. */
  owner: Owner
}

/** CollectionElement defines all types that can be returned as a collection element. */
export type CollectionElement = Dataproduct

/** CollectionElementType defines all possible types that can be stored as a collection element. */
export enum CollectionElementType {
  Dataproduct = 'dataproduct',
}

/** Dataproduct contains metadata on a datasource. */
export type Dataproduct = {
  __typename?: 'Dataproduct'
  /** access contains list of users, groups and service accounts which have access to the dataproduct */
  access: Array<Access>
  /** access collections that dataproduct is part of */
  collections: Array<Collection>
  /** created is the timestamp for when the dataproduct was created */
  created: Scalars['Time']
  /** datasource contains metadata on the datasource */
  datasource: Datasource
  /** description of the dataproduct */
  description?: Maybe<Scalars['String']>
  /** id is the identifier for the dataproduct */
  id: Scalars['ID']
  /** keywords for the dataproduct used as tags. */
  keywords: Array<Scalars['String']>
  /** lastModified is the timestamp for when the dataproduct was last modified */
  lastModified: Scalars['Time']
  /** name of the dataproduct */
  name: Scalars['String']
  /** owner of the dataproduct. Changes to the dataproduct can only be done by a member of the owner. */
  owner: Owner
  /** pii indicates whether it is personal identifiable information in the dataproduct */
  pii: Scalars['Boolean']
  /** repo is the url of the repository containing the code to create the dataproduct */
  repo?: Maybe<Scalars['String']>
  /** requesters contains list of users, groups and service accounts which can request access to the dataproduct */
  requesters: Array<Scalars['String']>
}

/** Dataproduct contains metadata on a datasource. */
export type DataproductCollectionsArgs = {
  limit?: Maybe<Scalars['Int']>
  offset?: Maybe<Scalars['Int']>
}

/** Datasource defines types that can be returned as a dataproduct datasource. */
export type Datasource = BigQuery

/** GCPProject contains metadata on a GCP project. */
export type GcpProject = {
  __typename?: 'GCPProject'
  /** group is owner group of GCP project */
  group: Group
  /** id of GCP project */
  id: Scalars['String']
}

/** Group contains metadata on a GCP group */
export type Group = {
  __typename?: 'Group'
  /** email of the group */
  email: Scalars['String']
  /** name of the group */
  name: Scalars['String']
}

export type Mutation = {
  __typename?: 'Mutation'
  /**
   * addRequesterToDataproduct adds a requester to the dataproduct.
   *
   * Requires authentication.
   */
  addRequesterToDataproduct: Scalars['Boolean']
  /**
   *    addToCollection adds a new element to the collection.
   *
   * Requires authentication.
   */
  addToCollection: Scalars['Boolean']
  /**
   *    createCollection creates a new collection.
   *
   * Requires authentication.
   */
  createCollection: Collection
  /**
   * createDataproduct creates a new dataproduct
   *
   * Requires authentication.
   */
  createDataproduct: Dataproduct
  /**
   *    deleteCollection deletes a collection.
   *
   * Requires authentication.
   */
  deleteCollection: Scalars['Boolean']
  /**
   * deleteDataproduct deletes a dataproduct.
   *
   * Requires authentication.
   */
  deleteDataproduct: Scalars['Boolean']
  /** This mutation doesn't do anything. */
  dummy?: Maybe<Scalars['String']>
  /**
   * grantAccessToDataproduct grants access for a subject to the dataproduct.
   *
   * Requires authentication.
   */
  grantAccessToDataproduct: Access
  /**
   *    removeFromCollection removes a collection.
   *
   * Requires authentication.
   */
  removeFromCollection: Scalars['Boolean']
  /**
   * removeRequesterFromDataproduct removes a requester from the dataproduct.
   *
   * Requires authentication.
   */
  removeRequesterFromDataproduct: Scalars['Boolean']
  /**
   * revokeAccessToDataproduct revokes access for a subject to the dataproduct.
   *
   * Requires authentication.
   */
  revokeAccessToDataproduct: Scalars['Boolean']
  /**
   *    updateCollection updates the metadata of a collection.
   *
   * Requires authentication.
   */
  updateCollection: Collection
  /**
   * updateDataproduct updates an existing dataproduct
   *
   * Requires authentication.
   */
  updateDataproduct: Dataproduct
}

export type MutationAddRequesterToDataproductArgs = {
  dataproductID: Scalars['ID']
  subject: Scalars['String']
}

export type MutationAddToCollectionArgs = {
  elementID: Scalars['ID']
  elementType: CollectionElementType
  id: Scalars['ID']
}

export type MutationCreateCollectionArgs = {
  input: NewCollection
}

export type MutationCreateDataproductArgs = {
  input: NewDataproduct
}

export type MutationDeleteCollectionArgs = {
  id: Scalars['ID']
}

export type MutationDeleteDataproductArgs = {
  id: Scalars['ID']
}

export type MutationDummyArgs = {
  no?: Maybe<Scalars['String']>
}

export type MutationGrantAccessToDataproductArgs = {
  dataproductID: Scalars['ID']
  expires?: Maybe<Scalars['Time']>
  subject?: Maybe<Scalars['String']>
  subjectType?: Maybe<SubjectType>
}

export type MutationRemoveFromCollectionArgs = {
  elementID: Scalars['ID']
  elementType: CollectionElementType
  id: Scalars['ID']
}

export type MutationRemoveRequesterFromDataproductArgs = {
  dataproductID: Scalars['ID']
  subject: Scalars['String']
}

export type MutationRevokeAccessToDataproductArgs = {
  id: Scalars['ID']
}

export type MutationUpdateCollectionArgs = {
  id: Scalars['ID']
  input: UpdateCollection
}

export type MutationUpdateDataproductArgs = {
  id: Scalars['ID']
  input: UpdateDataproduct
}

/** NewBigQuery contains metadata for creating a new bigquery data source */
export type NewBigQuery = {
  /** dataset is the name of the dataset. */
  dataset: Scalars['String']
  /** projectID is the GCP project ID that contains the dataset. */
  projectID: Scalars['String']
  /** table is the name of the table */
  table: Scalars['String']
}

/**
 * NewCollection contains metadata for creating a new collection.
 * Group must match one of the groups the authenticated user is part of.
 */
export type NewCollection = {
  /** description of the colection. */
  description?: Maybe<Scalars['String']>
  /** group the collection belongs to. Used for authorization. */
  group: Scalars['String']
  /** keywords for the collection used as tags. */
  keywords?: Maybe<Array<Scalars['String']>>
  /** name of the collection. */
  name: Scalars['String']
  /** owner Teamkatalogen URL for the collection. */
  teamkatalogenURL?: Maybe<Scalars['String']>
}

/** NewDataproduct contains metadata for creating a new dataproduct */
export type NewDataproduct = {
  /** bigquery contains metadata for the bigquery datasource added to the dataproduct. */
  bigquery: NewBigQuery
  /** description of the dataproduct */
  description?: Maybe<Scalars['String']>
  /** owner group email for the dataproduct. */
  group: Scalars['String']
  /** keywords for the dataproduct used as tags. */
  keywords?: Maybe<Array<Scalars['String']>>
  /** name of dataproduct */
  name: Scalars['String']
  /** pii indicates whether it is personal identifiable information in the dataproduct */
  pii: Scalars['Boolean']
  /** repo is the url of the repository containing the code to create the dataproduct */
  repo?: Maybe<Scalars['String']>
  /** requesters contains list of users, groups and service accounts which can request access to the dataproduct */
  requesters?: Maybe<Array<Scalars['String']>>
  /** owner Teamkatalogen URL for the dataproduct. */
  teamkatalogenURL?: Maybe<Scalars['String']>
}

/** Owner contains metadata on the owner of the dataproduct. */
export type Owner = {
  __typename?: 'Owner'
  /** owner group is the email for the group. */
  group: Scalars['String']
  /** teamkatalogenURL is url for the team in the NAV team catalog. */
  teamkatalogenURL?: Maybe<Scalars['String']>
}

export type Query = {
  __typename?: 'Query'
  /** collection returns the given collection. */
  collection: Collection
  /** collections returns a list of collections. Pagination done using the arguments. */
  collections: Array<Collection>
  /** dataproduct returns the given dataproduct. */
  dataproduct: Dataproduct
  /** dataproducts returns a list of dataproducts. Pagination done using the arguments. */
  dataproducts: Array<Dataproduct>
  /**
   * gcpGetDatasets returns all datasets for a given project.
   *
   * Requires authentication.
   */
  gcpGetDatasets: Array<Scalars['String']>
  /**
   * gcpGetTables returns all tables for a given dataset.
   *
   * Requires authentication.
   */
  gcpGetTables: Array<BigQueryTable>
  /** search through existing dataproducts and collections. */
  search: Array<SearchResultRow>
  /** searches teamkatalogen for teams where team name matches query input */
  teamkatalogen: Array<TeamkatalogenResult>
  /** userInfo returns information about the logged in user. */
  userInfo: UserInfo
  /** version returns the API version. */
  version: Scalars['String']
}

export type QueryCollectionArgs = {
  id: Scalars['ID']
}

export type QueryCollectionsArgs = {
  limit?: Maybe<Scalars['Int']>
  offset?: Maybe<Scalars['Int']>
}

export type QueryDataproductArgs = {
  id: Scalars['ID']
}

export type QueryDataproductsArgs = {
  limit?: Maybe<Scalars['Int']>
  offset?: Maybe<Scalars['Int']>
}

export type QueryGcpGetDatasetsArgs = {
  projectID: Scalars['String']
}

export type QueryGcpGetTablesArgs = {
  datasetID: Scalars['String']
  projectID: Scalars['String']
}

export type QuerySearchArgs = {
  q?: Maybe<SearchQuery>
}

export type QueryTeamkatalogenArgs = {
  q: Scalars['String']
}

export type SearchQuery = {
  /** keyword filters results on the keyword. */
  keyword?: Maybe<Scalars['String']>
  /** limit the number of returned search results. */
  limit?: Maybe<Scalars['Int']>
  /** offset the list of returned search results. Used as pagination with PAGE-INDEX * limit. */
  offset?: Maybe<Scalars['Int']>
  /**
   * text is used as freetext search.
   *
   * Use " to identify phrases. Example: "hello world"
   *
   * Use - to exclude words. Example "include this -exclude -those"
   *
   * Use OR as a keyword for the OR operator. Example "night OR day"
   */
  text?: Maybe<Scalars['String']>
}

export type SearchResult = Collection | Dataproduct

export type SearchResultRow = {
  __typename?: 'SearchResultRow'
  excerpt: Scalars['String']
  result: SearchResult
}

/** SubjectType defines all possible types that can request access to a dataproduct. */
export enum SubjectType {
  Group = 'group',
  ServiceAccount = 'serviceAccount',
  User = 'user',
}

/** TableColumn contains metadata on a BigQuery table column. */
export type TableColumn = {
  __typename?: 'TableColumn'
  /** description of column. */
  description: Scalars['String']
  /** mode of column (NULLABLE, REQUIRED or REPEATED). */
  mode: Scalars['String']
  /** name of column. */
  name: Scalars['String']
  /** type is the datatype of the column. */
  type: Scalars['String']
}

export type TeamkatalogenResult = {
  __typename?: 'TeamkatalogenResult'
  /** team description. */
  description: Scalars['String']
  /** team name. */
  name: Scalars['String']
  /** url to team in teamkatalogen. */
  url: Scalars['String']
}

/** UpdateCollection contains data for updating the metadata of a collection. */
export type UpdateCollection = {
  /** description of the collection. */
  description?: Maybe<Scalars['String']>
  /** keywords for the collection used as tags. */
  keywords?: Maybe<Array<Scalars['String']>>
  /** name of the collection. */
  name: Scalars['String']
  /** owner Teamkatalogen URL for the collection. */
  teamkatalogenURL?: Maybe<Scalars['String']>
}

/** UpdateDataproduct contains metadata for updating a dataproduct */
export type UpdateDataproduct = {
  /** description of the dataproduct */
  description?: Maybe<Scalars['String']>
  /** keywords for the dataproduct used as tags. */
  keywords?: Maybe<Array<Scalars['String']>>
  /** name of dataproduct */
  name: Scalars['String']
  /** pii indicates whether it is personal identifiable information in the dataproduct */
  pii: Scalars['Boolean']
  /** repo is the url of the repository containing the code to create the dataproduct */
  repo?: Maybe<Scalars['String']>
  /** requesters contains list of users, groups and service accounts which can request access to the dataproduct */
  requesters?: Maybe<Array<Scalars['String']>>
  /** owner Teamkatalogen URL for the dataproduct. */
  teamkatalogenURL?: Maybe<Scalars['String']>
}

/** UserInfo contains metadata on a logged in user */
export type UserInfo = {
  __typename?: 'UserInfo'
  /** accessable is a list of dataproducts which the user has explicit access to */
  accessable: Array<Dataproduct>
  /** collections is a list of collections with one of the users groups as owner */
  collections: Array<Collection>
  /** dataproducts is a list of dataproducts with one of the users groups as owner */
  dataproducts: Array<Dataproduct>
  /** email of user. */
  email: Scalars['String']
  /** gcpProjects is GCP projects the user is a member of. */
  gcpProjects: Array<GcpProject>
  /** groups the user is a member of. */
  groups: Array<Group>
  /** name of user. */
  name: Scalars['String']
}

export type AddRequesterMutationVariables = Exact<{
  dataproductID: Scalars['ID']
  subject: Scalars['String']
}>

export type AddRequesterMutation = {
  __typename?: 'Mutation'
  addRequesterToDataproduct: boolean
}

export type DataproductAccessQueryVariables = Exact<{
  id: Scalars['ID']
}>

export type DataproductAccessQuery = {
  __typename?: 'Query'
  dataproduct: {
    __typename?: 'Dataproduct'
    id: string
    name: string
    requesters: Array<string>
    owner: { __typename?: 'Owner'; group: string }
    access: Array<{
      __typename?: 'Access'
      id: string
      subject: string
      granter: string
      expires?: any | null | undefined
      created: any
      revoked?: any | null | undefined
    }>
  }
}

export type DataproductRequestersQueryVariables = Exact<{
  id: Scalars['ID']
}>

export type DataproductRequestersQuery = {
  __typename?: 'Query'
  dataproduct: { __typename?: 'Dataproduct'; requesters: Array<string> }
}

export type GrantAccessMutationVariables = Exact<{
  dataproductID: Scalars['ID']
  subject: Scalars['String']
  subjectType: SubjectType
  expires?: Maybe<Scalars['Time']>
}>

export type GrantAccessMutation = {
  __typename?: 'Mutation'
  grantAccessToDataproduct: { __typename?: 'Access'; id: string }
}

export type RemoveRequesterMutationVariables = Exact<{
  dataproductID: Scalars['ID']
  subject: Scalars['String']
}>

export type RemoveRequesterMutation = {
  __typename?: 'Mutation'
  removeRequesterFromDataproduct: boolean
}

export type RevokeAccessMutationVariables = Exact<{
  id: Scalars['ID']
}>

export type RevokeAccessMutation = {
  __typename?: 'Mutation'
  revokeAccessToDataproduct: boolean
}

export type AddToCollectionMutationVariables = Exact<{
  id: Scalars['ID']
  elementID: Scalars['ID']
  elementType: CollectionElementType
}>

export type AddToCollectionMutation = {
  __typename?: 'Mutation'
  addToCollection: boolean
}

export type CollectionQueryVariables = Exact<{
  id: Scalars['ID']
}>

export type CollectionQuery = {
  __typename?: 'Query'
  collection: {
    __typename?: 'Collection'
    id: string
    name: string
    description?: string | null | undefined
    created: any
    keywords: Array<string>
    lastModified: any
    owner: { __typename?: 'Owner'; group: string }
    elements: Array<{
      __typename: 'Dataproduct'
      id: string
      name: string
      description?: string | null | undefined
    }>
  }
}

export type CreateCollectionMutationVariables = Exact<{
  input: NewCollection
}>

export type CreateCollectionMutation = {
  __typename?: 'Mutation'
  createCollection: { __typename?: 'Collection'; id: string }
}

export type DeleteCollectionMutationVariables = Exact<{
  id: Scalars['ID']
}>

export type DeleteCollectionMutation = {
  __typename?: 'Mutation'
  deleteCollection: boolean
}

export type RemoveFromCollectionMutationVariables = Exact<{
  id: Scalars['ID']
  elementID: Scalars['ID']
  elementType: CollectionElementType
}>

export type RemoveFromCollectionMutation = {
  __typename?: 'Mutation'
  removeFromCollection: boolean
}

export type UpdateCollectionMutationVariables = Exact<{
  id: Scalars['ID']
  input: UpdateCollection
}>

export type UpdateCollectionMutation = {
  __typename?: 'Mutation'
  updateCollection: { __typename?: 'Collection'; id: string }
}

export type CreateDataproductMutationVariables = Exact<{
  input: NewDataproduct
}>

export type CreateDataproductMutation = {
  __typename?: 'Mutation'
  createDataproduct: { __typename?: 'Dataproduct'; id: string }
}

export type DataproductQueryVariables = Exact<{
  id: Scalars['ID']
}>

export type DataproductQuery = {
  __typename?: 'Query'
  dataproduct: {
    __typename?: 'Dataproduct'
    id: string
    lastModified: any
    name: string
    description?: string | null | undefined
    created: any
    repo?: string | null | undefined
    pii: boolean
    keywords: Array<string>
    collections: Array<{ __typename: 'Collection'; id: string; name: string }>
    owner: {
      __typename?: 'Owner'
      group: string
      teamkatalogenURL?: string | null | undefined
    }
    datasource: {
      __typename?: 'BigQuery'
      projectID: string
      dataset: string
      table: string
      lastModified: any
      created: any
      expires?: any | null | undefined
      tableType: BigQueryType
      type: 'BigQuery'
      schema: Array<{
        __typename?: 'TableColumn'
        name: string
        description: string
        mode: string
        type: string
      }>
    }
  }
}

export type DataproductSummaryQueryVariables = Exact<{
  id: Scalars['ID']
}>

export type DataproductSummaryQuery = {
  __typename?: 'Query'
  dataproduct: {
    __typename?: 'Dataproduct'
    id: string
    lastModified: any
    name: string
    description?: string | null | undefined
    created: any
    pii: boolean
    keywords: Array<string>
    datasource: { __typename?: 'BigQuery'; type: 'BigQuery' }
  }
}

export type DeleteDataproductMutationVariables = Exact<{
  id: Scalars['ID']
}>

export type DeleteDataproductMutation = {
  __typename?: 'Mutation'
  deleteDataproduct: boolean
}

export type GcpGetDatasetsQueryVariables = Exact<{
  projectID: Scalars['String']
}>

export type GcpGetDatasetsQuery = {
  __typename?: 'Query'
  gcpGetDatasets: Array<string>
}

export type GcpGetTablesQueryVariables = Exact<{
  projectID: Scalars['String']
  datasetID: Scalars['String']
}>

export type GcpGetTablesQuery = {
  __typename?: 'Query'
  gcpGetTables: Array<{
    __typename?: 'BigQueryTable'
    name: string
    type: BigQueryType
    description: string
  }>
}

export type UpdateDataproductMutationVariables = Exact<{
  id: Scalars['ID']
  input: UpdateDataproduct
}>

export type UpdateDataproductMutation = {
  __typename?: 'Mutation'
  updateDataproduct: { __typename?: 'Dataproduct'; id: string }
}

export type SearchContentQueryVariables = Exact<{
  q: SearchQuery
}>

export type SearchContentQuery = {
  __typename?: 'Query'
  search: Array<{
    __typename?: 'SearchResultRow'
    excerpt: string
    result:
      | {
          __typename: 'Collection'
          id: string
          name: string
          description?: string | null | undefined
          created: any
          lastModified: any
          owner: {
            __typename?: 'Owner'
            group: string
            teamkatalogenURL?: string | null | undefined
          }
        }
      | {
          __typename: 'Dataproduct'
          id: string
          name: string
          description?: string | null | undefined
          created: any
          lastModified: any
          owner: {
            __typename?: 'Owner'
            group: string
            teamkatalogenURL?: string | null | undefined
          }
        }
  }>
}

export type TeamkatalogenQueryVariables = Exact<{
  q: Scalars['String']
}>

export type TeamkatalogenQuery = {
  __typename?: 'Query'
  teamkatalogen: Array<{
    __typename?: 'TeamkatalogenResult'
    name: string
    url: string
  }>
}

export type UserInfoDetailsQueryVariables = Exact<{ [key: string]: never }>

export type UserInfoDetailsQuery = {
  __typename?: 'Query'
  userInfo: {
    __typename?: 'UserInfo'
    name: string
    email: string
    groups: Array<{ __typename?: 'Group'; name: string; email: string }>
    gcpProjects: Array<{
      __typename?: 'GCPProject'
      id: string
      group: { __typename?: 'Group'; name: string; email: string }
    }>
  }
}

export type UserInfoAccessableDataproductQueryVariables = Exact<{
  [key: string]: never
}>

export type UserInfoAccessableDataproductQuery = {
  __typename?: 'Query'
  userInfo: {
    __typename?: 'UserInfo'
    accessable: Array<{
      __typename: 'Dataproduct'
      id: string
      name: string
      description?: string | null | undefined
      created: any
      lastModified: any
      owner: {
        __typename?: 'Owner'
        group: string
        teamkatalogenURL?: string | null | undefined
      }
    }>
  }
}

export type UserInfoUserProductsQueryVariables = Exact<{ [key: string]: never }>

export type UserInfoUserProductsQuery = {
  __typename?: 'Query'
  userInfo: {
    __typename?: 'UserInfo'
    dataproducts: Array<{
      __typename: 'Dataproduct'
      id: string
      name: string
      description?: string | null | undefined
      created: any
      lastModified: any
      owner: {
        __typename?: 'Owner'
        group: string
        teamkatalogenURL?: string | null | undefined
      }
    }>
    collections: Array<{
      __typename: 'Collection'
      id: string
      name: string
      description?: string | null | undefined
      created: any
      lastModified: any
      owner: {
        __typename?: 'Owner'
        group: string
        teamkatalogenURL?: string | null | undefined
      }
    }>
  }
}

export const AddRequesterDocument = gql`
  mutation AddRequester($dataproductID: ID!, $subject: String!) {
    addRequesterToDataproduct(dataproductID: $dataproductID, subject: $subject)
  }
`
export type AddRequesterMutationFn = Apollo.MutationFunction<
  AddRequesterMutation,
  AddRequesterMutationVariables
>

/**
 * __useAddRequesterMutation__
 *
 * To run a mutation, you first call `useAddRequesterMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddRequesterMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addRequesterMutation, { data, loading, error }] = useAddRequesterMutation({
 *   variables: {
 *      dataproductID: // value for 'dataproductID'
 *      subject: // value for 'subject'
 *   },
 * });
 */
export function useAddRequesterMutation(
  baseOptions?: Apollo.MutationHookOptions<
    AddRequesterMutation,
    AddRequesterMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    AddRequesterMutation,
    AddRequesterMutationVariables
  >(AddRequesterDocument, options)
}
export type AddRequesterMutationHookResult = ReturnType<
  typeof useAddRequesterMutation
>
export type AddRequesterMutationResult =
  Apollo.MutationResult<AddRequesterMutation>
export type AddRequesterMutationOptions = Apollo.BaseMutationOptions<
  AddRequesterMutation,
  AddRequesterMutationVariables
>
export const DataproductAccessDocument = gql`
  query DataproductAccess($id: ID!) {
    dataproduct(id: $id) {
      id
      name
      owner {
        group
      }
      access {
        id
        subject
        granter
        expires
        created
        revoked
      }
      requesters
    }
  }
`

/**
 * __useDataproductAccessQuery__
 *
 * To run a query within a React component, call `useDataproductAccessQuery` and pass it any options that fit your needs.
 * When your component renders, `useDataproductAccessQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useDataproductAccessQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDataproductAccessQuery(
  baseOptions: Apollo.QueryHookOptions<
    DataproductAccessQuery,
    DataproductAccessQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<
    DataproductAccessQuery,
    DataproductAccessQueryVariables
  >(DataproductAccessDocument, options)
}
export function useDataproductAccessLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    DataproductAccessQuery,
    DataproductAccessQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<
    DataproductAccessQuery,
    DataproductAccessQueryVariables
  >(DataproductAccessDocument, options)
}
export type DataproductAccessQueryHookResult = ReturnType<
  typeof useDataproductAccessQuery
>
export type DataproductAccessLazyQueryHookResult = ReturnType<
  typeof useDataproductAccessLazyQuery
>
export type DataproductAccessQueryResult = Apollo.QueryResult<
  DataproductAccessQuery,
  DataproductAccessQueryVariables
>
export const DataproductRequestersDocument = gql`
  query DataproductRequesters($id: ID!) {
    dataproduct(id: $id) {
      requesters
    }
  }
`

/**
 * __useDataproductRequestersQuery__
 *
 * To run a query within a React component, call `useDataproductRequestersQuery` and pass it any options that fit your needs.
 * When your component renders, `useDataproductRequestersQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useDataproductRequestersQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDataproductRequestersQuery(
  baseOptions: Apollo.QueryHookOptions<
    DataproductRequestersQuery,
    DataproductRequestersQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<
    DataproductRequestersQuery,
    DataproductRequestersQueryVariables
  >(DataproductRequestersDocument, options)
}
export function useDataproductRequestersLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    DataproductRequestersQuery,
    DataproductRequestersQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<
    DataproductRequestersQuery,
    DataproductRequestersQueryVariables
  >(DataproductRequestersDocument, options)
}
export type DataproductRequestersQueryHookResult = ReturnType<
  typeof useDataproductRequestersQuery
>
export type DataproductRequestersLazyQueryHookResult = ReturnType<
  typeof useDataproductRequestersLazyQuery
>
export type DataproductRequestersQueryResult = Apollo.QueryResult<
  DataproductRequestersQuery,
  DataproductRequestersQueryVariables
>
export const GrantAccessDocument = gql`
  mutation GrantAccess(
    $dataproductID: ID!
    $subject: String!
    $subjectType: SubjectType!
    $expires: Time
  ) {
    grantAccessToDataproduct(
      dataproductID: $dataproductID
      subject: $subject
      subjectType: $subjectType
      expires: $expires
    ) {
      id
    }
  }
`
export type GrantAccessMutationFn = Apollo.MutationFunction<
  GrantAccessMutation,
  GrantAccessMutationVariables
>

/**
 * __useGrantAccessMutation__
 *
 * To run a mutation, you first call `useGrantAccessMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useGrantAccessMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [grantAccessMutation, { data, loading, error }] = useGrantAccessMutation({
 *   variables: {
 *      dataproductID: // value for 'dataproductID'
 *      subject: // value for 'subject'
 *      subjectType: // value for 'subjectType'
 *      expires: // value for 'expires'
 *   },
 * });
 */
export function useGrantAccessMutation(
  baseOptions?: Apollo.MutationHookOptions<
    GrantAccessMutation,
    GrantAccessMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<GrantAccessMutation, GrantAccessMutationVariables>(
    GrantAccessDocument,
    options
  )
}
export type GrantAccessMutationHookResult = ReturnType<
  typeof useGrantAccessMutation
>
export type GrantAccessMutationResult =
  Apollo.MutationResult<GrantAccessMutation>
export type GrantAccessMutationOptions = Apollo.BaseMutationOptions<
  GrantAccessMutation,
  GrantAccessMutationVariables
>
export const RemoveRequesterDocument = gql`
  mutation RemoveRequester($dataproductID: ID!, $subject: String!) {
    removeRequesterFromDataproduct(
      dataproductID: $dataproductID
      subject: $subject
    )
  }
`
export type RemoveRequesterMutationFn = Apollo.MutationFunction<
  RemoveRequesterMutation,
  RemoveRequesterMutationVariables
>

/**
 * __useRemoveRequesterMutation__
 *
 * To run a mutation, you first call `useRemoveRequesterMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useRemoveRequesterMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [removeRequesterMutation, { data, loading, error }] = useRemoveRequesterMutation({
 *   variables: {
 *      dataproductID: // value for 'dataproductID'
 *      subject: // value for 'subject'
 *   },
 * });
 */
export function useRemoveRequesterMutation(
  baseOptions?: Apollo.MutationHookOptions<
    RemoveRequesterMutation,
    RemoveRequesterMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    RemoveRequesterMutation,
    RemoveRequesterMutationVariables
  >(RemoveRequesterDocument, options)
}
export type RemoveRequesterMutationHookResult = ReturnType<
  typeof useRemoveRequesterMutation
>
export type RemoveRequesterMutationResult =
  Apollo.MutationResult<RemoveRequesterMutation>
export type RemoveRequesterMutationOptions = Apollo.BaseMutationOptions<
  RemoveRequesterMutation,
  RemoveRequesterMutationVariables
>
export const RevokeAccessDocument = gql`
  mutation RevokeAccess($id: ID!) {
    revokeAccessToDataproduct(id: $id)
  }
`
export type RevokeAccessMutationFn = Apollo.MutationFunction<
  RevokeAccessMutation,
  RevokeAccessMutationVariables
>

/**
 * __useRevokeAccessMutation__
 *
 * To run a mutation, you first call `useRevokeAccessMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useRevokeAccessMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [revokeAccessMutation, { data, loading, error }] = useRevokeAccessMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useRevokeAccessMutation(
  baseOptions?: Apollo.MutationHookOptions<
    RevokeAccessMutation,
    RevokeAccessMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    RevokeAccessMutation,
    RevokeAccessMutationVariables
  >(RevokeAccessDocument, options)
}
export type RevokeAccessMutationHookResult = ReturnType<
  typeof useRevokeAccessMutation
>
export type RevokeAccessMutationResult =
  Apollo.MutationResult<RevokeAccessMutation>
export type RevokeAccessMutationOptions = Apollo.BaseMutationOptions<
  RevokeAccessMutation,
  RevokeAccessMutationVariables
>
export const AddToCollectionDocument = gql`
  mutation addToCollection(
    $id: ID!
    $elementID: ID!
    $elementType: CollectionElementType!
  ) {
    addToCollection(id: $id, elementID: $elementID, elementType: $elementType)
  }
`
export type AddToCollectionMutationFn = Apollo.MutationFunction<
  AddToCollectionMutation,
  AddToCollectionMutationVariables
>

/**
 * __useAddToCollectionMutation__
 *
 * To run a mutation, you first call `useAddToCollectionMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddToCollectionMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addToCollectionMutation, { data, loading, error }] = useAddToCollectionMutation({
 *   variables: {
 *      id: // value for 'id'
 *      elementID: // value for 'elementID'
 *      elementType: // value for 'elementType'
 *   },
 * });
 */
export function useAddToCollectionMutation(
  baseOptions?: Apollo.MutationHookOptions<
    AddToCollectionMutation,
    AddToCollectionMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    AddToCollectionMutation,
    AddToCollectionMutationVariables
  >(AddToCollectionDocument, options)
}
export type AddToCollectionMutationHookResult = ReturnType<
  typeof useAddToCollectionMutation
>
export type AddToCollectionMutationResult =
  Apollo.MutationResult<AddToCollectionMutation>
export type AddToCollectionMutationOptions = Apollo.BaseMutationOptions<
  AddToCollectionMutation,
  AddToCollectionMutationVariables
>
export const CollectionDocument = gql`
  query Collection($id: ID!) {
    collection(id: $id) {
      id
      name
      description
      created
      keywords
      lastModified
      owner {
        group
      }
      elements {
        ... on Dataproduct {
          __typename
          id
          name
          description
        }
      }
    }
  }
`

/**
 * __useCollectionQuery__
 *
 * To run a query within a React component, call `useCollectionQuery` and pass it any options that fit your needs.
 * When your component renders, `useCollectionQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCollectionQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useCollectionQuery(
  baseOptions: Apollo.QueryHookOptions<
    CollectionQuery,
    CollectionQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<CollectionQuery, CollectionQueryVariables>(
    CollectionDocument,
    options
  )
}
export function useCollectionLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    CollectionQuery,
    CollectionQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<CollectionQuery, CollectionQueryVariables>(
    CollectionDocument,
    options
  )
}
export type CollectionQueryHookResult = ReturnType<typeof useCollectionQuery>
export type CollectionLazyQueryHookResult = ReturnType<
  typeof useCollectionLazyQuery
>
export type CollectionQueryResult = Apollo.QueryResult<
  CollectionQuery,
  CollectionQueryVariables
>
export const CreateCollectionDocument = gql`
  mutation createCollection($input: NewCollection!) {
    createCollection(input: $input) {
      id
    }
  }
`
export type CreateCollectionMutationFn = Apollo.MutationFunction<
  CreateCollectionMutation,
  CreateCollectionMutationVariables
>

/**
 * __useCreateCollectionMutation__
 *
 * To run a mutation, you first call `useCreateCollectionMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateCollectionMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createCollectionMutation, { data, loading, error }] = useCreateCollectionMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateCollectionMutation(
  baseOptions?: Apollo.MutationHookOptions<
    CreateCollectionMutation,
    CreateCollectionMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    CreateCollectionMutation,
    CreateCollectionMutationVariables
  >(CreateCollectionDocument, options)
}
export type CreateCollectionMutationHookResult = ReturnType<
  typeof useCreateCollectionMutation
>
export type CreateCollectionMutationResult =
  Apollo.MutationResult<CreateCollectionMutation>
export type CreateCollectionMutationOptions = Apollo.BaseMutationOptions<
  CreateCollectionMutation,
  CreateCollectionMutationVariables
>
export const DeleteCollectionDocument = gql`
  mutation deleteCollection($id: ID!) {
    deleteCollection(id: $id)
  }
`
export type DeleteCollectionMutationFn = Apollo.MutationFunction<
  DeleteCollectionMutation,
  DeleteCollectionMutationVariables
>

/**
 * __useDeleteCollectionMutation__
 *
 * To run a mutation, you first call `useDeleteCollectionMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteCollectionMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteCollectionMutation, { data, loading, error }] = useDeleteCollectionMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteCollectionMutation(
  baseOptions?: Apollo.MutationHookOptions<
    DeleteCollectionMutation,
    DeleteCollectionMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    DeleteCollectionMutation,
    DeleteCollectionMutationVariables
  >(DeleteCollectionDocument, options)
}
export type DeleteCollectionMutationHookResult = ReturnType<
  typeof useDeleteCollectionMutation
>
export type DeleteCollectionMutationResult =
  Apollo.MutationResult<DeleteCollectionMutation>
export type DeleteCollectionMutationOptions = Apollo.BaseMutationOptions<
  DeleteCollectionMutation,
  DeleteCollectionMutationVariables
>
export const RemoveFromCollectionDocument = gql`
  mutation removeFromCollection(
    $id: ID!
    $elementID: ID!
    $elementType: CollectionElementType!
  ) {
    removeFromCollection(
      id: $id
      elementID: $elementID
      elementType: $elementType
    )
  }
`
export type RemoveFromCollectionMutationFn = Apollo.MutationFunction<
  RemoveFromCollectionMutation,
  RemoveFromCollectionMutationVariables
>

/**
 * __useRemoveFromCollectionMutation__
 *
 * To run a mutation, you first call `useRemoveFromCollectionMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useRemoveFromCollectionMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [removeFromCollectionMutation, { data, loading, error }] = useRemoveFromCollectionMutation({
 *   variables: {
 *      id: // value for 'id'
 *      elementID: // value for 'elementID'
 *      elementType: // value for 'elementType'
 *   },
 * });
 */
export function useRemoveFromCollectionMutation(
  baseOptions?: Apollo.MutationHookOptions<
    RemoveFromCollectionMutation,
    RemoveFromCollectionMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    RemoveFromCollectionMutation,
    RemoveFromCollectionMutationVariables
  >(RemoveFromCollectionDocument, options)
}
export type RemoveFromCollectionMutationHookResult = ReturnType<
  typeof useRemoveFromCollectionMutation
>
export type RemoveFromCollectionMutationResult =
  Apollo.MutationResult<RemoveFromCollectionMutation>
export type RemoveFromCollectionMutationOptions = Apollo.BaseMutationOptions<
  RemoveFromCollectionMutation,
  RemoveFromCollectionMutationVariables
>
export const UpdateCollectionDocument = gql`
  mutation updateCollection($id: ID!, $input: UpdateCollection!) {
    updateCollection(id: $id, input: $input) {
      id
    }
  }
`
export type UpdateCollectionMutationFn = Apollo.MutationFunction<
  UpdateCollectionMutation,
  UpdateCollectionMutationVariables
>

/**
 * __useUpdateCollectionMutation__
 *
 * To run a mutation, you first call `useUpdateCollectionMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateCollectionMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateCollectionMutation, { data, loading, error }] = useUpdateCollectionMutation({
 *   variables: {
 *      id: // value for 'id'
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateCollectionMutation(
  baseOptions?: Apollo.MutationHookOptions<
    UpdateCollectionMutation,
    UpdateCollectionMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    UpdateCollectionMutation,
    UpdateCollectionMutationVariables
  >(UpdateCollectionDocument, options)
}
export type UpdateCollectionMutationHookResult = ReturnType<
  typeof useUpdateCollectionMutation
>
export type UpdateCollectionMutationResult =
  Apollo.MutationResult<UpdateCollectionMutation>
export type UpdateCollectionMutationOptions = Apollo.BaseMutationOptions<
  UpdateCollectionMutation,
  UpdateCollectionMutationVariables
>
export const CreateDataproductDocument = gql`
  mutation createDataproduct($input: NewDataproduct!) {
    createDataproduct(input: $input) {
      id
    }
  }
`
export type CreateDataproductMutationFn = Apollo.MutationFunction<
  CreateDataproductMutation,
  CreateDataproductMutationVariables
>

/**
 * __useCreateDataproductMutation__
 *
 * To run a mutation, you first call `useCreateDataproductMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateDataproductMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createDataproductMutation, { data, loading, error }] = useCreateDataproductMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateDataproductMutation(
  baseOptions?: Apollo.MutationHookOptions<
    CreateDataproductMutation,
    CreateDataproductMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    CreateDataproductMutation,
    CreateDataproductMutationVariables
  >(CreateDataproductDocument, options)
}
export type CreateDataproductMutationHookResult = ReturnType<
  typeof useCreateDataproductMutation
>
export type CreateDataproductMutationResult =
  Apollo.MutationResult<CreateDataproductMutation>
export type CreateDataproductMutationOptions = Apollo.BaseMutationOptions<
  CreateDataproductMutation,
  CreateDataproductMutationVariables
>
export const DataproductDocument = gql`
  query Dataproduct($id: ID!) {
    dataproduct(id: $id) {
      id
      lastModified
      name
      description
      created
      repo
      pii
      keywords
      collections {
        __typename
        id
        name
      }
      owner {
        group
        teamkatalogenURL
      }
      datasource {
        type: __typename
        ... on BigQuery {
          projectID
          dataset
          table
          lastModified
          created
          expires
          tableType
          schema {
            name
            description
            mode
            type
          }
        }
      }
    }
  }
`

/**
 * __useDataproductQuery__
 *
 * To run a query within a React component, call `useDataproductQuery` and pass it any options that fit your needs.
 * When your component renders, `useDataproductQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useDataproductQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDataproductQuery(
  baseOptions: Apollo.QueryHookOptions<
    DataproductQuery,
    DataproductQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<DataproductQuery, DataproductQueryVariables>(
    DataproductDocument,
    options
  )
}
export function useDataproductLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    DataproductQuery,
    DataproductQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<DataproductQuery, DataproductQueryVariables>(
    DataproductDocument,
    options
  )
}
export type DataproductQueryHookResult = ReturnType<typeof useDataproductQuery>
export type DataproductLazyQueryHookResult = ReturnType<
  typeof useDataproductLazyQuery
>
export type DataproductQueryResult = Apollo.QueryResult<
  DataproductQuery,
  DataproductQueryVariables
>
export const DataproductSummaryDocument = gql`
  query DataproductSummary($id: ID!) {
    dataproduct(id: $id) {
      id
      lastModified
      name
      description
      created
      pii
      keywords
      datasource {
        type: __typename
      }
    }
  }
`

/**
 * __useDataproductSummaryQuery__
 *
 * To run a query within a React component, call `useDataproductSummaryQuery` and pass it any options that fit your needs.
 * When your component renders, `useDataproductSummaryQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useDataproductSummaryQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDataproductSummaryQuery(
  baseOptions: Apollo.QueryHookOptions<
    DataproductSummaryQuery,
    DataproductSummaryQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<
    DataproductSummaryQuery,
    DataproductSummaryQueryVariables
  >(DataproductSummaryDocument, options)
}
export function useDataproductSummaryLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    DataproductSummaryQuery,
    DataproductSummaryQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<
    DataproductSummaryQuery,
    DataproductSummaryQueryVariables
  >(DataproductSummaryDocument, options)
}
export type DataproductSummaryQueryHookResult = ReturnType<
  typeof useDataproductSummaryQuery
>
export type DataproductSummaryLazyQueryHookResult = ReturnType<
  typeof useDataproductSummaryLazyQuery
>
export type DataproductSummaryQueryResult = Apollo.QueryResult<
  DataproductSummaryQuery,
  DataproductSummaryQueryVariables
>
export const DeleteDataproductDocument = gql`
  mutation deleteDataproduct($id: ID!) {
    deleteDataproduct(id: $id)
  }
`
export type DeleteDataproductMutationFn = Apollo.MutationFunction<
  DeleteDataproductMutation,
  DeleteDataproductMutationVariables
>

/**
 * __useDeleteDataproductMutation__
 *
 * To run a mutation, you first call `useDeleteDataproductMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteDataproductMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteDataproductMutation, { data, loading, error }] = useDeleteDataproductMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteDataproductMutation(
  baseOptions?: Apollo.MutationHookOptions<
    DeleteDataproductMutation,
    DeleteDataproductMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    DeleteDataproductMutation,
    DeleteDataproductMutationVariables
  >(DeleteDataproductDocument, options)
}
export type DeleteDataproductMutationHookResult = ReturnType<
  typeof useDeleteDataproductMutation
>
export type DeleteDataproductMutationResult =
  Apollo.MutationResult<DeleteDataproductMutation>
export type DeleteDataproductMutationOptions = Apollo.BaseMutationOptions<
  DeleteDataproductMutation,
  DeleteDataproductMutationVariables
>
export const GcpGetDatasetsDocument = gql`
  query gcpGetDatasets($projectID: String!) {
    gcpGetDatasets(projectID: $projectID)
  }
`

/**
 * __useGcpGetDatasetsQuery__
 *
 * To run a query within a React component, call `useGcpGetDatasetsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGcpGetDatasetsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGcpGetDatasetsQuery({
 *   variables: {
 *      projectID: // value for 'projectID'
 *   },
 * });
 */
export function useGcpGetDatasetsQuery(
  baseOptions: Apollo.QueryHookOptions<
    GcpGetDatasetsQuery,
    GcpGetDatasetsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<GcpGetDatasetsQuery, GcpGetDatasetsQueryVariables>(
    GcpGetDatasetsDocument,
    options
  )
}
export function useGcpGetDatasetsLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    GcpGetDatasetsQuery,
    GcpGetDatasetsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<GcpGetDatasetsQuery, GcpGetDatasetsQueryVariables>(
    GcpGetDatasetsDocument,
    options
  )
}
export type GcpGetDatasetsQueryHookResult = ReturnType<
  typeof useGcpGetDatasetsQuery
>
export type GcpGetDatasetsLazyQueryHookResult = ReturnType<
  typeof useGcpGetDatasetsLazyQuery
>
export type GcpGetDatasetsQueryResult = Apollo.QueryResult<
  GcpGetDatasetsQuery,
  GcpGetDatasetsQueryVariables
>
export const GcpGetTablesDocument = gql`
  query gcpGetTables($projectID: String!, $datasetID: String!) {
    gcpGetTables(projectID: $projectID, datasetID: $datasetID) {
      name
      type
      description
    }
  }
`

/**
 * __useGcpGetTablesQuery__
 *
 * To run a query within a React component, call `useGcpGetTablesQuery` and pass it any options that fit your needs.
 * When your component renders, `useGcpGetTablesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGcpGetTablesQuery({
 *   variables: {
 *      projectID: // value for 'projectID'
 *      datasetID: // value for 'datasetID'
 *   },
 * });
 */
export function useGcpGetTablesQuery(
  baseOptions: Apollo.QueryHookOptions<
    GcpGetTablesQuery,
    GcpGetTablesQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<GcpGetTablesQuery, GcpGetTablesQueryVariables>(
    GcpGetTablesDocument,
    options
  )
}
export function useGcpGetTablesLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    GcpGetTablesQuery,
    GcpGetTablesQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<GcpGetTablesQuery, GcpGetTablesQueryVariables>(
    GcpGetTablesDocument,
    options
  )
}
export type GcpGetTablesQueryHookResult = ReturnType<
  typeof useGcpGetTablesQuery
>
export type GcpGetTablesLazyQueryHookResult = ReturnType<
  typeof useGcpGetTablesLazyQuery
>
export type GcpGetTablesQueryResult = Apollo.QueryResult<
  GcpGetTablesQuery,
  GcpGetTablesQueryVariables
>
export const UpdateDataproductDocument = gql`
  mutation updateDataproduct($id: ID!, $input: UpdateDataproduct!) {
    updateDataproduct(id: $id, input: $input) {
      id
    }
  }
`
export type UpdateDataproductMutationFn = Apollo.MutationFunction<
  UpdateDataproductMutation,
  UpdateDataproductMutationVariables
>

/**
 * __useUpdateDataproductMutation__
 *
 * To run a mutation, you first call `useUpdateDataproductMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateDataproductMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateDataproductMutation, { data, loading, error }] = useUpdateDataproductMutation({
 *   variables: {
 *      id: // value for 'id'
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateDataproductMutation(
  baseOptions?: Apollo.MutationHookOptions<
    UpdateDataproductMutation,
    UpdateDataproductMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useMutation<
    UpdateDataproductMutation,
    UpdateDataproductMutationVariables
  >(UpdateDataproductDocument, options)
}
export type UpdateDataproductMutationHookResult = ReturnType<
  typeof useUpdateDataproductMutation
>
export type UpdateDataproductMutationResult =
  Apollo.MutationResult<UpdateDataproductMutation>
export type UpdateDataproductMutationOptions = Apollo.BaseMutationOptions<
  UpdateDataproductMutation,
  UpdateDataproductMutationVariables
>
export const SearchContentDocument = gql`
  query searchContent($q: SearchQuery!) {
    search(q: $q) {
      excerpt
      result {
        ... on Collection {
          __typename
          id
          name
          description
          created
          lastModified
          owner {
            group
            teamkatalogenURL
          }
        }
        ... on Dataproduct {
          __typename
          id
          name
          description
          created
          lastModified
          owner {
            group
            teamkatalogenURL
          }
        }
      }
    }
  }
`

/**
 * __useSearchContentQuery__
 *
 * To run a query within a React component, call `useSearchContentQuery` and pass it any options that fit your needs.
 * When your component renders, `useSearchContentQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useSearchContentQuery({
 *   variables: {
 *      q: // value for 'q'
 *   },
 * });
 */
export function useSearchContentQuery(
  baseOptions: Apollo.QueryHookOptions<
    SearchContentQuery,
    SearchContentQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<SearchContentQuery, SearchContentQueryVariables>(
    SearchContentDocument,
    options
  )
}
export function useSearchContentLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    SearchContentQuery,
    SearchContentQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<SearchContentQuery, SearchContentQueryVariables>(
    SearchContentDocument,
    options
  )
}
export type SearchContentQueryHookResult = ReturnType<
  typeof useSearchContentQuery
>
export type SearchContentLazyQueryHookResult = ReturnType<
  typeof useSearchContentLazyQuery
>
export type SearchContentQueryResult = Apollo.QueryResult<
  SearchContentQuery,
  SearchContentQueryVariables
>
export const TeamkatalogenDocument = gql`
  query Teamkatalogen($q: String!) {
    teamkatalogen(q: $q) {
      name
      url
    }
  }
`

/**
 * __useTeamkatalogenQuery__
 *
 * To run a query within a React component, call `useTeamkatalogenQuery` and pass it any options that fit your needs.
 * When your component renders, `useTeamkatalogenQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useTeamkatalogenQuery({
 *   variables: {
 *      q: // value for 'q'
 *   },
 * });
 */
export function useTeamkatalogenQuery(
  baseOptions: Apollo.QueryHookOptions<
    TeamkatalogenQuery,
    TeamkatalogenQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<TeamkatalogenQuery, TeamkatalogenQueryVariables>(
    TeamkatalogenDocument,
    options
  )
}
export function useTeamkatalogenLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    TeamkatalogenQuery,
    TeamkatalogenQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<TeamkatalogenQuery, TeamkatalogenQueryVariables>(
    TeamkatalogenDocument,
    options
  )
}
export type TeamkatalogenQueryHookResult = ReturnType<
  typeof useTeamkatalogenQuery
>
export type TeamkatalogenLazyQueryHookResult = ReturnType<
  typeof useTeamkatalogenLazyQuery
>
export type TeamkatalogenQueryResult = Apollo.QueryResult<
  TeamkatalogenQuery,
  TeamkatalogenQueryVariables
>
export const UserInfoDetailsDocument = gql`
  query userInfoDetails {
    userInfo {
      name
      email
      groups {
        name
        email
      }
      gcpProjects {
        id
        group {
          name
          email
        }
      }
    }
  }
`

/**
 * __useUserInfoDetailsQuery__
 *
 * To run a query within a React component, call `useUserInfoDetailsQuery` and pass it any options that fit your needs.
 * When your component renders, `useUserInfoDetailsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useUserInfoDetailsQuery({
 *   variables: {
 *   },
 * });
 */
export function useUserInfoDetailsQuery(
  baseOptions?: Apollo.QueryHookOptions<
    UserInfoDetailsQuery,
    UserInfoDetailsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<UserInfoDetailsQuery, UserInfoDetailsQueryVariables>(
    UserInfoDetailsDocument,
    options
  )
}
export function useUserInfoDetailsLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    UserInfoDetailsQuery,
    UserInfoDetailsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<
    UserInfoDetailsQuery,
    UserInfoDetailsQueryVariables
  >(UserInfoDetailsDocument, options)
}
export type UserInfoDetailsQueryHookResult = ReturnType<
  typeof useUserInfoDetailsQuery
>
export type UserInfoDetailsLazyQueryHookResult = ReturnType<
  typeof useUserInfoDetailsLazyQuery
>
export type UserInfoDetailsQueryResult = Apollo.QueryResult<
  UserInfoDetailsQuery,
  UserInfoDetailsQueryVariables
>
export const UserInfoAccessableDataproductDocument = gql`
  query userInfoAccessableDataproduct {
    userInfo {
      accessable {
        __typename
        id
        name
        description
        created
        lastModified
        owner {
          group
          teamkatalogenURL
        }
      }
    }
  }
`

/**
 * __useUserInfoAccessableDataproductQuery__
 *
 * To run a query within a React component, call `useUserInfoAccessableDataproductQuery` and pass it any options that fit your needs.
 * When your component renders, `useUserInfoAccessableDataproductQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useUserInfoAccessableDataproductQuery({
 *   variables: {
 *   },
 * });
 */
export function useUserInfoAccessableDataproductQuery(
  baseOptions?: Apollo.QueryHookOptions<
    UserInfoAccessableDataproductQuery,
    UserInfoAccessableDataproductQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<
    UserInfoAccessableDataproductQuery,
    UserInfoAccessableDataproductQueryVariables
  >(UserInfoAccessableDataproductDocument, options)
}
export function useUserInfoAccessableDataproductLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    UserInfoAccessableDataproductQuery,
    UserInfoAccessableDataproductQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<
    UserInfoAccessableDataproductQuery,
    UserInfoAccessableDataproductQueryVariables
  >(UserInfoAccessableDataproductDocument, options)
}
export type UserInfoAccessableDataproductQueryHookResult = ReturnType<
  typeof useUserInfoAccessableDataproductQuery
>
export type UserInfoAccessableDataproductLazyQueryHookResult = ReturnType<
  typeof useUserInfoAccessableDataproductLazyQuery
>
export type UserInfoAccessableDataproductQueryResult = Apollo.QueryResult<
  UserInfoAccessableDataproductQuery,
  UserInfoAccessableDataproductQueryVariables
>
export const UserInfoUserProductsDocument = gql`
  query userInfoUserProducts {
    userInfo {
      dataproducts {
        __typename
        id
        name
        description
        created
        lastModified
        owner {
          group
          teamkatalogenURL
        }
      }
      collections {
        __typename
        id
        name
        description
        created
        lastModified
        owner {
          group
          teamkatalogenURL
        }
      }
    }
  }
`

/**
 * __useUserInfoUserProductsQuery__
 *
 * To run a query within a React component, call `useUserInfoUserProductsQuery` and pass it any options that fit your needs.
 * When your component renders, `useUserInfoUserProductsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useUserInfoUserProductsQuery({
 *   variables: {
 *   },
 * });
 */
export function useUserInfoUserProductsQuery(
  baseOptions?: Apollo.QueryHookOptions<
    UserInfoUserProductsQuery,
    UserInfoUserProductsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useQuery<
    UserInfoUserProductsQuery,
    UserInfoUserProductsQueryVariables
  >(UserInfoUserProductsDocument, options)
}
export function useUserInfoUserProductsLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    UserInfoUserProductsQuery,
    UserInfoUserProductsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions }
  return Apollo.useLazyQuery<
    UserInfoUserProductsQuery,
    UserInfoUserProductsQueryVariables
  >(UserInfoUserProductsDocument, options)
}
export type UserInfoUserProductsQueryHookResult = ReturnType<
  typeof useUserInfoUserProductsQuery
>
export type UserInfoUserProductsLazyQueryHookResult = ReturnType<
  typeof useUserInfoUserProductsLazyQuery
>
export type UserInfoUserProductsQueryResult = Apollo.QueryResult<
  UserInfoUserProductsQuery,
  UserInfoUserProductsQueryVariables
>
