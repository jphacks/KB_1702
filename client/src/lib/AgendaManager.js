class AgendaManager {

  constructor() {
    this.data = [];
    this.create();
  }

  create() {
    this.data.push({
      id: this.data.length + 1,
      title: '',
      goal: '',
      time: 10,
      start_at: '',
      end_at: ''
    })
  }

  destroy(id) {
    this.data.map((item, index) => {
      if (item.id === id) this.data.splice(index, 1);
    })
  }

  postFormat(name, agenda) {
    const data = {
      room: {
        room_id: "",
        name,
        progress: 1,
        start_at: "",
        end_at: "",
        agenda
      }
    }
    return data;
  }

}

export default new AgendaManager;
