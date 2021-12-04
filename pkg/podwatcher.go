package pkg

import (
	"fmt"
	"net"
	"time"
)

func getListOfPods(headlesssvc string) ([]string, error) {
	var pods = make([]string, 0)
	ips, err := net.LookupIP(headlesssvc)
	if err != nil {
		return nil, fmt.Errorf("Error in nslookup %+v", err)
	}
	for _, ip := range ips {
		pods = append(pods, ip.String())
	}
	return pods, nil
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func Watch(headlessService string, interval time.Duration, newPodCallBack func(address []string), failPodCallback func(address []string)) {
	var pods []string
	for {
		currentPods, err := getListOfPods(headlessService)
		if err != nil {
			fmt.Printf("Error in retriving pods %+v\n", err)
		} else {
			newPods := difference(currentPods, pods)
			if len(newPods) > 0 {
				newPodCallBack(newPods)
			}
			failedPods := difference(pods, currentPods)
			if len(failedPods) > 0 {
				failPodCallback(failedPods)
			}
			pods = currentPods
		}

		time.Sleep(interval * time.Second)
	}
}
