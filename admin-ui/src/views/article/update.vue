<template>
  <el-card>
    <div class="container">
      <div class="left">
        <el-page-header @back="$router.push('/articles')">
          <template #content>
            <el-text class="mx-1" size="large">{{ $t("article.addArticleContent") }}</el-text>
          </template>
        </el-page-header>
        <div class="edit-size" style="border: 1px solid #ccc;margin-top: 15px;">
          <div id="vditor" v-loading="loading.vditorLoading"></div>
        </div>
      </div>

      <div class="right">
        <el-text class="mx-1" size="large">{{ $t("article.addArticleParameters") }}</el-text>
        <hr />
        <div class="form">
          <el-form ref="ruleFormRef" :rules="rules" :label-position="labelPosition" label-width="100px"
            :model="formLabelAlign" style="max-width: 460px">
            <el-form-item :label="addArticleParameters.author_name">
              <el-input v-model="formLabelAlign.author_name" disabled />
            </el-form-item>

            <el-form-item :label="addArticleParameters.title" prop="title">
              <el-input v-model="formLabelAlign.title" placeholder="Please input 标题" clearable />
            </el-form-item>

            <el-form-item :label="addArticleParameters.description" prop="description">
              <el-input v-model="formLabelAlign.description" placeholder="Please input 摘要" clearable />
            </el-form-item>

            <!-- <el-form-item label="发布于">
              <div class="demo-date-picker">
                <div class="block">
                  <el-date-picker v-model="input" type="date" placeholder="Pick a date"
                    :default-value="new Date(2010, 9, 1)" />
                </div>
              </div>
            </el-form-item>

            <el-form-item label="更新于">
              <div class="demo-date-picker">
                <div class="block">
                  <el-date-picker v-model="input" type="date" placeholder="Pick a date"
                    :default-value="new Date(2010, 9, 1)" />
                </div>
              </div>
            </el-form-item> -->

            <el-form-item :label="addArticleParameters.tags">
              <el-tag v-for="tag in dynamicTags" :key="tag" class="mx-1" closable :disable-transitions="false"
                @close="handleClose(tag)" style="margin-top: 5px;">
                <span v-if="tag === ''">
                  null
                </span>
                <span v-else>
                  {{ tag }}
                </span>
              </el-tag>
              <el-input v-if="inputVisible" ref="InputRef" v-model="inputValue" class="ml-1 w-20" size="small"
                @keyup.enter="handleInputConfirm" @blur="handleInputConfirm" />
              <el-button v-else class="button-new-tag ml-1" size="small" @click="showInput">
                + New Tag
              </el-button>
            </el-form-item>

            <el-form-item :label="addArticleParameters.category" prop="category">
              <el-select v-model="formLabelAlign.category" class="m-2" placeholder="Select">
                <el-option v-for="item in categoryList" :key="item.category_id" :label="item.category_name"
                  :value="item.category_id" />
              </el-select>
            </el-form-item>

          </el-form>
        </div>
        <div class="btns">
          <el-button type="primary" style="float: right" @click="publicFun(ruleFormRef)" :disabled="loading.isDisabled"
            :loading="loading.saveLoading">{{
              $t("Btns.save")
            }}</el-button>
        </div>
      </div>
    </div>
  </el-card>
</template>
<script lang="ts" setup>
import "@wangeditor/editor/dist/css/style.css"; // 引入 css

// eslint-disable-next-line no-unused-vars
import { onBeforeUnmount, reactive, ref, shallowRef, nextTick, onMounted } from "vue";
// eslint-disable-next-line no-unused-vars
import { ElInput, ElMessage, FormInstance, FormRules } from 'element-plus'
import { Session } from "@/utils/storage"
import { useArticleApi } from "@/api/articles"
import { useCategoryApi } from "@/api/category";
import { useI18n } from "vue-i18n";
// eslint-disable-next-line no-unused-vars
import router from "@/router";
import { useRoute } from "vue-router";
import Vditor from "vditor";
import "vditor/dist/index.css";

const loading = reactive({
  vditorLoading: false,
  isDisabled: false,
  saveLoading: false
})

// 表单校验
const ruleFormRef = ref<FormInstance>()
const checkTitle = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error(t("tips.rules.checkTitle")))
  } else {
    if (formLabelAlign.title !== '') {
      if (!ruleFormRef.value) return
      ruleFormRef.value.validateField('checkPass', () => null)
    }
    callback()
  }
}
const checkDescription = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error(t("tips.rules.checkDescription")))
  } else {
    if (formLabelAlign.description !== '') {
      if (!ruleFormRef.value) return
      ruleFormRef.value.validateField('checkPass', () => null)
    }
    callback()
  }
}
const checkCategory = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error(t("tips.rules.checkCategory")))
  } else {
    if (formLabelAlign.category !== '') {
      if (!ruleFormRef.value) return
      ruleFormRef.value.validateField('checkPass', () => null)
    }
    callback()
  }
}
const rules = reactive<FormRules>({
  title: [{ required: true, validator: checkTitle, trigger: 'blur' }],
  description: [{ required: true, validator: checkDescription, trigger: 'blur' }],
  category: [{ required: true, validator: checkCategory, trigger: 'change' }],
})

// 双语引入
const { t } = useI18n();
const addArticleParameters = {
  author_name: t("article.Parameters.author_name"),
  title: t("article.Parameters.title"),
  description: t("article.Parameters.description"),
  tags: t("article.Parameters.tags"),
  category: t("article.Parameters.category"),
}

const labelPosition = ref("top");

// 用户信息
const userInfo = Session.get("userInfo")
// console.log(userInfo)

// Vditor markdown编辑器
const contentEditor = ref();
const setVditor = () => {
  contentEditor.value = new Vditor("vditor", {
    height: 650,
    mode: 'sv', // 显示预览内容
    toolbarConfig: {
      pin: true,
    },
    cache: {
      enable: false,
    },
    after: () => {
      contentEditor.value.setValue("······内容正在加载中······");
    },
  });
}

// 表单信息
const formLabelAlign = reactive({
  author_name: userInfo.name,
  title: '',
  description: '',
  category: '',
});

// 接口请求参数
const saveInfo = (status: number) => {
  const category = []
  category[0] = String(formLabelAlign.category)
  const parms = {
    article_id: route.query.article_id,
    // author_id: userInfo.user_id,
    author_name: formLabelAlign.author_name,
    title: formLabelAlign.title,
    description: formLabelAlign.description,
    content: contentEditor.value.getValue(),
    tags: JSON.stringify(dynamicTags.value),
    category: JSON.stringify(category),
    keyword: '',
    next: '',
    prev: '',
    status: status,
  }
  // console.log(formLabelAlign.category, parms.category)

  return parms
}
// eslint-disable-next-line no-unused-vars
const articlesApi = useArticleApi()

// 发布功能
const publicFun = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid, fields) => {
    if (valid) {
      const parms = saveInfo(0)
      loading.saveLoading = true
      articlesApi.updataArticle(parms)
        .then((res: any) => {
          if (res.code === '2000') {
            // console.log(res)
            ElMessage({
              type: 'success',
              message: '更新成功',
            });
            loading.saveLoading = false
            router.push('/articles')
          }
        })
        .catch((err: any) => {
          loading.saveLoading = false

          ElMessage({
            type: 'warning',
            message: err.message,
          });
        })
      // console.log(parms)
    } else {
      // console.log('error submit!', fields)
    }
  })
}

// 添加 tags
const inputValue = ref('')
const dynamicTags = ref([''])
const inputVisible = ref(false)
const InputRef = ref<InstanceType<typeof ElInput>>()

const showInput = () => {
  inputVisible.value = true
  nextTick(() => {
    InputRef.value!.input!.focus()
  })
}
// 删除标签
const handleClose = (tag: string) => {
  dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1)
}
const handleInputConfirm = () => {
  if (inputValue.value) {
    if (dynamicTags.value[0] === '') {
      dynamicTags.value[0] = inputValue.value
    } else {
      dynamicTags.value.push(inputValue.value)
    }
  }
  inputVisible.value = false
  inputValue.value = ''
}

// 分类 变量定义
const categoryApi = useCategoryApi()
const categoryList = ref()
// 分类列表 获取
const getcategoryList = () => {
  const parms = {
    user_id: userInfo.user_id,
  }
  categoryApi.getCategoryListALL(parms)
    .then((res: any) => {
      if (res.code === '2000') {
        categoryList.value = res.data
        // console.log(categoryList.value)
      }
    })
    .catch((err: any) => {
      ElMessage({
        type: 'warning',
        message: err.message,
      });
    })
}
const route = useRoute()
// 获取文章内容
const getArticleContent = () => {
  const parms = {
    article_id: route.query.article_id
  }
  // console.log(parms.article_id)
  loading.vditorLoading = true
  loading.isDisabled = true
  articlesApi.getArticle(parms)
    .then((res: any) => {
      if (res.code === '2000') {
        // console.log(res)
        contentEditor.value.setValue(res.data.Content)
        formLabelAlign.title = res.data.title
        formLabelAlign.description = res.data.description
        if (res.data.category) {
          formLabelAlign.category = JSON.parse(res.data.category)
        }
        dynamicTags.value = JSON.parse(res.data.tags)
      }
      loading.vditorLoading = false
      loading.isDisabled = false
    })
    .catch((err: any) => {
      ElMessage({
        type: 'warning',
        message: err.message,
      });
    })
}

// 预加载
onMounted(() => {
  getArticleContent()
  getcategoryList()
  setVditor()
})
</script>

<style lang="scss">
// .edit-size {
//   width: 60%;
// }
.el-card {
  .container {
    max-width: 1500px;
    margin: 0px;
    display: flex;
  }

  .left {
    flex: 8;
  }

  .right {
    margin-left: 30px;
    flex: 2;

    .btns {
      margin-top: 20px;

      .el-button {
        margin: 10px;
      }
    }
  }
}
</style>
