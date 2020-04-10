export class Container {
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
}
