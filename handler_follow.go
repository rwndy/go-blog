package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rwndy/go-blog/internal/database"
)

func handleFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	feedUrl := cmd.args[0]

	currentUser := s.cfg.CurrentUserName

	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)

	if err != nil {
		return fmt.Errorf("couldn't get the feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't follow the feed: %w", err)
	}
	fmt.Printf("%s is now following %s\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handleFollowing(s *state, cmd command) error {
	username := s.cfg.CurrentUserName

	user, err := s.db.GetUser(context.Background(), username)

	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("couldn't get follow: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("You're not following any feeds yet.")
		return nil
	}

	fmt.Printf("Feeds followed by %s:\n", user.Name)

	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
