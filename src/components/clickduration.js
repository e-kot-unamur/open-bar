export function handleClickDuration(elem) {
  let timer;

  function handleMousedown() {
    timer = Date.now();
  };

  function handleMouseup() {
    if (Date.now() - timer > 400) {
      elem.dispatchEvent(new CustomEvent('longclick'))
    } else {
      // shortclick instead of click => preventing collision
      elem.dispatchEvent(new CustomEvent('shortclick'))
    }
  };

  elem.addEventListener('pointerdown', handleMousedown);
  elem.addEventListener('pointerup', handleMouseup);

  return {
    destroy() {
      elem.removeEventListener('pointerdown', handleMousedown);
      elem.removeEventListener('pointerup', handleMousedown);
    }
  };
}