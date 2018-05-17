package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"
	myresourceclientset "github.com/xiaoheigou/mycrd/pkg/client/clientset/versioned"
	myresourceinformer_v1 "github.com/xiaoheigou/mycrd/pkg/client/informers/externalversions/myresource/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

func getKubernetesClient() (kubernetes.Interface, myresourceclientset.Interface) {
	// construct the path to resolve to `~/.kube/config`
	// kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	kubeConfigPath := "/etc/kubernetes/kubectl.kubeconfig"

	// create the config from the path
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	// generate the client based off of the config
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}
	myresourceClient, err := myresourceclientset.NewForConfig(config)

	log.Info("Successfully constructed k8s client")
	return client, myresourceClient
}

func main() {
	// ## 1.components
	// get the Kubernetes client for connectivity
	client, myresourceClient := getKubernetesClient()
	// list, err := client.CoreV1().Nodes().List(metav1.ListOptions{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, node := range list.Items {
	// 	fmt.Printf("Node : %s \n", node.Name)
	// }
	informer := myresourceinformer_v1.NewMyResourceInformer(
		myresourceClient,
		meta_v1.NamespaceAll,
		0,
		cache.Indexers{},
	)
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			// convert the resource object into a key (in this case
			// we are just doing it in the format of 'namespace/name')
			key, err := cache.MetaNamespaceKeyFunc(obj)
			log.Infof("Add myresource: %s", key)
			if err == nil {
				// add the key to the queue for the handler to get
				queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			log.Infof("Update myresource: %s", key)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// DeletionHandlingMetaNamsespaceKeyFunc is a helper function that allows
			// us to check the DeletedFinalStateUnknown existence in the event that
			// a resource was deleted but it is still contained in the index
			//
			// this then in turn calls MetaNamespaceKeyFunc
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			log.Infof("Delete myresource: %s", key)
			if err == nil {
				queue.Add(key)
			}
		},
	})

	// ## 2.controller
	controller := Controller{
		logger:    log.NewEntry(log.New()),
		clientset: client,
		informer:  informer,
		queue:     queue,
		handler:   &TestHandler{},
	}
	// use a channel to synchronize the finalization for a graceful shutdown
	stopCh := make(chan struct{})
	defer close(stopCh)

	// run the controller loop to process items
	go controller.Run(stopCh)

	// use a channel to handle OS signals to terminate and gracefully shut
	// down processing
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	<-sigTerm
}
