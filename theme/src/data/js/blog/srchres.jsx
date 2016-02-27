(function() {
  var classNames = require('classnames');

  var SearchResultPanel = React.createClass({
    componentWillMount: function() {
      self = this;
      $(window).resize(function() { self.forceUpdate(); });
    },
    render: function(){
      var searchBoxHeight = $('.srch-box').height();
      return (
        <ScrollArea
          speed={0.8}
          style={{height:window.innerHeight - searchBoxHeight -20}}
          className="area"
          contentClassName="content"
          horizontal={false}
          ref="area">
          <Content { ... this.props } ref="content"/>
        </ScrollArea>
      );
    },
    focusCurrent: function() {
      this.refs.content.focusCurrent();
    }
  });

  var Content = React.createClass({
    render: function(){
      var pathname = decodeURIComponent(window.location.pathname);
      return (
        <div>
          <ul className="srch-res-list" ref="list">
            {this.props.blogList.map(function(blog, index){
              var classes = {'active': blog.link.indexOf(pathname) === 0};
              return (
                <li key={index} className={classNames(classes)}>
                  <datetime>
                    {new Date(blog.date).toString('MMM d yyyy')}
                  </datetime>
                  <a
                    className="title"
                    onClick={this.titleClickHandler}
                    href={blog.link}>
                    {blog.title}
                  </a>
                  <div>
                    {blog.tags.map(function(tag, index){
                      return (
                        <span className="tag" key={index}>
                          { tag }
                        </span>
                      );
                    })}
                  </div>
                </li>
              );
            }, this)}
          </ul>
        </div>
      );
    },
    titleClickHandler: function(e) {
      e.stopPropagation();
      e.preventDefault();
      this.props.loadNewBlog($(e.target).attr('href'));
    },
    focusCurrent: function() {
      var pathname = decodeURIComponent(window.location.pathname);
      $('ul.srch-res-list > li > a').each(function(i, e) {
        if($(e).attr('href').indexOf(pathname) < 0) return;

        var height = $('ul.srch-res-list > li').height();
        console.log(this.context.scrollArea);
        this.context.scrollArea.scrollYTo(height * (i + 1));
      }.bind(this));
    }
  });

  Content.contextTypes = {
    scrollArea: React.PropTypes.object
  };

  module.exports = SearchResultPanel;
}());
