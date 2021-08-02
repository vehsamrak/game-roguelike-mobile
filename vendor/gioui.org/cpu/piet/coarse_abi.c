// Code generated by gioui.org/cpu/cmd/compile DO NOT EDIT.

//go:build linux && (arm64 || arm || amd64)
// +build linux
// +build arm64 arm amd64

#include <stdint.h>
#include <stddef.h>
#include "../abi.h"
#include "../runtime.h"
#include "coarse_abi.h"

const struct program_info coarse_program_info = {
	.has_cbarriers = 1,
	.min_memory_size = 100000,
	.desc_set_size = sizeof(struct coarse_descriptor_set_layout),
	.workgroup_size_x = 128,
	.workgroup_size_y = 1,
	.workgroup_size_z = 1,
	.begin = coarse_coroutine_begin,
	.await = coarse_coroutine_await,
	.destroy = coarse_coroutine_destroy,
};