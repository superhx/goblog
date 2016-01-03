(function() {
  var ArticleHeader = React.createClass({
    handleTagSearch: function(e){
      e.stopPropagation();
      this.props.search($(e.target).attr('data-tag'));
    },
    render: function(){
      return (
        <header>
          <h1>
            {this.props.title}
          </h1>
          <div>
            <time>
              {this.props.date}
            </time>
            <ul>
              {this.props.tags.map(function(tag, index){
                return (
                  <li key={index} onClick={this.handleTagSearch} data-tag={tag}>
                    {tag}
                  </li>
                );
              }, this)}
            </ul>
          </div>
        </header>
      );
    }
  });

  module.exports = ArticleHeader;
}());
