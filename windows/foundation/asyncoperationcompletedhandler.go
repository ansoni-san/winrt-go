// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package foundation

import (
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/go-ole/go-ole"
)

/*
#include <stdint.h>

// Note: these functions have a different signature but because they are only
// used as function pointers (and never called) and because they use C name
// mangling, the signature doesn't really matter.
void winrt_AsyncOperationCompletedHandler_Invoke(void);
void winrt_AsyncOperationCompletedHandler_QueryInterface(void);
uint64_t winrt_AsyncOperationCompletedHandler_AddRef(void);
uint64_t winrt_AsyncOperationCompletedHandler_Release(void);

// The Vtable structure for WinRT AsyncOperationCompletedHandler interfaces.
typedef struct {
	void *QueryInterface;
	void *AddRef;
	void *Release;
	void *Invoke;
} AsyncOperationCompletedHandlerVtbl_t;

// The Vtable itself. It can be kept constant.
static const AsyncOperationCompletedHandlerVtbl_t winrt_AsyncOperationCompletedHandlerVtbl = {
	(void*)winrt_AsyncOperationCompletedHandler_QueryInterface,
	(void*)winrt_AsyncOperationCompletedHandler_AddRef,
	(void*)winrt_AsyncOperationCompletedHandler_Release,
	(void*)winrt_AsyncOperationCompletedHandler_Invoke,
};

// A small helper function to get the Vtable.
const AsyncOperationCompletedHandlerVtbl_t * winrt_getAsyncOperationCompletedHandlerVtbl(void) {
	return &winrt_AsyncOperationCompletedHandlerVtbl;
}
*/
import "C"

const GUIDAsyncOperationCompletedHandler string = "fcdcf02c-e5d8-4478-915a-4d90b74b83a5"
const SignatureAsyncOperationCompletedHandler string = "delegate({fcdcf02c-e5d8-4478-915a-4d90b74b83a5})"

type AsyncOperationCompletedHandler struct {
	ole.IUnknown
	sync.Mutex
	refs int64
	IID  ole.GUID
}

type AsyncOperationCompletedHandlerCallback func(instance *AsyncOperationCompletedHandler, asyncInfo *IAsyncOperation, asyncStatus AsyncStatus)

var callbacksAsyncOperationCompletedHandler map[unsafe.Pointer]AsyncOperationCompletedHandlerCallback = make(map[unsafe.Pointer]AsyncOperationCompletedHandlerCallback)

func NewAsyncOperationCompletedHandler(iid *ole.GUID, callback AsyncOperationCompletedHandlerCallback) *AsyncOperationCompletedHandler {
	inst := (*AsyncOperationCompletedHandler)(C.malloc(C.size_t(unsafe.Sizeof(AsyncOperationCompletedHandler{}))))
	inst.RawVTable = (*interface{})((unsafe.Pointer)(C.winrt_getAsyncOperationCompletedHandlerVtbl()))
	inst.IID = *iid // copy contents

	callbacksAsyncOperationCompletedHandler[unsafe.Pointer(inst)] = callback

	inst.addRef()
	return inst
}

// addRef increments the reference counter by one
func (r *AsyncOperationCompletedHandler) addRef() int64 {

	return atomic.AddInt64(&(r.refs), 1)
}

// removeRef decrements the reference counter by one. If it was already zero, it will just return zero.
func (r *AsyncOperationCompletedHandler) removeRef() int64 {

	return atomic.AddInt64(&(r.refs), -1)
}
