/*
 * Author: x.H
 */

var SearchBox = React.createClass({
        render: function(){
                return (
			<div className="src-box">
				<span/>
				<input type="text" ref="search" value={this.props.searchValue} onChange={this.props.handleSearch} />
			</div>
		);
        }
});

var SearchResultPanel = React.createClass({
	render: function(){
		return (
			<div className="src-res-list">
				<ul>
					{this.props.blogList.map(function(blog, index){
						return (
							<li key={index}>
								<time>{blog.date}</time>
								<a href={blog.link}>{blog.title}</a>
								<div>
									{blog.tags.map(function(tag, index){
										return (
											<span key={index}>{ tag }</span>
										);
									})}
								</div>
							</li>
						)
					})}
				</ul>
			</div>
		);
	}
});

var SearchContainer = React.createClass({
	getInitialState: function(){
		return {
			searchValue: '',
			blogList: [],
			searchResult: []
		};
	},
	componentDidMount: function(){
		$.ajax({
			url: '/blogs.json',
			dataType: 'json',
			cache: false,
			success: function(data){
				this.setState({blogList: data});
				this.setState({searchResult: this.state.blogList})
			}.bind(this)
		});
	},
	search: function(srch){
		this.setState({'searchValue': srch});
		this.setState({'searchResult': this.state.blogList.filter(function(e, i){
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
	render: function(){
		return (
			<div>
				<SearchBox searchValue={this.state.searchValue} handleSearch={this.handleInputSearch} />
				<SearchResultPanel blogList={this.state.searchResult} />
			</div>
		);
	}
});

React.render(
	<SearchContainer />,
	document.getElementById('container')
);
