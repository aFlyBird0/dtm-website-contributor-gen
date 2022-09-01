# dtm-website-contributor-gen

DevStream 的官网有个[社区英雄榜](https://www.devstream.io/community/community-heroes/openSourceContributor/Associate), 底层是由 json 文件渲染的。
这个仓库用来从贡献者表格直接生成相应的 json 文件，并且直接提个 pr 到 [website](https://github.com/devstream-io/website) 仓库。

## 使用方法
视图切换到 **"下载专用（脱敏）"**，下载贡献者表格，保存成 `csv` 格式，放到此项目内某文件夹，并且设置 `/.github/workflows/pr-to-website.yml` 中的 `--data` 为该文件夹，该文件夹下的所有 `csv` 文件都会被读取并融合成一个 json。
csv 文件模板在 `/secretExample` 文件下。

之后就是更新 `csv` 文件，`git add, commit, push`, 等待 pr 自动生成。


## 自定义配置

发起 PR 的 GitHub Actions 官方文档：https://github.com/peter-evans/create-pull-request

### 1. 由 DevStream 其他组员使用
1. 需要 fork 一下 website 仓库到自己账号下。
2. 全局替换 `aFlyBird0` 成自己的 GitHub ID，通读一遍 `/.github/workflows/pr-to-website.yml`，修改觉得有必要的地方。
3. 给仓库设置一个名为 `TOKEN_WRITE` 的 secrets，值是一个拥有仓库写权限的 GitHub Personal Access Token。

### 2. 由其他社区使用
在 `1. 由 DevStream 其他组员使用` 的基础上，针对 `csv` 的格式、仓库的地址、社区自己的 website 的 json 文件存储位置等内容，修改代码和 `/.github/workflows/pr-to-website.yml` 文件。

## 自动化
可以配合额外的脚本执行，将 `csv` 下载和 git 操作集成到脚本里，然后把脚本放到桌面。

这样就能实现在飞书表格页面点一下下载，然后再到桌面双击运行一下脚本，后续自动执行 `csv` 位置移动、git add、git commit、git push、json 转换、提 pr 操作。

参考脚本如下，请自行替换相应信息：

```shell
rm -rf ~/Code/Go/tmp/dtm-website-contributor-gen/secretExample/**.csv
mv ~/Downloads/"开源社区贡献者与证书管理_所有证书_下载专用（脱敏）.csv" ~/Code/Go/tmp/dtm-website-contributor-gen/secretExample/
cd ~/Code/Go/tmp/dtm-website-contributor-gen
git add .
git commit -m "feat: update contributors info(auto by shell)" -s
git push origin
```
