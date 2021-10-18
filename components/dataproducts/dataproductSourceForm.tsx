import {
  FieldErrors,
  FieldValues,
  UseFormRegister,
  UseFormWatch,
} from 'react-hook-form'
import { useContext, useEffect, useState } from 'react'
import { components } from '../../lib/schema/schema'
import Tree, {
  Node as TreeNode,
  NodeId as TreeNodeId,
  NodeList,
} from '@naisutech/react-tree'
import { AuthState } from '../../lib/context'
import { Fieldset } from '@navikt/ds-react'

interface DataproductSourceFormProps {
  register: UseFormRegister<FieldValues>
  watch: UseFormWatch<FieldValues>
  errors: FieldErrors<FieldValues>
}

interface ProjectTableResponse {
  DataproductID: string
  Dataset: string
  ProjectID: string
  TableName: string
}

export const DataproductSourceForm = ({
  register,
  watch,
  errors,
}: DataproductSourceFormProps) => {
  const [teams, setTeams] = useState<components['schemas']['Group'][]>([])
  type TeamID = string
  type ProjectID = string
  type DatasetID = string
  type TableID = string

  const [teamProjects, setTeamProjects] = useState<Record<TeamID, ProjectID[]>>(
    {}
  )
  const [projectContents, setProjectContents] = useState<
    Record<ProjectID, ProjectTableResponse[]>
  >({})
  const [datasets, setDatasets] = useState<Record<ProjectID, DatasetID[]>>({})
  const [tables, setTables] = useState<Record<DatasetID, TableID[]>>({})

  const [loadedTeams, setLoadedTeams] = useState<string[]>([])
  const [loadedProjects, setLoadedProjects] = useState<string[]>([])

  const [nodes, setNodes] = useState<NodeList>([])

  const user = useContext(AuthState).user

  const teamName = (id: string) => id.split(':', 2)[1]
  const projectName = (id: string) => id.split(':', 3)[2]
  const datasetName = (id: string) => id.split(':', 4)[3]
  const tableName = (id: string) => id.split(':', 5)[4]

  useEffect(() => {
    if (!user?.groups) return
    setTeams(user.groups.filter((g) => g.email !== 'all-users@nav.no'))
  }, [user])

  // Processes the full state to create a NodeList for the <Tree>
  useEffect(() => {
    const teamsList = teams.map(
      (t): TreeNode => ({
        id: `team:${t.email}`,
        parentId: null,
        label: t.name,
      })
    )
    for (const [team, projectList] of Object.entries(teamProjects)) {
      projectList.forEach((project) =>
        teamsList.push({
          id: project,
          parentId: team,
          label: projectName(project),
        })
      )
    }
    for (const [project, datasetList] of Object.entries(datasets)) {
      datasetList.forEach((datasetID) => {
        teamsList.push({
          id: datasetID,
          parentId: project,
          label: datasetName(datasetID),
        })
      })
    }
    for (const [dataset, tableList] of Object.entries(tables)) {
      tableList.forEach((tableID) => {
        teamsList.push({
          id: tableID,
          parentId: dataset,
          label: tableName(tableID),
        })
      })
    }
    setNodes(teamsList)
  }, [teams, teamProjects, datasets, tables])

  // When the list of opened teams changes, make sure the relevant Record is populated.
  useEffect(() => {
    for (const teamID of loadedTeams.filter((t) => !(t in teamProjects))) {
      fetch(`/api/groups/${teamName(teamID)}/gcp_projects`)
        .then((res) => res.json())
        .then((p: string[]) => {
          setTeamProjects({
            ...teamProjects,
            [teamID]: p.map((p) => `project:${teamName(teamID)}:${p}`),
          })
        })
    }
  }, [loadedTeams])

  // When the list of opened projects changes, make sure the relevant Record is populated.
  useEffect(() => {
    for (const projectID of loadedProjects.filter(
      (p) => !(p in projectContents)
    )) {
      fetch(`/api/gcp/${projectName(projectID)}/tables`)
        .then((res) => res.json())
        .then((d: ProjectTableResponse[]) => {
          setProjectContents({
            ...projectContents,
            [projectID]: d,
          })
        })
    }
  }, [loadedProjects])

  useEffect(() => {
    const datasetsInResponse = (
      projectID: string,
      res: ProjectTableResponse[]
    ) => {
      return Array.from(
        new Set(
          res.map(
            (s) =>
              `dataset:${teamName(projectID)}:${projectName(projectID)}:${
                s.Dataset
              }`
          )
        )
      )
    }

    const tablesInDataset = (
      datasetID: string,
      res: ProjectTableResponse[]
    ) => {
      return Array.from(
        new Set(
          res
            .filter(
              (s) =>
                s.Dataset === datasetName(datasetID) &&
                s.ProjectID === projectName(datasetID)
            )
            .map(
              (s) =>
                `dataset:${teamName(datasetID)}:${projectName(
                  datasetID
                )}:${datasetName(datasetID)}:${s.TableName}`
            )
        )
      )
    }

    for (const project in projectContents) {
      if (!(project in datasets)) {
        setDatasets({
          ...datasets,
          [project]: datasetsInResponse(project, projectContents[project]),
        })

        for (const dataset of datasetsInResponse(
          project,
          projectContents[project]
        )) {
          console.log(dataset)
          setTables((tables) => ({
            ...tables,
            [dataset]: tablesInDataset(dataset, projectContents[project]),
          }))
        }
      }
    }
  }, [projectContents])
  useEffect(() => console.log('tables:', tables), [tables])
  const openCloseHandler = (nodeIds: TreeNodeId[]) => {
    const nodeIsTeam = (id: string) => id.split(':', 1)[0] === 'team'
    const nodeIsProject = (id: string) => id.split(':', 1)[0] === 'project'

    // Get the full node objects from just the ID
    const openNodes = nodes.filter((n) => nodeIds.includes(n.id))

    for (const { id } of openNodes) {
      if (nodeIsTeam(id as string) && !(id in loadedTeams)) {
        setLoadedTeams([...loadedTeams, id as string])
      }
      if (nodeIsProject(id as string) && !(id in loadedProjects)) {
        setLoadedProjects([...loadedProjects, id as string])
      }
    }
  }

  return (
    <Fieldset legend="Datakilde" errorPropagation={false}>
      <Tree theme="light" nodes={nodes} onOpenClose={openCloseHandler} />
    </Fieldset>
  )
}
