import 'whatwg-fetch';
import React, { Component } from 'react';
import Form from "react-jsonschema-form";

import { assignDeep } from './assignDeep';
import Navbar from './Navbar';
import ResultTable from './ResultTable';

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
    .then(resp => {
      const { status } = resp;

      return resp.json()
        .then(m => {
          m.status = status;
          return m;
        })
    })
    .then(({ status, statements }) => {
      if (status !== 200) {
        throw new Error("Error running task");
      }

      this.setState({
        lastResult: {
          success: true,
          statements,
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
    const schema = assignDeep({
      properties: {},
    }, this.props.task.schema, {
      type: 'object',
      title: null,
    });
    const firstProperty = Object.keys(schema.properties)[0];
    const uiSchema = assignDeep({
      [firstProperty]: {
        "ui:autofocus": true,
      },
    }, this.props.task.uiSchema);
    let result = "";

    if (lastResult) {
      const { success, statements } = lastResult;
      if (success) {
        if (statements) {
          result = (
            <div ref={r => this.resultDiv = r} className="panel panel-success">
              <div className="panel-heading">Success!</div>
              {statements.map((s, i) =>
                <ResultTable key={i} caption={`Statement ${i+1}`} data={s.data} />
              )}
            </div>
          )
        } else {
          result = (
            <div ref={r => this.resultDiv = r} className="alert alert-success">
              Success!
            </div>
          );
        }
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
