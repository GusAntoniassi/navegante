import React, { useEffect } from 'react';

import { Container } from '../Container';
import 'stylesheets/06-components/container-list.scss';

import { Container as EntityContainer, ContainerService } from 'core';

export const ContainerList = () => {
    const [containers, setContainers] = React.useState<Array<EntityContainer>>([]);

    useEffect(() => {
        ContainerService.getContainers().then(apiContainers => {
            setContainers(apiContainers);
        }).catch((err) => {
            // 404 is expected from the API when there are no containers running
            if (err.response && err.response.status === 404) {
                return;
            }

            // @TODO: Implement a better error treatment with toasts on screen
            console.error(err);
        });
    }, []);

    return <div className="container container-list">
        {/* @TODO: Implement a cute lil' loading class here */}
        {containers.length === 0
            ? <h1>You don't have any containers running yet!</h1>
            : containers.map((container, index) => <Container key={index} container={container}/>)
        }
    </div>;
};
