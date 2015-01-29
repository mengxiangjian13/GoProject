package library

import (
	"errors"
)

// 音乐对象类
type Music struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

// 音乐管理类
type MusicManager struct {
	musics []Music
}

// 新建音乐管理类
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}

// 音乐列表长度
func (m *MusicManager) Len() int {
	return len(m.musics)
}

// 根据序列号查找音乐
func (m *MusicManager) Get(index int) (*Music, error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("index out of range in Get.")
	}
	return &m.musics[index], nil
}

// 根据名称查找音乐
func (m *MusicManager) Find(name string) *Music {
	if len(m.musics) == 0 {
		return nil
	}
	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}
	return nil
}

// 添加音乐
func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics, *music)
}

// 移除音乐
func (m *MusicManager) Remove(index int) *Music {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removeMusic := &m.musics[index]

	if index == 0 {
		// 删除第一个元素
		if len(m.musics) == 1 {
			m.musics = make([]Music, 0)
		} else {
			m.musics = m.musics[index+1:]
		}
	} else if index == len(m.musics)-1 {
		// 删除最后一个元素
		m.musics = m.musics[:len(m.musics)-1]
	} else {
		// 中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	}

	return removeMusic
}

func (m *MusicManager) RemoveByName(name string) {
	for i := 0; i < len(m.musics); i++ {
		music := m.musics[i]
		if music.Name == name {
			m.Remove(i)
			break
		}
	}
}
