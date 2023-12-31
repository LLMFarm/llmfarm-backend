// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"llmfarm/internal/dao/internal"
)

// internalChainTemplateSceneDao is internal type for wrapping internal DAO implements.
type internalChainTemplateSceneDao = *internal.ChainTemplateSceneDao

// chainTemplateSceneDao is the data access object for table chainTemplateScene.
// You can define custom methods on it to extend its functionality as you wish.
type chainTemplateSceneDao struct {
	internalChainTemplateSceneDao
}

var (
	// ChainTemplateScene is globally public accessible object for table chainTemplateScene operations.
	ChainTemplateScene = chainTemplateSceneDao{
		internal.NewChainTemplateSceneDao(),
	}
)

// Fill with you ideas below.
