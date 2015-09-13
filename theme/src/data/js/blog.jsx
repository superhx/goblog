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


var ArticleHeader = React.createClass({
	render: function(){
		return (
			<div>
				<h1>{this.props.title}</h1>
				<div>
					<time>{this.props.date}</time>
					<ul>
						{this.props.tags.map(function(tag, index){
							return (
								<li key={index}>{tag}</li>
							);
						})}
					</ul>
				</div>
			</div>
		);
	}
});


var ArticleContent = React.createClass({
	render:function(){
		return (
			<div dangerouslySetInnerHTML={{__html: this.props.content}}/>
		);
	}
});


var ArticleContainer = React.createClass({
	componentWillMount: function(){
		var tags = $('b-tag').toArray().map(function(tag){ return $(tag).text(); });;
		this.setState({
			title: $('b-title').text(),
			date: $('b-date').text(),
			tags: tags,
			content: $('b-content').html()
		});
	},
	render: function(){
		return (
			<div>
				<ArticleHeader title={this.state.title} date={this.state.date} tags={this.state.tags}/>
				<ArticleContent content={this.state.content}/>
			</div>
		);
	}
});

React.render(
	<div>
		<SearchContainer />
		<ArticleContainer />
	</div>,
	document.getElementById('container')
);
