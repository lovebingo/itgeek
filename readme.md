项目地址:[https://github.com/ecdiy/itgeek]

## 安装
### 安装数据库
```angular2html
新建库XXX 导入db\geek.sql 库名与参数DbDsn DB名一致
```

###  源码安装
 ```angular2html
go get github.com/ecdiy/gpa
go get github.com/gin-gonic/gin
go get github.com/cihub/seelog
go get github.com/hunterhug/go_image/graphics
go get github.com/ecdiy/itgeek
.....

```
go打包
cd .\gk\cmd
go build  -ldflags "-s -w" -o geek.exe
 
 ```angular2html
 cd ui
 npm install
 
 后台管理打包
 npm run build-admin
 web打包
 npm run build-web
 H5打包
 npm run build-m
```
### build.cmd 打包所有
```
打包windows与linux下所有可以运行程序与资源
```

### IDE开发
```angular2html
设置nginx 代理webpack项目
intelj导入项目 
..... 此处省略N多 类似nginx,intelj设置之类
```


 
 ### 可执行包
 [https://github.com/ecdiy/itgeek/releases]
 下最新版

#### *** 源码安装与可执行包两者选一即可 ***

#### 参数说明
```angular2html
BindAddr=:9000
UploadDir=./upload/  #上传图片的目录
DbDriver=mysql       #数据库驱动名称
DbDsn=root:root@tcp(127.0.0.1:3306)/gk?timeout=30s&charset=utf8mb4&parseTime=true  
ImgHost=http://192.168.x.x:9000  #不能设置成127.0.0.1
MultiSite=0         #这个参数不要设置成1，这个多站点标识，如果要设置成多站点与开发者单独联系。



以上是默认值，如果与之一样可以不用配置
如:
 
geek BindAddr=:80 ImgHost=http://192.168.x.x DbDsn=root:root@tcp(127.0.0.1:3306)/gkx?timeout=30s&charset=utf8mb4&parseTime=true

端口号 
ImgHost不能设置成127.0.0.1，否则会导致上传图片有路径问题，不能正确显示图片
数据库


```

#### 下一步
```angular2html 
http://192.168.x.x/admin
管理你的站点信息，分类需要设置，否则发不了主题
```

#### 功能详情介绍：
[http://itgeek.top/p/topic/list,300,308,1] 


#### 目录结构
```angular2html
      gk                    后端GO代码
      db                    数据库脚本
      ui/web                pc web端
      ui/m                  h5
      tools                 工具    
      tools/post-update     编译参考，可以放在.git/hooks/ 在push代码后自动编译，有前后端的编译,ITGeek.top网站用的自动发布脚本，有需要的可以参考。
      tools/nginx-dev.conf  开发时nginx配置
      tools/post-update     开发环境nginx配置
```

#### 其它工具
[https://nodejs.org/en/]



#### 数据库增量升级: 
/db/geek.sql是当前版本下的数据库脚本，如果需要做增量升级的，参考这个主题
[http://www.itgeek.top/p/topic/detail,32,10021]
 
#### 使用过程有问题请到 [http://www.itgeek.top] 相应的板块反馈，截图。

#### 功能列表
|功能分类|子功能|
|-|-|
|主题|发表, 附加, 评论, 感谢,收藏|
|用户|设置，上传头像，登录奖励，个人主页|
|上传图片|压缩|
|积分|发表（-20）, 附加（-20）, 评论+-5 , 感谢+-10|
|消息|评论|
|笔记|分类，笔记|
|分享|微博|



#### 网站Web效果

![image](https://github.com/ecdiy/itgeek/blob/master/doc/web.gif?raw=true)

#### 网站H5效果

![image](https://github.com/ecdiy/itgeek/blob/master/doc/h5.gif?raw=true)

#### 后台管理

![image](https://github.com/ecdiy/itgeek/blob/master/doc/admin.gif?raw=true)


#### FAQ
##### baidu统计代码，修改发表后的index.html,找到对应位置,注意有部分代码修改

#### QQ群/收录申请：620063196

#### 使用该项目搭建的网站  
http://www.itgeek.top



#### If you find this project helpful, maybe you can buy me a coffee 
![image](https://github.com/ecdiy/itgeek/blob/master/doc/wxpay.jpg?raw=true)
