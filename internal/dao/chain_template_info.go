// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"llmfarm/internal/dao/internal"
)

// internalChainTemplateInfoDao is internal type for wrapping internal DAO implements.
type internalChainTemplateInfoDao = *internal.ChainTemplateInfoDao

// chainTemplateInfoDao is the data access object for table chainTemplateInfo.
// You can define custom methods on it to extend its functionality as you wish.
type chainTemplateInfoDao struct {
	internalChainTemplateInfoDao
}

var (
	// ChainTemplateInfo is globally public accessible object for table chainTemplateInfo operations.
	ChainTemplateInfo = chainTemplateInfoDao{
		internal.NewChainTemplateInfoDao(),
	}
)

// Fill with you ideas below.
