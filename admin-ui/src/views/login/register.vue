<template>
  <main class="main-content mt-0">
    <div class="page-header align-items-start min-vh-50 pt-5 pb-11 m-3 border-radius-lg" style="
        background-image: url('img/signup-cover.jpg');
        background-position: top;
      ">
      <span class="mask bg-gradient-dark opacity-6"></span>
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-lg-5 text-center mx-auto">
            <h1 class="text-white mb-2 mt-5">{{ $t("register.tip4") }}</h1>
            <p class="text-lead text-white">
              {{ $t("register.tip5") }}
            </p>
          </div>
        </div>
      </div>
    </div>
    <div class="container register-form">
      <div class="row mt-lg-n10 mt-md-n11 mt-n10 justify-content-center">
        <div class="col-xl-4 col-lg-5 col-md-7 mx-auto">
          <div class="card z-index-0">
            <div class="card-header text-center pt-4">
              <span class="font-weight-bolder">{{ $t("register.title") }}</span>
              <el-dropdown :show-timeout="70" :hide-timeout="50" trigger="click" @command="onLanguageChange"
                style="float: right;margin-top: -3px;">
                <el-button v-if="showLang === 'zh-CN'" circle>
                  Zh
                </el-button>
                <el-button v-if="showLang === 'en-US'" circle>
                  En
                </el-button>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="zh">简体中文</el-dropdown-item>
                    <el-dropdown-item command="en">English</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>

            <div class="card-body">
              <form role="form">
                <div class="mb-3">
                  <input v-model="state.account" type="text" class="form-control" placeholder="Name" aria-label="Name"
                    required />
                </div>
                <div class="mb-3">
                  <input v-model="state.password" type="text" class="form-control" placeholder="password"
                    aria-label="Password" required />
                </div>
                <div class="mb-3">
                  <input v-model="state.wallet" type="password" class="form-control" placeholder="wallet"
                    aria-label="Password" />
                </div>
                <div class="form-check form-check-info text-start">
                  <input class="form-check-input" type="checkbox" value="" id="flexCheckDefault" checked />
                  <label class="form-check-label" for="flexCheckDefault">
                    {{ $t("register.tip1") }}
                    <a href="javascript:;" class="text-dark font-weight-bolder">{{ $t("register.tip2") }}</a>
                  </label>
                </div>
                <div class="text-center">
                  <button type="button" class="btn bg-gradient-dark w-100 my-4 mb-2" @click="register()">
                    {{ $t("register.title") }}
                  </button>
                </div>
                <p class="text-sm mt-3 mb-0">
                  {{ $t("register.tip3") }}
                  <a href="javascript:;" class="text-dark font-weight-bolder">
                    <router-link to="/login" >
                      {{ $t("login.title") }}
                    </router-link>
                  </a>
                </p>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script lang="ts" setup>
import { ref } from "vue";

import { useLoginApi } from "@/api/login"
import router from "@/router";
import { ElMessage } from "element-plus";

import { useAppStore } from "@/store";
import { getLocal } from "@/i18n";

const appStore = useAppStore();
const showLang = ref(getLocal())
// 语言切换
const onLanguageChange = (lang: string) => {
  showLang.value = getLocal()
  if (lang === 'zh') {
    showLang.value = 'zh-CN'
    appStore.setLanguage("zh-CN");
  } else if (lang === 'en') {
    showLang.value = 'en-US'
    appStore.setLanguage("en-US");
  }
}
const state = {
  account: '',
  password: '',
  wallet: ''
}

// eslint-disable-next-line no-unused-vars
const RegisterApi = useLoginApi();

const register = () => {
  if (!state.account) {
    ElMessage.error('Name不能为空')
    return
  }
  if (!state.password) {
    ElMessage.error('password不能为空')
    return
  }
  const parms = {
    wallet: state.wallet,
    account: state.account,
    password: state.password
  }
  RegisterApi.register(parms)
    .then((res: any) => {
      if (res.code === '2000') {
        // console.log(res)
        ElMessage({
          type: 'success',
          message: "注册成功!跳转登录",
        });
        router.push('login')
      }
    })
    .catch((err: any) => {
      ElMessage({
        type: 'warning',
        message: err.message,
      });
    })

  // console.log(parms)
}
</script>

<style lang="scss" scoped></style>
