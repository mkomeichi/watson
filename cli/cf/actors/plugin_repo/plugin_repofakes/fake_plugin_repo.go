// This file was generated by counterfeiter
package plugin_repofakes

import (
	"sync"

	clipr "github.com/cloudfoundry-incubator/cli-plugin-repo/models"
	"github.com/cloudfoundry/cli/cf/actors/plugin_repo"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakePluginRepo struct {
	GetPluginsStub        func([]models.PluginRepo) (map[string][]clipr.Plugin, []string)
	getPluginsMutex       sync.RWMutex
	getPluginsArgsForCall []struct {
		arg1 []models.PluginRepo
	}
	getPluginsReturns struct {
		result1 map[string][]clipr.Plugin
		result2 []string
	}
}

func (fake *FakePluginRepo) GetPlugins(arg1 []models.PluginRepo) (map[string][]clipr.Plugin, []string) {
	fake.getPluginsMutex.Lock()
	fake.getPluginsArgsForCall = append(fake.getPluginsArgsForCall, struct {
		arg1 []models.PluginRepo
	}{arg1})
	fake.getPluginsMutex.Unlock()
	if fake.GetPluginsStub != nil {
		return fake.GetPluginsStub(arg1)
	} else {
		return fake.getPluginsReturns.result1, fake.getPluginsReturns.result2
	}
}

func (fake *FakePluginRepo) GetPluginsCallCount() int {
	fake.getPluginsMutex.RLock()
	defer fake.getPluginsMutex.RUnlock()
	return len(fake.getPluginsArgsForCall)
}

func (fake *FakePluginRepo) GetPluginsArgsForCall(i int) []models.PluginRepo {
	fake.getPluginsMutex.RLock()
	defer fake.getPluginsMutex.RUnlock()
	return fake.getPluginsArgsForCall[i].arg1
}

func (fake *FakePluginRepo) GetPluginsReturns(result1 map[string][]clipr.Plugin, result2 []string) {
	fake.GetPluginsStub = nil
	fake.getPluginsReturns = struct {
		result1 map[string][]clipr.Plugin
		result2 []string
	}{result1, result2}
}

var _ plugin_repo.PluginRepo = new(FakePluginRepo)
