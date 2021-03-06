// Downloaded from https://github.com/dgryski/go-stadtx/raw/3c3d9b328c24a9b5ecd370654cd6e9d60a85752d/stadtx.go

// Package stadtx implements Stadtx Hash
/*

   https://github.com/demerphq/BeagleHash

*/
package stadtx

func rotl64(x uint64, r uint64) uint64 {
	return (((x) << (r)) | ((x) >> (64 - r)))
}

func rotr64(x uint64, r uint64) uint64 {
	return (((x) >> (r)) | ((x) << (64 - r)))
}

func scramble64(v, prime uint64) uint64 {
	v ^= (v >> 13)
	v ^= (v << 35)
	v ^= (v >> 30)
	v *= prime
	v ^= (v >> 19)
	v ^= (v << 15)
	v ^= (v >> 46)
	return v
}

func SeedState(seed []uint64) State {

	var state State

	state[0] = seed[0] ^ 0x43f6a8885a308d31
	state[1] = seed[1] ^ 0x3198a2e03707344a
	state[2] = seed[0] ^ 0x4093822299f31d00
	state[3] = seed[1] ^ 0x82efa98ec4e6c894

	if state[0] == 0 {
		state[0] = 1
	}
	if state[1] == 0 {
		state[1] = 2
	}
	if state[2] == 0 {
		state[2] = 4
	}
	if state[3] == 0 {
		state[3] = 8
	}

	state[0] = scramble64(state[0], 0x801178846e899d17)
	state[0] = scramble64(state[0], 0xdd51e5d1c9a5a151)
	state[1] = scramble64(state[1], 0x93a7d6c8c62e4835)
	state[1] = scramble64(state[1], 0x803340f36895c2b5)
	state[2] = scramble64(state[2], 0xbea9344eb7565eeb)
	state[2] = scramble64(state[2], 0xcd95d1e509b995cd)
	state[3] = scramble64(state[3], 0x9999791977e30c13)
	state[3] = scramble64(state[3], 0xaab8b6b05abfc6cd)

	return state
}

type State [4]uint64
