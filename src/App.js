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

    this.state = {
      filteredTasks: filterTasks(this.props.tasks, search),
      focusedTask: this.props.tasks[0],
      search,
      detailTask: null,
    };
  }

  handleFocus(task) {
    this.setState({
      focusedTask: task,
    });
  }

  handleSearch(search) {
    const filteredTasks = filterTasks(this.props.tasks, search);

    this.setState({
      search,
      filteredTasks,
      focusedTask: filteredTasks[0],
    });
  }

  handleSelect(task) {
    this.setState({
      search: '',
      detailTask: task,
      focusedTask: task,
    });
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
        onFocusTask={this.handleFocus.bind(this)}
        onSelectTask={this.handleSelect.bind(this)}
        onSearch={this.handleSearch.bind(this)}
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
