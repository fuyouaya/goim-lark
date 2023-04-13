// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.6
// source: pb_auth/auth.proto

package pb_auth

import (
	pb_enum "lark/pkg/proto/pb_enum"
	pb_user "lark/pkg/proto/pb_user"
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SignUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegPlatform pb_enum.PLATFORM_TYPE `protobuf:"varint,1,opt,name=reg_platform,json=regPlatform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"reg_platform,omitempty"` // 注册平台
	Nickname    string                `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`                                                      // 昵称
	Password    string                `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                                                      // 密码
	Firstname   string                `protobuf:"bytes,4,opt,name=firstname,proto3" json:"firstname,omitempty"`                                                    // firstname
	Lastname    string                `protobuf:"bytes,5,opt,name=lastname,proto3" json:"lastname,omitempty"`                                                      // lastname
	Gender      int32                 `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`                                                         // 性别
	BirthTs     int64                 `protobuf:"varint,7,opt,name=birth_ts,json=birthTs,proto3" json:"birth_ts,omitempty"`                                        // 生日
	Email       string                `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`                                                            // Email
	Mobile      string                `protobuf:"bytes,9,opt,name=mobile,proto3" json:"mobile,omitempty"`                                                          // 手机号
	AvatarKey   string                `protobuf:"bytes,10,opt,name=avatar_key,json=avatarKey,proto3" json:"avatar_key,omitempty"`                                  // 头像
	CityId      int64                 `protobuf:"varint,11,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`                                          // 城市ID
	Code        int64                 `protobuf:"varint,12,opt,name=code,proto3" json:"code,omitempty"`                                                            // 验证码
	Udid        string                `protobuf:"bytes,13,opt,name=udid,proto3" json:"udid,omitempty"`                                                             // udid
	ServerId    int64                 `protobuf:"varint,14,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`                                    // server id
}

func (x *SignUpReq) Reset() {
	*x = SignUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpReq) ProtoMessage() {}

func (x *SignUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpReq.ProtoReflect.Descriptor instead.
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpReq) GetRegPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.RegPlatform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

func (x *SignUpReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SignUpReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpReq) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *SignUpReq) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *SignUpReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *SignUpReq) GetBirthTs() int64 {
	if x != nil {
		return x.BirthTs
	}
	return 0
}

func (x *SignUpReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SignUpReq) GetAvatarKey() string {
	if x != nil {
		return x.AvatarKey
	}
	return ""
}

func (x *SignUpReq) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *SignUpReq) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SignUpReq) GetUdid() string {
	if x != nil {
		return x.Udid
	}
	return ""
}

func (x *SignUpReq) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type SignUpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg          string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	UserInfo     *pb_user.UserInfo `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	AccessToken  *Token            `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken *Token            `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *SignUpResp) Reset() {
	*x = SignUpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResp) ProtoMessage() {}

func (x *SignUpResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResp.ProtoReflect.Descriptor instead.
func (*SignUpResp) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SignUpResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SignUpResp) GetUserInfo() *pb_user.UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *SignUpResp) GetAccessToken() *Token {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *SignUpResp) GetRefreshToken() *Token {
	if x != nil {
		return x.RefreshToken
	}
	return nil
}

type SignInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountType      pb_enum.ACCOUNT_TYPE  `protobuf:"varint,1,opt,name=account_type,json=accountType,proto3,enum=pb_enum.ACCOUNT_TYPE" json:"account_type,omitempty"` // 账户类型 1:手机号 2:lark账户
	Platform         pb_enum.PLATFORM_TYPE `protobuf:"varint,2,opt,name=platform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"platform,omitempty"`                         // 平台
	Account          string                `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`                                                       // 手机号/lark账户
	Udid             string                `protobuf:"bytes,4,opt,name=udid,proto3" json:"udid,omitempty"`                                                             // 设备唯一编号
	VerificationCode string                `protobuf:"bytes,5,opt,name=verification_code,json=verificationCode,proto3" json:"verification_code,omitempty"`             // 验证码
	Password         string                `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`                                                     // 密码
}

func (x *SignInReq) Reset() {
	*x = SignInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInReq) ProtoMessage() {}

func (x *SignInReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInReq.ProtoReflect.Descriptor instead.
func (*SignInReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SignInReq) GetAccountType() pb_enum.ACCOUNT_TYPE {
	if x != nil {
		return x.AccountType
	}
	return pb_enum.ACCOUNT_TYPE(0)
}

func (x *SignInReq) GetPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.Platform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

func (x *SignInReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SignInReq) GetUdid() string {
	if x != nil {
		return x.Udid
	}
	return ""
}

func (x *SignInReq) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

func (x *SignInReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg          string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	UserInfo     *pb_user.UserInfo `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	AccessToken  *Token            `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken *Token            `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *SignInResp) Reset() {
	*x = SignInResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResp) ProtoMessage() {}

func (x *SignInResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResp.ProtoReflect.Descriptor instead.
func (*SignInResp) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SignInResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SignInResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SignInResp) GetUserInfo() *pb_user.UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *SignInResp) GetAccessToken() *Token {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *SignInResp) GetRefreshToken() *Token {
	if x != nil {
		return x.RefreshToken
	}
	return nil
}

type RefreshTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"` // refresh token
}

func (x *RefreshTokenReq) Reset() {
	*x = RefreshTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenReq) ProtoMessage() {}

func (x *RefreshTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenReq.ProtoReflect.Descriptor instead.
func (*RefreshTokenReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshTokenReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg         string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	AccessToken *Token `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *RefreshTokenResp) Reset() {
	*x = RefreshTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResp) ProtoMessage() {}

func (x *RefreshTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResp.ProtoReflect.Descriptor instead.
func (*RefreshTokenResp) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshTokenResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RefreshTokenResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RefreshTokenResp) GetAccessToken() *Token {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"token"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"` // 用户token
	// @inject_tag: json:"expire"
	Expire int64 `protobuf:"varint,2,opt,name=expire,proto3" json:"expire"` // token过期时间戳（秒）
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetExpire() int64 {
	if x != nil {
		return x.Expire
	}
	return 0
}

type SignOutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64                 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`                                      // uid
	Platform pb_enum.PLATFORM_TYPE `protobuf:"varint,2,opt,name=platform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"platform,omitempty"` // 平台
}

func (x *SignOutReq) Reset() {
	*x = SignOutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutReq) ProtoMessage() {}

func (x *SignOutReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutReq.ProtoReflect.Descriptor instead.
func (*SignOutReq) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{7}
}

func (x *SignOutReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SignOutReq) GetPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.Platform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

type SignOutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SignOutResp) Reset() {
	*x = SignOutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutResp) ProtoMessage() {}

func (x *SignOutResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutResp.ProtoReflect.Descriptor instead.
func (*SignOutResp) Descriptor() ([]byte, []int) {
	return file_pb_auth_auth_proto_rawDescGZIP(), []int{8}
}

func (x *SignOutResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SignOutResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_pb_auth_auth_proto protoreflect.FileDescriptor

var file_pb_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x12, 0x70,
	0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x52, 0x65, 0x71, 0x12, 0x39, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x54, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4b, 0x65, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x64, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x64, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0xca,
	0x01, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf0, 0x01, 0x0a, 0x09,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x38, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x64, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x64, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xca,
	0x01, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x36, 0x0a, 0x0f, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x6b, 0x0a, 0x10, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x31, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x35, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x52, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x4f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x33, 0x0a, 0x0b, 0x53,
	0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x32, 0xe7, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x31, 0x0a, 0x06, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x06,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x43, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x18, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f,
	0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_auth_auth_proto_rawDescOnce sync.Once
	file_pb_auth_auth_proto_rawDescData = file_pb_auth_auth_proto_rawDesc
)

func file_pb_auth_auth_proto_rawDescGZIP() []byte {
	file_pb_auth_auth_proto_rawDescOnce.Do(func() {
		file_pb_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_auth_auth_proto_rawDescData)
	})
	return file_pb_auth_auth_proto_rawDescData
}

var file_pb_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_auth_auth_proto_goTypes = []interface{}{
	(*SignUpReq)(nil),          // 0: pb_auth.SignUpReq
	(*SignUpResp)(nil),         // 1: pb_auth.SignUpResp
	(*SignInReq)(nil),          // 2: pb_auth.SignInReq
	(*SignInResp)(nil),         // 3: pb_auth.SignInResp
	(*RefreshTokenReq)(nil),    // 4: pb_auth.RefreshTokenReq
	(*RefreshTokenResp)(nil),   // 5: pb_auth.RefreshTokenResp
	(*Token)(nil),              // 6: pb_auth.Token
	(*SignOutReq)(nil),         // 7: pb_auth.SignOutReq
	(*SignOutResp)(nil),        // 8: pb_auth.SignOutResp
	(pb_enum.PLATFORM_TYPE)(0), // 9: pb_enum.PLATFORM_TYPE
	(*pb_user.UserInfo)(nil),   // 10: pb_user.UserInfo
	(pb_enum.ACCOUNT_TYPE)(0),  // 11: pb_enum.ACCOUNT_TYPE
}
var file_pb_auth_auth_proto_depIdxs = []int32{
	9,  // 0: pb_auth.SignUpReq.reg_platform:type_name -> pb_enum.PLATFORM_TYPE
	10, // 1: pb_auth.SignUpResp.user_info:type_name -> pb_user.UserInfo
	6,  // 2: pb_auth.SignUpResp.access_token:type_name -> pb_auth.Token
	6,  // 3: pb_auth.SignUpResp.refresh_token:type_name -> pb_auth.Token
	11, // 4: pb_auth.SignInReq.account_type:type_name -> pb_enum.ACCOUNT_TYPE
	9,  // 5: pb_auth.SignInReq.platform:type_name -> pb_enum.PLATFORM_TYPE
	10, // 6: pb_auth.SignInResp.user_info:type_name -> pb_user.UserInfo
	6,  // 7: pb_auth.SignInResp.access_token:type_name -> pb_auth.Token
	6,  // 8: pb_auth.SignInResp.refresh_token:type_name -> pb_auth.Token
	6,  // 9: pb_auth.RefreshTokenResp.access_token:type_name -> pb_auth.Token
	9,  // 10: pb_auth.SignOutReq.platform:type_name -> pb_enum.PLATFORM_TYPE
	0,  // 11: pb_auth.Auth.SignUp:input_type -> pb_auth.SignUpReq
	2,  // 12: pb_auth.Auth.SignIn:input_type -> pb_auth.SignInReq
	4,  // 13: pb_auth.Auth.RefreshToken:input_type -> pb_auth.RefreshTokenReq
	7,  // 14: pb_auth.Auth.SignOut:input_type -> pb_auth.SignOutReq
	1,  // 15: pb_auth.Auth.SignUp:output_type -> pb_auth.SignUpResp
	3,  // 16: pb_auth.Auth.SignIn:output_type -> pb_auth.SignInResp
	5,  // 17: pb_auth.Auth.RefreshToken:output_type -> pb_auth.RefreshTokenResp
	8,  // 18: pb_auth.Auth.SignOut:output_type -> pb_auth.SignOutResp
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pb_auth_auth_proto_init() }
func file_pb_auth_auth_proto_init() {
	if File_pb_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpReq); i {
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
		file_pb_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResp); i {
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
		file_pb_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInReq); i {
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
		file_pb_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResp); i {
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
		file_pb_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenReq); i {
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
		file_pb_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenResp); i {
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
		file_pb_auth_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_pb_auth_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutReq); i {
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
		file_pb_auth_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutResp); i {
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
			RawDescriptor: file_pb_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_auth_auth_proto_goTypes,
		DependencyIndexes: file_pb_auth_auth_proto_depIdxs,
		MessageInfos:      file_pb_auth_auth_proto_msgTypes,
	}.Build()
	File_pb_auth_auth_proto = out.File
	file_pb_auth_auth_proto_rawDesc = nil
	file_pb_auth_auth_proto_goTypes = nil
	file_pb_auth_auth_proto_depIdxs = nil
}