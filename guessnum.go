package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  const (
    minNumber = 0
    maxNumber = 100

    maxTries = 10
  )

  setup()
  secretNumber := generateSecretNumber(minNumber, maxNumber)

  lowest := minNumber
  highest := maxNumber

  fmt.Println("Guess the number!")

  for tries := maxTries ; tries >= 0 ; tries-- {
    fmt.Printf("Known range is %d - %d: ", lowest, highest)
    var guess int
    numScanned, error := fmt.Scanf("%d", &guess)

    if error != nil || numScanned != 1 {
      fmt.Println("There was a problem")
    } else {
      switch {
      case guess == secretNumber:
        fmt.Println("That's it!")
        return
      case guess < secretNumber:
        fmt.Print("Too low! ")
        lowest = max(lowest, guess + 1)
      case guess > secretNumber:
        fmt.Print("Too high! ")
        highest = min(highest, guess - 1)
      }

      fmt.Printf("%d tries left...\n", tries)
    }
  }

  fmt.Println("Sorry! Looks like you ran out of tries! " +
              "The secret number was %d!", secretNumber)

}

func setup() {
  rand.Seed(time.Now().UTC().Unix())
}

func generateSecretNumber(minNumber, maxNumber int) int {
  numberRange := maxNumber - minNumber
  return rand.Intn(numberRange) + minNumber
}

func min(a, b int) int {
  if (a < b) { return a }
  return b
}

func max(a, b int) int {
  if (a > b) { return a }
  return b
}
