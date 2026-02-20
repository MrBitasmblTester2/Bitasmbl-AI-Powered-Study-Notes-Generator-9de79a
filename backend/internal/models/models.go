package models
type MaterialType string
type Material struct{ID string `json:"id"`;Title string `json:"title"`;Content string `json:"content"`;Type MaterialType `json:"type"`}
type Notes struct{ID string `json:"id"`;MaterialID string `json:"materialId"`;Summary string `json:"summary"`;Keywords []string `json:"keywords"`}