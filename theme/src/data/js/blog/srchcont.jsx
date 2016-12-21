import React from 'react'
import classNames from 'classnames'
import SearchBox from './srchbox'
import SearchResultPanel from './srchres'


export default class SearchContainer extends React.Component {
  constructor(props) {
    super(props)
    this.setInitialState();

    this.handleInputSearch = this.handleInputSearch.bind(this)
  }

  setInitialState() {
    this.state = {
      searchValue: '',
      blogList: [],
      searchResult: [],
    } 
  }

  render(){
    return (
      <div
        className={this.getClassNames()}
        ref="container">

        <SearchBox
          searchValue={this.state.searchValue}
          handleSearch={this.handleInputSearch}/>

        <SearchResultPanel
          blogList={this.state.searchResult}
          loadNewBlog={this.props.loadNewBlog}
          ref="resPanel"/>

      </div>
    )
  }

  componentDidMount(){
    $.ajax({
      url: '/blogs.json',
      dataType: 'json',
      cache: false,
      success: function(data){
        this.setState({blogList: data})
        this.setState({searchResult: this.state.blogList}, function() {
          this.refs.resPanel.focusCurrent()
        })
      }.bind(this)
    })
  }

  search(srch){
    this.setState({'searchValue': srch})
    this.setState({'searchResult': this.state.blogList.filter(function(e){
      let keyword = srch.toLowerCase()
      if(e.title.toLowerCase().indexOf(keyword) >= 0) return true
      for(let i=0; i < e.tags.length; i++){
        if(e.tags[i].toLowerCase().indexOf(keyword) >= 0) return true
      }
      return false
    })})
  }

  handleInputSearch(e){
    this.search(e.target.value)
  }

  getClassNames() {
    let classes = { 'search-container': true }
    if (this.props.open) {
      classes.open = true
      return classNames({ 'search-container': true, 'open': true})
    }else if($(this.refs.container).hasClass('open')){
      classes.close = true
    }
    return classNames(classes)
  }
}
