import { yupResolver } from '@hookform/resolvers/yup';
import { Button, DatePicker, Heading, Loader, Radio, RadioGroup, TextField, useDatepicker} from '@navikt/ds-react'
import { useRouter } from 'next/router'
import { useContext, useState } from 'react'
import { Controller, useForm } from 'react-hook-form'
import AsyncSelect from 'react-select/async'
import * as yup from 'yup'
import { DatasetQuery } from '../../../lib/schema/datasetQuery'
import { UserState } from '../../../lib/context'
import ErrorMessage from '../../lib/error';
import { PollyInput, SubjectType } from '../../../lib/rest/access';
import { useSearchPolly } from '../../../lib/rest/polly';
import { set } from 'lodash';

const tomorrow = () => {
  const date = new Date()
  date.setDate(date.getDate() + 1)
  return date
}

const currentDate = (currentDate: any) => {
    if (typeof currentDate === 'string') return new Date(currentDate)
    else if (currentDate instanceof Date) {
      return currentDate
    }
    return undefined
}

const schema = yup
  .object({
    subject: yup
      .string()
      .required(
        'Du må skrive inn e-postadressen til hvem tilgangen gjelder for'
      )
      .email('E-postadresssen er ikke gyldig'),
    subjectType: yup
      .mixed<SubjectType>()
      .required('Du må velge hvem tilgangen gjelder for')
      .oneOf([SubjectType.User, SubjectType.Group, SubjectType.ServiceAccount]),
    accessType: yup
      .string()
      .required('Du må velge hvor lenge du ønsker tilgang')
      .oneOf(['eternal', 'until']),
    expires: yup
      .date()
      .nullable()
      .when('accessType', {
        is: 'until',
        then: ()=>yup.date().required('Du må angi en utløpsdato for tilgang'),
        otherwise: ()=>yup.date().nullable()
      })
    })
  .required()

export type AccessRequestFormInput = {
  id?: string
  datasetID: string
  expires?: Date
  polly?: PollyInput
  subject?: string
  subjectType?: SubjectType
  status?: string
  reason?: string
}

interface AccessRequestFormProps {
  accessRequest?: AccessRequestFormInput
  dataset: DatasetQuery
  isEdit: boolean
  onSubmit: (requestData: AccessRequestFormInput) => void
  error: Error | null
  setModal: (value: boolean) => void
}

interface AccessRequestFields {
  subject: string
  subjectType: SubjectType
  accessType: string
  expires?: Date | null | undefined
}

const AccessRequestFormV2 = ({
  setModal,
  accessRequest,
  dataset,
  isEdit,
  onSubmit,
  error,
}: AccessRequestFormProps) => {
  const [searchText, setSearchText] = useState('')
  const [polly, setPolly] = useState<PollyInput | undefined | null>(null)
  const [submitted, setSubmitted] = useState(false)
  const [loadOptionsCallBack, setLoadOptionsCallBack] = useState<Function | null>(null)
  const router = useRouter()
  const userInfo = useContext(UserState)

  const {
    register,
    handleSubmit,
    control,
    formState: { errors },
    setValue,
    getValues,
  } = useForm({
    resolver: yupResolver(schema),
    defaultValues: {
      subject: accessRequest?.subject ? accessRequest.subject : userInfo?.email ? userInfo.email :  "",
      subjectType: accessRequest?.subjectType
        ? accessRequest.subjectType
        : SubjectType.User,
      accessType: !isEdit || accessRequest?.expires ? 'until' : 'eternal',
      expires: isEdit && accessRequest?.expires ? accessRequest.expires : tomorrow(),
    },
  })

  const { datepickerProps, inputProps, selectedDay } = useDatepicker({
    defaultSelected: currentDate(getValues("expires")),
    fromDate: new Date(tomorrow()),
    onDateChange: (d: Date | undefined) => setValue("expires", d),
  });

  const {
    searchResult,
    searchError,
    loading,
  } = useSearchPolly(searchText)

  const onSubmitForm = (data: AccessRequestFields) => {
    setSubmitted(true)
    const accessRequest: AccessRequestFormInput = {
      datasetID: dataset.id,
      subject: data.subject,
      subjectType: data.subjectType,
      polly: polly??undefined,
      expires: data.accessType === 'until' ? data.expires? new Date(data.expires) : undefined: undefined,
    }
    onSubmit(accessRequest)
  }

  interface Option {
    value: string
    label: string
  }

  console.log(searchResult)
  loadOptionsCallBack?.(
    searchResult
      ? searchResult.map((el) => {
          return { value: el.externalID, label: el.name }
        })
      : []
  )

  const loadOptions = (
    input: string,
    callback: (options: Option[]) => void
  ) => {
    console.log(input)
    setSearchText(input)
    setLoadOptionsCallBack(() => callback)
  }

  const onInputChange = (newOption: Option | null) => {
    newOption != null
      ? searchResult &&
        setPolly(searchResult.find((e) => e.externalID == newOption.value))
      : setPolly(null)
  }

  return (
    <div className="h-full">
      <Heading level="1" size="large" className="pb-8">
        Tilgangssøknad for {dataset.name}
      </Heading>
      <form
        onSubmit={handleSubmit(onSubmitForm)}
        className="flex flex-col gap-10 h-[90%]"
      >
        <div>
          <Controller
            name="subjectType"
            control={control}
            render={({ field }) => (
              <RadioGroup
                {...field}
                legend="Hvem gjelder tilgangen for?"
                error={errors?.subjectType?.message}
              >
                <Radio disabled={isEdit} value={SubjectType.User}>
                  Bruker
                </Radio>
                <Radio disabled={isEdit} value={SubjectType.Group}>
                  Gruppe
                </Radio>
                <Radio disabled={isEdit} value={SubjectType.ServiceAccount}>
                  Servicebruker
                </Radio>
              </RadioGroup>
            )}
          />
          <TextField
            {...register('subject')}
            disabled={isEdit}
            className="hidden-label"
            label="E-post-adresse"
            placeholder="Skriv inn e-post-adresse"
            error={errors?.subject?.message?.toString()}
            size="medium"
          />
        </div>
        <div>
          <Controller
            name="accessType"
            control={control}
            render={({ field }) => (
              <RadioGroup
                {...field}
                legend="Hvor lenge ønsker du tilgang?"
                error={errors?.accessType?.message}
              >
                <Radio value="until">Til dato</Radio>
                <DatePicker {...datepickerProps}>
                  <DatePicker.Input 
                    {...inputProps} 
                    label="" 
                    disabled={field.value === 'eternal'} 
                    error={errors?.expires?.message?.toString()} 
                  />
                </DatePicker>
                <Radio value="eternal">For alltid</Radio>
              </RadioGroup>
            )}
          />
          <div>
            <label className="navds-label">
              Velg behandling fra behandlingskatalogen
            </label>
            <AsyncSelect
              className="pt-2"
              classNamePrefix="select"
              isClearable
              placeholder="Skriv inn navnet på behandlingen"
              noOptionsMessage={({ inputValue }) =>
                inputValue ? 'Finner ikke behandling' : null
              }
              loadingMessage={() => 'Søker etter behandling...'}
              loadOptions={loadOptions}
              isLoading={loading}
              onChange={onInputChange}
              menuIsOpen={true}
            />
          </div>
        </div>
        { error && <ErrorMessage error={error} /> }
        {submitted && !error && <div>Vennligst vent...<Loader size="small"/></div>}
        <div className="flex flex-row gap-4 grow items-end pb-8">
          <Button
            type="button"
            variant="secondary"
            onClick={() => {
              setModal(false)
              router.push(`/user/requests`)
            }}
          >
            Avbryt
          </Button>
          <Button type="submit" disabled={submitted}>Lagre</Button>
        </div>
      </form>
    </div>
  )
}

export default AccessRequestFormV2
