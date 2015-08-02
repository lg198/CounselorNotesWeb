Backend = {};

Backend.searchStudents = function(query, callback) {
    var data ={"Function":"searchStudents", "query":query};
    var stringy = JSON.stringify(data);
    console.log(stringy);
    $.ajax({
    type: "POST",
    url: "/handle",
    processData: false,
    contentType: 'application/json',
    data: stringy,
    success:callback});
}


//////////////////////////////////////////////////////////////////////////////////////////////////////////

$(function() {
    $(".list").on("click", "li", function(evt) {
        var selected = $(this);
        if (!selected.hasClass("active")) {
            $(".list li.active").each(function() {
               $(this).toggleClass("active");
            });
            selected.toggleClass("active");
        }
        /*$(".list > li").each(function() {
            if (!$(this).is(selected) && $(this).hasClass("active")  ) {
                $(this).removeClass("active");
            }
        }); */
    });
    
    $("#student-search").on("keyup change", function() {
        if ($(this).val().length <= 0) {
            return;
        }
        Backend.searchStudents($(this).val(), function(ret) {
            alert(ret);
        });
    });
});

