// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package types

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/dragonfax/evernote-sdk-go/limits"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = limits.GoUnusedProtection__

const CLASSIFICATION_RECIPE_USER_NON_RECIPE = "000"
const CLASSIFICATION_RECIPE_USER_RECIPE = "001"
const CLASSIFICATION_RECIPE_SERVICE_RECIPE = "002"
const EDAM_NOTE_SOURCE_WEB_CLIP = "web.clip"
const EDAM_NOTE_SOURCE_WEB_CLIP_SIMPLIFIED = "Clearly"
const EDAM_NOTE_SOURCE_MAIL_CLIP = "mail.clip"
const EDAM_NOTE_SOURCE_MAIL_SMTP_GATEWAY = "mail.smtp"

func init() {
}
