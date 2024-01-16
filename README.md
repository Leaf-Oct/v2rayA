Original project: https://github.com/v2rayA/v2rayA

In my PC (Ubuntu)，once I run v2rayA in webGUI, I still need to set "Network Proxy"(in gnome-control-center) to "manual"。And need to set it to "none" manually when I disconnect proxy in webGUI. This means that I need two steps to visit google successfully, first run v2ray, then set system proxy in setting-panel.

Can I do them in one step? Set system proxy doesn't need higher permission. So I try adding some code in `service/server/controller/connection.go`，which will execute "gsettings" command to change system proxy settings when turn on or off v2ray service.

# v2rayA [![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/v2rayA/v2raya)](https://hub.docker.com/r/mzz2017/v2raya) [![Travis (.org)](https://img.shields.io/travis/v2rayA/v2rayA?label=travis-ci%20build)](https://travis-ci.org/v2rayA/v2rayA)

[**English**](https://github.com/Leaf-Oct/v2rayA/blob/master/README.md)&nbsp;&nbsp;&nbsp;[**简体中文**](https://github.com/Leaf-Oct/v2rayA/blob/master/README_zh.md)

v2rayA is a V2Ray client supporting global transparent proxy on Linux and system proxy on Windows and macOS, it is compatible with SS, SSR, Trojan(trojan-go), Tuic and [Juicity](https://github.com/juicity) protocols. [[SSR protocol list]](https://github.com/v2rayA/shadowsocksR/blob/main/README.md#ss-encrypting-algorithm)

We are committed to providing the simplest operation and meet most needs.

Thanks to the advantages of Web GUI, you can not only use it on your local computer, but also easily deploy it on a router or NAS.

Project：https://github.com/v2rayA/v2rayA


## Usage

v2rayA mainly provides the following methods of installation:

1. Install from apt-source or AUR
2. Docker
3. Our self-built [scoop bucket](https://github.com/v2rayA/v2raya-scoop) (for Windows users)
4. Our self-built [homebrew tap](https://github.com/v2rayA/homebrew-v2raya)
5. Our self-built [OpenWrt repo](https://github.com/v2rayA/v2raya-openwrt) and OpenWrt's official repo(from OpenWrt version 22.03)
6. Microsoft winget: https://winstall.app/apps/v2rayA.v2rayA
7. Ubuntu Snap: https://snapcraft.io/v2raya
8. Binary file and installation package from GitHub releases

See [**v2rayA - Docs**](https://v2raya.org/en/docs/prologue/introduction/)


## Screenshot

<img src="https://i.loli.net/2020/04/19/gt3NqOMiafYbp7L.png" border="0">

## Statement

1. The program does not store any user data in the cloud, and all user data is stored in local.
2. **Do not use this project for illegal purposes.**

## Credits

[hq450/fancyss](https://github.com/hq450/fancyss)

[ToutyRater/v2ray-guide](https://github.com/ToutyRater/v2ray-guide/blob/master/routing/sitedata.md)

[nadoo/glider](https://github.com/nadoo/glider)

[Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat)

[zfl9/ss-tproxy](https://github.com/zfl9/ss-tproxy/blob/master/ss-tproxy)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/v2rayA/v2rayA.svg)](https://starchart.cc/v2rayA/v2rayA)

## License

[![License: AGPL v3-only](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
