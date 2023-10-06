package discovery

type RemoteFileMetadata struct {
	Url       string
	Filename  string
	UpdatedAt string
}

type FileFinder interface {
	FindFiles() []RemoteFileMetadata
}
