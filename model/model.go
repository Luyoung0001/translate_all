package model

type Words struct {
	ID              string  `gorm:"column:vc_id"`               // 单词id	57067c89a172044907c6698e
	Vocabulary      string  `gorm:"column:vc_vocabulary"`       // 单词	superspecies
	PhoneticUK      string  `gorm:"column:vc_phonetic_uk"`      // uk英音音标	[su:pərsˈpi:ʃi:z]
	PhoneticUS      string  `gorm:"column:vc_phonetic_us"`      // us美音音标	[supɚsˈpiʃiz]
	Frequency       float32 `gorm:"column:vc_frequency"`        // 词频	0.000000
	Difficulty      int     `gorm:"column:vc_difficulty"`       // 难度	1
	AcknowledgeRate float32 `gorm:"column:vc_acknowledge_rate"` // 认识率 0.664
}

type WordTranslation struct {
	ID          uint   `gorm:"primaryKey"`
	Word        string `gorm:"column:word"`
	Translation string `gorm:"column:translation"`
}

type Book struct {
	ID            string  `gorm:"column:bk_id"`              // 单词书id	d645920e395fedad7bbbed0e
	ParentID      string  `gorm:"column:bk_parent_id"`       // 单词书父分类id（无则为0）	6512bd43d9caa6e02c990b0a
	Level         int     `gorm:"column:bk_level"`           // 等级	2
	Order         float32 `gorm:"column:bk_order"`           // 排序	2.000000
	Name          string  `gorm:"column:bk_name"`            // 书名	人教版高中英语1 - 必修
	ItemNum       int     `gorm:"column:bk_item_num"`        // 单词个数	315
	DirectItemNum int     `gorm:"column:bk_direct_item_num"` // 单词个数	315
	Author        string  `gorm:"column:bk_author"`          // 作者	刘道义
	Book          string  `gorm:"column:bk_book"`            // 完整书名	人教版普通高中课程标准实验教科书 英语 1 必修
	Comment       string  `gorm:"column:bk_comment"`         // 描述	黑体：本单元重点词汇和短语；无“△”：课标词汇，要求掌握；有“△”：不要求掌握（会出现大量缩写、人名、地名和短语，请选背）。
	Organization  string  `gorm:"column:bk_orgnization"`     // 组织	人民教育出版社 课程教材研究所；英语课程教材研究开发中心
	Publisher     string  `gorm:"column:bk_publisher"`       // 出版社	人民教育出版社
	Version       string  `gorm:"column:bk_version"`         // 版本	2007年1月第2版
	Flag          string  `gorm:"column:bk_flag"`            // 标记	默认：152;黑体：97;前△：66
}

type RelationBookWord struct {
	ID     string `gorm:"column:bv_id"`      // 关系id	58450c828958a37d5c10f763
	BookID string `gorm:"column:bv_book_id"` // 单词书id	d645920e395fedad7bbbed0e
	VocID  string `gorm:"column:bv_voc_id"`  // 单词id	57067b9ca172044907c615d7
	Flag   string `gorm:"column:bv_flag"`    // 分组	4
	Tag    string `gorm:"column:bv_tag"`     // 分组名	Unit 1
	Order  int    `gorm:"column:bv_order"`   // 排序	1
}
