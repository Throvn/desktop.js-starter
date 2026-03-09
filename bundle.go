package main

func Bundle(name string, platform string) {
	if name == "" || name == "." {
		Fatal("Please supply an app name.\nE.g. djs bundle MyCoolApp")
	}

	switch (platform) {
	case "darwin":
		BundleDarwin(name)
	case "linux":
		BundleLinux(name)
	default:
		Fatalf("Cannot bundle desktop.js applications for "+platform+" right now.\nPlease contribute if you want to support more platforms.")
	}
}