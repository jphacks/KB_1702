class AgendaManager {

  constructor() {
    this.data = [];
  }

  create() {
    this.data.push({
      id: this.data.length + 1,
      title: '',
      goal: '',
      time: 0
    })
  }

  destroy(id) {
    this.data.map((item, index) => {
      if (item.id === id) this.data.splice(index, 1);
    })
  }

}

export default new AgendaManager;
