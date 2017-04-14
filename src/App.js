import React, { Component } from 'react';
import fuzzy from 'fuzzy';

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
        {highlights.map(s => s.split('<<', 2)).map(s =>
          <span key={s}>
            <span style={{color: 'red'}}>{s[0]}</span>
            {s[1]}
          </span>
        )}
      </span>
    )
  }
}

class TaskPreview extends Component {
  render() {
    return (
      <li className="task-preview">
        <h3><HighlightText text={this.props.task.title} /></h3>
        <p><HighlightText text={this.props.task.description} /></p>
      </li>
    )
  }
}

class TaskList extends Component {
  constructor(props) {
    super(props);
    this.state = { search: '' };

    this.handleSearchChange = this.handleSearchChange.bind(this);
  }

  handleSearchChange(e) {
    this.setState({ search: e.target.value });
  }

  render() {
    const results = fuzzy
      .filter(this.state.search, this.props.tasks, fuzzyOptions)
      .map(r => {
        const strings = r.string.split(separator);
        return Object.assign({}, r.original, {
          title: strings[0],
          description: strings[1],
        });
      });
    return (
      <div className="task-list">
        <input type="search" value={this.state.search} onChange={this.handleSearchChange} />
        <ol>
          {results.map(task =>
            <TaskPreview key={task.name} task={task} />,
          )}
        </ol>
      </div>
    )
  }
}

class App extends Component {
  render() {
    return (
      <div className="app">
        <TaskList tasks={this.props.tasks} />
      </div>
    );
  }
}

export default App;
