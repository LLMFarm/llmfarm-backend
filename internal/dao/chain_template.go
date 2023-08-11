// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"llmfarm/internal/dao/internal"
)

// internalChainTemplateDao is internal type for wrapping internal DAO implements.
type internalChainTemplateDao = *internal.ChainTemplateDao

// chainTemplateDao is the data access object for table chainTemplate.
// You can define custom methods on it to extend its functionality as you wish.
type chainTemplateDao struct {
	internalChainTemplateDao
}

var (
	// ChainTemplate is globally public accessible object for table chainTemplate operations.
	ChainTemplate = chainTemplateDao{
		internal.NewChainTemplateDao(),
	}
)

// Fill with you ideas below.