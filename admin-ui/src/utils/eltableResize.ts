// 解决eltable无法自适应宽度问题
export function onHandleElTableResize() {
  const debounce = (fn:any, delay:any) => {
    let timer:any = null;
    return function(this: any, ...args:any[]) {
      const context = this;
      clearTimeout(timer);
      timer = setTimeout(function () {
        fn.apply(context, args);
      }, delay);
    };
  };

  const _ResizeObserver = window.ResizeObserver;
  window.ResizeObserver = class ResizeObserver extends _ResizeObserver {
    constructor(callback:any) {
      callback = debounce(callback, 16);
      super(callback);
    }
  };
}
