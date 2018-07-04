package errors

import (
	"errors"
)

// Different error macros
var (
	ErrVolCreateFail           = errors.New("unable to create volume")
	ErrVolNotFound             = errors.New("volume not found")
	ErrVolNotStarted           = errors.New("volume not started")
	ErrPeerNotFound            = errors.New("peer not found")
	ErrJSONParsingFailed       = errors.New("unable to parse the request")
	ErrEmptyVolName            = errors.New("volume name is empty")
	ErrInvalidVolName          = errors.New("invalid volume name")
	ErrEmptyBrickList          = errors.New("brick list is empty")
	ErrInvalidBrickPath        = errors.New("invalid brick path, brick path should be in host:<brick> format")
	ErrVolExists               = errors.New("volume already exists")
	ErrVolAlreadyStarted       = errors.New("volume already started")
	ErrVolAlreadyStopped       = errors.New("volume already stopped")
	ErrWrongGraphType          = errors.New("graph: incorrect graph type")
	ErrDeviceIDNotFound        = errors.New("failed to get device id")
	ErrBrickIsMountPoint       = errors.New("brick path is already a mount point")
	ErrBrickUnderRootPartition = errors.New("brick path is under root partition")
	ErrBrickNotDirectory       = errors.New("brick path is not a directory")
	ErrBrickPathAlreadyInUse   = errors.New("brick path is already in use by other gluster volume")
	ErrNoHostnamesPresent      = errors.New("no hostnames present")
	ErrBrickPathConvertFail    = errors.New("failed to convert the brickpath to absolute path")
	ErrBrickNotLocal           = errors.New("brickpath doesn't belong to localhost")
	ErrBrickPathTooLong        = errors.New("brickpath too long")
	ErrSubDirPathTooLong       = errors.New("sub directory path is too long")
	ErrIPAddressNotFound       = errors.New("failed to find IP address")
	ErrPeerLocalNode           = errors.New("peer being added is the local node")
	ErrProcessNotFound         = errors.New("process is not running or is inaccessible")
	ErrProcessAlreadyRunning   = errors.New("process is already running")
	ErrBitrotAlreadyEnabled    = errors.New("bitrot is already enabled")
	ErrBitrotAlreadyDisabled   = errors.New("bitrot is already disabled")
	ErrBitrotNotEnabled        = errors.New("bitrot is not enabled")
	ErrQuotadNotRunning        = errors.New("quotad is not running")
	ErrQuotadNotEnabled        = errors.New("quotad is not enabled")
	ErrUnknownValue            = errors.New("unknown value specified")
	ErrGetFailed               = errors.New("failed to get value from the store")
	ErrUnmarshallFailed        = errors.New("failed to unmarshall from json")
	ErrClusterNotFound         = errors.New("cluster instance not found in store")
	ErrDuplicateBrickPath      = errors.New("duplicate brick entry")
	ErrRestrictedKeyFound      = errors.New("key names starting with '_' are restricted in metadata field")
	ErrVolFileNotFound         = errors.New("volume file not found")
	ErrEmptySnapName           = errors.New("snapshot name is empty")
	ErrSnapExists              = errors.New("snapshot already exists")
	ErrSnapNotFound            = errors.New("snapshot not found")
	ErrInvalidVolFlags         = errors.New("invalid volume flags")
	ErrMetadataSizeOutOfBounds = errors.New("metadata size exceeds max allowed size of 4KB")
	ErrFetchingVolfileContent  = errors.New("unable to fetch volfile content")
	ErrPidFileNotFound         = errors.New("pid file not found")
)
