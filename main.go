package main

import( 
  "fmt"
  "os"
  "github.com/bwmarrin/discordgo"
)

func main() {
  botSession, err := discordgo.New("Bot " + os.Getenv("token"))

  if err != nil{
    panic(err)
  }
  botSession.AddHandler(messageHandler)

  err = botSession.Open()
  if err != nil{
    fmt.Println("Error opening session ", err)
    panic(err)
  }

  fmt.Println("discord session opened. Type CTRL + C to stop")

  for{}
  botSession.Close()


}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate){

  raccoonReader, err := os.Open("raccoon.jpg")
  if err != nil{
    panic(err)
  }


  if m.Content == "ping"{
    s.ChannelMessageSend(m.ChannelID, "pong")
  } else if m.Content == "levitate"{
    s.ChannelFileSend(m.ChannelID, "raccoon.jpg", raccoonReader )
  }
}