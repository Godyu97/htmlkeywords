package main

import (
	"github.com/Godyu97/htmlkeywords/keywords"
	"log"
)

func main() {
	ext := keywords.NewExtractor(
		[]string{
			"项目编号",
			"招标项目编号",
			"采购项目编号",
			"采购项目编码",
			"项目标号",
			"采购单编号",
			"询价书编码",
			"询价单号",
			"标段/包编号",
			"标段（包）编号",
			"标的编号",
			"招标编号",
			"项目代码",
			"交易编号",
			"公告编号",
			"反拍单号",
			"工程编号",
			//"编号",
			"宗地编号",
			"标段标号",
			"比选编号",
			"招标编号及包号",
			"采购编号",
			"e采项目编号",
			"招募编号",
			"合同编号",
			"标段编号",
			"序列号",
			"场次号",
			"投资项目代码",
			"日期",
			"电话",
			"机构",
		},
		keywords.WithFilter(
			func(s string) string {
				s = keywords.DefaultFilter(s)
				return s
			}),
	)
	err := ext.ExtractKeywordsFromHtml(`<div><div><p><span></span><span>相关标段</span><span>收起</span></p><table><tr><td style=\"text-align: center;width: 15%;\">标段（包）名称</td><td style=\"text-align: center;width: 10%;\">标段（包）编号</td></tr></table></div><p><span></span><span>挂牌公告</span></p><div><ul><li>项目名称:</li><li>国家税务总局上海市金山区税务局部分资产（报废办公家具及信息设备一批）</li></ul><ul><li>项目编号:</li><li>GR2023SH1000885</li><li>转让价格(万元):</li><li>0.766</li></ul><ul><li>标的所属行业:</li><li>机械设备</li><li>标的所在地区:</li><li>上海嘉定区</li></ul><ul><li>信息披露起始日期:</li><li>2023-10-18</li><li>信息披露期满日期:</li><li>2023-10-24</li></ul><ul><li>受托机构:</li><li>受托机构名称:上海国际商品拍卖有限公司&nbsp&nbsp&nbsp&nbsp&nbsp受托机构联系人:李强&nbsp&nbsp&nbsp&nbsp&nbsp联系电话:13918874993</li></ul><ul><li>交易机构:</li><li>业务联系人:无&nbsp&nbsp&nbsp&nbsp&nbsp联系电话:无&nbsp&nbsp&nbsp&nbsp&nbsp负责人:高文骏&nbsp&nbsp&nbsp&nbsp&nbsp联系电话:62657272-154</li></ul></div><br/></div><div>原文链接：<a href=\"https://www.shggzy.com/jyxxcqgg/713520?cExt=eyJhbGciOiJIUzI1NiJ9.eyJwYXRoIjoiL2p5eHhjcSIsInBhZ2VObyI6MSwiZXhwIjoxNjk5MjYwNzAxMjI4fQ.J3kzqy1P5teq7Rbn3bJa3gg1blN-tu0ww2wKS6TCaVw&isIndex=\" target=\"_blank\">https://www.shggzy.com/jyxxcqgg/713520?cExt=eyJhbGciOiJIUzI1NiJ9.eyJwYXRoIjoiL2p5eHhjcSIsInBhZ2VObyI6MSwiZXhwIjoxNjk5MjYwNzAxMjI4fQ.J3kzqy1P5teq7Rbn3bJa3gg1blN-tu0ww2wKS6TCaVw&isIndex=</a></div>`)
	if err == nil {
		log.Println(ext.GetContent())
		log.Println(ext.GetResult(true))
	}

}
