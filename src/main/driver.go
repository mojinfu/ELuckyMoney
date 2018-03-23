package main
import (
	"ELuckyMoney/src/UselessDownload"
	"ELuckyMoney/src/UselessHelper"
	"fmt"
	"math/rand"
	"encoding/json"
	"github.com/rlds/rlog"
	"net/url"
	"strconv"
	"time"
	"strings"
)

func (this *registerStruct)converjson()[]byte{
	body, err := json.Marshal(this)
	if err != nil {
		rlog.V(1).Info("func converjson error:" + err.Error())
		return nil
	}
	return body
}
func (this *luckydoorStruct)converjson()[]byte{
	body, err := json.Marshal(this)
	if err != nil {
		rlog.V(1).Info("func converjson error:" + err.Error())
		return nil
	}
	return body
}
func areyourealhungry(myHungryMan *HungryManStruct)bool{
	if len(myHungryMan.WechatCookie)==0{
		return false
	}
	cookieAfterUrldecode,isdecodeOk :=url.QueryUnescape(myHungryMan.WechatCookie)
	if isdecodeOk!=nil{
		return false
	}
	rlog.V(1).Info("cookieAfterUrldecode:"+cookieAfterUrldecode)
	myHungryMan.openid = UselessHelper.GetJsonValue(cookieAfterUrldecode , "openid")
	myHungryMan.eleme_key = UselessHelper.GetJsonValue(cookieAfterUrldecode , "eleme_key")
	myHungryMan.nickname = ""
	rlog.V(1).Info("myHungryMan:"+myHungryMan.nickname)
	
	if !myHungryMan.realHungry&&len(myHungryMan.Phone)!=11{
		myHungryMan.Phone =getAFakeNum()
	}
	if len(myHungryMan.eleme_key)!=0&&len(myHungryMan.openid)!=0{
		myHungryMan.realHungry=true
		return true
	}
	return false
}
func getAFakeNum()string{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return "1575711"+fmt.Sprintf("%d",r.Intn(10))+
	fmt.Sprintf("%d",r.Intn(10))+
	fmt.Sprintf("%d",r.Intn(10))+
	fmt.Sprintf("%d",r.Intn(10))
}
func getAFakeName()string{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return "摸金校尉"+fmt.Sprintf("%d",r.Intn(1000))
}
func (myHungryMan *HungryManStruct)register()error{
	var myregister registerStruct = registerStruct{Sign:myHungryMan.eleme_key, Phone:myHungryMan.Phone}
	UselessDownload.Download_PUT_Json(
		fmt.Sprintf(myEOrigin+`/restapi/v1/weixin/%s/phone`,myHungryMan.openid),
		string(myregister.converjson()),
		myHungryMan.WechatCookie)
	//rlog.V(1).Info("registerJosn:"+registerJosn)
	return nil
}

func (myHungryMan *HungryManStruct)luckydoor(myELuckyMoney *ELuckyMoneyStruct)(string,error){
	var myLuckydoor luckydoorStruct
	myLuckydoor.Device_id=""
	myLuckydoor.Hardware_id=""
	myLuckydoor.Method = "phone"
	myLuckydoor.Group_sn= myELuckyMoney.sn
	myLuckydoor.Phone = myHungryMan.Phone
	myLuckydoor.Platform = 0
	myLuckydoor.Sign = myHungryMan.eleme_key
	myLuckydoor.Track_id = "undefined"
	myLuckydoor.Unionid ="o_PVDuIDTHd67WEarVE7vWW22XdM"
	myLuckydoor.Weixin_avatar ="http://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJYk0mY3BuCX8sfBkQwxAqvdniawPnqJXNx2JuXNibMicvzdjz7BgfsibYcibfG6uFhKpKLSVLEEEugTKQ/0"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if len(myHungryMan.nickname)==0{
		myLuckydoor.Weixin_username = "摸金校尉"+fmt.Sprintf("%d",r.Intn(1000))
	}else{
		myLuckydoor.Weixin_username = myHungryMan.nickname
	}
	
	luckydoorJosn , _:=UselessDownload.Download_POST_Json(
		fmt.Sprintf(myEOrigin+`/restapi/marketing/promotion/weixin/%s`,myHungryMan.openid),
		string(myLuckydoor.converjson()),
		myELuckyMoney.sn,
		myHungryMan.WechatCookie)
	return luckydoorJosn,nil
}

func (myELuckyMoney *ELuckyMoneyStruct)Parse_E_LuckyMoney()error{
	myELuckyMoney.E_Url = strings.Replace(myELuckyMoney.E_Url,"#","&",1)
	myE_Url , idParseOK:=url.Parse(myELuckyMoney.E_Url)
	if idParseOK!=nil{
		return idParseOK
	}
	myELuckyMoney.sn=myE_Url.Query().Get("sn")
	lucky_number,err:=strconv.Atoi(myE_Url.Query().Get("lucky_number"))  
	if err!=nil{
		return err
	}
	myELuckyMoney.lucky_number = lucky_number
	return nil
}