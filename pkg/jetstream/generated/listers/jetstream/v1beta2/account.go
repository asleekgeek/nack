// Copyright 2025 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1beta2

import (
	jetstreamv1beta2 "github.com/nats-io/nack/pkg/jetstream/apis/jetstream/v1beta2"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// AccountLister helps list Accounts.
// All objects returned here must be treated as read-only.
type AccountLister interface {
	// List lists all Accounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*jetstreamv1beta2.Account, err error)
	// Accounts returns an object that can list and get Accounts.
	Accounts(namespace string) AccountNamespaceLister
	AccountListerExpansion
}

// accountLister implements the AccountLister interface.
type accountLister struct {
	listers.ResourceIndexer[*jetstreamv1beta2.Account]
}

// NewAccountLister returns a new AccountLister.
func NewAccountLister(indexer cache.Indexer) AccountLister {
	return &accountLister{listers.New[*jetstreamv1beta2.Account](indexer, jetstreamv1beta2.Resource("account"))}
}

// Accounts returns an object that can list and get Accounts.
func (s *accountLister) Accounts(namespace string) AccountNamespaceLister {
	return accountNamespaceLister{listers.NewNamespaced[*jetstreamv1beta2.Account](s.ResourceIndexer, namespace)}
}

// AccountNamespaceLister helps list and get Accounts.
// All objects returned here must be treated as read-only.
type AccountNamespaceLister interface {
	// List lists all Accounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*jetstreamv1beta2.Account, err error)
	// Get retrieves the Account from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*jetstreamv1beta2.Account, error)
	AccountNamespaceListerExpansion
}

// accountNamespaceLister implements the AccountNamespaceLister
// interface.
type accountNamespaceLister struct {
	listers.ResourceIndexer[*jetstreamv1beta2.Account]
}
