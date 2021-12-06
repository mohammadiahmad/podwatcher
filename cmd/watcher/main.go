package cmd

import (
	"fmt"
	pw "github.com/mohammadiahmad/podwatcher/pkg"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

const (
	use   = "watch"
	short = "notify pod added or removed in k8s using headless service"
)

func Watcher() *cobra.Command {
	// nolint: exhaustivestruct
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	cmd.PersistentFlags().String("headless", "", "headless serivce name")
	cmd.PersistentFlags().Duration("interval", 1, "time interval for checking list of pods")
	return cmd
}

func main(cmd *cobra.Command, _ []string) {
	headlessSVC, _ := cmd.PersistentFlags().GetString("headless")
	interval, _ := cmd.PersistentFlags().GetDuration("interval")

	go pw.Watch(headlessSVC, interval, newPodCallBack, failPodCallBack)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

}

func newPodCallBack(pods []string) {
	for _, pod := range pods {
		fmt.Printf("Pod added with address: %s\n", pod)
	}
}

func failPodCallBack(pods []string) {
	for _, pod := range pods {
		fmt.Printf("Pod remove with address: %s\n", pod)
	}
}
