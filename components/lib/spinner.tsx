import styled from 'styled-components'
import { Loader } from '@navikt/ds-react'

const CenteredSpinner = styled.div`
  margin: 20% auto;
  width: fit-content;
`

export const DataProductSpinner = () => (
  <CenteredSpinner>
    <Loader size="2xlarge" transparent />
  </CenteredSpinner>
)

export default DataProductSpinner
