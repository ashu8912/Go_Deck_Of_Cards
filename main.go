package main

func main(){
    cards:= newDeck()
    cards.saveToFile("my_cards")
    d:=newDeckFromFile("my_cards")
    //d.print()
    d.shuffle()
    d.saveToFile("my_cards")
    d.print()
}