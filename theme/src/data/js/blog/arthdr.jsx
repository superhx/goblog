import React from 'react'


export default class ArticleHeader extends React.Component {
  constructor(props) {
    super(props)

    this.handleTagSearch = this.handleTagSearch.bind(this)
  }


  handleTagSearch(e){
    e.stopPropagation()
    this.props.search($(e.target).attr('data-tag'))
  }

  render(){
    return (
      <header>
        <h1>
          {this.props.title}
        </h1>
        <div>
          <time>
            {this.props.date}
          </time>
          <ul>
            {this.props.tags.map(function(tag, index){
              return (
                <li key={index} onClick={this.handleTagSearch} data-tag={tag}>
                  {tag}
                </li>
              )
            }, this)}
          </ul>
        </div>
      </header>
    )
  }
}
