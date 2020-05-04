import * as React from 'react';
import { TiChevronLeft } from 'react-icons/ti';

import 'stylesheets/06-components/app-header.scss';

type Props = {
    sidebarCollapsed: boolean
}

export const SidebarToggleButton = (props: Props) => {
    const [sidebarCollapsed, setSidebarCollapsed] = React.useState<boolean>(props.sidebarCollapsed);

    const doSidebarCollapse = React.useCallback(() => {
        const sidebar = document.getElementsByClassName('app-sidebar');
        if (!sidebar || sidebar.length === 0) {
            return;
        }

        if (sidebarCollapsed) {
            sidebar[0].classList.add('-collapsed');
        } else {
            sidebar[0].classList.remove('-collapsed');
        }
    }, [sidebarCollapsed]);

    // Runs whenever sidebarCollapsed is changed, and executes
    // doSidebarCollapse
    React.useEffect(() => {
        doSidebarCollapse();
    }, [sidebarCollapsed, doSidebarCollapse]);

    // Triggered when the parent prop changes (window width is below a certain threshold)
    // and updates the sidebarCollapsed with the value
    React.useEffect(() => {
        setSidebarCollapsed(props.sidebarCollapsed);
    }, [props.sidebarCollapsed]);

    // Triggered when the button is clicked
    const sidebarToggle = (e: React.MouseEvent) => {
        setSidebarCollapsed(!sidebarCollapsed);
        doSidebarCollapse();
    };

    return (
        <div className={"sidebar-togglebutton " + (sidebarCollapsed ? '-collapsed' : '')} onClick={sidebarToggle}>
            <TiChevronLeft className="icon" />
        </div>
    );
};
