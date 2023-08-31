package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"xiudong/showstart"
)

var cfgFile string
var globalConfig *showstart.WapEncryptConfigV3

//var globalClient = resty.New()

var globalClient = getRetryClient()

func getRetryClient() *resty.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	httpClient := http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}

	return resty.NewWithClient(&httpClient)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "showstart",
	Long: `showstart cli sample`,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig) // todo: 其实这里也可以不要 自己进行 config 读取也可以的
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.showstart.yaml)",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//log.Println(home)
		// 家目录 .showstart.yaml windows 如何设置我不太清楚
		// todo: 可以改为当前目录
		viper.AddConfigPath(home)
		viper.SetConfigName(".showstart") // 不需要后缀
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("使用的配置文件:", viper.ConfigFileUsed())
		globalConfig = &showstart.WapEncryptConfigV3{
			Sign:        viper.GetString("sign"),
			StFlpv:      viper.GetString("st_flpv"),
			Token:       viper.GetString("token"),
			UserId:      viper.GetUint32("userId"),
			AccessToken: viper.GetString("accessToken"),
			IdToken:     viper.GetString("idToken"),
		}
		log.Printf("globalConfig is %+v\n", globalConfig)
	}
}
