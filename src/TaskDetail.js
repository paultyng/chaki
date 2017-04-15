import React, { Component } from 'react';
import Form from "react-jsonschema-form";

class TaskDetail extends Component {
  handleExit() {
    this.props.onExit && this.props.onExit();
  }

  render() {
    const formData = Object.assign({}, this.props.task.defaultData);
    const uiSchema = Object.assign({}, this.props.task.uiSchema);
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
        <Form schema={schema}
          uiSchema={uiSchema}
          formData={formData} />
      </div>
    );
  }
}

export default TaskDetail;
