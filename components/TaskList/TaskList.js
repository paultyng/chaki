import React, { PropTypes } from 'react';

import TaskPreview from './TaskPreview';

class TaskList extends React.Component {
  render() {
    return (
      <div>
        {this.props.tasks.map(task =>
          <TaskPreview key={task.name} task={task} />,
        )}
      </div>
    );
  }
}

export default TaskList;
