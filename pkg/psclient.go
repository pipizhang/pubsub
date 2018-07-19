package pkg

import (
	"errors"

	emitter "github.com/emitter-io/go"
)

type (
	PSClient struct {
		Client emitter.Emitter
	}
)

var (
	ErrPSConnection = errors.New("Failed to connect Pub/Sub service")
	ErrPSPublish    = errors.New("Failed to publish message")
)

// NewPSClient returns an instance of PSClient
func NewPSClient() *PSClient {
	// Create the option with default vaules
	opt := emitter.NewClientOptions()
	opt.AddBroker(Conf.Emitter.Address)
	c := emitter.NewClient(opt)
	return &PSClient{Client: c}
}

// Sync Connect
func (p *PSClient) Connect() error {
	sToken := p.Client.Connect()
	if sToken.Wait() && sToken.Error() != nil {
		AppLog.Error(ErrPSConnection)
		return ErrPSConnection
	}
	return nil
}

// Disconnect
func (p *PSClient) Disconnect(waitTime uint) {
	if p.Client.IsConnected() {
		p.Client.Disconnect(waitTime)
	}
}

// Sync Publish function
func (p *PSClient) Publish(key string, channel string, payload interface{}) error {
	sToken := p.Client.Publish(key, channel, payload)
	if sToken.Wait() && sToken.Error() != nil {
		AppLog.Error(ErrPSPublish)
		return ErrPSPublish
	}
	return nil
}

// Sync PublishWithTTL function
func (p *PSClient) PublishWithTTL(key string, channel string, payload interface{}, ttl int) error {
	sToken := p.Client.PublishWithTTL(key, channel, payload, ttl)
	if sToken.Wait() && sToken.Error() != nil {
		AppLog.Error(ErrPSPublish)
		return ErrPSPublish
	}
	return nil
}
