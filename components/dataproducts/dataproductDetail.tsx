import { useContext, useState } from 'react'
import { useRouter } from 'next/router'
import apiDELETE from '../../lib/api/delete'
import ErrorMessage from '../lib/error'
import LoaderSpinner from '../lib/spinner'
import Link from 'next/link'
import { Box, Tab, Tabs } from '@mui/material'
import TabPanel from '../lib/tabPanel'
import ReactMarkdown from 'react-markdown'
import DataproductTableSchema from './dataproductTableSchema'
import { DataproductSchema } from '../../lib/schema/schema_types'
import styled from 'styled-components'
import EditDataproduct from './editDataproduct'
import DotMenu from '../lib/editMenu'
import { PiiIkon } from '../lib/piiIkon'
import { MicroCard } from '@navikt/ds-react'
import { MetadataTable } from './metadataTable'
import { RepoKnapp } from '../widgets/RepoKnapp'

const StyledDiv = styled.div`
  display: flex;
  margin: 40px 0;
  justify-content: space-between;
  align-items: center;
`

export interface DataproductDetailProps {
  product: DataproductSchema
}

export const DataproductDetail = ({ product }: DataproductDetailProps) => {
  const [edit, setEdit] = useState(false)
  const [activeTab, setActiveTab] = useState(0)
  const router = useRouter()
  const deleteDataproduct = async (id: string) => {
    try {
      await apiDELETE(`/api/dataproducts/${id}`)
      await router.push('/')
    } catch (e: any) {
      setBackendError(e.toString())
    }
  }
  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setActiveTab(newValue)
  }
  const [backendError, setBackendError] = useState()

  if (!product) return <LoaderSpinner />

  return edit ? (
    <EditDataproduct dataproduct={product} close={() => setEdit(false)} />
  ) : (
    <div>
      {backendError && <ErrorMessage error={backendError} />}
      <StyledDiv>
        <h1>{product.name}</h1>
        <DotMenu
          onEdit={() => setEdit(true)}
          onDelete={async () => await deleteDataproduct(product.id)}
        />
      </StyledDiv>

      <MetadataTable product={product} />

      <div></div>
      <ReactMarkdown>
        {product.description || '*ingen beskrivelse*'}
      </ReactMarkdown>
      <Box sx={{ maxWidth: 480, bgcolor: 'background.paper' }}>
        <Tabs
          value={activeTab}
          onChange={handleChange}
          variant="scrollable"
          scrollButtons="auto"
          aria-label="scrollable auto tabs example"
        >
          <Tab label="Skjema" value={0} />
          <Tab label="Lineage" />
        </Tabs>
      </Box>
      <TabPanel index={0} value={activeTab}>
        <DataproductTableSchema product={product} />
      </TabPanel>
      <TabPanel index={1} value={activeTab}>
        <div>Placeholder</div>
      </TabPanel>
    </div>
  )
}
export default DataproductDetail
