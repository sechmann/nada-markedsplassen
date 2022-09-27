import { EllipsisCircleH } from '@navikt/ds-icons'
import { Button } from '@navikt/ds-react'
import { Dropdown, DropdownContext } from '@navikt/ds-react-internal'
import { useRouter } from 'next/router'
import { useState } from 'react'
import { GET_DATAPRODUCT } from '../../../lib/queries/dataproduct/dataproduct'
import {
  DataproductQuery,
  useDeleteDatasetMutation,
} from '../../../lib/schema/graphql'
import DeleteModal from '../../lib/deleteModal'

interface IDatasetOwnerMenuProps {
  datasetName: string
  datasetId: string
  dataproduct: DataproductQuery['dataproduct']
  setEdit: (value: boolean) => void
}

const DatasetOwnerMenu = ({
  datasetName,
  datasetId,
  dataproduct,
  setEdit,
}: IDatasetOwnerMenuProps) => {
  const [isOpen, setIsOpen] = useState(false)
  const [anchorEl, setAnchorEl] = useState<Element | null>(null)
  const [showDelete, setShowDelete] = useState(false)
  const [deleteError, setDeleteError] = useState('')
  const router = useRouter()

  const handleMenuButtonClick = (e: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(e.currentTarget)
    setIsOpen(!isOpen)
  }

  const [deleteDataset] = useDeleteDatasetMutation({
    variables: { id: datasetId },
    awaitRefetchQueries: true,
    refetchQueries: [
      {
        query: GET_DATAPRODUCT,
        variables: {
          id: dataproduct?.id,
        },
      },
    ],
  })

  const onDelete = async () => {
    try {
      await deleteDataset()
      await router.push(
        `/dataproduct/${dataproduct?.id}/${dataproduct?.slug}/info`
      )
    } catch (e: any) {
      setDeleteError(e.toString())
    }
  }

  return (
    <>
      <DropdownContext.Provider
        value={{ isOpen, setIsOpen, anchorEl, setAnchorEl }}
      >
        <Button
          variant="tertiary"
          onClick={handleMenuButtonClick}
        >
          <EllipsisCircleH />
        </Button>
        <Dropdown.Menu>
          <Dropdown.Menu.GroupedList>
            <Dropdown.Menu.GroupedList.Item onClick={() => setEdit(true)}>
              Endre datasett
            </Dropdown.Menu.GroupedList.Item>
            <Dropdown.Menu.GroupedList.Item onClick={() => setShowDelete(true)}>
              Slett datasett
            </Dropdown.Menu.GroupedList.Item>
          </Dropdown.Menu.GroupedList>
        </Dropdown.Menu>
      </DropdownContext.Provider>
      <DeleteModal
        name={datasetName}
        resource="datasett"
        error={deleteError}
        open={showDelete}
        onCancel={() => setShowDelete(false)}
        onConfirm={onDelete}
      ></DeleteModal>
    </>
  )
}

export default DatasetOwnerMenu
