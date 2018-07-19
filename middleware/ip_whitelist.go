package middleware

import (
	"net/http"
	"sync"

	"github.com/jpillora/ipfilter"
	"github.com/labstack/echo"
)

type (
	// IPWhitelistConfig defines the config for IPWhitelist middleware.
	IPWhitelistConfig struct {
		Enable   bool
		List     []string
		IPDBFile string
	}

	gIPFilter struct {
		Client        *ipfilter.IPFilter
		IsInitialized bool
		rwmutex       sync.RWMutex
	}
)

var (
	DefaultIPWhitelistConfig = IPWhitelistConfig{
		Enable: true,
	}
	GIPFilter = &gIPFilter{
		IsInitialized: false,
	}
)

func IPWhitelist(ipList []string) echo.MiddlewareFunc {
	c := DefaultIPWhitelistConfig
	c.List = ipList
	return IPWhitelistWithConfig(c)
}

func IPWhitelistWithConfig(config IPWhitelistConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Enable {

				if !GIPFilter.IsInitialized {
					GIPFilter.rwmutex.Lock()
					GIPFilter.Client, err = ipfilter.New(ipfilter.Options{
						AllowedIPs:     config.List,
						BlockByDefault: true,
						IPDBPath:       config.IPDBFile,
					})
					GIPFilter.IsInitialized = true
					GIPFilter.rwmutex.Unlock()
					if err != nil {
						return echo.NewHTTPError(http.StatusInternalServerError)
					}
				}

				ip := c.RealIP()
				if !GIPFilter.Client.Allowed(ip) {
					return echo.ErrForbidden
				}

			}
			return next(c)
		}
	}
}
