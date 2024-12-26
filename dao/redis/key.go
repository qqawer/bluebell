package redis

const (
	Prefix                 = "webapp:"
	KeyPostTimeZSet        = "post:time"   //zest;帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  //zest;帖子及投票的分数
	KeyPostVotedASetSuffix = "post:voted:" //zest;记录用及投票类型;参数是post id
	KeyCommunitySetSuffix  = "community:"  //set;保存每个分区下帖子的id
)

//给redis key 加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
