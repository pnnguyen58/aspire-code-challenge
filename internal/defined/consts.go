package defined

const (
	SCHEDULED State = "SCHEDULED"
	APPROVED  State = "APPROVED"
	PENDING   State = "PENDING"
	CANCELED  State = "PENDING"
)

const (
	DAILY    RepaymentType = "daily"
	WEEKLY   RepaymentType = "weekly"
	MONTHLY  RepaymentType = "monthly"
	ANNUALLY RepaymentType = "annually"
)

const (
	PRECISION  = 2
	BACTH_SIZE = 10
)
