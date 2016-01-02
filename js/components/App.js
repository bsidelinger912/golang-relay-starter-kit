import React from 'react';
import Relay from 'react-relay';

class App extends React.Component {
  render() {
    return (
      <div>
        <h1>Hello {this.props.viewer.name}</h1>
        <p>
          email: {this.props.viewer.email}
        </p>
        <h3>Past visits</h3>
        <ul>
          {(this.props.viewer.visits.edges || []).map(edge =>
            <Visit data={edge.node} />
          )}
        </ul>
      </div>
    );
  }
}

class Visit extends React.Component {
  render() {
    return (
      <li>{this.props.data.location}</li>
    );
  }
}

export default Relay.createContainer(App, {
  fragments: {
    viewer: () => Relay.QL`
      fragment on User {
        name,
        email,
        visits(first: 10) {
          edges {
            node {
              id,
              location,
            },
          },
        },
      }
    `,
  },
});
