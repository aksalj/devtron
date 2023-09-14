// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	pg "github.com/go-pg/pg"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/devtron-labs/devtron/pkg/cluster/repository"
)

// EnvironmentRepository is an autogenerated mock type for the EnvironmentRepository type
type EnvironmentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: mappings
func (_m *EnvironmentRepository) Create(mappings *repository.Environment) error {
	ret := _m.Called(mappings)

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.Environment) error); ok {
		r0 = rf(mappings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *EnvironmentRepository) FindAll() ([]repository.Environment, error) {
	ret := _m.Called()

	var r0 []repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]repository.Environment, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []repository.Environment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllActive provides a mock function with given fields:
func (_m *EnvironmentRepository) FindAllActive() ([]*repository.Environment, error) {
	ret := _m.Called()

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*repository.Environment, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*repository.Environment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllActiveEnvOnlyDetails provides a mock function with given fields:
func (_m *EnvironmentRepository) FindAllActiveEnvOnlyDetails() ([]*repository.Environment, error) {
	ret := _m.Called()

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*repository.Environment, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*repository.Environment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllActiveWithFilter provides a mock function with given fields:
func (_m *EnvironmentRepository) FindAllActiveWithFilter() ([]*repository.Environment, error) {
	ret := _m.Called()

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*repository.Environment, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*repository.Environment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByClusterId provides a mock function with given fields: clusterId
func (_m *EnvironmentRepository) FindByClusterId(clusterId int) ([]*repository.Environment, error) {
	ret := _m.Called(clusterId)

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]*repository.Environment, error)); ok {
		return rf(clusterId)
	}
	if rf, ok := ret.Get(0).(func(int) []*repository.Environment); ok {
		r0 = rf(clusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(clusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByClusterIdAndNamespace provides a mock function with given fields: namespaceClusterPair
func (_m *EnvironmentRepository) FindByClusterIdAndNamespace(namespaceClusterPair []*repository.ClusterNamespacePair) ([]*repository.Environment, error) {
	ret := _m.Called(namespaceClusterPair)

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func([]*repository.ClusterNamespacePair) ([]*repository.Environment, error)); ok {
		return rf(namespaceClusterPair)
	}
	if rf, ok := ret.Get(0).(func([]*repository.ClusterNamespacePair) []*repository.Environment); ok {
		r0 = rf(namespaceClusterPair)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func([]*repository.ClusterNamespacePair) error); ok {
		r1 = rf(namespaceClusterPair)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByClusterIds provides a mock function with given fields: clusterIds
func (_m *EnvironmentRepository) FindByClusterIds(clusterIds []int) ([]*repository.Environment, error) {
	ret := _m.Called(clusterIds)

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) ([]*repository.Environment, error)); ok {
		return rf(clusterIds)
	}
	if rf, ok := ret.Get(0).(func([]int) []*repository.Environment); ok {
		r0 = rf(clusterIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(clusterIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByClusterIdsWithFilter provides a mock function with given fields: clusterIds
func (_m *EnvironmentRepository) FindByClusterIdsWithFilter(clusterIds []int) ([]*repository.Environment, error) {
	ret := _m.Called(clusterIds)

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) ([]*repository.Environment, error)); ok {
		return rf(clusterIds)
	}
	if rf, ok := ret.Get(0).(func([]int) []*repository.Environment); ok {
		r0 = rf(clusterIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(clusterIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEnvName provides a mock function with given fields: envName
func (_m *EnvironmentRepository) FindByEnvName(envName string) ([]*repository.Environment, error) {
	ret := _m.Called(envName)

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*repository.Environment, error)); ok {
		return rf(envName)
	}
	if rf, ok := ret.Get(0).(func(string) []*repository.Environment); ok {
		r0 = rf(envName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(envName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEnvNameAndClusterIds provides a mock function with given fields: envName, clusterIds
func (_m *EnvironmentRepository) FindByEnvNameAndClusterIds(envName string, clusterIds []int) ([]*repository.Environment, error) {
	ret := _m.Called(envName, clusterIds)

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []int) ([]*repository.Environment, error)); ok {
		return rf(envName, clusterIds)
	}
	if rf, ok := ret.Get(0).(func(string, []int) []*repository.Environment); ok {
		r0 = rf(envName, clusterIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []int) error); ok {
		r1 = rf(envName, clusterIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *EnvironmentRepository) FindById(id int) (*repository.Environment, error) {
	ret := _m.Called(id)

	var r0 *repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*repository.Environment, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *repository.Environment); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIdentifier provides a mock function with given fields: identifier
func (_m *EnvironmentRepository) FindByIdentifier(identifier string) (*repository.Environment, error) {
	ret := _m.Called(identifier)

	var r0 *repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*repository.Environment, error)); ok {
		return rf(identifier)
	}
	if rf, ok := ret.Get(0).(func(string) *repository.Environment); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIds provides a mock function with given fields: ids
func (_m *EnvironmentRepository) FindByIds(ids []*int) ([]*repository.Environment, error) {
	ret := _m.Called(ids)

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func([]*int) ([]*repository.Environment, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]*int) []*repository.Environment); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func([]*int) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: name
func (_m *EnvironmentRepository) FindByName(name string) (*repository.Environment, error) {
	ret := _m.Called(name)

	var r0 *repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*repository.Environment, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *repository.Environment); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByNameOrIdentifier provides a mock function with given fields: name, identifier
func (_m *EnvironmentRepository) FindByNameOrIdentifier(name string, identifier string) (*repository.Environment, error) {
	ret := _m.Called(name, identifier)

	var r0 *repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*repository.Environment, error)); ok {
		return rf(name, identifier)
	}
	if rf, ok := ret.Get(0).(func(string, string) *repository.Environment); ok {
		r0 = rf(name, identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByNames provides a mock function with given fields: envNames
func (_m *EnvironmentRepository) FindByNames(envNames []string) ([]*repository.Environment, error) {
	ret := _m.Called(envNames)

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]*repository.Environment, error)); ok {
		return rf(envNames)
	}
	if rf, ok := ret.Get(0).(func([]string) []*repository.Environment); ok {
		r0 = rf(envNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(envNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByNamespaceAndClusterName provides a mock function with given fields: namespaces, clusterName
func (_m *EnvironmentRepository) FindByNamespaceAndClusterName(namespaces string, clusterName string) (*repository.Environment, error) {
	ret := _m.Called(namespaces, clusterName)

	var r0 *repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*repository.Environment, error)); ok {
		return rf(namespaces, clusterName)
	}
	if rf, ok := ret.Get(0).(func(string, string) *repository.Environment); ok {
		r0 = rf(namespaces, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespaces, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEnvClusterInfosByIds provides a mock function with given fields: _a0
func (_m *EnvironmentRepository) FindEnvClusterInfosByIds(_a0 []int) ([]*repository.EnvCluserInfo, error) {
	ret := _m.Called(_a0)

	var r0 []*repository.EnvCluserInfo
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) ([]*repository.EnvCluserInfo, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]int) []*repository.EnvCluserInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.EnvCluserInfo)
		}
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEnvLinkedWithCiPipelines provides a mock function with given fields: externalCi, ciPipelineIds
func (_m *EnvironmentRepository) FindEnvLinkedWithCiPipelines(externalCi bool, ciPipelineIds []int) ([]*repository.Environment, error) {
	ret := _m.Called(externalCi, ciPipelineIds)

	var r0 []*repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(bool, []int) ([]*repository.Environment, error)); ok {
		return rf(externalCi, ciPipelineIds)
	}
	if rf, ok := ret.Get(0).(func(bool, []int) []*repository.Environment); ok {
		r0 = rf(externalCi, ciPipelineIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(bool, []int) error); ok {
		r1 = rf(externalCi, ciPipelineIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIdsByNames provides a mock function with given fields: envNames
func (_m *EnvironmentRepository) FindIdsByNames(envNames []string) ([]int, error) {
	ret := _m.Called(envNames)

	var r0 []int
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]int, error)); ok {
		return rf(envNames)
	}
	if rf, ok := ret.Get(0).(func([]string) []int); ok {
		r0 = rf(envNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(envNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: environment
func (_m *EnvironmentRepository) FindOne(environment string) (*repository.Environment, error) {
	ret := _m.Called(environment)

	var r0 *repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*repository.Environment, error)); ok {
		return rf(environment)
	}
	if rf, ok := ret.Get(0).(func(string) *repository.Environment); ok {
		r0 = rf(environment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(environment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneByNamespaceAndClusterId provides a mock function with given fields: namespace, clusterId
func (_m *EnvironmentRepository) FindOneByNamespaceAndClusterId(namespace string, clusterId int) (*repository.Environment, error) {
	ret := _m.Called(namespace, clusterId)

	var r0 *repository.Environment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) (*repository.Environment, error)); ok {
		return rf(namespace, clusterId)
	}
	if rf, ok := ret.Get(0).(func(string, int) *repository.Environment); ok {
		r0 = rf(namespace, clusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(namespace, clusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnection provides a mock function with given fields:
func (_m *EnvironmentRepository) GetConnection() *pg.DB {
	ret := _m.Called()

	var r0 *pg.DB
	if rf, ok := ret.Get(0).(func() *pg.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg.DB)
		}
	}

	return r0
}

// MarkEnvironmentDeleted provides a mock function with given fields: mappings, tx
func (_m *EnvironmentRepository) MarkEnvironmentDeleted(mappings *repository.Environment, tx *pg.Tx) error {
	ret := _m.Called(mappings, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.Environment, *pg.Tx) error); ok {
		r0 = rf(mappings, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: mappings
func (_m *EnvironmentRepository) Update(mappings *repository.Environment) error {
	ret := _m.Called(mappings)

	var r0 error
	if rf, ok := ret.Get(0).(func(*repository.Environment) error); ok {
		r0 = rf(mappings)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewEnvironmentRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewEnvironmentRepository creates a new instance of EnvironmentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEnvironmentRepository(t mockConstructorTestingTNewEnvironmentRepository) *EnvironmentRepository {
	mock := &EnvironmentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
