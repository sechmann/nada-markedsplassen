import { gql } from 'graphql-tag'

export const REVOKE_ACCESS = gql`
  mutation RevokeAccess($id: ID!) {
    revokeAccessToDataproduct(id: $id)
  }
`
