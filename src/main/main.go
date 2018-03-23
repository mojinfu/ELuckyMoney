package main

import (
	"strings"
	 "fmt"
	"time"
	"encoding/json"
	. "ELuckyMoney/src/loadFile"
	"github.com/rlds/rlog"
	"net/http"
	"ELuckyMoney/src/UselessHelper"
	"io/ioutil"
	"bytes"
)
var myFakeHungryMenArr []*HungryManStruct
func main (){

	if Loadconf() {
		UselessHelper.MkAlldir(PrivateConf.LogDir)
		rlog.LogInit(PrivateConf.LogLevel, PrivateConf.LogDir, PrivateConf.MaxLogLen_m, 1)

		for _,myFakeHungryMenCookie :=range PrivateConf.Fakehungrymen{
			myFakeHungryMan:=&HungryManStruct{WechatCookie: myFakeHungryMenCookie}
			if !areyourealhungry(myFakeHungryMan){
				continue 
			}
			myFakeHungryMan.register()
			time.Sleep(time.Second*1)
			myFakeHungryMenArr=append(myFakeHungryMenArr,myFakeHungryMan)
		}
		myFakeLastEmptyHungryMan:=&HungryManStruct{WechatCookie: "2344",openid:"lastone"}
		myFakeHungryMenArr=append(myFakeHungryMenArr,myFakeLastEmptyHungryMan)
		profServeMux := http.NewServeMux()
		profServeMux.HandleFunc("/e/luckymoney", Get_E_LuckyMoney)
		server := &http.Server{
			Addr:           PrivateConf.HttpHost,
			Handler:        profServeMux,
			ReadTimeout:    100 * time.Second,
			WriteTimeout:   100 * time.Second,
			MaxHeaderBytes: 4 << 20,
		}
		server.SetKeepAlivesEnabled(false)
		rlog.V(1).Info( "即将启动服务:[" + PrivateConf.HttpHost + "]")
		server.ListenAndServe()
	}
}
func Get_E_LuckyMoney(w http.ResponseWriter, r *http.Request){
	defer UselessHelper.RecoverMedicine("Get_E_LuckyMoney")
	err := r.ParseForm()
	if err!=nil{
		w.Write([]byte("err"))
		rlog.V(1).Info("invalid url  err:"+err.Error())
		rlog.V(1).Info("_______________________________________________")
		return 
	}
	askBody, ioutilerr := ioutil.ReadAll(r.Body)
	if r.Body!=nil{
		defer r.Body.Close()
	}
	if ioutilerr != nil {
		rlog.V(1).Info("ioutilerr"+ioutilerr.Error())
		rlog.V(1).Info("_______________________________________________")
		return 
	} 
	var myrequestList string =bytes.NewBuffer(askBody).String()
	var myELuckyMoney ELuckyMoneyStruct
	err3 := json.Unmarshal([]byte(myrequestList), &myELuckyMoney)
	if err3 != nil {
		rlog.V(1).Info(err3)
		rlog.V(1).Info("_______________________________________________")
		return 
	}
	
	rlog.V(1).Info("myELuckyMoney.E_Url-----:",myELuckyMoney.E_Url)
	rlog.V(1).Info("myELuckyMoney.Phone-----:",myELuckyMoney.LuckyBoyPhone)
	myELuckyMoney.E_Url = strings.Replace(myELuckyMoney.E_Url,"#","?",1)
	myELuckyMoney.registerSuccess = false
	myELuckyMoney.luckyMoneySuccess = false
	myELuckyMoney.EGonaLucky = false
	myELuckyMoney.isUsedLucky = false
	if errE:=myELuckyMoney.Parse_E_LuckyMoney();errE!=nil{
		fmt.Fprint(w,"it is not a real E LuckyMoney.")
		return
	}
	if len(myELuckyMoney.LuckyBoyPhone)!=11{
		fmt.Fprint(w,"it is not a real phone number.")
		return
	}

	for rankFakeHungryMan,myFakeHungryMan:=range myFakeHungryMenArr{
		if myFakeHungryMan.openid=="lastone"{
			rlog.V(1).Info("摸金校尉人手不足")
			fmt.Fprint(w,"摸金校尉人手不足 抱歉")
			break
		}
		if myELuckyMoney.EGonaLucky {
			rlog.V(1).Info("pick up LuckyBoyPhone.")
			myFakeHungryMan.Phone= myELuckyMoney.LuckyBoyPhone
			if len(myELuckyMoney.HungryManName)>0{
				myFakeHungryMan.nickname= myELuckyMoney.HungryManName
			}else{
				myFakeHungryMan.nickname= "笑纳了"
			}
			myFakeHungryMan.register()
			time.Sleep(time.Second*1)
		}
		
		
		myluckydoorreturnjson ,luckydoorerr:=myFakeHungryMan.luckydoor(&myELuckyMoney)
		if luckydoorerr!=nil{
			continue
		}
		if myELuckyMoney.EGonaLucky {
			rlog.V(1).Info("take off LuckyBoyPhone.")
			myFakeHungryMan.Phone = getAFakeNum()
			myFakeHungryMan.nickname = getAFakeName()
			myFakeHungryMan.register()
			
		}
		var myluckydoorReturnJsonStruct luckydoorReturnJsonStruct
		err := json.Unmarshal([]byte(myluckydoorreturnjson), &myluckydoorReturnJsonStruct)
		if err != nil {
			rlog.V(1).Info("myluckydoorReturnJsonStruct json.Unmarshal:",err)
			rlog.V(1).Info("myluckydoorreturnjson:"+myluckydoorreturnjson)
			fmt.Fprint(w,"something err!!")
			break 
		}
		rlog.V(1).Info("lucky ranking-----------:",myELuckyMoney.lucky_number)
		rlog.V(1).Info("now ranking-------------:",len(myluckydoorReturnJsonStruct.Promotion_records))
		if myluckydoorReturnJsonStruct.Ret_code == fivemost{
			//rlog.V(1).Info("this lucky boy can not get E, he got 5!!!")
			//fmt.Fprint(w,"you already got 5 Elucky money today!!")
			rlog.V(1).Info("this lucky have been used 5 times!!!")
			continue
		}
		if myluckydoorReturnJsonStruct.Ret_code == alreadyCome{
			rlog.V(1).Info("this lucky boy Came here once!!!")
			//fmt.Fprint(w,"you already used this E lucky money!!")
			continue
		}
		if len(myluckydoorReturnJsonStruct.Promotion_records)==0{
			if strings.Contains(myluckydoorreturnjson,"TOO_BUSY"){
				rlog.V(1).Info("TOO_BUSY err:"+myluckydoorreturnjson)
				time.Sleep(time.Second*5)
				continue
			}else{
				rlog.V(1).Info("unknown err:"+myluckydoorreturnjson)
				continue
			}
		}
		if myELuckyMoney.lucky_number < len(myluckydoorReturnJsonStruct.Promotion_records){
			myELuckyMoney.isUsedLucky = true
			rlog.V(1).Info("Used lucky !!!")
			fmt.Fprint(w,"this is a used Elucky!")
			break
		}
		if myELuckyMoney.EGonaLucky{
			rlog.V(1).Info("FakeHungryMan ret.~~~~~~~~~~~~~~")


			if (myluckydoorReturnJsonStruct.Ret_code == successEdoor4||myluckydoorReturnJsonStruct.Ret_code == successEdoor3)&&myluckydoorReturnJsonStruct.Account == myELuckyMoney.LuckyBoyPhone{
				rlog.V(1).Info("lucky success!!!")
				fmt.Fprint(w,"lucky success!!")
				myELuckyMoney.isUsedLucky = true
				myELuckyMoney.EGonaLucky = false
				myELuckyMoney.luckyMoneySuccess = true
				break
			}else{
				rlog.V(1).Info("lucky err")
				rlog.V(1).Info("myluckydoorreturnjson:"+myluckydoorreturnjson)
				fmt.Fprint(w,"lucky err")
				fmt.Fprint(w,"myluckydoorreturnjson:"+myluckydoorreturnjson)
				break
			}
		}else{
			rlog.V(1).Info("FakeHungryMan ret.***************")
			if (myluckydoorReturnJsonStruct.Ret_code!=successEdoor4)&&(myluckydoorReturnJsonStruct.Ret_code!=successEdoor3){
				rlog.V(1).Info("Ret_code:",myluckydoorReturnJsonStruct.Ret_code)
				if myluckydoorReturnJsonStruct.Ret_code!=2{
					rlog.V(1).Info("myluckydoorReturnJsonStruct.Ret_code:",myluckydoorReturnJsonStruct.Ret_code)
				}
				continue
			}
		}
		if myELuckyMoney.lucky_number == len(myluckydoorReturnJsonStruct.Promotion_records)+1{
			myELuckyMoney.EGonaLucky = true
			rlog.V(1).Info("gona lucky !!!")
		}else if  myELuckyMoney.lucky_number <= len(myluckydoorReturnJsonStruct.Promotion_records){
			myELuckyMoney.isUsedLucky = true
			rlog.V(1).Info("Used lucky !!!")
			fmt.Fprint(w,"this is a used Elucky!")
			break
		}
		if rankFakeHungryMan == len(myFakeHungryMenArr)-1{
			rlog.V(1).Info("次数耗尽  明日再来")
			fmt.Fprint(w,"次数耗尽  明日再来")
			break
		}
	}
}



