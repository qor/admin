$(document).ready(function () {
	let rotated = false;

	$('.collapsible-header').click(function () {
	  const $header = $(this);
	  const $content = $header.next();
	  const $arrowIcon = $header.find('.tyk-icon.tykon');
	  const contentIsVisible = $content.css('display') === 'block' ? true : false;

  	$arrowIcon.removeClass(contentIsVisible ? 'tykon-arrowdown' : 'tykon-arrowup').addClass(contentIsVisible ? 'tykon-arrowup' : 'tykon-arrowdown');
	  $content.slideToggle(400, function () {});
	  rotated = !rotated;
	});
});

/* On form cancel */
function onCancel() {
	let paths = window.location.pathname.split('/').filter(path => path);
	let backURL = window.location.origin + '/' + paths[0] + '/' + paths[1];

	/* If it has opened on a new page, needs to go back with page reload otherwise without */
	if($('.qor-slideout').hasClass('is-slided')) {
		window.history.pushState({}, null, backURL);
	} else {
		window.location.href = backURL;
	}
}