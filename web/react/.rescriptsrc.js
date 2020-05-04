const ExtraWatchWebpackPlugin = require('extra-watch-webpack-plugin');
const { appendWebpackPlugin } = require("@rescripts/utilities");

const path = require('path');

module.exports = [
  ['use-babel-config', '.babelrc'],
  ['use-eslint-config', '.eslintrc.js'],

  appendWebpackPlugin(
    new ExtraWatchWebpackPlugin({
      dirs: [ path.resolve('../core/dist') ]
    })
  )
];
