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

function CodeReferenceOnChange(inputId, value, crType) {
  const codeRefContainer = document.getElementById(`${inputId}_container`);
  codeRefContainer.style.display = 'flex';
  const codeRefEl = document.getElementById(inputId);
  let crValue;
  switch (crType) {
    case 'content':
      crValue = `safe .blocks.${value}.Content`;
      break;
    case 'image':
      crValue = `.blocks.${value}.Image.URL`;
      break;
    default:
      crValue = `.blocks.${value}.Content`;
      break;
  }
  codeRefEl.value = value == '' ? '' : `{{ ${crValue} }}`;
}

function copyToClipboard(elemId) {
  const codeRefEl = document.getElementById(elemId);
  navigator.clipboard.writeText(codeRefEl.value);
}

function hideHelpSection() {
  let overviewToggleButton = document.querySelector('.toggle-help-section');
  let overviewToggleSection = document.querySelector('.help-section');

  overviewToggleButton[0].style.display = 'none';
  overviewToggleSection[0].style.display = 'none';
  localStorage.setItem('hideHelpSection', 'hidden');
}

document.addEventListener('DOMContentLoaded', function() {
  let overviewToggleButton = document.querySelector('.toggle-help-section');
  let overviewToggleSection = document.querySelector('.help-section');

  if(localStorage.getItem('hideHelpSection') === 'hidden') {
    overviewToggleButton[0].style.display = 'none';
    overviewToggleSection[0].style.display = 'none';
  }
}, false);
