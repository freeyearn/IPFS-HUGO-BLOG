import request from "@/utils/request";
// import { objectToFormData } from "@/utils/formDataFormat";

export function useNewsApi() {
  return {
    getNewsList: (params:any) => {
      return request({
        url: '/news-list',
        method: 'get',
        params,
      });
    }
  };
}
