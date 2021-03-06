// Copyright 2018 NetApp, Inc. All Rights Reserved.

package core

import (
	"github.com/netapp/trident/config"
	"github.com/netapp/trident/frontend"
	"github.com/netapp/trident/storage"
	"github.com/netapp/trident/storage_class"
)

type Orchestrator interface {
	Bootstrap() error
	AddFrontend(f frontend.Plugin)
	GetVersion() string

	AddStorageBackend(configJSON string) (*storage.BackendExternal, error)
	GetBackend(backend string) *storage.BackendExternal
	ListBackends() []*storage.BackendExternal
	OfflineBackend(backend string) (bool, error)

	AddVolume(volumeConfig *storage.VolumeConfig) (*storage.VolumeExternal, error)
	CloneVolume(volumeConfig *storage.VolumeConfig) (*storage.VolumeExternal, error)
	GetVolume(volume string) *storage.VolumeExternal
	GetDriverTypeForVolume(vol *storage.VolumeExternal) string
	GetVolumeType(vol *storage.VolumeExternal) config.VolumeType
	ListVolumes() []*storage.VolumeExternal
	DeleteVolume(volume string) (found bool, err error)
	ListVolumesByPlugin(pluginName string) []*storage.VolumeExternal
	AttachVolume(volumeName, mountpoint string, options map[string]string) error
	DetachVolume(volumeName, mountpoint string) error
	ListVolumeSnapshots(volumeName string) ([]*storage.SnapshotExternal, error)
	ReloadVolumes() error

	AddStorageClass(scConfig *storageclass.Config) (*storageclass.External, error)
	GetStorageClass(scName string) *storageclass.External
	ListStorageClasses() []*storageclass.External
	DeleteStorageClass(scName string) (bool, error)
}
