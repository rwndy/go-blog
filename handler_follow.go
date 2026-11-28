package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rwndy/go-blog/internal/database"
)

func handleFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	feedUrl := cmd.args[0]

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

func handleFollowing(s *state, cmd command, user database.User) error {

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

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	feedUrl := cmd.args[0]

	currentFeed, err := s.db.GetFeedByURL(context.Background(), feedUrl)

	if err != nil {
		return fmt.Errorf("couldn't get the feed: %w", err)
	}

	rows, err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: currentFeed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't unfollow the feed: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("%s is not following %s", user.Name, currentFeed.Name)
	}

	fmt.Printf("%s unfollowed %s\n", user.Name, currentFeed.Name)

	return nil

}
