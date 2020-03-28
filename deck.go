package main
import (
    "fmt"
    "strings"
    "io/ioutil"
    "os"
    "math/rand"
    "time"
)

type deck []string

func newDeck() deck{
    cards:=deck{}
    cardSuits:=[]string{"Spades","Diamonds","Hearts","Clubs"}
    cardValues:=[]string{"Ace","Two","Three","Four"}

    for _,suit:=range cardSuits{
        for _,value:=range cardValues{
            cards=append(cards,value+" of "+suit)
        }
    }
    return cards
}

func (c deck) print(){
    for i,values:=range c{
        fmt.Println(i,values)
    }
}

func deal(d deck,handSize int)(deck,deck){
     return d[:handSize],d[handSize:]
}

func (d deck) toString() string{
    return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error{
        return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}
func (d deck) shuffle(){
    //the problem with shuffle method is that it uses a constant seed to generate 
    //randomness and that is why for same input we get same random value everytime
    //  rand.Shuffle(len(d),func(i,j int){
    //      d[i],d[j]=d[j],d[i]
    //  })
    source:=rand.NewSource(time.Now().UnixNano())
    r:=rand.New(source)
    for i:=range d{
        var random int=r.Intn(len(d)-1)
        d[random],d[i]=d[i],d[random]
    }
}
func newDeckFromFile(filename string) deck{
   bs,err:=ioutil.ReadFile(filename)
   if err!=nil{
      fmt.Println("Error",err)
      os.Exit(1)
   }
   return deck(strings.Split(string(bs),","))
}