import * as React from 'react';
import { SidebarToggleButton } from './SidebarToggleButton';

import 'stylesheets/06-components/app-header.scss';
import sassVariables from 'stylesheets/01-settings/exports.module.scss';
import useWindowDimensions from 'hooks/useWindowDimensions';

export const Header = () => {
    const { width: windowWidth } = useWindowDimensions();
    const collapseBreakpoint = parseInt(sassVariables['breakpoint-medium'], 10);

    return (
        <header className="app-header">
            {/*
                @TODO: As an improvement, move the `collapsed` state to the sidebar component
                and figure out the proper way to pass it to the toggle button
            */}
            <SidebarToggleButton sidebarCollapsed={windowWidth < collapseBreakpoint}/>
            <h2 className="pagename">Containers</h2>
        </header>
    );
};
