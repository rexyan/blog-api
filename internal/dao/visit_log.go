// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog-api/internal/dao/internal"
)

// internalVisitLogDao is internal type for wrapping internal DAO implements.
type internalVisitLogDao = *internal.VisitLogDao

// visitLogDao is the data access object for table visit_log.
// You can define custom methods on it to extend its functionality as you wish.
type visitLogDao struct {
	internalVisitLogDao
}

var (
	// VisitLog is globally public accessible object for table visit_log operations.
	VisitLog = visitLogDao{
		internal.NewVisitLogDao(),
	}
)

// Fill with you ideas below.
