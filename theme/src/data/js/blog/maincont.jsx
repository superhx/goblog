(function() {
  var ArticleContainer = require('./artcont.jsx');
  var classNames = require('classnames');

  var MainContainer = React.createClass({
    render: function() {
      return (
        <div className={this.getClassNames()} ref="container" onClick={this.handleClick}>
          <ArticleContainer search={this.props.search}/>
        </div>
      );
    },
    getClassNames: function() {
      classes = {'main-container': true};
      if (this.props.offset) {
        classes['offset-right'] = true;
      }else if($(this.refs.container).hasClass('offset-right')){
        classes['offset-none'] = true;
      }
      return classNames(classes);
    },
    handleClick: function() {
      this.props.closeSearchContainer();
    }
  });

  module.exports = MainContainer;
}());
