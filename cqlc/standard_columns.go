// THIS FILE WAS AUTOGENERATED - ANY EDITS TO THIS WILL BE LOST WHEN IT IS REGENERATED

package cqlc

import (
	"github.com/gocql/gocql"
	"speter.net/go/exp/math/dec/inf"
	"time"
)



type StringColumn interface {
	Column
	To(value *string) ColumnBinding
}

type EqualityStringColumn interface {
	StringColumn
	Eq(value string) Condition
}

type PartitionedStringColumn interface {
	PartitionedColumn
	EqualityStringColumn
}

type LastPartitionedStringColumn interface {
	PartitionedStringColumn
	In(value ...string) Condition
}

type ClusteredStringColumn interface {
	ClusteredColumn
	EqualityStringColumn
	Gt(value string) Condition
	Lt(value string) Condition
	Ge(value string) Condition
	Le(value string) Condition
}

type LastClusteredStringColumn interface {
	ClusteredStringColumn
	In(value ...string) Condition
}



type Int32Column interface {
	Column
	To(value *int32) ColumnBinding
}

type EqualityInt32Column interface {
	Int32Column
	Eq(value int32) Condition
}

type PartitionedInt32Column interface {
	PartitionedColumn
	EqualityInt32Column
}

type LastPartitionedInt32Column interface {
	PartitionedInt32Column
	In(value ...int32) Condition
}

type ClusteredInt32Column interface {
	ClusteredColumn
	EqualityInt32Column
	Gt(value int32) Condition
	Lt(value int32) Condition
	Ge(value int32) Condition
	Le(value int32) Condition
}

type LastClusteredInt32Column interface {
	ClusteredInt32Column
	In(value ...int32) Condition
}



type Int64Column interface {
	Column
	To(value *int64) ColumnBinding
}

type EqualityInt64Column interface {
	Int64Column
	Eq(value int64) Condition
}

type PartitionedInt64Column interface {
	PartitionedColumn
	EqualityInt64Column
}

type LastPartitionedInt64Column interface {
	PartitionedInt64Column
	In(value ...int64) Condition
}

type ClusteredInt64Column interface {
	ClusteredColumn
	EqualityInt64Column
	Gt(value int64) Condition
	Lt(value int64) Condition
	Ge(value int64) Condition
	Le(value int64) Condition
}

type LastClusteredInt64Column interface {
	ClusteredInt64Column
	In(value ...int64) Condition
}



type Float32Column interface {
	Column
	To(value *float32) ColumnBinding
}

type EqualityFloat32Column interface {
	Float32Column
	Eq(value float32) Condition
}

type PartitionedFloat32Column interface {
	PartitionedColumn
	EqualityFloat32Column
}

type LastPartitionedFloat32Column interface {
	PartitionedFloat32Column
	In(value ...float32) Condition
}

type ClusteredFloat32Column interface {
	ClusteredColumn
	EqualityFloat32Column
	Gt(value float32) Condition
	Lt(value float32) Condition
	Ge(value float32) Condition
	Le(value float32) Condition
}

type LastClusteredFloat32Column interface {
	ClusteredFloat32Column
	In(value ...float32) Condition
}



type Float64Column interface {
	Column
	To(value *float64) ColumnBinding
}

type EqualityFloat64Column interface {
	Float64Column
	Eq(value float64) Condition
}

type PartitionedFloat64Column interface {
	PartitionedColumn
	EqualityFloat64Column
}

type LastPartitionedFloat64Column interface {
	PartitionedFloat64Column
	In(value ...float64) Condition
}

type ClusteredFloat64Column interface {
	ClusteredColumn
	EqualityFloat64Column
	Gt(value float64) Condition
	Lt(value float64) Condition
	Ge(value float64) Condition
	Le(value float64) Condition
}

type LastClusteredFloat64Column interface {
	ClusteredFloat64Column
	In(value ...float64) Condition
}



type TimestampColumn interface {
	Column
	To(value *time.Time) ColumnBinding
}

type EqualityTimestampColumn interface {
	TimestampColumn
	Eq(value time.Time) Condition
}

type PartitionedTimestampColumn interface {
	PartitionedColumn
	EqualityTimestampColumn
}

type LastPartitionedTimestampColumn interface {
	PartitionedTimestampColumn
	In(value ...time.Time) Condition
}

type ClusteredTimestampColumn interface {
	ClusteredColumn
	EqualityTimestampColumn
	Gt(value time.Time) Condition
	Lt(value time.Time) Condition
	Ge(value time.Time) Condition
	Le(value time.Time) Condition
}

type LastClusteredTimestampColumn interface {
	ClusteredTimestampColumn
	In(value ...time.Time) Condition
}



type TimeUUIDColumn interface {
	Column
	To(value *gocql.UUID) ColumnBinding
}

type EqualityTimeUUIDColumn interface {
	TimeUUIDColumn
	Eq(value gocql.UUID) Condition
}

type PartitionedTimeUUIDColumn interface {
	PartitionedColumn
	EqualityTimeUUIDColumn
}

type LastPartitionedTimeUUIDColumn interface {
	PartitionedTimeUUIDColumn
	In(value ...gocql.UUID) Condition
}

type ClusteredTimeUUIDColumn interface {
	ClusteredColumn
	EqualityTimeUUIDColumn
	Gt(value gocql.UUID) Condition
	Lt(value gocql.UUID) Condition
	Ge(value gocql.UUID) Condition
	Le(value gocql.UUID) Condition
}

type LastClusteredTimeUUIDColumn interface {
	ClusteredTimeUUIDColumn
	In(value ...gocql.UUID) Condition
}



type UUIDColumn interface {
	Column
	To(value *gocql.UUID) ColumnBinding
}

type EqualityUUIDColumn interface {
	UUIDColumn
	Eq(value gocql.UUID) Condition
}

type PartitionedUUIDColumn interface {
	PartitionedColumn
	EqualityUUIDColumn
}

type LastPartitionedUUIDColumn interface {
	PartitionedUUIDColumn
	In(value ...gocql.UUID) Condition
}

type ClusteredUUIDColumn interface {
	ClusteredColumn
	EqualityUUIDColumn
	Gt(value gocql.UUID) Condition
	Lt(value gocql.UUID) Condition
	Ge(value gocql.UUID) Condition
	Le(value gocql.UUID) Condition
}

type LastClusteredUUIDColumn interface {
	ClusteredUUIDColumn
	In(value ...gocql.UUID) Condition
}



type BooleanColumn interface {
	Column
	To(value *bool) ColumnBinding
}

type EqualityBooleanColumn interface {
	BooleanColumn
	Eq(value bool) Condition
}

type PartitionedBooleanColumn interface {
	PartitionedColumn
	EqualityBooleanColumn
}

type LastPartitionedBooleanColumn interface {
	PartitionedBooleanColumn
	In(value ...bool) Condition
}

type ClusteredBooleanColumn interface {
	ClusteredColumn
	EqualityBooleanColumn
	Gt(value bool) Condition
	Lt(value bool) Condition
	Ge(value bool) Condition
	Le(value bool) Condition
}

type LastClusteredBooleanColumn interface {
	ClusteredBooleanColumn
	In(value ...bool) Condition
}



type DecimalColumn interface {
	Column
	To(value **inf.Dec) ColumnBinding
}

type EqualityDecimalColumn interface {
	DecimalColumn
	Eq(value *inf.Dec) Condition
}

type PartitionedDecimalColumn interface {
	PartitionedColumn
	EqualityDecimalColumn
}

type LastPartitionedDecimalColumn interface {
	PartitionedDecimalColumn
	In(value ...*inf.Dec) Condition
}

type ClusteredDecimalColumn interface {
	ClusteredColumn
	EqualityDecimalColumn
	Gt(value *inf.Dec) Condition
	Lt(value *inf.Dec) Condition
	Ge(value *inf.Dec) Condition
	Le(value *inf.Dec) Condition
}

type LastClusteredDecimalColumn interface {
	ClusteredDecimalColumn
	In(value ...*inf.Dec) Condition
}



type BytesColumn interface {
	Column
	To(value *[]byte) ColumnBinding
}

type EqualityBytesColumn interface {
	BytesColumn
	Eq(value []byte) Condition
}

type PartitionedBytesColumn interface {
	PartitionedColumn
	EqualityBytesColumn
}

type LastPartitionedBytesColumn interface {
	PartitionedBytesColumn
	In(value ...[]byte) Condition
}

type ClusteredBytesColumn interface {
	ClusteredColumn
	EqualityBytesColumn
	Gt(value []byte) Condition
	Lt(value []byte) Condition
	Ge(value []byte) Condition
	Le(value []byte) Condition
}

type LastClusteredBytesColumn interface {
	ClusteredBytesColumn
	In(value ...[]byte) Condition
}




type StringSliceColumn interface {
	Column
}

type Int32SliceColumn interface {
	Column
}

type Int64SliceColumn interface {
	Column
}

type Float32SliceColumn interface {
	Column
}

type Float64SliceColumn interface {
	Column
}

type TimestampSliceColumn interface {
	Column
}

type TimeUUIDSliceColumn interface {
	Column
}

type UUIDSliceColumn interface {
	Column
}

type BooleanSliceColumn interface {
	Column
}

type DecimalSliceColumn interface {
	Column
}

type BytesSliceColumn interface {
	Column
}


type SetValueStep interface {
	Executable
	SelectWhereStep
	Apply(cols ...ColumnBinding) SetValueStep
	IfExists(cols ...ColumnBinding) CompareAndSwap

	
	SetString(col StringColumn, value string) SetValueStep
	
	SetInt32(col Int32Column, value int32) SetValueStep
	
	SetInt64(col Int64Column, value int64) SetValueStep
	
	SetFloat32(col Float32Column, value float32) SetValueStep
	
	SetFloat64(col Float64Column, value float64) SetValueStep
	
	SetTimestamp(col TimestampColumn, value time.Time) SetValueStep
	
	SetTimeUUID(col TimeUUIDColumn, value gocql.UUID) SetValueStep
	
	SetUUID(col UUIDColumn, value gocql.UUID) SetValueStep
	
	SetBoolean(col BooleanColumn, value bool) SetValueStep
	
	SetDecimal(col DecimalColumn, value *inf.Dec) SetValueStep
	
	SetBytes(col BytesColumn, value []byte) SetValueStep
	

	SetMap(col MapColumn, value map[string]string) SetValueStep
	//SetArray(col ArrayColumn, value []string) SetValueStep

	
	SetStringSlice(col StringSliceColumn, value []string) SetValueStep
	
	SetInt32Slice(col Int32SliceColumn, value []int32) SetValueStep
	
	SetInt64Slice(col Int64SliceColumn, value []int64) SetValueStep
	
	SetFloat32Slice(col Float32SliceColumn, value []float32) SetValueStep
	
	SetFloat64Slice(col Float64SliceColumn, value []float64) SetValueStep
	
	SetTimestampSlice(col TimestampSliceColumn, value []time.Time) SetValueStep
	
	SetTimeUUIDSlice(col TimeUUIDSliceColumn, value []gocql.UUID) SetValueStep
	
	SetUUIDSlice(col UUIDSliceColumn, value []gocql.UUID) SetValueStep
	
	SetBooleanSlice(col BooleanSliceColumn, value []bool) SetValueStep
	
	SetDecimalSlice(col DecimalSliceColumn, value []*inf.Dec) SetValueStep
	
	SetBytesSlice(col BytesSliceColumn, value [][]byte) SetValueStep
	
}

func (c *Context) SetMap(col MapColumn, value map[string]string) SetValueStep {
	set(c, col, value)
	return c
}

/*func (c *Context) SetArray(col ArrayColumn, value []string) SetValueStep {
	set(c, col, value)
	return c
}*/


func (c *Context) SetString(col StringColumn, value string) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetInt32(col Int32Column, value int32) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetInt64(col Int64Column, value int64) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetFloat32(col Float32Column, value float32) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetFloat64(col Float64Column, value float64) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetTimestamp(col TimestampColumn, value time.Time) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetTimeUUID(col TimeUUIDColumn, value gocql.UUID) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetUUID(col UUIDColumn, value gocql.UUID) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetBoolean(col BooleanColumn, value bool) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetDecimal(col DecimalColumn, value *inf.Dec) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetBytes(col BytesColumn, value []byte) SetValueStep {
	set(c, col, value)
	return c
}



func (c *Context) SetStringSlice(col StringSliceColumn, value []string) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetInt32Slice(col Int32SliceColumn, value []int32) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetInt64Slice(col Int64SliceColumn, value []int64) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetFloat32Slice(col Float32SliceColumn, value []float32) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetFloat64Slice(col Float64SliceColumn, value []float64) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetTimestampSlice(col TimestampSliceColumn, value []time.Time) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetTimeUUIDSlice(col TimeUUIDSliceColumn, value []gocql.UUID) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetUUIDSlice(col UUIDSliceColumn, value []gocql.UUID) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetBooleanSlice(col BooleanSliceColumn, value []bool) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetDecimalSlice(col DecimalSliceColumn, value []*inf.Dec) SetValueStep {
	set(c, col, value)
	return c
}

func (c *Context) SetBytesSlice(col BytesSliceColumn, value [][]byte) SetValueStep {
	set(c, col, value)
	return c
}

