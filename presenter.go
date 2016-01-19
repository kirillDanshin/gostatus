package main

import "fmt"

// RepoFilter is a repo filter.
type RepoFilter func(r *Repo) (show bool)

// RepoPresenter is a repo presenter.
// All implementations must be read-only and safe for concurrent execution.
type RepoPresenter func(r *Repo) string

// PorcelainPresenter is a simple porcelain repo presenter to humans.
var PorcelainPresenter RepoPresenter = func(r *Repo) string {
	if r.vcs == nil {
		// Go package not under VCS.
		return "????" + " " + r.Root
	}

	var s string
	if r.Local.Branch != r.DefaultBranch {
		s += "b"
	} else {
		s += " "
	}
	if r.Local.Status != "" {
		s += "*"
	} else {
		s += " "
	}
	if r.Remote.Revision == "" {
		s += "!"
	} else if !*fFlag && (r.Local.RemoteURL != r.Remote.RepoURL) {
		s += "#"
	} else if r.Local.Revision != r.Remote.Revision {
		if !r.LocalContainsRemoteRevision {
			s += "+"
		} else {
			s += "-"
		}
	} else {
		s += " "
	}
	if r.Local.Stash != "" {
		s += "$"
	} else {
		s += " "
	}
	s += " " + r.Root + "/..."
	return s
}

// DebugPresenter produces debug output.
var DebugPresenter RepoPresenter = func(r *Repo) string {
	var s string
	s += fmt.Sprintf("Path=%q", r.Path)
	s += fmt.Sprintf("\tRoot=%q", r.Root)
	if r.VCS == nil {
		s += fmt.Sprintf("\tVCS=%v", r.VCS)
		return s
	}
	s += fmt.Sprintf("\tDefaultBranch=%q", r.VCS.DefaultBranch())
	s += fmt.Sprintf("\tLocal.Status=%q", r.Local.Status)
	s += fmt.Sprintf("\tLocal.Branch=%q", r.Local.Branch)
	s += fmt.Sprintf("\tLocal.Revision=%q", r.Local.Revision)
	s += fmt.Sprintf("\tLocal.Stash=%q", r.Local.Stash)
	s += fmt.Sprintf("\tLocal.RemoteURL=%q", r.Local.RemoteURL)
	s += fmt.Sprintf("\tRemote.RepoURL=%q", r.Remote.RepoURL)
	s += fmt.Sprintf("\tRemote.Revision=%q", r.Remote.Revision)
	s += fmt.Sprintf("\tLocalContainsRemoteRevision=%v", r.LocalContainsRemoteRevision)
	return s
}