import { gql } from 'graphql-tag'

export const GET_PRODUCT_AREAS = gql`
  query ProductAreas {
    productAreas {
      id
      name
      areaType
      dataproducts{
        id
        name
        description
        owner{
          group
        }
      }
      stories{
        id
        name
        created
        lastModified
        keywords
        owner {
            group
            teamkatalogenURL
        }
      }
      quartoStories{
        id
        name
        created
        lastModified
        keywords
      }
      insightProducts{
        id
        name
        created
        lastModified
        keywords
        type
        group
        teamkatalogenURL        
        link
      }
    teams{
        id
        name
        dataproducts{
          id
          name
        }
        stories{
          id
          name
        }
        quartoStories{
          id
          name
        }
        insightProducts{
          id
          name
        }
      }
    }
  }
`
