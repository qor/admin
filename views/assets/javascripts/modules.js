import { Collapsible } from './modules/collapsible.js';
import { ProfileMenu } from './modules/profile-menu.js';

let settingsCollapsablePanels = Collapsible();

const profileMenu = ProfileMenu({id: "profile-menu"});
document.getElementById("profile-icon").addEventListener("mouseover", profileMenu.show);
document.getElementById("logout-btn").addEventListener("mouseout", profileMenu.hide);
