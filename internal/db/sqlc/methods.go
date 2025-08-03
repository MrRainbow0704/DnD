package sqlc

func (c CharacterScore) GetValue() int64 {
	return c.Value
}

func (c CharacterScore) GetOperation() string {
	return c.Operation.(string)
}

func (c CharacterSpeed) GetValue() int64 {
	return c.Value
}

func (c CharacterSpeed) GetOperation() string {
	return c.Operation.(string)
}
