import React, { useEffect } from 'react';

import { Container } from '../Container';
import { LoadingIndicator } from 'components/ui/LoadingIndicator';
import { RefreshInterval } from 'components/ui/RefreshInterval';

import 'stylesheets/06-components/container-list.scss';

import { Container as EntityContainer, ContainerService } from 'core';

export const ContainerList = () => {
    const [containers, setContainers] = React.useState<Array<EntityContainer>>([]);
    const [loading, setLoading] = React.useState<boolean>(false);
    const [refreshIntervalSeconds, setRefreshIntervalSeconds] = React.useState<number>(30);

    const handleRefreshChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setRefreshIntervalSeconds(parseInt(event.target.value, 10));
    };

    const getContainers = () => {
        setLoading(true);
        ContainerService.getContainers().then(apiContainers => {
            setContainers(apiContainers);
        }).catch((err) => {
            // 404 is expected from the API when there are no containers running
            if (err.response && err.response.status === 404) {
                return;
            }

            // @TODO: Implement a better error treatment with toasts on screen
            console.error(err);
        }).finally(() => {
            setLoading(false);
        });
    };

    useEffect(() => {
        if (refreshIntervalSeconds <= 0) {
            return;
        }

        getContainers();

        const interval = setInterval(() => {
            getContainers();
        }, refreshIntervalSeconds*1000);

        // Clears the interval after this hook is "done"
        return () => clearInterval(interval);
    }, [refreshIntervalSeconds]);

    return (
        <div className="container container-list">
            <LoadingIndicator visible={loading}></LoadingIndicator>
            <div className="container text-right">
                <RefreshInterval
                    defaultValue={refreshIntervalSeconds}
                    handleChange={handleRefreshChange}
                ></RefreshInterval>
            </div>
            {containers.length === 0
                ? <h1>You don't have any containers running yet!</h1>
                : containers.map((container, index) => <Container key={index} container={container}/>)
            }
        </div>
    );
};
