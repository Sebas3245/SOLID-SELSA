package main

// Bad structure
type BadDropbox interface {
	storeFile(name string) string
	getFile(name string) string
	listServers(region string) []string
}

type BadGoogleDrive interface {
	storeFile(name string) string
	getFile(name string) string
}

type BadCloudProvider interface {
	storeFile(name string) string
	getFile(name string) string
	listServers(region string) []string
}

// Correct structure
type Dropbox interface {
	storeFile(name string) string
	getFile(name string) string
	listServers(region string) []string
}

type GoogleDrive interface {
	storeFile(name string) string
	getFile(name string) string
}

type CloudHostingProvider interface {
	listServers(region string) []string
}

type CloudStorageProvider interface {
	storeFile(name string) string
	getFile(name string) string
}
