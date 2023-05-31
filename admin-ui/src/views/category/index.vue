<template>
  <el-card>
    <div>
      <el-text class="mx-1" size="default">{{
        $t("category.categoryList")
      }}</el-text>

      <el-card style="margin-top: 10px;">
        <el-form label-width="90px" size="default">
          <el-row :gutter="10" class="negtive-marginBottom">
            <el-col :lg="6" :md="8" :sm="12">
              <el-form-item label="文章分类">
                <div style="margin-top: -6px;">
                  <el-select v-model="searchData.categpry" class="m-2" placeholder="Select" clearable>
                    <el-option v-for="item in categoryList" :key="item.category_id" :label="item.category_name"
                      :value="item.category_id" />
                  </el-select>
                </div>
              </el-form-item>
            </el-col>
            <el-col :lg="6" :md="8" :sm="12">
              <el-button type="primary" :icon="Search" @click="searchFun()" title="搜索"></el-button>

            </el-col>
            <el-col :lg="6" :md="8" :sm="12">

            </el-col>
            <el-col :lg="6" :md="8" :sm="12">
              <el-button type="primary" style="float: right" :icon="Plus" @click="dialogFormVisible = true" title="新增">
              </el-button>
            </el-col>
          </el-row>
        </el-form>
      </el-card>
    </div>

    <div style="margin-top: 20px;">
      <el-card>
        <el-table :data="tableData" v-loading="loading.tableLoading" style="width: 100%">
          <el-table-column prop="category_id" :label="CategoryListParameters.category_id" width="300" />
          <el-table-column prop="category_name" :label="CategoryListParameters.category_name" width="300" />
          <el-table-column prop="create_time" :label="CategoryListParameters.create_time" min-width="200" />

          <el-tableColumn :label="CategoryListParameters.action" width="200">
            <!-- eslint-disable-next-line vue/no-unused-vars -->
            <template #default="scoped">
              <!-------------------------------- 修改按钮 ---------------------------------->
              <el-button :icon="Edit" type="warning" size="default" title="修改" @click="
                showUpdateDialog(
                  scoped.row.user_id,
                  scoped.row.category_id,
                  scoped.row.category_name
                )
                "></el-button>

              <!------------------------------- 删除按钮 ------------------------------>
              <el-button :icon="Delete" type="danger" size="default" title="删除"
                @click="delCategory(scoped.row.category_id)"></el-button>
            </template>
          </el-tableColumn>
        </el-table>
        <div>
          <el-pagination background class="pagination" layout="prev, pager, next" :total="state.total"
            :current-page="state.currentPage" :page-size="state.pageSize" @current-change="currentChange"
            @size-change="handleSizeChange" style="margin-top: 10px" />
        </div>
      </el-card>
    </div>

  </el-card>

  <!-- 新增 分类 -->
  <el-dialog v-model="dialogFormVisible" :title="addCategoryParameters.addCategoryTitle" style="width: 450px">
    <el-form ref="ruleFormRef" :model="addCategoryParms" :rules="rules">
      <el-form-item :label="addCategoryParameters.addCategoryUser" label-width="100" prop="user_id">
        <el-input v-model="addCategoryParms.user_id" autocomplete="off" style="width: 300px" disabled />
      </el-form-item>
      <el-form-item :label="addCategoryParameters.addCategoryName" label-width="100" prop="category_name">
        <el-input v-model="addCategoryParms.category_name" autocomplete="off" style="width: 300px" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogFormVisible = false">{{
          $t("Btns.cancle")
        }}</el-button>
        <el-button type="primary" @click="addCategory(ruleFormRef)" :loading="loading.addLoading">
          {{ $t("Btns.ok") }}
        </el-button>
      </span>
    </template>
  </el-dialog>
  <!-- 修改分类 -->
  <el-dialog v-model="updateDialogFormVisible" :title="updateCategoryParameters.updateCategoryTitle" style="width: 450px">
    <el-form ref="ruleFormRef" :rules="rules" :model="updateCategoryParms">
      <el-form-item :label="updateCategoryParameters.updateCategoryUser" label-width="100px">
        <el-input v-model="updateCategoryParms.user_id" autocomplete="off" style="width: 300px" disabled />
      </el-form-item>
      <el-form-item :label="updateCategoryParameters.updateCategoryOldName" label-width="100px">
        <el-input v-model="updateCategoryParms.old_category_name" autocomplete="off" style="width: 300px" disabled />
      </el-form-item>
      <el-form-item :label="updateCategoryParameters.updateCategoryNewName" label-width="100px" prop="new_category_name">
        <el-input v-model="updateCategoryParms.new_category_name" autocomplete="off" style="width: 300px" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="updateDialogFormVisible = false">{{
          $t("Btns.cancle")
        }}</el-button>
        <el-button type="primary" @click="updateCategory(ruleFormRef)" :loading="loading.updateLoading">
          {{ $t("Btns.ok") }}
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {
  ComponentInternalInstance,
  getCurrentInstance,
  onBeforeMount,
  reactive,
  ref
} from "vue";
import { useCategoryApi } from "@/api/category";
import { Plus, Edit, Delete, Search } from "@element-plus/icons-vue";
import { Session } from "@/utils/storage";
import { ElMessage, ElMessageBox, FormInstance, FormRules } from "element-plus";
import { useI18n } from "vue-i18n";

// loadings
const loading = reactive({
  addLoading: false,
  updateLoading: false,
  tableLoading: false
})
// 搜索 部分
const searchData = reactive({
  categpry: '',
});
// 搜索 分类列表
const categoryList = ref()

// 表单校验
const ruleFormRef = ref<FormInstance>();
const checkCategoryName = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error(t("tips.rules.checkCategoryName")));
  } else {
    if (addCategoryParms.value.category_name !== "") {
      if (!ruleFormRef.value) return;
      ruleFormRef.value.validateField("checkPass", () => null);
    }
    callback();
  }
};
const checkCategoryNameNew = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback(new Error(t("tips.rules.checkCategoryNameNew")));
  } else {
    if (updateCategoryParms.value.new_category_name !== "") {
      if (!ruleFormRef.value) return;
      ruleFormRef.value.validateField("checkPass", () => null);
    }
    callback();
  }
};

const rules = reactive<FormRules>({
  category_name: [
    { required: true, validator: checkCategoryName, trigger: "blur" }
  ],
  new_category_name: [
    { required: true, validator: checkCategoryNameNew, trigger: "blur" }
  ]
});

// formatter
const { appContext } = getCurrentInstance() as ComponentInternalInstance;
const proxy = appContext.config.globalProperties;
// eslint-disable-next-line no-unused-vars
const formatter = proxy.$formatter;

const { t } = useI18n();
const addCategoryParameters = {
  addCategoryTitle: t("category.addCategory"),
  addCategoryUser: t("category.addCategoryUser"),
  addCategoryName: t("category.addCategoryName")
};
const updateCategoryParameters = {
  updateCategoryTitle: t("category.updateParameters.updateCategoryTitle"),
  updateCategoryUser: t("category.updateParameters.updateCategoryUser"),
  updateCategoryOldName: t("category.updateParameters.updateCategoryOldName"),
  updateCategoryNewName: t("category.updateParameters.updateCategoryNewName")
};
const CategoryListParameters = {
  category_id: t("category.categoryListParameters.category_id"),
  user_id: t("category.categoryListParameters.user_id"),
  category_name: t("category.categoryListParameters.category_name"),
  create_time: t("category.categoryListParameters.create_time"),
  action: t("category.categoryListParameters.action")
};
const tableData = ref();

const CategoryApi = useCategoryApi();

const state = {
  total: 5,
  currentPage: 1,
  pageSize: 10
};
// 获取 分类列表
const getCategoryList = () => {
  const parms = {
    user_id: Session.get("userInfo").user_id,
    Page: state.currentPage,
    Size: state.pageSize,
    category_id: searchData.categpry
  };
  loading.tableLoading = true;
  CategoryApi.getCategoryList(parms)
    .then((res: any) => {
      if (res.code === "2000") {
        tableData.value = res.data;
        state.total = res.totalCount;
        loading.tableLoading = false;
      }
    })
    .catch((err: any) => {
      ElMessage({
        type: "warning",
        message: err
      });
    });
};

// 搜索选择器 分类列表获取
const getcategorySearchList = () => {
  const parms = {
    user_id: Session.get("userInfo").user_id,
  }
  CategoryApi.getCategoryListALL(parms)
    .then((res: any) => {
      if (res.code === '2000') {
        categoryList.value = res.data
        // console.log(categoryList.value)
      }
    })
    .catch((err: any) => {
      ElMessage({
        type: 'warning',
        message: err,
      });
    })
}

// 搜索功能
const searchFun = () => {
  getCategoryList()
}
// 修改分类
const updateDialogFormVisible = ref(false);
const updateCategoryParms = ref({
  user_id: "",
  category_id: "",
  old_category_name: "",
  new_category_name: ""
});
const showUpdateDialog = (userId: any, categoryId: any, categoryName: any) => {
  updateDialogFormVisible.value = true;
  updateCategoryParms.value.user_id = userId;
  updateCategoryParms.value.category_id = categoryId;
  updateCategoryParms.value.old_category_name = categoryName;
};
const updateCategory = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate((valid, fields) => {
    if (valid) {
      const categoryId = updateCategoryParms.value.category_id;
      const parms = {
        user_id: updateCategoryParms.value.user_id,
        category_name: updateCategoryParms.value.new_category_name
      };
      loading.updateLoading = true
      CategoryApi.updataCategory(categoryId, parms)
        .then((res: any) => {
          if (res.code === "2000") {
            // console.log(res);
            updateDialogFormVisible.value = false;
            updateCategoryParms.value.new_category_name = "";
            getCategoryList();
          }
          loading.updateLoading = false
        })
        .catch((err: any) => {
          loading.updateLoading = false

          ElMessage({
            type: "warning",
            message: err
          });
        });
      // console.log(parms);
    } else {
      // console.log("error submit!", fields);
    }
  });
};

// 删除分类
const delCategory = (categoryId: any) => {
  const parms = {
    category_id: categoryId
  };
  ElMessageBox({
    closeOnClickModal: false,
    closeOnPressEscape: false,
    title: t("tips.Title"),
    message: t("tips.delContent"),
    showCancelButton: true,
    confirmButtonText: t("Btns.sure"),
    cancelButtonText: t("Btns.cancle"),
    buttonSize: "default",
    beforeClose: (action, instance, done) => {
      if (action === "confirm") {
        instance.confirmButtonLoading = true;
        instance.confirmButtonText = t("tips.deleting");
        setTimeout(() => {
          CategoryApi.delCategory(parms)
            .then((res: any) => {
              if (res.code === "2000") {
                // console.log(res.data);
                done();
                instance.confirmButtonLoading = false;
              }
            })
            .catch((err: any) => {
              ElMessage({
                type: "warning",
                message: err
              });
              done();
              instance.confirmButtonLoading = false;
            });
        }, 700);
      } else {
        done();
      }
    }
  })
    .then(async () => {
      getCategoryList();
    })
    .catch(() => { });
};

// 添加分类
const dialogFormVisible = ref(false);
const addCategoryParms = ref({
  user_id: Session.get("userInfo").user_id,
  category_name: ""
});
const addCategory = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate((valid) => {
    if (valid) {
      const parms = {
        user_id: addCategoryParms.value.user_id,
        category_name: addCategoryParms.value.category_name
      };
      // console.log(parms);
      loading.addLoading = true
      CategoryApi.addCategory(parms)
        .then((res: any) => {
          if (res.code === "2000") {
            // console.log(res);
            getCategoryList();
            dialogFormVisible.value = false;
            addCategoryParms.value.category_name = "";
          }
          loading.addLoading = false
        })
        .catch((err: any) => {
          loading.addLoading = false
          ElMessage({
            type: "warning",
            message: err.err
          });
        });
    } else {
      // console.log("error submit!");
      return false;
    }
  });
};

// 分页
const currentChange = (currentPage: number) => {
  state.currentPage = currentPage;
  getCategoryList();
};
const handleSizeChange = (size: number) => {
  state.pageSize = size;
};

// 预加载
onBeforeMount(() => {
  getCategoryList();
  getcategorySearchList()
});
</script>

<style lang="scss" scoped>
.set-content {
  cursor: default;
  width: 100%;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

// 添加分类 弹窗样式
.addCategoryDialog {
  width: 450px;
}

.addCategoryInput {
  width: 300px;
}

.el-col {
  margin-top: 10px;
}
</style>
