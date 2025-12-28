package plot

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/soluchok/go-bayesopt/gp"
)

// SaveAll saves all dimension graphs of the gaussian process to a temp
// directory and prints their file names.
func SaveAll(gp *gp.GP) (string, error) {
	dir, err := ioutil.TempDir("", "gp-plots")
	if err != nil {
		return "", err
	}
	dims := gp.Dims()
	for i := 0; i < dims; i++ {
		name := fmt.Sprintf("%d.svg", i)
		fpath := path.Join(dir, name)
		f, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return "", err
		}
		defer f.Close()
		if err := GP(gp, f, i); err != nil {
			return "", err
		}
		f.Close()
	}
	return dir, nil
}

// GP renders a plot of the gaussian process for the specified dimension.
func GP(gp *gp.GP, w io.Writer, dim int) error {
	return errors.New("not implemented")
}
