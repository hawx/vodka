$(document).ready(function(){

    $('.summary ul').addClass('closed');

    $('.summary > li > a').click(function() {
        $('.summary ul').addClass('closed');
        var ul = $(this).next();
        ul.toggleClass("closed");
    });

    function openLeaf(name) {
        $('.summary ul').addClass('closed');
        $('.summary a[href="#x'+name+'"]').next().removeClass('closed');
    }

    $('.docs > h1').appear(function() {
        openLeaf($(this).attr('id'));
    }, {one: false});


    var toScroll = $('ul').position().top;
    $(window).scroll(function() {
        if (toScroll >= $(window).scrollTop()) {
            if ($('.summary').hasClass('fixed')) {
                $('.summary').removeClass('fixed');
            }
        } else {
            if (!$('.summary').hasClass('fixed')) {
                $('.summary').addClass('fixed');
            }
        }
    });
});
