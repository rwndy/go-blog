package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rwndy/go-blog/internal/database"
)

func handleFeedRSS(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("%+v\n", feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}
	name := cmd.args[0]
	url := cmd.args[1]
	id := uuid.New()
	now := time.Now()

	// TODO 1: ambil current user dari DB
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		return fmt.Errorf("couldn't get the user: %w", err)
	}

	// TODO 2: panggil CreateFeed
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        id,
		CreatedAt: now.UTC(),
		UpdatedAt: now.UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	// TODO 3: print fields feed baru
	fmt.Println("Feed created successfully:")
	fmt.Printf("* ID:        %s\n", feed.ID)
	fmt.Printf("* CreatedAt: %v\n", feed.CreatedAt)
	fmt.Printf("* UpdatedAt: %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:      %s\n", feed.Name)
	fmt.Printf("* URL:       %s\n", feed.Url)
	fmt.Printf("* UserID:    %s\n", feed.UserID)

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* Name: %s\n", feed.Name)
		fmt.Printf("* URL:  %s\n", feed.Url)
		fmt.Printf("* User: %s\n", feed.UserName)
		fmt.Println()
	}

	return nil
}
