/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2025 WireGuard LLC. All Rights Reserved.
 */

package device

import "github.com/snehesht/wireguard-go/conn"

/* Reduce memory consumption for Android */

const (
	QueueStagedSize    = conn.IdealBatchSize
	QueueHandshakeSize = 1024
	MaxSegmentSize     = (1 << 16) - 1 // largest possible UDP datagram
)

// Sizes a consumer may set before NewDevice, as iOS already does: the per-peer
// outbound and the inbound queue depths, and the bound on every buffer pool
// (0 is unbounded, which lets a receiver that falls behind hold QueueInboundSize
// batches of 64 KB buffers). wirepeer's fork; upstream keeps them const here.
var (
	QueueOutboundSize          int    = 1024
	QueueInboundSize           int    = 1024
	PreallocatedBuffersPerPool uint32 = 4096
)
