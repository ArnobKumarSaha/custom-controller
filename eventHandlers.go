package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

// enqueuemessi takes a messi resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than messi.
func (c *Controller) messiAdderFunction(obj interface{}) {
	fmt.Println("messiAdderFunction is called")
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

// This function is same as MessiAdderFunction except the print message.
func (c *Controller) messiDeleteFunction(obj interface{}) {
	fmt.Println("messiDeleteFunction is called")
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

// handleObject will take any resource implementing metav1.Object and attempt
// to find the messi resource that 'owns' it. It does this by looking at the
// objects metadata.ownerReferences field for an appropriate OwnerReference.
// It then enqueues that messi resource to be processed. If the object does not
// have an appropriate OwnerReference, it will simply be skipped.
func (c *Controller) deploymentAdderFunction(obj interface{}) {
	fmt.Println("deploymentAdderFunction is called")
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		klog.V(4).Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	klog.V(4).Infof("Processing object: %s", object.GetName())
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		// If this object is not owned by a messi, we should not do anything more
		// with it.
		if ownerRef.Kind != "Messi" {
			return
		}

		messi, err := c.messiLister.Messis(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			klog.V(4).Infof("ignoring orphaned object '%s' of messi '%s'", object.GetSelfLink(), ownerRef.Name)
			return
		}

		c.messiAdderFunction(messi)
		return
	}
}

// This function is same as DeploymentAdderFunction except the print message.
func (c *Controller) deploymentDeleteFunction(obj interface{}) {
	fmt.Println("deploymentDeleteFunction is called")
	var object metav1.Object
	var ok bool
	fmt.Println("Before the 1st if in deplDeleteFunc")
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		klog.V(4).Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	fmt.Println("After the 1st if in deplDeleteFunc")
	klog.V(4).Infof("Processing object: %s", object.GetName())
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		fmt.Println("Inside the 2nd if in deplDeleteFunc")
		// If this object is not owned by a messi, we should not do anything more
		// with it.
		if ownerRef.Kind != "Messi" {
			return
		}

		messi, err := c.messiLister.Messis(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			klog.V(4).Infof("ignoring orphaned object '%s' of messi '%s'", object.GetSelfLink(), ownerRef.Name)
			return
		}

		c.messiDeleteFunction(messi)
		return
	}
	fmt.Println("Before the 2nd if in deplDeleteFunc")
}

func (c *Controller) serviceAdderFunction(obj interface{}) {
	fmt.Println("serviceAdderFunction is called")
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		klog.V(4).Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	klog.V(4).Infof("Processing object: %s", object.GetName())
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		// If this object is not owned by a messi, we should not do anything more
		// with it.
		if ownerRef.Kind != "Messi" {
			return
		}

		messi, err := c.messiLister.Messis(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			klog.V(4).Infof("ignoring orphaned object '%s' of messi '%s'", object.GetSelfLink(), ownerRef.Name)
			return
		}

		c.messiAdderFunction(messi)
		return
	}
}

// This function is same as ServiceAdderFunction except the print message.
func (c *Controller) serviceDeleteFunction(obj interface{}) {
	fmt.Println("serviceDeleteFunction is called")
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		klog.V(4).Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	klog.V(4).Infof("Processing object: %s", object.GetName())
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		// If this object is not owned by a messi, we should not do anything more
		// with it.
		if ownerRef.Kind != "Messi" {
			return
		}

		messi, err := c.messiLister.Messis(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			klog.V(4).Infof("ignoring orphaned object '%s' of messi '%s'", object.GetSelfLink(), ownerRef.Name)
			return
		}

		c.messiDeleteFunction(messi)
		return
	}
}
