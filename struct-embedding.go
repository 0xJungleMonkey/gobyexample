// go supports embedding of struct and interfaces to express a more seamless composition of types. This is not to be confused with // go:embed which is a go directive introduced in go version 1.16+ to embed files and folders into the application binary
package main
import "fmt"
type base struct {
	num int
}
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v",b.num)

}
//container embeds a base 
// an embedding looks like a field without a name
type container struct {
	base
	str string
}
//when creating struct with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name
func main(){
	co := container{
		base: base{1},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:",co.base.num)
	fmt.Println("describe:", co.describe())
	
	type describer interface {
		describe() string
	}
	//Embedding structs with methods may be used to bestow interface implementations onto other structs. Here we see that a container now implements the describer interface because it embeds base.
	var d describer = co
	fmt.Println("describe:",d.describe())
}