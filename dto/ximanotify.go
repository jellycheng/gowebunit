package dto

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"gowebunit/utils"
)

type XimaNotifyDto struct {
	ResultCode int    `json:"result_code"`
	ResultMsg  string `json:"result_msg"`
	CpOrderID  string `json:"cp_order_id"`
	OutOrderID string `json:"out_order_id"`
	Fee        int    `json:"fee"`
	PayType    int    `json:"pay_type"`
	OutExt     string `json:"out_ext"`
	Sign       string `json:"sign"`
	OrderType  int    `json:"order_type"`
	ProductId  string `json:"product_id"`
	OpType     int    `json:"op_type"`
	FirstPay   int    `json:"first_pay"`
}

func (m XimaNotifyDto) CheckSign(appsecret string) bool {
	if m.PayType == 2 { //2-点播业务
		//MD5（result_code+cp_order_id+out_order_id+fee+pay_type+app_secret）
		str := fmt.Sprintf("%d%s%s%d%d%s", m.ResultCode, m.CpOrderID, m.OutOrderID, m.Fee, m.PayType, appsecret)
		utils.Writelog("点播拼接串：" + str)
		myMd5 := gosupport.Md5(str)
		if myMd5 == m.Sign {
			return true
		}
	} else if m.PayType == 1 { // 1- 包月业务
		// MD5（result_code+cp_order_id+out_order_id+fee+pay_type+order_type+product_id+op_type+app_secret）
		str := fmt.Sprintf("%d%s%s%d%d%d%s%d%s",
			m.ResultCode, m.CpOrderID, m.OutOrderID, m.Fee, m.PayType, m.OrderType, m.ProductId, m.OpType, appsecret)
		utils.Writelog("包月拼接串：" + str)
		myMd5 := gosupport.Md5(str)
		if myMd5 == m.Sign {
			return true
		}
	}
	return false
}
