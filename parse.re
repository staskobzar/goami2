// -*-go-*-
//
//go:generate re2go ami_re2c.re -o ami_re2c.go -i

package goami2

import "fmt"

// Parse AMI message as string into *Message structure
func Parse(data string) (*Message, error) {
	msg := NewMessage()
	var cur, mar int
	var ns, ne, vs, ve int
	/*!stags:re2c format = "\tvar @@ int\n"; */
	for { /*!re2c
				re2c:yyfill:enable = 0;
				re2c:define:YYCTYPE     = byte;
				re2c:define:YYPEEK      = "data[cur]";
				re2c:define:YYSKIP      = "cur += 1";
				re2c:define:YYBACKUP    = "mar = cur";
				re2c:define:YYRESTORE   = "cur = mar";
				re2c:define:YYSTAGP     = "@@{tag} = cur";
				re2c:define:YYSTAGN     = "@@{tag} = -1";
				re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";
				re2c:tags = 1;

				CRLF  = "\r\n";
		        alpha = [a-zA-Z];
				name  = alpha (alpha | "-")+;
				value = [^\r\n]+;

				*    { break }
				CRLF { return msg, nil }
				@ns name @ne ":" [ ]* @vs value? @ve CRLF {
					msg.AddField(data[ns:ne], data[vs:ve])
				}
		*/
	}
	return nil, fmt.Errorf("%w: invalid input: %q", ErrAMI, data[cur:])
}

// vi: ft=go
