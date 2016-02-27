(function() {
  var ArticleHeader = require('./arthdr.jsx');
  var ArticleContent = require('./art.jsx');

  var ArticleContainer = React.createClass({
    render: function(){
      return (
        <div className="article-container">
          <ArticleHeader
            title={this.props.title}
            date={this.props.date}
            tags={this.props.tags}
            search={this.props.search}/>
          <ArticleContent content={this.props.content}/>
          <div id="disqus_thread"></div>
        </div>
      );
    }
  });

  module.exports = ArticleContainer;
}());
