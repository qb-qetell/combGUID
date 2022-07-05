package combGUID
import "testing"
import "fmt"
import "time"

func Test_ (t *testing.T) {
	for {
		_ba00 := CombGUID_Estb ("syst", 28)
		if _ba00 == nil {
			fmt.Println ("no datat!")
		} else {
			_bb00 := fmt.Sprintf ("%s: %s", _ba00.SmplFrmt (), _ba00.StndFrmt ())
			fmt.Println (_bb00)
		}
		time.Sleep (time.Second * 1)
	}
}
