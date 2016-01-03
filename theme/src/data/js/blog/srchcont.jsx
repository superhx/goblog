(function() {
  var classNames = require('classnames');
  var SearchBox = require('./srchbox.jsx');
  var SearchResultPanel = require('./srchres.jsx');

  var SearchContainer = React.createClass({
    getInitialState: function(){
      return {
        searchValue: '',
        blogList: [],
        searchResult: [],
      };
    },
    render: function(){
      return (
        <div
          className={this.getClassNames()}
          ref="container">

          <SearchBox
            searchValue={this.state.searchValue}
            handleSearch={this.handleInputSearch}/>

          <SearchResultPanel
            blogList={this.state.searchResult}
            loadNewBlog={this.props.loadNewBlog}/>

        </div>
      );
    },
    componentDidMount: function(){
      $.ajax({
        url: '/blogs.json',
        dataType: 'json',
        cache: false,
        success: function(data){
          this.setState({blogList: data});
          this.setState({searchResult: this.state.blogList});
        }.bind(this)
      });
    },
    search: function(srch){
      this.setState({'searchValue': srch});
      this.setState({'searchResult': this.state.blogList.filter(function(e){
        srch = srch.toLowerCase();
        if(e.title.toLowerCase().indexOf(srch) >= 0) return true;
        for(var i=0; i < e.tags.length; i++){
          if(e.tags[i].toLowerCase().indexOf(srch) >= 0) return true;
        }
        return false;
      })});
    },
    handleInputSearch: function(e){
      this.search(e.target.value);
    },
    getClassNames: function() {
      classes = { 'search-container': true };
      if (this.props.open) {
        classes.open = true;
        return classNames({ 'search-container': true, 'open': true});
      }else if($(this.refs.container).hasClass('open')){
        classes.close = true;
      }
      return classNames(classes);
    }
  });

  module.exports = SearchContainer;
}());
