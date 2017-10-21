class FormatUtil {

  zeroSuppress(number) {
    return ('00' + number).slice(-2);
  }

}

export default new FormatUtil;
