⸻

🐊 Gator — RSS Feed Aggregator CLI

Gator is a command-line RSS feed aggregator written in Go.
It allows users to register, follow RSS feeds, aggregate posts in the background, and browse content directly from the terminal.

This project was built as part of the Boot.dev backend curriculum, and demonstrates:
	•	SQL migrations (Goose)
	•	Type-safe queries (SQLC)
	•	Middleware patterns
	•	Background workers
	•	CLI design
	•	PostgreSQL integration

⸻

🚀 Features
	•	User registration & login
	•	Add & follow RSS feeds
	•	Periodic feed aggregation
	•	Persistent post storage
	•	Browse posts from followed feeds
	•	Middleware-based authentication
	•	Fully CLI-driven workflow

⸻

📦 Requirements

Before running Gator, make sure you have:
	•	Go 1.21+
	•	PostgreSQL 14+
	•	Git

⸻

🛠 Installation

1. Clone the repository

git clone https://github.com/YOUR_USERNAME/gator.git
cd gator


⸻

2. Install dependencies

go mod tidy


⸻

3. Create the database

CREATE DATABASE gator;


⸻

4. Set up migrations

Run all migrations:

goose -dir migrations up


⸻

⚙️ Configuration

Gator uses a config file located at:

~/.gatorconfig.json

Create it manually:

{
  "db_url": "postgres://postgres:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}


⸻

🧪 Development Usage

During development, you can run:

go run .

But in production you should use:

go build
./gator

or

go install
gator


⸻

📌 Commands

🔐 User Management

Register a user

gator register alice

Login

gator login alice


⸻

📰 Feeds

Add a feed

gator addfeed "Hacker News" https://hnrss.org/newest

Follow a feed

gator follow https://hnrss.org/newest

Unfollow a feed

gator unfollow https://hnrss.org/newest

List followed feeds

gator following


⸻

🔄 Aggregation

Start the background feed fetcher:

gator agg 30s

This will:
	•	Fetch RSS feeds every 30 seconds
	•	Store new posts
	•	Avoid duplicates
	•	Respect feed update timing

Stop it anytime with Ctrl+C.

⸻

📚 Browse Posts

View recent posts from followed feeds:

gator browse

Limit results:

gator browse 10


⸻

🧠 Architecture Overview
	•	Postgres – persistent storage
	•	Goose – schema migrations
	•	SQLC – type-safe SQL queries
	•	Middleware – auth handling
	•	Ticker-based worker – feed scraping
	•	CLI-first design

⸻

🏗 Project Structure

.
├── cmd/                # CLI commands
├── internal/database   # SQLC generated code
├── migrations/         # Goose migrations
├── sql/                # SQL definitions
├── scrape.go           # Feed aggregation logic
├── middleware.go       # Auth middleware
├── main.go             # CLI entrypoint
└── README.md


⸻

🧪 Example Workflow

gator register alice
gator login alice
gator addfeed "HN" https://hnrss.org/newest
gator agg 1m

In another terminal:

gator browse


⸻

🧠 Notes
	•	Duplicate posts are automatically ignored
	•	Feeds are fetched in a round-robin fashion
	•	Safe to stop/restart aggregation
	•	Designed to scale cleanly

⸻

📜 License

MIT — use freely, modify freely.

⸻

🙌 Credits

Built as part of the Boot.dev Backend Developer Path
Designed and implemented by YOU 🚀

⸻

✅ Next Steps (Optional)
	•	Add pagination
	•	Add search
	•	Add web API
	•	Add concurrent workers
	•	Add tests

⸻