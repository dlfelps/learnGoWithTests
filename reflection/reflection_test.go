package main

import (
  "testing"
  "slices"
)

func TestWalk(t *testing.T) {

  cases := []struct {
    Name          string
    Input         interface{}
    ExpectedCalls []string
  }{
    {
      "struct with one string field",
      struct{ Name string }{"Chris"},
      []string{"Chris"},
    },
    {
      "struct with two string fields",
      struct {
        Name string
        City string
      }{"Chris", "London"},
      []string{"Chris", "London"},
    },
    {
      "struct with non string field",
      struct {
        Name string
        Age  int
      }{"Chris", 33},
      []string{"Chris"},
    },
    {
      "nested fields",
      Person{
        "Chris",
        Profile{33, "London"},
      },
      []string{"Chris", "London"},
    },
    {
      "pointers to things",
      &Person{
        "Chris",
        Profile{33, "London"},
      },
      []string{"Chris", "London"},
    },
    {
      "slices",
      []Profile{
        {33, "London"},
        {34, "Reykjavík"},
      },
      []string{"London", "Reykjavík"},
    },
    {
      "arrays",
      [2]Profile{
        {33, "London"},
        {34, "Reykjavík"},
      },
      []string{"London", "Reykjavík"},
    },
  }

  for _, test := range cases {
    t.Run(test.Name, func(t *testing.T) {
      var got []string
      walk(test.Input, func(input string) {
        got = append(got, input)
      })

      if !slices.Equal(got, test.ExpectedCalls) {
        t.Errorf("got %v, want %v", got, test.ExpectedCalls)
      }
    })
  }

  t.Run("with maps", func(t *testing.T) {
    aMap := map[string]string{
      "Foo": "Bar",
      "Baz": "Boz",
    }

    var got []string
    walk(aMap, func(input string) {
      got = append(got, input)
    })

    assertContains(t, got, "Bar")
    assertContains(t, got, "Boz")
  })

  t.Run("with channels", func(t *testing.T) {
    aChannel := make(chan Profile)

    go func() {
      aChannel <- Profile{33, "Berlin"}
      aChannel <- Profile{34, "Katowice"}
      close(aChannel)
    }()

    var got []string
    want := []string{"Berlin", "Katowice"}

    walk(aChannel, func(input string) {
      got = append(got, input)
    })

    if !slices.Equal(got, want) {
      t.Errorf("got %v, want %v", got, want)
    }
  })

  t.Run("with function", func(t *testing.T) {
    aFunction := func() (Profile, Profile) {
      return Profile{33, "Berlin"}, Profile{34, "Katowice"}
    }

    var got []string
    want := []string{"Berlin", "Katowice"}

    walk(aFunction, func(input string) {
      got = append(got, input)
    })

    if !slices.Equal(got, want) {
      t.Errorf("got %v, want %v", got, want)
    }
  })
}

type Person struct {
  Name    string
  Profile Profile
}

type Profile struct {
  Age  int
  City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
  t.Helper()
  contains := false
  for _, x := range haystack {
    if x == needle {
      contains = true
      break
    }
  }
  if !contains {
    t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
  }
}