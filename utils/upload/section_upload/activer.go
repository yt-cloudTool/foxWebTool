package section_upload

import (
	"errors"
	"sync"
)
// 记录活动 变量
var Activer activer
// 锁
var ActiverLock sync.Mutex

// 记录活动uid 定义
type activer struct {
	member 		map[string]int64		// 活动成员 [uid]额定接收切片数
	memLimit 	int64					// 活动任务数限制 (key数)
	memCount	int64					// 每个切片大小限制 (value值)
}


// ****************************************************************************
// 初始化内容
func (this *activer) Init() {
	this.member = make(map[string]int64)
	// 默认限制
	this.memLimit = TaskLimitNum
	this.memCount = TaskCountNum
}

// 设置限制
func (this *activer) Setting(memLimit int64, memCount int64) {
	this.memLimit = memLimit
	this.memCount = memCount
}
// ****************************************************************************

// 获取member大小
func (this *activer) Len() int64 {
	ActiverLock.Lock()
	ret := int64(len(this.member))
	ActiverLock.Unlock()
	return ret
}

// 判断有无key
func (this *activer) HasKey(uid string) bool {
	ActiverLock.Lock()
	if _, ok := this.member[uid]; ok {
		ActiverLock.Unlock()
		return true
	} else {
		ActiverLock.Unlock()
		return false
	}
}
// ****************************************************************************

// 添加/修改活动uid ( uid, 初始值数量 )
func (this *activer) Set(uid string, count int64) error {
	
	ActiverLock.Lock()
	
	// 判断 count 是否超出限制
	if count > this.memCount {
		return errors.New("memCount out of range")
	}
	
	// 有此成员则直接修改
	if _, ok := this.member[uid]; ok {
		this.member[uid] = count
	} else {
	// 没有此成员先判断成员数是否会超出限制
		if (int64(len(this.member))+1) > this.memLimit {
			return errors.New("memLimit out of range")
		} else {
			this.member[uid] = count
		}
	}
	
	ActiverLock.Unlock()
	
	return nil
}

// 获取 (带判断)
func (this *activer) Get(uid string) (int64, error){
	
	ActiverLock.Lock()
	defer ActiverLock.Unlock()
	
	if v, ok := this.member[uid]; ok {
		return v, nil
	} else {
		return -1, errors.New("!key")
	}
	
}

// 获取 (不带判断)
func (this *activer) GetV(uid string) int64 {
	ActiverLock.Lock()
	ret := this.member[uid]
	ActiverLock.Unlock()
	
	return ret
}

// 删除
func (this *activer) Del(uid string) {
	ActiverLock.Lock()
	delete(this.member, uid)
	ActiverLock.Unlock()
}