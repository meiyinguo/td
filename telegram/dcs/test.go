package dcs

import "github.com/gotd/td/tg"

// Staging returns staging DC list.
//
// Deprecated: Use Test().
func Staging() List {
	return Test()
}

// Test returns test DC list.
func Test() List {
	return List{
		Options: []tg.DCOption{
			{
				ID:        1,
				IPAddress: "149.154.175.10",
				Port:      443,
			},
			{
				ID:        1,
				Ipv6:      true,
				IPAddress: "2001:0b28:f23d:f001:0000:0000:0000:000e",
				Port:      443,
			},
			{
				ID:        2,
				IPAddress: "149.154.167.40",
				Port:      443,
			},
			{
				ID:        2,
				Ipv6:      true,
				IPAddress: "2001:067c:04e8:f002:0000:0000:0000:000e",
				Port:      443,
			},
			{
				ID:        3,
				IPAddress: "149.154.175.117",
				Port:      443,
			},
			{
				ID:        3,
				Ipv6:      true,
				IPAddress: "2001:0b28:f23d:f003:0000:0000:0000:000e",
				Port:      443,
			},
		},
		Domains: map[int]string{
			1: "wss://pluto.web.telegram.org/apiws_test",
			2: "wss://venus.web.telegram.org/apiws_test",
			3: "wss://aurora.web.telegram.org/apiws_test",
			4: "wss://vesta.web.telegram.org/apiws_test",
			5: "wss://flora.web.telegram.org/apiws_test",
		},
		Test: true,
	}
}

func Live() List {
	return List{
		Options: []tg.DCOption{
			{
				ID:        1,
				IPAddress: "149.154.175.50",
				Port:      443,
			},
			{
				ID:        1,
				Ipv6:      true,
				IPAddress: "2001:b28:f23d:f001:0000:0000:0000:000a",
				Port:      443,
			},
			{
				ID:        2,
				IPAddress: "149.154.167.51",
				Port:      443,
			},
			{
				ID:        2,
				Ipv6:      true,
				IPAddress: "2001:67c:4e8:f002:0000:0000:0000:000a",
				Port:      443,
			},
			{
				ID:        3,
				IPAddress: "149.154.175.100",
				Port:      443,
			},
			{
				ID:        3,
				Ipv6:      true,
				IPAddress: "2001:b28:f23d:f003:0000:0000:0000:000a",
				Port:      443,
			},
			{
				ID:        4,
				IPAddress: "149.154.167.91",
				Port:      443,
			},
			{
				ID:        4,
				Ipv6:      true,
				IPAddress: "2001:67c:4e8:f004:0000:0000:0000:000a",
				Port:      443,
			},
			{
				ID:        5,
				IPAddress: "149.154.171.5",
				Port:      443,
			},
			{
				ID:        4,
				Ipv6:      true,
				IPAddress: "2001:b28:f23f:f005:0000:0000:0000:000a",
				Port:      443,
			},
		},
		Test: false,
	}

}
