export class Byte {
  constructor(private _value: number) {}

  get value(): number {
    return this._value;
  }

  public toHuman(): string {
    let self = this._value;

    if (self < 1024) {
      return `${self}B`;
    }

    let count = 0;

    while (self > 1024) {
      self /= 1024.0;
      count++;
    }

    return `${self.toFixed(2)} ${'KMGTPE'[count - 1]}iB`;
  }
}
