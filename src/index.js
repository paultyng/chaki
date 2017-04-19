import createHistory from 'history/createBrowserHistory';
import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

import './index.css';

ReactDOM.render(
  <App history={ createHistory() } />,
  document.getElementById('root')
);
