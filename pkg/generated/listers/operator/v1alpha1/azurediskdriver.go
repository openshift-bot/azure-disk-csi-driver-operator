// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/azure-disk-csi-driver-operator/pkg/apis/operator/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AzureDiskDriverLister helps list AzureDiskDrivers.
type AzureDiskDriverLister interface {
	// List lists all AzureDiskDrivers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AzureDiskDriver, err error)
	// Get retrieves the AzureDiskDriver from the index for a given name.
	Get(name string) (*v1alpha1.AzureDiskDriver, error)
	AzureDiskDriverListerExpansion
}

// azureDiskDriverLister implements the AzureDiskDriverLister interface.
type azureDiskDriverLister struct {
	indexer cache.Indexer
}

// NewAzureDiskDriverLister returns a new AzureDiskDriverLister.
func NewAzureDiskDriverLister(indexer cache.Indexer) AzureDiskDriverLister {
	return &azureDiskDriverLister{indexer: indexer}
}

// List lists all AzureDiskDrivers in the indexer.
func (s *azureDiskDriverLister) List(selector labels.Selector) (ret []*v1alpha1.AzureDiskDriver, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AzureDiskDriver))
	})
	return ret, err
}

// Get retrieves the AzureDiskDriver from the index for a given name.
func (s *azureDiskDriverLister) Get(name string) (*v1alpha1.AzureDiskDriver, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("azurediskdriver"), name)
	}
	return obj.(*v1alpha1.AzureDiskDriver), nil
}