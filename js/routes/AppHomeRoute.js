import Relay from 'react-relay';

export default class extends Relay.Route {
  static queries = {
    viewer: () => Relay.QL`
      query {
        viewer(id: $userId)
      }
    `,
  };

  static paramDefinitions = {
    userId: {required: true},
  };

  static routeName = 'AppHomeRoute';
}
