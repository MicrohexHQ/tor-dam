package damlib

/*
 * Copyright (c) 2017-2018 Dyne.org Foundation
 * tor-dam is written and maintained by Ivan J. <parazyd@dyne.org>
 *
 * This file is part of tor-dam
 *
 * This source code is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This software is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this source code. If not, see <http://www.gnu.org/licenses/>.
 */

import "os"

// Workdir holds the path to the directory where we will Chdir on startup.
var Workdir = os.Getenv("HOME") + "/.dam"

// RsaBits holds the size of our RSA private key. Tor standard is 1024.
const RsaBits = 1024

// PrivKeyPath holds the name of where our private key is.
const PrivKeyPath = "dam-private.key"

// SeedPath holds the name of where our public key is.
const SeedPath = "dam-private.seed"

// PubSubChan is the name of the pub/sub channel we're publishing to in Redis.
const PubSubChan = "tordam"

// PostMsg holds the message we are signing with our private key.
const PostMsg = "I am a DAM node!"

// WelcomeMsg holds the message we return when welcoming a node.
const WelcomeMsg = "Welcome to the DAM network!"

// ProxyAddr is the address of our Tor SOCKS port.
const ProxyAddr = "127.0.0.1:9050"

// TorPortMap is a comma-separated string holding the mapping of ports
// to be opened by the Tor Hidden Service. Format is "remote:local".
const TorPortMap = "80:49371,13010:13010,13011:13011,5000:5000"

// DirPort is the port where dam-dir will be listening.
const DirPort = 49371

// Testnet is flipped with a flag in dam-dir and represents if all new
// nodes are initially marked valid or not.
var Testnet = false

// Noremote is flipped with a flag in dam-client and disables fetching
// remote entry points (directories) if enabled.
var Noremote = false
