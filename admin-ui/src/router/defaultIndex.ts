export default [

  {
    path: "/",
    redirect: "/dashboard",
    component: () => import(/* webpackChunkName: "login" */ '@/layout/index.vue'),
    children: [

      {
        path: '/dashboard',
        name: 'Dashboard',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "dashboard" */ '@/views/dashboard/dashboard.vue'),
        meta: {
          role: 'common',
          title: 'dashboard'
        }
      },
      // articles route
      {
        path: '/add-article',
        name: 'AddArticle',
        component: () => import('@/views/article/add.vue'),
        meta: {
          role: 'common',
          title: 'add article'
        }
      },
      {
        path: '/update-article',
        name: 'UpdateArticle',
        component: () => import('@/views/article/update.vue'),
        meta: {
          role: 'common',
          title: 'update article'
        }
      },
      {
        path: '/articles',
        name: 'Articles',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "articles" */ '@/views/article/index.vue'),
        meta: {
          role: 'common',
          title: 'blog list'
        }
      },
      // category route
      {
        path: '/category',
        name: 'category',
        component: () => import('@/views/category/index.vue'),
        meta: {
          role: 'common',
          title: 'category'
        }
      },
    ],
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/login.vue'),
    meta: {
      role: 'common',
      title: 'login'
    }
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('@/views/login/register.vue'),
    meta: {
      role: 'common',
      title: 'register'
    }
  },
]
