<template>
  <aside class="sidenav bg-white navbar navbar-vertical navbar-expand-xs border-0 border-radius-xl my-3 fixed-start ms-4"
    id="sidenav-main">
    <div class="sidenav-header">
      <i class="fas fa-times p-3 cursor-pointer text-secondary opacity-5 position-absolute end-0 top-0 d-none d-xl-none"
        aria-hidden="true" id="iconSidenav"></i>
      <a class="navbar-brand m-0" href=" https://demos.creative-tim.com/argon-dashboard/pages/dashboard.html "
        target="_blank">
        <img src="@/assets/img/logo-ct-dark.png" class="navbar-brand-img h-100" alt="main_logo" />
        <span class="ms-1 font-weight-bold">{{ $t("menus.model")}}</span>
      </a>
    </div>
    <hr class="horizontal dark mt-0" />
    <div class="collapse navbar-collapse w-auto" id="sidenav-collapse-main">
      <ul class="navbar-nav">
        <li class="nav-item " :class="{ active: $route.path === '/' }">
          <router-link to="/" class="nav-link">
            <div
              class="icon icon-shape icon-sm border-radius-md text-center me-2 d-flex align-items-center justify-content-center">
              <i class="ni ni-tv-2 text-primary text-sm opacity-10"></i>
            </div>
            <span class="nav-link-text ms-1">{{ $t("menus.dashboard")}}</span>
          </router-link>
        </li>
        <li class="nav-item" :class="{ active: $route.path === '/articles' }">
          <router-link to="/articles" class="nav-link">
            <div
              class="icon icon-shape icon-sm border-radius-md text-center me-2 d-flex align-items-center justify-content-center">
              <i class="ni ni-calendar-grid-58 text-warning text-sm opacity-10"></i>
            </div>
            <span class="nav-link-text ms-1">{{ $t("menus.blogs")}}</span>
          </router-link>
        </li>
        <li class="nav-item" :class="{ active: $route.path === '/category' }">
          <router-link to="/category" class="nav-link">
            <div
              class="icon icon-shape icon-sm border-radius-md text-center me-2 d-flex align-items-center justify-content-center">
              <i class="ni ni-calendar-grid-58 text-warning text-sm opacity-10"></i>
            </div>
            <span class="nav-link-text ms-1">{{ $t("menus.category")}}</span>
          </router-link>
        </li>
      </ul>
    </div>
  </aside>
</template>

<script lang="ts" setup>

</script>

<style lang="scss" scoped>
.nav-item.active {
  background-color: #e1e6ff;
  border-radius: 10px;
}
</style>
