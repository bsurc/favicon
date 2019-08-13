// Copyright (c) 2019, Boise State University All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package favicon

import (
	"net/http"
)

// hexdump favicon.ico, then swap little and big
var icon = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
	0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x40, 0x08, 0x06, 0x00, 0x00, 0x00, 0xaa, 0x69, 0x71,
	0xde, 0x00, 0x00, 0x00, 0x19, 0x74, 0x45, 0x58, 0x74, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72,
	0x65, 0x00, 0x41, 0x64, 0x6f, 0x62, 0x65, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x71, 0xc9, 0x65, 0x3c, 0x00, 0x00, 0x03, 0x28, 0x69, 0x54, 0x58, 0x74, 0x58, 0x4d,
	0x4c, 0x3a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6f, 0x62, 0x65, 0x2e, 0x78, 0x6d, 0x70, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x3c, 0x3f, 0x78, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x3d, 0x22, 0xef, 0xbb, 0xbf, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x57, 0x35,
	0x4d, 0x30, 0x4d, 0x70, 0x43, 0x65, 0x68, 0x69, 0x48, 0x7a, 0x72, 0x65, 0x53, 0x7a, 0x4e, 0x54,
	0x63, 0x7a, 0x6b, 0x63, 0x39, 0x64, 0x22, 0x3f, 0x3e, 0x20, 0x3c, 0x78, 0x3a, 0x78, 0x6d, 0x70,
	0x6d, 0x65, 0x74, 0x61, 0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x78, 0x3d, 0x22, 0x61, 0x64,
	0x6f, 0x62, 0x65, 0x3a, 0x6e, 0x73, 0x3a, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x22, 0x20, 0x78, 0x3a,
	0x78, 0x6d, 0x70, 0x74, 0x6b, 0x3d, 0x22, 0x41, 0x64, 0x6f, 0x62, 0x65, 0x20, 0x58, 0x4d, 0x50,
	0x20, 0x43, 0x6f, 0x72, 0x65, 0x20, 0x35, 0x2e, 0x36, 0x2d, 0x63, 0x31, 0x34, 0x30, 0x20, 0x37,
	0x39, 0x2e, 0x31, 0x36, 0x30, 0x34, 0x35, 0x31, 0x2c, 0x20, 0x32, 0x30, 0x31, 0x37, 0x2f, 0x30,
	0x35, 0x2f, 0x30, 0x36, 0x2d, 0x30, 0x31, 0x3a, 0x30, 0x38, 0x3a, 0x32, 0x31, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x3e, 0x20, 0x3c, 0x72, 0x64, 0x66, 0x3a, 0x52, 0x44, 0x46,
	0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x72, 0x64, 0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x77, 0x33, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x31, 0x39,
	0x39, 0x39, 0x2f, 0x30, 0x32, 0x2f, 0x32, 0x32, 0x2d, 0x72, 0x64, 0x66, 0x2d, 0x73, 0x79, 0x6e,
	0x74, 0x61, 0x78, 0x2d, 0x6e, 0x73, 0x23, 0x22, 0x3e, 0x20, 0x3c, 0x72, 0x64, 0x66, 0x3a, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x64, 0x66, 0x3a, 0x61,
	0x62, 0x6f, 0x75, 0x74, 0x3d, 0x22, 0x22, 0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x78, 0x6d,
	0x70, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61, 0x64, 0x6f,
	0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x61, 0x70, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x22,
	0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3d, 0x22, 0x68, 0x74,
	0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61, 0x64, 0x6f, 0x62, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x78, 0x61, 0x70, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6d, 0x6d, 0x2f, 0x22, 0x20, 0x78,
	0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x73, 0x74, 0x52, 0x65, 0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70,
	0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61, 0x64, 0x6f, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x78, 0x61, 0x70, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x2f, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x23, 0x22, 0x20, 0x78, 0x6d, 0x70, 0x3a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x6f, 0x6c, 0x3d, 0x22, 0x41, 0x64, 0x6f,
	0x62, 0x65, 0x20, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x20, 0x43, 0x43, 0x20,
	0x32, 0x30, 0x31, 0x38, 0x20, 0x28, 0x4d, 0x61, 0x63, 0x69, 0x6e, 0x74, 0x6f, 0x73, 0x68, 0x29,
	0x22, 0x20, 0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x44, 0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e, 0x69, 0x69, 0x64, 0x3a, 0x37, 0x37, 0x36, 0x38,
	0x36, 0x32, 0x37, 0x39, 0x37, 0x37, 0x44, 0x31, 0x31, 0x31, 0x45, 0x38, 0x42, 0x43, 0x45, 0x38,
	0x38, 0x35, 0x39, 0x43, 0x46, 0x45, 0x38, 0x41, 0x38, 0x32, 0x36, 0x34, 0x22, 0x20, 0x78, 0x6d,
	0x70, 0x4d, 0x4d, 0x3a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x3d, 0x22,
	0x78, 0x6d, 0x70, 0x2e, 0x64, 0x69, 0x64, 0x3a, 0x37, 0x37, 0x36, 0x38, 0x36, 0x32, 0x37, 0x41,
	0x37, 0x37, 0x44, 0x31, 0x31, 0x31, 0x45, 0x38, 0x42, 0x43, 0x45, 0x38, 0x38, 0x35, 0x39, 0x43,
	0x46, 0x45, 0x38, 0x41, 0x38, 0x32, 0x36, 0x34, 0x22, 0x3e, 0x20, 0x3c, 0x78, 0x6d, 0x70, 0x4d,
	0x4d, 0x3a, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x20, 0x73, 0x74,
	0x52, 0x65, 0x66, 0x3a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x3d, 0x22,
	0x78, 0x6d, 0x70, 0x2e, 0x69, 0x69, 0x64, 0x3a, 0x37, 0x37, 0x36, 0x38, 0x36, 0x32, 0x37, 0x37,
	0x37, 0x37, 0x44, 0x31, 0x31, 0x31, 0x45, 0x38, 0x42, 0x43, 0x45, 0x38, 0x38, 0x35, 0x39, 0x43,
	0x46, 0x45, 0x38, 0x41, 0x38, 0x32, 0x36, 0x34, 0x22, 0x20, 0x73, 0x74, 0x52, 0x65, 0x66, 0x3a,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e,
	0x64, 0x69, 0x64, 0x3a, 0x37, 0x37, 0x36, 0x38, 0x36, 0x32, 0x37, 0x38, 0x37, 0x37, 0x44, 0x31,
	0x31, 0x31, 0x45, 0x38, 0x42, 0x43, 0x45, 0x38, 0x38, 0x35, 0x39, 0x43, 0x46, 0x45, 0x38, 0x41,
	0x38, 0x32, 0x36, 0x34, 0x22, 0x2f, 0x3e, 0x20, 0x3c, 0x2f, 0x72, 0x64, 0x66, 0x3a, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3e, 0x20, 0x3c, 0x2f, 0x72, 0x64, 0x66,
	0x3a, 0x52, 0x44, 0x46, 0x3e, 0x20, 0x3c, 0x2f, 0x78, 0x3a, 0x78, 0x6d, 0x70, 0x6d, 0x65, 0x74,
	0x61, 0x3e, 0x20, 0x3c, 0x3f, 0x78, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x65, 0x6e, 0x64,
	0x3d, 0x22, 0x72, 0x22, 0x3f, 0x3e, 0xe1, 0x3b, 0x64, 0x4d, 0x00, 0x00, 0x0c, 0xe8, 0x49, 0x44,
	0x41, 0x54, 0x78, 0xda, 0xec, 0x5b, 0x09, 0x90, 0x14, 0xd5, 0x19, 0xfe, 0x5e, 0x77, 0xcf, 0xb9,
	0x07, 0x7b, 0x0e, 0x7b, 0x01, 0x22, 0xcb, 0xa1, 0xa8, 0x80, 0x22, 0x18, 0x83, 0xa2, 0x1c, 0x51,
	0x14, 0x10, 0x04, 0x0f, 0x34, 0x1e, 0x10, 0x8b, 0x08, 0x18, 0x8d, 0xa6, 0x42, 0x11, 0x2b, 0xa6,
	0x28, 0x31, 0x56, 0xa5, 0x4c, 0x55, 0x34, 0x85, 0x37, 0x28, 0x82, 0x1c, 0xde, 0x98, 0x28, 0x41,
	0x51, 0x90, 0x12, 0x04, 0x85, 0x02, 0x89, 0xb8, 0x48, 0xb0, 0x04, 0x76, 0x17, 0x58, 0x66, 0x77,
	0x76, 0x67, 0xe7, 0x3e, 0xfb, 0xe5, 0x7f, 0xdd, 0xb3, 0xc0, 0xe2, 0xce, 0x4c, 0xf7, 0x86, 0x24,
	0x4b, 0xb1, 0x5d, 0xf5, 0x76, 0xae, 0xee, 0x7e, 0xef, 0xbf, 0xbe, 0xff, 0xfb, 0xff, 0xd7, 0xcb,
	0x38, 0xe7, 0x38, 0x97, 0x0f, 0x09, 0xe7, 0xf8, 0xd1, 0xad, 0x80, 0x6e, 0x05, 0x74, 0x2b, 0xa0,
	0x5b, 0x01, 0xe7, 0xf4, 0xc1, 0xf8, 0x2c, 0xd7, 0xe9, 0xdf, 0xfd, 0x9e, 0xc6, 0x5c, 0x70, 0xe5,
	0x30, 0xe4, 0x30, 0x60, 0x6b, 0xa1, 0xb3, 0xe8, 0x9b, 0xcc, 0xd9, 0x92, 0x75, 0x39, 0xc9, 0x38,
	0xe2, 0xf4, 0xb7, 0x05, 0x9c, 0xd5, 0x42, 0xb5, 0x7d, 0x85, 0x78, 0xfe, 0x0e, 0x92, 0x69, 0x1f,
	0x58, 0x92, 0x7e, 0x63, 0x6d, 0xd2, 0x77, 0xa8, 0x80, 0x9f, 0xd2, 0x2f, 0xeb, 0xa0, 0x04, 0xf2,
	0xb7, 0xb5, 0x5e, 0x84, 0xcf, 0x7d, 0x83, 0x51, 0x95, 0x73, 0x08, 0x56, 0xba, 0x23, 0xe7, 0x52,
	0xd7, 0x92, 0x9f, 0xa7, 0xfb, 0x9a, 0xc1, 0xc1, 0x12, 0x28, 0xb4, 0xb6, 0xa0, 0xd4, 0xe6, 0xc1,
	0x40, 0x9b, 0x9b, 0xc3, 0x12, 0xad, 0x45, 0xc2, 0xba, 0x02, 0xd1, 0xa2, 0x97, 0x68, 0xbd, 0x75,
	0xda, 0xc5, 0x69, 0x14, 0x20, 0x22, 0xc3, 0x06, 0xc5, 0xbf, 0xe2, 0x40, 0x22, 0xef, 0x96, 0x29,
	0xdf, 0x3e, 0x8e, 0x9a, 0xa3, 0xd7, 0x03, 0xf9, 0xfb, 0xe9, 0x02, 0x52, 0x2a, 0x57, 0xb2, 0xba,
	0x43, 0xd7, 0xf0, 0x6d, 0xf1, 0x47, 0x85, 0x53, 0x4a, 0xe0, 0x02, 0xfb, 0x31, 0x8c, 0x2a, 0xd8,
	0x8d, 0x39, 0x15, 0x1f, 0x60, 0x60, 0x5e, 0x5d, 0x13, 0x82, 0x25, 0x8f, 0x90, 0x57, 0xac, 0x00,
	0xfd, 0x96, 0x46, 0x01, 0xd0, 0x05, 0xb5, 0x7a, 0x96, 0x91, 0xe6, 0xee, 0x99, 0xf1, 0xcd, 0x42,
	0xac, 0x39, 0x32, 0x05, 0xc8, 0x23, 0x25, 0x48, 0xd1, 0x94, 0x12, 0xce, 0x02, 0x0d, 0x70, 0x99,
	0x74, 0x40, 0x6b, 0x55, 0xed, 0x40, 0xac, 0x07, 0x72, 0x2d, 0x7e, 0xbe, 0xb0, 0xfa, 0x79, 0xf6,
	0x9b, 0x5e, 0xef, 0x03, 0xa1, 0xe2, 0xdb, 0xa0, 0x5a, 0xde, 0x94, 0x17, 0x0e, 0xcb, 0x49, 0x87,
	0x0e, 0x40, 0x22, 0x77, 0x2d, 0xdd, 0x45, 0x9a, 0xde, 0xfb, 0xe3, 0xd1, 0x61, 0x44, 0xb1, 0xb5,
	0xe1, 0x3a, 0x0e, 0x25, 0xc8, 0x68, 0xe8, 0x13, 0x30, 0xb5, 0x8b, 0x8f, 0x04, 0x19, 0x2c, 0x06,
	0xc8, 0x21, 0xc2, 0xb2, 0x26, 0xc4, 0xb8, 0x95, 0x7d, 0x5c, 0x3f, 0x8d, 0x97, 0xda, 0x8f, 0xb0,
	0x11, 0xc5, 0x5f, 0x4f, 0x41, 0x3c, 0x77, 0x59, 0x7a, 0x05, 0x68, 0x02, 0x92, 0x12, 0x92, 0xf6,
	0xcf, 0x90, 0x90, 0xdd, 0xe3, 0xcb, 0xb7, 0xdf, 0x58, 0x20, 0x79, 0xd9, 0x47, 0xc7, 0x26, 0xd0,
	0x4d, 0x29, 0x14, 0x64, 0xa1, 0x04, 0x39, 0xe5, 0x6b, 0x67, 0xc3, 0xa0, 0xb5, 0x2a, 0x7e, 0x7a,
	0x89, 0xb2, 0xf5, 0xcd, 0x97, 0xf3, 0x7b, 0x8b, 0x3f, 0x97, 0x0b, 0xec, 0x9e, 0xc6, 0xec, 0x69,
	0x50, 0x68, 0x31, 0xe1, 0x7c, 0x1e, 0xbe, 0x92, 0x9b, 0x7e, 0x5d, 0xfd, 0x66, 0xf2, 0xed, 0x61,
	0x8f, 0x00, 0xd1, 0x12, 0x8e, 0x48, 0x99, 0xae, 0xdd, 0xb3, 0xe6, 0x20, 0x63, 0xaa, 0x56, 0xc0,
	0xd2, 0x0c, 0x35, 0x56, 0xc8, 0x96, 0x37, 0x5d, 0x45, 0x0a, 0x89, 0x4e, 0x31, 0xc6, 0x03, 0x84,
	0x12, 0x54, 0xfb, 0xdf, 0xd0, 0x5a, 0x36, 0x79, 0x5a, 0xc5, 0x16, 0xdf, 0xfb, 0x43, 0xe7, 0xb3,
	0x1c, 0xd5, 0xca, 0x11, 0xea, 0x75, 0x96, 0x29, 0x01, 0x3a, 0x2e, 0x50, 0x7a, 0xff, 0x46, 0xac,
	0x5d, 0x45, 0x2f, 0xc9, 0x30, 0xa0, 0x68, 0xf9, 0x53, 0x5e, 0x07, 0x6f, 0xf9, 0x88, 0xc9, 0xae,
	0x6d, 0xb5, 0xdb, 0x46, 0xce, 0x62, 0x95, 0x72, 0x88, 0xc3, 0xdf, 0x5f, 0x57, 0x10, 0xce, 0xa2,
	0xb2, 0x5a, 0x4a, 0xc2, 0x93, 0xc8, 0x25, 0x05, 0x30, 0x9b, 0x39, 0x26, 0x28, 0x80, 0x05, 0x6c,
	0x3f, 0x7c, 0xe5, 0xd7, 0x5f, 0x9c, 0x5b, 0xbb, 0x77, 0xdd, 0xf0, 0x87, 0xd8, 0xf0, 0x1e, 0x7b,
	0x89, 0x72, 0x14, 0x1a, 0xd7, 0xfe, 0x19, 0x19, 0x92, 0xfe, 0x6a, 0x80, 0xa1, 0x75, 0xbc, 0x0e,
	0x09, 0x45, 0x02, 0x18, 0x25, 0x1e, 0x31, 0x9f, 0xcf, 0x98, 0x46, 0x20, 0xf6, 0x21, 0x50, 0x35,
	0xec, 0x92, 0x82, 0xc3, 0x7b, 0x16, 0xf7, 0x5f, 0x7c, 0xe1, 0x15, 0xbb, 0x16, 0x73, 0x72, 0x2b,
	0xa6, 0xf1, 0x84, 0x0e, 0x49, 0x11, 0xfd, 0x84, 0x24, 0x72, 0x59, 0x82, 0x07, 0x63, 0x25, 0x50,
	0xc5, 0xe2, 0x19, 0x4f, 0x47, 0xa2, 0x98, 0x31, 0x62, 0x25, 0x00, 0xda, 0x42, 0xa9, 0x9a, 0x98,
	0xaa, 0x25, 0xa8, 0xa7, 0x3b, 0xa3, 0xde, 0xac, 0x5a, 0x70, 0xbe, 0xe3, 0xa8, 0x78, 0x7b, 0x5c,
	0xc9, 0x4c, 0xb3, 0xf8, 0x49, 0xa1, 0x49, 0x80, 0xf6, 0xf7, 0x89, 0xe5, 0x20, 0x86, 0xfc, 0x17,
	0x1b, 0xae, 0x83, 0xc6, 0x2d, 0x35, 0x7e, 0x20, 0xa7, 0xb9, 0x15, 0xb1, 0x48, 0x12, 0xfc, 0xbe,
	0xaa, 0xf7, 0xf8, 0x03, 0x95, 0x1f, 0x63, 0x7f, 0xa0, 0x37, 0x4b, 0x72, 0x95, 0xa5, 0x23, 0x75,
	0xfc, 0x34, 0x05, 0xf0, 0x0e, 0x14, 0x62, 0x93, 0x63, 0x38, 0x18, 0x75, 0xf1, 0xc7, 0xeb, 0xa6,
	0xf3, 0x86, 0x78, 0x9e, 0xa4, 0xa5, 0x3a, 0xc3, 0x8c, 0x54, 0xc5, 0x25, 0xce, 0x23, 0xe2, 0xcd,
	0x5e, 0xc5, 0xf8, 0x45, 0xbc, 0x3d, 0x49, 0xb2, 0x35, 0xce, 0x3d, 0x1c, 0xe8, 0x53, 0xb5, 0xc2,
	0x3d, 0x9a, 0x13, 0x61, 0x62, 0xe0, 0x52, 0x66, 0xaf, 0x51, 0x02, 0x78, 0xeb, 0xd8, 0x44, 0xe9,
	0xda, 0xa2, 0xbd, 0x98, 0xd4, 0x6b, 0x0b, 0x10, 0xc4, 0x17, 0x34, 0xf5, 0x7a, 0xfa, 0x35, 0x27,
	0x0b, 0xb7, 0x6e, 0xf3, 0xf3, 0x53, 0xbf, 0x23, 0x93, 0xb3, 0x22, 0x38, 0xf9, 0xbc, 0x9d, 0xad,
	0x83, 0xa4, 0x57, 0x8e, 0x4e, 0xa2, 0xbb, 0xfc, 0x90, 0x9d, 0xa0, 0x89, 0x10, 0x8e, 0xe7, 0xc1,
	0x69, 0x6f, 0xe0, 0xe3, 0xf3, 0xf6, 0x31, 0x24, 0xec, 0x9f, 0x2b, 0x9d, 0x03, 0x2f, 0xb2, 0xb8,
	0x12, 0x9f, 0xb9, 0xcc, 0x73, 0x25, 0x12, 0xb1, 0x62, 0xa6, 0x4d, 0x9e, 0xb1, 0xb0, 0xa4, 0xb5,
	0x2b, 0x21, 0xd4, 0xc7, 0x9c, 0x98, 0xbc, 0xe3, 0x19, 0xbc, 0x33, 0x64, 0x01, 0x6e, 0x2e, 0xdd,
	0x92, 0x44, 0xb8, 0x64, 0x51, 0xe7, 0x90, 0x9c, 0xe6, 0xb2, 0xf8, 0xc6, 0xf9, 0x9a, 0xed, 0x0f,
	0x6e, 0xf0, 0x0d, 0xd4, 0xc3, 0xc0, 0x48, 0x61, 0x2b, 0x80, 0x9c, 0x18, 0xe1, 0x04, 0xd7, 0x6e,
	0xe6, 0xca, 0x39, 0x12, 0x43, 0xd8, 0xf5, 0xa9, 0x72, 0xa2, 0x32, 0x32, 0x43, 0x31, 0x15, 0xff,
	0xa4, 0x78, 0xd2, 0xd2, 0x7f, 0xf5, 0xf1, 0xb1, 0x9a, 0x65, 0x0d, 0x79, 0x91, 0xb0, 0x8e, 0x4c,
	0xe7, 0xca, 0xf9, 0x28, 0x10, 0xd7, 0xb0, 0xe4, 0x21, 0x1d, 0x54, 0x39, 0xcc, 0x7b, 0x21, 0x9d,
	0x6f, 0xf3, 0xff, 0x71, 0x59, 0xfd, 0xb5, 0xa8, 0xf3, 0x0f, 0x4a, 0x59, 0x5f, 0x32, 0x80, 0x19,
	0x36, 0x4d, 0x51, 0xf7, 0xb8, 0x36, 0x89, 0x2f, 0xe8, 0x8f, 0xf4, 0x83, 0x62, 0xba, 0x90, 0x13,
	0x0a, 0xb3, 0x06, 0xee, 0xdd, 0xd2, 0x7c, 0x29, 0xf6, 0x07, 0xfb, 0xd1, 0x42, 0x1a, 0x0c, 0xc7,
	0x1d, 0xa2, 0xa5, 0xf8, 0x49, 0xd1, 0x4e, 0x8c, 0x29, 0xd9, 0x01, 0x44, 0x0a, 0x9e, 0x3e, 0xc9,
	0xd2, 0x92, 0xa7, 0x80, 0x62, 0x86, 0x79, 0xdb, 0xce, 0x97, 0x43, 0x97, 0xd3, 0xed, 0x46, 0xac,
	0x76, 0x5f, 0x43, 0xef, 0x23, 0xc6, 0xd7, 0x1e, 0xef, 0x81, 0x0a, 0x7b, 0x03, 0x46, 0x17, 0xed,
	0xa2, 0xf7, 0x8e, 0x35, 0x42, 0x29, 0xe6, 0x43, 0x80, 0x25, 0x7b, 0xd1, 0x1a, 0x6e, 0x58, 0x2a,
	0xc0, 0x4f, 0x2c, 0x4a, 0x09, 0x6b, 0xa8, 0x6a, 0xa8, 0x3a, 0x23, 0x26, 0x76, 0x57, 0xcf, 0x8d,
	0xe2, 0xd3, 0x17, 0x48, 0xe4, 0xec, 0xd2, 0x28, 0xb5, 0xc8, 0x1c, 0x9a, 0x27, 0x64, 0xb3, 0xbc,
	0x48, 0x7f, 0x36, 0xfd, 0xa3, 0xa5, 0xf5, 0x97, 0xbb, 0xfc, 0x03, 0xb1, 0xdd, 0x7b, 0x31, 0x19,
	0xc0, 0x6b, 0xc0, 0x83, 0x52, 0x3c, 0x86, 0xd2, 0xf5, 0x6d, 0x55, 0x6b, 0x91, 0x6f, 0x0d, 0x1d,
	0x47, 0xa0, 0xec, 0x0d, 0xa1, 0x74, 0x73, 0x3c, 0x40, 0x95, 0xc5, 0xe4, 0xd3, 0xdc, 0x81, 0x0a,
	0xdb, 0xdb, 0xde, 0x61, 0x1c, 0x54, 0x6b, 0xa7, 0x45, 0xfe, 0xd3, 0xad, 0x1f, 0xcf, 0x47, 0xae,
	0xa3, 0x9e, 0xdf, 0x56, 0x40, 0xd6, 0x8f, 0x3b, 0x9f, 0x13, 0xa5, 0xa8, 0x2e, 0x7c, 0xe2, 0x14,
	0x0d, 0xa5, 0x1b, 0xd2, 0xc9, 0xe2, 0x86, 0x33, 0x2a, 0xd5, 0xf9, 0xc4, 0x37, 0x84, 0xf5, 0xb9,
	0xd5, 0x58, 0x08, 0x0a, 0xe1, 0x93, 0x39, 0x60, 0x4a, 0x0b, 0x9f, 0xe7, 0xda, 0x40, 0xf3, 0x5b,
	0x97, 0xd0, 0x3d, 0xc3, 0x42, 0xb1, 0x92, 0xe9, 0xa2, 0xc2, 0x12, 0x9d, 0xbd, 0xc4, 0x33, 0x0a,
	0xd1, 0x50, 0x39, 0xd3, 0x62, 0x9a, 0x1b, 0x05, 0x9f, 0x22, 0xdc, 0x59, 0xfc, 0x15, 0x2b, 0xca,
	0x75, 0xd7, 0x93, 0xfb, 0x7d, 0x44, 0xf3, 0x5b, 0xe9, 0x7b, 0x92, 0x40, 0xa2, 0x81, 0x8e, 0x86,
	0xa5, 0xfd, 0x60, 0xe2, 0xd5, 0x0e, 0x47, 0xe3, 0x1f, 0x10, 0xb7, 0xf5, 0x7c, 0xa5, 0xf1, 0x6a,
	0x4e, 0xc6, 0x30, 0x86, 0x1f, 0x42, 0x71, 0xa1, 0x2a, 0xfc, 0xaa, 0x7c, 0x1d, 0xeb, 0x97, 0x7f,
	0x98, 0x0a, 0xdb, 0xc2, 0xe7, 0xda, 0x14, 0xaf, 0x98, 0x03, 0xbf, 0xc0, 0x18, 0x55, 0x65, 0x17,
	0x68, 0xda, 0x17, 0xae, 0x6f, 0x94, 0xb0, 0x24, 0xf2, 0xe1, 0x90, 0x83, 0xfc, 0xb1, 0x3e, 0xab,
	0x85, 0x28, 0x39, 0xb0, 0x7b, 0xfe, 0x9e, 0x4a, 0x7f, 0xdc, 0x44, 0xbf, 0x47, 0x25, 0x73, 0x39,
	0x69, 0x0c, 0x9c, 0x5d, 0xf3, 0x10, 0x9a, 0xc2, 0x64, 0x00, 0x67, 0x9d, 0xee, 0x95, 0xd9, 0x84,
	0x0f, 0x57, 0xc2, 0x66, 0xf5, 0xe0, 0x77, 0x7d, 0x57, 0x0b, 0xeb, 0xbf, 0x04, 0xc6, 0x8e, 0xb6,
	0xdd, 0xde, 0x1c, 0x06, 0x58, 0x7d, 0x73, 0xb7, 0xb6, 0x0c, 0xc5, 0x3f, 0x05, 0xf2, 0xda, 0xdd,
	0x66, 0x5a, 0x33, 0xb0, 0xdb, 0xdd, 0xec, 0xa9, 0x86, 0x09, 0x3c, 0x59, 0x6f, 0x29, 0x64, 0x5c,
	0xb9, 0xe2, 0xc7, 0xba, 0xe3, 0x19, 0xc9, 0x80, 0x44, 0xbf, 0xc7, 0xa4, 0x08, 0xbe, 0xa2, 0xf9,
	0x77, 0xfa, 0x86, 0x02, 0xce, 0xfa, 0x2c, 0x06, 0x60, 0x7a, 0xf3, 0x26, 0x5e, 0x20, 0xbc, 0x8f,
	0xaf, 0x1a, 0xfa, 0x30, 0x2b, 0xb3, 0xb6, 0x1c, 0x44, 0xa0, 0x72, 0x81, 0x86, 0x3d, 0xa9, 0xc3,
	0x04, 0x11, 0x62, 0x65, 0x84, 0x43, 0xe3, 0x5e, 0x77, 0x5f, 0xab, 0xc7, 0xbd, 0x44, 0xcc, 0x8b,
	0x5b, 0x8c, 0x5d, 0xaa, 0x04, 0xd1, 0x42, 0xa1, 0xf2, 0xcc, 0xc1, 0x9f, 0x33, 0xed, 0xda, 0xac,
	0xa0, 0x97, 0xee, 0x10, 0xf3, 0x12, 0xea, 0x3b, 0xea, 0xf4, 0xac, 0xd1, 0x61, 0xf8, 0xf1, 0x54,
	0x46, 0xa1, 0x39, 0x44, 0xc9, 0x1e, 0x2e, 0xe7, 0x7f, 0x19, 0xbc, 0x88, 0xdd, 0x5c, 0xbe, 0x15,
	0xf0, 0xbb, 0xee, 0x20, 0x8f, 0x68, 0x97, 0x36, 0x8c, 0x85, 0x80, 0x58, 0xb4, 0xb5, 0x65, 0xaa,
	0x27, 0xd8, 0xb3, 0xc7, 0xaa, 0xe6, 0x91, 0xc4, 0xfc, 0x9a, 0x99, 0x31, 0xf0, 0x3b, 0x85, 0xb8,
	0x88, 0xc3, 0x71, 0xec, 0x0c, 0x75, 0x42, 0x19, 0x3a, 0xe6, 0x2f, 0x5c, 0xcf, 0x48, 0x89, 0x3c,
	0x4a, 0xb9, 0xc5, 0xa8, 0x24, 0xc6, 0xf7, 0xe4, 0xb0, 0x87, 0xd9, 0xdd, 0x15, 0x9b, 0x81, 0x40,
	0xf1, 0xed, 0xb4, 0x8e, 0xed, 0xa7, 0x2b, 0xdf, 0x78, 0x08, 0x58, 0xc2, 0xf7, 0xad, 0x6c, 0xf8,
	0x19, 0x02, 0xa1, 0x4a, 0x86, 0xbc, 0x1f, 0x8c, 0x81, 0xdf, 0x7f, 0xaf, 0xdb, 0xd9, 0xb1, 0xd5,
	0x49, 0x68, 0x4a, 0xaf, 0xbc, 0xda, 0xee, 0xe6, 0x37, 0x94, 0xaf, 0x97, 0x16, 0xf4, 0x5e, 0xc5,
	0xca, 0x1d, 0xcd, 0x07, 0xe1, 0xef, 0x39, 0x97, 0x0c, 0xb6, 0xfe, 0x64, 0xc6, 0x81, 0xc9, 0x10,
	0x50, 0x82, 0xa3, 0xa9, 0x76, 0xbe, 0x74, 0xd9, 0xf1, 0x71, 0x3a, 0xf8, 0x75, 0xa9, 0xd2, 0x5f,
	0x58, 0xdd, 0x06, 0x27, 0x53, 0xf9, 0xa2, 0xbe, 0xaf, 0xaa, 0xd7, 0x95, 0x7c, 0x29, 0xf5, 0x51,
	0x5a, 0x59, 0xae, 0xc5, 0xbf, 0x07, 0x09, 0xc7, 0x5a, 0x04, 0x2a, 0x9e, 0x26, 0xab, 0x7b, 0xd3,
	0x85, 0x9d, 0x74, 0xb2, 0xea, 0x4b, 0x37, 0x44, 0xea, 0xf3, 0xcd, 0xda, 0x41, 0x9c, 0x7b, 0xb7,
	0xff, 0x82, 0x14, 0xef, 0xee, 0x4a, 0xfb, 0x20, 0x4c, 0x43, 0xfa, 0xa8, 0xaa, 0xb0, 0x5d, 0x2d,
	0xc3, 0x18, 0x4b, 0x5a, 0x58, 0x6e, 0x9e, 0x9f, 0x21, 0x5c, 0xbc, 0x14, 0x91, 0x92, 0x85, 0x94,
	0x6a, 0xbd, 0x99, 0x2c, 0x96, 0xdd, 0x8f, 0x39, 0x2f, 0x85, 0xc4, 0x27, 0x2f, 0x13, 0xcc, 0x4f,
	0xc4, 0x97, 0x1c, 0x34, 0xa8, 0x80, 0x54, 0x11, 0x27, 0x38, 0xc0, 0x7f, 0x3c, 0xd4, 0xcc, 0x73,
	0xd2, 0x39, 0x49, 0x02, 0xda, 0x95, 0xcd, 0x23, 0xa4, 0xc1, 0x5b, 0xdf, 0xc1, 0x53, 0x07, 0x6e,
	0xe1, 0xc8, 0xf1, 0x3c, 0x09, 0x25, 0x32, 0xea, 0xe4, 0x75, 0xbc, 0x43, 0x5e, 0x93, 0x05, 0x04,
	0xb5, 0xaa, 0xeb, 0x46, 0x7f, 0xa4, 0xb8, 0x60, 0x75, 0xcb, 0x70, 0x7a, 0xdf, 0x92, 0xea, 0x04,
	0x1b, 0xb1, 0x4a, 0xfc, 0x04, 0x18, 0x41, 0x4a, 0x15, 0x3d, 0xbc, 0x13, 0x9e, 0xc3, 0x35, 0xf6,
	0x49, 0x9e, 0xd7, 0x9c, 0xa1, 0xdc, 0x4d, 0xed, 0x01, 0x38, 0x6b, 0xe9, 0x5c, 0x3f, 0xe6, 0xef,
	0x5b, 0x80, 0x4a, 0x6b, 0x53, 0xee, 0x1d, 0x65, 0x9b, 0x96, 0x23, 0x58, 0x31, 0x88, 0x14, 0x94,
	0xb6, 0x71, 0xa9, 0x64, 0xde, 0x58, 0x20, 0x05, 0x58, 0x43, 0xf7, 0xaf, 0xa8, 0x1f, 0x83, 0x96,
	0xe0, 0x79, 0x06, 0xab, 0xae, 0x14, 0xf9, 0x20, 0xea, 0xdb, 0xc7, 0xe6, 0x55, 0x27, 0x97, 0x6d,
	0x92, 0x90, 0x64, 0x94, 0x94, 0x74, 0xe4, 0x36, 0x03, 0x1f, 0x32, 0x59, 0x5e, 0xa2, 0xe9, 0x3e,
	0xf4, 0x5e, 0xc4, 0xbf, 0x8f, 0x96, 0xb0, 0xac, 0x4d, 0x0f, 0xc2, 0x02, 0x28, 0x3e, 0xb1, 0xa1,
	0xc3, 0x1e, 0x3c, 0x30, 0x0f, 0x13, 0x8a, 0x77, 0xf6, 0x2d, 0xb4, 0xb4, 0xce, 0x22, 0x43, 0xbc,
	0xa0, 0xa5, 0x45, 0x73, 0x0a, 0x10, 0x2b, 0x08, 0x5f, 0x46, 0x2b, 0x1e, 0xf9, 0xba, 0x56, 0x75,
	0x45, 0xcd, 0xc5, 0x65, 0xc4, 0x85, 0x47, 0xcf, 0x5b, 0xc9, 0x66, 0x0f, 0x78, 0x37, 0x81, 0x10,
	0x6a, 0x44, 0x12, 0xec, 0x04, 0x78, 0x46, 0x89, 0xfc, 0x16, 0xcc, 0x6e, 0x1a, 0x50, 0x39, 0xf4,
	0xeb, 0x3f, 0xf1, 0x38, 0xb3, 0xb1, 0xac, 0x5d, 0x68, 0xbd, 0x59, 0x03, 0x4f, 0xb8, 0x0a, 0x1b,
	0x1a, 0x47, 0xe0, 0xd6, 0xb2, 0x4f, 0xa7, 0x20, 0x91, 0xfb, 0x42, 0x27, 0x3c, 0x40, 0xa4, 0x3e,
	0xef, 0x9c, 0xdd, 0xbe, 0x01, 0xd8, 0xe6, 0x1d, 0xaa, 0xef, 0x12, 0x1b, 0xa5, 0xbe, 0xb1, 0x02,
	0x14, 0xda, 0x9b, 0x30, 0xb5, 0x74, 0x23, 0x83, 0x4f, 0x7e, 0x02, 0xf1, 0xe2, 0x27, 0xe8, 0x7b,
	0x87, 0x69, 0x78, 0xe7, 0xa2, 0xd9, 0x77, 0x74, 0x99, 0x3b, 0xee, 0xbc, 0x87, 0x8b, 0xed, 0x2d,
	0x8d, 0x7e, 0x1b, 0xe9, 0xfa, 0xc6, 0x34, 0x2f, 0xdc, 0x46, 0xa0, 0x7d, 0x6b, 0xf9, 0xa7, 0xd5,
	0x64, 0x7d, 0x9a, 0x9b, 0x85, 0xd3, 0x28, 0x80, 0xa5, 0x23, 0x1b, 0xf9, 0xb0, 0xf0, 0x49, 0x6b,
	0x34, 0xe6, 0x97, 0xaa, 0xba, 0x8c, 0x34, 0x1e, 0x99, 0xae, 0x80, 0xf1, 0x15, 0xeb, 0x50, 0x6a,
	0xf7, 0x46, 0x10, 0x2a, 0x5d, 0xac, 0x23, 0x19, 0x02, 0xa6, 0xd1, 0x5d, 0x8a, 0x15, 0x12, 0x0c,
	0x4d, 0x79, 0xe6, 0xd8, 0x44, 0x24, 0x84, 0x02, 0xb4, 0xa6, 0xab, 0x64, 0x58, 0x7f, 0x47, 0x12,
	0x05, 0xe2, 0x85, 0x6a, 0x0e, 0x5e, 0x90, 0x4e, 0x01, 0x1d, 0xa7, 0x41, 0x11, 0xe7, 0x8a, 0x7f,
	0x72, 0x28, 0x54, 0xe8, 0x5a, 0xda, 0x74, 0x25, 0x55, 0x5d, 0x5e, 0x18, 0x4e, 0x7d, 0xc2, 0x05,
	0xa5, 0x18, 0x9f, 0xe3, 0xfa, 0x44, 0xf4, 0xdd, 0xd7, 0x90, 0xf2, 0x3c, 0x9d, 0x23, 0x0e, 0x22,
	0xfd, 0x7a, 0x1f, 0xa9, 0x0d, 0xf7, 0xec, 0xf1, 0x0f, 0xcf, 0x15, 0x29, 0x00, 0x36, 0x07, 0xa2,
	0x49, 0x55, 0x88, 0xc7, 0x24, 0x32, 0x8a, 0x64, 0x32, 0x0d, 0x0a, 0xf0, 0x0b, 0xce, 0x7c, 0x8b,
	0x90, 0xdf, 0x13, 0x3c, 0x8f, 0x69, 0x0a, 0x50, 0x65, 0x63, 0xe0, 0x47, 0xa8, 0x7f, 0x61, 0xfe,
	0x77, 0xec, 0x9a, 0xc2, 0x6f, 0xe8, 0x7d, 0x8f, 0xe5, 0x27, 0xd9, 0x17, 0x37, 0x38, 0x4e, 0x9c,
	0x2e, 0x43, 0x49, 0x4c, 0x7f, 0xd7, 0x7d, 0x35, 0xa2, 0x04, 0xa8, 0xda, 0xbe, 0x9e, 0xf1, 0xba,
	0x45, 0xfb, 0x5b, 0x90, 0xea, 0xfd, 0x93, 0x41, 0x43, 0x19, 0x14, 0xd0, 0x41, 0xcd, 0x2f, 0x87,
	0x87, 0xd0, 0x02, 0xc6, 0x2c, 0x6d, 0x18, 0xa7, 0x6f, 0x84, 0x1a, 0xb5, 0xa0, 0x48, 0x45, 0x09,
	0x27, 0xe6, 0xb8, 0x88, 0x7b, 0x2b, 0xb1, 0x1a, 0xa8, 0xf6, 0xcd, 0x27, 0x68, 0xaa, 0xd1, 0xd1,
	0x26, 0x80, 0x1c, 0xb9, 0x86, 0xea, 0xfe, 0x41, 0x2f, 0x37, 0x5e, 0xad, 0x35, 0x54, 0xcd, 0x15,
	0x50, 0xfa, 0xd6, 0x42, 0x95, 0xad, 0x49, 0x7c, 0x20, 0xeb, 0x31, 0x6f, 0x06, 0x05, 0x74, 0xc8,
	0xfc, 0x7e, 0x51, 0x13, 0xec, 0x83, 0x2f, 0x5a, 0x87, 0x90, 0x27, 0x78, 0x61, 0xa6, 0xee, 0xb7,
	0x51, 0xce, 0x9e, 0x5a, 0xb2, 0x59, 0xd4, 0xdd, 0xaf, 0x93, 0x40, 0xea, 0xc9, 0xc6, 0xa7, 0x19,
	0x6a, 0x2b, 0x36, 0x3c, 0xbc, 0x33, 0x36, 0xb7, 0x0e, 0x46, 0x8d, 0x7f, 0xa0, 0x68, 0xbb, 0x9b,
	0x7b, 0x26, 0x21, 0xf5, 0x10, 0x47, 0xb5, 0xfd, 0x98, 0x78, 0x39, 0x94, 0xa9, 0xe1, 0x28, 0x69,
	0xe4, 0x84, 0x9f, 0x52, 0x5d, 0x71, 0xe4, 0x41, 0x49, 0x4e, 0x5b, 0x45, 0x85, 0x4f, 0x92, 0xac,
	0x69, 0xd8, 0xf5, 0x84, 0xa0, 0x84, 0x59, 0x37, 0x16, 0x6f, 0x47, 0xa5, 0xa3, 0xf1, 0x28, 0xd5,
	0xe1, 0xcf, 0x9e, 0xa8, 0x04, 0x4d, 0x0d, 0xd1, 0x3d, 0x8e, 0x8c, 0x80, 0x0d, 0x77, 0x3f, 0xe7,
	0x1e, 0x27, 0xb6, 0xe7, 0x59, 0x47, 0x45, 0x4c, 0xd6, 0xee, 0x2f, 0xe1, 0xee, 0xd0, 0xbc, 0x03,
	0xa4, 0x4c, 0x69, 0x4f, 0xa6, 0x90, 0x6b, 0xdf, 0x15, 0x16, 0x2e, 0xac, 0x04, 0xc6, 0xc6, 0xa2,
	0xb9, 0x15, 0xcb, 0x9b, 0x47, 0xc2, 0x70, 0xcb, 0x49, 0x43, 0x1c, 0x3b, 0x1c, 0x72, 0x80, 0x3f,
	0x21, 0xba, 0x2e, 0x0e, 0x8a, 0x5f, 0xd5, 0xfd, 0x22, 0xa9, 0x37, 0x27, 0x7b, 0xd7, 0xe2, 0xf4,
	0xe6, 0x21, 0xad, 0xca, 0x8a, 0xab, 0x37, 0xba, 0x2f, 0xb3, 0xac, 0x3d, 0x3e, 0x96, 0xc3, 0xd1,
	0xc0, 0x4c, 0xd7, 0x1e, 0x64, 0xb8, 0x0b, 0x9d, 0x75, 0x18, 0x24, 0x9a, 0x26, 0x49, 0xe7, 0x96,
	0x4c, 0xd7, 0x9f, 0x56, 0x0e, 0x0b, 0xf0, 0xf3, 0xcf, 0x79, 0xa7, 0x61, 0x2c, 0xea, 0xc4, 0xae,
	0xaf, 0xf3, 0xb0, 0xc1, 0xa6, 0x27, 0xd7, 0x2c, 0x97, 0xe7, 0xac, 0xe7, 0x9b, 0x5a, 0x2e, 0xe2,
	0x1f, 0x34, 0x5f, 0xd6, 0x53, 0x55, 0xf9, 0xed, 0x9d, 0x29, 0x74, 0x99, 0xd8, 0xf2, 0x0a, 0x55,
	0xe0, 0xd5, 0xa3, 0x93, 0x78, 0xcc, 0x12, 0x64, 0x5a, 0xed, 0x61, 0xc6, 0xfd, 0x85, 0x27, 0x52,
	0x28, 0x4e, 0x73, 0xad, 0x84, 0x45, 0x89, 0x7d, 0x4f, 0xa0, 0xbc, 0x31, 0x13, 0x7e, 0x28, 0xed,
	0xfa, 0xee, 0x52, 0xa4, 0x3f, 0xad, 0x62, 0xfc, 0xf2, 0xc6, 0xd1, 0x30, 0xb7, 0x69, 0xa1, 0xef,
	0xfc, 0xb8, 0x93, 0x4e, 0x69, 0xde, 0x81, 0x07, 0x44, 0x0c, 0xf3, 0x0e, 0xc2, 0xce, 0xc4, 0x56,
	0x2e, 0x63, 0xb0, 0x36, 0x31, 0x8d, 0x7b, 0x98, 0x7a, 0x1e, 0x89, 0x6b, 0xd6, 0x67, 0x56, 0x0f,
	0xbf, 0xb7, 0x68, 0x2b, 0x43, 0xc2, 0xfa, 0x5a, 0xaa, 0x95, 0x9c, 0x56, 0x8e, 0x53, 0x88, 0x90,
	0x28, 0x7c, 0xbc, 0xf7, 0xff, 0x2b, 0x50, 0xc5, 0x3e, 0x69, 0x1e, 0x61, 0x3e, 0xef, 0xb6, 0x61,
	0x88, 0xd6, 0x2b, 0x4c, 0x5b, 0xf5, 0x30, 0xc3, 0x82, 0x9c, 0x00, 0x33, 0x33, 0x2e, 0x94, 0xd0,
	0xda, 0x60, 0x33, 0xca, 0xd7, 0xb1, 0xf3, 0xf3, 0x0f, 0x45, 0x11, 0x22, 0x37, 0x68, 0x97, 0x86,
	0xd3, 0xf2, 0x00, 0x91, 0x82, 0x28, 0xf2, 0xac, 0xc9, 0x69, 0x6f, 0x13, 0xf3, 0x4b, 0x24, 0x73,
	0x44, 0x26, 0xe8, 0x64, 0xdd, 0x7f, 0x26, 0xba, 0x25, 0xac, 0x13, 0x73, 0x6b, 0x0f, 0x75, 0x69,
	0xf5, 0xfb, 0xfc, 0xaa, 0xf7, 0x04, 0x09, 0xfb, 0x90, 0xc4, 0x3b, 0x98, 0x95, 0x35, 0x9f, 0x70,
	0x0f, 0x29, 0x70, 0x43, 0x32, 0x92, 0xd3, 0xe7, 0xe5, 0xa6, 0xab, 0xb8, 0x56, 0x51, 0x75, 0xbd,
	0x67, 0x3f, 0xb3, 0x48, 0x22, 0x7a, 0xff, 0x7d, 0x70, 0x27, 0x59, 0x7f, 0x48, 0xe1, 0x7e, 0xf2,
	0x84, 0xd2, 0xc7, 0xd2, 0x55, 0x80, 0x3f, 0xf6, 0x00, 0xad, 0x82, 0xf2, 0xdf, 0xf5, 0x21, 0x15,
	0x3d, 0x87, 0xfc, 0xd5, 0xda, 0x83, 0x44, 0x86, 0x98, 0x5f, 0x97, 0x11, 0x3e, 0xaa, 0xf5, 0xfe,
	0x4b, 0xec, 0xc7, 0xf8, 0x53, 0xd5, 0x4b, 0x88, 0x81, 0xda, 0xff, 0x4a, 0xe0, 0x5d, 0x63, 0xe8,
	0x52, 0xdd, 0xfa, 0x91, 0xf3, 0xe9, 0xfd, 0x4d, 0x2f, 0x8a, 0x9e, 0x1f, 0x23, 0x00, 0xea, 0x74,
	0xdb, 0xfa, 0xff, 0xd0, 0x0f, 0x14, 0xc2, 0x87, 0x7a, 0x8b, 0xed, 0x36, 0xfe, 0xee, 0x90, 0x47,
	0x59, 0xb9, 0xcd, 0xb3, 0x9f, 0x90, 0x7f, 0xbe, 0x5e, 0x7f, 0x65, 0xdf, 0xe9, 0xd2, 0x51, 0xc6,
	0xd2, 0x3a, 0xb3, 0x36, 0xe2, 0x92, 0x3f, 0x6b, 0x1e, 0x9e, 0xea, 0xf9, 0x9d, 0xa9, 0x58, 0xfe,
	0x6f, 0xb6, 0x02, 0xf5, 0x66, 0x28, 0x7c, 0xfd, 0x90, 0x47, 0xd9, 0xe2, 0xb5, 0x61, 0xbf, 0x65,
	0x57, 0x15, 0xd6, 0xf8, 0xe0, 0x2f, 0xbb, 0x9d, 0xc2, 0x21, 0x6a, 0x34, 0x86, 0x45, 0x16, 0xb0,
	0xc1, 0x92, 0x98, 0xb1, 0xba, 0xf6, 0x7a, 0x84, 0x62, 0xc5, 0xa9, 0xae, 0x8f, 0xa5, 0x2b, 0x49,
	0x9a, 0x4a, 0x2a, 0xa9, 0xc7, 0x5e, 0x45, 0xb8, 0x8a, 0xe7, 0xfd, 0x28, 0xdd, 0x29, 0x4a, 0x90,
	0x4f, 0xef, 0xb9, 0x81, 0x2d, 0x1a, 0xf0, 0x2c, 0xaa, 0x9d, 0x8d, 0x87, 0xe0, 0x2b, 0x9f, 0x48,
	0x1e, 0xfc, 0xed, 0x8f, 0x1e, 0xe7, 0xc9, 0xa8, 0x00, 0x39, 0x30, 0xca, 0x13, 0x2e, 0xea, 0xf7,
	0xd8, 0x91, 0xa9, 0x7a, 0xdd, 0x2f, 0xb6, 0x92, 0x32, 0x29, 0x80, 0xff, 0x8f, 0x85, 0x6f, 0xdb,
	0xe5, 0xa1, 0xb0, 0x74, 0xb0, 0x38, 0x5c, 0x44, 0xcd, 0xfb, 0x3a, 0x8e, 0x60, 0x64, 0x7e, 0x0d,
	0x6e, 0x2d, 0xfe, 0x92, 0x5d, 0x5a, 0x40, 0xa1, 0x9e, 0x94, 0x5f, 0x82, 0xaf, 0x62, 0x11, 0x9d,
	0x5e, 0x6f, 0x04, 0xf8, 0x4e, 0xf3, 0x00, 0xe9, 0xcf, 0x6a, 0xac, 0x07, 0x96, 0xf4, 0x5d, 0x0a,
	0x1b, 0x91, 0x19, 0x85, 0xb7, 0x31, 0x16, 0x66, 0x58, 0x7e, 0xde, 0x19, 0xc1, 0x0c, 0xdc, 0x5c,
	0x7c, 0xb4, 0x91, 0xe0, 0x39, 0x72, 0x18, 0x4e, 0x62, 0x84, 0x45, 0xb6, 0x16, 0xb8, 0xac, 0x01,
	0x14, 0xc9, 0xad, 0x11, 0xc8, 0xea, 0xf7, 0x24, 0xf8, 0x5a, 0x84, 0x8a, 0xdf, 0x07, 0xb7, 0xed,
	0xd4, 0x63, 0x5e, 0xea, 0x44, 0x24, 0xcd, 0x2c, 0x17, 0x3b, 0x26, 0x25, 0x70, 0x34, 0x46, 0x0d,
	0xac, 0x98, 0x9b, 0x27, 0x35, 0x86, 0x88, 0x10, 0x4b, 0xdf, 0x16, 0x43, 0x82, 0x7e, 0x0d, 0xd0,
	0xab, 0x97, 0x06, 0xd5, 0xb7, 0x4a, 0x3d, 0x15, 0x3b, 0x7b, 0x28, 0x04, 0xbe, 0xa3, 0x70, 0x38,
	0xa4, 0x13, 0x1d, 0x59, 0x7b, 0x12, 0x4d, 0x57, 0x00, 0xcf, 0xfe, 0xb4, 0x49, 0x3b, 0x05, 0x74,
	0xff, 0xe3, 0x24, 0xba, 0x15, 0xd0, 0xad, 0x80, 0x6e, 0x05, 0x74, 0x2b, 0xa0, 0x5b, 0x01, 0xe7,
	0xea, 0xf1, 0x6f, 0x01, 0x06, 0x00, 0xaa, 0x78, 0x85, 0x03, 0xcb, 0x00, 0x92, 0x3d, 0x00, 0x00,
	0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}

func Favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/vnd.microsoft.icon")
	w.Write(icon)
}
