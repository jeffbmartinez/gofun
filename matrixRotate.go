package main

import (
  "fmt"
)

type Matrix struct {
  Width int
  Height int
  matrix [][]int
}

// Create an empty matrix
func NewMatrix(width, height int) *Matrix {
  matrix := make([][]int, height)
  for row := range matrix {
    matrix[row] = make([]int, width)
  }

  return &Matrix{width, height, matrix}
}

// Display a matrix
func (m Matrix) Display() {
  for row := range m.matrix {
    for column := range m.matrix[row] {
      fmt.Printf("%d\t", m.matrix[row][column])
    }
    fmt.Println()
  }
}

// Assign numbers to the matrix
func (m *Matrix) Fill() {
  num := 0
  for row := range m.matrix {
    for column := range m.matrix[row] {
      m.matrix[row][column] = num
      num++
    }
  }
}

// Flip matrix on negative diagonal
func (m Matrix) FlipOnNegativeDiagonal() (*Matrix) {
  newWidth := m.Height
  newHeight := m.Width

  flippedMatrix := NewMatrix(newWidth, newHeight)
  for row := range m.matrix {
    for column := range m.matrix[row] {
      flippedMatrix.matrix[column][row] = m.matrix[row][column]
    }
  }

  return flippedMatrix
}

// Flip matrix on vertical axis
func (m Matrix) FlipOnVerticalAxis() *Matrix {
  flippedMatrix := NewMatrix(m.Width, m.Height)

  for row := range m.matrix {
    for column := range m.matrix[row] {
      flippedMatrix.matrix[row][m.Width - column - 1] = m.matrix[row][column]
    }
  }

  return flippedMatrix
}

// Return a new instance of a matrix, rotated 90 degrees clockwise
func (m Matrix) Rotate90() *Matrix {
  return m.FlipOnNegativeDiagonal().FlipOnVerticalAxis()
}

func main() {
  const width int = 4
  const height int = 11

  fmt.Println("------------------------")

  matrix := NewMatrix(width, height)
  matrix.Fill()
  matrix.Display()

  for i := 0 ; i < 4 ; i++ {
    fmt.Println("------------------------")
    matrix = matrix.Rotate90()
    matrix.Display()
  }

  fmt.Println("------------------------")
}
