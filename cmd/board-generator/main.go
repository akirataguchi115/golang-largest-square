package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"math/rand"
	"os"
	"time"
)

type argT struct {
	cli.Helper

	Width             int    `cli:"*W,width" usage:"width of the board, has to be greater than 0"`
	Height            int    `cli:"*H,height" usage:"height of the board, has to be greater than 0"`
	Block             string `cli:"b,block" usage:"block character" dft:"O"`
	Space             string `cli:"s,space" usage:"space character" dft:"."`
	Draw              string `cli:"d,draw" usage:"character to draw the square" dft:"*"`
	BlockDistribution int    `cli:"p,block-distribution" usage:"Block probability distribution" dft:"5"`
}

func (argv *argT) Validate(ctx *cli.Context) error {
	if argv.Width < 1 {
		return fmt.Errorf("width %d out of range (>0)", argv.Width)
	}

	if argv.Height < 1 {
		return fmt.Errorf("height %d out of range (>0)", argv.Height)
	}

	if argv.BlockDistribution > 100 {
		return fmt.Errorf("block distribution %d out of range (>100)", argv.BlockDistribution)
	}

	return nil
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)

		ctx.String("%d%s%s%s\n", argv.Height, argv.Space, argv.Block, argv.Draw)

		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		board := ""

		for y := 0; y < argv.Height; y++ {
			line := ""

			for x := 0; x < argv.Width; x++ {
				if r1.Intn(100) < argv.BlockDistribution {
					line += argv.Block
				} else {
					line += argv.Space
				}
			}

			ctx.String("%s\n", line)

			board += line
		}

		// ctx.String("%x", sha256.Sum256([]byte(board)))

		return nil
	}))
}
