import request from "@/utils/request";
// import { objectToFormData } from "@/utils/formDataFormat";

export function UserInfo() {
  return {
    getUserInfo: (params: any) => {
      return request({
        url: '/user/user_info',
        method: 'get',
        headers: params
      });
    }
  };
}
