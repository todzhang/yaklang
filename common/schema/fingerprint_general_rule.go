package schema

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type CPE struct {
	Part     string `yaml:"part,omitempty" json:"part"`
	Vendor   string `yaml:"vendor,omitempty" json:"vendor"`
	Product  string `yaml:"product,omitempty" json:"product"`
	Version  string `yaml:"version,omitempty" json:"version"`
	Update   string `yaml:"update,omitempty" json:"update"`
	Edition  string `yaml:"edition,omitempty" json:"edition"`
	Language string `yaml:"language,omitempty" json:"language"`
}

func (c *CPE) Init() {
	if c.Part == "" {
		c.Part = "a"
	}

	setWildstart := func(raw *string) {
		if *raw == "" {
			*raw = "*"
		}
	}

	setWildstart(&c.Vendor)
	setWildstart(&c.Product)
	setWildstart(&c.Version)
	setWildstart(&c.Update)
	setWildstart(&c.Edition)
	setWildstart(&c.Language)
}

func (c *CPE) String() string {
	c.Init()
	raw := fmt.Sprintf("cpe:/%s:%s:%s:%s:%s:%s:%s", c.Part, c.Vendor, c.Product, c.Version, c.Update, c.Edition, c.Language)
	raw = strings.ReplaceAll(raw, " ", "_")
	raw = strings.ToLower(raw)
	return raw
}

type GeneralRule struct {
	gorm.Model
	*CPE
	RuleName        string `gorm:"unique_index"`
	WebPath         string
	ExtInfo         string
	MatchExpression string
}

func (g *GeneralRule) String() string {
	items := []string{fmt.Sprintf("name:%s", g.RuleName)}
	cpeStr := g.CPE.String()
	items = append(items, fmt.Sprintf("cpe:%s", cpeStr))

	if g.WebPath != "" {
		items = append(items, fmt.Sprintf("webpath:%s", g.WebPath))
	}
	if g.ExtInfo != "" {
		items = append(items, fmt.Sprintf("info:%s", g.ExtInfo))
	}
	items = append(items, fmt.Sprintf("rule:%s", g.MatchExpression))
	return strings.Join(items, " ")
}