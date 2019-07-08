/*
Copyright © 2019 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package plugin

type PluginManifest struct {
	Name    string
	Version string
}

type PluginManagerConfig struct {
	PluginDirs []string
	RootCmd    *cobra.Command
}

type PluginManager struct {
	config  PluginManagerConfig
	plugins []*PluginManifest
}

const (
	pluginExt = ".px"
)

func NewPluginManager(config *PluginManagerConfig) *PluginManager {
	return &PluginManager{
		config:  *config,
		plugins: make([]*PluginManifest, 0),
	}
}

func (p *PluginManager) load() error {

	for _, pluginDir := range p.config.PluginDirs {
		files, err := ioutil.ReadDir(pluginDir)
		if err != nil {
			continue
		}

		for _, file := range files {
			if filepath.Ext(file) == pluginExt {
				manifest, err := p.loadPlugin(file)
				if err != nil {
					util.Eprintf("%v\n", err)
					continue
				}
				p.plugins = append(p.plugins, manifest)
			}
		}
	}
}

func (p *PluginManager) List() ([]Plugin, error) {

}

func (p *PluginManager) loadPlugin(soPath string) (*PluginManifest, error) {
	// Open plugin library
	p, err := plugin.Open(soPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open plugin %s: %v\n", soPath, err)
	}

	// Get Plugin Manifest to get name and info
	m, err := p.Lookup("PluginManifest")
	if err != nil {
		return nil, fmt.Errorf("Plugin ___ does not have init function\n")
	}

	// Confirm the interface is correct
	manifestMap, ok := *m.(*map[string]string)
	if !ok {
		return nil, fmt.Errorf("Plugin %s failed to get manifest", soPath)
	}
	manifest := &PluginManifest{}
	if name, ok := manifestMap["name"]; !ok {
		return nil, fmt.Errorf("%s plugin is missing a required name", soPath)
	} else {
		manifest.Name = name
	}
	if name, ok := manifestMap["version"]; !ok {
		return nil, fmt.Errorf("%s plugin is missing a required version", soPath)
	} else {
		manifest.Name = version
	}

	// Get access to plugin init function
	f, err := p.Lookup("PluginInit")
	if err != nil {
		return nil, fmt.Errorf("Plugin %s does not have init function", soPath)
	}

	// Confirm the interface is correct
	pinit, ok := f.(func(*cobra.Command))
	if !ok {
		return nil, fmt.Errorf("Plugin %s does not have the correct init function", soPath)
	}

	// Finally, register the handler
	pinit(rootCmd)

	return manifest, nil
}
