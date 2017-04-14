import React, { PropTypes } from 'react';

class TaskPreview extends React.Component {
  render() {
    return (
      <div>
        <h2>{this.props.task.title}</h2>
      </div>
    );
  }
}

export default TaskPreview;
