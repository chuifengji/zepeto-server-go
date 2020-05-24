# MindaZepeto　Server 简单说明

## 目录说明

config :配置文件
models:数据模型
route:路由
dbConnect:数据库连接
repoHandler:数据库层事件处理
controller:路由层事件处理
util:公共项（备用）

## 数据库模型说明

1. Decoration

- id  
- url

2. User

- openid
- avatar
- name
- college
- major
- class
- personnal_image

3. Background

- id  
- name  
- url
 
 4. Appearance_controllers

 - id
 - type
 - url

## 路由

1. decoration/add
添加饰品
2. decoration/delete
删除饰品
3. decoration/modify
修改饰品图片
4. decoration/get-list
获取饰品列表


5. background/add

添加背景

6. background/delete

删除背景

7. background/modify

修改背景图片

8. background/get-bg-list

获取背景列表

9. appearance/add

增加外貌小物件

10. appearance/modify

修改外貌小物件

11. appearance/delete
删除外貌小物件

12. personal/insert-personal-info

添加个人信息

13. personal/update-personal-image

添加、修改个人形象图片
（此处需要借助七牛云sdk,直接修改url还是修改url对应的图片呢？未定）

14. personal/get-friends-list

获取朋友列表

15. personal/get-self-info

获取个人信息

16. personal/get-search

-合照肯定要保存在用户自己手机上了，不然我们的服务器撑不住。
待确定：模糊搜索？加入班级？

