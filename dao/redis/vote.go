package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200  --&gt; 200张赞成票可以给你的帖子续一天

/*
	 投票的几种情况：
	   direction=1时，有两种情况：
	   	1. 之前没有投过票，现在投赞成票    --&gt; 更新分数和投票记录  差值的绝对值：1  +432
	   	2. 之前投反对票，现在改投赞成票    --&gt; 更新分数和投票记录  差值的绝对值：2  +432*2
	   direction=0时，有两种情况：
	   	1. 之前投过反对票，现在要取消投票  --&gt; 更新分数和投票记录  差值的绝对值：1  +432
		2. 之前投过赞成票，现在要取消投票  --&gt; 更新分数和投票记录  差值的绝对值：1  -432
	   direction=-1时，有两种情况：
	   	1. 之前没有投过票，现在投反对票    --&gt; 更新分数和投票记录  差值的绝对值：1  -432
	   	2. 之前投赞成票，现在改投反对票    --&gt; 更新分数和投票记录  差值的绝对值：2  -432*2

	   投票的限制：
	   每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
	   	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	   	2. 到期之后删除那个 KeyPostVotedZSetPF
*/
const (
	oneWeekInStance = 7 * 24 * 3600
	scorePerVote    = 432 //每一票值多少分
)

func CreatePost(postID int64) error {
	pipeline := RedisClient.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	//1.判断投票限制
	postTime,err := RedisClient.ZScore(getRedisKey(KeyPostTimeZSet), postID).Result()
	if err!=nil{
		return err
	}
	if float64(time.Now().Unix())-postTime > oneWeekInStance {
		return errors.New("投票时间已过")
	}
	//2和3需要放在一个pipeline事务中操作
	//2.更新帖子的分数
	//先查询该用户之前的投票纪录
	ov := RedisClient.ZScore(getRedisKey(KeyPostVotedASetSuffix+postID), userID).Val()

	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算两次投票的差值
	pipeline:=RedisClient.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	
	//3.记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedASetSuffix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedASetSuffix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_,err=pipeline.Exec()
	return err
}
