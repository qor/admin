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

/* should not hide this section with such a method but rather find the issue in the Go template
$(document).ready(function () {
	if($('.qor-page__body').siblings('.qor-table-container')) {
		$('.qor-page__header').remove();
	}
})
*/
