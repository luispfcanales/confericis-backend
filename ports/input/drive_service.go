package input

import (
	"context"

	"github.com/luispfcanales/confericis-backend/model"
)

type DriveService interface {
	// File operations
	// UploadFile(ctx context.Context, name string, content io.Reader, parentID string) (*model.DriveFile, error)
	// GetFileByID(ctx context.Context, fileID string) (*model.DriveFile, error)
	ListFiles(ctx context.Context, parentID string) ([]*model.DriveFile, error)
	// UpdateFile(ctx context.Context, fileID string, content io.Reader) (*model.DriveFile, error)
	// DeleteFile(ctx context.Context, fileID string) error

	// Folder operations
	// CreateFolder(ctx context.Context, name string, parentID string) (*model.DriveFolder, error)
	// GetFolderByID(ctx context.Context, folderID string) (*model.DriveFolder, error)
	ListFolders(ctx context.Context, parentID string) ([]*model.DriveFolder, error)
	// UpdateFolder(ctx context.Context, folder *model.DriveFolder) error
	// DeleteFolder(ctx context.Context, folderID string) error
}
