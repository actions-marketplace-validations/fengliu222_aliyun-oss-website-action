package config

import (
	"fmt"
	"os"

	"aliyun-oss-website-action/utils"

	"github.com/fangbinwei/aliyun-oss-go-sdk/oss"
	"github.com/joho/godotenv"
)

var (
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Folder          string
	Exclude         []string
	BucketName      string
	IsCname         bool
	Client          *oss.Client
	Bucket          *oss.Bucket
	SkipSetting     bool

	IndexPage         string
	NotFoundPage      string
	HTMLCacheControl  string
	ImageCacheControl string
	OtherCacheControl string
	PageDataCacheControl string
)

func init() {
	godotenv.Load(".env")
	godotenv.Load(".env.local")

	Endpoint = os.Getenv("ENDPOINT")
	IsCname = os.Getenv("CNAME") == "true"
	AccessKeyID = os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret = os.Getenv("ACCESS_KEY_SECRET")
	Folder = os.Getenv("FOLDER")
	Exclude = utils.GetActionInputAsSlice(os.Getenv("EXCLUDE"))
	BucketName = os.Getenv("BUCKET")
	SkipSetting = os.Getenv("SKIP_SETTING") == "true"

	IndexPage = utils.Getenv("INDEX_PAGE", "index.html")
	NotFoundPage = utils.Getenv("NOT_FOUND_PAGE", "404.html")
	HTMLCacheControl = utils.Getenv("HTML_CACHE_CONTROL", "public, max-age=0, must-revalidate")
	ImageCacheControl = utils.Getenv("IMAGE_CACHE_CONTROL", "public, max-age=31536000, immutable")
	OtherCacheControl = utils.Getenv("OTHER_CACHE_CONTROL", "public, max-age=31536000, immutable")
	PageDataCacheControl = utils.Getenv("PAGE_DATA_CACHE_CONTROL", "public, max-age=0, must-revalidate")

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("current directory: %s\n", currentPath)
	fmt.Printf("endpoint: %s\nbucketName: %s\nfolder: %s\nexclude: %v\nindexPage: %s\nnotFoundPage: %s\nisCname: %t\nskipSetting: %t\n",
		Endpoint, BucketName, Folder, Exclude, IndexPage, NotFoundPage, IsCname, SkipSetting)
	fmt.Printf("HTMLCacheControl: %s\nimageCacheControl: %s\notherCacheControl: %s\n",
		HTMLCacheControl, ImageCacheControl, OtherCacheControl)

	Client, err = oss.New(Endpoint, AccessKeyID, AccessKeySecret, oss.UseCname(IsCname))
	if err != nil {
		utils.HandleError(err)
	}

	Bucket, err = Client.Bucket(BucketName)
	if err != nil {
		utils.HandleError(err)
	}
}
