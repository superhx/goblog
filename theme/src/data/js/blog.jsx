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
										)
									})}
								</div>
							</li>
						)
					})}
				</ul>
			</div>
		);
	}
})

var SearchContainer = React.createClass({
	getInitialState: function(){
		return {
			searchValue: '',
			blogList: []
		};
	},
	componentDidMount: function(){
		$.ajax({
			url: '/blogs.json',
			dataType: 'json',
			cache: false,
			success: function(data){
				this.setState({blogList: data});
			}.bind(this)
		});
	},
	filterBlogs: function(){
		return this.state.blogList;
	},
	render: function(){
		srcRes = this.filterBlogs();
		return (
			<div>
				<SearchBox searchValue={this.state.searchValue} />
				<SearchResultPanel blogList={srcRes} />
			</div>
		)
	}
});

console.log(287);
React.render(
	<SearchContainer />,
	document.getElementById('container')
);
