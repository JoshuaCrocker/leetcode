// A valid number can be split up into these components (in order):
//
//     A decimal number or an integer.
//     (Optional) An 'e' or 'E', followed by an integer.
//
// A decimal number can be split up into these components (in order):
//
//     (Optional) A sign character (either '+' or '-').
//     One of the following formats:
//         One or more digits, followed by a dot '.'.
//         One or more digits, followed by a dot '.', followed by one or more digits.
//         A dot '.', followed by one or more digits.
//
// An integer can be split up into these components (in order):
//
//     (Optional) A sign character (either '+' or '-').
//     One or more digits.
//
// For example, all the following are valid numbers: ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"], while the following are not valid numbers: ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"].
//
// Given a string s, return true if s is a valid number.
package lc65validnumber

func IsNumber(s string) bool {
	parser := &NumberInput{
		Input: s,
	}

	return parser.Parse() == nil
}

type NumberInput struct {
	Input string

	// TODO syntax tree
}

type NumberParser interface {
	Parse() error

	accept(expected string) bool
}

// type TokenType int8

// const (
// 	T_INTEGER       = iota
// 	T_FLOAT         = iota
// 	T_EXPONENTIAL   = iota
// 	T_DECIMAL_POINT = iota
// 	T_SIGN          = iota
// )

// type Token struct {
// 	Type TokenType
// 	Text string
// }

func (ni *NumberInput) Parse() error {
	// if ni.accept("+") {
	// 	// woo!
	// } else if ni.accept("-") {
	// 	// woo!
	// }

	// err := ni.digits()

	return nil
}
