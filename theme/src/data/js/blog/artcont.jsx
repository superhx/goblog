import React from 'react'
import ArticleHeader from './arthdr'
import ArticleContent from './art'


export default class ArticleContainer extends React.Component {
  constructor(props) {
    super(props)
  }

  render(){
    return (
      <div className="article-container">
        <ArticleHeader
          title={this.props.title}
          date={this.props.date}
          tags={this.props.tags}
          search={this.props.search}/>
        <ArticleContent content={this.props.content}/>
        <div id="disqus_thread"></div>
      </div>
    )
  }
}
