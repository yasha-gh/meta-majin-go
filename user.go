package main

import (
	"fmt"
	"strings"

	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	userv1 "github.com/yasha-gh/meta-majin/gen/user/v1"
	"github.com/yasha-gh/meta-majin/gen/user/v1/userv1connect"
)

type User userv1.User

type DeviceuserInfo userv1.DeviceUserInfo
func (m *User) Me() *User {
	return m
}

func (m *MetaMajin) DeviceUserInfo() (deviceUserInfo DeviceuserInfo, err error) {
	client := m.Clients.User
	// client := userv1connect.NewUserServiceClient(
	// 	http.DefaultClient,
	// 	m.GetConnectString(),
	// )
	res, err := client.DeviceUserInfo(context.Background(), connect.NewRequest(&userv1.DeviceUserInfoRequest{}))
	if err != nil {
		log.Println("Failed to get current user", err)
		return DeviceuserInfo{}, err
	}
	fmt.Println("Current user result", res.Msg)
	return DeviceuserInfo{
		Username: res.Msg.DeviceUserInfo.Username,
		UserId: res.Msg.DeviceUserInfo.UserId,
		Os: res.Msg.DeviceUserInfo.Os,
		OsDescription: res.Msg.DeviceUserInfo.OsDescription,
		Hostname: res.Msg.DeviceUserInfo.Hostname,
	}, nil
}

// message AddUserRequest {
//     // User fields
//     string username = 1; // Meta majin username
//     optional string email = 3;
//     optional string first_name = 4;
//     optional string last_name = 5;
//     // Device fields
//     optional string device_user_id = 6; // Device specific user ID (e.g. Linux user ID)
//     optional string device_username = 7; // Device specific username (e.g. Linux username)
//     AddDeviceRequest local_device = 8;
// }
type AddUserParams struct {
	Username string `json:"username,required"`
	Email *string `json:"email"`
	FirstName *string `json:"firstName"`
	LastName *string `json:"lastName"`
	LocalDevice AddDeviceParams `json:"addDeviceParams,required"`
}
// func (a *AddUserParams) Void() *AddUserParams { return a }

func (m *MetaMajin) AddUser(newUser AddUserParams) (user User, err error) {
	if strings.TrimSpace(newUser.Username) == "" {
		return User{}, fmt.Errorf("Invalid input, username is required")
	}
	// fmt.Println("Get connections", m.GetConnectString())
	client := userv1connect.NewUserServiceClient(
		http.DefaultClient,
		m.GetConnectString(),
	)
	reqData := userv1.AddUserRequest{
		Username: newUser.Username,
		Email: newUser.Email,
		FirstName: newUser.FirstName,
		LastName: newUser.LastName,
		LocalDevice: &userv1.AddDeviceRequest{
			DisplayName: newUser.LocalDevice.DisplayName,
			Hostname: newUser.LocalDevice.Hostname,
			TimezoneCode: &newUser.LocalDevice.TimezoneCode,
			TimezoneName: &newUser.LocalDevice.TimezoneName,
			TimezoneOffset: &newUser.LocalDevice.TimezoneOffset,
			DeviceUsername: newUser.LocalDevice.DeviceUsername,
			DeviceUserId: newUser.LocalDevice.DeviceUserId,
		},
	}
	res, err := client.AddUser(context.Background(), connect.NewRequest(&reqData))
	if err != nil {
		log.Println("Failed to add user", err)
		return User{}, err
	}
	fmt.Println("Add user result", res.Msg)
	return User{
		Id: res.Msg.User.Id,
		Username: res.Msg.User.Username,
		Email: res.Msg.User.Email,
		FirstName: res.Msg.User.FirstName,
		LastName: res.Msg.User.LastName,
		Devices: []*userv1.Device{},
	}, nil
}

type AddDeviceParams struct {
	UserId string `json:"userId"` // Meta Majin user ID (Current host)
    DisplayName string `json:"displayName,required"` // Current device display name
    Hostname string `json:"hostname,required"`
    TimezoneCode string `json:"timezoneCode"`
    TimezoneOffset int32 `json:"timezoneOffset"`
    TimezoneName string `json:"timezoneName"`
    DeviceUsername string `json:"deviceUsername"`
    DeviceUserId string `json:"deviceUserId"`
}
func (m *MetaMajin) AddDevice(newDevice AddDeviceParams) (user Device, err error) {
	if newDevice.UserId == "" || newDevice.DisplayName == "" || newDevice.Hostname == "" {
			return Device{}, fmt.Errorf("Invalid input")
		}
		// fmt.Println("Get connections", m.GetConnectString())
		client := userv1connect.NewUserServiceClient(
			http.DefaultClient,
			m.GetConnectString(),
		)
		tzOffset := int32(newDevice.TimezoneOffset)
		reqData := userv1.AddDeviceRequest{
			UserId: &newDevice.UserId,
			DisplayName: newDevice.DisplayName,
			Hostname: newDevice.Hostname,
			TimezoneCode: &newDevice.TimezoneCode,
			TimezoneOffset: &tzOffset,
			TimezoneName: &newDevice.TimezoneName,
		}
		// reqData := userv1.AddUserRequest{
		// 	Username: newUser.Username,
		// 	Email: newUser.Email,
		// 	FirstName: newUser.FirstName,
		// 	LastName: newUser.LastName,
		// }
		res, err := client.AddDevice(context.Background(), connect.NewRequest(&reqData))
		if err != nil {
			log.Println("Failed to add user", err)
			return Device{}, err
		}
		fmt.Println("Add device result", res.Msg)
		return Device{
			UserId: res.Msg.Device.UserId,
			DisplayName: res.Msg.Device.DisplayName,
			Hostname: res.Msg.Device.Hostname,
			TimezoneCode: res.Msg.Device.TimezoneCode,
			TimezoneName: res.Msg.Device.TimezoneName,
			TimezoneOffset: res.Msg.Device.TimezoneOffset,
		}, nil
}
