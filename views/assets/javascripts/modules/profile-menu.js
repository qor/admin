export function ProfileMenu ({id}) {
  const $element = document.getElementById(id);

  const toggleMenu = (action) => {
    $element.style.visibility = action;
  }

  const show = () => {
    toggleMenu("visible");
  }

  const hide = () => {
    toggleMenu("hidden");
  }

  return {
    show,
    hide
  }
};
