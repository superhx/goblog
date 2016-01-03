(function() {
  var ArticleHeader = require('./arthdr.jsx');
  var ArticleContent = require('./art.jsx');

  var ArticleContainer = React.createClass({
    componentWillMount: function(){
      var tags = $('b-tag').toArray().map(function(tag){ return $(tag).text(); });
      this.setState({
        title: $('b-title').text(),
        date: $('b-time').text(),
        tags: tags,
        content: $('b-content').html()
      });
    },
    render: function(){
      return (
        <div className="article-container">
          <ArticleHeader
            title={this.state.title}
            date={this.state.date}
            tags={this.state.tags}
            search={this.props.search}/>
          <ArticleContent content={this.state.content}/>
        </div>
      );
    }
  });

  module.exports = ArticleContainer;
}());
