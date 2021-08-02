#!/bin/sh

# SPDX-License-Identifier: Unlicense OR MIT

set -e

OBJCOPY_ARM64=$ANDROID_SDK_ROOT/ndk/21.3.6528147/toolchains/aarch64-linux-android-4.9/prebuilt/linux-x86_64/aarch64-linux-android/bin/objcopy
OBJCOPY_ARM=$ANDROID_SDK_ROOT/ndk/21.3.6528147/toolchains/arm-linux-androideabi-4.9/prebuilt/linux-x86_64/arm-linux-androideabi/bin/objcopy

export CGO_ENABLED=1
export GOARCH=386
export VK_ICD_FILENAMES=../swiftshader/build.32bit/Linux/vk_swiftshader_icd.json

export SWIFTSHADER_TRIPLE=armv7a-none-eabi
go run ../cmd/compile -arch arm -objcopy $OBJCOPY_ARM -layout "0:buffer,1:buffer,2:image,3:image" $GIO/gpu/shaders/kernel4.comp
go run ../cmd/compile -arch arm -objcopy $OBJCOPY_ARM -layout "0:buffer,1:buffer" $GIO/gpu/shaders/coarse.comp
go run ../cmd/compile -arch arm -objcopy $OBJCOPY_ARM -layout "0:buffer,1:buffer" $GIO/gpu/shaders/binning.comp
go run ../cmd/compile -arch arm -objcopy $OBJCOPY_ARM -layout "0:buffer,1:buffer" $GIO/gpu/shaders/backdrop.comp
go run ../cmd/compile -arch arm -objcopy $OBJCOPY_ARM -layout "0:buffer,1:buffer" $GIO/gpu/shaders/path_coarse.comp
go run ../cmd/compile -arch arm -objcopy $OBJCOPY_ARM -layout "0:buffer,1:buffer" $GIO/gpu/shaders/tile_alloc.comp
go run ../cmd/compile -arch arm -objcopy $OBJCOPY_ARM -layout "0:buffer,1:buffer,2:buffer,3:buffer" $GIO/gpu/shaders/elements.comp

export GOARCH=amd64
export VK_ICD_FILENAMES=../swiftshader/build.64bit/Linux/vk_swiftshader_icd.json
export SWIFTSHADER_TRIPLE=x86_64-unknown-none-gnu

go run ../cmd/compile -arch amd64 -layout "0:buffer,1:buffer,2:image,3:image" $GIO/gpu/shaders/kernel4.comp
go run ../cmd/compile -arch amd64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/coarse.comp
go run ../cmd/compile -arch amd64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/binning.comp
go run ../cmd/compile -arch amd64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/backdrop.comp
go run ../cmd/compile -arch amd64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/path_coarse.comp
go run ../cmd/compile -arch amd64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/tile_alloc.comp
go run ../cmd/compile -arch amd64 -layout "0:buffer,1:buffer,2:buffer,3:buffer" $GIO/gpu/shaders/elements.comp

export SWIFTSHADER_TRIPLE=aarch64-unknown-linux-gnu

go run ../cmd/compile -arch arm64 -objcopy $OBJCOPY_ARM64 -layout "0:buffer,1:buffer,2:image,3:image" $GIO/gpu/shaders/kernel4.comp
go run ../cmd/compile -arch arm64 -objcopy $OBJCOPY_ARM64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/coarse.comp
go run ../cmd/compile -arch arm64 -objcopy $OBJCOPY_ARM64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/binning.comp
go run ../cmd/compile -arch arm64 -objcopy $OBJCOPY_ARM64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/backdrop.comp
go run ../cmd/compile -arch arm64 -objcopy $OBJCOPY_ARM64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/path_coarse.comp
go run ../cmd/compile -arch arm64 -objcopy $OBJCOPY_ARM64 -layout "0:buffer,1:buffer" $GIO/gpu/shaders/tile_alloc.comp
go run ../cmd/compile -arch arm64 -objcopy $OBJCOPY_ARM64 -layout "0:buffer,1:buffer,2:buffer,3:buffer" $GIO/gpu/shaders/elements.comp
