package main

import (
	"fmt"
	"minik8s/pkg/netSupport/tools"
)

func main() {
	fmt.Println(tools.GetDocker0IpAndMask())
	//err := boot.BootFlannel()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//select {}
	//clientConfig := client.Config{Host: "127.0.0.1:8080"}
	////restClient := client.RESTClient{
	////	Base: "http://" + clientConfig.Host,
	////}
	//data, err := ioutil.ReadFile("D:\\goLandProject\\minik8s\\minik8s\\test\\pod\\gpuPod.yaml")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//pod := &object.Pod{}
	//err = yaml.Unmarshal([]byte(data), &pod)
	//fmt.Println(*pod)
	//pod.UID = uuid.NewV4().String()
	//manager := podManager.NewPodManager(clientConfig)
	//manager.StartPodManager()
	//err = manager.AddPod(pod)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//select {}
	//for {
	//time.Sleep(10 * time.Second)
	//res, err2 := manager.GetPodSnapShoot(pod.Name)
	//if err2 != nil {
	//	fmt.Println(err2)
	//} else {
	//	fmt.Println(res)
	//}
	//}
}
