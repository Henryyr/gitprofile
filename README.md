# Gitprofile

A CLI to manage multiple git profiles.

## Installation

### macOS (via Homebrew)

```bash
brew install henryyr/tap/gitprofile
```

### Windows (via Scoop)

```bash
# Add the bucket first (only needs to be done once)
scoop bucket add henryyr https://github.com/henryyr/scoop-bucket.git

# Then install the app
scoop install gitprofile
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