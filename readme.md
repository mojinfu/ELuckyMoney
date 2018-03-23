

# 功能
    帮助广大贫困群众，薅饿了么的羊毛。（饿了么拼手气红包领取）

# 使用方法
    服务开启后，会部署一个http服务在配置文件描述的相应端口。
    使用者发送http请求到该服务，服务自动开始领取饿了么红包

## 服务开启所需参数
    配置文件在/bin/config.json 
    HttpHost字段配置服务部署的IP和端口
    Fakehungrymen配置抢饿了么红包的人手，实则为cookie，可自行扩展。

## 使用所需参数

    请求Method：POST

    请求URL：  http://[HttpHost]/e/luckymoney

    请求Header："Content-Type:application/json; charset=UTF-8"

    请求Body：
    {
        "name": "摸金符",                
        "phone": "15967184143",
        "e_url": "https://h5.ele.me/hongbao/?from=singlemessage&isappinstalled=0#hardware_id=&is_lucky_group=True&lucky_number=9&track_id=&platform=0&sn=29e55676462cf442&theme_id=2097&device_id="
    }   

### 请求Body参数说明：
        注意：请求Body参数为一个json
        1. name： 为最终领取红包人的名字，会显示在饿了么红包的页面中。可缺省 可为空
        2. phone：使用者的饿了么红包账户
        3. e_url：一个饿了么红包的链接。 （在浏览器中打开时的URL）

        
