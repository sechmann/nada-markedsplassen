import SearchResult from './searchResult'
import styled from 'styled-components'
import { Loader } from '@navikt/ds-react'
import useSWR from 'swr'
import fetcher from '../../lib/api/fetcher'
import { SearchResultEntry } from '../../lib/schema/schema_types'
import { useRouter } from 'next/router'
import SearchResultLink from './searchResult'

const NoResultsYetBox = styled.div`
  margin: 0 auto;
`

export function Results() {
  const router = useRouter()
  let { q } = router.query
  if (typeof q !== 'string') q = ''

  const { data, error } = useSWR<SearchResultEntry[], Error>(
    `/api/search?q=${q}`,
    fetcher
  )

  if (error) {
    return (
      <NoResultsYetBox>
        <h1>Error</h1>
      </NoResultsYetBox>
    )
  }

  if (!data) {
    return (
      <NoResultsYetBox>
        <Loader transparent />
      </NoResultsYetBox>
    )
  }

  return (
    <div>
      {!data.length ? (
        <div>Ingen resultater funnet</div>
      ) : (
        data.map((d) => {
          return <SearchResultLink key={d.id} result={d} />
        })
      )}
    </div>
  )
}

export default Results
