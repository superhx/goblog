(function() {
  var MenuButton = React.createClass({
    render: function() {
      return (
        <div className="menu-button" onClick={this.props.toggleSearchContainer}>
          <span></span>
          <span></span>
          <span></span>
        </div>
      );
    }
  });

  module.exports = MenuButton;
}());
