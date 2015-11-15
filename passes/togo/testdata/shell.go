package translated

import "github.com/krageon/php/passes/togo/internal/phpctx"

func Shell(ctx phpctx.PHPContext) {
    ctx.Shell(`ls -al`)
}
