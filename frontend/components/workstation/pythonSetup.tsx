import React from 'react';
import {BodyLong, BodyShort, Heading, VStack} from "@navikt/ds-react";

// Need to add a python section with the following documentation:
// We have configured a global pip config under /etc/pip.conf, for use with pip, uv, etc.
// However, using this config requires authentication, which means you need to run the following commands:
//
//
// $ gcloud auth login --update-adc
// $ pypi-auth
// Where pypi-auth will configure the .netrc file located in your $HOME directory for the pypi artifact registry using an auth token obtained from glcoud auth login --update-adc, this is a short-lived token, e.g., 1 hour, and you will need to refresh after that period by running pypi-auth again

const PythonSetup: React.FC = () => {
    return (
        <div className="basis-1/2">
            <VStack gap="5">
                <Heading size="medium">Autentisering mot pypi-proxy</Heading>
                <BodyShort size="medium">Vi har laget en global deny regel mot <strong>pypi.org</strong>, for å kunne laste ned Python pakker ønsker vi at dere går via vår pypi-proxy.</BodyShort>
                <BodyShort size="medium">Vi har konfigurert en global pip config <strong>/etc/pip.conf</strong>, som ikke tar stilling til hvilken package manager (uv, pip, poetry) du bruker, så lenge den er kompatibel med pip.</BodyShort>
                <BodyShort size="medium">For å kunne bruke denne configen må du kjøre følgende kommandoer:</BodyShort>
                <code>
                    $ gcloud auth login --update-adc
                    <br />
                    $ pypi-auth
                </code>
                <Heading size="small">gcloud auth login --update-adc</Heading>
                <BodyShort size="medium">Med denne kommandoen logger du inn på Google Cloud med din personlige bruker og oppdaterer Application Default Credentials (ADC).</BodyShort>
                <Heading size="small">pypi-auth</Heading>
                <BodyShort size="medium">Denne kommandoen vil konfigurere <strong>.netrc</strong> filen i ditt <strong>$HOME</strong> directory for pypi artifact registry ved å bruke en autentiseringsnøkkel som er hentet fra <code>gcloud auth login --update-adc</code>. Denne nøkkelen er kortlevd, 1 time, og du må oppdatere den etter denne perioden ved å kjøre <code>pypi-auth</code> på nytt.</BodyShort>
            </VStack>
        </div>
    );
};

export default PythonSetup;
