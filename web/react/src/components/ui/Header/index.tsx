import * as React from 'react';
import { AiOutlineMenu } from 'react-icons/ai';

import 'stylesheets/06-components/app-header.scss';

export const Header = () => {
    return (
        <header className="app-header">
            <AiOutlineMenu className="menuicon" />
            <h2 className="pagename">Containers</h2>
        </header>
    );
};
