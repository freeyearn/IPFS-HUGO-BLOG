import request from "@/utils/request";
import { objectToFormData } from "@/utils/formDataFormat";

export function useCategoryApi() {
  return {
    /*
      分类列表查询  分页
      参数
       user_id
       Page
       Size
    */
    getCategoryList: (params: any) => {
      return request({
        url: '/category/list/page',
        method: 'get',
        params
      });
    },
    /*
     分类列表查询  ALL
     参数
      user_id
   */
    getCategoryListALL: (params: any) => {
      return request({
        url: '/category/list',
        method: 'get',
        params
      });
    },
    // 分类删除 category_id
    delCategory: (params: any) => {
      return request({
        url: '/category',
        method: 'delete',
        params,
      });
    },
    /*
     分类修改
     参数
      category_id
      category_name
      is_deleted
   */
    updataCategory: (categoryId: any, params: any) => {
      return request({
        url: '/category' + '/' + categoryId,
        method: 'put',
        data: objectToFormData(params),
      });
    },
    /*
     分类添加
     参数
      user_id
      category_name
   */
    addCategory: (params: any) => {
      return request({
        url: '/category',
        method: 'post',
        data: objectToFormData(params),
      });
    },
    /*
      分类查询
      参数
       user_id
       category_id
    */
    getCategoryCount: (params: any) => {
      return request({
        url: '/category',
        method: 'get',
        params
      });
    },

  };
}
