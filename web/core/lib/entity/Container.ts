import { Stat } from './Stat';

export class Container {
  constructor(obj: any) {
    this.id = obj.id;
    this.cmd = obj.cmd;
    this.entrypoint = obj.entrypoint;
    this.created = new Date(obj.created);
    this.name = obj.name;
    this.state = obj.state;
    this.status = obj.status;
    this.image = obj.image;
    this.ports = obj.ports;
    this.labels = obj.labels;
    this.volumes = obj.volumes;
    this.networks = obj.networks;
    this.statistics = new Stat(obj.statistics);
  }

  id: string

  cmd: string

  entrypoint: string

  created: Date

  name: string

  state: string

  status: string

  image: string

  ports: Array<string>

  labels: Array<object>

  volumes: Array<string>

  networks: Array<string>

  statistics: Stat
}
