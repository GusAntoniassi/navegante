import * as React from 'react';

import 'stylesheets/06-components/app-sidebar.scss';

export const Sidebar = () => {
    return (
        <aside className="app-sidebar">
            <h1 className="brand">Navegante</h1>
            <hr className="ruler"></hr>
            <ul className="list">
                <li className="item -active">Containers</li>
                <li className="item">Networking</li>
                <li className="item">Images</li>
                <li className="item">Volumes</li>
                <li className="item">Logs</li>
                <li className="item">Notifications</li>
            </ul>
        </aside>
    );
};
