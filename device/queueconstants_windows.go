/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2025 WireGuard LLC. All Rights Reserved.
 */

package device

const (
	QueueStagedSize    = 128
	QueueHandshakeSize = 1024
	MaxSegmentSize     = 2048 - 32 // largest possible UDP datagram
)

// Sizes a consumer may set before NewDevice: the per-peer outbound and the
// inbound queue depths, and the bound on every buffer pool (0 is unbounded,
// which lets a receiver that falls behind hold QueueInboundSize times the
// batch size of 64 KB buffers). wirepeer's fork; upstream keeps them const.
var (
	QueueOutboundSize          int    = 1024
	QueueInboundSize           int    = 1024
	PreallocatedBuffersPerPool uint32 = 0
)
