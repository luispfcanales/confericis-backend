package model

type DriveFile struct {
	ID            string `json:"id"`           // Google Drive file ID
	Name          string `json:"name"`         // File name
	MimeType      string `json:"mimeType"`     // File MIME type
	Description   string `json:"description"`   // File description
	WebViewLink   string `json:"webViewLink"`  // URL to view the file in browser
	DownloadLink  string `json:"downloadLink"` // Direct download URL
	ThumbnailLink string `json:"thumbnailLink"`// Thumbnail URL if available
	Size          int64  `json:"size"`         // File size in bytes
	CreatedTime   string `json:"createdTime"`  // Creation timestamp
	ModifiedTime  string `json:"modifiedTime"` // Last modification timestamp
	Parents       []string `json:"parents"`    // Parent folder IDs
}

type DriveFolder struct {
	ID          string `json:"id"`          // Folder ID in Drive
	Name        string `json:"name"`        // Folder name
	Description string `json:"description"` // Folder description
	ParentID    string `json:"parentId"`   // Parent folder ID
}
