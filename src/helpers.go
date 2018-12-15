package main

func min(x, y int) int {
  if x < y { return x }
  return y
}

func max(x, y int) int {
  if x > y { return x }
  return y
}

func contains(s []string, e string) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}