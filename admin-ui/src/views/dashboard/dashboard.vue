<template>
  <div>
    <div class="row">
      <div class="col-xl-3 col-sm-6 mb-xl-0 mb-4">
        <div class="card">
          <div class="card-body p-3">
            <div class="row">
              <div class="col-8">
                <!-- 卡片一 -->
                <div class="numbers" v-loading="loading.cards.card1">
                  <p class="text-sm mb-0 text-uppercase font-weight-bold">
                    {{ cardData.card1.title }}
                  </p>
                  <h5 class="font-weight-bolder">{{ cardData.card1.value }} 篇</h5>
                  <p class="mb-0">
                    <span class="text-success text-sm font-weight-bolder">+55%</span>
                    相比上周
                  </p>
                </div>
              </div>
              <div class="col-4 text-end">
                <div class="icon icon-shape bg-gradient-primary shadow-primary text-center rounded-circle">
                  <i class="ni ni-money-coins text-lg opacity-10" aria-hidden="true"></i>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-xl-3 col-sm-6 mb-xl-0 mb-4">
        <div class="card">
          <div class="card-body p-3">
            <div class="row">
              <div class="col-8">
                <!-- 卡片二 -->
                <div class="numbers" v-loading="loading.cards.card2">
                  <p class="text-sm mb-0 text-uppercase font-weight-bold">
                    {{ cardData.card2.title }}
                  </p>
                  <h5 class="font-weight-bolder">{{ cardData.card2.value }}个</h5>
                  <p class="mb-0">
                    <span class="text-success text-sm font-weight-bolder">+3%</span>
                    相比上月
                  </p>
                </div>
              </div>
              <div class="col-4 text-end">
                <div class="icon icon-shape bg-gradient-danger shadow-danger text-center rounded-circle">
                  <i class="ni ni-world text-lg opacity-10" aria-hidden="true"></i>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-xl-3 col-sm-6 mb-xl-0 mb-4">
        <div class="card">
          <div class="card-body p-3">
            <div class="row">
              <div class="col-8">
                <!-- 卡片三 -->
                <div class="numbers" v-loading="loading.cards.card3">
                  <p class="text-sm mb-0 text-uppercase font-weight-bold">
                    今日浏览数
                  </p>
                  <h5 class="font-weight-bolder">+3,462</h5>
                  <p class="mb-0">
                    <span class="text-danger text-sm font-weight-bolder">-2%</span>
                    相比昨天
                  </p>
                </div>
              </div>
              <div class="col-4 text-end">
                <div class="icon icon-shape bg-gradient-success shadow-success text-center rounded-circle">
                  <i class="ni ni-paper-diploma text-lg opacity-10" aria-hidden="true"></i>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-xl-3 col-sm-6">
        <div class="card">
          <div class="card-body p-3">
            <div class="row">
              <div class="col-8">
                <!-- 卡片四 -->
                <div class="numbers" v-loading="loading.cards.card4">
                  <p class="text-sm mb-0 text-uppercase font-weight-bold">
                    今日评论数
                  </p>
                  <h5 class="font-weight-bolder">+103,430</h5>
                  <p class="mb-0">
                    <span class="text-success text-sm font-weight-bolder">+5%</span>
                    相比昨天
                  </p>
                </div>
              </div>
              <div class="col-4 text-end">
                <div class="icon icon-shape bg-gradient-warning shadow-warning text-center rounded-circle">
                  <i class="ni ni-cart text-lg opacity-10" aria-hidden="true"></i>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row mt-4">
      <div class="col-lg-7 mb-lg-0 mb-4">
        <div class="card z-index-2 h-100">
          <div class="card-header pb-0 pt-3 bg-transparent">
            <h6 class="text-capitalize">Sales overview</h6>
            <p class="text-sm mb-0">
              <i class="fa fa-arrow-up text-success"></i>
              <span class="font-weight-bold">4% more</span> in 2021
            </p>
          </div>
          <div class="card-body p-3">
            <div class="chart">
              <canvas id="chart-line" class="chart-canvas" height="300"></canvas>
            </div>
          </div>
        </div>
      </div>
      <div class="col-lg-5">
        <div class="card card-carousel overflow-hidden h-100 p-0">
          <div id="carouselExampleCaptions" class="carousel slide h-100" data-bs-ride="carousel">
            <div class="carousel-inner border-radius-lg h-100">
              <!-- <div class="carousel-item h-100 active" style="
                  /* background-image: url('/img/carousel-1.jpg'); */
                  /* background-size: cover; */
                ">
                <div class="carousel-caption d-none d-md-block bottom-0 text-start start-0 ms-5">
                  <div class="icon icon-shape icon-sm bg-white text-center border-radius-md mb-3">
                    <i class="ni ni-camera-compact text-dark opacity-10"></i>
                  </div>
                  <h5 class="text-white mb-1">Get started with Argon</h5>
                  <p>
                    There’s nothing I really wanted to do in life that I wasn’t
                    able to get good at.
                  </p>
                </div>
              </div> -->
              <div class="carousel-item h-100" style="
                  background-image: url('/img/carousel-2.jpg');
                  background-size: cover;
                ">
                <div class="carousel-caption d-none d-md-block bottom-0 text-start start-0 ms-5">
                  <div class="icon icon-shape icon-sm bg-white text-center border-radius-md mb-3">
                    <i class="ni ni-bulb-61 text-dark opacity-10"></i>
                  </div>
                  <h5 class="text-white mb-1">
                    Faster way to create web pages
                  </h5>
                  <p>
                    That’s my skill. I’m not really specifically talented at
                    anything except for the ability to learn.
                  </p>
                </div>
              </div>
              <div class="carousel-item h-100" style="
                  background-image: url('/img/carousel-3.jpg');
                  background-size: cover;
                ">
                <div class="carousel-caption d-none d-md-block bottom-0 text-start start-0 ms-5">
                  <div class="icon icon-shape icon-sm bg-white text-center border-radius-md mb-3">
                    <i class="ni ni-trophy text-dark opacity-10"></i>
                  </div>
                  <h5 class="text-white mb-1">
                    Share with us your design tips!
                  </h5>
                  <p>
                    Don’t be afraid to be wrong because you can’t learn anything
                    from a compliment.
                  </p>
                </div>
              </div>
            </div>
            <button class="carousel-control-prev w-5 me-3" type="button" data-bs-target="#carouselExampleCaptions"
              data-bs-slide="prev">
              <span class="carousel-control-prev-icon" aria-hidden="true"></span>
              <span class="visually-hidden">Previous</span>
            </button>
            <button class="carousel-control-next w-5 me-3" type="button" data-bs-target="#carouselExampleCaptions"
              data-bs-slide="next">
              <span class="carousel-control-next-icon" aria-hidden="true"></span>
              <span class="visually-hidden">Next</span>
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- <div class="row mt-4">
      <div class="col-lg-7 mb-lg-0 mb-4">
        <div class="card">
          <div class="card-header pb-0 p-3">
            <div class="d-flex justify-content-between">
              <h6 class="mb-2">Sales by Country</h6>
            </div>
          </div>
          <div class="table-responsive">
            <table class="table align-items-center">
              <tbody>
                <tr>
                  <td class="w-30">
                    <div class="d-flex px-2 py-1 align-items-center">
                      <div>
                        <img
                          src="@/assets/img/icons/flags/US.png"
                          alt="Country flag"
                        />
                      </div>
                      <div class="ms-4">
                        <p class="text-xs font-weight-bold mb-0">Country:</p>
                        <h6 class="text-sm mb-0">United States</h6>
                      </div>
                    </div>
                  </td>
                  <td>
                    <div class="text-center">
                      <p class="text-xs font-weight-bold mb-0">Sales:</p>
                      <h6 class="text-sm mb-0">2500</h6>
                    </div>
                  </td>
                  <td>
                    <div class="text-center">
                      <p class="text-xs font-weight-bold mb-0">Value:</p>
                      <h6 class="text-sm mb-0">$230,900</h6>
                    </div>
                  </td>
                  <td class="align-middle text-sm">
                    <div class="col text-center">
                      <p class="text-xs font-weight-bold mb-0">Bounce:</p>
                      <h6 class="text-sm mb-0">29.9%</h6>
                    </div>
                  </td>
                </tr>
                <tr>
                  <td class="w-30">
                    <div class="d-flex px-2 py-1 align-items-center">
                      <div>
                        <img
                          src="@/assets/img/icons/flags/DE.png"
                          alt="Country flag"
                        />
                      </div>
                      <div class="ms-4">
                        <p class="text-xs font-weight-bold mb-0">Country:</p>
                        <h6 class="text-sm mb-0">Germany</h6>
                      </div>
                    </div>
                  </td>
                  <td>
                    <div class="text-center">
                      <p class="text-xs font-weight-bold mb-0">Sales:</p>
                      <h6 class="text-sm mb-0">3.900</h6>
                    </div>
                  </td>
                  <td>
                    <div class="text-center">
                      <p class="text-xs font-weight-bold mb-0">Value:</p>
                      <h6 class="text-sm mb-0">$440,000</h6>
                    </div>
                  </td>
                  <td class="align-middle text-sm">
                    <div class="col text-center">
                      <p class="text-xs font-weight-bold mb-0">Bounce:</p>
                      <h6 class="text-sm mb-0">40.22%</h6>
                    </div>
                  </td>
                </tr>
                <tr>
                  <td class="w-30">
                    <div class="d-flex px-2 py-1 align-items-center">
                      <div>
                        <img
                          src="@/assets/img/icons/flags/GB.png"
                          alt="Country flag"
                        />
                      </div>
                      <div class="ms-4">
                        <p class="text-xs font-weight-bold mb-0">Country:</p>
                        <h6 class="text-sm mb-0">Great Britain</h6>
                      </div>
                    </div>
                  </td>
                  <td>
                    <div class="text-center">
                      <p class="text-xs font-weight-bold mb-0">Sales:</p>
                      <h6 class="text-sm mb-0">1.400</h6>
                    </div>
                  </td>
                  <td>
                    <div class="text-center">
                      <p class="text-xs font-weight-bold mb-0">Value:</p>
                      <h6 class="text-sm mb-0">$190,700</h6>
                    </div>
                  </td>
                  <td class="align-middle text-sm">
                    <div class="col text-center">
                      <p class="text-xs font-weight-bold mb-0">Bounce:</p>
                      <h6 class="text-sm mb-0">23.44%</h6>
                    </div>
                  </td>
                </tr>
                <tr>
                  <td class="w-30">
                    <div class="d-flex px-2 py-1 align-items-center">
                      <div>
                        <img
                          src="@/assets/img/icons/flags/BR.png"
                          alt="Country flag"
                        />
                      </div>
                      <div class="ms-4">
                        <p class="text-xs font-weight-bold mb-0">Country:</p>
                        <h6 class="text-sm mb-0">Brasil</h6>
                      </div>
                    </div>
                  </td>
                  <td>
                    <div class="text-center">
                      <p class="text-xs font-weight-bold mb-0">Sales:</p>
                      <h6 class="text-sm mb-0">562</h6>
                    </div>
                  </td>
                  <td>
                    <div class="text-center">
                      <p class="text-xs font-weight-bold mb-0">Value:</p>
                      <h6 class="text-sm mb-0">$143,960</h6>
                    </div>
                  </td>
                  <td class="align-middle text-sm">
                    <div class="col text-center">
                      <p class="text-xs font-weight-bold mb-0">Bounce:</p>
                      <h6 class="text-sm mb-0">32.14%</h6>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <div class="col-lg-5">
        <div class="card">
          <div class="card-header pb-0 p-3">
            <h6 class="mb-0">Categories</h6>
          </div>
          <div class="card-body p-3">
            <ul class="list-group">
              <li
                class="list-group-item border-0 d-flex justify-content-between ps-0 mb-2 border-radius-lg"
              >
                <div class="d-flex align-items-center">
                  <div
                    class="icon icon-shape icon-sm me-3 bg-gradient-dark shadow text-center"
                  >
                    <i class="ni ni-mobile-button text-white opacity-10"></i>
                  </div>
                  <div class="d-flex flex-column">
                    <h6 class="mb-1 text-dark text-sm">Devices</h6>
                    <span class="text-xs">
                      250 in stock,
                      <span class="font-weight-bold">346+ sold</span>
                    </span>
                  </div>
                </div>
                <div class="d-flex">
                  <button
                    class="btn btn-link btn-icon-only btn-rounded btn-sm text-dark icon-move-right my-auto"
                  >
                    <i class="ni ni-bold-right" aria-hidden="true"></i>
                  </button>
                </div>
              </li>
              <li
                class="list-group-item border-0 d-flex justify-content-between ps-0 mb-2 border-radius-lg"
              >
                <div class="d-flex align-items-center">
                  <div
                    class="icon icon-shape icon-sm me-3 bg-gradient-dark shadow text-center"
                  >
                    <i class="ni ni-tag text-white opacity-10"></i>
                  </div>
                  <div class="d-flex flex-column">
                    <h6 class="mb-1 text-dark text-sm">Tickets</h6>
                    <span class="text-xs">
                      123 closed,
                      <span class="font-weight-bold">15 open</span>
                    </span>
                  </div>
                </div>
                <div class="d-flex">
                  <button
                    class="btn btn-link btn-icon-only btn-rounded btn-sm text-dark icon-move-right my-auto"
                  >
                    <i class="ni ni-bold-right" aria-hidden="true"></i>
                  </button>
                </div>
              </li>
              <li
                class="list-group-item border-0 d-flex justify-content-between ps-0 mb-2 border-radius-lg"
              >
                <div class="d-flex align-items-center">
                  <div
                    class="icon icon-shape icon-sm me-3 bg-gradient-dark shadow text-center"
                  >
                    <i class="ni ni-box-2 text-white opacity-10"></i>
                  </div>
                  <div class="d-flex flex-column">
                    <h6 class="mb-1 text-dark text-sm">Error logs</h6>
                    <span class="text-xs">
                      1 is active,
                      <span class="font-weight-bold">40 closed</span>
                    </span>
                  </div>
                </div>
                <div class="d-flex">
                  <button
                    class="btn btn-link btn-icon-only btn-rounded btn-sm text-dark icon-move-right my-auto"
                  >
                    <i class="ni ni-bold-right" aria-hidden="true"></i>
                  </button>
                </div>
              </li>
              <li
                class="list-group-item border-0 d-flex justify-content-between ps-0 border-radius-lg"
              >
                <div class="d-flex align-items-center">
                  <div
                    class="icon icon-shape icon-sm me-3 bg-gradient-dark shadow text-center"
                  >
                    <i class="ni ni-satisfied text-white opacity-10"></i>
                  </div>
                  <div class="d-flex flex-column">
                    <h6 class="mb-1 text-dark text-sm">Happy users</h6>
                    <span class="text-xs font-weight-bold">+ 430</span>
                  </div>
                </div>
                <div class="d-flex">
                  <button
                    class="btn btn-link btn-icon-only btn-rounded btn-sm text-dark icon-move-right my-auto"
                  >
                    <i class="ni ni-bold-right" aria-hidden="true"></i>
                  </button>
                </div>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div> -->
  </div>
</template>

<script lang="ts" setup>
// eslint-disable-next-line no-unused-vars
import { onMounted, reactive, ref } from 'vue';
import { useCategoryApi } from "@/api/category";
import { useArticleApi } from "@/api/articles";
import { Session } from '@/utils/storage';
// eslint-disable-next-line no-unused-vars
import { ElMessage } from 'element-plus';

// 接口引入
// eslint-disable-next-line no-unused-vars
const UseCategoryApi = useCategoryApi()
// eslint-disable-next-line no-unused-vars
const UseArticleApi = useArticleApi()

// 加载状态
const loading = reactive({
  cards: {
    card1: false,
    card2: false,
    card3: false,
    card4: false
  }
})

// 卡片内容
const cardData = reactive({
  card1: {
    title: "文章数量",
    value: 0
  },
  card2: {
    title: "分类数量",
    value: 0
  },
})

// 卡片数据 接口调用
const getCardDatas = () => {
  const params = {
    author_id: Session.get("userID")
  }
  loading.cards.card1 = true
  loading.cards.card2 = true
  loading.cards.card3 = true
  loading.cards.card4 = true

  // 文章数量
  UseArticleApi.getArticleCount(params)
    .then((res: any) => {
      if (res.code === '2000') {
        cardData.card1.value = res.data.count
      }
      loading.cards.card1 = false
    })
    .catch((err: any) => {
      ElMessage({
        type: 'warning',
        message: err.message,
      });
    })
  // 分类数量
  // UseCategoryApi.(params)
  // .then((res: any) => {
  //   if (res.code === '2000') {
  //     cardData.card1.value = res.data.count
  //   }
  // })
  // .catch((err: any) => {
  //   ElMessage({
  //     type: 'warning',
  //     message: err.message,
  //   });
  // })
}

// 预加载
onMounted(() => {
  getCardDatas()
})
</script>

<style lang="scss" scoped></style>
