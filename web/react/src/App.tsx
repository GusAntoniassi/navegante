import React from 'react';
import { Header } from './components/ui/Header';
import { Sidebar } from './components/ui/Sidebar';
import { Container } from './components/containers/Container';

import './stylesheets/app.scss';
import './stylesheets/06-components/app-content.scss';
import './stylesheets/06-components/container-list.scss';

import { Container as EntityContainer } from 'core';
const container = new EntityContainer();
container.id = "abc123456";
container.name = "foobar";
container.ports = ["80:8000/TCP"];
container.volumes = ["/var/lib/data", "/home/foo/bar:/usr/local/application"];

function App() {
  return (
    <>
      <Sidebar/>
      <div className="app-content">
        <Header/>
        <div className="container container-list">
          <Container container={container}/>
          <Container container={container}/>
          <Container container={container}/>
          <Container container={container}/>
          <Container container={container}/>
        </div>
      </div>
    </>
  );
}

export default App;
