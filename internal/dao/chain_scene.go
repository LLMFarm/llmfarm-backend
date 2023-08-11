// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"llmfarm/internal/dao/internal"
)

// internalChainSceneDao is internal type for wrapping internal DAO implements.
type internalChainSceneDao = *internal.ChainSceneDao

// chainSceneDao is the data access object for table chainScene.
// You can define custom methods on it to extend its functionality as you wish.
type chainSceneDao struct {
	internalChainSceneDao
}

var (
	// ChainScene is globally public accessible object for table chainScene operations.
	ChainScene = chainSceneDao{
		internal.NewChainSceneDao(),
	}
)

// Fill with you ideas below.