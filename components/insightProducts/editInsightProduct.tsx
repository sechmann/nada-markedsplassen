import { yupResolver } from '@hookform/resolvers/yup'
import { FieldValues, useForm } from 'react-hook-form'
import ErrorMessage from '../lib/error'
import { useRouter } from 'next/router'
import { useMutation } from '@apollo/client'
import TeamkatalogenSelector from '../lib/teamkatalogenSelector'
import DescriptionEditor from '../lib/DescriptionEditor'
import {
    Button,
    Heading,
    TextField,
    Select,
    Checkbox,
} from '@navikt/ds-react'
import amplitudeLog from '../../lib/amplitude'
import * as yup from 'yup'
import { useContext, useState } from 'react'
import TagsSelector from '../lib/tagsSelector'
import { UserState } from "../../lib/context";
import { UPDATE_INSIGHT_PRODUCT_METADATA } from '../../lib/queries/insightProducts/editInsightProduct'
import { useUpdateInsightProductMetadataMutation } from '../../lib/schema/graphql'
import { USER_INFO } from '../../lib/queries/userInfo/userInfo'

const schema = yup.object().shape({
    name: yup.string().nullable().required('Skriv inn navnet på innsiktsproduktet'),
    description: yup.string(),
    teamkatalogenURL: yup.string().required('Du må velge team i teamkatalogen'),
    keywords: yup.array().of(yup.string()),
    group: yup.string(),
    sensitiveInfo: yup.boolean().oneOf([true], 'Bekreft at innsiktsproduktet ikke inneholder personsensitive eller identifiserende opplysninger'),
    type: yup.string().nullable().required('Velg en type for innsiktsproduktet'),
    link: yup
        .string()
        .required('Du må legge til en lenke til innsiktsproduktet')
        .url('Lenken må være en gyldig URL'), // Add this line to validate the link as a URL    type: yup.string().required('Du må velge en type for innsiktsproduktet'),
})

export interface EditInsightProductMetadataFields {
    id: string
    name: string
    description: string
    keywords: string[]
    teamkatalogenURL: string
    group: string
    type: string
    link: string
}

export const EditInsightProductMetadataForm = ({ id, name, description, type, link, keywords, teamkatalogenURL, group }: EditInsightProductMetadataFields) => {
    const router = useRouter()
    const [productAreaID, setProductAreaID] = useState<string>('')
    const [teamID, setTeamID] = useState<string>('')
    const userInfo = useContext(UserState)
    const [updateInsightProductQuery, { loading, error }] = useUpdateInsightProductMetadataMutation()
    const {
        register,
        handleSubmit,
        watch,
        formState,
        setValue,
        control,
    } = useForm({
        resolver: yupResolver(schema),
        defaultValues: {
            name: name,
            description: description,
            keywords: keywords,
            teamkatalogenURL: teamkatalogenURL,
            group: group,
            sensitiveInfo: false,
            type: type,
            link: link,
        },
    })

    const { errors } = formState
    const kw = watch('keywords')

    const onDeleteKeyword = (keyword: string) => {
        setValue(
            'keywords',
            kw.filter((k: string) => k !== keyword)
        )
    }

    const onAddKeyword = (keyword: string) => {
        kw
            ? setValue('keywords', [...kw, keyword])
            : setValue('keywords', [keyword])
    }

    const valueOrNull = (val: string) => (val == '' ? null : val)

    const onSubmit = async (data: any) => {
        const editInsightProductData = {
            variables: {
                id: id,
                name: data.name,
                description: data.description,
                type: data.type,
                link: data.link,
                keywords: data.keywords,
                teamkatalogenURL: data.teamkatalogenURL,
                productAreaID: productAreaID,
                teamID: teamID,
                group: data.group,
            },
            refetchQueries: [
                {
                    query: USER_INFO,
                }]
        }

        updateInsightProductQuery(editInsightProductData).then(() => {
            amplitudeLog('skjema fullført', { skjemanavn: 'endre-innsiktsprodukt' })
            router.back()
        }).catch(e => {
            console.log(e)
            amplitudeLog('skjemainnsending feilet', {
                skjemanavn: 'endre-innsiktsprodukt',
            })
        })
    }

    const onCancel = () => {
        amplitudeLog('Klikker på: Avbryt',
            {
                pageName: 'endre-innsiktsprodukt',
            })
        router.back()
    }

    const onError = (errors: any) => {
        amplitudeLog('skjemavalidering feilet', {
            skjemanavn: 'endre-innsiktsprodukt',
            feilmeldinger: Object.keys(errors)
                .map((errorKey) => errorKey)
                .join(','),
        })
    }

    return (
        <div className="mt-8 md:w-[46rem]">
            <Heading level="1" size="large">
                Rediger innsiktsprodukt
            </Heading>
            <form
                className="pt-12 flex flex-col gap-10"
                onSubmit={handleSubmit(onSubmit, onError)}
            >
                <TextField
                    className="w-full"
                    label="Navn på innsiktsprodukt"
                    {...register('name')}
                    error={errors.name?.message?.toString()}
                />
                <DescriptionEditor
                    label="Beskrivelse av hva innsiktsproduktet kan brukes til"
                    name="description"
                    control={control}
                />
                <TeamkatalogenSelector
                    gcpGroups={userInfo?.gcpProjects.map(it => it.group.email)}
                    register={register}
                    watch={watch}
                    errors={errors}
                    setProductAreaID={setProductAreaID}
                    setTeamID={setTeamID}
                />
                <Select
                    className="w-full"
                    label="Velg type innsiktsprodukt"
                    {...register('type')}
                    error={errors.type?.message?.toString()}
                >
                    <option value="">Velg type</option>
                    <option value="Amplitude">Amplitude</option>
                    <option value="Grafana">Grafana</option>
                    <option value="HotJar">HotJar</option>
                    <option value="TaskAnalytics">TaskAnalytics</option>
                    <option value="Metabase">Metabase</option>
                    <option value="Annet">Annet</option>
                </Select>
                <TextField
                    className="w-full"
                    label="Lenke til innsiktsprodukt"
                    {...register('link')}
                    error={errors.link?.message?.toString()}
                />
                <TagsSelector
                    onAdd={onAddKeyword}
                    onDelete={onDeleteKeyword}
                    tags={kw || []}
                />
                <Checkbox
                    {...register('sensitiveInfo')}
                    errorId={errors.sensitiveInfo?.message?.toString()}
                >
                    Jeg bekrefter at innsiktsproduktet ikke inneholder personsensitive eller identifiserende opplysninger
                </Checkbox>
                {error && <ErrorMessage error={error} />}
                <div className="flex flex-row gap-4 mb-16">
                    <Button type="button" variant="secondary" onClick={onCancel}>
                        Avbryt
                    </Button>
                    <Button type="submit" disabled={loading}>Lagre</Button>
                </div>
            </form>
        </div>
    )
}
