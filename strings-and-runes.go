//A go string is a read-only slice of bytes. The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
//In other languages, strings are made of "characters". In go the concept of a character is called a rune.
//It is an integer that represents a unicode code point. 
package main
import (
	"fmt" 
"unicode/utf8"
)

func main(){
	const s ="你好"
	fmt.Println(s)
	fmt.Println(len(s))
//strings are equivalent to []byte, this will produce the length of the raw bytes stored within
//indexing into a string produces the raw byte values at each index. 
	for i := 0; i< len(s);i++{
		fmt.Printf("%v ", s)
	}
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
//A range loop handles strings specially and decodes each rune along with its offset in the string.
for idx, runeValue := range s {
	fmt.Println(runeValue, idx)
	fmt.Printf("%c starts at %d\n", runeValue, idx)
}
fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width
//This demonstrates passing a rune value to a function.

        // examineRune(runeValue)
    }
	str := "Welcome to GeeksforGeeks"
  
    // Accessing the bytes of the given string
    for c := 0; c < len(str); c++ {
  
        fmt.Printf("\nCharacter = %c Bytes = %v", str[c], str[c])
    }
}
// func examineRune(r rune) {
// 	//Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly.
	
// 		if r == '你' {
// 			fmt.Println("found tee")
// 		} else if r == '好' {
// 			fmt.Println("found so sua")
// 		}
// 	}
