import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';

const tasks = [
  {
    "name": "update-user-email",
    "title": "Update User's Email",
    "description": "Update a user's email address.",
    "schema": {
      "properties": {
        "id": {
          "type": "integer"
        },
        "email": {
          "pattern": "@",
          "type": "string"
        }
      }
    }
  },
  {
    "name": "update-order-status",
    "title": "Update Order Status",
    "description": "Change an order's status.",
    "schema": {
      "properties": {
        "number": {
          "type": "string",
          "pattern": "[0-9]+"
        },
        "status": {
          "type": "string",
          "enum": [
            "shipped",
            "delivered",
            "cancelled"
          ]
        }
      }
    }
  }
]

ReactDOM.render(
  <App tasks={tasks} />,
  document.getElementById('root')
);
