package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/luispfcanales/confericis-backend/model"
)

type driveService struct {
	apiKey   string
	folderID string
}

func NewDriveService(apiKey, folderID string) *driveService {
	return &driveService{
		apiKey:   apiKey,
		folderID: folderID,
	}
}

type fileList struct {
	Files []*model.DriveFile `json:"files"`
}

func (s *driveService) ListFiles(ctx context.Context, parentID string) ([]*model.DriveFile, error) {
	url := fmt.Sprintf("https://www.googleapis.com/drive/v3/files?q='%s'+in+parents&fields=files(id,name,mimeType,webContentLink,exportLinks)&key=%s",
		parentID,
		s.apiKey,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var files fileList
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	result := make([]*model.DriveFile, 0, len(files.Files))
	for _, file := range files.Files {
		driveFile := &model.DriveFile{
			ID:           file.ID,
			Name:         file.Name,
			MimeType:     file.MimeType,
			DownloadLink: fmt.Sprintf("https://drive.google.com/uc?export=download&id=%s", file.ID),
			WebViewLink:  file.WebViewLink,
		}
		result = append(result, driveFile)
	}

	return result, nil
}

// Implement other DriveService interface methods here
func (s *driveService) GetFileByID(ctx context.Context, fileID string) (*model.DriveFile, error) {
	// TODO: Implement
	return nil, nil
}

func (s *driveService) UploadFile(ctx context.Context, name string, content io.Reader, parentID string) (*model.DriveFile, error) {
	// TODO: Implement
	return nil, nil
}

func (s *driveService) UpdateFile(ctx context.Context, fileID string, content io.Reader) (*model.DriveFile, error) {
	// TODO: Implement
	return nil, nil
}

func (s *driveService) DeleteFile(ctx context.Context, fileID string) error {
	// TODO: Implement
	return nil
}

func (s *driveService) CreateFolder(ctx context.Context, name string, parentID string) (*model.DriveFolder, error) {
	// TODO: Implement
	return nil, nil
}

func (s *driveService) GetFolderByID(ctx context.Context, folderID string) (*model.DriveFolder, error) {
	// TODO: Implement
	return nil, nil
}

func (s *driveService) ListFolders(ctx context.Context, parentID string) ([]*model.DriveFolder, error) {
	// If no parentID provided, use the service's folderID
	if parentID == "" {
		parentID = s.folderID
	}

	url := fmt.Sprintf("https://www.googleapis.com/drive/v3/files?q='%s'+in+parents+and+mimeType='application/vnd.google-apps.folder'&fields=files(id,name,description)&key=%s",
		parentID,
		s.apiKey,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var result struct {
		Files []*model.DriveFolder `json:"files"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	// Set parentID for each folder
	for _, folder := range result.Files {
		folder.ParentID = parentID
	}

	return result.Files, nil
}

func (s *driveService) UpdateFolder(ctx context.Context, folder *model.DriveFolder) error {
	// TODO: Implement
	return nil
}

func (s *driveService) DeleteFolder(ctx context.Context, folderID string) error {
	// TODO: Implement
	return nil
}
