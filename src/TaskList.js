import React, { Component } from 'react';

import TaskPreview from './TaskPreview';

import './TaskList.css';

class TaskList extends Component {
  handleKeyDown(e) {
    const i = this.props.tasks.indexOf(this.props.focusedTask)
    let delta = 0;
    switch(e.keyCode) {
      case 13: //enter
        this.props.onSelectTask && this.props.onSelectTask(this.props.focusedTask);
        return;
      case 38: //up
        delta = -1;
        break;
      case 40: //down
        delta = 1;
        break;
      default:
        //nothing to do
    }

    if (delta !== 0 && this.props.onFocusTask) {
      const newI = i + delta;
      if (newI >= 0 && newI < this.props.tasks.length) {
        this.props.onFocusTask(this.props.tasks[newI]);
      }
    }
  }

  componentDidMount() {
    this.searchInput.focus();
  }

  render() {
    return (
      <div className="task-list">
        <div className="header">
          <input placeholder="Search tasks"
            ref={r => this.searchInput = r}
            type="search"
            value={this.props.search}
            onKeyDown={this.handleKeyDown.bind(this)}
            onChange={(e) => this.props.onSearch && this.props.onSearch(e.target.value)}
            />
        </div>
        <ol>
          {this.props.tasks.map(task =>
            <TaskPreview
              key={task.name}
              task={task}
              focus={this.props.focusedTask === task}
              onSelect={this.props.onSelectTask} />,
          )}
        </ol>
      </div>
    )
  }
}

export default TaskList;
