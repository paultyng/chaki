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
      <li className="task-preview" onClick={this.handleSelect.bind(this)}>
        <h3><HighlightText text={this.props.task.highlightTitle} /></h3>
        <p><HighlightText text={this.props.task.highlightDescription} /></p>
      </li>
    )
  }
}

export default TaskPreview;
