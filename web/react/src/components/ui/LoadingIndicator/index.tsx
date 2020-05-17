import * as React from 'react';

import { AiOutlineLoading } from 'react-icons/ai';

import 'stylesheets/06-components/loading-indicator.scss';

type Props = {
    visible: boolean;
}

export const LoadingIndicator: React.FC<Props> = ({
    visible
}) => {
    if (!visible) return <></>;

    return (
        <div className="loading-indicator">
            <AiOutlineLoading className="spin"/>
        </div>
    );
};
