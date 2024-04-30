package network

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/chromedp/cdproto/cdp"
)

// EventDataReceived fired when data chunk was received over the network.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-dataReceived
type EventDataReceived struct {
	RequestID         RequestID          `json:"requestId"`         // Request identifier.
	Timestamp         *cdp.MonotonicTime `json:"timestamp"`         // Timestamp.
	DataLength        int64              `json:"dataLength"`        // Data chunk length.
	EncodedDataLength int64              `json:"encodedDataLength"` // Actual bytes received (might be less than dataLength for compressed encodings).
	Data              string             `json:"data,omitempty"`    // Data that was received.
}

// EventEventSourceMessageReceived fired when EventSource message is
// received.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-eventSourceMessageReceived
type EventEventSourceMessageReceived struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	EventName string             `json:"eventName"` // Message type.
	EventID   string             `json:"eventId"`   // Message identifier.
	Data      string             `json:"data"`      // Message content.
}

// EventLoadingFailed fired when HTTP request has failed to load.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-loadingFailed
type EventLoadingFailed struct {
	RequestID       RequestID          `json:"requestId"`                 // Request identifier.
	Timestamp       *cdp.MonotonicTime `json:"timestamp"`                 // Timestamp.
	Type            ResourceType       `json:"type"`                      // Resource type.
	ErrorText       string             `json:"errorText"`                 // User friendly error message.
	Canceled        bool               `json:"canceled,omitempty"`        // True if loading was canceled.
	BlockedReason   BlockedReason      `json:"blockedReason,omitempty"`   // The reason why loading was blocked, if any.
	CorsErrorStatus *CorsErrorStatus   `json:"corsErrorStatus,omitempty"` // The reason why loading was blocked by CORS, if any.
}

// EventLoadingFinished fired when HTTP request has finished loading.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-loadingFinished
type EventLoadingFinished struct {
	RequestID         RequestID          `json:"requestId"`         // Request identifier.
	Timestamp         *cdp.MonotonicTime `json:"timestamp"`         // Timestamp.
	EncodedDataLength float64            `json:"encodedDataLength"` // Total number of bytes received for this request.
}

// EventRequestServedFromCache fired if request ended up loading from cache.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-requestServedFromCache
type EventRequestServedFromCache struct {
	RequestID RequestID `json:"requestId"` // Request identifier.
}

// EventRequestWillBeSent fired when page is about to send HTTP request.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-requestWillBeSent
type EventRequestWillBeSent struct {
	RequestID            RequestID           `json:"requestId"`                  // Request identifier.
	LoaderID             cdp.LoaderID        `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched from worker.
	DocumentURL          string              `json:"documentURL"`                // URL of the document this request is loaded for.
	Request              *Request            `json:"request"`                    // Request data.
	Timestamp            *cdp.MonotonicTime  `json:"timestamp"`                  // Timestamp.
	WallTime             *cdp.TimeSinceEpoch `json:"wallTime"`                   // Timestamp.
	Initiator            *Initiator          `json:"initiator"`                  // Request initiator.
	RedirectHasExtraInfo bool                `json:"redirectHasExtraInfo"`       // In the case that redirectResponse is populated, this flag indicates whether requestWillBeSentExtraInfo and responseReceivedExtraInfo events will be or were emitted for the request which was just redirected.
	RedirectResponse     *Response           `json:"redirectResponse,omitempty"` // Redirect response data.
	Type                 ResourceType        `json:"type,omitempty"`             // Type of this resource.
	FrameID              cdp.FrameID         `json:"frameId,omitempty"`          // Frame identifier.
	HasUserGesture       bool                `json:"hasUserGesture,omitempty"`   // Whether the request is initiated by a user gesture. Defaults to false.
}

// EventResourceChangedPriority fired when resource loading priority is
// changed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-resourceChangedPriority
type EventResourceChangedPriority struct {
	RequestID   RequestID          `json:"requestId"`   // Request identifier.
	NewPriority ResourcePriority   `json:"newPriority"` // New priority
	Timestamp   *cdp.MonotonicTime `json:"timestamp"`   // Timestamp.
}

// EventSignedExchangeReceived fired when a signed exchange was received over
// the network.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-signedExchangeReceived
type EventSignedExchangeReceived struct {
	RequestID RequestID           `json:"requestId"` // Request identifier.
	Info      *SignedExchangeInfo `json:"info"`      // Information about the signed exchange response.
}

// EventResponseReceived fired when HTTP response is available.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-responseReceived
type EventResponseReceived struct {
	RequestID    RequestID          `json:"requestId"`         // Request identifier.
	LoaderID     cdp.LoaderID       `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched from worker.
	Timestamp    *cdp.MonotonicTime `json:"timestamp"`         // Timestamp.
	Type         ResourceType       `json:"type"`              // Resource type.
	Response     *Response          `json:"response"`          // Response data.
	HasExtraInfo bool               `json:"hasExtraInfo"`      // Indicates whether requestWillBeSentExtraInfo and responseReceivedExtraInfo events will be or were emitted for this request.
	FrameID      cdp.FrameID        `json:"frameId,omitempty"` // Frame identifier.
}

// EventWebSocketClosed fired when WebSocket is closed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketClosed
type EventWebSocketClosed struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
}

// EventWebSocketCreated fired upon WebSocket creation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketCreated
type EventWebSocketCreated struct {
	RequestID RequestID  `json:"requestId"`           // Request identifier.
	URL       string     `json:"url"`                 // WebSocket request URL.
	Initiator *Initiator `json:"initiator,omitempty"` // Request initiator.
}

// EventWebSocketFrameError fired when WebSocket message error occurs.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketFrameError
type EventWebSocketFrameError struct {
	RequestID    RequestID          `json:"requestId"`    // Request identifier.
	Timestamp    *cdp.MonotonicTime `json:"timestamp"`    // Timestamp.
	ErrorMessage string             `json:"errorMessage"` // WebSocket error message.
}

// EventWebSocketFrameReceived fired when WebSocket message is received.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketFrameReceived
type EventWebSocketFrameReceived struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	Response  *WebSocketFrame    `json:"response"`  // WebSocket response data.
}

// EventWebSocketFrameSent fired when WebSocket message is sent.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketFrameSent
type EventWebSocketFrameSent struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	Response  *WebSocketFrame    `json:"response"`  // WebSocket response data.
}

// EventWebSocketHandshakeResponseReceived fired when WebSocket handshake
// response becomes available.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketHandshakeResponseReceived
type EventWebSocketHandshakeResponseReceived struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	Response  *WebSocketResponse `json:"response"`  // WebSocket response data.
}

// EventWebSocketWillSendHandshakeRequest fired when WebSocket is about to
// initiate handshake.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webSocketWillSendHandshakeRequest
type EventWebSocketWillSendHandshakeRequest struct {
	RequestID RequestID           `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime  `json:"timestamp"` // Timestamp.
	WallTime  *cdp.TimeSinceEpoch `json:"wallTime"`  // UTC Timestamp.
	Request   *WebSocketRequest   `json:"request"`   // WebSocket request data.
}

// EventWebTransportCreated fired upon WebTransport creation.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webTransportCreated
type EventWebTransportCreated struct {
	TransportID RequestID          `json:"transportId"`         // WebTransport identifier.
	URL         string             `json:"url"`                 // WebTransport request URL.
	Timestamp   *cdp.MonotonicTime `json:"timestamp"`           // Timestamp.
	Initiator   *Initiator         `json:"initiator,omitempty"` // Request initiator.
}

// EventWebTransportConnectionEstablished fired when WebTransport handshake
// is finished.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webTransportConnectionEstablished
type EventWebTransportConnectionEstablished struct {
	TransportID RequestID          `json:"transportId"` // WebTransport identifier.
	Timestamp   *cdp.MonotonicTime `json:"timestamp"`   // Timestamp.
}

// EventWebTransportClosed fired when WebTransport is disposed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-webTransportClosed
type EventWebTransportClosed struct {
	TransportID RequestID          `json:"transportId"` // WebTransport identifier.
	Timestamp   *cdp.MonotonicTime `json:"timestamp"`   // Timestamp.
}

// EventRequestWillBeSentExtraInfo fired when additional information about a
// requestWillBeSent event is available from the network stack. Not every
// requestWillBeSent event will have an additional requestWillBeSentExtraInfo
// fired for it, and there is no guarantee whether requestWillBeSent or
// requestWillBeSentExtraInfo will be fired first for the same request.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-requestWillBeSentExtraInfo
type EventRequestWillBeSentExtraInfo struct {
	RequestID                     RequestID            `json:"requestId"`                               // Request identifier. Used to match this information to an existing requestWillBeSent event.
	AssociatedCookies             []*AssociatedCookie  `json:"associatedCookies"`                       // A list of cookies potentially associated to the requested URL. This includes both cookies sent with the request and the ones not sent; the latter are distinguished by having blockedReasons field set.
	Headers                       Headers              `json:"headers"`                                 // Raw request headers as they will be sent over the wire.
	ConnectTiming                 *ConnectTiming       `json:"connectTiming"`                           // Connection timing information for the request.
	ClientSecurityState           *ClientSecurityState `json:"clientSecurityState,omitempty"`           // The client security state set for the request.
	SiteHasCookieInOtherPartition bool                 `json:"siteHasCookieInOtherPartition,omitempty"` // Whether the site has partitioned cookies stored in a partition different than the current one.
}

// EventResponseReceivedExtraInfo fired when additional information about a
// responseReceived event is available from the network stack. Not every
// responseReceived event will have an additional responseReceivedExtraInfo for
// it, and responseReceivedExtraInfo may be fired before or after
// responseReceived.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-responseReceivedExtraInfo
type EventResponseReceivedExtraInfo struct {
	RequestID                RequestID                      `json:"requestId"`                          // Request identifier. Used to match this information to another responseReceived event.
	BlockedCookies           []*BlockedSetCookieWithReason  `json:"blockedCookies"`                     // A list of cookies which were not stored from the response along with the corresponding reasons for blocking. The cookies here may not be valid due to syntax errors, which are represented by the invalid cookie line string instead of a proper cookie.
	Headers                  Headers                        `json:"headers"`                            // Raw response headers as they were received over the wire.
	ResourceIPAddressSpace   IPAddressSpace                 `json:"resourceIPAddressSpace"`             // The IP address space of the resource. The address space can only be determined once the transport established the connection, so we can't send it in requestWillBeSentExtraInfo.
	StatusCode               int64                          `json:"statusCode"`                         // The status code of the response. This is useful in cases the request failed and no responseReceived event is triggered, which is the case for, e.g., CORS errors. This is also the correct status code for cached requests, where the status in responseReceived is a 200 and this will be 304.
	HeadersText              string                         `json:"headersText,omitempty"`              // Raw response header text as it was received over the wire. The raw text may not always be available, such as in the case of HTTP/2 or QUIC.
	CookiePartitionKey       string                         `json:"cookiePartitionKey,omitempty"`       // The cookie partition key that will be used to store partitioned cookies set in this response. Only sent when partitioned cookies are enabled.
	CookiePartitionKeyOpaque bool                           `json:"cookiePartitionKeyOpaque,omitempty"` // True if partitioned cookies are enabled, but the partition key is not serializeable to string.
	ExemptedCookies          []*ExemptedSetCookieWithReason `json:"exemptedCookies,omitempty"`          // A list of cookies which should have been blocked by 3PCD but are exempted and stored from the response with the corresponding reason.
}

// EventTrustTokenOperationDone fired exactly once for each Trust Token
// operation. Depending on the type of the operation and whether the operation
// succeeded or failed, the event is fired before the corresponding request was
// sent or after the response was received.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-trustTokenOperationDone
type EventTrustTokenOperationDone struct {
	Status           TrustTokenOperationDoneStatus `json:"status"` // Detailed success or error status of the operation. 'AlreadyExists' also signifies a successful operation, as the result of the operation already exists und thus, the operation was abort preemptively (e.g. a cache hit).
	Type             TrustTokenOperationType       `json:"type"`
	RequestID        RequestID                     `json:"requestId"`
	TopLevelOrigin   string                        `json:"topLevelOrigin,omitempty"`   // Top level origin. The context in which the operation was attempted.
	IssuerOrigin     string                        `json:"issuerOrigin,omitempty"`     // Origin of the issuer in case of a "Issuance" or "Redemption" operation.
	IssuedTokenCount int64                         `json:"issuedTokenCount,omitempty"` // The number of obtained Trust Tokens on a successful "Issuance" operation.
}

// EventSubresourceWebBundleMetadataReceived fired once when parsing the .wbn
// file has succeeded. The event contains the information about the web bundle
// contents.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-subresourceWebBundleMetadataReceived
type EventSubresourceWebBundleMetadataReceived struct {
	RequestID RequestID `json:"requestId"` // Request identifier. Used to match this information to another event.
	Urls      []string  `json:"urls"`      // A list of URLs of resources in the subresource Web Bundle.
}

// EventSubresourceWebBundleMetadataError fired once when parsing the .wbn
// file has failed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-subresourceWebBundleMetadataError
type EventSubresourceWebBundleMetadataError struct {
	RequestID    RequestID `json:"requestId"`    // Request identifier. Used to match this information to another event.
	ErrorMessage string    `json:"errorMessage"` // Error message
}

// EventSubresourceWebBundleInnerResponseParsed fired when handling requests
// for resources within a .wbn file. Note: this will only be fired for resources
// that are requested by the webpage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-subresourceWebBundleInnerResponseParsed
type EventSubresourceWebBundleInnerResponseParsed struct {
	InnerRequestID  RequestID `json:"innerRequestId"`            // Request identifier of the subresource request
	InnerRequestURL string    `json:"innerRequestURL"`           // URL of the subresource resource.
	BundleRequestID RequestID `json:"bundleRequestId,omitempty"` // Bundle request identifier. Used to match this information to another event. This made be absent in case when the instrumentation was enabled only after webbundle was parsed.
}

// EventSubresourceWebBundleInnerResponseError fired when request for
// resources within a .wbn file failed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-subresourceWebBundleInnerResponseError
type EventSubresourceWebBundleInnerResponseError struct {
	InnerRequestID  RequestID `json:"innerRequestId"`            // Request identifier of the subresource request
	InnerRequestURL string    `json:"innerRequestURL"`           // URL of the subresource resource.
	ErrorMessage    string    `json:"errorMessage"`              // Error message
	BundleRequestID RequestID `json:"bundleRequestId,omitempty"` // Bundle request identifier. Used to match this information to another event. This made be absent in case when the instrumentation was enabled only after webbundle was parsed.
}

// EventReportingAPIReportAdded is sent whenever a new report is added. And
// after 'enableReportingApi' for all existing reports.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-reportingApiReportAdded
type EventReportingAPIReportAdded struct {
	Report *ReportingAPIReport `json:"report"`
}

// EventReportingAPIReportUpdated [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-reportingApiReportUpdated
type EventReportingAPIReportUpdated struct {
	Report *ReportingAPIReport `json:"report"`
}

// EventReportingAPIEndpointsChangedForOrigin [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#event-reportingApiEndpointsChangedForOrigin
type EventReportingAPIEndpointsChangedForOrigin struct {
	Origin    string                  `json:"origin"` // Origin of the document(s) which configured the endpoints.
	Endpoints []*ReportingAPIEndpoint `json:"endpoints"`
}
