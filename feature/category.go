package feature

type categoryType string

const (
	CategoryPermanent categoryType = "permanent"
	CategoryTemporary categoryType = "temporary"
)

func (ct categoryType) String() string {
	return map[categoryType]string{
		CategoryPermanent: "p",
		CategoryTemporary: "t",
	}[ct]
}
