// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog-api/internal/dao/internal"
)

// internalCityVisitorDao is internal type for wrapping internal DAO implements.
type internalCityVisitorDao = *internal.CityVisitorDao

// cityVisitorDao is the data access object for table city_visitor.
// You can define custom methods on it to extend its functionality as you wish.
type cityVisitorDao struct {
	internalCityVisitorDao
}

var (
	// CityVisitor is globally public accessible object for table city_visitor operations.
	CityVisitor = cityVisitorDao{
		internal.NewCityVisitorDao(),
	}
)

// Fill with you ideas below.
