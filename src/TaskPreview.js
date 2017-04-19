import React, { Component } from 'react';

import HighlightText from './HighlightText';

class TaskPreview extends Component {
  handleSelect(e) {
    if (this.props.onSelect) {
      this.props.onSelect(this.props.task);
    }
  }

  render() {
    return (
      <li className={`task-preview list-group-item ${this.props.focus ? 'active' : ''}`} onClick={this.handleSelect.bind(this)}>
        <h3 className="list-group-item-heading"><HighlightText text={this.props.task.highlightTitle} /></h3>
        <p className="list-group-item-text"><HighlightText text={this.props.task.highlightDescription} /></p>
      </li>
    )
  }
}

export default TaskPreview;
