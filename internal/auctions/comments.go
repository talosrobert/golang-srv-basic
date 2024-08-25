package auctions

import (
	"time"
)

type ItemComment struct {
	CommentText   string
	CommentSentBy *AuctionUser
	CommentSentAt time.Time
	ParentComment *ItemComment
	ChildComments []*ItemComment
}
