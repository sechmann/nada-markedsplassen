import * as React from 'react'
import { Select, Table } from '@navikt/ds-react'
import LoaderSpinner from '../../lib/spinner'
import { ApolloError } from '@apollo/client'
import {
  ColumnType,
  DEFAULT_COLUMN_TAG,
  PIITagNames,
  PIITagOptions,
  PIITagType,
} from './useColumnTags'
import {PersonopplysningerDetaljert} from "./helptext";

interface AnnotateDatasetTableProps {
  loading: boolean
  error: ApolloError | undefined
  columns: ColumnType[] | undefined
  tags: Map<string, PIITagType> | undefined
  annotateColumn: (columnName: string, tag: PIITagType) => void
}

const AnnotateDatasetTable = ({
  loading,
  error,
  columns,
  tags,
  annotateColumn,
}: AnnotateDatasetTableProps) => {
  if (loading) {
    return <LoaderSpinner />
  }

  if (error) {
    console.log(error)
    return <div>Kan ikke hent skjemainformasjon</div>
  }

  if (!columns) return <div>Ingen skjemainformasjon</div>

  return (
    <div className="mb-3 w-[91vw] overflow-auto">
      <p className="flex gap-2 items-center mb-2">Klassifiser personopplysningene <PersonopplysningerDetaljert /></p>
      <Table className="w-[60rem]" size="small">
        <Table.Header>
          <Table.Row>
            <Table.HeaderCell>Name</Table.HeaderCell>
            <Table.HeaderCell>Mode</Table.HeaderCell>
            <Table.HeaderCell>Type</Table.HeaderCell>
            <Table.HeaderCell>Description</Table.HeaderCell>
            <Table.HeaderCell>Personopplysning</Table.HeaderCell>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {columns?.map((row) => (
            <Table.Row key={row.name}>
              <Table.DataCell>{row.name}</Table.DataCell>
              <Table.DataCell>{row.mode}</Table.DataCell>
              <Table.DataCell>{row.type}</Table.DataCell>
              <Table.DataCell>{row.description}</Table.DataCell>
              <Table.DataCell className="w-60">
                <Select
                  className="w-full"
                  size="small"
                  label=""
                  value={
                    tags && tags.has(row.name)
                      ? tags.get(row.name)
                      : DEFAULT_COLUMN_TAG
                  }
                  onChange={(e) =>
                    annotateColumn(row.name, e.target.value as PIITagType)
                  }
                >
                  {Array.from(PIITagOptions).map(([key, name]) => (
                    <option value={key} key={key}>
                      {name}
                    </option>
                  ))}
                </Select>
              </Table.DataCell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
    </div>
  )
}
export default AnnotateDatasetTable
