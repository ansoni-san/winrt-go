// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package enumeration

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const SignatureDeviceInformationCustomPairing string = "rc(Windows.Devices.Enumeration.DeviceInformationCustomPairing;{85138c02-4ee6-4914-8370-107a39144c0e})"

type DeviceInformationCustomPairing struct {
	ole.IUnknown
}

func (impl *DeviceInformationCustomPairing) PairAsync(pairingKindsSupported DevicePairingKinds) (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiDeviceInformationCustomPairing))
	defer itf.Release()
	v := (*iDeviceInformationCustomPairing)(unsafe.Pointer(itf))
	return v.PairAsync(pairingKindsSupported)
}

func (impl *DeviceInformationCustomPairing) AddPairingRequested(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiDeviceInformationCustomPairing))
	defer itf.Release()
	v := (*iDeviceInformationCustomPairing)(unsafe.Pointer(itf))
	return v.AddPairingRequested(handler)
}

func (impl *DeviceInformationCustomPairing) RemovePairingRequested(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiDeviceInformationCustomPairing))
	defer itf.Release()
	v := (*iDeviceInformationCustomPairing)(unsafe.Pointer(itf))
	return v.RemovePairingRequested(token)
}

const GUIDiDeviceInformationCustomPairing string = "85138c02-4ee6-4914-8370-107a39144c0e"
const SignatureiDeviceInformationCustomPairing string = "{85138c02-4ee6-4914-8370-107a39144c0e}"

type iDeviceInformationCustomPairing struct {
	ole.IInspectable
}

type iDeviceInformationCustomPairingVtbl struct {
	ole.IInspectableVtbl

	PairAsync                               uintptr
	PairWithProtectionLevelAsync            uintptr
	PairWithProtectionLevelAndSettingsAsync uintptr
	AddPairingRequested                     uintptr
	RemovePairingRequested                  uintptr
}

func (v *iDeviceInformationCustomPairing) VTable() *iDeviceInformationCustomPairingVtbl {
	return (*iDeviceInformationCustomPairingVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iDeviceInformationCustomPairing) PairAsync(pairingKindsSupported DevicePairingKinds) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().PairAsync,
		uintptr(unsafe.Pointer(v)),     // this
		uintptr(pairingKindsSupported), // in DevicePairingKinds
		uintptr(unsafe.Pointer(&out)),  // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iDeviceInformationCustomPairing) AddPairingRequested(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddPairingRequested,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iDeviceInformationCustomPairing) RemovePairingRequested(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemovePairingRequested,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
