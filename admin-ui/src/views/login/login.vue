<template>
  <div>
    <main class="main-content mt-0 box-size">
      <section>
        <div class="page-header min-vh-80">
          <div class="container login-form">
            <div class="row">
              <div class="col-xl-4 col-lg-5 col-md-7 d-flex flex-column mx-lg-0 mx-auto">
                <div class="card card-plain">
                  <div class="card-header pb-0 text-start">
                    <div class="mb-3">
                      <span class="font-weight-bolder">
                        {{
                          $t("login.title")
                        }}
                      </span>
                      <el-dropdown :show-timeout="70" :hide-timeout="50" trigger="click" @command="onLanguageChange"
                        style="float: right; margin-top: -3px">
                        <el-button v-if="showLang === 'zh-CN'" circle>Zh</el-button>
                        <el-button v-if="showLang === 'en-US'" circle>En</el-button>
                        <template #dropdown>
                          <el-dropdown-menu>
                            <el-dropdown-item command="zh">简体中文</el-dropdown-item>
                            <el-dropdown-item command="en">English</el-dropdown-item>
                          </el-dropdown-menu>
                        </template>
                      </el-dropdown>
                    </div>
                    <p class="mb-1">{{ $t("login.tip1") }}</p>
                  </div>
                  <div class="card-body">
                    <el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="0px"
                      class="demo-ruleForm">
                      <el-form-item prop="pass">
                        <el-label>{{ $t("login.account") }}</el-label>
                        <el-input v-model="ruleForm.account" type="text" autocomplete="off" />
                      </el-form-item>
                      <el-form-item prop="checkPass">
                        <el-label>{{ $t("login.password") }}</el-label>
                        <el-input v-model="ruleForm.password" type="password" autocomplete="off" show-password />
                      </el-form-item>
                      <el-form-item>
                        <el-button type="primary" @click="login(ruleFormRef)" class="mx-auto d-block"
                          :loading="loginLoading">{{ $t("login.title")
                          }}</el-button>
                      </el-form-item>
                    </el-form>
                  </div>
                  <div class="card-footer text-center pt-0 px-lg-2 px-1">
                    <p class="mb-4 text-sm mx-auto">
                      {{ $t("login.tip2") }}
                      <a href="javascript:;" class="text-primary text-gradient font-weight-bold">
                        <router-link to="/register">{{ $t("register.title") }}</router-link>
                      </a>
                    </p>
                  </div>
                </div>
              </div>
              <div
                class="col-6 d-lg-flex d-none h-100 my-auto pe-0 position-absolute top-0 end-0 text-center justify-content-center flex-column">
                <div
                  class="position-relative bg-gradient-primary h-100 m-3 px-7 border-radius-lg d-flex flex-column justify-content-center overflow-hidden"
                  style="
                    background-image: url('img/signin-ill.jpg');
                    background-size: cover;
                  ">
                  <span class="mask bg-gradient-primary opacity-6"></span>
                  <h4 class="mt-5 text-white font-weight-bolder position-relative">"Attention is the new currency"</h4>
                  <p class="text-white position-relative">
                    The more effortless the writing looks, the more effort the
                    writer actually put into the process.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import { useLoginApi } from "@/api/login";
import { UserInfo } from "@/api/userInfo";
import { ElMessage } from "element-plus";
import { Session } from "@/utils/storage";
import { rsaEncrypt } from "@/utils/rsaEncrypt";
import router from "@/router";

import type { FormInstance, FormRules } from "element-plus";
import systemConfig from "@/config";

import { useAppStore } from "@/store";
import { getLocal } from "@/i18n";

const loginLoading = ref(false)
const appStore = useAppStore();
const showLang = ref(getLocal());
// 语言切换
const onLanguageChange = (lang: string) => {
  showLang.value = getLocal();
  if (lang === "zh") {
    showLang.value = "zh-CN";
    appStore.setLanguage("zh-CN");
  } else if (lang === "en") {
    showLang.value = "en-US";
    appStore.setLanguage("en-US");
  }
};

const ruleFormRef = ref<FormInstance>();

const validateAccount = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("Please input the password"));
  }
};
const validatePass = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error("Please input the password again"));
  }
};

const ruleForm = reactive({
  account: "",
  password: "",
});

const rules = reactive<FormRules>({
  account: [{ validator: validateAccount, trigger: "blur" }],
  password: [{ validator: validatePass, trigger: "blur" }],
});

const LoginApi = useLoginApi();
const getUserInfoAPI = UserInfo();

const login = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async (valid) => {
    if (valid) {
      const parms = {
        account: ruleForm.account,
        password: rsaEncrypt(ruleForm.password),
      };
      loginLoading.value = true
      await LoginApi.signIn(parms)
        .then(async (res: any) => {
          if (res.code === "2000") {
            // console.log(res);
            localStorage.setItem("account", res.data.Name);
            localStorage.setItem(systemConfig.TOKEN_NAME, res.data.Token);

            Session.set("userName", res.data.Name);
            Session.set("userID", res.data.UserId);
            Session.set(systemConfig.TOKEN_NAME, res.data.Token);
          }
          loginLoading.value = false
        })
        .catch((err: any) => {
          loginLoading.value = false

          ElMessage({
            type: "warning",
            message: err.message,
          });
        });
      console.log(Session.get(systemConfig.TOKEN_NAME))
      if (Session.get(systemConfig.TOKEN_NAME)) {
        const parms = {
          token: Session.get(systemConfig.TOKEN_NAME),
        };

        await getUserInfoAPI
          .getUserInfo(parms)
          .then((res: any) => {
            if (res.code === "2000") {
              Session.set("userInfo", res.data);
            }
          })
          .catch((err: any) => {
            ElMessage({
              type: "warning",
              message: err.message,
            });
          });
      }

      router.push("/dashboard");
    }
  });
};
</script>

<style lang="scss" scoped>
.login-form {
  margin-top: 100px;
}
</style>
