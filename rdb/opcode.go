package rdb

const (
	RDB_OPCODE_IDLE = iota + 248
	RDB_OPCODE_FREQ
	RDB_OPCODE_AUX
	RDB_OPCODE_RESIZEDB
	RDB_OPCODE_EXPIRETIME_MS
	RDB_OPCODE_EXPIRETIME
	RDB_OPCODE_SELECTDB
	RDB_OPCODE_EOF
)

const (
	rdbExpireTimeLen   = 8
	rdbMagicVersionLen = 4
	rdbLfuLen          = 1
)
