package uuid

import (
	"sync"
	"time"
)

const (
	twepoch            = int64(1648189834236)        // 开始时间戳 2022-03-25T06:30:34.236Z
	workerIdBits       = 10                          // 机器码所占字节数
	sequenceBits       = 12                          // 序列号所占字节数
	maxWorkerId        = -1 ^ (-1 << workerIdBits)   // 最大机器id
	workerIdShift      = sequenceBits                // 机器码左移位数
	timestampLeftShift = sequenceBits + workerIdBits // 时间戳左移位数
	sequenceMask       = -1 ^ (-1 << sequenceBits)   // 序列号掩码
	// maxNextIdsNum      = 1024                        // 最大生成id数量
)

type SnowflakeAlgorithm struct {
	sync.Mutex
	timestamp int64
	workerId  int64
	sequence  int64
}

func (s *SnowflakeAlgorithm) Generate() int64 {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now
	return (now-twepoch)<<timestampLeftShift | (s.workerId << workerIdShift) | s.sequence
}

func Snowflake(workerId int64) int64 {
	if workerId < 0 || workerId > maxWorkerId {
		workerId = 0
	}

	sf := &SnowflakeAlgorithm{
		workerId: workerId,
	}

	return sf.Generate()
}
