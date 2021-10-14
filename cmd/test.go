package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
)

var endPoints []string = []string{"env", "docs", "dump", "info", "docs/", "beans", "trace", "caches", "eureka", "flyway", "health", "api/doc", "hystrix", "jolokia", "logfile", "loggers", "metrics", "monitor", "refresh", "swagger", "actuator", "api.html", "api-docs", "doc.html", "heapdump", "mappings", "sessions", "shutdown", "httptrace", "liquibase", "swagger/ui", "autoconfig", "conditions", "entity/all", "env/(name)", "prometheus", "swagger-ui", "threaddump", "api/swagger", "v2/api-docs", "v1/api-docs", "v3/api-docs", "auditevents", "configprops", "actuator/env", "jolokia/list", "actuator/info", "api/swaggerui", "heapdump.json", "swagger/codes", "actuator/beans", "actuator/trace", "api/index.html", "libs/swaggerui", "api/swagger/ui", "hystrix.stream", "scheduledtasks", "actuator/health", "api/v2/api-docs", "swagger-ui.html", "swagger-ui/html", "v2/swagger.json", "actuator/jolokia", "actuator/logfile", "actuator/loggers", "actuator/metrics", "Swagger/ui/index", "druid/index.html", "druid/login.html", "gateway/actuator", "actuator/heapdump", "actuator/mappings", "distv2/index.html", "intergrationgraph", "actuator/httptrace", "sw/swagger-ui.html", "swagger/index.html", "%20/swagger-ui.html", "actuator/conditions", "actuator/threaddump", "api/swagger-ui.html", "static/swagger.json", "actuator/auditevents", "actuator/configprops", "gateway/actuator/env", "user/swagger-ui.html", "v1.1/swagger-ui.html", "v1.2/swagger-ui.html", "v1.3/swagger-ui.html", "v1.4/swagger-ui.html", "v1.5/swagger-ui.html", "v1.6/swagger-ui.html", "v1.7/swagger-ui.html", "v1.8/swagger-ui.html", "v1.9/swagger-ui.html", "v2.0/swagger-ui.html", "v2.1/swagger-ui.html", "v2.2/swagger-ui.html", "v2.3/swagger-ui.html", "druid/websession.html", "gateway/actuator/info", "swagger-ui/index.html", "gateway/actuator/beans", "gateway/actuator/trace", "swagger-dubbo/api-docs", "actuator/hystrix.stream", "actuator/scheduledtasks", "cloudfoundryapplication", "gateway/actuator/health", "swagger/swagger-ui.html", "system/druid/index.html", "swagger/v1/swagger.json", "swagger/v2/swagger.json", "actuator/swagger-ui.html", "gateway/actuator/jolokia", "gateway/actuator/logfile", "gateway/actuator/loggers", "gateway/actuator/metrics", "template/swagger-ui.html", "gateway/actuator/heapdump", "gateway/actuator/mappings", "swagger/static/index.html", "gateway/actuator/httptrace", "gateway/actuator/conditions", "gateway/actuator/threaddump", "gateway/actuator/auditevents", "gateway/actuator/configprops", "gateway/actuator/hystrix.stream", "gateway/actuator/scheduledtasks", "webpage/system/druid/index.html", "dubbo-provider/distv2/index.html", "gateway/actuator/swagger-ui.html", "swagger-resources/configuration/ui", "swagger-resources/configuration/security", "spring-security-rest/api/swagger-ui.html", "spring-security-oauth-resource/swagger-ui.html"}
var threadNumber = 1

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		var wg sync.WaitGroup
		threadLimiter := make(chan struct{}, threadNumber)
		for _, endPoint := range endPoints {
			wg.Add(1)
			threadLimiter <- struct{}{}
			go func(endPoint string) {
				defer func() {
					wg.Done()
					<-threadLimiter
				}()
				u, err := concatUrl(host, endPoint)
				if err != nil {
					fmt.Println(err)
					return
				}
				getEndpoint(u)
			}(endPoint)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getEndpoint(u string) {
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(u + "	status:" + strconv.Itoa(resp.StatusCode) + "	content-length:" + strconv.Itoa(int(resp.ContentLength)))
}
