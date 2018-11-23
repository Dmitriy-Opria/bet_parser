package log

const (
	FieldCategory           = "category"
	FieldApplication        = "application"
	FieldService            = "service"
	FieldLocalAddress       = "localaddr"
	FieldLocalPort          = "localport"
	FieldDuration           = "duration"
	FieldRemoteAddress      = "remoteaddr"
	FieldRemotePort         = "remoteport"
	FieldRequestBodyLen     = "reqlen"
	FieldRequestMethod      = "method"
	FieldRequestURI         = "uri"
	FieldRequestProtocol    = "proto"
	FieldResponseStatusCode = "status"
	FieldAuthnUsername      = "authuser"
	FieldAuthnGroups        = "authgroups"
	FieldAuthnRoles         = "authnroles"
	FieldUserAgent          = "useragent"
	FieldFunc               = "func"
	FieldEntityType         = "entitytype"
	FieldTraceID            = "traceid"
	FieldSpanID             = "spanid"
	FieldParentSpanID       = "parentspanid"
	FieldRequestUUID        = "requestuuid"
)

// Deprecated fields.
const (
	FieldHost = "host"
)
