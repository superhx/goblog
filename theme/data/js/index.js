$.getJSON("./category.json", function(data) {

});
var jsonData = [{
    "date": "2012-01-10T00:00:00Z",
    "link": "2012/1/10/markdown_help_9/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-09T00:00:00Z",
    "link": "2012/1/9/markdown_help_8/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-09T00:00:00Z",
    "link": "2012/1/9/markdown_help_8/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
},{
    "date": "2012-01-09T00:00:00Z",
    "link": "2012/1/9/markdown_help_8/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
},{
    "date": "2012-01-09T00:00:00Z",
    "link": "2012/1/9/markdown_help_8/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
},{
    "date": "2012-01-09T00:00:00Z",
    "link": "2012/1/9/markdown_help_8/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
},{
    "date": "2012-01-09T00:00:00Z",
    "link": "2012/1/9/markdown_help_8/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
},{
    "date": "2012-01-09T00:00:00Z",
    "link": "2012/1/9/markdown_help_8/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
},{
    "date": "2012-01-08T00:00:00Z",
    "link": "2012/1/8/markdown_help_7/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-07T00:00:00Z",
    "link": "2012/1/7/markdown_help_6/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-07T00:00:00Z",
    "link": "2012/1/7/markdown_help_6/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-07T00:00:00Z",
    "link": "2012/1/7/markdown_help_6/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-07T00:00:00Z",
    "link": "2012/1/7/markdown_help_6/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-07T00:00:00Z",
    "link": "2012/1/7/markdown_help_6/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-06T00:00:00Z",
    "link": "2012/1/6/markdown_help_5/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-05T00:00:00Z",
    "link": "2012/1/5/markdown_help_4/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-04T00:00:00Z",
    "link": "2012/1/4/markdown_help_3/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-03T00:00:00Z",
    "link": "2012/1/3/markdown_help_2/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-02T00:00:00Z",
    "link": "2012/1/2/markdown_help_1/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}, {
    "date": "2012-01-01T00:00:00Z",
    "link": "2012/1/1/markdown_help_0/index.html",
    "tags": ["java", "golang"],
    "title": "Markdown help"
}];
(function(data) {
    $.each(data, function(index, val) {
    	var date = Date.parse(val.date);
    	var tag = '';
    	for (var i = val.tags.length - 1; i >= 0; i--) {
    		tag += '<a class="node" href="#">'+val.tags[i]+'</a>'
    	};
      var string='<li class="menu-li"><span class="time">'+format(date)+'</span><div class="content"><a class="title" href="'+val.link+'">' +val.title+'</a><div class="tags"><span>'+tag+'</span></div></div></li>'
    	$('.ul-div ul').append(string);
    });
})(jsonData);

(function(){
    var a = /\b.*(iphone|ipad|android).*\b/;
    if(!(a.test(navigator.userAgent.toLowerCase()))){

    }
}());


var toggleBlog = function(){
    // var localhref = window.location.href;
    var localhref = 'localhost://blog/2012/1/9/markdown_help_8/index.html'
    var blogList = $(".menu-li");
    var blogTop , num;
    for (var i = 0; i < blogList.length; i++) {
        if (localhref.toLowerCase().indexOf(totalList.eq(i).find('a').attr('href').toLowerCase())>= 0) {
            blogTop = totalList.eq(i).offset();
            num = i;
            break;
        }
    }
    return {
        getTop : function(){return blogTop.top;},
        getNum : function(){return num;}
    };
}

var rightTop = parseInt($('.right-menu').offset().top);
 $(document).scroll(function (e) {
    var pos = document.body.scrollTop;
    if(pos >= rightTop){
        $('.right-menu').addClass('fixed');
    }else{
        $('.right-menu').removeClass('fixed');
        $('.right-menu-ul a').removeClass('active');
                $('.right-menu-ul ul').removeClass('active');
                $('.right-menu-ul li ul li').removeClass('active');
    }
 });
$("#toggle").click(
        function(event) {
            event.preventDefault();
            if($('.search').val() == ''){
                $('.ul-div').scrollTop(toggleBlog().getTop());
                $('.menu-li').eq(toggleBlog().getNum()).addClass('active');
            }
                $(this).find(".top").toggleClass("active");
                $(this).find(".middle").toggleClass("active");
                $(this).find(".bottom").toggleClass("active");
                $(".main-content").toggleClass("active");
                $("#overlay").toggleClass("open");
                $(".right-menu").toggleClass("active");
        }
    ),

    $(".main-content").click(function() {
        if ($(".main-content").attr("class").split(" ").length == 2) {
            $("#toggle").trigger("click");
        };
    });
(function() {
    var getEle = $('.content').find('h1,h2');
    var a = getEle.length;
    for (var i = 0; i < a;) {
        var tag = "tag" + i;
        getEle.eq(i).addClass(tag);
        getEle.eq(i).attr('id', tag);
        var temp = '<li ><a class="a' + i + '" href="#tag' + i + '">' + getEle.eq(i).html() + '</a>   <ul class="nav">';
        i = i + 1;
        while ((i < a) && (getEle.eq(i).is('h2'))) {
            var tag = "tag" + i;
            getEle.eq(i).addClass(tag);
            getEle.eq(i).attr('id', tag);
            temp += '<li> <a class = "a' + i + ' sub" href = "#tag' + i + '" > ' + getEle.eq(i).html() + ' </a> </li>';
            i++;
        }
        temp += '</ul></li>';
        $(".right-menu-ul").append(temp);
    }

    $('.content').find('h1,h2,article').each(function() {
        new Waypoint({
            element: this,
            handler: function(direction) {
                var num = '.a' + $(this.element).attr("class").charAt(3);
                $('.right-menu-ul a').removeClass('active');
                $('.right-menu-ul ul').removeClass('active');
                $('.right-menu-ul li ul li').removeClass('active');
                if ($(num).hasClass("sub")) {
                    $(num).parent().parent().addClass("active");
                    $(num).parent().parent().siblings("a").addClass('active');
                }
                $(num).addClass("active");
                $(num).siblings("ul").addClass("active");

            },
            continuous: false
        })
    })

})();

var totalList = $(".menu-li");
var input = $('.search');
input.bind('keyup', sendKeyWord);

function sendKeyWord(e) {
    dealData($.trim(input.val()));
}

function dealData(text) {
    for (var i = 0; i < totalList.length; i++) {
        if (totalList.eq(i).text().toLowerCase().indexOf(text.toLowerCase()) < 0) {
            totalList.eq(i).hide();
        } else {
            totalList.eq(i).show();
        }
    }
    $('.ul-div').scrollTop(0);
}


var scroll = function(event, scroller) {
    var k = event.wheelDelta ? event.wheelDelta : -event.detail * 10;
    scroller.scrollTop = scroller.scrollTop - k;
    return false;
};

var isUp=true;
var lastScrollTop=0;
var searchDiv=$('.search-div')
var topThreshold=searchDiv.height();
$('.ul-div').scroll(function(e){
  var top=$(e.delegateTarget).scrollTop();
  if(top<topThreshold) return;
  if(isUp^(top<lastScrollTop)){
    console.log(top);
    searchDiv.fadeToggle();
  }
  isUp=top<lastScrollTop;
  lastScrollTop=top;
});

// var isUp=false;
// var lastScrollTop=0;
// var searchDiv=$('.search-div')
// var topThresh=searchDiv.height();
// function searchScroll(top){
//   console.log(top);
//   if(top<topThreshold) return;
//   if(isUp^(top>topThreshold)) searchDiv.fadeToggle();
//   isUp=top<lastScrollTop;
//   lastScrollTop=top;
// }



$('.ul-div').perfectScrollbar();
$('.node').on('click', searchNode);

function searchNode(e) {
    if ($(".main-content").attr("class").split(" ").length == 1) {
        $("#toggle").trigger("click");
    };
    e.stopPropagation();
    $('.search').val($(e.target).html());
    dealData($(e.target).html());
}

$('.menu-li').mouseenter(function(e){
  var target=$(e.delegateTarget);
  setTimeout(function(){
    if(target.is(':hover')){
      target.find('div.tags').animate({'margin-top':'5px'},400);
    }
  },500);
  // $(e.delegateTarget).find('div.tags').animate({'margin-top':'5px'},200);
}).mouseleave(function(e){
  $(e.delegateTarget).find('div.tags').animate({'margin-top':'-25px'},200);
});

function format(date){
  return date.toString('MMM d');
}
