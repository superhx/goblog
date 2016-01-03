(function() {
  var MenuButton = require('./blog/menu.jsx');
  var SearchContainer = require('./blog/srchcont.jsx');
  var MainContainer = require('./blog/maincont.jsx');

  var Blog = React.createClass({
    getInitialState: function() {
      var tags = $('b-tag').toArray().map(function(tag){ return $(tag).text(); });
      return {
        isMenuOpen: false,
        blogContent: $('b-content').html(),
        blogTitle: $('b-title').text(),
        blogDate: $('b-time').text(),
        blogTags: tags
      };
    },
    render: function(){
      return (
        <div>
          <MenuButton
            toggleSearchContainer={this.toggleSearchContainer}/>

          <SearchContainer
            ref='searchContainer'
            open={this.state.isMenuOpen}
            loadNewBlog={this.loadNewBlog}/>

          <MainContainer
            search={this.search}
            offset={this.state.isMenuOpen}
            closeSearchContainer={this.closeSearchContainer}
            content={this.state.blogContent}
            title={this.state.blogTitle}
            date={this.state.blogDate}
            tags={this.state.blogTags}/>
        </div>
      );
    },
    search: function(srch){
      this.refs.searchContainer.search(srch);
      this.openSearchContainer();
    },
    toggleSearchContainer: function() {
      this.setState({isMenuOpen: !this.state.isMenuOpen});
    },
    openSearchContainer: function() {
      if(!this.state.isMenuOpen) this.toggleSearchContainer();
    },
    closeSearchContainer: function() {
      if(this.state.isMenuOpen) this.toggleSearchContainer();
    },
    checkUrlIdentity: function(blogLink) {
      var oldUrl = decodeURIComponent(window.location.pathname);
      return blogLink.indexOf(oldUrl) == 0;
    },
    loadNewBlog: function(blogLink) {
      if(this.checkUrlIdentity(blogLink)) return;
      history.pushState('', '', blogLink);
      $.ajax({
        url: blogLink,
        method: 'GET',
        success: function(html) {
          var blog = $(html);
          var tags = blog.find('b-tag').toArray().map(function(tag){ return $(tag).text(); });
          this.setState({
            blogContent: blog.find('b-content').html(),
            blogTitle: blog.find('b-title').text(),
            blogDate: blog.find('b-time').text(),
            blogTags: tags
          });
        }.bind(this)
      });
    }
  });

  ReactDOM.render(
    <Blog />,
    document.getElementById('container')
  );
}());
