import { Heading, Link, Alert, Modal } from "@navikt/ds-react";
import { isAfter, parseISO } from "date-fns";
import humanizeDate from "../../../lib/humanizeDate";
import { NewAccessRequest, SubjectType, UserInfoDetailsQuery } from "../../../lib/schema/graphql";
import DatasetTableSchema from "./datasetTableSchema";
import styled from "styled-components";
import Explore from "../../../components/dataproducts/explore";
import BigQueryLogo from "../../lib/icons/bigQueryLogo";
import IconBox from "../../lib/icons/iconBox";
import * as React from "react";
import DatasetMetadata from "./datasetMetadata";
import SpacedDiv from "../../lib/spacedDiv";
import { DatasetQuery } from "../../../lib/schema/datasetQuery";
import { useState } from "react";
import {KeywordBox, KeywordPill} from "../../lib/keywordList";
import DatasetOwnerMenu from "./datasetOwnerMenu";
import NewAccessRequestForm from "../accessRequest/newAccessRequest";


interface EntryProps {
    dataset: DatasetQuery
    userInfo: UserInfoDetailsQuery['userInfo'] | undefined
    isOwner: boolean
}

const findAccessType = (groups: UserInfoDetailsQuery['userInfo']['groups'] | undefined, dataset: DatasetQuery, isOwner: boolean) => {
    if (!groups) return {type: "utlogget"}
    if (isOwner) return {type: "owner"}
    const activeAccess = dataset.access.filter(a => (!a.revoked && (!a.expires || isAfter(parseISO(a.expires), Date.now()))))[0]
    if (activeAccess) return {type: "user", expires: activeAccess.expires}
    return {type: "none"}
}

const Section = styled.section`
    margin-bottom: 0.75rem;
    display: flex;
    flex-direction: column;
`

interface DatasetAlertProps {
    narrow?: boolean
}

const DatasetAlert = styled(Alert)<DatasetAlertProps>`
    width: ${(props) => props.narrow ? 'fit-content' : 'unset'};
    margin-bottom: 0.75rem;
`

const Article = styled.article`
    border-bottom: 1px solid #ddd;
    margin-bottom: 0.75rem;
    &:last-child {
        border-bottom: 0px;
    }
`

const DatasetHeading = styled(Heading)`
    display: inline-flex;
    align-items: center;
    gap: 0.75rem;
`

const MainView = styled.div`
    display: block;
    padding-top: 2rem;
    padding-right: 2rem;
`

const DatasetContainer = styled.div`
    display: flex;
`

const HeadingContainer = styled.div`
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 0.5rem;
`

const ModalLink = styled.a`
    cursor: pointer;
`
const Dataset = ({dataset, userInfo, isOwner}: EntryProps) => {
    const accessType = findAccessType(userInfo?.groups, dataset, isOwner)
    const [accessRequested, setAccessRequested] = useState(false);

    const defaultAccessRequestValues: NewAccessRequest = {
        owner: userInfo?.email,
        datasetID: dataset.id,
        expires: "",
        polly: null,
        subject: userInfo?.email,
        subjectType: SubjectType.User,
    }


    return <DatasetContainer>
        <Modal 
            open={accessRequested}
            aria-label='Søk om tilgang til datasettet'
            onClose={() => setAccessRequested(false)}
        >
            <Modal.Content>
                <NewAccessRequestForm dataset={dataset} newAccessRequest={defaultAccessRequestValues}/>
            </Modal.Content>
        </Modal>
        <MainView>
            {accessType.type === 'utlogget' && <DatasetAlert size="small" variant="info">
                Du er ikke innlogget
            </DatasetAlert>}
            {accessType.type === 'user' && <DatasetAlert size="small" variant="success">
                Du har tilgang{accessType.expires && ` til: ${humanizeDate(accessType.expires)}`}. <ModalLink onClick={() => setAccessRequested(true)}>Søk om tilgang</ModalLink>
            </DatasetAlert>}
            {accessType.type === 'none' && <DatasetAlert size="small" variant="info">
                Du har ikke tilgang til datasettet. <ModalLink onClick={() => setAccessRequested(true)}>Søk om tilgang</ModalLink>
            </DatasetAlert>}
            <HeadingContainer>
                <DatasetHeading level="2" size="large">
                    <IconBox size={42} inline={true}>
                        <BigQueryLogo/>
                    </IconBox>
                    {dataset.name}
                </DatasetHeading>
                {isOwner && <DatasetOwnerMenu />}
            </HeadingContainer>
            {dataset.pii 
                ? <DatasetAlert size="small" variant="warning" narrow={true}>Inneholder persondata</DatasetAlert>
                : <DatasetAlert size="small" variant="success" narrow={true}>Inneholder <b>ikke</b> persondata</DatasetAlert>
            }
            <Section>
                <Article>
                    <DatasetMetadata datasource={dataset.datasource}/>
                    <SpacedDiv>
                        <KeywordBox>
                            {dataset.keywords.map((keyword, idx) => <KeywordPill key={idx} keyword={keyword}>{keyword}</KeywordPill>)}
                        </KeywordBox>
                    </SpacedDiv>
                    <DatasetTableSchema datasource={dataset.datasource} />
                </Article>
                {userInfo && accessType.type !== "none" && <Article>
                    <Heading spacing level="3" size="small">
                        Utforsk
                    </Heading>
                    <Explore dataproductId={dataset.id} dataset={dataset} isOwner={accessType.type === "owner"}/>
                </Article>}
            </Section>
        </MainView>
    </DatasetContainer>
}


export default Dataset