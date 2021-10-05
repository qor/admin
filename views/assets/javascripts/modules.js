import { Collapsible } from './modules/collapsible.js';

let settingsCollapsablePanels = Collapsible();

let overviewSelectors = {
  clickableSelector: 'toggle-help-section__cta',
  contentSelector: 'help-section'
}

let overviewSection = Collapsible(overviewSelectors)