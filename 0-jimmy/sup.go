package main

import "fmt"

const (
  spanish = "Spanish"
  french = "French"

  englishHelloPrefix = "Deez "
  spanishHelloPrefix = "Estaz "
  frenchHelloPrefix = "Deezeaux "
)

func Hello(name, language string) string {
  if name == "" {
    name = "nuts"
  }
    
  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
  switch language {
  case spanish:
    prefix = spanishHelloPrefix
  case french:
    prefix = frenchHelloPrefix
  default:
    prefix = englishHelloPrefix
  }
  return
}

func main() {
  fmt.Println(Hello("nuts", ""))
}
