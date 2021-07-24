package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	mrand "math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	//PathSeparator
	PathSeparator = string(os.PathSeparator)
	DevNull       = os.DevNull
)

func Get_Bin_Path() (string, error) {
	file, _ := exec.LookPath(os.Args[0])
	bin_path, _ := filepath.Abs(file)
	bin_path, err := filepath.EvalSymlinks(bin_path)
	if nil != err {
		return "", err
	}
	return bin_path, nil
}

func Get_Bin_Dir() string {
	binPath, _ := Get_Bin_Path()
	return filepath.Dir(binPath)
}

func Get_Project_Dir() string {
	dirs := strings.Split(Get_Bin_Dir(), PathSeparator)
	project_path := strings.Join(dirs[:len(dirs)-1], PathSeparator)
	return project_path
}

func DateToStr(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func Format_Date(time time.Time, format string) string {
	return time.Format(format)
}

func MkdirByFile(file string) error {
	fileDir := filepath.Dir(file)
	if !IsDir(fileDir) {
		if err := os.Mkdir(fileDir, os.ModePerm); nil != err {
			return err
		}
	}
	return nil
}

func WritePidFile(file, pid string) error {
	if err := MkdirByFile(file); nil != err {
		return err
	}

	pidfile, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer pidfile.Close()
	_, err = io.WriteString(pidfile, pid)
	if err != nil {
		return err
	}
	return nil
}

type StringUtils string

func (s *StringUtils) Set(v string) {
	if v != "" {
		*s = StringUtils(v)
	} else {
		s.Clear()
	}
}

func (s *StringUtils) Clear() {
	*s = StringUtils(0x1E)
}

func (s StringUtils) Exist() bool {
	return string(s) != string(0x1E)
}

func (s StringUtils) Bool() (bool, error) {
	v, err := strconv.ParseBool(s.String())
	return bool(v), err
}

func (s StringUtils) Float32() (float32, error) {
	v, err := strconv.ParseFloat(s.String(), 32)
	return float32(v), err
}

func (s StringUtils) Float64() (float64, error) {
	return strconv.ParseFloat(s.String(), 64)
}

func (s StringUtils) Int() (int, error) {
	v, err := strconv.ParseInt(s.String(), 10, 32)
	return int(v), err
}

func (s StringUtils) Int8() (int8, error) {
	v, err := strconv.ParseInt(s.String(), 10, 8)
	return int8(v), err
}

func (s StringUtils) Int16() (int16, error) {
	v, err := strconv.ParseInt(s.String(), 10, 16)
	return int16(v), err
}

func (s StringUtils) Int32() (int32, error) {
	v, err := strconv.ParseInt(s.String(), 10, 32)
	return int32(v), err
}

func (s StringUtils) Int64() (int64, error) {
	v, err := strconv.ParseInt(s.String(), 10, 64)
	return int64(v), err
}

func (s StringUtils) Uint() (uint, error) {
	v, err := strconv.ParseUint(s.String(), 10, 32)
	return uint(v), err
}

func (s StringUtils) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(s.String(), 10, 8)
	return uint8(v), err
}

func (s StringUtils) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(s.String(), 10, 16)
	return uint16(v), err
}

func (s StringUtils) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(s.String(), 10, 32)
	return uint32(v), err
}

func (s StringUtils) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(s.String(), 10, 64)
	return uint64(v), err
}

func (s StringUtils) ToTitleLower() string {
	str := strings.ToLower(s.String()[:1]) + s.String()[1:]
	return str
}

func (s StringUtils) ToTitleUpper() string {
	str := strings.ToUpper(s.String()[:1]) + s.String()[1:]
	return str
}

func (s StringUtils) RegexpSQLVal() (bool, error) {
	r := "^[0-9a-zA-Z\\s-_\u4E00-\u9FA5'=@.?]+$"
	b, err := regexp.MatchString(r, s.String())
	return bool(b), err
}

func (s StringUtils) ContainsNum() (bool, error) {
	r := "^.*\\d+.*+$"
	b, err := regexp.MatchString(r, s.String())
	return bool(b), err
}

func (s StringUtils) ContainsBool(sep string) bool {
	index := strings.Index(s.String(), sep)
	return index > -1
}

func (s StringUtils) RegexpSQLSgin() (bool, error) {
	r := "^[<>=!?]+$"
	b, err := regexp.MatchString(r, s.String())
	return bool(b), err
}

func (s StringUtils) String() string {
	if s.Exist() {
		return string(s)
	}
	return ""
}

func (s StringUtils) MD5() string {
	m := md5.New()
	m.Write([]byte(s.String()))
	return hex.EncodeToString(m.Sum(nil))
}

func (s StringUtils) SHA1() string {
	sha := sha1.New()
	sha.Write([]byte(s.String()))
	return hex.EncodeToString(sha.Sum(nil))
}

func (s StringUtils) SHA256() string {
	sha := sha256.New()
	sha.Write([]byte(s.String()))
	return hex.EncodeToString(sha.Sum(nil))
}

func (s StringUtils) SHA512() string {
	sha := sha512.New()
	sha.Write([]byte(s.String()))
	return hex.EncodeToString(sha.Sum(nil))
}

func (s StringUtils) HMAC_SHA1(key string) string {
	mc := hmac.New(sha1.New, []byte(key))
	mc.Write([]byte(s.String()))
	return hex.EncodeToString(mc.Sum(nil))
}

func (s StringUtils) HMAC_SHA256(key string) string {
	mc := hmac.New(sha256.New, []byte(key))
	mc.Write([]byte(s.String()))
	return hex.EncodeToString(mc.Sum(nil))
}

func (s StringUtils) HMAC_SHA512(key string) string {
	mc := hmac.New(sha512.New, []byte(key))
	mc.Write([]byte(s.String()))
	return hex.EncodeToString(mc.Sum(nil))
}

func (s StringUtils) Base64Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(s.String()))
}

func (s StringUtils) Base64Decode() (string, error) {
	v, err := base64.StdEncoding.DecodeString(s.String())
	return string(v), err
}

func RandInt64(start, end int64) int64 {
	if start >= end {
		return end
	}
	rnd := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	return rnd.Int63n(end-start) + start
}

func GenerateMacAddr() string {
	mac := []int64{
		0x00, 0xf8, 0x3e, RandInt64(0x00, 0xff), RandInt64(0x00, 0xff), RandInt64(0x00, 0xff),
	}
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

func AESEncode(msg, key string) (string, error) {
	if len(key) == 16 {
		var iv = []byte(key)[:aes.BlockSize]
		c := make([]byte, len(msg))
		be, err := aes.NewCipher([]byte(key))
		if err != nil {
			return "", err
		}
		e := cipher.NewCFBEncrypter(be, iv)
		e.XORKeyStream(c, []byte(msg))
		b64 := base64.StdEncoding.EncodeToString(c)
		b64 = strings.Replace(b64, "/", "-", -1)
		return b64, nil
	} else {
		return "", fmt.Errorf("%s", "Key length is not equal to 16.")
	}
}

func AESDecode(enmsg, key string) (string, error) {
	if len(key) == 16 {
		enmsg = strings.Replace(enmsg, "-", "/", -1)
		msg, err := base64.StdEncoding.DecodeString(enmsg)
		if nil != err {
			return "", err
		}
		var iv = []byte(key)[:aes.BlockSize]
		d := make([]byte, len(msg))
		var bd cipher.Block
		bd, err = aes.NewCipher([]byte(key))
		if err != nil {
			return "", err
		}
		e := cipher.NewCFBDecrypter(bd, iv)
		e.XORKeyStream(d, msg)
		return string(d), nil
	} else {
		return "", fmt.Errorf("%s", "Key length is not equal to 16.")
	}
}

func (s StringUtils) GenerateRandStr32() string {
	b := make([]byte, 24)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func ToStr(value interface{}) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', 2, 32)
	case float64:
		s = strconv.FormatFloat(v, 'f', 2, 64)
	case int:
		s = strconv.FormatInt(int64(v), 10)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(int64(v), 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

func ConvertUTF8(src []byte) ([]byte, error) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GetCharset("UTF-8").NewEncoder()))
	return data, err
}

func GetCharset(charset string) encoding.Encoding {
	switch strings.ToUpper(charset) {
	case "GB18030":
		return simplifiedchinese.GB18030
	case "GB2312", "HZ-GB2312":
		return simplifiedchinese.HZGB2312
	case "GBK":
		return simplifiedchinese.GBK
	case "BIG5":
		return traditionalchinese.Big5
	case "EUC-JP":
		return japanese.EUCJP
	case "ISO2022JP":
		return japanese.ISO2022JP
	case "SHIFTJIS":
		return japanese.ShiftJIS
	case "EUC-KR":
		return korean.EUCKR
	case "UTF8", "UTF-8":
		return encoding.Nop
	case "UTF16-BOM", "UTF-16-BOM":
		return unicode.UTF16(unicode.BigEndian, unicode.UseBOM)
	case "UTF16-BE-BOM", "UTF-16-BE-BOM":
		return unicode.UTF16(unicode.BigEndian, unicode.UseBOM)
	case "UTF16-LE-BOM", "UTF-16-LE-BOM":
		return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM)
	case "UTF16", "UTF-16":
		return unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	case "UTF16-BE", "UTF-16-BE":
		return unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	case "UTF16-LE", "UTF-16-LE":
		return unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	//case "UTF32", "UTF-32":
	//	return simplifiedchinese.GBK
	default:
		return nil
	}
}

func IsFile(p string) bool {
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func IsDir(p string) bool {
	fi, err := os.Stat(p)
	if err != nil {
		return false
	} else {
		return fi.IsDir()
	}
	return false
}

func MkDirAll(dir string) error {
	if IsDir(dir) {
		if err := os.RemoveAll(dir); nil != err {
			return err
		}
	}
	if err := os.MkdirAll(dir, os.ModePerm); nil != err {
		return err
	}
	return nil
}

func Get_File_Size(fileSize int64) string {
	ff_size := float64(fileSize)
	var (
		fs   string
		pb_s float64 = 1024 << 40
		tb_s float64 = 1024 << 30
		gb_s float64 = 1024 << 20
		mb_s float64 = 1024 << 10
	)
	if ff_size > pb_s {
		f := ff_size / pb_s
		fs = ToStr(f) + " PB"
	} else if ff_size > tb_s {
		f := ff_size / tb_s
		fs = ToStr(f) + " TB"
	} else if ff_size > gb_s {
		f := ff_size / gb_s
		fs = ToStr(f) + " GB"
	} else if ff_size > mb_s {
		f := ff_size / mb_s
		fs = ToStr(f) + " MB"
	} else {
		f := ff_size / 1024
		fs = ToStr(f) + " KB"
	}
	return fs
}

func Daemon(nochdir, noclose int) (int, error) {
	if syscall.Getppid() == 1 {
		syscall.Umask(0)
		if nochdir == 0 {
			binPath, _ := Get_Bin_Path()
			os.Chdir(binPath)
		}
		return 0, nil
	}
	files := make([]*os.File, 3, 6)
	if noclose == 0 {
		devNull, err := os.OpenFile(DevNull, os.O_RDWR, 0)
		if err != nil {
			return 1, err
		}
		files[0], files[1], files[2] = devNull, devNull, devNull
	} else {
		files[0], files[1], files[2] = os.Stdin, os.Stdout, os.Stderr
	}
	sysattrs := syscall.SysProcAttr{Setsid: true}
	attrs := os.ProcAttr{Dir: Get_Bin_Dir(), Env: os.Environ(), Files: files, Sys: &sysattrs}
	binPath, _ := Get_Bin_Path()
	proc, err := os.StartProcess(binPath, os.Args, &attrs)
	if err != nil {
		return -1, err
	}
	proc.Release()
	os.Exit(0)
	return 0, nil
}
