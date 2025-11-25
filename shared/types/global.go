// SPDX-FileCopyrightText: 2025 Massimo Ambrosi
//
// SPDX-License-Identifier: Apache-2.0

package types

// GlobalFlags represents the flags used by all commands.
type GlobalFlags struct {
	ConfigPath  string
	LogLevel    string
	KeepTempDir bool
}
