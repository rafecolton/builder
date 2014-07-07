package parser

import (
	"github.com/modcloth/docker-builder/parser/uuid"
)

import (
	"path/filepath"

	"github.com/Sirupsen/logrus"
)

/*
Parser is a struct that contains a Builderfile and knows how to parse it both
as raw text and to convert toml to a Builderfile struct.  It also knows how to
tell if the Builderfile is valid (openable) or nat.
*/
type Parser struct {
	filename string
	*logrus.Logger
	uuidGenerator uuid.UUIDGenerator
	top           string
}

var logger *logrus.Logger

/*
NewParser returns an initialized Parser.  Not currently necessary, as no
default values are assigned to a new Parser, but useful to have in case we need
to change this.
*/
func NewParser(filename string, l *logrus.Logger) (*Parser, error) {
	logger = l
	return &Parser{
		Logger:        l,
		filename:      filename,
		uuidGenerator: uuid.NewUUIDGenerator(true),
		top:           filepath.Dir(filename),
	}, nil
}

/*
NextUUID returns the next UUID generated by the parser's uuid generator.  This
will either be a random uuid (normal behavior) or the same uuid every time if
the generator is "seeded" (used for tests)
*/
func (parser *Parser) NextUUID() (string, error) {
	ret, err := parser.uuidGenerator.NextUUID()
	if err != nil {
		return "", err
	}

	return ret, nil
}

/*
SeedUUIDGenerator turns this parser's uuidGenerator into a seeded generator.
All calls to NextUUID() will produce the same uuid after this function is
called and until RandomizeUUIDGenerator() is called.
*/
func (parser *Parser) SeedUUIDGenerator() {
	parser.uuidGenerator = uuid.NewUUIDGenerator(false)
}

/*
RandomizeUUIDGenerator turns this parser's uuidGenerator into a random
generator.  All calls to NextUUID() will produce a random uuid after this
function is called and until SeedUUIDGenerator() is called.
*/
func (parser *Parser) RandomizeUUIDGenerator() {
	parser.uuidGenerator = uuid.NewUUIDGenerator(true)
}
