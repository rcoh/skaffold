/*
Copyright 2019 The Skaffold Authors

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

package shared

import (
	plugin "github.com/hashicorp/go-plugin"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion: 1,

	MagicCookieKey:   "SKAFFOLD_BUILDER_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is a map of all accepted plugins
var PluginMap = map[string]plugin.Plugin{
	"docker": &BuilderPlugin{},
	"bazel":  &BuilderPlugin{},
}
