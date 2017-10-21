class Logger {

  log(...args) {
    console.log.apply(console, [`[LOG] `, ...args]);
  }

  info(...args) {
    console.info.apply(console, [`[INFO] `, ...args]);
  }

  warn(...args) {
    console.warn.apply(console, [`[WARN] `, ...args]);
  }

  error(...args) {
    console.error.apply(console, [`[ERROR] `, ...args]);
  }

  biglog(...args) {
    const text = `%c${args.join(' ')}`;
    console.log.apply(console, [text, 'font-size: 20px']);
  }

  cycle(...args) {
    const text = `%c[LIFE CYCLE]\n${args.join(' ')}`;
    console.log.apply(console, [text, 'color: #bbb;font-size: 10px']);
  }

}

export default new Logger;
