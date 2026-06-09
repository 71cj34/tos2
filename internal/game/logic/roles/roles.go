package roles
type AttackLevel int
type DefenseLevel int

const (
    None AttackLevel = iota
    Basic
    Powerful
    Unstoppable
)

type Role interface {
    GetAttack() AttackLevel
    GetDefense() DefenseLevel
    NightAction(target ...int) (string, bool)
    DayAction(target ...int) (string, bool)
}

type CommonBase struct {
    Attack  AttackLevel
    Defense DefenseLevel
}