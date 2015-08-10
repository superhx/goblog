var a=/\b.*(iphone|ipad|android).*\b/;
if(a.test(navigator.userAgent.toLowerCase())){
	$('.dashboard').css('height','1000px');
	$('.dashboard').css('width','200px');
	$('.dashboard').css('font-size','150px');
	$('.node').css('width','200px');
	$('.node').css('height','200px');
	$('.node').css('border-radius','100px');
	$(".node").eq(2).animate({
	top :"700px",
	opacity: "1"},
	500);
setTimeout(function(){
	$(".node").eq(1).animate({
	top :"400px",
	opacity: "1"},
	300);
},200);

setTimeout(function(){
	$(".node").eq(0).animate({
	top :"100px",
	opacity: "1"},
	100);
},450);


}else{

$(".node").eq(2).animate({
	left :"550px",
	opacity: "1"},
	500);
setTimeout(function(){
	$(".node").eq(1).animate({
	left :"325px",
	opacity: "1"},
	300);
},200);

setTimeout(function(){
	$(".node").eq(0).animate({
	left :"100px",
	opacity: "1"},
	100);
},450);
}

PNotify.prototype.options.styling = "fontawesome";
$("#deploy").click(function(){
	$.get("/dashboard/deploy",function(data){
		//console.log(data);
        if(data.status=="fail"){
            new PNotify({title:"Deploy Failed",text: data.error, type: "error"});
        }else{
            new PNotify({title: "Deploy Success", type: "success"});
        }

	});
})