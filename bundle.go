package main

// Turn naming ambiguities of system architectures into the golang naming scheme.
func normalizeArch(arch string) string {
	switch arch {
	case "aarch64":
		return "arm64"
	case "x86_64":
		return "amd64"
	default:
		return arch
	}
}

// Turn golang naming scheme into platform specific scheme.
func DenormalizeArch(arch string) string {
	switch arch {
	case "arm64":
		return "aarch64"
	case "amd64":
		return "x86_64"
	default:
		return arch
	}
}

func Bundle(name string, platform string, arch string) {
	if name == "" || name == "." {
		Fatal("Please supply an app name.\nE.g. djs bundle MyCoolApp")
	}

	switch platform {
	case "darwin":
		BundleDarwin(name)
	case "linux":
		BundleLinux(name, normalizeArch(arch))
	default:
		Fatalf("Cannot bundle desktop.js applications for %s right now.\nPlease contribute if you want to support more platforms.", platform)
	}
}
