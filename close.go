package gitdb

func (g *GitDB) Close() error {
	return g.DB.Close()
}
