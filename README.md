# 🕸️ WebScrapper

A high-performance, concurrent web crawler built with Go. Designed for efficiency and scalability, it leverages **Redis** for distributed task management and **Bloom Filters** for memory-efficient URL deduplication.

---

## 🚀 Features

- **Concurrent Architecture**: Utilizes Go's goroutines with a scalable worker/crawler model to maximize throughput.
- **Distributed Queue**: Uses Redis to manage the URL queue, allowing for horizontal scaling and persistence.
- **Smart Deduplication**: Implements Redis-backed Bloom Filters to ensure no URL is crawled twice, even across millions of links.
- **Robots.txt Compliance**: Automatically identifies and respects `robots.txt` directives to ensure polite crawling.
- **Real-time Monitoring**: Built-in benchmarking tool to track Pages Per Second (PPS) and discovery rates.

---

## 🛠️ Tech Stack

- **Language**: [Go](https://go.dev/) (v1.26.1+)
- **Database**: [Redis](https://redis.io/) (with RedisBloom module)
- **HTML Parsing**: [GoQuery](https://github.com/PuerkitoBio/goquery)
- **Environment Management**: [Godotenv](https://github.com/joho/godotenv)

---

## 📦 Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/0Blackraven/webscrapper.git
   cd limiter
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**:
   Create a `.env` file in the root directory:
   ```env
   REDIS_ADDR=your-redis-address:port
   REDIS_PASS=your-redis-password
   ```

---

## 🖥️ Usage

### Running the Crawler
Start the crawler by specifying a target URL, the number of concurrent workers (link processors), and common crawlers (HTTP fetchers):

```bash
go run main.go -u "https://example.com" -w 10 -c 5
```

**Flags:**
- `-u`: The starting URL to crawl.
- `-w`: Number of worker goroutines for link processing.
- `-c`: Number of crawler goroutines for HTTP fetching.

### Running the Benchmark
To test the performance of your setup:

```bash
go run real_benchmark.go -u "https://books.toscrape.com/" -w 20 -c 20 -d 60
```

---

## 🏗️ Project Structure

- `main.go`: Application entry point.
- `crawler/`: Core crawling logic and worker management.
- `utils/`: 
  - `redis/`: Redis connection and Bloom Filter implementation.
  - `htmlParser.go`: Logic for extracting links from HTML.
  - `queueManager.go`: Redis-backed job queue.
  - `robotResolver.go`: Robots.txt parsing.
- `real_benchmark.go`: Performance testing tool.

---

Basic ref point := https://github.com/tonywangcn/distributed-web-crawler/blob/master/go/src/crawler/crawler.go