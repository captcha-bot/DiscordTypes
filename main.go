package discordtypes

import (
	"encoding/json"
	"io"
	"time"
)

type ChannelType int
type VerificationLevel int
type MfaLevel int
type PremiumTier int

type Identify struct {
	Op   int `json:"op"`
	Data struct {
		Token      string `json:"token"`
		Properties struct {
			OS      string `json:"os"`
			Browser string `json:"browser"`
			Device  string `json:"device"`
		} `json:"properties"`
		Version        int      `json:"v"`
		LargeThreshold int      `json:"large_threshold"`
		Compress       bool     `json:"compress"`
		Shard          [2]int64 `json:"shard"`
		Intents        int64    `json:"intents"`
	} `json:"d"`
}

type AllowedMentionType string

// The types of mentions used in MessageAllowedMentions.
const (
	AllowedMentionTypeRoles    AllowedMentionType = "roles"
	AllowedMentionTypeUsers    AllowedMentionType = "users"
	AllowedMentionTypeEveryone AllowedMentionType = "everyone"
)

type MessageSend struct {
	Content         string                  `json:"content,omitempty"`
	Embed           *MessageEmbed           `json:"embed,omitempty"`
	TTS             bool                    `json:"tts"`
	Files           []*File                 `json:"-"`
	AllowedMentions *MessageAllowedMentions `json:"allowed_mentions,omitempty"`

	// TODO: Remove this when compatibility is not required.
	File *File `json:"-"`
}

type MessageAllowedMentions struct {
	// The mention types that are allowed to be parsed in this message.
	// Please note that this is purposely **not** marked as omitempty,
	// so if a zero-value MessageAllowedMentions object is provided no
	// mentions will be allowed.
	Parse []AllowedMentionType `json:"parse"`

	// A list of role IDs to allow. This cannot be used when specifying
	// AllowedMentionTypeRoles in the Parse slice.
	Roles []string `json:"roles,omitempty"`

	// A list of user IDs to allow. This cannot be used when specifying
	// AllowedMentionTypeUsers in the Parse slice.
	Users []string `json:"users,omitempty"`
}

type GatewayIdentify struct {
	Type      string `json:"t"`
	Operation int16  `json:"op"`
	RawData   struct {
		Token      string `json:"token"`
		Shard      int64  `json:"shard"`
		ShardCount int64  `json:"shardCount"`
	} `json:"d"`
}

type heartbeatOp struct {
	Op   int   `json:"op"`
	Data int64 `json:"d"`
}

type PingEvent struct {
	Operation int           `json:"op"`
	Data      time.Duration `json:"d"`
}
type Event struct {
	Operation int             `json:"op"`
	Sequence  int64           `json:"s"`
	Type      string          `json:"t"`
	RawData   json.RawMessage `json:"d"`
}

type GatewayStatusUpdate struct {
	Operation int    `json:"op"`
	Type      string `json:"t"`
	RawData   string `json:"d"`
}

type GatewayEvent struct {
	Operation int             `json:"op"`
	Type      string          `json:"t"`
	RawData   json.RawMessage `json:"d"`
}

type Handshake struct {
	Operation int `json:"op"`
	RawData   struct {
		HeartbeatInterval int      `json:"heartbeat_interval"`
		Trace             []string `json:"_trace"`
	} `json:"d"`
}

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
)

type PermissionOverwrite struct {
	ID    string `json:"id"`
	Type  int    `json:"type"`
	Deny  string `json:"deny"`
	Allow string `json:"allow"`
}

type Guild struct {
	ID                       string            `json:"id"`
	Name                     string            `json:"name"`
	Icon                     string            `json:"icon"`
	Region                   string            `json:"region"`
	AfkChannelID             string            `json:"afk_channel_id"`
	EmbedChannelID           string            `json:"embed_channel_id"`
	OwnerID                  string            `json:"owner_id"`
	JoinedAt                 string            `json:"joined_at"`
	Splash                   string            `json:"splash"`
	AfkTimeout               int               `json:"afk_timeout"`
	MemberCount              int               `json:"member_count"`
	VerificationLevel        VerificationLevel `json:"verification_level"`
	Large                    bool              `json:"large"`
	Roles                    []*Role           `json:"roles"`
	Members                  []*Member         `json:"members"`
	Channels                 []*Channel        `json:"channels"`
	Unavailable              bool              `json:"unavailable"`
	Features                 []string          `json:"features"`
	WidgetEnabled            bool              `json:"widget_enabled"`
	WidgetChannelID          string            `json:"widget_channel_id"`
	SystemChannelID          string            `json:"system_channel_id"`
	VanityURLCode            string            `json:"vanity_url_code"`
	Description              string            `json:"description"`
	Banner                   string            `json:"banner"`
	PremiumTier              PremiumTier       `json:"premium_tier"`
	PremiumSubscriptionCount int               `json:"premium_subscription_count"`
}

type Channel struct {
	ID                   string                 `json:"id"`
	GuildID              string                 `json:"guild_id"`
	Name                 string                 `json:"name"`
	Topic                string                 `json:"topic"`
	Type                 ChannelType            `json:"type"`
	LastMessageID        string                 `json:"last_message_id"`
	LastPinTimestamp     string                 `json:"last_pin_timestamp"`
	NSFW                 bool                   `json:"nsfw"`
	Icon                 string                 `json:"icon"`
	Position             int                    `json:"position"`
	Bitrate              int                    `json:"bitrate"`
	Recipients           []*User                `json:"recipients"`
	Messages             []*Message             `json:"-"`
	PermissionOverwrites []*PermissionOverwrite `json:"permission_overwrites"`
	UserLimit            int                    `json:"user_limit"`
	ParentID             string                 `json:"parent_id"`
	RateLimitPerUser     int                    `json:"rate_limit_per_user"`
}

type Member struct {
	GuildID      string   `json:"guild_id"`
	JoinedAt     string   `json:"joined_at"`
	Nick         string   `json:"nick"`
	Deaf         bool     `json:"deaf"`
	Mute         bool     `json:"mute"`
	User         *User    `json:"user"`
	Roles        []string `json:"roles"`
	PremiumSince string   `json:"premium_since"`
}

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Locale        string `json:"locale"`
	Discriminator string `json:"discriminator"`
	Token         string `json:"token"`
	Verified      bool   `json:"verified"`
	MFAEnabled    bool   `json:"mfa_enabled"`
	Bot           bool   `json:"bot"`
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Managed     bool   `json:"managed"`
	Mentionable bool   `json:"mentionable"`
	Hoist       bool   `json:"hoist"`
	Color       int    `json:"color"`
	Position    int    `json:"position"`
	Permissions string `json:"permissions"`
}

type MessageType int

// Block contains the valid known MessageType values
const (
	MessageTypeDefault MessageType = iota
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypeChannelPinnedMessage
	MessageTypeGuildMemberJoin
	MessageTypeUserPremiumGuildSubscription
	MessageTypeUserPremiumGuildSubscriptionTierOne
	MessageTypeUserPremiumGuildSubscriptionTierTwo
	MessageTypeUserPremiumGuildSubscriptionTierThree
	MessageTypeChannelFollowAdd
)

// A Message stores all data related to a specific Discord message.
type Message struct {
	ID              string          `json:"id"`
	ChannelID       string          `json:"channel_id"`
	GuildID         string          `json:"guild_id,omitempty"`
	Content         string          `json:"content"`
	Timestamp       string          `json:"timestamp"`
	EditedTimestamp string          `json:"edited_timestamp"`
	MentionRoles    []string        `json:"mention_roles"`
	Author          *User           `json:"author"`
	Embeds          []*MessageEmbed `json:"embeds"`
	Mentions        []*User         `json:"mentions"`
	Pinned          bool            `json:"pinned"`
	Type            MessageType     `json:"type"`
	WebhookID       string          `json:"webhook_id"`
	Member          *Member         `json:"member"`
	MentionChannels []*Channel      `json:"mention_channels"`
	Flags           int             `json:"flags"`
}

// File stores info about files you e.g. send in messages.
type File struct {
	Name        string
	ContentType string
	Reader      io.Reader
}

type MessageAttachment struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Filename string `json:"filename"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Size     int    `json:"size"`
}

// MessageEmbedFooter is a part of a MessageEmbed struct.
type MessageEmbedFooter struct {
	Text         string `json:"text,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// MessageEmbedImage is a part of a MessageEmbed struct.
type MessageEmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// MessageEmbedThumbnail is a part of a MessageEmbed struct.
type MessageEmbedThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// MessageEmbedVideo is a part of a MessageEmbed struct.
type MessageEmbedVideo struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// MessageEmbedProvider is a part of a MessageEmbed struct.
type MessageEmbedProvider struct {
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

// MessageEmbedAuthor is a part of a MessageEmbed struct.
type MessageEmbedAuthor struct {
	URL          string `json:"url,omitempty"`
	Name         string `json:"name,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// MessageEmbedField is a part of a MessageEmbed struct.
type MessageEmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type WebhookParams struct {
	Content         string                  `json:"content,omitempty"`
	Username        string                  `json:"username,omitempty"`
	AvatarURL       string                  `json:"avatar_url,omitempty"`
	TTS             bool                    `json:"tts,omitempty"`
	File            string                  `json:"file,omitempty"`
	Embeds          []*MessageEmbed         `json:"embeds,omitempty"`
	AllowedMentions *MessageAllowedMentions `json:"allowed_mentions,omitempty"`
}

// An MessageEmbed stores data for message embeds.
type MessageEmbed struct {
	URL         string                 `json:"url,omitempty"`
	Type        string                 `json:"type,omitempty"`
	Title       string                 `json:"title,omitempty"`
	Description string                 `json:"description,omitempty"`
	Timestamp   string                 `json:"timestamp,omitempty"`
	Color       int                    `json:"color,omitempty"`
	Footer      *MessageEmbedFooter    `json:"footer,omitempty"`
	Image       *MessageEmbedImage     `json:"image,omitempty"`
	Thumbnail   *MessageEmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *MessageEmbedVideo     `json:"video,omitempty"`
	Provider    *MessageEmbedProvider  `json:"provider,omitempty"`
	Author      *MessageEmbedAuthor    `json:"author,omitempty"`
	Fields      []*MessageEmbedField   `json:"fields,omitempty"`
}

// A Ready stores all data for the websocket READY event.
type Ready struct {
	Version         int        `json:"v"`
	SessionID       string     `json:"session_id"`
	User            *User      `json:"user"`
	PrivateChannels []*Channel `json:"private_channels"`
	Guilds          []*Guild   `json:"guilds"`
}

// ChannelCreate is the data for a ChannelCreate event.
type ChannelCreate struct {
	*Channel
}

// ChannelUpdate is the data for a ChannelUpdate event.
type ChannelUpdate struct {
	*Channel
}

// ChannelDelete is the data for a ChannelDelete event.
type ChannelDelete struct {
	*Channel
}

// GuildCreate is the data for a GuildCreate event.
type GuildCreate struct {
	*Guild
}

// GuildUpdate is the data for a GuildUpdate event.
type GuildUpdate struct {
	*Guild
}

// GuildDelete is the data for a GuildDelete event.
type GuildDelete struct {
	*Guild
}

// GuildMemberAdd is the data for a GuildMemberAdd event.
type GuildMemberAdd struct {
	*Member
}

// GuildMemberUpdate is the data for a GuildMemberUpdate event.
type GuildMemberUpdate struct {
	*Member
}

// GuildMemberRemove is the data for a GuildMemberRemove event.
type GuildMemberRemove struct {
	*Member
}

// GuildRoleCreate is the data for a GuildRoleCreate event.
type GuildRoleCreate struct {
	*GuildRole
}

// GuildRoleUpdate is the data for a GuildRoleUpdate event.
type GuildRoleUpdate struct {
	*GuildRole
}

// A GuildRoleDelete is the data for a GuildRoleDelete event.
type GuildRoleDelete struct {
	RoleID  string `json:"role_id"`
	GuildID string `json:"guild_id"`
}

// MessageCreate is the data for a MessageCreate event.
type MessageCreate struct {
	*Message
}

type Resume struct {
	Op   int `json:"op"`
	Data struct {
		Token  string `json:"token"`
		SessID string `json:"session_id"`
		Seq    int64  `json:"seq"`
	} `json:"d"`
}

// Resumed is the data for a Resumed event.
type Resumed struct {
	Trace []string `json:"_trace"`
}

// A GuildRole stores data for guild roles.
type GuildRole struct {
	Role    *Role  `json:"role"`
	GuildID string `json:"guild_id"`
}
