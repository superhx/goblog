var webpack = require('webpack');

module.exports = {
  entry: './src/data/js/main.jsx',
  output: {
    // Output the bundled file.
    path: './data/js',
    // Use the name specified in the entry key as name for the bundle file.
    filename: 'blog.min.js'
  },
  module: {
    loaders: [
      {
        // Test for js or jsx files.
        test: /\.jsx?$/,
        exclude: /node_modules/,
        loader: 'babel',
        query: {
          presets: ['es2015', 'react']
        }
      }
    ]
  },
  externals: {
    // Don't bundle the 'react' npm package with the component.
    'react': 'React',
    'react-dom': 'ReactDOM'
  },
  resolve: {
    // Include empty string '' to resolve files by their explicit extension
    // (e.g. require('./somefile.ext')).
    // Include '.js', '.jsx' to resolve files by these implicit extensions
    // (e.g. require('underscore')).
    extensions: ['', '.js', '.jsx']
  }
};
