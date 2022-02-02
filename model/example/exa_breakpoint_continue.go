package example

import "rebuildServer/global"

// file struct 文件结构体
type ExaFile struct {
	global.GVA_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

type ExaFileChunk struct {
	global.GVA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
