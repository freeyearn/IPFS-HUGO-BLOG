<<<<<<< HEAD
<h1 align="center"><img src="https://github.com/freeyearn/IPFS-HUGO-BLOG/blob/dev/.github/images/ipfs.png?raw=true" height="30" width="30"> IPFS_HUGO_BLOG</h1>
=======
<h1 align="center"><img src=".github/images/ipfs.png?raw=true" height="30" width="30"> IPFS_HUGO_BLOG</h1>
>>>>>>> dev

<div align="center">

[English](README.en.md) / 简体中文

Quickly and freely deploy your IPFS-based Blog system to build an always-on article publishing website.

快速免费部署你的基于IPFS的Blog系统，搭建永不down机的文章发布网站

[Demo](http://libertypress.quwancode.com/) / [Issues](https://github.com/freeyearn/IPFS-HUGO-BLOG/issues) / [Buy Me a Coffee](https://www.buymeacoffee.com/iDealStudio)

[演示](http://libertypress.quwancode.com/) / [反馈](https://github.com/freeyearn/IPFS-HUGO-BLOG/issues) / [打赏开发者](.github/images/support.png)


![cover](.github/images/cover.png?raw=true)

</div>


# IPFS_HUGO_Blog

#### 项目介绍

IPFS_HUGO_Blog一款去中心化的，永不down机的静态Blog管理发布管理系统；静态Blog生成部分，基于hugo系统，增加了基于web的发布管理界面，方便用户发布管理自己的Blog文档；静态Blog页面存储到IPFS分布式网络节点上，用户可以使用基于IPFS网络协议的浏览器访问Blog内容；Blog内容一经发布，将永久存储于IPFS分布式网络节点，不属于任何平台；Blog系统基于Hugo，用户可以方便地更换自己喜欢的模板主题； 

Blog管理发布系统采用前后端分离模式，前端程序编译后也可以部署在IPFS网络上，供用户访问；前端程序接口由一个“文章编译发布服务”提供（该服务和一个轻量级的文章数据库需要集中部署）；文章数据库仅记录Blog文章相关元数据，用于索引和文章分类检索；

#### 系统特点和优势
1.  易于使用：增加了Web管理界面来发布和管理文档；
2.  快速：生成静态页面网站的速度非常快，可以在几秒钟内生成一个包含数千个页面的网站。
3.  灵活：Hugo支持多种主题和插件，可以根据需要进行自定义和扩展；
4.  分布式：服务器端基于Hugo生成的静态网站和文档，可以轻松地部署到IPFS节点上，不需要集中的服务器存储页面和文档；
5.  方便部署：提供基于Docker的部署方式，能快速地部署IPFS节点和服务器端处理程序；


#### 安装教程
#### 一、Docker方式快速部署
1. 进入部署目标服务器，确保服务器上安装了docker以及docker-compose工具;

2. 在部署目标服务器的根目录创建相应的文件夹；
'''
  mkdir -p /docker/deploy
  mkdir -p /docker/deploy
'''
 
  下载项目并进入项目目录 
  '''
  cd /docker/deploy
  git clone https://github.com/freeyearn/IPFS-HUGO-BLOG.git    
  ''' 

3. 进入代码目录：
  cd /docker/deploy/IPFS-HUGO-BLOG
  
  构建应用的运行环境镜像
  docker build . 

  运行启动服务
  docker-compose up -d

4. 修改配置文件`configs/config.yaml`中的ipfs url设置，通过容器名访问
    ```
    # ipfs url
    ipfs:
      Url: http://ipfs:5001
    ```
5. 根据 docker-compose.yaml 参数，修改相应的配置；
  修改配置文件`configs/database.yaml`，检查数据库配置：
    ```
   mysql:
     database: hugo-blog
     username: root
     password: admin123456
     dsn: root:admin123@tcp(ipfs-mysql:3306)/hugo-blog?charset=utf8&parseTime=True&loc=Local
   ```
   说明：
    范例中使用的是云数据库产品（https://planetscale.com/）；
    可更换为自行安装部署的MySQL数据库；

6. 项目测试：
   见使用说明，项目基于ipfs默认网关来访问
   

#### 二、手动部署
1. 进入到部署服务器，安装Go运行环境；centos下Go配置参考：
   ```
   yum install epel-release
   yum repolist
   yum install -y golang
   go version
   ```
2. 下载项目并进入项目目录
3. 编译go程序，运行`go build`生成可执行文件，默认生成`IPFS-Blog-Hugo`可执行文件
4. 安装ipfs结点，参考[https://blog.csdn.net/sgl520lxl/article/details/125932213]()
5. 启动ipfs结点后，运行项目，先设置`IPFS-Blog-Hugo`文件的权限，`chmod 744 IPFS-Blog-Hugo`，随后运行：`nohup ./IPFS-Blog-Hugo &`
6. 数据库配置：
   将管理系统所需要的数据库，通过存放在“config/databse”目录下的sql脚本文件“ipfs-hugo-blog.sql”恢复；并在 `configs/database.yaml`中做好相应配置：
   ```
   mysql:
     database: hugo-blog
     username: root
     password: admin123
     dsn: your IP
     port：your port
   ```
   说明：
   范例中使用的是云数据库产品（https://planetscale.com/）；
   可更换为自行安装部署的MySQL数据库；
   
7. 项目测试：见使用说明，项目基于ipfs默认网关来访问


#### 使用说明
##### 项目技术说明
- 项目有一个服务程序，负责对HUGO博客文档进行定时编译打包，生成静态BLOG发布文件夹：public，位于：`resources/public`；
- 通过ipfs将该文件夹发布出去`ipfs add [dir]`，ipfs先在本地发布一份，同时会将文档发布到默认网关`ipfs.io`；
- 通过生成对运行日志，在logs文件夹下，可以查看到每次发布产生的cid；
- 博客可以直接通过访问默认网关：`http://ipfs.io/ipfs/[cid]` 或`http://ipfs.io/ipns/[cid]`
- 也可以通过ipns链接来访问，具体操作配置见下文；

##### 查看博客
1. 项目初始有默认的模板与文章，在项目启动时有定时任务来打包发布，想访问hugo发布的博客，需要通过ipfs访问。可以通过logs日志查看具体的CID；
   ```
   {"level":"info","msg":"[Blog published with CID: QmYGQaqgUmkZkcSNfNZsoHehmbK7xaSgzG46J4GHFeYicR\n]","time":"2023-05-26 18:34:01"}
   ```
2. 在支持IPFS的浏览器（Brave等）中输入`ipfs://QmYGQaqgUmkZkcSNfNZsoHehmbK7xaSgzG46J4GHFeYicR`即可访问，但是由于博客需要修改，CID会变化，所以可以采用IPNS方法；

3. 项目发布需要配置一个指定的访问地址，那么就需要使用到ipns的技术，项目中服务程序已经完成了ipns的发布过程，在配置时需要先获取到ipns的链接。可以通过在命令行环境下，运行`ipfs name publish [CID]`，这个CID可以从日志文件中获取。命令如下：
   ```
   [root@ecs-3180 ~]# ipfs name publish QmTYbcVWwC418hGTzbY2nQu6ZYSXwon4r5YCddZv9JjhhN
   Published to k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s: /ipfs/QmTYbcVWwC418hGTzbY2nQu6ZYSXwon4r5YCddZv9JjhhN
   ```

4. 获取到ipns地址：`k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s`，在Brave浏览器中可以通过ipns访问：`ipns://k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s/`；对应的博客管理后台为：`ipns://k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s/admin/`

5. 可以通过ipfs网关访问：`http://ipfs.io/ipns/k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s`

#### 常见问题及配置方案
##### 1. 博客项目域名配置
如果为了方便记忆和推广，可以博客项目配置一个域名。并且自己部署一个ipfs网关，将防火墙对应端口打开（默认为8080），便于快速访问。
具体步骤如下:
1. 获取一个域名，这里假设权威域名为`ipfs.com`
2. 配置DNS解析

| 主机记录       | 记录类型 | 记录值                         |
|---------------|------|-----------------------------|
| blog          | A    | xxx.xxx.xxx.xxx（你的主机）       |
| _dnslink.blog | TXT | dnslink=/ipns/（你的ipns对应cid）  |

3. 配置nginx(如果要自行配置) 

proxy_pass字段配置的ip+端口号，其中端口号为ipfs网关的端口8080，可以根据实际需要，通过`ipfs config edit`进行修改。在防火墙与安全组中打开对应端口。 

nginx配置如下
```nginx
server {
   listen  80;
   server_name blog.ipfs.com;
   
   location / {
    proxy_pass   http://127.0.0.1:8080;
    proxy_set_header    Host       $http_host;
    proxy_set_header    X-Real-IP        $remote_addr;
    proxy_set_header    X-Forwarded-For  $proxy_add_x_forwarded_for;
    proxy_set_header    HTTP_X_FORWARDED_FOR $remote_addr;
    proxy_set_header    X-Forwarded-Proto $scheme;
   }
   
   # api location 
   location /api {
        proxy_pass http://127.0.0.1:8000;
        proxy_set_header Host $host:80;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
   }
}
```
至此配置完成，可以直接通过域名访问： 

前台： `http://blog.ipfs.com` 

发布系统： `http://blog.ipfs.com/admin`

当然这个发布方案可能在某种程度上还存在集中中心化的情况，但是响应速度加快了，可以根据需要进行配置。

##### 2. IPFS请求头配置
ipfs网关默认情况下对于跨域问题有限制，http请求方法有限，通过以下步骤可以进行修改：
1. 在命令行输入`ipfs config edit`。如果有报错：`Error: ENV variable $EDITOR not set`，可以先输入`export EDITOR=/usr/bin/vim`
2. 修改Gateway配置选项，改成如下配置：
   ```json
   "Gateway": {
       "APICommands": [],
       "HTTPHeaders": {
         "Access-Control-Allow-Headers": [
           "Authorization"
         ],
         "Access-Control-Allow-Methods": [
           "GET",
           "POST",
           "PUT",
           "DELETE",
           "OPTIONS"
         ],
         "Access-Control-Allow-Origin": [
           "*"
         ],
         "Access-Control-Allow-Credentials": [
           "true"
         ],
         "Access-Control-Expose-Headers": [
           "Location"
         ]
       },
       "NoDNSLink": false,
       "NoFetch": false,
       "PathPrefixes": [],
       "PublicGateways": null,
       "RootRedirect": ""
     },
   ```

3. 将ipfs节点停止，再重新运行`ipfs daemon`
#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request



#### 捐助和支持
<hr/>

如果您觉得这个项目对您的工作有帮助✨，可以通过扫描下面的二维码进行捐助，鼓励我们将工作做的更好！
<div align="center">
<img src=".github/images/support.png" alt="support" width="200" height="200" />
<p>~支付宝支持一下这个项目👆~</p>
</div>
