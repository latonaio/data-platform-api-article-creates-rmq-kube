package requests

type Like struct {
	Article					int		`json:"Article"`
	BusinessPartner			int		`json:"BusinessPartner"`
	Like					*bool	`json:"Like"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
}
