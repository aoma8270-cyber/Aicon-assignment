package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewItem(t *testing.T) {
	tests := []struct {
		name          string
		itemName      string
		category      string
		brand         string
		purchasePrice int
		purchaseDate  string
		wantErr       bool
		expectedErr   string
	}{
		{
			name:          "正常系: 有効なアイテム作成",
			itemName:      "ロレックス デイトナ",
			category:      "時計",
			brand:         "ROLEX",
			purchasePrice: 1500000,
			purchaseDate:  "2023-01-15",
			wantErr:       false,
		},
		{
			name:          "異常系: 名前が空",
			itemName:      "",
			category:      "時計",
			brand:         "ROLEX",
			purchasePrice: 1500000,
			purchaseDate:  "2023-01-15",
			wantErr:       true,
			expectedErr:   "name is required",
		},
		{
			name:          "異常系: 名前が100文字超過",
			itemName:      "ロレックス デイトナ 16520 18K イエローゴールド ブラック文字盤 自動巻き クロノグラフ メンズ 腕時計 1988年製 ヴィンテージ 希少 コレクション アイテム",
			category:      "時計",
			brand:         "ROLEX",
			purchasePrice: 1500000,
			purchaseDate:  "2023-01-15",
			wantErr:       true,
			expectedErr:   "name must be 100 characters or less",
		},
		{
			name:          "異常系: カテゴリーが空",
			itemName:      "ロレックス デイトナ",
			category:      "",
			brand:         "ROLEX",
			purchasePrice: 1500000,
			purchaseDate:  "2023-01-15",
			wantErr:       true,
			expectedErr:   "category is required",
		},
		{
			name:          "異常系: 無効なカテゴリー",
			itemName:      "ロレックス デイトナ",
			category:      "無効なカテゴリー",
			brand:         "ROLEX",
			purchasePrice: 1500000,
			purchaseDate:  "2023-01-15",
			wantErr:       true,
			expectedErr:   "category must be one of: 時計, バッグ, ジュエリー, 靴, その他",
		},
		{
			name:          "異常系: ブランドが空",
			itemName:      "ロレックス デイトナ",
			category:      "時計",
			brand:         "",
			purchasePrice: 1500000,
			purchaseDate:  "2023-01-15",
			wantErr:       true,
			expectedErr:   "brand is required",
		},
		{
			name:          "異常系: ブランドが100文字超過",
			itemName:      "ロレックス デイトナ",
			category:      "時計",
			brand:         "ROLEX SA Geneva Switzerland Official Authorized Dealer Store Premium Collection Limited Edition Special",
			purchasePrice: 1500000,
			purchaseDate:  "2023-01-15",
			wantErr:       true,
			expectedErr:   "brand must be 100 characters or less",
		},
		{
			name:          "異常系: 購入価格が負の値",
			itemName:      "ロレックス デイトナ",
			category:      "時計",
			brand:         "ROLEX",
			purchasePrice: -1,
			purchaseDate:  "2023-01-15",
			wantErr:       true,
			expectedErr:   "purchase_price must be 0 or greater",
		},
		{
			name:          "異常系: 購入日が空",
			itemName:      "ロレックス デイトナ",
			category:      "時計",
			brand:         "ROLEX",
			purchasePrice: 1500000,
			purchaseDate:  "",
			wantErr:       true,
			expectedErr:   "purchase_date is required",
		},
		{
			name:          "異常系: 無効な日付形式",
			itemName:      "ロレックス デイトナ",
			category:      "時計",
			brand:         "ROLEX",
			purchasePrice: 1500000,
			purchaseDate:  "2023/01/15",
			wantErr:       true,
			expectedErr:   "purchase_date must be in YYYY-MM-DD format",
		},
		{
			name:          "正常系: 購入価格が0",
			itemName:      "ギフト品",
			category:      "その他",
			brand:         "不明",
			purchasePrice: 0,
			purchaseDate:  "2023-01-15",
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := NewItem(tt.itemName, tt.category, tt.brand, tt.purchasePrice, tt.purchaseDate)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				assert.Nil(t, item)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, item)

			// フィールドの値をチェック
			assert.Equal(t, tt.itemName, item.Name)
			assert.Equal(t, tt.category, item.Category)
			assert.Equal(t, tt.brand, item.Brand)
			assert.Equal(t, tt.purchasePrice, item.PurchasePrice)
			assert.Equal(t, tt.purchaseDate, item.PurchaseDate)

			// CreatedAt と UpdatedAt がセットされているかチェック
			assert.False(t, item.CreatedAt.IsZero())
			assert.False(t, item.UpdatedAt.IsZero())
		})
	}
}

func TestItem_Update(t *testing.T) {
	// 初期アイテムを作成
	item, err := NewItem("初期アイテム", "時計", "初期ブランド", 100000, "2023-01-01")
	require.NoError(t, err)

	originalUpdatedAt := item.UpdatedAt
	time.Sleep(1 * time.Millisecond) // UpdatedAt の変更を確認するため

	tests := []struct {
		name        string
		newName     string
		newCategory string
		newBrand    string
		newPrice    int
		newDate     string
		wantErr     bool
		expectedErr string
	}{
		{
			name:        "正常系: 全フィールド更新",
			newName:     "更新されたアイテム",
			newCategory: "バッグ",
			newBrand:    "更新されたブランド",
			newPrice:    200000,
			newDate:     "2023-12-31",
			wantErr:     false,
		},
		{
			name:        "異常系: 無効なカテゴリー",
			newName:     "更新されたアイテム",
			newCategory: "無効なカテゴリー",
			newBrand:    "更新されたブランド",
			newPrice:    200000,
			newDate:     "2023-12-31",
			wantErr:     true,
			expectedErr: "category must be one of: 時計, バッグ, ジュエリー, 靴, その他",
		},
		{
			name:        "異常系: 負の価格",
			newName:     "更新されたアイテム",
			newCategory: "バッグ",
			newBrand:    "更新されたブランド",
			newPrice:    -1,
			newDate:     "2023-12-31",
			wantErr:     true,
			expectedErr: "purchase_price must be 0 or greater",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := item.Update(tt.newName, tt.newCategory, tt.newBrand, tt.newPrice, tt.newDate)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				return
			}

			assert.NoError(t, err)

			// 更新後の値をチェック
			assert.Equal(t, tt.newName, item.Name)
			assert.Equal(t, tt.newCategory, item.Category)
			assert.Equal(t, tt.newBrand, item.Brand)
			assert.Equal(t, tt.newPrice, item.PurchasePrice)
			assert.Equal(t, tt.newDate, item.PurchaseDate)

			// UpdatedAt が更新されているかチェック
			assert.True(t, item.UpdatedAt.After(originalUpdatedAt))
		})
	}
}

func TestItem_Validate(t *testing.T) {
	tests := []struct {
		name        string
		item        *Item
		wantErr     bool
		expectedErr string
	}{
		{
			name: "正常系: 有効なアイテム",
			item: &Item{
				Name:          "ロレックス デイトナ",
				Category:      "時計",
				Brand:         "ROLEX",
				PurchasePrice: 1500000,
				PurchaseDate:  "2023-01-15",
			},
			wantErr: false,
		},
		{
			name: "異常系: 複数のバリデーションエラー",
			item: &Item{
				Name:          "",
				Category:      "",
				Brand:         "",
				PurchasePrice: -1,
				PurchaseDate:  "",
			},
			wantErr:     true,
			expectedErr: "name is required, category is required, brand is required, purchase_price must be 0 or greater, purchase_date is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.item.Validate()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestIsValidCategory(t *testing.T) {
	tests := []struct {
		name     string
		category string
		want     bool
	}{
		{"有効なカテゴリー: 時計", "時計", true},
		{"有効なカテゴリー: バッグ", "バッグ", true},
		{"有効なカテゴリー: ジュエリー", "ジュエリー", true},
		{"有効なカテゴリー: 靴", "靴", true},
		{"有効なカテゴリー: その他", "その他", true},
		{"無効なカテゴリー: 衣服", "衣服", false},
		{"無効なカテゴリー: 空文字", "", false},
		{"無効なカテゴリー: 英語", "watch", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidCategory(tt.category)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsValidDateFormat(t *testing.T) {
	tests := []struct {
		name    string
		dateStr string
		want    bool
	}{
		{"有効な日付: 2023-01-15", "2023-01-15", true},
		{"有効な日付: 2023-12-31", "2023-12-31", true},
		{"有効な日付: 2024-02-29", "2024-02-29", true}, // うるう年
		{"無効な日付: 2023/01/15", "2023/01/15", false},
		{"無効な日付: 2023-1-15", "2023-1-15", false},
		{"無効な日付: 15-01-2023", "15-01-2023", false},
		{"無効な日付: 2023-13-01", "2023-13-01", false},
		{"無効な日付: 2023-02-30", "2023-02-30", false},
		{"無効な日付: 空文字", "", false},
		{"無効な日付: 無効な形式", "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidDateFormat(tt.dateStr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetValidCategories(t *testing.T) {
	categories := GetValidCategories()
	expected := []string{"時計", "バッグ", "ジュエリー", "靴", "その他"}

	assert.Equal(t, expected, categories)
	assert.Len(t, categories, 5)
}

func TestItem_ApplyPatch(t *testing.T) {
	// ヘルパー関数
	strPtr := func(s string) *string { return &s }
	intPtr := func(i int) *int { return &i }

	tests := []struct {
		name          string
		patchName     *string
		patchBrand    *string
		patchPrice    *int
		wantErr       bool
		expectedErr   string
		expectedName  string
		expectedBrand string
		expectedPrice int
	}{
		{
			name:          "正常系: nameのみ更新",
			patchName:     strPtr("更新された名前"),
			patchBrand:    nil,
			patchPrice:    nil,
			wantErr:       false,
			expectedName:  "更新された名前",
			expectedBrand: "ROLEX",
			expectedPrice: 1000000,
		},
		{
			name:          "正常系: brandのみ更新",
			patchName:     nil,
			patchBrand:    strPtr("OMEGA"),
			patchPrice:    nil,
			wantErr:       false,
			expectedName:  "初期アイテム",
			expectedBrand: "OMEGA",
			expectedPrice: 1000000,
		},
		{
			name:          "正常系: purchase_priceのみ更新",
			patchName:     nil,
			patchBrand:    nil,
			patchPrice:    intPtr(2000000),
			wantErr:       false,
			expectedName:  "初期アイテム",
			expectedBrand: "ROLEX",
			expectedPrice: 2000000,
		},
		{
			name:          "正常系: 複数フィールド同時更新",
			patchName:     strPtr("新しい名前"),
			patchBrand:    strPtr("CARTIER"),
			patchPrice:    intPtr(500000),
			wantErr:       false,
			expectedName:  "新しい名前",
			expectedBrand: "CARTIER",
			expectedPrice: 500000,
		},
		{
			name:          "正常系: nameの前後空白がトリムされる",
			patchName:     strPtr("  スペース付き名前  "),
			patchBrand:    nil,
			patchPrice:    nil,
			wantErr:       false,
			expectedName:  "スペース付き名前",
			expectedBrand: "ROLEX",
			expectedPrice: 1000000,
		},
		{
			name:        "異常系: nameを空文字に更新",
			patchName:   strPtr(""),
			patchBrand:  nil,
			patchPrice:  nil,
			wantErr:     true,
			expectedErr: "name is required",
		},
		{
			name:        "異常系: brandを空文字に更新",
			patchName:   nil,
			patchBrand:  strPtr(""),
			patchPrice:  nil,
			wantErr:     true,
			expectedErr: "brand is required",
		},
		{
			name:        "異常系: purchase_priceを負の値に更新",
			patchName:   nil,
			patchBrand:  nil,
			patchPrice:  intPtr(-100),
			wantErr:     true,
			expectedErr: "purchase_price must be 0 or greater",
		},
		{
			name:          "正常系: purchase_priceを0に更新",
			patchName:     nil,
			patchBrand:    nil,
			patchPrice:    intPtr(0),
			wantErr:       false,
			expectedName:  "初期アイテム",
			expectedBrand: "ROLEX",
			expectedPrice: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 各テストケースで新しいアイテムを作成
			item, err := NewItem("初期アイテム", "時計", "ROLEX", 1000000, "2023-01-01")
			require.NoError(t, err)

			originalUpdatedAt := item.UpdatedAt
			originalCategory := item.Category
			originalPurchaseDate := item.PurchaseDate
			originalCreatedAt := item.CreatedAt
			time.Sleep(1 * time.Millisecond) // UpdatedAtの変更を確認するため

			err = item.ApplyPatch(tt.patchName, tt.patchBrand, tt.patchPrice)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				return
			}

			require.NoError(t, err)

			// 更新対象フィールドの値を確認
			assert.Equal(t, tt.expectedName, item.Name)
			assert.Equal(t, tt.expectedBrand, item.Brand)
			assert.Equal(t, tt.expectedPrice, item.PurchasePrice)

			// 不変フィールドが変更されていないことを確認
			assert.Equal(t, originalCategory, item.Category)
			assert.Equal(t, originalPurchaseDate, item.PurchaseDate)
			assert.Equal(t, originalCreatedAt, item.CreatedAt)

			// UpdatedAtが更新されていることを確認
			assert.True(t, item.UpdatedAt.After(originalUpdatedAt))
		})
	}
}
