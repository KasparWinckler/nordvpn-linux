// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: settings.proto

package pb

import (
	config "github.com/NordSecurity/nordvpn-linux/config"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int64     `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Data *Settings `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SettingsResponse) Reset() {
	*x = SettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsResponse) ProtoMessage() {}

func (x *SettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsResponse.ProtoReflect.Descriptor instead.
func (*SettingsResponse) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{0}
}

func (x *SettingsResponse) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SettingsResponse) GetData() *Settings {
	if x != nil {
		return x.Data
	}
	return nil
}

type AutoconnectData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled     bool               `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Country     string             `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	City        string             `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	ServerGroup config.ServerGroup `protobuf:"varint,4,opt,name=server_group,json=serverGroup,proto3,enum=config.ServerGroup" json:"server_group,omitempty"`
}

func (x *AutoconnectData) Reset() {
	*x = AutoconnectData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoconnectData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoconnectData) ProtoMessage() {}

func (x *AutoconnectData) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoconnectData.ProtoReflect.Descriptor instead.
func (*AutoconnectData) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{1}
}

func (x *AutoconnectData) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AutoconnectData) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *AutoconnectData) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AutoconnectData) GetServerGroup() config.ServerGroup {
	if x != nil {
		return x.ServerGroup
	}
	return config.ServerGroup(0)
}

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Technology           config.Technology     `protobuf:"varint,1,opt,name=technology,proto3,enum=config.Technology" json:"technology,omitempty"`
	Firewall             bool                  `protobuf:"varint,2,opt,name=firewall,proto3" json:"firewall,omitempty"`
	KillSwitch           bool                  `protobuf:"varint,3,opt,name=kill_switch,json=killSwitch,proto3" json:"kill_switch,omitempty"`
	AutoConnectData      *AutoconnectData      `protobuf:"bytes,4,opt,name=auto_connect_data,json=autoConnectData,proto3" json:"auto_connect_data,omitempty"`
	Ipv6                 bool                  `protobuf:"varint,5,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	Meshnet              bool                  `protobuf:"varint,6,opt,name=meshnet,proto3" json:"meshnet,omitempty"`
	Routing              bool                  `protobuf:"varint,7,opt,name=routing,proto3" json:"routing,omitempty"`
	Fwmark               uint32                `protobuf:"varint,8,opt,name=fwmark,proto3" json:"fwmark,omitempty"`
	Analytics            bool                  `protobuf:"varint,9,opt,name=analytics,proto3" json:"analytics,omitempty"`
	Dns                  []string              `protobuf:"bytes,10,rep,name=dns,proto3" json:"dns,omitempty"`
	ThreatProtectionLite bool                  `protobuf:"varint,11,opt,name=threat_protection_lite,json=threatProtectionLite,proto3" json:"threat_protection_lite,omitempty"`
	Protocol             config.Protocol       `protobuf:"varint,12,opt,name=protocol,proto3,enum=config.Protocol" json:"protocol,omitempty"`
	LanDiscovery         bool                  `protobuf:"varint,13,opt,name=lan_discovery,json=lanDiscovery,proto3" json:"lan_discovery,omitempty"`
	Allowlist            *Allowlist            `protobuf:"bytes,14,opt,name=allowlist,proto3" json:"allowlist,omitempty"`
	Obfuscate            bool                  `protobuf:"varint,15,opt,name=obfuscate,proto3" json:"obfuscate,omitempty"`
	VirtualLocation      bool                  `protobuf:"varint,16,opt,name=virtualLocation,proto3" json:"virtualLocation,omitempty"`
	PostquantumVpn       bool                  `protobuf:"varint,17,opt,name=postquantum_vpn,json=postquantumVpn,proto3" json:"postquantum_vpn,omitempty"`
	UserSettings         *UserSpecificSettings `protobuf:"bytes,18,opt,name=user_settings,json=userSettings,proto3" json:"user_settings,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{2}
}

func (x *Settings) GetTechnology() config.Technology {
	if x != nil {
		return x.Technology
	}
	return config.Technology(0)
}

func (x *Settings) GetFirewall() bool {
	if x != nil {
		return x.Firewall
	}
	return false
}

func (x *Settings) GetKillSwitch() bool {
	if x != nil {
		return x.KillSwitch
	}
	return false
}

func (x *Settings) GetAutoConnectData() *AutoconnectData {
	if x != nil {
		return x.AutoConnectData
	}
	return nil
}

func (x *Settings) GetIpv6() bool {
	if x != nil {
		return x.Ipv6
	}
	return false
}

func (x *Settings) GetMeshnet() bool {
	if x != nil {
		return x.Meshnet
	}
	return false
}

func (x *Settings) GetRouting() bool {
	if x != nil {
		return x.Routing
	}
	return false
}

func (x *Settings) GetFwmark() uint32 {
	if x != nil {
		return x.Fwmark
	}
	return 0
}

func (x *Settings) GetAnalytics() bool {
	if x != nil {
		return x.Analytics
	}
	return false
}

func (x *Settings) GetDns() []string {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *Settings) GetThreatProtectionLite() bool {
	if x != nil {
		return x.ThreatProtectionLite
	}
	return false
}

func (x *Settings) GetProtocol() config.Protocol {
	if x != nil {
		return x.Protocol
	}
	return config.Protocol(0)
}

func (x *Settings) GetLanDiscovery() bool {
	if x != nil {
		return x.LanDiscovery
	}
	return false
}

func (x *Settings) GetAllowlist() *Allowlist {
	if x != nil {
		return x.Allowlist
	}
	return nil
}

func (x *Settings) GetObfuscate() bool {
	if x != nil {
		return x.Obfuscate
	}
	return false
}

func (x *Settings) GetVirtualLocation() bool {
	if x != nil {
		return x.VirtualLocation
	}
	return false
}

func (x *Settings) GetPostquantumVpn() bool {
	if x != nil {
		return x.PostquantumVpn
	}
	return false
}

func (x *Settings) GetUserSettings() *UserSpecificSettings {
	if x != nil {
		return x.UserSettings
	}
	return nil
}

type UserSpecificSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Notify bool  `protobuf:"varint,2,opt,name=notify,proto3" json:"notify,omitempty"`
	Tray   bool  `protobuf:"varint,3,opt,name=tray,proto3" json:"tray,omitempty"`
}

func (x *UserSpecificSettings) Reset() {
	*x = UserSpecificSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSpecificSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpecificSettings) ProtoMessage() {}

func (x *UserSpecificSettings) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSpecificSettings.ProtoReflect.Descriptor instead.
func (*UserSpecificSettings) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{3}
}

func (x *UserSpecificSettings) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UserSpecificSettings) GetNotify() bool {
	if x != nil {
		return x.Notify
	}
	return false
}

func (x *UserSpecificSettings) GetTray() bool {
	if x != nil {
		return x.Tray
	}
	return false
}

var File_settings_proto protoreflect.FileDescriptor

var file_settings_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x91, 0x01, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0xb2, 0x05, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x32, 0x0a, 0x0a, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54,
	0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x0a, 0x74, 0x65, 0x63, 0x68, 0x6e,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x12, 0x3f, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x68, 0x6e,
	0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x68, 0x6e, 0x65,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x77, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x66, 0x77, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x14, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x5f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x6c, 0x61, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x2b, 0x0a, 0x09,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x09,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x66,
	0x75, 0x73, 0x63, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x62,
	0x66, 0x75, 0x73, 0x63, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d,
	0x5f, 0x76, 0x70, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x74,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x75, 0x6d, 0x56, 0x70, 0x6e, 0x12, 0x3d, 0x0a, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x54, 0x0a, 0x14, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x72, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x74, 0x72, 0x61, 0x79, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f,
	0x72, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x6e, 0x6f, 0x72, 0x64, 0x76,
	0x70, 0x6e, 0x2d, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_settings_proto_rawDescOnce sync.Once
	file_settings_proto_rawDescData = file_settings_proto_rawDesc
)

func file_settings_proto_rawDescGZIP() []byte {
	file_settings_proto_rawDescOnce.Do(func() {
		file_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_settings_proto_rawDescData)
	})
	return file_settings_proto_rawDescData
}

var file_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_settings_proto_goTypes = []interface{}{
	(*SettingsResponse)(nil),     // 0: pb.SettingsResponse
	(*AutoconnectData)(nil),      // 1: pb.AutoconnectData
	(*Settings)(nil),             // 2: pb.Settings
	(*UserSpecificSettings)(nil), // 3: pb.UserSpecificSettings
	(config.ServerGroup)(0),      // 4: config.ServerGroup
	(config.Technology)(0),       // 5: config.Technology
	(config.Protocol)(0),         // 6: config.Protocol
	(*Allowlist)(nil),            // 7: pb.Allowlist
}
var file_settings_proto_depIdxs = []int32{
	2, // 0: pb.SettingsResponse.data:type_name -> pb.Settings
	4, // 1: pb.AutoconnectData.server_group:type_name -> config.ServerGroup
	5, // 2: pb.Settings.technology:type_name -> config.Technology
	1, // 3: pb.Settings.auto_connect_data:type_name -> pb.AutoconnectData
	6, // 4: pb.Settings.protocol:type_name -> config.Protocol
	7, // 5: pb.Settings.allowlist:type_name -> pb.Allowlist
	3, // 6: pb.Settings.user_settings:type_name -> pb.UserSpecificSettings
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_settings_proto_init() }
func file_settings_proto_init() {
	if File_settings_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoconnectData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSpecificSettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_settings_proto_goTypes,
		DependencyIndexes: file_settings_proto_depIdxs,
		MessageInfos:      file_settings_proto_msgTypes,
	}.Build()
	File_settings_proto = out.File
	file_settings_proto_rawDesc = nil
	file_settings_proto_goTypes = nil
	file_settings_proto_depIdxs = nil
}
