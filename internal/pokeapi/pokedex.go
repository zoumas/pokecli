package pokeapi

func (c *Client) AddToPokedex(name string, pokemon *PokemonResponse) {
	c.Pokedex[name] = pokemon
}
