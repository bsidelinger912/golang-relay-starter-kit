var getbabelRelayPlugin = require('babel-relay-plugin');
var schema = require('../models/schema.json');

module.exports = getbabelRelayPlugin(schema.data);
