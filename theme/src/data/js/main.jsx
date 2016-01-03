(function() {
  var MenuButton = require('./blog/menu.jsx');
  var SearchContainer = require('./blog/srchcont.jsx');
  var MainContainer = require('./blog/maincont.jsx');

  var Blog = React.createClass({
    getInitialState: function() {
      return {
        'isMenuOpen': false
      };
    },
    render: function(){
      return (
        <div>
          <MenuButton
            toggleSearchContainer={this.toggleSearchContainer}/>

          <SearchContainer
            ref='searchContainer'
            open={this.state.isMenuOpen}/>

          <MainContainer
            search={this.search}
            offset={this.state.isMenuOpen}
            closeSearchContainer={this.closeSearchContainer}/>
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
    }
  });

  ReactDOM.render(
    <Blog />,
    document.getElementById('container')
  );
}());
