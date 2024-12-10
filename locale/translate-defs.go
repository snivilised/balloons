package locale

// CLIENT-TODO: Should be updated to use url of the implementing project,
// so should not be left as balloons. (this should be set by auto-check)
const AstrolibSourceID = "github.com/snivilised/balloons"

type balloonsTemplData struct{}

func (td balloonsTemplData) SourceID() string {
	return AstrolibSourceID
}
