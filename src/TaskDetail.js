import 'whatwg-fetch';
import React, { Component } from 'react';
import Form from "react-jsonschema-form";

import Navbar from './Navbar';

class TaskDetail extends Component {
  constructor(props) {
    super(props);

    this.state = {};
  }

  handleKeyDown = (e) => {
    switch (e.keyCode) {
      case 27:
        this.props.onExit && this.props.onExit();
        e.preventDefault();
        break;
      default:
        //nothing to do
    }
  }

  handleSubmit = ({ formData }) => {
    const { task: { name } } = this.props;
    fetch(`/api/tasks/${name}/run`, {
      credentials: 'same-origin',
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        data: formData,
      }),
    })
    .then(({ status }) => {
      if (status !== 200) {
        throw new Error("Error running task");
      }

      this.setState({
        lastResult: {
          success: true,
        },
      });
    })
    .catch(() => {
      this.setState({
        lastResult: {
          success: false
        },
      });
    });
  }

  componentDidUpdate = (_prevProps, prevState) => {
    if (this.state.lastResult !== prevState.lastResult) {
      this.ensureResultVisible();
    }
  }

  ensureResultVisible = () => {
    if (this.resultDiv) {
      this.resultDiv.scrollIntoView();
    }
  }

  render() {
    const { lastResult } = this.state;
    const schema = Object.assign({
    }, this.props.task.schema, {
      type: 'object',
      title: null,
    });
    const firstProperty = Object.keys(schema.properties)[0];
    const uiSchema = Object.assign({
      [firstProperty]: {
        "ui:autofocus": true,
      },
    }, this.props.task.uiSchema);
    let result = "";

    console.log(lastResult);

    if (lastResult) {
      if (lastResult.success) {
        result = (
          <div ref={r => this.resultDiv = r} className="alert alert-success">
            Success!
          </div>
        );
      } else {
        result = (
          <div ref={r => this.resultDiv = r} className="alert alert-danger">
            An error occurred
          </div>
        );
      }
    }

    return (
      <div className="task-detail" onKeyDown={this.handleKeyDown}>
        <Navbar />
        <ol className="breadcrumb">
          <li><a href="/">Home</a></li>
          <li className="active">{this.props.task.title}</li>
        </ol>
        <p>{this.props.task.description}</p>
        <Form schema={schema}
          uiSchema={uiSchema}
          onSubmit={this.handleSubmit}
          showErrorList={false}
        />
        {result}
      </div>
    );
  }
}

export default TaskDetail;
