package mymath

import (
	"math"
)

func Abs(x float64) float64                       { return math.Abs(x) }
func Acos(x float64) float64                      { return math.Acos(x) }
func Acosh(x float64) float64                     { return math.Acosh(x) }
func Asin(x float64) float64                      { return math.Asin(x) }
func Asinh(x float64) float64                     { return math.Asinh(x) }
func Atan(x float64) float64                      { return math.Atan(x) }
func Atan2(y, x float64) float64                  { return math.Atan2(y, x) }
func Atanh(x float64) float64                     { return math.Atanh(x) }
func Cbrt(x float64) float64                      { return math.Cbrt(x) }
func Ceil(x float64) float64                      { return math.Ceil(x) }
func Copysign(f, sign float64) float64            { return math.Copysign(f, sign) }
func Cos(x float64) float64                       { return math.Cos(x) }
func Cosh(x float64) float64                      { return math.Cosh(x) }
func Dim(x, y float64) float64                    { return math.Dim(x, y) }
func Erf(x float64) float64                       { return math.Erf(x) }
func Erfc(x float64) float64                      { return math.Erfc(x) }
func Erfcinv(x float64) float64                   { return math.Erfcinv(x) }
func Erfinv(x float64) float64                    { return math.Erfinv(x) }
func Exp(x float64) float64                       { return math.Exp(x) }
func Exp2(x float64) float64                      { return math.Exp2(x) }
func Expm1(x float64) float64                     { return math.Expm1(x) }
func FMA(x, y, z float64) float64                 { return math.FMA(x, y, z) }
func Float32bits(f float32) uint32                { return math.Float32bits(f) }
func Float32frombits(b uint32) float32            { return math.Float32frombits(b) }
func Float64bits(f float64) uint64                { return math.Float64bits(f) }
func Float64frombits(b uint64) float64            { return math.Float64frombits(b) }
func Floor(x float64) float64                     { return math.Floor(x) }
func Frexp(f float64) (frac float64, exp int)     { return math.Frexp(f) }
func Gamma(x float64) float64                     { return math.Gamma(x) }
func Hypot(p, q float64) float64                  { return math.Hypot(p, q) }
func Ilogb(x float64) int                         { return math.Ilogb(x) }
func Inf(sign int) float64                        { return math.Inf(sign) }
func IsInf(f float64, sign int) bool              { return math.IsInf(f, sign) }
func IsNaN(f float64) (is bool)                   { return math.IsNaN(f) }
func J0(x float64) float64                        { return math.J0(x) }
func J1(x float64) float64                        { return math.J1(x) }
func Jn(n int, x float64) float64                 { return math.Jn(n, x) }
func Ldexp(frac float64, exp int) float64         { return math.Ldexp(frac, exp) }
func Lgamma(x float64) (lgamma float64, sign int) { return math.Lgamma(x) }
func Log(x float64) float64                       { return math.Log(x) }
func Log10(x float64) float64                     { return math.Log10(x) }
func Log1p(x float64) float64                     { return math.Log1p(x) }
func Log2(x float64) float64                      { return math.Log2(x) }
func Logb(x float64) float64                      { return math.Logb(x) }
func Max(x, y float64) float64                    { return math.Max(x, y) }
func Min(x, y float64) float64                    { return math.Min(x, y) }
func Mod(x, y float64) float64                    { return math.Mod(x, y) }
func Modf(f float64) (int float64, frac float64)  { return math.Modf(f) }
func NaN() float64                                { return math.NaN() }
func Nextafter(x, y float64) (r float64)          { return math.Nextafter(x, y) }
func Nextafter32(x, y float32) (r float32)        { return math.Nextafter32(x, y) }
func Pow(x, y float64) float64                    { return math.Pow(x, y) }
func Pow10(n int) float64                         { return math.Pow10(n) }
func Remainder(x, y float64) float64              { return math.Remainder(x, y) }
func Round(x float64) float64                     { return math.Round(x) }
func RoundToEven(x float64) float64               { return math.RoundToEven(x) }
func Signbit(x float64) bool                      { return math.Signbit(x) }
func Sin(x float64) float64                       { return math.Sin(x) }
func Sincos(x float64) (sin, cos float64)         { return math.Sincos(x) }
func Sinh(x float64) float64                      { return math.Sinh(x) }
func Sqrt(x float64) float64                      { return math.Sqrt(x) }
func Tan(x float64) float64                       { return math.Tan(x) }
func Tanh(x float64) float64                      { return math.Tanh(x) }
func Trunc(x float64) float64                     { return math.Trunc(x) }
func Y0(x float64) float64                        { return math.Y0(x) }
func Y1(x float64) float64                        { return math.Y1(x) }
func Yn(n int, x float64) float64                 { return math.Yn(n, x) }
