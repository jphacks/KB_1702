class FormatUtil {

  timer(second) {
    const sec = Math.trunc(second % 60);
    const min = Math.trunc(second / 60);
    const hour = Math.trunc(second / 3600);
    return `${this.zeroSuppress(hour)}:${this.zeroSuppress(min)}:${this.zeroSuppress(sec)}`;
  }

  zeroSuppress(number) {
    return ('00' + number).slice(-2);
  }

}

export default new FormatUtil;
