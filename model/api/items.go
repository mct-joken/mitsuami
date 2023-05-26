package api

// CreateItemRequestJSON 備品の追加(リクエスト)
type CreateItemRequestJSON struct {
	ID          string `json:"ID"`
	Type        int    `json:"type"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// CreateItemResponseJSON 備品の追加(レスポンス)
type CreateItemResponseJSON = GetItemResponseJSON

// GetItemResponseJSON 備品の取得(レスポンス)
type GetItemResponseJSON struct {
	ID          string `json:"id"`
	Type        int    `json:"type"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Status      struct {
		Status  int    `json:"status"`
		User    string `json:"user,omitempty"`
		DueDate string `json:"dueDate,omitempty"`
	} `json:"status"`
}

// StartUsingItemResponseJSON 貸出開始(レスポンス
type StartUsingItemResponseJSON struct {
	Status string `json:"status"`
}

// EndUsingItemResponseJSON 返却完了(レスポンス)
type EndUsingItemResponseJSON = StartUsingItemResponseJSON

// GetItemListResponseJSON 備品リストの取得(レスポンス)
type GetItemListResponseJSON = []GetItemResponseJSON
