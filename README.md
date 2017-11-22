# mackerel-plugin-katsubushi

Mackerel plugin for [Katsubushi](https://github.com/kayac/go-katsubushi).

## Install

Use [mkr](https://github.com/mackerelio/mkr).

```console
# mkr plugin install kayac/mackerel-plugin-katsubushi@v0.0.0
```

## Synopsis

```shell
Usage of /opt/mackerel-agent/plugins/bin/mackerel-plugin-katsubushi:
  -host string
    	Hostname (default "localhost")
  -metric-key-prefix string
    	Metric key prefix (default "katsubushi")
  -port string
    	Port (default "11212")
  -socket string
    	Server socket (overrides hosts and port)
  -tempfile string
    	Temp file name
```

## Example of mackerel-agent.conf

```
[plugin.metrics.katsubushi]
command = "/opt/mackerel-agent/plugins/bin/mackerel-plugin-katsubushi/mackerel-plugin-katsubushi"
```

## How to release

[goxc](https://github.com/laher/goxc) and [ghr](https://github.com/tcnksm/ghr) are used to release.

### Release by manually

1. Install goxc and ghr by `make setup`
2. Edit CHANGELOG.md, git commit, git push
3. `git tag vx.y.z`
4. GITHUB_TOKEN=... make release
5. See https://github.com/mackerelio/mackerel-plugin-katsubushi/releases

## Author

KAYAC Inc.

## License

Copyright 2014 Hatena Co., Ltd.
Copyright 2017 KAYAC Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
