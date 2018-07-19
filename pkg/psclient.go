package pkg

import (
	"errors"

	emitter "github.com/emitter-io/go"
)

type (
	// PSClient is a wrapper of Emitter client
	PSClient struct {
		Client emitter.Emitter
	}
)

var (
	// ErrPSConnection is a connection error
	ErrPSConnection = errors.New("Failed to connect Pub/Sub service")
	// ErrPSPublish is a publish error
	ErrPSPublish = errors.New("Failed to publish message")
)

// NewPSClient returns an instance of PSClient
func NewPSClient() *PSClient {
	// Create the option with default vaules
	opt := emitter.NewClientOptions()
	opt.AddBroker(Conf.Emitter.Address)
	c := emitter.NewClient(opt)
	return &PSClient{Client: c}
}

// Connect offers sync function
func (p *PSClient) Connect() error {
	sToken := p.Client.Connect()
	if sToken.Wait() && sToken.Error() != nil {
		AppLog.Error(ErrPSConnection)
		return ErrPSConnection
	}
	return nil
}

// Disconnect offers sync function
func (p *PSClient) Disconnect(waitTime uint) {
	if p.Client.IsConnected() {
		p.Client.Disconnect(waitTime)
	}
}

// Publish offers sync function
func (p *PSClient) Publish(key string, channel string, payload interface{}) error {
	sToken := p.Client.Publish(key, channel, payload)
	if sToken.Wait() && sToken.Error() != nil {
		AppLog.Error(ErrPSPublish)
		return ErrPSPublish
	}
	return nil
}

// PublishWithTTL offers sync function
func (p *PSClient) PublishWithTTL(key string, channel string, payload interface{}, ttl int) error {
	sToken := p.Client.PublishWithTTL(key, channel, payload, ttl)
	if sToken.Wait() && sToken.Error() != nil {
		AppLog.Error(ErrPSPublish)
		return ErrPSPublish
	}
	return nil
}
