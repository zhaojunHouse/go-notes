package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getHourMinute(timeStr string) (string, error){
	logTime, err := time.ParseInLocation("2006-01-02 15:04:05",timeStr,time.Local)
	if err != nil {
		fmt.Println("generateFiveMinuteData-ParseInLocation-failed", err.Error())
		return "",err
	}
	minute := strconv.Itoa(logTime.Minute())
	minuteLastBit := minute[len(minute)-1:]
	if minuteLastBit < "5" { //  f_0000 = [0,5)  f_0005 = [5,10) f_0010 = [10,15) f_0015 = [15,20)
		minute = strings.Join([]string{minute[0:len(minute)-1],"0"},"")
	}
	if minuteLastBit >= "5" {
		minute = strings.Join([]string{minute[0:len(minute)-1],"5"},"")
	}
	min, err := strconv.Atoi(minute)
	if err != nil {
		return "", err
	}
	hourMinute := fmt.Sprintf("%02d%02d",logTime.Hour(),min)
	return hourMinute,nil
}


func main(){

}

func getFileSize(){
	urls := []string{
		"https://log-server.cn-bj.ufileos.com/8a152503c6c4c584d9f351c3938f7be2_182200046/202103111558_v29.douyinvod.com_CDN.log_part3.gz",
		"https://log-server.cn-bj.ufileos.com/7dab188bad9d6a8b4bcc8c8c83a8a90b_182534774/202103111557_v29.douyinvod.com_CDN.log_part12.gz",
		"https://log-server.cn-bj.ufileos.com/379cecd1ba579b7d24bac63fd277f667_183187465/202103111559_v29.douyinvod.com_CDN.log_part10.gz",
		"https://log-server.cn-bj.ufileos.com/099f6c4c72d917ddc98e560642ec9452_183703730/202103111555_v29.douyinvod.com_CDN.log_part28.gz",
		"https://log-server.cn-bj.ufileos.com/f6a6834a7afe1df130f3cef828aba4ba_183776446/202103111555_v29.douyinvod.com_CDN.log_part22.gz",
	}
	res,_ := batchGetFileSize(urls)
	fmt.Println(res)
}

type FileSizeResponse struct {
	Url  string `json:"url"`
	Size float64 `json:"size"`
	Err  error    `json:"err"`
}

func batchGetFileSize(urls []string) (map[string]float64,error){
	ch := make(chan *FileSizeResponse)

	for _,url := range urls{
		go func(filePath string) {
			res := GetFileSize(filePath)
			ch <- res
		}(url)
	}
	fileSizeMap := make(map[string]float64)
	var fileSizeErr error
	var i int
	for  v := range ch {
		fmt.Println(v)
		var size float64
		if v.Err == nil {
			size = v.Size
		}else{
			fileSizeErr = v.Err
			fmt.Println("getFileSize-filed", v.Err.Error())
		}
		fileSizeMap[getSplitLastItem(v.Url)] = size
		i++
		if len(urls) == i {
			close(ch)
		}
	}
	if fileSizeErr != nil {
		return nil, fileSizeErr
	}
	return fileSizeMap,nil
}

func getSplitLastItem(str string) string{
	urlSplits := strings.Split(str,"/")
	return urlSplits[len(urlSplits)-1]
}

func GetFileSize(url string) *FileSizeResponse{
	if url == "" {
		return &FileSizeResponse{
			Url:  url,
			Size: 0,
			Err:  nil,
		}
	}
	header, err := http.Head(url)
	if err != nil {
		fmt.Println("GetFileSize-Head-failed",err.Error())
		return &FileSizeResponse{
			Url:  url,
			Size: 0,
			Err:  err,
		}
	}
	var fileSize float64
	if header != nil{
		fileSize = math.Round(float64(header.ContentLength)/1024/1024)
	}
	return &FileSizeResponse{
		Url:  url,
		Size: fileSize,
		Err:  err,
	}
}


func IP6toIntV3(IPv6Address net.IP) *big.Int {
	IPv6Int := big.NewInt(0)
	IPv6Int.SetBytes(IPv6Address.To16())
	return IPv6Int
}

func ipv4IntToStr(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func ipv4StrToInt(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}


//func main() {
//	provIDStr:="1q"
//	provID,_ := strconv.Atoi(provIDStr)
//	fmt.Println(provID)
//
//	ipv6 := "2001:db8:0:1::101/45" // 0::/8     2001:250:21C::/48    2001:218:0:2000::180/123       2001:200:905::/49
//	start, end , _ := Ipv6CIDRToRange(ipv6)
//	fmt.Println(start)
//	fmt.Println(end)
//}
// 2001:0db8:0000:0001:0000:0000:0000:0101
func Ipv6CIDRToRange(ipv6 string) (string, string ,error) {
	idx := strings.Index(ipv6,"/")
	ip := ipv6[0:idx]
	prefix := ipv6[idx+1:]
	prefixInt,_ := strconv.Atoi(prefix)
	ipv6AllStr := getFormatIpv6(ip)
	// 十六进制字符串转二进制        二进制转十六进制
	startIPBin,err := hexTobin(ipv6AllStr)
	if err != nil {
		fmt.Println(err.Error())
		return "","", err
	}

	num1 := 128-prefixInt
	num1StrSlice := make([]string, 0, num1)
	for i:= 0 ;i< num1; i++{
		num1StrSlice = append(num1StrSlice, "1")
	}
	num1Str := strings.Join(num1StrSlice,"")
	endIPBin := strings.Join([]string{startIPBin[0:prefixInt] , num1Str},"")

	endIP16, err := binToHex(endIPBin)
	if err != nil {
		fmt.Println(err.Error())
		return "","", err
	}
	return ipv6AllStr,endIP16,nil
}
// 十六进制补齐
func getFormatIpv6(ipv6 string) string{
	ipv6Slice := strings.Split(ipv6,":")
	newIpv6Slice := make([]string,0,8)

	var i int
	for _,v := range ipv6Slice {
		switch len(v) {
		case 0:
			if i == 0 {
				for i:=0;i<8-len(ipv6Slice)+1;i++{
					newIpv6Slice = append(newIpv6Slice,"0000")
				}
				i++
			}else{
				newIpv6Slice = append(newIpv6Slice,"0000")
			}
			break
		case 4:
			newIpv6Slice = append(newIpv6Slice,v)
			break
		default:
			newIpv6Slice = append(newIpv6Slice,fmt.Sprintf("%04s",v) )
			break
		}
	}

	return strings.Join(newIpv6Slice,":")
}

// 十六进制转二进制     2001:021c:0000:2000:0000:0000:0000:0180
//10000000000001 1000011100 0000 10000000000000 000 110000000
func hexTobin(ip string) (string, error) {
	if len(ip) != 39 {
		return "", errors.New("invalid ip fro hexTobin")
	}
	ipSlice := strings.Split(ip,":")
	binStrSlice := make([]string,0,8)
	for _,ip16 := range ipSlice{
		ip64,err := strconv.ParseInt(ip16,16,64)
		if err != nil {
			return "", err
		}
		ipBin := fmt.Sprintf("%016s",strconv.FormatInt(ip64,2))
		binStrSlice = append(binStrSlice,ipBin)
	}
	return strings.Join(binStrSlice,""),nil
}

// 二进制转十六进制
func binToHex(ipBin string) (string, error) {
	if len(ipBin) != 128{
		return "", errors.New("invalid ip")
	}
	ip16Slice := make([]string,0,8)
	for i:= 0 ;i< 8;i++{
		start := 16*i
		end := 16*(i+1)
		one := ipBin[start:end]
		//fmt.Println("one---",one,"start--",start,"end--",end)
		ip64, err := strconv.ParseInt(one,2,64)
		if err != nil {
			return "",err
		}
		ipTo16 := fmt.Sprintf("%04s",strconv.FormatInt(ip64,16))
		ip16Slice = append(ip16Slice, ipTo16)
	}
	return strings.Join(ip16Slice,":"),nil
}








func IP6toInt(IPv6Address net.IP) *big.Int {
	IPv6Int := big.NewInt(0)
	ipv6AllStr := getFormatIpv6(IPv6Address.To16().String())
	fmt.Println("ToInt-----",IPv6Address.To16(),"ipString----",IPv6Address.To16().String(),"ipAllString---",ipv6AllStr)


	IPv6Int.SetBytes([]byte(ipv6AllStr))// 错误
	fmt.Println("ToBigInt-----",IPv6Int)
	return IPv6Int
}



//
//public static BigInteger ipv6toInt(String ipv6) {
//	int compressIndex = ipv6.indexOf("::");
//	if (compressIndex != -1) {
//	String part1s = ipv6.substring(0, compressIndex);
//	String part2s = ipv6.substring(compressIndex + 1);
//	BigInteger part1 = ipv6toInt(part1s);
//	BigInteger part2 = ipv6toInt(part2s);
//	int part1hasDot = 0;
//	char ch[] = part1s.toCharArray();
//	for (char c : ch) {
//	if (c == ':') {
//	part1hasDot++;
//	}
//	}
//	// ipv6 has most 7 dot
//	return part1.shiftLeft(16 * (7 - part1hasDot )).add(part2);
//	}
//
//	String[] str = ipv6.split(":");
//	BigInteger big = BigInteger.ZERO;
//	for (int i = 0; i < str.length; i++) {
//	//::1
//	if (str[i].isEmpty()) {
//	str[i] = "0";
//	}
//	big = big.add(BigInteger.valueOf(Long.valueOf(str[i], 16))
//	.shiftLeft(16 * (str.length - i - 1)));
//	}
//	return big;
//}


func ipv6Toint(ipv6 string ) int64{
	compressIndex := strings.Index(ipv6,"::")
	if compressIndex != -1 {
		part1Str := ipv6[0:compressIndex]
		part2Str := ipv6[compressIndex+1:]
		part1Int := ipv6Toint(part1Str)
		part2Int := ipv6Toint(part2Str)

		part1hasDot := 0;
		for i:=0;i<len(part1Str);i++ {
			if part1Str[i:i+1] == ":" {
				part1hasDot++
			}
		}
		return part1Int << (16*(7-part1hasDot))+part2Int
	}


	splitSlice := strings.Split(ipv6,":")
	var big int64
	for i:=0;i<len(splitSlice);i++{
		if splitSlice[i] == "" {
			splitSlice[i] = "0";
		}
		st := fmt.Sprintf("%x", splitSlice[i])
		stInt ,_ := strconv.Atoi(st)
		big += int64(stInt) << (16*len(splitSlice) -i -1)
	}

	return big
}


func HourMinuteHelper(){
	logTime, err := time.ParseInLocation("2006-01-02 15:04:05","2021-02-28 00:23:00",time.Local)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	minute := strconv.Itoa(logTime.Minute())
	minuteLastBit := minute[len(minute)-1:]
	if minuteLastBit < "5" { //  f_0000 = [0,5)  f_0005 = [5,10) f_0010 = [10,15) f_0015 = [15,20)
		minute = strings.Join([]string{minute[0:len(minute)-1],"0"},"")
	}
	if minuteLastBit >= "5" {
		minute = strings.Join([]string{minute[0:len(minute)-1],"5"},"")
	}
	min, err := strconv.Atoi(minute)
	if err != nil {
		return
	}
	hourMinute := fmt.Sprintf("%02d%02d",logTime.Hour(),min)
	fmt.Println(hourMinute)
}