package storage

import (
	"errors"
	"io"

	"github.com/sliverarmory/external-armory/consts"
	"github.com/sliverarmory/external-armory/util"
)

var (
	ErrStorageNotInitialized = errors.New("the storage provider has not been initialized")
	ErrDoesNotExist          = errors.New("the file does not exist")
)

type PackageEntry struct {
	Name         string
	CommandName  string
	ManifestData []byte
}

/*
This is the interface for the armory's storage backends. It features discrete functions for reading and writing the different
kinds of files that we would need from the storage backend. This shifts the need to remember the paths for each type of file
onto the storage provider as opposed to having a single read or write file function where a path or other specifier would have
to be supplied by the caller.
*/
type StorageProvider interface {
	// New sets up a new instance of the storage provider with a given base path
	New(string, bool) error
	// Returns whether the directories had to be created on the storage provider (useful to know whether initial setup needs to be run)
	IsNew() bool
	// Returns the paths configured for this provider and an error if the provider is not initialized
	Paths() (*StoragePaths, error)
	// Indicates whether the storage provider can notify the main process of a change in the package directories
	AutoRefreshEnabled() (bool, error)
	// Returns the channels for file events and errors. An error is returned if the provider is not initialized or there was an error in setting up auto refresh
	AutoRefreshChannels() (chan string, chan error, error)
	// Removes the root directory and all of its subdirectories
	Destroy() error
	// Returns whether the storage provider has been initialized successfully
	Initialized() bool
	// Returns the root directory (base path) of the storage provider
	BasePath() string
	// Checks for the existance of a file in the storage provider
	CheckFile(string) error
	// Sets a different path for the config file than the default
	SetConfigPath(string) error
	// Returns the raw contents of the configuration file
	ReadConfig() ([]byte, error)
	// Stores the provided data in the configuration file
	WriteConfig([]byte) error
	// Reads the package signing key from storage
	ReadPackageSigningKey() ([]byte, error)
	// Writes the package signing key to storage
	WritePackageSigningKey([]byte) error
	// Reads the TLS certificate key file from storage
	ReadTLSCertificateKey() ([]byte, error)
	// Writes the TLS certificate key file to storage
	WriteTLSCertificateKey([]byte) error
	// Reads the TLS certificate file from storage
	ReadTLSCertificateCrt() ([]byte, error)
	// Writes the TLS certificate file to storage
	WriteTLSCertificateCrt([]byte) error
	// Reads the bundle file from storage
	ReadBundleFile() ([]byte, error)
	// Writes the bundle file to storage
	WriteBundleFile([]byte) error
	// Checks for the existence of a package with a given name and returns its package type and any errors encountered
	CheckPackage(string) (consts.PackageType, error)
	// Returns the archive for a package with the given name
	ReadPackage(string) ([]byte, error)
	// Writes the package with the given name to storage
	WritePackage(string, []byte) error
	// Writes the package with the given file name to storage
	WritePackageWithFileName(string, []byte) error
	// Remove a package with a given name from storage. Removing a package does not remove its signature, so be sure to remove its signature too.
	RemovePackage(string) error
	// Return a list of packages as a map of the command name to a PackageEntry (name, command name, manifest)
	ListPackages(consts.PackageType) (map[string]PackageEntry, []error)
	// Get the signature for a given package (this package does not sign the package)
	ReadPackageSignature(string) ([]byte, error)
	// Write the signature for a given package (this function does not sign the package)
	WritePackageSignature(string, []byte) error
	// Remove the package signature from storage for a given package
	RemovePackageSignature(string) error
	// Return the package index (this function does not generate the index)
	ReadIndex() ([]byte, error)
	// Write the package index to storage
	WriteIndex([]byte) error
	// Read the package index signature (this function does not sign the index)
	ReadIndexSignature() ([]byte, error)
	// Write the package index signature to storage (this function does not sign the index)
	WriteIndexSignature([]byte) error
	// Return the named logging backend (io.Writer)
	GetLogger(string) (io.Writer, error)
	// For Vault
	// Return the custom CA PEM file for the configured Vault
	ReadVaultCA() ([]byte, error)
	// Store the custom CA PEM file for the configured Vault
	WriteVaultCA([]byte) error
}

// Common functions
func derivePackageTypeFromArchive(archiveData []byte) consts.PackageType {
	// Try to find the alias manifest
	manifest, err := util.ReadFileFromTarGzMemory(archiveData, consts.AliasArchiveManifestFilePath)
	if err != nil || len(manifest) == 0 {
		// Then this might be an extension
		manifest, err = util.ReadFileFromTarGzMemory(archiveData, consts.ExtensionArchiveManifestFilePath)

		if err != nil || len(manifest) == 0 {
			return consts.UnknownPackageType
		} else {
			return consts.ExtensionPackageType
		}
	} else {
		return consts.AliasPackageType
	}
}

// An object for the paths - this is common for all StorageProviders
type StoragePaths struct {
	Aliases           string
	Extensions        string
	PackageSignatures string
	Certificates      string
	Bundles           string
	Config            string
	Index             string
	IndexSignature    string
	PackageSigningKey string
	CertificateKey    string
	CertificateCrt    string
	VaultCAPEM        string
}

func (sp *StoragePaths) Directories() map[string]string {
	return map[string]string{
		"aliases":            sp.Aliases,
		"extensions":         sp.Extensions,
		"package signatures": sp.PackageSignatures,
		"certificates":       sp.Certificates,
	}
}

func (sp *StoragePaths) Files() map[string]string {
	return map[string]string{
		"bundle information":     sp.Bundles,
		"configuration":          sp.Config,
		"index":                  sp.Index,
		"index signature":        sp.IndexSignature,
		"package signing key":    sp.PackageSigningKey,
		"TLS certificate key":    sp.CertificateKey,
		"signed TLS certificate": sp.CertificateCrt,
		"Vault CA PEM file":      sp.VaultCAPEM,
	}
}