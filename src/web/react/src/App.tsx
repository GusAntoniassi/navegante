import React from 'react';
import { Header } from './components/ui/Header';
import { Sidebar } from './components/ui/Sidebar';
import { Container } from './components/containers/Container';

import { Container as EntityContainer } from 'core';
const container = new EntityContainer();
container.id = "abc123456";
container.name = "foobar";
container.ports = ["80:8000/TCP"];
container.volumes = ["/var/lib/data", "/home/foo/bar:/usr/local/application"];

function App() {
  return (
    <div className="App">
      <Header/>
      <Sidebar/>
      <Container container={container}/>
    </div>
  );
}

export default App;
