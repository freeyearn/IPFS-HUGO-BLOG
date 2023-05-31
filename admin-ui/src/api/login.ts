import request from "@/utils/request";
import { objectToFormData } from "@/utils/formDataFormat";

export function useLoginApi() {
  return {
    signIn: (params: any) => {
      return request({
        url: '/user/login',
        method: 'post',
        data: objectToFormData(params),
      });
    },
    register: (params: any) => {
      return request({
        url: '/user/register',
        method: 'post',
        data: objectToFormData(params),
      });
    },
  };
}
