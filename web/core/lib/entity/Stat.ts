import { Byte } from './Byte';

export class Stat {
  constructor(obj: any) {
    this.cpuPercent = obj.cpuPercent;
    this.memoryPercent = obj.memoryPercent;
    this.memoryUsage = new Byte(obj.memoryUsage);
    this.memoryTotal = new Byte(obj.memoryTotal);
    this.networkInput = new Byte(obj.networkInput);
    this.networkOutput = new Byte(obj.networkOutput);
    this.blockRead = new Byte(obj.blockRead);
    this.blockWrite = new Byte(obj.blockWrite);
  }

  cpuPercent: number

  memoryPercent: number

  memoryUsage: Byte

  memoryTotal: Byte

  networkInput: Byte

  networkOutput: Byte

  blockRead: Byte

  blockWrite: Byte
}
