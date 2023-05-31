# IPFS_HUGO_Blog

#### introduction

IPFS_HUGO_Blog is a decentralized, always-on static Blog management and publishing system. The static Blog generation part is based on the Hugo system and includes a web-based publishing and management interface that makes it easy for users to publish and manage their own Blog documents. The static Blog pages are stored on IPFS distributed network nodes, and users can access the Blog content using a browser based on the IPFS network protocol. Once published, Blog content is permanently stored on IPFS distributed network nodes and is not owned by any platform. The Blog system is based on Hugo, and users can easily switch to their preferred template theme.

The Blog management and publishing system uses a front-end and back-end separation model. The front-end program can also be deployed on the IPFS network after compilation for user access. The front-end program interface is provided by an “article compilation and publishing service” (which needs to be centrally deployed with a lightweight article database). The article database only records metadata related to Blog articles for indexing and article classification retrieval.

#### System features and advantages.
1.  Easy to use: A web-based management interface has been added for publishing and managing documents；
2.  Fast: The speed of generating a static website is very fast, and a website containing thousands of pages can be generated in a few seconds;
3.  Flexible: Hugo supports multiple themes and plugins, allowing for customization and extension as needed；
4.  Distributed: The server-side generated static website and documents based on Hugo can be easily deployed to IPFS nodes without requiring centralized server storage of pages and documents；
5.  Convenient deployment: A deployment method based on Docker is provided, which can quickly deploy IPFS nodes and server-side processing programs；


#### Installation Tutorial
#### 一、Docker based rapid deployment
1. Access the target deployment server and ensure that docker and docker-compose tools are installed on the server.
2. Download the project and navigate to the project directory.

   `git clone github地址`

3. update the config file: `configs/config.yaml` ,ipfs url config，So that it can be accessed via the container name.
    ```
    # ipfs url
    ipfs:
      Url: http://ipfs:5001
    ```
4. update the config file:`configs/database.yaml`，check the database config:
    ```
   mysql:
     database: hugo-blog
     username: root
     password: admin123
     dsn: root:admin123@tcp(ipfs-mysql:3306)/hugo-blog?charset=utf8&parseTime=True&loc=Local
   ```
5. run `docker-compose up -d`,start the server.
6. test：See instructions for use. The project is based on the default IPFS gateway for access.
   

#### 二、Manual deployment
1. Access the deployment server and install the Go runtime environment. For reference on configuring Go on CentOS, please see this: 
   ```
   yum install epel-release
   yum repolist
   yum install -y golang
   go version
   ```
2. Download the project and enter the project directory.
3. Compile the Go program，run:`go build`,so can Generate the executable file，default file is `IPFS-Blog-Hugo.exe`
4. isntall ipfs node，refer to:[https://blog.csdn.net/sgl520lxl/article/details/125932213]()
5. After starting the IPFS node, run the project, first set the permissions of the IPFS-Blog-Hugo file with chmod 744 IPFS-Blog-Hugo, and then run nohup ./IPFS-Blog-Hugo &.
6. Project testing: See instructions for use. The project is based on the default IPFS gateway for access.


#### Instructions for use
##### Project Technical Description
- The project has a service program that compiles and packages Blog documents using Hugo on a schedule, generating the public folder, which is stored in resources/public in the project.
- Then, publish the folder through IPFS using ipfs add [dir]. IPFS publishes a local copy first, and then the webpage can be accessed through the default gateway ipfs.io.
- The project generates runtime logs in the logs folder, and the published CID can be found by checking the logs.
- The Blog can be accessed directly through the default gateway: http://ipfs.io/ipfs/[cid] or http://ipfs.io/ipns/[cid].
- Access can also be done through an IPNS link, with specific operations described in the following text.

##### Browse Blog
1. The project has default templates and articles at the start, and there is a scheduled task to package and publish the project upon startup. To access the Hugo-published blog, it is necessary to access it through IPFS. The specific CID can be viewed through the logs.
   ```
   {"level":"info","msg":"[Blog published with CID: QmYGQaqgUmkZkcSNfNZsoHehmbK7xaSgzG46J4GHFeYicR\n]","time":"2023-05-26 18:34:01"}
   ```
2. To access it, simply enter ipfs://QmYGQaqgUmkZkcSNfNZsoHehmbK7xaSgzG46J4GHFeYicR in a browser that supports IPFS (such as Brave). However, since the blog may need to be modified, the CID may change. Therefore, the IPNS method can be used instead.

3. For project publishing, a specific access address needs to be configured, which requires the use of IPNS technology. The service program in the project has already completed the IPNS publishing process, and the IPNS link needs to be obtained before configuration. This can be done by running ipfs name publish [CID] in the command-line environment, with the CID obtained from the log file. The command is as follows:
   ```
   [root@ecs-3180 ~]# ipfs name publish QmTYbcVWwC418hGTzbY2nQu6ZYSXwon4r5YCddZv9JjhhN
   Published to k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s: /ipfs/QmTYbcVWwC418hGTzbY2nQu6ZYSXwon4r5YCddZv9JjhhN
   ```

4. so we can get the ipns address：`k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s`，you can access it using an IPNS address in the Brave browser：`ipns://k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s/`；

 The management system address for the Blog is:`ipns://k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s/admin/`

5. you can access through the ipfs gateway: `http://ipfs.io/ipns/k51qzi5uqu5dkb2xif9u7uithy4a5j9nua2aofhw7aj2k1f8ugk3cq0l18313s`

#### Frequently Asked Questions and Configuration Solutions
##### 1. Blog domain name configuration
If for the sake of convenience and promotion, a domain name can be configured for the blog project. Additionally, a personal IPFS gateway can be deployed and the corresponding firewall port (default is 8080) can be opened for quick access. The specific steps are as follows:
1. Obtain a domain name, assuming the authoritative domain name is ipfs.com.
2. Configure DNS resolution.

| Host Record   | RECORD | record value                  |
|---------------|------|-----------------------------|
| blog          | A    | xxx.xxx.xxx.xxx（your host）       |
| _dnslink.blog | TXT | dnslink=/ipns/（your ipns cid）  |

3. Configure nginx (if you want to configure it yourself)

The proxy_pass field is configured with an IP address and port number, where the port number is the IPFS gateway port 8080. This can be modified as needed using ipfs config edit. Open the corresponding port in the firewall and security group 

the nginx configuration like this:
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
With this, the configuration is complete and the Blog can be accessed directly through the domain name： 

the web site： `http://blog.ipfs.com` 

management system： `http://blog.ipfs.com/admin`

This publishing solution may still have a certain degree of centralization, but it speeds up response times and can be configured according to needs.

##### 2. IPFS request header configuration

By default, the IPFS gateway has restrictions on cross-domain issues and limited HTTP request methods. The following steps can be taken to modify these restrictions：
1. Enter ipfs config edit in the command line. If an error occurs stating “Error: ENV variable $EDITOR not set”, you can enter export EDITOR=/usr/bin/vim first.
2. update the Gateway config，like this：
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

3. stop the ipfs node, and run:`ipfs daemon`

#### How to Contribution:
<hr/>

1. Fork this repository
2. Create a new branch called Feat_xxx
3. Submit your code changes
4. Create a new Pull Request




#### Donation and Support
<hr/>

If you find this project helpful for your work, you can donate and support us by scanning the QR code below. Your encouragement will motivate us to make it even better! 
<div align="center">
<img src=".github/images/donate.png" alt="donate" width="200" height="200" />
<p>if you like our projects, buy us a coffee</p>
</div>
<div align="center">
<img src=".github/images/support.png" alt="support" width="200" height="200" />
<p>if you like our projects, support us with Alipay</p>
</div>