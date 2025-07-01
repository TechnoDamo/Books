# Table of Contents
1. [Scale from Zero to Million Users](#1-scale-from-zero-to-million-users)
2. [Back-of-the-Envelope Estimation](#2-back-of-the-envelope-estimation)
3. [A Framework for System Design Interviews](#3-a-framework-for-system-design-interviews)
4. [Design a Rate Limiter](#4-design-a-rate-limiter)
5. [Design Consistent Hashing](#5-design-consistent-hashing)
6. [Design a Key-Value Store](#6-design-a-key-value-store)
7. [Design a Unique ID Generator in Distributed Systems](#7-design-a-unique-id-generator-in-distributed-systems)
8. [Design a URL Shortener](#8-design-a-url-shortener)
9. [Design a Web Crawler](#9-design-a-web-crawler)
10. [Design a Notification System](#10-design-a-notification-system)
11. [Design a News Feed System](#11-design-a-news-feed-system)
12. [Design a Chat System](#12-design-a-chat-system)
13. [Design a Search Autocomplete System](#13-design-a-search-autocomplete-system)
14. [Design YouTube](#14-design-youtube)
15. [Design Google Drive](#15-design-google-drive)
16. [Further Studies](#16-further-studies)


---

# 1. Scale from Zero to Million Users
## Single Server Setup
## Database
## Choosing a Database
## Vertical Scaling vs Horizontal Scaling
## Load Balancer
## Database Replication
## Cache Tier
## Considerations for Using Cache
## Content Delivery Network (CDN)
## Considerations for Using a CDN
## Stateless Web Tier
## Stateful Architecture
## Stateless Architecture
## Data Centers
## Message Queue
## Logging, Metrics, Automation
## Database Scaling

# 2. Back-of-the-Envelope Estimation
## Power of two
## Latency number 
## Availability numbers

# 3. A Framework for System Design Interviews
## Efficient 4-step process 

# 4. Design a Rate Limiter
## Algorithms for rate limiting
### Token bucket algorithm
### Leaking bucket algorithm
### Fixed window counter algorithm
### Sliding window log algorithm
### Sliding window counter algorithm
## High-level architechture
## Rate limiting rules
## Exceeding the rate limit
## Rate limiter headers
## Detailed design
## Rate limiter in a distributed environment
### Race condition
### Sync issue
### Perfomance optimization
## Monitoring

# 5. Design Consistent Hashing
## The rehashing problem
## Consistent hashing
## Hash space and hash ring
## Hash servers
## Hash keys
## Server lookup
## Add a server
## Remove a server
## Two issues in the basic approach
## Virtual nodes
## Find affected keys

# 6. Design a Key-Value Store
## Single server key-value store
## Distributed key-values store
## CAP theorem
### Ideal situation
### Real-world distributed systems
## System components
## Data partition
## Data replication
## Consistency
### Consistency models
## Inconsistency resolution: versioning
## Handling failures
### Failure detection
### Handling temporary failures 
### Handling permanent failures
### Handling data center outage
### Write path
### Read path



# 7. Design a Unique ID Generator in Distributed Systems
## Multi-master replication
## UUID
## Ticket Server
## Twitter snowflake approach
## Deep dive
### Timestamp
### Sequence number

# 8. Design a URL Shortener
## API Endpoints
## URL redirecting
## URL shortening
## Data model
## Hash function
### Hash value length
### Hash + collision resolution
### Base 62 conversion
## URL shortening deep dive
## URL redirecting deep dive



# 9. Design a Web Crawler
### Seed URLs
### URL Frontier
### HTML Downloader
### DNS Resolver
### Content Parser
### Content seen?
### Content Storage
### URL Extractor
### URL Filter
### URL Seen?
### URL Storage
### Web crawler worlflow
## DFS vs BFS
## URL frontier
### Politeness
### Priority
### Freshness
### Storage for URL Frontier
## HTML Downloader
## Perfomance optimisation
### 1. Distributed crawl
### 2. Cache DNS Resolver
### 3. Locality
### 4. Short timeout
## Robustness
## Extensibility
## Detect andd avoid problematic content
### 1. Redundant content
### 2. Spider traps 
### 3. Data noise

# 10. Design a Notification System
## Different types of notifications
### iOS push notifications
### Android push notifications
### SMS message
### Email
## Contact info gathering flow
## Notification sending/receiving flow
## High-level design
### Service 1 to N
### Notification system
### Third-party services
### iOS, Android, SMS, Email
### High-level design (improved)
...
## Reliability
### Data loss prevention
### Will recipients receive a notification exactly once?
## Additional components and considerations
### Notification template
### Notification setting
### Rate limiting
### Retry mechanism 
### Security in push notifications
### Monitor queued notifications
### Events tracking



# 11. Design a News Feed System


# 12. Design a Chat System
Content for designing a chat system.

---

# 13. Design a Search Autocomplete System
Content for designing a search autocomplete system.

---

# 14. Design YouTube
Content for designing YouTube.

---

# 15. Design Google Drive
Content for designing Google Drive.

---

# 16. Further Studies
Content for further studies in system design.