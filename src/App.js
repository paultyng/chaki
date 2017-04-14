import React, { Component } from 'react';
import Form from "react-jsonschema-form";
import fuzzy from 'fuzzy';

import TaskPreview from './TaskPreview';

import './App.css';

class TaskList extends Component {
  render() {
    return (
      <div className="task-list">
        <div className="header">
          <input placeholder="Search tasks" type="search" value={this.props.search} onChange={(e) => this.props.onSearch && this.props.onSearch(e.target.value)} />
        </div>
        <ol>
          {this.props.tasks.map(task =>
            <TaskPreview key={task.name} task={task} onSelect={this.props.onSelectTask} />,
          )}
        </ol>
      </div>
    )
  }
}

class TaskDetail extends Component {
  // handleExit() {
  //   this.props.onExit && this.props.onExit();
  // }

  render() {
    const schema = Object.assign({
    }, this.props.task.schema, {
      type: 'object',
      title: null,
    });
    return (
      <div className="task-detail">
        <div className="header">
          <h1>{this.props.task.title}</h1>
        </div>
        <p>{this.props.task.description}</p>
        <Form schema={schema} />
      </div>
    );
  }
}

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
