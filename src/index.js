import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

import 'normalize.css/normalize.css';
import './index.css';

const tasks = [
  {
    "name": "update-user-email",
    "title": "Update User's Email",
    "description": "Update a user's email address.",
    "schema": {
      "properties": {
        "id": {
          "title": "User ID",
          "type": "integer"
        },
        "email": {
          "title": "Email",
          "pattern": "@",
          "type": "string"
        }
      }
    },
    "uiSchema": {}
  },
  {
    "name": "update-order-status",
    "title": "Update Order Status",
    "description": "Change an order's status.",
    "schema": {
      "properties": {
        "number": {
          "title": "Order Number",
          "type": "string",
          "pattern": "[0-9]+"
        },
        "status": {
          "title": "Status",
          "type": "string",
          "enum": [
            "shipped",
            "delivered",
            "cancelled"
          ],
          "default": "shipped"
        }
      }
    },
    "uiSchema": {}
  }
]

ReactDOM.render(
  <App tasks={tasks} />,
  document.getElementById('root')
);
