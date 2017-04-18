import 'whatwg-fetch';
import fuzzy from 'fuzzy';
import React, { Component } from 'react';

import TaskDetail from './TaskDetail';
import TaskList from './TaskList';

import './App.css';

const separator = '::';

const fuzzyOptions = {
  pre: '>>',
  post: '<<',
  extract: t => [t.title, t.description].join(separator),
}

function filterTasks(tasks, search) {
  return fuzzy
    .filter(search, tasks, fuzzyOptions)
    .map(r => {
      const strings = r.string.split(separator);
      return Object.assign({}, r.original, {
        highlightTitle: strings[0],
        highlightDescription: strings[1],
      });
    });
}

function getTaskForLocation (tasks, location) {
  const [task] = tasks.filter(({ name }) => name === location.pathname.replace('/', ''));
  return task;
}

class App extends Component {
  constructor(props) {
    super(props);

    const { history } = props;
    const initialLocation = history.location;
    this.unlisten = history.listen(this.handleRouteChange)

    this.state = {
      initialLocation,
      tasks: [],
      filteredTasks: [],
      search: '',
    };
  }

  componentDidMount() {
    fetch('/api/tasks', {
      credentials: 'same-origin',
    })
      .then(r => {
        switch(r.status) {
          case 200:
            return r.json();
          case 401:
            window.location.replace(`/login?redirect_url=${encodeURIComponent(this.props.history.location.pathname)}`);
            throw new Error("Unauthorized");
          default:
            throw new Error(`Unexpected response ${r.status}`);
        }
      })
      .then(r => {
        const tasks = [];
        for (const name in r.tasks) {
          const t = r.tasks[name];
          Object.assign(t, { name });
          tasks.push(t);
        }
        return tasks;
      })
      .then(tasks => {
        const { search, focusedTask, initialLocation } = this.state;
        const [initialTask] = tasks;
        const detailTask = getTaskForLocation(tasks, initialLocation);

        this.setState({
          tasks,
          filteredTasks: filterTasks(tasks, search),
          focusedTask: focusedTask || initialTask,
          detailTask,
        })
      });
  }

  handleFocus = (task) => {
    this.setState({
      focusedTask: task,
    });
  }

  handleSearch = (search) => {
    const filteredTasks = filterTasks(this.state.tasks, search);
    this.setState({
      search,
      filteredTasks,
      focusedTask: filteredTasks[0],
    });
  }

  handleSelect = (task) => {
    const { history } = this.props;
    if (task) {
      history.push(task.name);
    } else {
      history.push('/');
    }
  }

  handleRouteChange = (location) => {
    const { tasks } = this.state;
    const task = getTaskForLocation(tasks, location);
    this.setState({
      search: '',
      focusedTask: task,
      filteredTasks: filterTasks(tasks, ''),
      detailTask: task,
    });
  }

  componentWillUnmount() {
    this.unlisten();
  }

  render() {
    let body;
    if (this.state.detailTask != null) {
      body = <TaskDetail
        task={this.state.detailTask}
        onExit={() => this.handleSelect(null)}
        />;
    } else {
      body = <TaskList
        tasks={this.state.filteredTasks}
        search={this.state.search}
        focusedTask={this.state.focusedTask}
        onFocusTask={this.handleFocus}
        onSelectTask={this.handleSelect}
        onSearch={this.handleSearch}
        />;
    }

    return (
      <div className="app">
        {body}
      </div>
    );
  }
}

export default App;
