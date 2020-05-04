import axios from 'axios';
import { plainToClass } from 'class-transformer';
import { Container } from '../entity';

export class ContainerService {
  static async getContainers(): Promise<Array<Container>> {
    const containers: Array<Container> = [];

    await axios.get('http://localhost:5000/v1/containers')
      .then((response) => { // @TODO: Use a constant source for the API
        if (response.data) {
          for (let i = 0; i < response.data.length; i++) {
            const container = plainToClass(Container, response.data[i]);
            containers.push(container);
          }
        }
      });

    return containers;
  }
}
