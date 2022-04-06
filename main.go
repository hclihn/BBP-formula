package main

import (
	"fmt"
  "math/big"
  "runtime"
  "time"
  "math"
)

// References:
// https://en.wikipedia.org/wiki/Bailey%E2%80%93Borwein%E2%80%93Plouffe_formula
// https://www.davidhbailey.com/dhbpapers/digits.pdf
// https://www.davidhbailey.com/dhbpapers/bbp-formulas.pdf
// http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.68.2817&rep=rep1&type=pdf

func GenWorkerPi(p uint) func(id int, result chan *big.Float) {
	B1 := new(big.Float).SetPrec(p).SetInt64(1)
	B2 := new(big.Float).SetPrec(p).SetInt64(2)
	B4 := new(big.Float).SetPrec(p).SetInt64(4)
	B5 := new(big.Float).SetPrec(p).SetInt64(5)
	B6 := new(big.Float).SetPrec(p).SetInt64(6)
	B8 := new(big.Float).SetPrec(p).SetInt64(8)
	B16 := new(big.Float).SetPrec(p).SetInt64(16)

	return func(id int, result chan *big.Float) {
		Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))

		C1 := new(big.Float).SetPrec(p).SetInt64(1)
		for i := 0; i < id; i++ {
			C1.Mul(C1, B16)
		}

		C2 := new(big.Float).SetPrec(p).Mul(B8, Bn)

		T1 := new(big.Float).SetPrec(p).Add(C2, B1)
		T1.Quo(B4, T1)

		T2 := new(big.Float).SetPrec(p).Add(C2, B4)
		T2.Quo(B2, T2)

		T3 := new(big.Float).SetPrec(p).Add(C2, B5)
		T3.Quo(B1, T3)

		T4 := new(big.Float).SetPrec(p).Add(C2, B6)
		T4.Quo(B1, T4)

		R := new(big.Float).SetPrec(p).Sub(T1, T2)
		R.Sub(R, T3).Sub(R, T4).Quo(R, C1)

		result <- R
	}
}

// ln2
func GenWorkerLn2(p uint) func(id int, result chan *big.Float) {
	B1 := new(big.Float).SetPrec(p).SetInt64(1)
	B2 := new(big.Float).SetPrec(p).SetInt64(2)

	return func(id int, result chan *big.Float) {
		Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))

		C1 := new(big.Float).SetPrec(p).SetInt64(2)
		for i := 0; i < id; i++ {
			C1.Mul(C1, B2)
		}

		T1 := new(big.Float).SetPrec(p).Add(Bn, B1)
		T1.Quo(B1, T1)

		R := new(big.Float).SetPrec(p).Quo(B1, C1)
    R.Mul(R, T1)

		result <- R
	}
}

// ln3
func GenWorkerLn3(p uint) func(id int, result chan *big.Float) {
	B1 := new(big.Float).SetPrec(p).SetInt64(1)
	B2 := new(big.Float).SetPrec(p).SetInt64(2)
  B4 := new(big.Float).SetPrec(p).SetInt64(4)

	return func(id int, result chan *big.Float) {
		Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))

		C1 := new(big.Float).SetPrec(p).SetInt64(1)
		for i := 0; i < id; i++ {
			C1.Mul(C1, B4)
		}
    
    C2 := new(big.Float).SetPrec(p).Mul(B2, Bn)
    
		T1 := new(big.Float).SetPrec(p).Add(C2, B1)
		T1.Quo(B1, T1)

		R := new(big.Float).SetPrec(p).Quo(B1, C1)
    R.Mul(R, T1)

		result <- R
	}
}

// ln5
func GenWorkerLn5(p uint) func(id int, result chan *big.Float) {
	B1 := new(big.Float).SetPrec(p).SetInt64(1)
	B2 := new(big.Float).SetPrec(p).SetInt64(2)
  B3 := new(big.Float).SetPrec(p).SetInt64(3)
  B4 := new(big.Float).SetPrec(p).SetInt64(4)
  B16 := new(big.Float).SetPrec(p).SetInt64(16)

	return func(id int, result chan *big.Float) {
		Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))

		C1 := new(big.Float).SetPrec(p).SetInt64(1)
		for i := 0; i < id; i++ {
			C1.Mul(C1, B16)
		}
    
    C2 := new(big.Float).SetPrec(p).Mul(B4, Bn)
    
		T1 := new(big.Float).SetPrec(p).Add(C2, B1)
		T1.Quo(B1, T1)

    T2 := new(big.Float).SetPrec(p).Add(C2, B2)
		T2.Quo(B1, T2)

    T3 := new(big.Float).SetPrec(p).Add(C2, B3)
		T3.Mul(T3, B4).Quo(B1, T3)

		R := new(big.Float).SetPrec(p).Add(T1, T2)
    R.Add(R, T3).Quo(R, C1)

		result <- R
	}
}

//atan(1/2)
func GenWorkerAtanHalf(p uint) func(id int, result chan *big.Float) {
  B1 := new(big.Float).SetPrec(p).SetInt64(1)
	B3 := new(big.Float).SetPrec(p).SetInt64(3)
	B4 := new(big.Float).SetPrec(p).SetInt64(4)
	B16 := new(big.Float).SetPrec(p).SetInt64(16)

	return func(id int, result chan *big.Float) {
		Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))

		C1 := new(big.Float).SetPrec(p).SetInt64(2)
		for i := 0; i < id; i++ {
			C1.Mul(C1, B16)
		}

		C2 := new(big.Float).SetPrec(p).Mul(B4, Bn)

		T1 := new(big.Float).SetPrec(p).Add(C2, B1)
		T1.Quo(B1, T1)

		T2 := new(big.Float).SetPrec(p).Add(C2, B3)
    T2.Mul(T2, B4)
		T2.Quo(B1, T2)

		R := new(big.Float).SetPrec(p).Sub(T1, T2)
		R.Quo(R, C1)

		result <- R
	}
}

//atan(1/2)
func GenWorkerAtanThird(p uint) func(id int, result chan *big.Float) {
  B1 := new(big.Float).SetPrec(p).SetInt64(1)
	B2 := new(big.Float).SetPrec(p).SetInt64(2)
	B4 := new(big.Float).SetPrec(p).SetInt64(4)
  B5 := new(big.Float).SetPrec(p).SetInt64(5)
  B8 := new(big.Float).SetPrec(p).SetInt64(8)
	B16 := new(big.Float).SetPrec(p).SetInt64(16)

	return func(id int, result chan *big.Float) {
		Bn := new(big.Float).SetPrec(p).SetInt64(int64(id))

		C1 := new(big.Float).SetPrec(p).SetInt64(1)
		for i := 0; i < id; i++ {
			C1.Mul(C1, B16)
		}

		C2 := new(big.Float).SetPrec(p).Mul(B8, Bn)

		T1 := new(big.Float).SetPrec(p).Add(C2, B1)
		T1.Quo(B1, T1)

		T2 := new(big.Float).SetPrec(p).Add(C2, B2)
		T2.Quo(B1, T2)

    T3 := new(big.Float).SetPrec(p).Add(C2, B4)
		T3.Mul(T3, B2).Quo(B1, T3)

    T4 := new(big.Float).SetPrec(p).Add(C2, B5)
		T4.Mul(T4, B4).Quo(B1, T4)

		R := new(big.Float).SetPrec(p).Sub(T1, T2)
		R.Sub(R, T3).Sub(R, T4).Quo(R, C1)

		result <- R
	}
}

type ConstMap map[int64]*big.Float

func (c *ConstMap) Register(x int64, p uint) *big.Float {
  if *c == nil { // empty map
    *c = make(ConstMap)
  }
  if bf, ok := (*c)[x]; ok { // existing one
    return bf
  }
  // new one
  bf := new(big.Float).SetPrec(p).SetInt64(x)
  (*c)[x] = bf
  return bf
}

// set Num or Den to 0 if not using it
type Fraction struct {
  Num int // numerator, not used if 0 or 1
  Den uint // denominator, not used if 0 or 1
}

func (f Fraction) UseNum() bool {
  return f.Num != 0 && f.Den != 1
}

func (f Fraction) UseDen() bool {
  return f.Den != 0 && f.Den != 1
}

func (f Fraction) UseIt() bool {
  return f.UseNum() || f.UseDen()
}

func (f Fraction) IsZero() bool {
  return f.Num == 0 && f.Den == 0
}

type BBPWorker func(id uint, result chan *big.Float)

// BBPFormula is defined as 
// P(s, b, m, A) = Sum_{k=0}^{Inf}(1/(b^k) * Sum_{j=1}^{m}(a_j / (m*k+j)^s))
// A = (a_1, a_2,..., a_m)
type BBPFormula struct {
  Mall Fraction // overall multiplier
  Power uint // power of the polymonials (s)
  Base uint // base (b)
  Mk uint // multiplier of k (m)
  Alist []Fraction // the a_j's (A). If shorter than Mk, it's treated as 0.
}

func (f BBPFormula) TermsNeeded(nBits uint) uint {
  n := uint(0)
	for x := f.Base; x > 1; x >>= 1 {
		n++
	}
	if f.Base&^(1<<n) != 0 { // not exact 2's power (like 9)
		n++
	}
	return (nBits + n - 1) / n + 1 // N+1 for k=0..N
  
 //return uint(math.Ceil(float64(nBits) / math.Log2(float64(f.Base))))+1
}

func (f BBPFormula) Check() error {
  if f.Mall.IsZero() {
    return fmt.Errorf("zero overall multiplier specified")
  }
  if f.Power == 0 {
    return fmt.Errorf("zero Power specified")
  }
  if f.Base <= 1 {
    return fmt.Errorf("invalid Base specified: needs to be larger than 1")
  }
  if f.Mk == 0 {
    return fmt.Errorf("zero Mk specified")
  }
  la := len(f.Alist) 
  if la == 0 {
    return fmt.Errorf("length of Alist (%d) too short: expected in [1..%d]", la, f.Mk)
  } else if la > int(f.Mk) {
    return fmt.Errorf("length of Alist (%d) too long: expected at most %d", la, f.Mk)
  }
  return nil
}

func (f BBPFormula) GenWorker(p uint) (BBPWorker, error) {
  if p == 0 {
    return nil, fmt.Errorf("zero precision p specified")
  }
  if err := f.Check(); err != nil {
    return nil, err
  }
  var cm ConstMap
  lAlist := len(f.Alist)
  Base := cm.Register(int64(f.Base), p)
  Mk := cm.Register(int64(f.Mk), p)
  B1 := cm.Register(1, p)
  
  Alist := make([][3]*big.Float, lAlist)
  // [3]*big.Float: the J, numerator, and denominator
  for j, aj := range f.Alist {
    if aj.IsZero() {
      continue
    }
    Alist[j][0] = cm.Register(int64(j+1), p) // 1-base J!
    if aj.UseIt() {
      if aj.UseNum() { 
        Alist[j][1] = cm.Register(int64(aj.Num), p)
      }
      if aj.UseDen() {
        Alist[j][2] = cm.Register(int64(aj.Den), p)
      }
    }
  }
  
  return func(id uint, result chan *big.Float) {
		Bk := new(big.Float).SetPrec(p).SetUint64(uint64(id))
		C1 := new(big.Float).SetPrec(p).SetInt64(1)
		for i := uint(0); i < id; i++ {
			C1.Mul(C1, Base)
		}
		C2 := new(big.Float).SetPrec(p).Mul(Mk, Bk)
    
    R := new(big.Float).SetPrec(p).SetInt64(0)
    for j, aj := range f.Alist {
      if aj.IsZero() {
        continue
      }
      T := new(big.Float).SetPrec(p).Add(C2, Alist[j][0])
      if f.Power > 1 {
        U := new(big.Float).SetPrec(p).Set(T)
        for i := uint(1); i < f.Power; i++ {
          U.Mul(U, T)
        }
        T.Set(U)
      }
      if aj.UseDen() {
        T.Mul(T, Alist[j][2])
      }
      if aj.UseNum() { 
        T.Quo(Alist[j][1], T)
      } else {
        T.Quo(B1, T)
      }
      R.Add(R, T)
    }
    R.Quo(R, C1)
    
		result <- R
	}, nil
}

func (f BBPFormula) Calculate(nBits uint) (*big.Float, error) {
  runtime.GOMAXPROCS(runtime.NumCPU())

  p := nBits + 32 // big float precision
  n := f.TermsNeeded(nBits)

  result := make(chan *big.Float, n)
	worker, err := f.GenWorker(p)
  if err != nil {
    return nil, fmt.Errorf("failed to generate worker function: %w", err)
  }

	value := new(big.Float).SetPrec(p).SetInt64(0)

	for i := uint(0); i < n; i++ {
		go worker(i, result)
	}
	for i := uint(0); i < n; i++ {
		value.Add(value, <-result)
	}
  if f.Mall.UseIt() {
    if f.Mall.UseNum() { 
      x := new(big.Float).SetPrec(p).SetInt64(int64(f.Mall.Num))
      value.Mul(value, x)
    }
    if f.Mall.UseDen() {
      x := new(big.Float).SetPrec(p).SetUint64(uint64(f.Mall.Den))
      value.Quo(value, x)
    }
  }
  return value, nil
}

func GetBytes(value *big.Float, nBytes uint) []byte {
  nb := 8 * int(nBytes)
  xx := new(big.Float).SetMantExp(value, nb - value.MantExp(nil))
  x, _ := xx.Int(nil)
  
  buf := make([]byte, nBytes)
  x.FillBytes(buf)
  return buf
}

var bbpNames = []string{
  "pi", "ln2", "ln3", "ln5", "atan(1/2)", "atan(1/3)",
  "ln7", "ln10",
}

var bbpList = []BBPFormula{
  // pi
  {Fraction{1, 0}, 1, 16, 8, []Fraction{
    {4, 0}, {0, 0}, {0, 0}, {-2, 0},
    {-1, 0}, {-1, 0}, 
  }},
  // ln2
  {Fraction{1, 16}, 1, 16, 4, []Fraction{
    {8, 0}, {4, 0}, {2, 0}, {1, 0},
  }},
  // ln3
  {Fraction{1, 0}, 1, 16, 4, []Fraction{
    {1, 0}, {0, 0}, {1, 4}, 
  }},
  // ln5
  {Fraction{1, 0}, 1, 16, 4, []Fraction{
    {1, 0}, {1, 0}, {1, 4}, 
  }},
  // atan(1/2)
  {Fraction{1, 2}, 1, 16, 4, []Fraction{
    {1, 0}, {0, 0}, {-1, 4}, 
  }},
  // atan(1/3)
  {Fraction{1, 0}, 1, 16, 8, []Fraction{
    {1, 0}, {-1, 0}, {0, 0}, {-1, 2},
    {-1, 4}, 
  }},
  // ln7
  {Fraction{3, 4}, 1, 8, 3, []Fraction{
    {2, 0}, {1, 0},
  }},
  // ln10
  {Fraction{1, 16}, 1, 16, 4, []Fraction{
    {24, 0}, {20, 0}, {6, 0}, {1, 0}, 
  }},
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

  const nBytes = 64
  const nb = 8 * nBytes

  p := uint(nb + 32) // big float precision
  n := nb/4 + 1 // hex digits
  nd := int(float64(nb) / math.Log2(10) + 1.0) // decimal digits

  // Pi: 3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709384460955058223172535940812848112
  start := time.Now()
	result := make(chan *big.Float, n)
	worker := GenWorkerPi(p)

	pi := new(big.Float).SetPrec(p).SetInt64(0)

	for i := 0; i < n; i++ {
		go worker(i, result)
	}

	for i := 0; i < n; i++ {
		pi.Add(pi, <-result)
	}

	dur := time.Since(start)
	fmt.Printf("take %v to calculate %d (%d hex) digits (%d/%d bits) long pi \n", dur, nd, n, nb, p)
	fmt.Printf("%[1]*.[2]*[3]f\n", 1, nd, pi)

  xx := new(big.Float).SetMantExp(pi, nb-pi.MantExp(nil))
  x, _ := xx.Int(nil)
  
  buf := make([]byte, nBytes)
  x.FillBytes(buf)
  fmt.Printf("Buf: % 02x \n", buf) 


  // ln2 = 0.69314718055994530941723212145817656807550013436025525412068000949339362196969471560586332699641868754200148102057068573368552023575813055703267075163507596
  start = time.Now()
	result = make(chan *big.Float, nb+1)
	worker = GenWorkerLn2(p)

	ln2 := new(big.Float).SetPrec(p).SetInt64(0)

	for i := 0; i < nb+1; i++ {
		go worker(i, result)
	}

	for i := 0; i < nb+1; i++ {
		ln2.Add(ln2, <-result)
	}

	dur = time.Since(start)
	fmt.Printf("take %v to calculate %d (%d hex) digits (%d/%d bits) long ln2 \n", dur, nd, n, nb, p)
	fmt.Printf("%[1]*.[2]*[3]f\n", 1, nd, ln2)

  xx = new(big.Float).SetMantExp(ln2, nb-ln2.MantExp(nil))
  x, _ = xx.Int(nil)
  
  buf = make([]byte, nBytes)
  x.FillBytes(buf)
  fmt.Printf("Buf: % 02x \n", buf) 

  // ln3 = 1.09861228866810969139524523692252570464749055782274945173469433363749429321860896687361575481373208878797002906595786574236800422593051982105280187076727741
  start = time.Now()
	result = make(chan *big.Float, nb/2+1)
	worker = GenWorkerLn3(p)

	ln3 := new(big.Float).SetPrec(p).SetInt64(0)

	for i := 0; i < nb/2+1; i++ {
		go worker(i, result)
	}

	for i := 0; i < nb/2+1; i++ {
		ln3.Add(ln3, <-result)
	}

	dur = time.Since(start)
	fmt.Printf("take %v to calculate %d (%d hex) digits (%d/%d bits) long ln3 \n", dur, nd, n, nb, p)
	fmt.Printf("%[1]*.[2]*[3]f\n", 1, nd, ln3)

  xx = new(big.Float).SetMantExp(ln3, nb-ln3.MantExp(nil))
  x, _ = xx.Int(nil)
  
  buf = make([]byte, nBytes)
  x.FillBytes(buf)
  fmt.Printf("Buf: % 02x \n", buf) 

  //Ln5 = 1.60943791243410037460075933322618763952560135426851772191264789147417898770765776463013387809317961079996630302171556289972400522932467619963361661746370573
  start = time.Now()
	result = make(chan *big.Float, n)
	worker = GenWorkerLn5(p)

	ln5 := new(big.Float).SetPrec(p).SetInt64(0)

	for i := 0; i < n; i++ {
		go worker(i, result)
	}

	for i := 0; i < n; i++ {
		ln5.Add(ln5, <-result)
	}

	dur = time.Since(start)
	fmt.Printf("take %v to calculate %d (%d hex) digits (%d/%d bits) long ln5 \n", dur, nd, n, nb, p)
	fmt.Printf("%[1]*.[2]*[3]f\n", 1, nd, ln5)

  xx = new(big.Float).SetMantExp(ln5, nb-ln5.MantExp(nil))
  x, _ = xx.Int(nil)
  
  buf = make([]byte, nBytes)
  x.FillBytes(buf)
  fmt.Printf("Buf: % 02x \n", buf) 

  // atan(1/2) = 0.46364760900080611621425623146121440202853705428612026381093308872019786416574170530060028398488789255652985225119083751350581818162501115547153056994410562
  start = time.Now()
	result = make(chan *big.Float, n)
	worker = GenWorkerAtanHalf(p)

	atan0_5 := new(big.Float).SetPrec(p).SetInt64(0)

	for i := 0; i < n; i++ {
		go worker(i, result)
	}

	for i := 0; i < n; i++ {
		atan0_5.Add(atan0_5, <-result)
	}

	dur = time.Since(start)
	fmt.Printf("take %v to calculate %d (%d hex) digits (%d/%d bits) long atan(1/2) \n", dur, nd, n, nb, p)
	fmt.Printf("%[1]*.[2]*[3]f\n", 1, nd, atan0_5)

  xx = new(big.Float).SetMantExp(atan0_5, nb-atan0_5.MantExp(nil))
  x, _ = xx.Int(nil)
  
  buf = make([]byte, nBytes)
  x.FillBytes(buf)
  fmt.Printf("Buf: % 02x \n", buf) 

  // atan(1/3) = 0.32175055439664219340140461435866131902075529555765619143280305935675623740581054435640842235064137443900716937712973914826764297076263440245980928208801466
  start = time.Now()
	result = make(chan *big.Float, n)
	worker = GenWorkerAtanThird(p)

	atan3rd := new(big.Float).SetPrec(p).SetInt64(0)

	for i := 0; i < n; i++ {
		go worker(i, result)
	}

	for i := 0; i < n; i++ {
		atan3rd.Add(atan3rd, <-result)
	}

	dur = time.Since(start)
	fmt.Printf("take %v to calculate %d (%d hex) digits (%d/%d bits) long atan(1/3) \n", dur, nd, n, nb, p)
	fmt.Printf("%[1]*.[2]*[3]f\n", 1, nd, atan3rd)

  xx = new(big.Float).SetMantExp(atan3rd, nb-atan3rd.MantExp(nil))
  x, _ = xx.Int(nil)
  
  buf = make([]byte, nBytes)
  x.FillBytes(buf)
  fmt.Printf("Buf: % 02x \n", buf) 

  for i, bbp := range bbpList {
    fmt.Printf("\n\n****** %s *******\n", bbpNames[i])
    start = time.Now()
    val, err := bbp.Calculate(nb)
    dur = time.Since(start)
    if err != nil {
      fmt.Printf("ERROR: %s!\n", err)
      return
    }
  	fmt.Printf("take %v to calculate %d (%d hex) digits (%d/%d bits) long %s \n", dur, nd, n, nb, p, bbpNames[i])
  	fmt.Printf("%[1]*.[2]*[3]f\n", 1, nd, val)
    buf := GetBytes(val, nBytes)
    fmt.Printf("Buf: % 02x \n", buf)
  }
}
