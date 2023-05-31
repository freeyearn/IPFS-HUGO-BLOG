import { createRouter, createWebHashHistory } from 'vue-router';
import defaultIndex from '@/router/defaultIndex';

import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import { Session } from '@/utils/storage';
import systemConfig from "@/config"

export const router = createRouter({
  history: createWebHashHistory(),
  routes: defaultIndex,
});

// 路由加载前
router.beforeEach(async (to, from, next) => {
  NProgress.configure({ showSpinner: false });
  if (to.meta.title) NProgress.start();
  const token = Session.get(systemConfig.TOKEN_NAME);
  const role = Session.get('role');
  if (to.path === '/login' && !token) {
    next();
    NProgress.done();
  } else if (to.path === '/register' && !token) {
    next();
    NProgress.done();
  } else {
    if (!token) {
      next(`/login?redirect=${to.path}&params=${JSON.stringify(to.query ? to.query : to.params)}`);
      Session.clear();
      NProgress.done();
    } else if (token && to.path === '/login') {
      if (role === 'admin' || role === 'common') {
        next('/');
        NProgress.done();
      } else {
        next(`/login?redirect=${to.path}&params=${JSON.stringify(to.query ? to.query : to.params)}`);
        Session.clear();
        NProgress.done();
      }
    } else {
      next();
    }
  }
});

// 路由加载后
router.afterEach(() => {
  NProgress.done();
});

// 导出路由
export default router;
