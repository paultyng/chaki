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

class App extends Component {
  constructor(props) {
    super(props);

    const search = '';

    const initialLocation = props.history.location;
    const [initialTask] = props.tasks;
    
    const detailTask = this.getTaskForLocation(initialLocation);
    this.unlisten = props.history.listen(this.handleRouteChange)

    this.state = {
      filteredTasks: filterTasks(props.tasks, search),
      focusedTask: initialTask,
      search,
      detailTask,
    };
  }

  handleFocus = (task) => {
    this.setState({
      focusedTask: task,
    });
  }

  handleSearch = (search) => {
    const filteredTasks = filterTasks(this.props.tasks, search);
    this.setState({
      search,
      filteredTasks,
      focusedTask: filteredTasks[0],
    });
  }

  handleSelect = ({ name }) => {
    const { history } = this.props;
    history.push(name);
  }

  handleRouteChange = (location) => {
    const task = this.getTaskForLocation(location);
    const { tasks } = this.props;
    this.setState({
      search: '',
      focusedTask: task,
      filteredTasks: filterTasks(tasks, ''),
      detailTask: task,
    });
  }

  getTaskForLocation = (location) => {
    const { tasks } = this.props;
    const [task] = tasks.filter(({ name }) => name === location.pathname.replace('/', ''));
    return task;
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
