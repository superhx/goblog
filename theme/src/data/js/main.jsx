import MenuButton from './blog/menu'
import SearchContainer from './blog/srchcont'
import MainContainer from './blog/maincont'
import React from 'react'
import ReactDom from 'react-dom'

class Blog extends React.Component{

  constructor(props) {
    super(props)
    this.setInitialState()

    this.toggleSearchContainer = this.toggleSearchContainer.bind(this)
    this.closeSearchContainer = this.closeSearchContainer.bind(this)
    this.loadNewBlog = this.loadNewBlog.bind(this)
    this.search = this.search.bind(this)
  }

  setInitialState() {
    let tags = $('b-tag').toArray().map(function(tag){ return $(tag).text() })
    this.state = {
      isMenuOpen: false,
      blogContent: $('b-content').html(),
      blogTitle: $('b-title').text(),
      blogDate: $('b-time').text(),
      blogTags: tags
    }
  }

  render(){
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
    )
  }

  componentDidMount() {
    hljs.initHighlightingOnLoad()
  }

  search(srch){
    this.refs.searchContainer.search(srch)
    this.openSearchContainer()
  }

  toggleSearchContainer() {
    this.setState({isMenuOpen: !this.state.isMenuOpen})
  }

  openSearchContainer() {
    if(!this.state.isMenuOpen) this.toggleSearchContainer()
  }

  closeSearchContainer() {
    if(this.state.isMenuOpen) this.toggleSearchContainer()
  }

  checkUrlIdentity(blogLink) {
    let oldUrl = decodeURIComponent(window.location.pathname)
    return blogLink.indexOf(oldUrl) === 0
  }

  loadNewBlog(blogLink) {
    if(this.checkUrlIdentity(blogLink)) return
    history.pushState('', '', blogLink)
    $.ajax({
      url: blogLink,
      method: 'GET',
      success: html => {
        let blog = $(html)
        let tags = blog.find('b-tag').toArray().map(function(tag){ return $(tag).text() })
        this.setState({
          blogContent: blog.find('b-content').html(),
          blogTitle: blog.find('b-title').text(),
          blogDate: blog.find('b-time').text(),
          blogTags: tags
        }, function() {
          $('html > head > title').text(this.state.blogTitle)
          $('pre code').each(function(i, block) { hljs.highlightBlock(block) })
        })
        this.scrollTopSmooth()
        this.reloadComment()
      }
    })
  }

  scrollTopSmooth() {
    $('body').animate({ scrollTop: 0 }, 500)
  }

  reloadComment() {
    DISQUS.reset({
        reload: true,
        config: function() {
          this.page.identifier = window.location.pathname
        }
    })
  }

}


ReactDOM.render(
  <Blog />,
  document.getElementById('container')
)
