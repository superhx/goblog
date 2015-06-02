$("#toggle").click(
	function(event){
		event.preventDefault(),
		$(this).find(".top").toggleClass("active"),
		$(this).find(".middle").toggleClass("active"),
		$(this).find(".bottom").toggleClass("active"),
		$("#overlay").toggleClass("open"),
		$(".main-content").toggleClass("active")
	}
	),

$(".main-content").click(function(){
	if($(".main-content").attr("class").split(" ").length==2){
			$("#toggle").trigger("click");
	};
});
	(function(){
		var a = $('.content').find('h1,h2').length;
		for(var i = 0;i < a;i++){
			var tag = "tag" + i;
			$(".content h1,.content h2").eq(i).addClass(tag);
			$(".content h1,.content h2").eq(i).attr('id',tag);
			var temp = '<li ><a class="a'+i+'" href="#tag'+i+'">'+$(".content h1,.content h2").eq(i).html()+'</a></li>';
			$(".right-menu-ul").append(temp);
		}
		console.log(a);
	})();
  $('.content').find('h1,h2').each(function() {
    new Waypoint({
      element: this,
      handler: function(direction) {
        var num = '.a'+$(this.element).attr("class").charAt(3);
        $('.right-menu-ul').find('a').filter(".active").each(function(){
        	$(this).removeClass("active");
        })
        $(num).addClass("active");
      },
      continuous: false
    })  
  })


