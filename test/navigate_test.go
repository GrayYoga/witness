package test

import (
	"strings"
	"testing"

	"github.com/GrayYoga/witness"

	"github.com/GrayYoga/witness/pkg/chrome"
)

func link(t *testing.T, sess *witness.Session, l string) {
	t.Helper()
	err := sess.Page.C(l, true).Click()
	check(t, err)
}

func checkNav(t *testing.T, sess *witness.Session, suffix string) {
	t.Helper()
	nav, err := sess.Page.GetNavigationEntry()
	check(t, err)
	if !strings.HasSuffix(nav.URL, suffix) {
		t.Fatalf("%s != %s", suffix, nav.URL)
	}
}

func TestNavigateHistory(t *testing.T) {
	t.Parallel()

	chrome, err := chrome.New("--disable-popup-blocking", "--headless")
	check(t, err)

	defer chrome.Close()
	sess, err := chrome.CDP.DefaultSession()
	check(t, err)

	check(t, sess.Page.Navigate(getFilepath("navigate.html")))

	link(t, sess, "[id='a1']")
	link(t, sess, "[id='a2']")
	link(t, sess, "[id='a3']")
	checkNav(t, sess, "#nav3")

	check(t, sess.Page.Back())
	checkNav(t, sess, "#nav2")

	check(t, sess.Page.Back())
	checkNav(t, sess, "#nav1")

	check(t, sess.Page.Back())
	checkNav(t, sess, "navigate.html")

	check(t, sess.Page.Forward())
	checkNav(t, sess, "#nav1")

	check(t, sess.Page.Forward())
	checkNav(t, sess, "#nav2")

	check(t, sess.Page.Forward())
	checkNav(t, sess, "#nav3")

	check(t, sess.Page.Forward())
	checkNav(t, sess, "#nav3")
}
