package signalr

import "time"

type party interface {
	onConnected(hc hubConnection)
	onDisconnected(hc hubConnection)

	invocationTarget(hc hubConnection) interface{}

	timeout() time.Duration
	setTimeout(timeout time.Duration)

	setHandshakeTimeout(timeout time.Duration)

	keepAliveInterval() time.Duration
	setKeepAliveInterval(interval time.Duration)

	chanReceiveTimeout() time.Duration
	setChanReceiveTimeout(interval time.Duration)

	streamBufferCapacity() uint
	setStreamBufferCapacity(capacity uint)

	allowReconnect() bool

	enableDetailedErrors() bool
	setEnableDetailedErrors(enable bool)

	loggers() (info StructuredLogger, dbg StructuredLogger)
	setLoggers(info StructuredLogger, dbg StructuredLogger)

	prefixLoggers() (info StructuredLogger, dbg StructuredLogger)

	maximumReceiveMessageSize() uint
	setMaximumReceiveMessageSize(size uint)
}

type partyBase struct {
	_timeout                   time.Duration
	_handshakeTimeout          time.Duration
	_keepAliveInterval         time.Duration
	_chanReceiveTimeout        time.Duration
	_streamBufferCapacity      uint
	_maximumReceiveMessageSize uint
	_enableDetailedErrors      bool
	info                       StructuredLogger
	dbg                        StructuredLogger
}

func (p *partyBase) timeout() time.Duration {
	return p._timeout
}

func (p *partyBase) setTimeout(timeout time.Duration) {
	p._timeout = timeout
}

func (p *partyBase) HandshakeTimeout() time.Duration {
	return p._handshakeTimeout
}

func (p *partyBase) setHandshakeTimeout(timeout time.Duration) {
	p._handshakeTimeout = timeout
}

func (p *partyBase) keepAliveInterval() time.Duration {
	return p._keepAliveInterval
}

func (p *partyBase) setKeepAliveInterval(interval time.Duration) {
	p._keepAliveInterval = interval
}

func (p *partyBase) chanReceiveTimeout() time.Duration {
	return p._chanReceiveTimeout
}

func (p *partyBase) setChanReceiveTimeout(interval time.Duration) {
	p._chanReceiveTimeout = interval
}

func (p *partyBase) streamBufferCapacity() uint {
	return p._streamBufferCapacity
}

func (p *partyBase) setStreamBufferCapacity(capacity uint) {
	p._streamBufferCapacity = capacity
}

func (p *partyBase) maximumReceiveMessageSize() uint {
	return p._maximumReceiveMessageSize
}

func (p *partyBase) setMaximumReceiveMessageSize(size uint) {
	p._maximumReceiveMessageSize = size
}

func (p *partyBase) enableDetailedErrors() bool {
	return p._enableDetailedErrors
}

func (p *partyBase) setEnableDetailedErrors(enable bool) {
	p._enableDetailedErrors = enable
}

func (p *partyBase) setLoggers(info StructuredLogger, dbg StructuredLogger) {
	p.info = info
	p.dbg = dbg
}

func (p *partyBase) loggers() (info StructuredLogger, debug StructuredLogger) {
	return p.info, p.dbg
}
