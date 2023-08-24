package my

import (
	"reflect"
	"testing"
)

func TestNewSSet(t *testing.T) {
	cases := []struct {
		Got, Want SSet
	}{
		{
			NewSSet(),
			NewSSet(),
		},
		{
			NewSSet("1"),
			NewSSet("1"),
		},
		{
			NewSSet("1", "2"),
			NewSSet("1", "2"),
		},
	}
	for _, c := range cases {
		got := c.Got
		want := c.Want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		Init  SSet
		Added []string
		Want  SSet
	}{
		{
			NewSSet(),
			[]string{"1"},
			NewSSet("1"),
		},
		{
			NewSSet("1"),
			[]string{"2"},
			NewSSet("1", "2"),
		},
		{
			NewSSet(),
			[]string{"1", "2", "3"},
			NewSSet("1", "2", "3"),
		},
	}
	for _, c := range cases {
		got := c.Init
		for _, s := range c.Added {
			got.Add(s)
		}
		want := c.Want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestAnd(t *testing.T) {
	cases := []struct {
		Init  SSet
		Other SSet
		Want  SSet
	}{
		{
			NewSSet(),
			NewSSet("1"),
			NewSSet(),
		},
		{
			NewSSet("1"),
			NewSSet("2"),
			NewSSet(),
		},
		{
			NewSSet("1", "2"),
			NewSSet("2", "3"),
			NewSSet("2"),
		},
	}
	for _, c := range cases {
		got := c.Init.And(c.Other)
		want := c.Want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestOr(t *testing.T) {
	cases := []struct {
		Init  SSet
		Other SSet
		Want  SSet
	}{
		{
			NewSSet(),
			NewSSet("1"),
			NewSSet("1"),
		},
		{
			NewSSet("1"),
			NewSSet("2"),
			NewSSet("1", "2"),
		},
		{
			NewSSet("1", "2"),
			NewSSet("2", "3"),
			NewSSet("1", "2", "3"),
		},
	}
	for _, c := range cases {
		got := c.Init.Or(c.Other)
		want := c.Want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		Init  SSet
		Other SSet
		Want  SSet
	}{
		{
			NewSSet(),
			NewSSet(),
			NewSSet(),
		},
		{
			NewSSet(),
			NewSSet("1"),
			NewSSet(),
		},
		{
			NewSSet("1"),
			NewSSet(),
			NewSSet("1"),
		},
		{
			NewSSet("1"),
			NewSSet("1"),
			NewSSet(),
		},
		{
			NewSSet("1", "2"),
			NewSSet("2", "3"),
			NewSSet("1"),
		},
	}
	for _, c := range cases {
		got := c.Init.Sub(c.Other)
		want := c.Want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestHas(t *testing.T) {
	cases := []struct {
		Set  SSet
		Key  string
		Want bool
	}{
		{
			NewSSet(),
			"1",
			false,
		},
		{
			NewSSet("1"),
			"1",
			true,
		},
		{
			NewSSet("1", "2", "3"),
			"2",
			true,
		},
		{
			NewSSet("1", "2", "3"),
			"4",
			false,
		},
	}
	for _, c := range cases {
		got := c.Set.Has(c.Key)
		want := c.Want
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
