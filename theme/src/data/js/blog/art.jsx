(function() {
  var ArticleContent = React.createClass({
    render:function(){
      return (
        <div dangerouslySetInnerHTML={{__html: this.props.content}}/>
      );
    }
  });

  module.exports = ArticleContent;
}());
