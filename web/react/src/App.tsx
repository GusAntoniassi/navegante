import React from 'react';
import { Header } from 'components/ui/Header';
import { Sidebar } from 'components/ui/Sidebar';

import 'stylesheets/app.scss';
import 'stylesheets/06-components/app-content.scss';

import { ContainerList } from 'components/containers/ContainerList';

function App() {
  return (
    <>
      <Sidebar/>
      <div className="app-content">
        <Header/>
        <ContainerList/>
      </div>
    </>
  );
}

export default App;
