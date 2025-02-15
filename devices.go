package majin

import (
	// "fmt"

	// "context"
	// "log"
	// "net/http"

	// devicev1 "github.com/yasha-gh/meta-majin/gen/device/v1"
	userv1 "github.com/yasha-gh/meta-majin/gen/user/v1"
	// "github.com/yasha-gh/meta-majin/gen/user/v1/userv1connect"

	// "connectrpc.com/connect"
)

type Device userv1.Device

// func (m *MetaMajin) ListDevices() (devices []*Device, err error) {
// 	fmt.Println("Get connections", m.GetConnectString())
// 	client := userv1connect.NewUserServiceClient(
// 		http.DefaultClient,
// 		m.GetConnectString(),
// 	)
// 	res, err := client.ListDevices(context.Background(), connect.NewRequest(&devicev1.ListDevicesRequest{}))
// 	if err != nil {
// 		log.Println("Failed to get devices", err)
// 		return nil, err
// 	}
// 	fmt.Println("Devices result", res.Msg)
// 	// (*devicev1.Device)(dev)
// 	devices = make([]*Device, 0)
// 	for _, device := range res.Msg.Devices {
// 		devices = append(devices, (*Device)(device))
// 	}
// 	return devices, nil
// }
