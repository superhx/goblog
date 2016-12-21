import React from 'react'
import ArticleContainer from './artcont'
import classNames from 'classnames'


export default class MainContainer extends React.Component {
  constructor(props) {
    super(props)

    this.handleClick = this.handleClick.bind(this)
  }

  render() {
    return (
      <div className={this.getClassNames()} ref="container" onClick={this.handleClick}>
        <ArticleContainer {...this.props}/>
      </div>
    )
  }

  getClassNames() {
    let classes = {'main-container': true}
    if (this.props.offset) {
      classes['offset-right'] = true
    }else if($(this.refs.container).hasClass('offset-right')){
      classes['offset-none'] = true
    }
    return classNames(classes)
  }

  handleClick() {
    this.props.closeSearchContainer()
  }
}
