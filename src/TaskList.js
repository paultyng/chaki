import React, { Component } from 'react';

import TaskPreview from './TaskPreview';
import Navbar from './Navbar';

class TaskList extends Component {
  handleKeyUp(e) {
    const i = this.props.tasks.indexOf(this.props.focusedTask)
    let delta = 0;
    switch(e.keyCode) {
      case 13: //enter
        this.props.onSelectTask && this.props.onSelectTask(this.props.focusedTask);
        e.preventDefault();
        return;
      case 38: //up
        delta = -1;
        e.preventDefault();
        break;
      case 40: //down
        delta = 1;
        e.preventDefault();
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
        <Navbar>
          <div className="navbar-form navbar-left">
            <div className="form-group">
              <input placeholder="Search tasks"
                className="form-control"
                ref={r => this.searchInput = r}
                type="search"
                value={this.props.search}
                onKeyUp={this.handleKeyUp.bind(this)}
                onChange={(e) => this.props.onSearch && this.props.onSearch(e.target.value)}
                />
              </div>
            </div>
        </Navbar>
        <ol className="list-group">
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
