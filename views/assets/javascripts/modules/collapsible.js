export function Collapsible (options) {
  let clickableElement = [...document.getElementsByClassName(options?.clickableSelector || 'collapsible-header')];
  let collapsibleContent = [...document.getElementsByClassName(options?.contentSelector || 'collapsible-content')];
  let ctaWrapper = document.querySelector(options?.textSelector);

  clickableElement.forEach((item,index) => {
    item.addEventListener('click', () => {
      let contentStyle = collapsibleContent[index].style;
      let initialText = ctaWrapper.querySelector(`${options?.textSelector}-text.initial`);
      let secondaryText = ctaWrapper.querySelector(`${options?.textSelector}-text.secondary`);
      let arrowIcon = ctaWrapper.querySelector('.tyk-icon');

      if(contentStyle.display === 'none' || window.getComputedStyle(collapsibleContent[index], null).display === 'none') {
        contentStyle.display = 'block';
        if(initialText) {
          initialText.style.display = 'inline';
          secondaryText.style.display = 'none';
        }
        arrowIcon?.classList.replace('tykon-arrowup', 'tykon-arrowdown')
      } else {
        contentStyle.display = 'none';
        if(initialText) {
          initialText.style.display = 'none';
          secondaryText.style.display = 'inline';
        }
        arrowIcon?.classList.replace('tykon-arrowdown', 'tykon-arrowup')
      }
    });
  });
};

