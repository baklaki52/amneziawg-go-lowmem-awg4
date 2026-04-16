//go:build !android && !ios && !windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2025 WireGuard LLC. All Rights Reserved.
 */

package device

import "github.com/amnezia-vpn/amneziawg-go/conn"

const (
	QueueStagedSize = conn.IdealBatchSize
	MaxSegmentSize  = (1 << 16) - 1 // largest possible UDP datagram
)

var (
	QueueOutboundSize          = 1024
	QueueInboundSize           = 1024
	QueueHandshakeSize         = 1024
	PreallocatedBuffersPerPool = uint32(0) // Disable and allow for infinite memory growth
)
