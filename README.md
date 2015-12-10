# go-ascii
The **ascii** package is fast Go implementation of character classification and conversion functions for ASCII bytes.

The functions in this package provide faster alternatives to the [Go](https://golang.org) [unicode](https://golang.org/pkg/unicode/) package.  It implements the subset of Is* and To* functions that make sense for the [ASCII](https://en.wikipedia.org/wiki/ASCII) character set.

#### Character Classification

Function | Definition
-------- | ----------
**IsASCII(c)**    | Return true if byte *c* is a valid ASCII character.
**IsAlnum(c)**    | Return true if byte *c* is a letter or decimal digit.
**IsControl(c)**  | Return true if byte *c* is a control character.
**IsDigit(c)**    | Return true if byte *c* is a decimal digit.
**IsGraph(c)**    | Return true if byte *c* is a character with a graphic represntation.
**IsHexDigit(c)** | Return true if byte *c* is a hexadecimal digit.
**IsLetter(c)**   | Return true if byte *c* is a letter.
**IsLower(c)**    | Return true if byte *c* is a lower case letter.
**IsPrint(c)**    | Return true if byte *c* is a character with a print represntation.
**IsPunct(c)**    | Return true if byte *c* is a punctuation character ( <code>! " # % & ' ( ) * , - . / : ; ? @ [ \\ ] _ { }</code> )
**IsSpace(c)**    | Return true if byte *c* is a space character (space, tab, vertical tab, form feed, carriage return, or linefeed).
**IsSymbol(c)**   | Return true if byte *c* is a symbol ( <code>$ + < = > ^ ` \ &#124; ~</code> )
**IsUpper(c)**    | Return true if byte *c* is an upper case letter.

#### Character Conversion

Function | Definition
-------- | ----------
**ToLower(c)** | Map upper case letters to lower case; return other characters unchanged.
**ToUpper(c)** | Map lower case letters to upper case; return other characters unchanged.

#### Notes

1. This package is suitable if and only if the data is really ASCII.  No attempt is made to detect non-ASCII charater sets or perform other sanity checks.  However, when the data is ASCII, these functions are about 2.5 times faster than the [unicode](https://golang.org/pkg/unicode/) package equivalents.

2. This package differentiates between punctuation and symbols as does the [unicode](https://golang.org/pkg/unicode/) package.  This differs from the [GNU C Library](http://www.gnu.org/software/libc/manual/html_node/Classification-of-Characters.html) and [C Standard Library](http://en.cppreference.com/w/c/string/byte/ispunct), which combine the two into the ispunct() function.

#### Dependencies

- **[GoConvey](https://github.com/smartystreets/goconvey)** for testing.
