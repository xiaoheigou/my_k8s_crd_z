package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/xiaoheigou/mycrd/pkg/apis/myresource/v1"
	myresourceclientset "github.com/xiaoheigou/mycrd/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{}, client kubernetes.Interface, myresourceclient *myresourceclientset.Clientset)
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}, client kubernetes.Interface, myresourceclient *myresourceclientset.Clientset) {
	list, _ := client.CoreV1().Nodes().List(metav1.ListOptions{})
	pods, _ := client.CoreV1().Pods("").List(metav1.ListOptions{})

	myresourceList, _ := myresourceclient.TrstringerV1().MyResources("").List(metav1.ListOptions{})
	myresource := obj.(*v1.MyResource)
	myresource.Status.ResourceNumber = make(map[string]int)
	myresource.Status.ResourceNumber["pod"] = len(pods.Items)
	myresource.Status.ResourceNumber["nodes"] = len(list.Items)
	myresource.Status.ResourceNumber["myresource"] = len(myresourceList.Items)
	_, err := myresourceclient.TrstringerV1().MyResources("default").Update(myresource)
	if err != nil {
		log.Errorf("err:%s", err)
	}
	log.Infof("TestHandler.ObjectCreated-----\n[nodeNumber:%d][podsNumber:%d][obj:%v]", len(list.Items), len(pods.Items), myresource)
	// fmt.Println(myresource.Spec.Message)
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}
