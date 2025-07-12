# mlmym 🌐

[![Go Version](https://img.shields.io/badge/Go-1.23.8%2B-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](LICENSE)

**A familiar desktop experience for [Lemmy](https://join-lemmy.org) that brings back the classic forum feel.**

Transform any Lemmy instance into a desktop-style interface reminiscent of old Reddit, complete with familiar navigation and layout.

![screenshot](https://raw.githubusercontent.com/Fedihosting-Foundation-Forks/mlmym/main/screenshot1.png?raw=true)

---

**📋 About This Fork**

This project was originally created by [rystaf](https://github.com/rystaf/mlmym). As the original project is no longer maintained, this fork is receiving limited life support to support the instance at [old.lemmy.world](https://old.lemmy.world/).

## ✨ Features

- 🖥️ **Classic desktop interface** - Familiar layout for Reddit refugees
- 🌍 **Multi-instance support** - Browse any Lemmy instance through one interface
- 🎯 **Single-instance mode** - Dedicated frontend for your preferred instance
- 🔄 **Hot reload development** - Fast iteration during development
- 📱 **Responsive design** - Works on desktop and mobile
- 🎨 **Customizable settings** - Dark mode, thumbnails, and more

## 🚀 Quick Start

```bash
# Clone and run immediately
git clone https://github.com/Fedihosting-Foundation-Forks/mlmym.git
cd mlmym
go run . --addr :8080
```

Then visit `http://localhost:8080` and enter any Lemmy instance URL to start browsing!

## 📦 Deployment

### 🐳 Docker (Recommended)

```bash
docker run -it -p "8080:8080" ghcr.io/fedihosting-foundation-forks/mlmym:latest
```

### 🔧 Running without Docker

**Prerequisites:** Go 1.23.8 or later

For build instructions and development setup, see [docs/development.md](docs/development.md).

## ⚙️ Configuration

### 🎯 Deployment Modes

mlmym supports two deployment modes:

#### 🌍 Multi-Instance Mode (Default)
Run without any configuration to browse **any** Lemmy instance through the same interface.

```bash
# Docker
docker run -it -p "8080:8080" ghcr.io/fedihosting-foundation-forks/mlmym:latest

# Direct
./mlmym --addr :8080
```

Users can enter any Lemmy instance URL (e.g., `lemmy.world`, `lemmy.ml`) and browse seamlessly.

#### 🎯 Single Instance Mode  
Set `LEMMY_DOMAIN` to create a dedicated frontend for one specific Lemmy instance.

```bash
# Docker
docker run -it -e LEMMY_DOMAIN='lemmy.world' -p "8080:8080" ghcr.io/fedihosting-foundation-forks/mlmym:latest

# Direct
LEMMY_DOMAIN='lemmy.world' ./mlmym --addr :8080
```

This mode provides a dedicated frontend for a specific Lemmy instance.

### 🎨 Default User Settings

Customize the default experience for all users by setting these environment variables:

| Environment Variable   | Default | Description |
|------------------------|---------|-------------|
| `LISTING`              | All     | Default post listing type |
| `SORT`                 | Hot     | Default post sorting |
| `COMMENT_SORT`         | Hot     | Default comment sorting |
| `DARK`                 | false   | Enable dark mode by default |
| `HIDE_THUMBNAILS`      | false   | Hide post thumbnails by default |
| `COLLAPSE_MEDIA`       | false   | Collapse media content by default |
| `LINKS_IN_NEW_WINDOW`  | false   | Open links in new windows |
| `COMMENTS_IN_NEW_WINDOW` | false | Open comments in new windows |

**Usage:** Set any environment variable to any value to enable it. Leave blank to disable.

```bash
# Example: Enable dark mode and hide thumbnails by default
DARK=1 HIDE_THUMBNAILS=1 ./mlmym --addr :8080
```
