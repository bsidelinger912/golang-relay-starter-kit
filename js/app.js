import 'babel/polyfill';

import App from './components/App';
import AppHomeRoute from './routes/AppHomeRoute';
import React from 'react';
import ReactDOM from 'react-dom';
import Relay from 'react-relay';
require('./app.scss');

ReactDOM.render(
  <Relay.RootContainer
    Component={App}
    route={new AppHomeRoute({ userId: '1' })}
  />,
  document.getElementById('root')
);
