import request from "@/utils/request";
import { objectToFormData } from "@/utils/formDataFormat";

export function useArticleApi() {
  return {
    // 添加文章
    addArticle: (params: any) => {
      return request({
        url: '/article',
        method: 'post',
        data: objectToFormData(params),
      });
    },
    // 文章列表查询
    getArticleList: (params: any) => {
      return request({
        url: '/article/list/page',
        method: 'get',
        params
      });
    },
    // 文章发表数量
    getArticleCount: (params: any) => {
      return request({
        url: '/article/count',
        method: 'get',
        params
      });
    },
    // 文章删除
    delArticle: (params: any) => {
      return request({
        url: '/article/' + params.article_id,
        method: 'delete',
      });
    },
    // 文章修改
    updataArticle: (params: any) => {
      return request({
        url: '/article' + '/' + params.article_id,
        method: 'put',
        data: objectToFormData(params),
      });
    },
    // 文章查询
    // getArticle: (params: any) => {
    //   return request({
    //     url: '/article' + '/' + params.article_id,
    //     method: 'get',
    //   });
    // },
    // 文章内容获取
    getArticle: (params: any) => {
      return request({
        url: '/article/fetch',
        method: 'get',
        params
      });
    },
    // 文章状态修改
    updateArticleStatus: (params: any) => {
      return request({
        url: '/article/status',
        method: 'put',
        data: objectToFormData(params),
      });
    },
  };
}
