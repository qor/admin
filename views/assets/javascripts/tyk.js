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
    console.log(codeRefContainer)
    if (value == '') {
        codeRefContainer.style.display = 'none';
        return;
    }
    codeRefContainer.style.display = 'flex';
    const codeRefEl = document.getElementById(inputId);
    let crValue;
    console.log(codeRefEl)
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
    codeRefEl.value = `{{ ${crValue} }}`;
}

function copyToClipboard(elemId) {
    const codeRefEl = document.getElementById(elemId);
    navigator.clipboard.writeText(codeRefEl.value);
}