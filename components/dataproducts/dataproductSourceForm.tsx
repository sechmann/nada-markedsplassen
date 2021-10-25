import { useGcpGetDatasetsQuery } from '../../lib/schema/graphql'
import { useContext, useState } from 'react'
import { UserState } from '../../lib/context'
import TreeView from '@mui/lab/TreeView'
import TreeItem from '@mui/lab/TreeItem'
import {
  FieldErrors,
  FieldValues,
  SetFieldValue,
  UseFormRegister,
  UseFormWatch,
} from 'react-hook-form'
import { Dataset } from './datasource/dataset'
import { Project } from './datasource/project'

interface DataproductSourceFormProps {
  register: UseFormRegister<FieldValues>
  watch: UseFormWatch<FieldValues>
  errors: FieldErrors<FieldValues>
  setValue: SetFieldValue<FieldValues>
}

export const DataproductSourceForm = ({
  register,
  errors,
  watch,
  setValue,
}: DataproductSourceFormProps) => {
  const user = useContext(UserState)

  const [activePaths, setActivePaths] = useState<string[]>([])
  register('bigquery.projectID')
  register('bigquery.dataset')
  register('bigquery.table')
  const group = watch('group')

  const teamProjects = user?.gcpProjects
    .filter((project) => project.group.email == group)
    .map((group) => group.id)

  const handleNodeSelect = (e: any, node: string) => {
    const [projectID, datasetID, tableID] = node.split('/')
    if (projectID && datasetID && tableID) {
      setValue('bigquery.projectID', projectID)
      setValue('bigquery.dataset', datasetID)
      setValue('bigquery.table', tableID)
    }
  }

  return (
    <TreeView
      onNodeSelect={handleNodeSelect}
      onNodeToggle={(x, n) => setActivePaths(n)}
    >
      {teamProjects?.map((projectID) => {
        return (
          <Project
            key={projectID}
            projectID={projectID}
            activePaths={activePaths}
          />
        )
      })}
    </TreeView>
  )
}
export default DataproductSourceForm
