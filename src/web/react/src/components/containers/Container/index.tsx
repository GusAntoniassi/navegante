import * as React from 'react';

import { Container as EntityContainer } from 'core';

type Props = {
    container: EntityContainer | null;
}

export const Container: React.FC<Props> = ({
    container
}) => {
    if (container == null) return (null);

    return (
        <dl>
            <dd>Name: {container.name}</dd>
            <dd>ID: {container.id}</dd>
            <dd>
                <dl>
                    <dt>Ports:</dt>
                    {container.ports.map((port, index) => <dd key={index}>{port}</dd>)}
                </dl>
            </dd>
            <dd>
                <dl>
                    <dt>Volumes:</dt>
                    {container.volumes.map((volume, index) => <dd key={index}>{volume}</dd>)}
                </dl>
            </dd>
        </dl>
    );
};