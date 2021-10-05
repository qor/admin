export function Collapsible (options) {
	let clickableElement = [...document.getElementsByClassName(options?.clickableSelector || 'collapsible-header')];
	let collapsibleContent = [...document.getElementsByClassName(options?.contentSelector || 'collapsible-content')]

	clickableElement.forEach((item,index) => {
  	item.addEventListener('click', () => {
      let contentStyle = collapsibleContent[index].style;
  		let arrowIcon = clickableElement[index].children[0]?.classList[0] === 'tyk-icon' ? clickableElement[index].children[0] : null;
			if(
					contentStyle.display.length === 0 ||
					contentStyle.display === 'none'
				) {
				contentStyle.display = 'block';
				arrowIcon?.classList.remove('tykon-arrowup');
				arrowIcon?.classList.add('tykon-arrowdown');
			} else {
				contentStyle.display = 'none';
				arrowIcon?.classList.remove('tykon-arrowdown');
				arrowIcon?.classList.add('tykon-arrowup');
			}
  	});
	});
};

