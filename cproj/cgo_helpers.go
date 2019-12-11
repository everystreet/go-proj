// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 11 Dec 2019 19:04:51 UTC.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package cproj

/*
#cgo LDFLAGS: -lproj
#include "proj.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// Ref returns a reference to C object as it is.
func (x *PJ_AREA) Ref() *C.PJ_AREA {
	if x == nil {
		return nil
	}
	return (*C.PJ_AREA)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *PJ_AREA) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewPJ_AREARef converts the C object reference into a raw struct reference without wrapping.
func NewPJ_AREARef(ref unsafe.Pointer) *PJ_AREA {
	return (*PJ_AREA)(ref)
}

// NewPJ_AREA allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewPJ_AREA() *PJ_AREA {
	return (*PJ_AREA)(allocPJ_AREAMemory(1))
}

// allocPJ_AREAMemory allocates memory for type C.PJ_AREA in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPJ_AREAMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPJ_AREAValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPJ_AREAValue = unsafe.Sizeof([1]C.PJ_AREA{})

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *PJ_AREA) PassRef() *C.PJ_AREA {
	if x == nil {
		x = (*PJ_AREA)(allocPJ_AREAMemory(1))
	}
	return (*C.PJ_AREA)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *PJ) Ref() *C.PJ {
	if x == nil {
		return nil
	}
	return (*C.PJ)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *PJ) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewPJRef converts the C object reference into a raw struct reference without wrapping.
func NewPJRef(ref unsafe.Pointer) *PJ {
	return (*PJ)(ref)
}

// NewPJ allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewPJ() *PJ {
	return (*PJ)(allocPJMemory(1))
}

// allocPJMemory allocates memory for type C.PJ in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPJMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPJValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPJValue = unsafe.Sizeof([1]C.PJ{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *PJ) PassRef() *C.PJ {
	if x == nil {
		x = (*PJ)(allocPJMemory(1))
	}
	return (*C.PJ)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *PJ_CONTEXT) Ref() *C.PJ_CONTEXT {
	if x == nil {
		return nil
	}
	return (*C.PJ_CONTEXT)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *PJ_CONTEXT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewPJ_CONTEXTRef converts the C object reference into a raw struct reference without wrapping.
func NewPJ_CONTEXTRef(ref unsafe.Pointer) *PJ_CONTEXT {
	return (*PJ_CONTEXT)(ref)
}

// NewPJ_CONTEXT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewPJ_CONTEXT() *PJ_CONTEXT {
	return (*PJ_CONTEXT)(allocPJ_CONTEXTMemory(1))
}

// allocPJ_CONTEXTMemory allocates memory for type C.PJ_CONTEXT in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPJ_CONTEXTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPJ_CONTEXTValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPJ_CONTEXTValue = unsafe.Sizeof([1]C.PJ_CONTEXT{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *PJ_CONTEXT) PassRef() *C.PJ_CONTEXT {
	if x == nil {
		x = (*PJ_CONTEXT)(allocPJ_CONTEXTMemory(1))
	}
	return (*C.PJ_CONTEXT)(unsafe.Pointer(x))
}

// allocPROJ_CRS_INFOMemory allocates memory for type C.PROJ_CRS_INFO in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPROJ_CRS_INFOMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPROJ_CRS_INFOValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPROJ_CRS_INFOValue = unsafe.Sizeof([1]C.PROJ_CRS_INFO{})

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PROJ_CRS_INFO) Ref() *C.PROJ_CRS_INFO {
	if x == nil {
		return nil
	}
	return x.ref48245dab
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PROJ_CRS_INFO) Free() {
	if x != nil && x.allocs48245dab != nil {
		x.allocs48245dab.(*cgoAllocMap).Free()
		x.ref48245dab = nil
	}
}

// NewPROJ_CRS_INFORef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPROJ_CRS_INFORef(ref unsafe.Pointer) *PROJ_CRS_INFO {
	if ref == nil {
		return nil
	}
	obj := new(PROJ_CRS_INFO)
	obj.ref48245dab = (*C.PROJ_CRS_INFO)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PROJ_CRS_INFO) PassRef() (*C.PROJ_CRS_INFO, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref48245dab != nil {
		return x.ref48245dab, nil
	}
	mem48245dab := allocPROJ_CRS_INFOMemory(1)
	ref48245dab := (*C.PROJ_CRS_INFO)(mem48245dab)
	allocs48245dab := new(cgoAllocMap)
	allocs48245dab.Add(mem48245dab)

	var cauth_name_allocs *cgoAllocMap
	ref48245dab.auth_name, cauth_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.auth_name)).Data)), cgoAllocsUnknown
	allocs48245dab.Borrow(cauth_name_allocs)

	var ccode_allocs *cgoAllocMap
	ref48245dab.code, ccode_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.code)).Data)), cgoAllocsUnknown
	allocs48245dab.Borrow(ccode_allocs)

	var cname_allocs *cgoAllocMap
	ref48245dab.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs48245dab.Borrow(cname_allocs)

	var cdeprecated_allocs *cgoAllocMap
	ref48245dab.deprecated, cdeprecated_allocs = (C.int)(x.deprecated), cgoAllocsUnknown
	allocs48245dab.Borrow(cdeprecated_allocs)

	var cbbox_valid_allocs *cgoAllocMap
	ref48245dab.bbox_valid, cbbox_valid_allocs = (C.int)(x.bbox_valid), cgoAllocsUnknown
	allocs48245dab.Borrow(cbbox_valid_allocs)

	var cwest_lon_degree_allocs *cgoAllocMap
	ref48245dab.west_lon_degree, cwest_lon_degree_allocs = (C.double)(x.west_lon_degree), cgoAllocsUnknown
	allocs48245dab.Borrow(cwest_lon_degree_allocs)

	var csouth_lat_degree_allocs *cgoAllocMap
	ref48245dab.south_lat_degree, csouth_lat_degree_allocs = (C.double)(x.south_lat_degree), cgoAllocsUnknown
	allocs48245dab.Borrow(csouth_lat_degree_allocs)

	var ceast_lon_degree_allocs *cgoAllocMap
	ref48245dab.east_lon_degree, ceast_lon_degree_allocs = (C.double)(x.east_lon_degree), cgoAllocsUnknown
	allocs48245dab.Borrow(ceast_lon_degree_allocs)

	var cnorth_lat_degree_allocs *cgoAllocMap
	ref48245dab.north_lat_degree, cnorth_lat_degree_allocs = (C.double)(x.north_lat_degree), cgoAllocsUnknown
	allocs48245dab.Borrow(cnorth_lat_degree_allocs)

	var carea_name_allocs *cgoAllocMap
	ref48245dab.area_name, carea_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.area_name)).Data)), cgoAllocsUnknown
	allocs48245dab.Borrow(carea_name_allocs)

	var cprojection_method_name_allocs *cgoAllocMap
	ref48245dab.projection_method_name, cprojection_method_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.projection_method_name)).Data)), cgoAllocsUnknown
	allocs48245dab.Borrow(cprojection_method_name_allocs)

	x.ref48245dab = ref48245dab
	x.allocs48245dab = allocs48245dab
	return ref48245dab, allocs48245dab

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x PROJ_CRS_INFO) PassValue() (C.PROJ_CRS_INFO, *cgoAllocMap) {
	if x.ref48245dab != nil {
		return *x.ref48245dab, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PROJ_CRS_INFO) Deref() {
	if x.ref48245dab == nil {
		return
	}
	hxfc4425b := (*sliceHeader)(unsafe.Pointer(&x.auth_name))
	hxfc4425b.Data = unsafe.Pointer(x.ref48245dab.auth_name)
	hxfc4425b.Cap = 0x7fffffff
	// hxfc4425b.Len = ?

	hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&x.code))
	hxf95e7c8.Data = unsafe.Pointer(x.ref48245dab.code)
	hxf95e7c8.Cap = 0x7fffffff
	// hxf95e7c8.Len = ?

	hxff2234b := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxff2234b.Data = unsafe.Pointer(x.ref48245dab.name)
	hxff2234b.Cap = 0x7fffffff
	// hxff2234b.Len = ?

	x.deprecated = (int32)(x.ref48245dab.deprecated)
	x.bbox_valid = (int32)(x.ref48245dab.bbox_valid)
	x.west_lon_degree = (float64)(x.ref48245dab.west_lon_degree)
	x.south_lat_degree = (float64)(x.ref48245dab.south_lat_degree)
	x.east_lon_degree = (float64)(x.ref48245dab.east_lon_degree)
	x.north_lat_degree = (float64)(x.ref48245dab.north_lat_degree)
	hxff73280 := (*sliceHeader)(unsafe.Pointer(&x.area_name))
	hxff73280.Data = unsafe.Pointer(x.ref48245dab.area_name)
	hxff73280.Cap = 0x7fffffff
	// hxff73280.Len = ?

	hxfa9955c := (*sliceHeader)(unsafe.Pointer(&x.projection_method_name))
	hxfa9955c.Data = unsafe.Pointer(x.ref48245dab.projection_method_name)
	hxfa9955c.Cap = 0x7fffffff
	// hxfa9955c.Len = ?

}

// allocPROJ_CRS_LIST_PARAMETERSMemory allocates memory for type C.PROJ_CRS_LIST_PARAMETERS in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPROJ_CRS_LIST_PARAMETERSMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPROJ_CRS_LIST_PARAMETERSValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPROJ_CRS_LIST_PARAMETERSValue = unsafe.Sizeof([1]C.PROJ_CRS_LIST_PARAMETERS{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PROJ_CRS_LIST_PARAMETERS) Ref() *C.PROJ_CRS_LIST_PARAMETERS {
	if x == nil {
		return nil
	}
	return x.ref577cd7ca
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PROJ_CRS_LIST_PARAMETERS) Free() {
	if x != nil && x.allocs577cd7ca != nil {
		x.allocs577cd7ca.(*cgoAllocMap).Free()
		x.ref577cd7ca = nil
	}
}

// NewPROJ_CRS_LIST_PARAMETERSRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPROJ_CRS_LIST_PARAMETERSRef(ref unsafe.Pointer) *PROJ_CRS_LIST_PARAMETERS {
	if ref == nil {
		return nil
	}
	obj := new(PROJ_CRS_LIST_PARAMETERS)
	obj.ref577cd7ca = (*C.PROJ_CRS_LIST_PARAMETERS)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PROJ_CRS_LIST_PARAMETERS) PassRef() (*C.PROJ_CRS_LIST_PARAMETERS, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref577cd7ca != nil {
		return x.ref577cd7ca, nil
	}
	mem577cd7ca := allocPROJ_CRS_LIST_PARAMETERSMemory(1)
	ref577cd7ca := (*C.PROJ_CRS_LIST_PARAMETERS)(mem577cd7ca)
	allocs577cd7ca := new(cgoAllocMap)
	allocs577cd7ca.Add(mem577cd7ca)

	var ctypesCount_allocs *cgoAllocMap
	ref577cd7ca.typesCount, ctypesCount_allocs = (C.size_t)(x.typesCount), cgoAllocsUnknown
	allocs577cd7ca.Borrow(ctypesCount_allocs)

	var ccrs_area_of_use_contains_bbox_allocs *cgoAllocMap
	ref577cd7ca.crs_area_of_use_contains_bbox, ccrs_area_of_use_contains_bbox_allocs = (C.int)(x.crs_area_of_use_contains_bbox), cgoAllocsUnknown
	allocs577cd7ca.Borrow(ccrs_area_of_use_contains_bbox_allocs)

	var cbbox_valid_allocs *cgoAllocMap
	ref577cd7ca.bbox_valid, cbbox_valid_allocs = (C.int)(x.bbox_valid), cgoAllocsUnknown
	allocs577cd7ca.Borrow(cbbox_valid_allocs)

	var cwest_lon_degree_allocs *cgoAllocMap
	ref577cd7ca.west_lon_degree, cwest_lon_degree_allocs = (C.double)(x.west_lon_degree), cgoAllocsUnknown
	allocs577cd7ca.Borrow(cwest_lon_degree_allocs)

	var csouth_lat_degree_allocs *cgoAllocMap
	ref577cd7ca.south_lat_degree, csouth_lat_degree_allocs = (C.double)(x.south_lat_degree), cgoAllocsUnknown
	allocs577cd7ca.Borrow(csouth_lat_degree_allocs)

	var ceast_lon_degree_allocs *cgoAllocMap
	ref577cd7ca.east_lon_degree, ceast_lon_degree_allocs = (C.double)(x.east_lon_degree), cgoAllocsUnknown
	allocs577cd7ca.Borrow(ceast_lon_degree_allocs)

	var cnorth_lat_degree_allocs *cgoAllocMap
	ref577cd7ca.north_lat_degree, cnorth_lat_degree_allocs = (C.double)(x.north_lat_degree), cgoAllocsUnknown
	allocs577cd7ca.Borrow(cnorth_lat_degree_allocs)

	var callow_deprecated_allocs *cgoAllocMap
	ref577cd7ca.allow_deprecated, callow_deprecated_allocs = (C.int)(x.allow_deprecated), cgoAllocsUnknown
	allocs577cd7ca.Borrow(callow_deprecated_allocs)

	x.ref577cd7ca = ref577cd7ca
	x.allocs577cd7ca = allocs577cd7ca
	return ref577cd7ca, allocs577cd7ca

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x PROJ_CRS_LIST_PARAMETERS) PassValue() (C.PROJ_CRS_LIST_PARAMETERS, *cgoAllocMap) {
	if x.ref577cd7ca != nil {
		return *x.ref577cd7ca, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PROJ_CRS_LIST_PARAMETERS) Deref() {
	if x.ref577cd7ca == nil {
		return
	}
	x.typesCount = (uint)(x.ref577cd7ca.typesCount)
	x.crs_area_of_use_contains_bbox = (int32)(x.ref577cd7ca.crs_area_of_use_contains_bbox)
	x.bbox_valid = (int32)(x.ref577cd7ca.bbox_valid)
	x.west_lon_degree = (float64)(x.ref577cd7ca.west_lon_degree)
	x.south_lat_degree = (float64)(x.ref577cd7ca.south_lat_degree)
	x.east_lon_degree = (float64)(x.ref577cd7ca.east_lon_degree)
	x.north_lat_degree = (float64)(x.ref577cd7ca.north_lat_degree)
	x.allow_deprecated = (int32)(x.ref577cd7ca.allow_deprecated)
}

// safeString ensures that the string is NULL-terminated, a NULL-terminated copy is created otherwise.
func safeString(str string) string {
	if len(str) > 0 && str[len(str)-1] != '\x00' {
		str = str + "\x00"
	} else if len(str) == 0 {
		str = "\x00"
	}
	return str
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	str = safeString(str)
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(h.Data), cgoAllocsUnknown
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}
