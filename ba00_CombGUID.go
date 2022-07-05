package combGUID
import "fmt"
import "github.com/thanhpk/randstr"
import "math/big"
import "strings"
import "time"

type CombGUID struct {
	iddd string
}
	func CombGUID_Estb (prfx string, sffxLngh uint64) (*CombGUID) {
		if sffxLngh < 6 { return nil }
		combGUID := &CombGUID { iddd: prfx }
		_ba00 := time.Now ().In (time.FixedZone ("UTC!0", 0))
		_bb00 := big.NewInt (_ba00.UnixNano ()).Text (35)
		_ca00 := randstr.String (int (sffxLngh - 6))
		combGUID.iddd = combGUID.iddd + (_bb00 + _ca00) [0:int(sffxLngh)]
		combGUID.iddd = strings.ToLower (combGUID.iddd)
		return combGUID
	}
	func (objc *CombGUID) SmplFrmt () (string) { return objc.iddd }
	func (objc *CombGUID) StndFrmt () (string) {
		if len (objc.iddd) != 32 { return "" }
		prt1 := objc.iddd [ 0: 8]
		prt2 := objc.iddd [ 8:12]
		prt3 := objc.iddd [12:16]
		prt4 := objc.iddd [16:20]
		prt5 := objc.iddd [20:32]
		rslt := fmt.Sprintf ("%s-%s-%s-%s-%s", prt1, prt2, prt3, prt4, prt5)
		return rslt
	}
