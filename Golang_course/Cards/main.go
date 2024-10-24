/*
Basic Go Types
1. bool(True or False)
2. string(includes chars "Hello World")
3. int(integer values i.e 0, -10000, 999999)
4. float64(similar to integers but with decimation 10.00001, 0.00009, -100.0003)
*/

package main 


func main() {
	// var card string = "Ace of Spades"
	cards := deck{"Ace of Diamonds",newCard()}
	cards  = append(cards, "Six of Spades")

	cards.print()

} 

func newCard() string{
	return "Five of Diamonds"
}