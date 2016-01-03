(function() {
  var SearchResultPanel = React.createClass({
    componentWillMount: function() {
      self = this;
      $(window).resize(function() { self.forceUpdate(); });
    },
    render: function(){
      return (
        <ScrollArea
          speed={0.8}
          style={{height:window.innerHeight}}
          className="area"
          contentClassName="content"
          horizontal={false}>
          <div>
            <ul className="srch-res-list">
              {this.props.blogList.map(function(blog, index){
                return (
                  <li key={index}>
                    <datetime>
                      {new Date(blog.date).toString('MMM d yyyy')}
                    </datetime>
                    <a href={blog.link}>
                      {blog.title}
                    </a>
                    <div>
                      {blog.tags.map(function(tag, index){
                        return (
                          <span key={index}>
                            { tag }
                          </span>
                        );
                      })}
                    </div>
                  </li>
                );
              })}
            </ul>
          </div>
        </ScrollArea>
      );
    }
  });

  module.exports = SearchResultPanel;
}());
