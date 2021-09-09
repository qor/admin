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

<<<<<<< HEAD
function CodeReferenceOnChange(inputId, value, crType) {
  const codeRefContainer = document.getElementById(`${inputId}_container`);
  if (value == '') {
    codeRefContainer.style.display = 'none';
    return;
  }
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
  codeRefEl.innerText = crValue;
}

function copyToClipboard(elemId) {
  const codeRefEl = document.getElementById(elemId);
  navigator.clipboard.writeText(codeRefEl.innerText);
}
=======
document.getElementById("profile-icon").addEventListener("mouseover", showProfileMenu);
document.getElementById("logout-btn").addEventListener("mouseout", hideProfilemenu);
document.getElementById("logout-btn").addEventListener("mouseleave", hideProfilemenu);

function showProfileMenu() {
  document.getElementById("profile-menu").style.visibility = "visible";
}

function hideProfilemenu() {
  document.getElementById("profile-menu").style.visibility = "hidden";
}
>>>>>>> fbabca1... rebasing
