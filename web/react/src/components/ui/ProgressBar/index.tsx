import * as React from 'react';

import 'stylesheets/06-components/progress-bar.scss';

type Props = {
    progress: number;
    inline?: boolean;
    displayLabel?: boolean;
}

export const ProgressBar: React.FC<Props> = ({
    progress,
    inline = false,
    displayLabel = true
}) => {
    const progressValue = progress.toFixed(2);
    const progressPercent = progressValue + '%';

    return (
        <div className={ 'progress-bar' + (inline ? ' -inline' : '') }>
            <div className="progress" style={{ width: progressPercent }}></div>
            { displayLabel ? <div className="label">{progressPercent}</div> : <></> }
        </div>
    );
};
