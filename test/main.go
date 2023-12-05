package main

import (
	"github.com/Godyu97/htmlkeywords/keywords"
	"log"
)

var ProjectCodeKeys = []string{
	//"项目编号",
	//"招标项目编号",
	//"采购项目编号",
	//"采购项目编码",
	//"项目标号",
	//"采购单编号",
	//"询价书编码",
	//"询价单号",
	//"询价单编号",
	//"标段/包编号",
	//"标段（包）编号",
	//"标的编号",
	//"招标编号",
	//"项目代码",
	//"交易编号",
	//"公告编号",
	//"反拍单号",
	//"工程编号",
	//"宗地编号",
	//"标段标号",
	//"比选编号",
	//"招标编号及包号",
	//"采购编号",
	//"e采项目编号",
	//"招募编号",
	//"合同编号",
	//"标段编号",
	//"序列号",
	//"场次号",
	//"投资项目代码",
	//"编号",
	"中标价(元)",
	//"中标价（元）",
}

func main() {
	ext := keywords.NewExtractor(
		ProjectCodeKeys,
		keywords.WithFilter(
			func(s string) string {
				s = keywords.DefaultFilter(s)
				return s
			}),
		keywords.WithFilter(keywords.DefaultFilter),
	)

	err := ext.ExtractKeywordsFromHtml(`<table> 
 <tbody>
  <tr> 
   <td >
    <table> 
     <tbody>
      <tr> 
       <td height="35" colspan="2" > </td> 
      </tr> 
      <tr> 
       <td width="28%" height="50" > 工程编号</td> 
       <td>  T0WZ202300505-P02</td> 
      </tr> 
      <tr> 
       <td height="50" > 建设单位名称</td> 
       <td>   </td> 
      </tr> 
      <tr> 
       <td height="50" > 工程名称</td> 
       <td>  新建西安至延安铁路西安至铜川段西安枢纽联络线工程国铁集团管理甲供物资第五批次采购  </td> 
      </tr> 
      <tr> 
       <td height="50" > 建设地点</td> 
       <td>   </td> 
      </tr> 
      <tr> 
       <td height="50" > 中标人</td> 
       <td>  
        <div>
         <div>
          <b><u>  新建西安至延安铁路西安至铜川段西安枢纽联络线工程国铁集团管理甲供物资第五批次采购   </u>中标结果公告</b>
         </div>
         <div>
          <b>招标编号：  <u>  T0WZ202300505</u></b>
         </div>
         <div>
              
          <u>  新建西安至延安铁路西安至铜川段西安枢纽联络线工程国铁集团管理甲供物资第五批次采购   </u>的招标中标候选人公示已结束。现就本次招标的中标结果公示如下：
         </div>
         <div>
          <b>一、中标人</b>
         </div>
         <div>
          包件编号：
          <u>P02</u>(包件名称：
          <u>信号电缆</u>)
          <div>
               中标人名称：
           <u>安徽铁信光电科技有限公司 </u>
          </div>
         </div>
         <div>
              中标价（元）：
          <u>14408342.2 </u>
         </div>
        </div>
        <div>
         <b>二、联系方式</b>
        </div>
        <div>
         <table>
          <tbody>
           <tr>
            <td>    招 标人： 中国铁路西安局集团有限公司</td>
           </tr>
           <tr>
            <td>    地址： 陕西省西安市碑林区南二环东段366号</td>
           </tr>
           <tr>
            <td>    联 系 人： 冷济福</td>
           </tr>
           <tr>
            <td>    联系电话： 13363973455</td>
           </tr>
           <tr>
            <td>    传真/邮箱： 029-82159987</td>
           </tr>
           <tr>
            <td>    招 标代理： 中国铁路物资股份有限公司</td>
           </tr>
           <tr>
            <td>    地址： 北京市丰台区太平街道凤凰嘴街5号院2号楼鼎兴大厦A座10层</td>
           </tr>
           <tr>
            <td>    联 系 人： 杜凯娟（售卖招标文件）、商学斌、赵艺程</td>
           </tr>
           <tr>
            <td>    联系电话： 010-51898508（售卖招标文件电话）、010-51895074、010-51898717</td>
           </tr>
           <tr>
            <td>    传真/邮箱： 010-51898566</td>
           </tr>
          </tbody>
         </table>
        </div>
        <table>
         <tbody>
          <tr>
           <td width="40%"> </td>
           <td width="50%" >日期：2023年08月07日 </td>
           <td width="10%"></td>
          </tr>
         </tbody>
        </table>  </td> 
      </tr> 
      <tr> 
       <td height="50" > 中标价(元)</td> 
       <td>  14408342.2</td> 
      </tr> 
      <tr> 
       <td height="50" > 公示开始时间</td> 
       <td>  2023-8-7</td> 
      </tr> 
     </tbody>
    </table></td> 
  </tr> 
 </tbody>
</table>`)
	if err == nil {
		log.Println(ext.GetContent())
		log.Println(ext.GetResult(true))
	}

}
