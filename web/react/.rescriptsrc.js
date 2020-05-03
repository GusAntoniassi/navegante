const ExtraWatchWebpackPlugin = require('extra-watch-webpack-plugin');
const path = require('path');

module.exports = [
  ['use-babel-config', '.babelrc'],
  ['use-eslint-config', '.eslintrc.js'],

  // add the 'core' folder to Webpack's watch mode
  config => {
    config.plugins.push(new ExtraWatchWebpackPlugin({
      files: [ path.resolve('../core/index.ts') ],
      dirs: [ path.resolve('../core/lib') ]
    }));
    return config;
  }
];
