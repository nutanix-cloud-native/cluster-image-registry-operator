package openshift

import (
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	roleBindingCloudCredsSecretReaderFileName = "rolebinding-cloud-creds-secret-reader.yaml.template"
)

var _ asset.WritableAsset = (*RoleBindingCloudCredsSecretReader)(nil)

// RoleBindingCloudCredsSecretReader is the variable to represent contents of corresponding file
type RoleBindingCloudCredsSecretReader struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *RoleBindingCloudCredsSecretReader) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *RoleBindingCloudCredsSecretReader) Name() string {
	return "RolebindingCloudCredsSecretReader"
}

// Generate generates the actual files by this asset
func (t *RoleBindingCloudCredsSecretReader) Generate(parents asset.Parents) error {
	fileName := roleBindingCloudCredsSecretReaderFileName
	data, err := content.GetOpenshiftTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *RoleBindingCloudCredsSecretReader) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *RoleBindingCloudCredsSecretReader) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, roleBindingCloudCredsSecretReaderFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
