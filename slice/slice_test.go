package slice

import "testing"

func TestShouldBeTrueWhenSlicesAreSame(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	result := AreSlicesEqual(a, b)
	want := true
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeFalseWhenStringSlicesAreNotSame(t *testing.T) {
	a := []string{"one", "two", "three"}
	b := []string{"four", "five", "six"}
	result := AreSlicesEqual(a, b)
	want := false
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueForEmptyStringSlices(t *testing.T) {
	a := []string{}
	b := []string{}
	result := AreSlicesEqual(a, b)
	want := true
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueWhenStringSlicesAreSame(t *testing.T) {
	a := []string{"one", "two", "three"}
	b := []string{"one", "two", "three"}
	result := AreSlicesEqual(a, b)
	want := true
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeFalseWhenSlicesAreNotSame(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	result := AreSlicesEqual(a, b)
	want := false
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueForEmptySlices(t *testing.T) {
	a := []int{}
	b := []int{}
	result := AreSlicesEqual(a, b)
	want := true
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
