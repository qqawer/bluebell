package redis

const (
	KeyPrefix              = "webapp:"
	KeyPostTimeZSet        = "post:time"//zest;帖子及发帖时间
	KeyPostScoreZset       = "post:score"//zest;帖子及投票的分数
	KeyPostVotedASetSuffix = "post:voted:"//zest;记录用及投票类型;参数是post id
)
