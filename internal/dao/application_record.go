// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"llmfarm/internal/dao/internal"
)

// internalApplicationRecordDao is internal type for wrapping internal DAO implements.
type internalApplicationRecordDao = *internal.ApplicationRecordDao

// applicationRecordDao is the data access object for table applicationRecord.
// You can define custom methods on it to extend its functionality as you wish.
type applicationRecordDao struct {
	internalApplicationRecordDao
}

var (
	// ApplicationRecord is globally public accessible object for table applicationRecord operations.
	ApplicationRecord = applicationRecordDao{
		internal.NewApplicationRecordDao(),
	}
)

// Fill with you ideas below.
