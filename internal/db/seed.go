package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/The-yohanna/social/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "david", "emma", "frank", "grace", "harry", "isla", "jack",
	"kate", "liam", "mia", "nathan", "olivia", "paul", "quinn", "rachel", "sammy", "tina",
	"ursula", "vincent", "wendy", "xander", "yara", "zack", "amy", "brian", "claire", "daniel",
	"ella", "felix", "george", "hannah", "ian", "julia", "kevin", "lily", "mark", "nina",
	"oscar", "penny", "quincy", "ruby", "steve", "tommy", "una", "victor", "willow", "zoe",
}

var titles = []string{
	"The Power of Daily Habits",
	"Learning to Embrace Change",
	"How to Cultivate Inner Peace",
	"Breaking Free from Self-Doubt",
	"The Art of Saying No",
	"Creating a Life with Purpose",
	"Finding Joy in the Little Things",
	"Why Rest Is Productive",
	"Mastering the Growth Mindset",
	"Letting Go of Perfectionism",
	"Small Steps to Big Changes",
	"Declutter Your Mind, Not Just Your Home",
	"Living with Intention Every Day",
	"Turning Setbacks into Strength",
	"Building Mental Resilience",
	"How to Stay Grounded in a Noisy World",
	"The Psychology of Self-Discipline",
	"Choosing Simplicity in a Complex World",
	"How to Reconnect with Yourself",
	"The Power of Gratitude in Daily Life",
}

var contents = []string{
	"In this post, we'll explore how to develop good habits that stick and transform your life.",
	"Discover the benefits of a minimalist lifestyle and how to declutter your home and mind.",
	"Learn practical tips for eating healthy on a budget without sacrificing flavor.",
	"Traveling doesn't have to be expensive. Here are some tips for seeing the world on a budget.",
	"Mindfulness meditation can reduce stress and improve your mental well-being. Here's how to get started.",
	"Increase your productivity with these simple and effective strategies.",
	"Set up the perfect home office to boost your work-from-home efficiency and comfort.",
	"A digital detox can help you reconnect with the real world and improve your mental health.",
	"Start your gardening journey with these basic tips for beginners.",
	"Transform your home with these fun and easy DIY projects.",
	"Yoga is a great way to stay fit and flexible. Here are some beginner-friendly poses to try.",
	"Sustainable living is good for you and the planet. Learn how to make eco-friendly choices.",
	"Master time management with these tips and get more done in less time.",
	"Nature has so much to offer. Discover the benefits of spending time outdoors.",
	"Whip up delicious meals with these simple and quick cooking recipes.",
	"Stay fit without leaving home with these effective at-home workout routines.",
	"Take control of your finances with these practical personal finance tips.",
	"Unleash your creativity with these inspiring writing prompts and exercises.",
	"Mental health is just as important as physical health. Learn how to take care of your mind.",
	"Learning new skills can be fun and rewarding. Here are some ideas to get you started.",
}

var tags = []string{
	"Self Improvement", "Minimalism", "Health", "Travel", "Mindfulness",
	"Productivity", "Home Office", "Digital Detox", "Gardening", "DIY",
	"Yoga", "Sustainability", "Time Management", "Nature", "Cooking",
	"Fitness", "Personal Finance", "Writing", "Mental Health", "Learning",
}

var comments = []string{
	"Great post! Thanks for sharing.",
	"I completely agree with your thoughts.",
	"Thanks for the tips, very helpful.",
	"Interesting perspective, I hadn't considered that.",
	"Thanks for sharing your experience.",
	"Well written, I enjoyed reading this.",
	"This is very insightful, thanks for posting.",
	"Great advice, I'll definitely try that.",
	"I love this, very inspirational.",
	"Thanks for the information, very useful.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)

	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}

	_ = tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Post creation failed:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Comment creation failed:", err)
			return
		}
	}

	log.Println("Seeding complete.")
}

func generateUsers(count int) []*store.User {
	users := make([]*store.User, count)

	for i := 0; i < count; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Role: store.Role{
				Name: "user",
			},
		}
	}

	return users
}

func generatePosts(count int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, count)

	for i := 0; i < count; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: titles[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(count int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, count)

	for i := 0; i < count; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: titles[rand.Intn(len(comments))],
		}
	}

	return cms
}
