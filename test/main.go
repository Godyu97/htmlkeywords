package main

import (
	"github.com/Godyu97/htmlkeywords/keywords"
	"log"
)

func main() {
	ext := keywords.NewExtractor(
		[]string{
			"电话",

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
			"机构",
		},
		keywords.WithFilter(
			func(s string) string {
				s = keywords.DefaultFilter(s)
				return s
			}),
	)
	err := ext.ExtractKeywordsFromHtml(`<div>                               <div>                            <table>                                <tr>                                    <th>招标公告名称：</th>                                    <td><font color='red'>秦皇岛市抚宁区2023年人造河综合整治工程勘察设计招标公告</font></td>                                </tr>                                <tr>                                    <th>招标公告编号：</th>                                    <td>I1301000075057306001G01</td>                                </tr>                                <tr>                                    <th>招标内容：</th>                                </tr>                                <tr>                                    <td><p>        </p><h2><font><font color='red'>秦皇岛市抚宁区2023年人造河综合整治工程勘察设计招标公告</font></font></h2>        <p>        <strong>            <span>1. 招标条件</span>        </strong>    </p>    <p>        <span>本招标项目<font>秦皇岛市抚宁区2023年人造河综合整治工程</font>已由<font>秦皇岛市抚宁区行政审批局</font>以<font>抚审批字【2023】40号</font>批准建设，项目业主为<font>秦皇岛市抚宁区水务局</font>，建设资金来自<font>省级以上资金及区财政</font>，出资比例为<font>100%</font>，招标人为 <font>秦皇岛市抚宁区水务局</font>。项目已具备招标条件，现对该项目的设计进行公开招标。</span>    </p>    <p>        <strong>            <span>2. 项目概况与招标范围</span>        </strong>    </p>    <p>        <span><font>2.1项目概况：2.1.1工程名称：秦皇岛市抚宁区2023年人造河综合整治工程勘察设计。2.1.2项目建设地点：秦皇岛市抚宁区留守镇。2.1.3建设规模：工程治理范围全长2911m，分别包含东河与西河，其中东河为留守营镇内友谊干渠至205国道，全长1996m；西河为友谊干渠至南街主街道，全长915m。现状输水涵洞进行清淤疏通、对破坏的混凝土涵洞盖板进行拆除重建，对现状河道进行清淤疏浚，并对局部段岸坡采取浆砌石挡墙防护，增加其稳定性。2.1.4勘察设计服务期限：合同签订后30日历天完成工程勘察、测量，初步设计编制、施工图设计。施工现场配合服务直至本项目交（竣）工验收后。2.1.5质量标准：满足国家行业及专业规范要求。</font><br><font>2.2招标范围：工程勘察、测量，初步设计编制、施工图设计；工程施工至竣工验收期间的设计服务等内容。</font>。</span>    </p>    <p>        <strong>            <span>3. 投标人资格要求</span>        </strong>    </p>    <p>        <span>3.1 本次招标对投标人的资格要求如下:<br><font> 3.1.1资质要求:（1）勘察：须具备工程勘察综合资质或工程勘察专业类（岩土工程勘察）乙级及以上资质；设计：须具备工程设计综合甲级资质或水利行业工程设计乙级及以上资质或具备水利行业（河道整治）专业乙级及以上资质。（2）财务要求：2022年度经会计师事务所或审计机构审计的财务会计报表，包括资产负债表、现金流量表、利润表和附注的扫描件。(投标人的成立时间少于投标人须知前附表规定年份的，应提供成立以来的财务会计报表扫描件)（3）信誉要求：在评标工作结束前，投标人未被“信用中国”网站 （www.creditchina.gov.cn）列入“失信被执行人”名单（以评标委员会现场查询为主）。（4）项目负责人的资格要求：本项目负责人须具有中级工程师专业技术职称（水利水电相关专业）及以上职称或执业资格证书。3.1.2特别提醒：本项目施行“双盲”评审和“分散”评标；投标文件技术标（设计方案）采用暗标方式编制及评审，即投标人在编制投标文件技术标部分时屏蔽投标人名称等信息，评标委员会依照招标文件的规定对投标文件技术标部分（设计方案）进行评审。    3.1.3接受异议的联系人和联系方式：招标代理机构：庞小东13933643968；招标人：马志君  0335-7811113。</font><br>。</span>    </p>    <p>        <span>3.2 本次招标<font>不接受</font>联合体投标。<font> </font></span>    </p>    <p>        <strong>            <span>4. 技术成果经济补偿</span>        </strong>    </p>    <p>        <span>本次招标对未中标人投标文件中的技术成果<font>不给予</font>经济补偿。给予经济补偿的，招标人将按如下标准支付经济补偿费用：<font>0</font></span>    </p>    <p>        <strong>            <span>5.招标文件的获取</span>        </strong>    </p>    <p>        <span>5.1 凡有意参加投标者，请于<font>2023-11-13 09:00</font>至<font>2023-11-18 17:30</font>（北京时间，下同），登录<font>惠招标电子招投标交易平台</font>下载电子招标文件。</span>    </p>    <p>        <span>5.2 招标文件售价<font>0</font>元，售后不退。技术资料押金<font>0</font>元，在退还技术资料时退还（不计利息）。 </span>    </p>    <p>        <strong>            <span>6. 投标文件的递交</span>        </strong>    </p>    <p>        <span>6.1 投标文件递交的截止时间为<font>2023-12-05 09:00:00</font>投标人应在截止时间前通过<font>惠招标</font>递交电子投标文件。</span>    </p>    <p>        <span>6.2 逾期送达的投标文件，电子招标投标交易平台将予以拒收。</span>    </p>    <p>        <strong>            <span>7. 发布公告的媒介</span>        </strong>    </p>    <p>        <span>本次招标公告同时在<font>河北省招标投标公共服务平台、惠招标电子招投标交易平台、秦皇岛市公共资源交易平台</font>上发布。</span>    </p>    <p>        <strong>            <span>8.联系方式</span>        </strong>    </p>    <p>    </p><table>        <tbody>            <tr>                <td>招标人：</td>                <td>                    <p>                        <a></a>                        <a></a>                        <a></a>                        <a></a>                        <a>                            <span><font>秦皇岛市抚宁区水务局</font></span>                        </a>                    </p>                </td>                <td>招标代理机构：</td>                <td>                    <span><font>秦皇岛广荣工程项目管理有限公司</font></span>                </td>            </tr>            <tr>                <td>地址：</td>                <td>                    <span><font>秦皇岛市抚宁区</font></span>                </td>                <td>地址：</td>                <td>                    <span><font>秦皇岛经济技术开发区泾河道6号二楼</font></span>                </td>            </tr>            <tr>                <td>邮编：</td>                <td>                    <p>                        <span><font>/</font></span>                    </p>                </td>                <td>邮编：</td>                <td>                    <span><font>/</font></span>                </td>            </tr>            <tr>                <td>联系人：</td>                <td>                    <p>                        <span><font>马志君</font></span>                    </p>                </td>                <td>联系人：</td>                <td>                    <span><font>庞小东</font></span>                </td>            </tr>            <tr>                <td>电话：</td>                <td>                    <p>                        <span><font>0335-7811113</font></span>                    </p>                </td>                <td>电话：</td>                <td>                    <span><font>13933643968</font></span>                </td>            </tr>            <tr>                <td>传真：</td>                <td>                    <p>                        <span><font>/</font></span>                    </p>                </td>                <td>传真：</td>                <td>                    <span><font>/</font></span>                </td>            </tr>            <tr>                <td>电子邮件：</td>                <td>                    <p>                        <span><font>/</font></span>                    </p>                </td>                <td>电子邮件：</td>                <td>                    <span><font>251507331@qq.com</font></span>                </td>            </tr>            <tr>                <td>网址：</td>                <td>                    <p>                        <span><font>/</font></span>                    </p>                </td>                <td>网址：</td>                <td>                    <span><font>/</font></span>                </td>            </tr>            <tr>                <td>开户银行：</td>                <td>                    <p>                        <span><font>/</font></span>                    </p>                </td>                <td>开户银行：</td>                <td>                    <span><font>/</font></span>                </td>            </tr>            <tr>                <td>账号：</td>                <td>                    <p>                        <span><font>/</font></span>                    </p>                </td>                <td>账号：</td>                <td>                    <span><font>/</font></span>                </td>            </tr>        </tbody>    </table>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 </td>                                </tr>                                </table>                                                        <div>                                                            </div>                        </div>                    <br>                </div>                `)
	if err == nil {
		log.Println(ext.GetContent())
		log.Println(ext.GetResult(true))
		log.Println(ext.GetItemsByWeight())
	}

}
