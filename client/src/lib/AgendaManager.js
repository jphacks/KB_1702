import _ from 'lodash';

const MOCK_AGENDAS = [
  {
    id: 'aaa',
    title: 'アイデア出し',
    goal: 'アイデアを10個だす',
    time: 10,
    margin: 0,
    child: ['bbb', 'ccc', 'ddd'],
    isParent: true
  },
  {
    id: 'bbb',
    title: 'アイデアうぇい',
    goal: 'アイデアうぇいぇういする',
    time: 10,
    margin: 2,
    child: [],
    isParent: false
  },
  {
    id: 'ccc',
    title: 'アイデアうぇい2',
    goal: 'アイデアうぇいぇういする',
    time: 10,
    margin: 2,
    child: [],
    isParent: false
  },
  {
    id: 'ddd',
    title: 'アイデアうぇい2',
    goal: 'アイデアうぇいぇういする',
    time: 10,
    margin: 2,
    child: [],
    isParent: false
  },
  {
    id: 'eee',
    title: 'eee',
    goal: 'eee',
    time: 10,
    margin: 0,
    child: [],
    isParent: true
  }
];

class AgendaManager {

  constructor() {
    this.data = [];
    this.vm = null;

    // test
    this.data = MOCK_AGENDAS;
  }

  create() {
    const id = (+new Date() + Math.floor(Math.random() * 999999)).toString(36);
    this.data.push({
      id,
      title: '',
      goal: '',
      time: 0,
      margin: 0,
      child: [],
      isParent: true
    })
  }

  child(targetKey) {
    const id = (+new Date() + Math.floor(Math.random() * 999999)).toString(36);
    this.data.map((item) => { if (targetKey === item.id) item.child.push(id) });
    this.data.push({
      id,
      title: '',
      goal: '',
      time: 0,
      margin: 2,
      child: [],
      isParent: false
    })
  }

  destroy(id) {
    let _data = this.data;

    this.data.map((item, index) => {
      if (item.isParent) {
        item.child.forEach((c) => {
          _.remove(_data, (n) => {
            return n.id === c;
          });
        });
      }
      if (item.id === id) this.data.splice(index, 1);
    })
  }

}

export default new AgendaManager;
