package Other

import "strings"
/*
回溯结束条件：
当前ip字段总共有4个
若当前递归的s为空，则表示当前ip是合理ip可加入结果集
回溯条件：
ip地址每个字段范围为0-255，即三位
并且要根据题意给出的s因此遍历条件为i=[1,2,3]
如果当前s长度小于i 表示s中没有此次遍历所需的i个字符直接退出
如果当前组成的ip字段刚好是三个并且大于255则不合理直接退出
如果当前组成的ip字段长度大于1并且第一位是0则不合理直接退出
之后的ip字段是合理的直接添加进ip string数组不断递归下一个ip字段
回溯返回现场：
取消当前ip字段
*/
func backTrace(ip []string,s string,res *[]string){
	if len(ip)==4{
		if len(s)==0{
			*res = append(*res,ip[0]+"."+ip[1]+"."+ip[2]+"."+ip[3])
		}
		return
	}

	for i:=1;i<=3;i++{
		if len(s)<i{
			return
		}
		tmp:=s[:i]
		if len(tmp)==3&&strings.Compare(tmp,"255")>0{
			return
		}
		if len(tmp)>1&&s[0]=='0'{
			return
		}
		ip = append(ip,tmp)
		backTrace(ip,s[i:],res)
		ip = ip[:len(ip)-1]
	}

}

func restoreIpAddresses(s string) []string {
	res:=make([]string,0)
	if len(s)<4||len(s)>12{
		return res
	}
	backTrace(make([]string,0),s,&res)
	return res
}

