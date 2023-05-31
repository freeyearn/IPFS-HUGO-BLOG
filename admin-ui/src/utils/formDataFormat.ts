/**
 * Created by dxz-dev on 2020/4/29.
 * explain:
 * */
export function objectToFormData(obj: any, form?: any, namespace?: any) {
  const fd = form || new FormData();
  let formKey;

  for (const property in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, property)) {
      const key = Array.isArray(obj) ? "[]" : `[${property}]`;
      if (namespace) {
        formKey = namespace + key;
      } else {
        formKey = property;
      }

      // if the property is an object, but not a File, use recursivity.
      if (
        typeof obj[property] === "object" &&
        !(obj[property] instanceof File)
      ) {
        objectToFormData(obj[property], fd, formKey);
      } else {
        // if it's a string or a File object
        fd.append(formKey, obj[property]);
      }
    }
  }
  return fd;
}
