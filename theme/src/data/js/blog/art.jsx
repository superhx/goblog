import React from 'react'

export default class ArticleContent extends React.Component{
  constructor(props) {
    super(props)
  }

  render() {
    return (
      <div dangerouslySetInnerHTML={{__html: this.props.content}}/>
    )
  }
}
