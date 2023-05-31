<template>
  <el-card>
    <div>
      <el-text class="mx-1" size="default">
        {{ $t("article.articleList") }}
      </el-text>
      <el-card style="margin-top: 10px">
        <el-form label-width="90px" size="default">
          <el-row :gutter="10" class="negtive-marginBottom">
            <el-col :lg="6" :md="8" :sm="12">
              <el-form-item label="文章标题">
                <el-input placeholder="请输入标题" clearable v-model="searchData.title" style="margin-top: -5px"></el-input>
              </el-form-item>
            </el-col>
            <el-col :lg="6" :md="8" :sm="12">
              <el-form-item label="文章分类">
                <div style="margin-top: -6px">
                  <el-select v-model="searchData.categpry" class="m-2" placeholder="Select" clearable>
                    <el-option v-for="item in categoryList" :key="item.category_id" :label="item.category_name"
                      :value="item.category_name" />
                  </el-select>
                </div>
              </el-form-item>
            </el-col>
            <el-col :lg="6" :md="8" :sm="12">
              <el-button type="primary" :icon="Search" @click="searchFun()" title="搜索"></el-button>
            </el-col>
            <el-col :lg="6" :md="8" :sm="12">
              <el-button type="primary" style="float: right" :icon="Plus" @click="$router.push('/add-article')"
                title="新增"></el-button>
            </el-col>
          </el-row>
        </el-form>
      </el-card>
    </div>

    <div style="margin-top: 20px">
      <el-card>
        <el-table :data="tableData" v-loading="loading" style="width: 100%">
          <!-- <el-table-column prop="article_id" :label="ArticleListParameters.article_id" width="160" /> -->
          <el-table-column prop="title" :label="ArticleListParameters.title" show-overflow-tooltip="title" width="400">
            <template #default="scope">
              <el-popover placement="top" trigger="hover" width="300px">
                <div class="set-popper">{{ scope.row.title }}</div>
                <template #reference>
                  <div class="set-content">{{ scope.row.title }}</div>
                </template>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column prop="description" :label="ArticleListParameters.description" width="250">
            <template #default="scope">
              <el-popover placement="top" trigger="hover" width="300px">
                <div class="set-popper">{{ scope.row.description }}</div>
                <template #reference>
                  <div class="set-content">{{ scope.row.description }}</div>
                </template>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column prop="category" :label="ArticleListParameters.categpry" width="150" />
          <el-table-column prop="publish_time" :label="ArticleListParameters.publish_time" width="150" />
          <el-table-column prop="update_time" :label="ArticleListParameters.update_time" width="150" />
          <el-table-column prop="status" :label="ArticleListParameters.status" :formatter="formatter.articleStatus"
            width="100" />

          <el-tableColumn :label="ArticleListParameters.action" min-width="300" align="center">
            <!-- eslint-disable-next-line vue/no-unused-vars -->
            <template #default="scoped">
              <!-------------------------------- 修改新闻按钮 ---------------------------------->
              <el-button :icon="Edit" type="warning" size="default" title="修改"
                @click="updateArticle(scoped.row.article_id)"></el-button>
              <!-------------------------------- 是否推荐 ---------------------------------->
              <span v-if="scoped.row.status === 0 && scoped.row.is_deleted === 0" style="margin-left: 10px">
                <el-button type="danger" @click="
                  ChangeIsPublish(scoped.row.article_id, scoped.row.status)
                  " size="default" :icon="Close" title="未发布" />
              </span>
              <span v-else-if="scoped.row.status === 1 && scoped.row.is_deleted === 0
                " style="margin-left: 10px">
                <el-button type="success" @click="
                  ChangeIsPublish(scoped.row.article_id, scoped.row.status)
                  " size="default" :icon="Check" title="已发布" />
              </span>

              <!------------------------------- 删除用户按钮 ------------------------------>
              <el-button :icon="Delete" type="danger" size="default" title="删除" @click="delArticle(scoped.row.article_id)"
                style="margin-left: 10px"></el-button>
            </template>
          </el-tableColumn>
        </el-table>
      </el-card>
      <div>
        <el-pagination background class="pagination" layout="prev, pager, next" :total="state.total"
          :current-page="state.currentPage" :page-size="state.pageSize" @current-change="currentChange"
          @size-change="handleSizeChange" style="margin-top: 10px" />
      </div>
    </div>

  </el-card>
</template>

<script lang="ts" setup>
import {
  ComponentInternalInstance,
  getCurrentInstance,
  onBeforeMount,
  // eslint-disable-next-line no-unused-vars
  onMounted,
  reactive,
  ref
} from "vue";
import { useArticleApi } from "@/api/articles";
import {
  Plus,
  Edit,
  Delete,
  Search,
  // eslint-disable-next-line no-unused-vars
  Check, Close
} from "@element-plus/icons-vue";
import { Session } from "@/utils/storage";
import { ElMessage, ElMessageBox } from "element-plus";
import router from "@/router";
import { useI18n } from "vue-i18n";
import { useCategoryApi } from "@/api/category";
// 接口注册
const ArticleApi = useArticleApi();

// 搜索 部分
const searchData = reactive({
  title: "",
  categpry: ""
});
// 搜索 分类列表
const categoryApi = useCategoryApi();
const categoryList = ref();

// 用户信息
const userInfo = Session.get("userInfo");

// formatter 部分
const { appContext } = getCurrentInstance() as ComponentInternalInstance;
const proxy = appContext.config.globalProperties;
const formatter = proxy.$formatter;

// 表格数据 加载动画
const loading = ref(false);

const tableData = ref();

// 双语 设置
const { t } = useI18n();
const ArticleListParameters = {
  article_id: t("article.articleListParameters.article_id"),
  title: t("article.articleListParameters.title"),
  description: t("article.articleListParameters.description"),
  categpry: t("article.articleListParameters.category"),
  publish_time: t("article.articleListParameters.publish_time"),
  update_time: t("article.articleListParameters.update_time"),
  status: t("article.articleListParameters.status"),
  action: t("article.articleListParameters.action")
};

// 分页参数
const state = {
  total: 5,
  currentPage: 1,
  pageSize: 10
};
// 获取 文章列表
const getArticleList = () => {
  const parms = {
    author_id: Session.get("userInfo").user_id,
    Page: state.currentPage,
    Size: state.pageSize,
    Order: "publish_time desc",
    category: searchData.categpry,
    title: searchData.title
  };
  loading.value = true;
  ArticleApi.getArticleList(parms)
    .then((res: any) => {
      if (res.code === "2000") {
        tableData.value = res.data;
        state.total = res.totalCount;

        for (let i = 0; i < tableData.value.length; i++) {
          tableData.value[i].publish_time = tableData.value[
            i
          ].publish_time.slice(0, 10);
          tableData.value[i].update_time = tableData.value[i].update_time.slice(
            0,
            10
          );
          if (tableData.value[i].category) {
            tableData.value[i].category = JSON.parse(
              tableData.value[i].category
            );
          }
        }
        loading.value = false;
        // console.log()
      }
    })
    .catch((err: any) => {
      ElMessage({
        type: "warning",
        message: err
      });
    });
};
// 分类列表 获取
const getcategoryList = () => {
  const parms = {
    user_id: userInfo.user_id
  };
  categoryApi
    .getCategoryListALL(parms)
    .then((res: any) => {
      if (res.code === "2000") {
        categoryList.value = res.data;
        // console.log(categoryList.value)
      }
    })
    .catch((err: any) => {
      ElMessage({
        type: "warning",
        message: err
      });
    });
};

// 修改文章
const updateArticle = (articleId: any) => {
  router.push({ name: "UpdateArticle", query: { article_id: articleId } });
};

// 删除文章
const delArticle = (articleId: any) => {
  const parms = {
    article_id: articleId
  };
  ElMessageBox({
    closeOnClickModal: false,
    closeOnPressEscape: false,
    title: "提示",
    message: "确认删除?",
    showCancelButton: true,
    confirmButtonText: "确认",
    cancelButtonText: "取消",
    buttonSize: "default",
    beforeClose: (action, instance, done) => {
      if (action === "confirm") {
        instance.confirmButtonLoading = true;
        instance.confirmButtonText = "正在删除";
        setTimeout(() => {
          ArticleApi.delArticle(parms)
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
      getArticleList();
    })
    .catch(() => { });
};

// 修改发布状态
// eslint-disable-next-line no-unused-vars
const ChangeIsPublish = (articleId: string, status: any) => {
  if (status === 0) {
    const params = {
      article_id: articleId,
      status: 1
    };
    console.log(params);

    loading.value = true
    ArticleApi.updateArticleStatus(params)
      .then((res: any) => {
        if (res.code === '2000') {
          ElMessage({
            type: 'success',
            message: '发布成功',
          });
          getArticleList()
          loading.value = false
        }
      })
      .catch((err: any) => {
        ElMessage({
          type: 'warning',
          message: err,
        })
      })
  } else if (status === 1) {
    const params = {
      article_id: articleId,
      status: 0
    };
    console.log(params);

    loading.value = true
    ArticleApi.updateArticleStatus(params)
      .then((res: any) => {
        if (res.code === '2000') {
          ElMessage({
            type: 'success',
            message: '修改成功',
          });
          getArticleList()
          loading.value = false
        }
      })
      .catch((err: any) => {
        ElMessage({
          type: 'warning',
          message: err,
        })
      })
  } else {
    ElMessage({
      type: "warning",
      message: "err.message"
    });
  }
};

// 查询列表
const searchFun = () => {
  getArticleList();
};

// 分页
const currentChange = (currentPage: number) => {
  state.currentPage = currentPage;
  getArticleList();
};
const handleSizeChange = (size: number) => {
  state.pageSize = size;
};

// 预加载
onBeforeMount(() => {
  getArticleList();
  getcategoryList();
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

.el-col {
  margin-top: 10px;
}
</style>
