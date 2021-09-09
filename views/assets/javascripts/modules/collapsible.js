export function Collapsible (options) {
	let clickableElement = [...document.getElementsByClassName(options?.clickableSelector || 'collapsible-header')];
	let collapsibleContent = [...document.getElementsByClassName(options?.contentSelector || 'collapsible-content')]

	clickableElement.forEach((item,index) => {
  	item.addEventListener('click', () => {
  		let arrowIcon = clickableElement[index].children[0].classList[0] === 'tyk-icon' ? clickableElement[index].children[0] : null;
			if(
					collapsibleContent[index].style.display.length === 0 ||
					collapsibleContent[index].style.display === 'none'
				) {
				collapsibleContent[index].style.display = 'block'
				arrowIcon?.classList.remove('tykon-arrowup');
				arrowIcon?.classList.add('tykon-arrowdown');
			} else {
				collapsibleContent[index].style.display = 'none'
				arrowIcon?.classList.remove('tykon-arrowdown');
				arrowIcon?.classList.add('tykon-arrowup');
			}
  	});
	});
};
