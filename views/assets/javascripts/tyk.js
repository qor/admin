/* On form cancel */
function onCancel() {
  let paths = window.location.pathname.split('/').filter((path) => path);
  let backURL = window.location.origin + '/' + paths[0] + '/' + paths[1];

  /* If it has opened on a new page, needs to go back with page reload otherwise without */
  if ($('.qor-slideout').hasClass('is-slided')) {
    window.history.pushState({}, null, backURL);
  } else {
    window.location.href = backURL;
  }
}

function CodeReferenceOnChange(inputEl, value) {
  if (inputEl.value.match(/\.{1,}/)) return;
  const labelValue =
    inputEl.parentElement.parentElement.parentElement.lastElementChild.lastElementChild.lastElementChild.innerText;
  inputEl.parentElement.parentElement.parentElement.lastElementChild.lastElementChild.lastElementChild.innerText =
    labelValue.replace(/\.blocks\..*?\./g, `.blocks.${value}.`);
}
