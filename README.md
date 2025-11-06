# Gitprofile

A CLI to manage multiple git profiles.

## Installation

To install, first clone this repository, then use the local installation files for your package manager.

```bash
git clone https://github.com/henryyr/gitprofile.git
cd gitprofile
```

### macOS (via Homebrew)

```bash
# From within the cloned repository
brew install --formula ./Formula/gitprofile.rb
```

### Windows (via Scoop)

```bash
# From within the cloned repository
scoop install ./scoop/gitprofile.json
```

### Linux (via .deb/.rpm)

Download the appropriate `.deb` or `.rpm` package from the [latest release page](https://github.com/henryyr/gitprofile/releases/latest) and install with your local package manager.

**For Debian/Ubuntu:**

```bash
sudo dpkg -i gitprofile_*.deb
```

**For Fedora/CentOS:**

```bash
sudo rpm -i gitprofile_*.rpm
```

## Usage

Once installed, you can use the `gitprofile` command globally.

```bash
gitprofile --help
```

### Add a new profile

```bash
gitprofile add
```

### List all profiles

```bash
gitprofile list
```

### Switch to a different profile

```bash
gitprofile use <username>
```

### Remove a profile

```bash
gitprofile remove <username>
```

### Show the current active profile

```bash
gitprofile current
```
