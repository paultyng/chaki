import React, { Component } from 'react';
import fuzzy from 'fuzzy';
import Form from "react-jsonschema-form";

import './App.css';

const separator = '::';

const fuzzyOptions = {
  pre: '>>',
  post: '<<',
  extract: t => [t.title, t.description].join(separator),
}

class HighlightText extends Component {
  render() {
    const highlights = this.props.text.split('>>');
    const first = highlights.shift();
    return (
      <span>
        {first}
        {highlights.map(s => s.split('<<', 2)).map((s, i) =>
          <span key={i}>
            <span style={{color: 'red'}}>{s[0]}</span>
            {s[1]}
          </span>
        )}
      </span>
    )
  }
}

class TaskPreview extends Component {
  handleSelect(e) {
    if (this.props.onSelect) {
      this.props.onSelect(this.props.task);
    }
  }

  render() {
    return (
      <li className="task-preview" onClick={this.handleSelect.bind(this)}>
        <h3><HighlightText text={this.props.task.highlightTitle} /></h3>
        <p><HighlightText text={this.props.task.highlightDescription} /></p>
      </li>
    )
  }
}

class TaskList extends Component {
  handleSearchChange(e) {
    this.setState({ search: e.target.value });
  }

  render() {
    const search = this.state && this.state.search ? this.state.search : '';
    const results = fuzzy
      .filter(search, this.props.tasks, fuzzyOptions)
      .map(r => {
        const strings = r.string.split(separator);
        return Object.assign({}, r.original, {
          highlightTitle: strings[0],
          highlightDescription: strings[1],
        });
      });
    return (
      <div className="task-list">
        <div className="header">
          <input placeholder="Search tasks" type="search" value={search} onChange={this.handleSearchChange.bind(this)} />
        </div>
        <ol>
          {results.map(task =>
            <TaskPreview key={task.name} task={task} onSelect={this.props.onSelect} />,
          )}
        </ol>
      </div>
    )
  }
}

class TaskDetail extends Component {
  handleExit() {
    this.props.onExit && this.props.onExit();
  }

  render() {
    const schema = Object.assign({
    }, this.props.task.schema, {
      type: 'object',
      title: null,
    });
    console.log(schema);
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

class App extends Component {
  handleSelect(task) {
    this.setState({ currentTask: task });
  }

  render() {
    let body = <TaskList tasks={this.props.tasks} onSelect={this.handleSelect.bind(this)} />;

    if (this.state != null && this.state.currentTask != null) {
      body = <TaskDetail task={this.state.currentTask} onExit={() => this.handleSelect(null)} />;
    }

    return (
      <div className="app">
        {body}
      </div>
    );
  }
}

export default App;
