package main

import (
    "testing"
    "os"
)

func TestNewDec(t *testing.T){
     cards:=newDeck()

     if len(cards)!=16{
         t.Errorf("Expected a value of 16 but got %v",len(cards))
     }
    if cards[0] != "Ace of Spades"{
        t.Errorf("Expected first card Ace of Spades but got %v",cards[0])
    }
    if(cards[len(cards)-1] != "Four of Clubs"){
        t.Errorf("Expected last card Four Of Clubs %v",cards[len(cards)-1])
    }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
   os.Remove("_decktesting")
   deck:=newDeck()
   deck.saveToFile("_decktesting")
   loadedDeck:=newDeckFromFile("_decktesting")
   if len(loadedDeck)!=16{
       t.Errorf("Expected 16 cards but got %v",len(loadedDeck))
   }
   os.Remove("_decktesting")
}