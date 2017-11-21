# mackerel-plugin-katsubushi

Mackerel plugin for [Katsubushi](https://github.com/kayac/go-katsubushi).

## Synopsis

```shell
mackerel-plugin-katsubushi [-metric-key-prefix=<prefix>]
```

## Example of mackerel-agent.conf

```
[plugin.metrics.sample]
command = "/path/to/mackerel-plugin-katsubushi"
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

KAYAC Inc. / Mackerel

## LICENCE

Apache 2.0
