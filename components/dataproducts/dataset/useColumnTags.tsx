import { useState } from 'react'
import {
  BigQueryType,
  useGcpGetColumnsQuery,
  DatasetQuery,
} from '../../../lib/schema/graphql'

export type PIITagType =
  | 'PII_DirekteIdentifiserende'
  | 'PII_SærligKategori'
  | 'PII_KanVæreIndirekteIdentifiserende'

export const PIITagNames = new Map([
  ['PII_DirekteIdentifiserende', 'Direkte identifiserende'],
  ['PII_SærligKategori', 'Særlig kategori'],
  ['PII_KanVæreIndirekteIdentifiserende', 'Kan være indirekte identifiserende'],
])

export const PIITagOptions = new Map([
  ['PII_DirekteIdentifiserende', 'Direkte identifiserende'],
  ['PII_SærligKategori', 'Særlig kategori'],
  ['PII_KanVæreIndirekteIdentifiserende', 'Ikke klassifisert'],
])

export const DEFAULT_COLUMN_TAG = 'PII_KanVæreIndirekteIdentifiserende'

type TagsMapType = Map<string, Map<string, PIITagType> | undefined>
type PseudoColumnsMapType = Map<string, Map<string, boolean>>

const buildTableKey = (projectID: string, datasetID: string, tableID: string) =>
  `${projectID}.${datasetID}.${tableID}`

export type ColumnType = {
  name: string
  type: BigQueryType
  mode: string
  description: string
}

export type AnnotateColumnListener = (column: string, tag: PIITagType) => void
export type PseudoColumnListener = (column: string, on: boolean) => void

export const useColumnTags = (
  projectID: string,
  datasetID: string,
  tableID: string,
  dataset?: DatasetQuery['dataset']
) => {
  const [tagsMap, setTagsMap] = useState<TagsMapType>(
    new Map<string, Map<string, PIITagType>>()
  )

  const [pseudoColumnsMap, setPseudoColumnsMap] = useState<PseudoColumnsMapType>(
    new Map<string, Map<string, boolean>>()
  )

  const columnsQuery = useGcpGetColumnsQuery({
    variables: {
      projectID: projectID,
      datasetID: datasetID,
      tableID: tableID,
    },
  })

  var tableKey = buildTableKey(projectID, datasetID, tableID)

  const initDefaultStatesForTable = () => {
    if (
      projectID &&
      datasetID &&
      tableID &&
      !tagsMap.has(tableKey) &&
      !columnsQuery.error &&
      !columnsQuery.loading &&
      columnsQuery.data
    ) {
      var newTagsMap = new Map<string, Map<string, PIITagType> | undefined>(
        tagsMap
      )
      var tags = new Map<string, PIITagType>()
      var tagsFromQuery = JSON.parse(dataset?.datasource.piiTags || '{}')

      var newPseudoColumnsMap = new Map<string, Map<string, boolean>>(pseudoColumnsMap)
      var pseudoColumns = new Map<string, boolean>()
      columnsQuery.data.gcpGetColumns.forEach((it) =>{
        tags.set(
          it.name,
          (!!tagsFromQuery[it.name] &&
            PIITagNames.has(tagsFromQuery[it.name]) &&
            tagsFromQuery[it.name]) ||
            DEFAULT_COLUMN_TAG
        )
        pseudoColumns.set(
          it.name, false
        )
      }
      )
      newTagsMap.set(tableKey, tags)
      newPseudoColumnsMap.set(tableKey, pseudoColumns)
      setTagsMap(newTagsMap)
      setPseudoColumnsMap(newPseudoColumnsMap)
    }
  }

  initDefaultStatesForTable()

  const annotateColumn = (column: string, tag: PIITagType) => {
    //we cannot directly update the tagsMap and set it back, because the reference to the object
    //will not change, and react do not know it should refresh the components
    var newTagsMap = new Map<string, Map<string, PIITagType> | undefined>(
      tagsMap.set(tableKey, tagsMap.get(tableKey)?.set(column, tag))
    )
    setTagsMap(newTagsMap)
  }

  const selectPseudoColumn = (column: string, on: boolean) => {
    var newPseudoColumnsMap = new Map<string, Map<string, boolean>>(
      pseudoColumnsMap.set(tableKey, pseudoColumnsMap.get(tableKey)?.set(column, on) || new Map<string, boolean>())
    )
    setPseudoColumnsMap(newPseudoColumnsMap)
  }

  return {
    columns:
      !columnsQuery.error && !columnsQuery.loading
        ? (columnsQuery.data?.gcpGetColumns as ColumnType[])
        : undefined,
    loading: columnsQuery.loading,
    error: columnsQuery.error,
    tags: tagsMap.get(tableKey),
    pseudoColumns: pseudoColumnsMap.get(tableKey) || new Map<string, boolean>(),
    annotateColumn: annotateColumn,
    selectPseudoColumn,
  }
}
