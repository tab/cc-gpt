package changelog

import (
	"fmt"

	"cmt/internal/commands"
	"cmt/internal/errors"
	"cmt/internal/utils"
)

type Command struct {
	Options commands.GenerateOptions
}

func NewCommand(options commands.GenerateOptions) *Command {
	return &Command{
		Options: options,
	}
}

func (c *Command) Generate() error {
	loader := utils.NewLoader()
	loader.Start()

	commits, err := c.Options.Client.Log(c.Options.Ctx, c.Options.Args)
	if err != nil {
		loader.Stop()
		errors.HandleGitLogError(err)
		return err
	}

	changelog, err := c.Options.Model.FetchChangelog(c.Options.Ctx, commits)
	loader.Stop()

	if err != nil {
		return fmt.Errorf("error requesting changelog: %w", err)
	}

	fmt.Printf("💬 Changelog: \n\n%s\n", changelog)
	return nil
}
