package lock

import (
	"github.com/go-redis/redis"
	"time"
)
/**
{
	CAP理论： 一致性，可用性，分区容错性

	在分布式系统环境下，一个方法在同一时间只能被一个机器的一个线程执行
	高可用的获取锁与释放锁
	高性能的获取锁与释放锁
	具备可重入特性
		重进入是指任意线程在获取到锁之后能够再次获取该锁而不会被锁阻塞，该特性的实现需要解决以下两个问题：
		线程再次获取锁：锁需要去识别获取锁的线程是否为当前占据锁的线程，如果是，则再次成功获取。
		锁的最终释放。线程重复 n 次获取了锁，随后在第 n 次释放该锁后，其它线程能够获取到该锁。锁的最终释放要求锁对于获取进行计数自增，计数表示当前锁被重复获取的次数，而锁被释放时，计数自减，当计数等于 0 时表示锁已经成功释放。
	具备锁失效机制，防止死锁
	具备非阻塞锁特性，即没有获取到锁将直接返回获取锁失败
}

{
	数据库实现分布式锁
	version
	unique唯一键
}

{
	redis实现分布式锁
	lock： set(key, value, EX, 过期时间)
	unlock 判断当前锁，del(key)


	注：
	setNx和expired原子性操作
	设置过期时间
	del误删，设置一个唯一串
	Get/判断/del 保证原子性（lua脚本， 增加守护进程）
}
*/


func Lock(redisDB *redis.Client, lockKey string, duration time.Duration)(bool , error){
	return redisDB.SetNX(lockKey, "", duration).Result()
}

func UnLock(redisDB *redis.Client, lockKey string)(int64, error) {
	return redisDB.Del(lockKey).Result()
}
