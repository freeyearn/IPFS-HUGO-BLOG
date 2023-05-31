const path = require("path");

const webpack = require("webpack");

function resolve(dir) {
  return path.join(__dirname, dir);
}
const name = "vue Admin Template"; // page title
const port = process.env.port || process.env.npm_config_port || 9528; // dev port

module.exports = {
  publicPath: "/admin/",

  outputDir: "../resources/static/admin",
  assetsDir: "static",
  productionSourceMap: false,

  devServer: {
    port: port,
    open: false,
    proxy: {
      "/api": {
        target: "http://139.9.128.8:8000", // 实际跨域请求的API地址
        secure: false, // https请求则使用true
        ws: true,
        changeOrigin: true // 跨域
      }
    }
  },
  configureWebpack: {
    name: name,
    resolve: {
      alias: {
        "@": resolve("src")
      }
    },
    plugins: [
      // 配置 jQuery 插件的参数
      new webpack.ProvidePlugin({
        $: "jquery",
        jQuery: "jquery",
        "window.jQuery": "jquery",
        Popper: ["popper.js", "default"]
      })
    ]
  }
};
