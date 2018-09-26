package spider

var conf = `

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head><meta http-equiv="Content-Type" content="text/html; charset=utf-8" /><title>

</title><link href="order/style.css" rel="stylesheet" /><link href="order/charge.css" rel="stylesheet" />
     <script src="js/jquery-1.9.1.min.js"></script>
        <script src="layer/layer.js"></script>
</head>

<body style="background-color:#fff">
    <input type="hidden" id="fbetsprofit" value="0.08"/>
    <input type="hidden" id="gcid" value="3"/>
    <input type="hidden" id="gpid" value="11"/>
     <input type="hidden" id="fieldnum" value="180919-20"/>

       <input type="hidden" id="gcname" value="上海11选5"/>
       <input type="hidden" id="gpname" value="5连号"/>

     <input type="hidden" id="timenow" value="2018-09-19 11:40:00"/>
     <input type="hidden" id="kjtime" value="2018-09-19 12:09:00"/>

     <input type="hidden" id="ordernumber" value="49041537328400812"/>

     <table style="background-color:#CD0000; color:#fff; width:100%;text-align:center; border:none;" rules=none>

            <tr style="font-size: 20px;background-color:#CD0000; line-height:50px; border:none;box-shadow:0 0 1px #000 inset;">
                <td width="10%" class="text-center"><strong>下单频道</strong></td>
                <td width="10%" class="text-center"><strong>下单形态</strong></td>
                 <td width="10%" class="text-center"><strong>下单期号</strong></td>
                         <td width="10%" class="text-center"><strong>剩余交易量</strong></td>
                 <td width="10%" class="text-center"><strong>反波收益</strong></td>
             
            </tr>

             <tr style="font-size: 18px; line-height:50px; background-color:#000; box-shadow:0px 0px 20px #000; border:none">
                    <td class="text-center" style=" border:none">上海11选5</td>
                    <td class="text-center">5连号</td>
                     <td class="text-center">180919-20</td>
                 
                       <td class="text-center">947900.00</td>

                 
                    <td class="text-center">0.80‰</td>
                   
                </tr>
          <tr style="font-size: 18px; line-height:2px; background-color:#CD0000; box-shadow:0px 0px 20px #000; border:none">
                    <td class="text-center" style=" border:none" colspan="7">&nbsp;</td>
                </tr>
     
    </table>

   



    <div class="yb_bg">
	    <div class="wrap">
		    <div class="border_yb">
		        <h2 class="data_tit">反波下单 <span style="font-size:13px;">(今日保本体验剩余 <span style="font-size:16px; color: red">14</span> 次)</span> <span style="float:right; font-size:16px;">下单剩余:&nbsp;<span class="buy_time" id="end_date_0"></span></span></h2>
		        <ul class="charge">
                    <li><span class="ch_t"><b>*</b>反波单号：</span><span class="ch_t">49041537328400812</span></li>
			        <li><span class="ch_t"><b>*</b>账户余额：</span><span class="ch_t" style="color:red">￥87.64</span></li>
			        <li><span class="ch_t"><b>*</b>下单金额：</span><input id="price" type="text" name="price" onchange="jsprice(this)" class="ch_input" onkeyup="value=value.replace(/[^\d]/g,'')">&nbsp;&nbsp;&nbsp;&nbsp;<span style="color:red">反波收益:&nbsp;&nbsp;￥<strong style="font-size:18px;" id="fbsy">0.00</strong> </span></li>
			        <li>
				        <span class="ch_t"><b></b>&nbsp;</span>
				        <ul class="charge_m">
				        	
                            <li val="z_1" name="liprice" id="liprice1" onclick="butshow('liprice1',500)">￥500</li>
				        	<li val="z_10" name="liprice" id="liprice2" onclick="butshow('liprice2',1000)">￥1000</li>
				        	<li val="z_20" name="liprice" id="liprice3" onclick="butshow('liprice3',2000)">￥2000</li>
				        	<li val="z_50" name="liprice" id="liprice4" onclick="butshow('liprice4',5000)">￥5000</li>
				        	<li val="z_100" name="liprice" id="liprice5" onclick="butshow('liprice5',10000)">￥10000</li>
				        	<li val="z_200" name="liprice" id="liprice6" onclick="butshow('liprice6',20000)">￥20000</li>
				        	<li val="z_300" name="liprice" id="liprice7" onclick="butshow('liprice7',30000)">￥30000</li>
				        	<li val="z_500" name="liprice" id="liprice8" onclick="butshow('liprice8',50000)">￥50000</li>
				        	<li val="z_800" name="liprice" id="liprice9" onclick="butshow('liprice9',100000)">￥100000</li>
				        </ul>
			        </li>
				
					<li><button type="submit" name="but0" id="but0" onclick="sub()" class="charge_btn" value="立即支付">确认下单</button></li>
				    <li><p style="margin-top:20px; margin-left:110px">下单说明：下单金额必须是100的倍数.</p></li>
		        </ul>
		    </div>
		</div>
	</div>

    

    <script>
    
        var timenow = document.getElementById("timenow").value;
      
        var arr = timenow.split(/[- : \/]/)
        var date = new Date(arr[0], arr[1] - 1, arr[2], arr[3], arr[4], arr[5]);
        var todayTime0 = date.getTime();

        var kjtime = document.getElementById("kjtime").value;
       

        var arr2 = kjtime.split(/[- : \/]/)
        var date2 = new Date(arr2[0], arr2[1] - 1, arr2[2], arr2[3], arr2[4], arr2[5]);
   
        show_date_time(date2.getTime(), 2, 0);

        function show_date_time(end, style, id) {
           
            //todayTime = new Date().getTime();
            var vars_name = 'todayTime' + id;
            var timea = eval(vars_name) + 100;

            eval("todayTime" + id + "=" + timea);

            timeold = ((end) - eval(vars_name));
            if (timeold <= 0) {
                //已经到开奖时间了 退出倒计时 开始开奖
                $("#end_date_" + id).html("已结束");
                vars_name = "but" + id;
                document.getElementById(vars_name).innerHTML = "已结束";
                document.getElementById(vars_name).disabled = true;
                return;
            }
            setTimeout("show_date_time(" + end + ',' + style + ',' + id + ")", 100);
            sectimeold = timeold / 1000;
            secondsold = Math.floor(sectimeold);
            msPerDay = 24 * 60 * 60 * 1000;
            e_daysold = timeold / msPerDay;
            daysold = Math.floor(e_daysold);
            e_hrsold = (e_daysold - daysold) * 24;
            hrsold = Math.floor(e_hrsold);
            e_minsold = (e_hrsold - hrsold) * 60;
            minsold = Math.floor((e_hrsold - hrsold) * 60);
            e_seconds = (e_minsold - minsold) * 60;
            seconds = Math.floor((e_minsold - minsold) * 60);
            ms = e_seconds - seconds;
            ms = new String(ms);
            ms1 = ms.substr(2, 1);
            ms2 = ms.substr(2, 2);
            hrsold1 = daysold * 24 + hrsold;
            if (style == 1) {
                $("#end_date_" + id).html((hrsold1 < 10 ? '0' + hrsold1 : hrsold1) + "小时" + (minsold < 10 ? '0' + minsold : minsold) + "分" + (seconds < 10 ? '0' + seconds : seconds) + "秒");
            } else if (style == 2) {
                $("#end_date_" + id).html("<em id='h'>" + (hrsold < 10 ? '0' + hrsold : hrsold) + "</em>:<em id='m'>" + (minsold < 10 ? '0' + minsold : minsold) + "</em>:<em id='s'>" + (seconds < 10 ? '0' + seconds : seconds) + "</em>");
            } else if (style == 3) {
                $("#end_date_" + id).html((hrsold1 < 10 ? '0' + hrsold1 : hrsold1) + "小时" + (minsold < 10 ? '0' + minsold : minsold) + "分" + (seconds < 10 ? '0' + seconds : seconds) + "." + ms1 + "秒");
            } else {
                $("#end_date_" + id).html((hrsold < 10 ? '0' + hrsold : hrsold) + "小时" + (minsold < 10 ? '0' + minsold : minsold) + "分" + (seconds < 10 ? '0' + seconds : seconds) + "秒." + ms2);
            }
        }

       
        function settime() {
            $.ajax({
                type: "POST",
                async: true,  // 设置ture 异步
                cache: false,
                url: "user/seltime",
                data: { "": 1 },
                dataType: "json",
                success: function (result) {
                    if (result != "") {

                       var arr = result.split(/[- : \/]/)
                       var date = new Date(arr[0], arr[1] - 1, arr[2], arr[3], arr[4], arr[5]);

                       todayTime0 = result.getTime();
                        setTimeout(settime, 30000);//更新校正时间
                    }

                }
            });
        }
        settime();

        function butshow(id,price) {

            document.getElementById("liprice1").className = "";
            document.getElementById("liprice2").className = "";
            document.getElementById("liprice3").className = "";
            document.getElementById("liprice4").className = "";
            document.getElementById("liprice5").className = "";
            document.getElementById("liprice6").className = "";
            document.getElementById("liprice7").className = "";
            document.getElementById("liprice8").className = "";
            document.getElementById("liprice9").className = "";

            document.getElementById(id).className = "on";

            document.getElementById("price").value = price;


            var fbetsprofit = document.getElementById("fbetsprofit").value;
          
            var fbsy = price * (fbetsprofit / 100);
            document.getElementById("fbsy").innerHTML = toDecimal2(fbsy);
        }

        function jsprice(docum) {

            var price= docum.value;
            var fbetsprofit = document.getElementById("fbetsprofit").value;

            var fbsy = price * (fbetsprofit / 100);
            document.getElementById("fbsy").innerHTML = toDecimal2(fbsy);
        }


        function toDecimal2(x) {
            var f = parseFloat(x);
            if (isNaN(f)) {
                return false;
            }
            var f = Math.round(x * 100) / 100;
            var s = f.toString();
            var rs = s.indexOf('.');
            if (rs < 0) {
                rs = s.length;
                s += '.';
            }
            while (s.length <= rs + 2) {
                s += '0';
            }
            return s;
        }

        function sub() {
            document.getElementById("but0").innerHTML = "正在提交…";
            document.getElementById("but0").disabled = true;
            var price = document.getElementById("price").value;
            var r = /^[1-9]\d*00$/;
            if (!r.test(price)) {
                //layer.msg('输入的下单金额不正确,下单金额为100的倍数.', { icon:5 });
                layer.alert("下单金额不正确,下单金额必须是100的倍数.");
                document.getElementById("but0").innerHTML = "确认下单";
                document.getElementById("but0").disabled = false;
                return false;
            }
            var fieldnum = document.getElementById("fieldnum").value;
            var gcname = document.getElementById("gcname").value;
            var gpname = document.getElementById("gpname").value;
   
            layer.confirm("是否确定下单：<Br>下单频道————" + gcname + "<Br>下单形态————<span style=\"color:red\">" + gpname + "</span><Br>下单期号————" + fieldnum + "<Br>下单金额————" + price +"元<Br>确定后无法撤销是否确定？", {
                btn: ['确定', '取消'] //按钮
                , cancel: function (index) {
                    //取消操作，点击右上角的X  
                    document.getElementById("but0").innerHTML = "确认下单";
                    document.getElementById("but0").disabled = false;
                }  
            }, function () {
                var gcid = document.getElementById("gcid").value;
                var gpid = document.getElementById("gpid").value;
                var ordernumber = document.getElementById("ordernumber").value;
                $.ajax({
                    type: "POST",
                    async: true,  // 设置ture 异步
                    cache: false,
                    url: "user/buyorder",
                    data: { "ordernum": ordernumber,"gcid": gcid, "gpid": gpid, "fieldnum": fieldnum, "buyprice": price },
                    dataType: "json",
                    success: function (result) {

                        if (result != "") {
                            if (eval(result).message == "ok") {
                                layer.alert('恭喜你,已下单成功.', function () {
                                    location.reload();
                                });

                            } 
                        else {
                        if(eval(result).message.indexOf('交易量不足')>0){
                             layer.alert("下单失败,交易量不足.");
                             document.getElementById("but0").disabled = false;
                               document.getElementById("but0").innerHTML = "确认下单";
                        }else{
                             layer.alert(eval(result).message);
                                                        document.getElementById("but0").disabled = false;
                                                        document.getElementById("but0").innerHTML = "确认下单";

                        }
                           
                            }

                        } else {
                            location.href = "login";
                        }

                    }
                });
                }, function (index) {
                    document.getElementById("but0").innerHTML = "确认下单";
                    document.getElementById("but0").disabled = false;
                    layer.close(index);
                
            });
          
          


        }

    </script>
</body>
</html>
`
