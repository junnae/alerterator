// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/nais/alerterator/pkg/apis/alerterator/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AlertLister helps list Alerts.
type AlertLister interface {
	// List lists all Alerts in the indexer.
	List(selector labels.Selector) (ret []*v1.Alert, err error)
	// Alerts returns an object that can list and get Alerts.
	Alerts(namespace string) AlertNamespaceLister
	AlertListerExpansion
}

// alertLister implements the AlertLister interface.
type alertLister struct {
	indexer cache.Indexer
}

// NewAlertLister returns a new AlertLister.
func NewAlertLister(indexer cache.Indexer) AlertLister {
	return &alertLister{indexer: indexer}
}

// List lists all Alerts in the indexer.
func (s *alertLister) List(selector labels.Selector) (ret []*v1.Alert, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Alert))
	})
	return ret, err
}

// Alerts returns an object that can list and get Alerts.
func (s *alertLister) Alerts(namespace string) AlertNamespaceLister {
	return alertNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlertNamespaceLister helps list and get Alerts.
type AlertNamespaceLister interface {
	// List lists all Alerts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Alert, err error)
	// Get retrieves the Alert from the indexer for a given namespace and name.
	Get(name string) (*v1.Alert, error)
	AlertNamespaceListerExpansion
}

// alertNamespaceLister implements the AlertNamespaceLister
// interface.
type alertNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Alerts in the indexer for a given namespace.
func (s alertNamespaceLister) List(selector labels.Selector) (ret []*v1.Alert, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Alert))
	})
	return ret, err
}

// Get retrieves the Alert from the indexer for a given namespace and name.
func (s alertNamespaceLister) Get(name string) (*v1.Alert, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("alert"), name)
	}
	return obj.(*v1.Alert), nil
}
