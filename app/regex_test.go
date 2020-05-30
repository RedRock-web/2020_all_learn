//@program: work
//@description: 
//@author: edte
//@create: 2020-05-30 19:26
package app

import (
	"testing"
	"work/database/dmysql"
)

func TestSaveNames(t *testing.T) {
	dmysql.Start()
	GetAllNames()
}
