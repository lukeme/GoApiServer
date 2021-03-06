package generate

import "fmt"

func getImportStr(name string)  string{
	imp := `import (
	"fmt"
	"github.com/gin-gonic/gin"
	"%s"
	"github.com/wuyan94zl/api/pkg/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)
`
	imp = fmt.Sprintf(imp,name)
	return imp
}
