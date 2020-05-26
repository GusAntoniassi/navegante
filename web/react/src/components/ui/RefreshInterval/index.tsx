import * as React from 'react';

type Props = {
    defaultValue: number,
    handleChange: (event: React.ChangeEvent<HTMLSelectElement>) => void
}

export const RefreshInterval: React.FC<Props> = ({
    defaultValue,
    handleChange
}) => {
    const allowedValues = [5, 10, 30, 60, 300, 600];

    return (
        <>
            <label style={{ marginRight: "5px" }}>Refresh interval:&nbsp;</label>
            <select defaultValue={defaultValue} onChange={handleChange}>
                <option value={0}>Off</option>
                {allowedValues.map((value, index) =>
                    <option value={value} key={index}>{value < 60 ? value + 's' : value/60 + 'm'}</option>)
                }
            </select>
            <hr></hr>
        </>
    );
};
