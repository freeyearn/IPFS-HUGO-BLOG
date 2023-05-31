<template>
  <!-- Navbar -->
  <nav class="navbar navbar-main navbar-expand-lg px-0 mx-4 shadow-none border-radius-xl" id="navbarBlur"
    data-scroll="false">
    <div class="container-fluid py-1 px-3">
      <div class="collapse navbar-collapse mt-sm-0 mt-2 me-md-0 me-sm-4" id="navbar">
        <div class="ms-md-auto pe-md-3 d-flex align-items-center">
          <!-- <div class="input-group">
            <span class="input-group-text text-body">
              <i class="fas fa-search" aria-hidden="true"></i>
            </span>
            <input type="text" class="form-control" placeholder="Type here..." />
          </div> -->
        </div>
        <ul class="navbar-nav justify-content-end">
          <!-- <li style="margin-right: 10px;">
            <el-dropdown :show-timeout="70" :hide-timeout="50" trigger="click" @command="onLanguageChange">
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
          </li>-->
          <li class="nav-item d-flex align-items-center">
            <el-dropdown split-button @command="onHandleCommandClick">
              <span class="el-dropdown-link">
                <i class="fa fa-user me-sm-1" style="color: black"></i>
                <span style="color: black">
                  {{
                    Session.get("userInfo").account
                  }}
                </span>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>{{ $t("user.userInfo") }}</el-dropdown-item>
                  <el-dropdown-item>
                    {{
                      $t("user.userSetting")
                    }}
                  </el-dropdown-item>
                  <el-dropdown-item divided command="/login">
                    {{
                      $t("user.exitLogin")
                    }}
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script lang="ts" setup>
import router from "@/router";
import { Session } from "@/utils/storage";
// eslint-disable-next-line no-unused-vars
import { ElMessage, ElMessageBox } from "element-plus";

import { useI18n } from "vue-i18n";

const { t } = useI18n();

// 用户设置选择
const onHandleCommandClick = (path: string) => {
  if (path === "/login") {
    ElMessageBox({
      closeOnClickModal: false,
      closeOnPressEscape: false,
      title: t("user.exitLogin"),
      message: t("tips.exitLoginContent"),
      showCancelButton: true,
      confirmButtonText: t("Btns.ok"),
      cancelButtonText: t("Btns.cancle"),
      buttonSize: "default",
      beforeClose: (action, instance, done) => {
        if (action === "confirm") {
          instance.confirmButtonLoading = true;
          instance.confirmButtonText = t("tips.exitingLogin");
          setTimeout(() => {
            done();
            setTimeout(() => {
              instance.confirmButtonLoading = false;
            }, 300);
          }, 700);
        } else {
          done();
        }
      },
    })
      .then(async () => {
        // 清除缓存/token等
        Session.clear();
        localStorage.clear();
        // 使用 reload 时，不需要调用 resetRoute() 重置路由
        window.location.reload();
      })
      .catch(() => { });
  } else {
    router.push(path);
  }
};
</script>

<style lang="scss" scoped></style>
