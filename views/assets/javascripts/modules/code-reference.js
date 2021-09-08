export function CodeReferenceOnChange(inputEl, value) {
  if (inputEl.value.match(/\.{1,}/)) return;
  const labelValue =
    inputEl.parentElement.parentElement.parentElement.lastElementChild.lastElementChild.lastElementChild.innerText;
  inputEl.parentElement.parentElement.parentElement.lastElementChild.lastElementChild.lastElementChild.innerText =
    labelValue.replace(/\.blocks\..*?\./g, `.blocks.${value}.`);
}
