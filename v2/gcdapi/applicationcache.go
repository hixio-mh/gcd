// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains ApplicationCache functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Detailed application cache resource information.
type ApplicationCacheApplicationCacheResource struct {
	Url  string `json:"url"`  // Resource url.
	Size int    `json:"size"` // Resource size.
	Type string `json:"type"` // Resource type.
}

// Detailed application cache information.
type ApplicationCacheApplicationCache struct {
	ManifestURL  string                                      `json:"manifestURL"`  // Manifest URL.
	Size         float64                                     `json:"size"`         // Application cache size.
	CreationTime float64                                     `json:"creationTime"` // Application cache creation time.
	UpdateTime   float64                                     `json:"updateTime"`   // Application cache update time.
	Resources    []*ApplicationCacheApplicationCacheResource `json:"resources"`    // Application cache resources.
}

// Frame identifier - manifest URL pair.
type ApplicationCacheFrameWithManifest struct {
	FrameId     string `json:"frameId"`     // Frame identifier.
	ManifestURL string `json:"manifestURL"` // Manifest URL.
	Status      int    `json:"status"`      // Application cache status.
}

//
type ApplicationCacheApplicationCacheStatusUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		FrameId     string `json:"frameId"`     // Identifier of the frame containing document whose application cache updated status.
		ManifestURL string `json:"manifestURL"` // Manifest URL.
		Status      int    `json:"status"`      // Updated application cache status.
	} `json:"Params,omitempty"`
}

//
type ApplicationCacheNetworkStateUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		IsNowOnline bool `json:"isNowOnline"` //
	} `json:"Params,omitempty"`
}

type ApplicationCache struct {
	target gcdmessage.ChromeTargeter
}

func NewApplicationCache(target gcdmessage.ChromeTargeter) *ApplicationCache {
	c := &ApplicationCache{target: target}
	return c
}

// Enables application cache domain notifications.
func (c *ApplicationCache) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ApplicationCache.enable"})
}

type ApplicationCacheGetApplicationCacheForFrameParams struct {
	// Identifier of the frame containing document whose application cache is retrieved.
	FrameId string `json:"frameId"`
}

// GetApplicationCacheForFrameWithParams - Returns relevant application cache data for the document in given frame.
// Returns -  applicationCache - Relevant application cache data for the document in given frame.
func (c *ApplicationCache) GetApplicationCacheForFrameWithParams(ctx context.Context, v *ApplicationCacheGetApplicationCacheForFrameParams) (*ApplicationCacheApplicationCache, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ApplicationCache.getApplicationCacheForFrame", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			ApplicationCache *ApplicationCacheApplicationCache
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.ApplicationCache, nil
}

// GetApplicationCacheForFrame - Returns relevant application cache data for the document in given frame.
// frameId - Identifier of the frame containing document whose application cache is retrieved.
// Returns -  applicationCache - Relevant application cache data for the document in given frame.
func (c *ApplicationCache) GetApplicationCacheForFrame(ctx context.Context, frameId string) (*ApplicationCacheApplicationCache, error) {
	var v ApplicationCacheGetApplicationCacheForFrameParams
	v.FrameId = frameId
	return c.GetApplicationCacheForFrameWithParams(ctx, &v)
}

// GetFramesWithManifests - Returns array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
// Returns -  frameIds - Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
func (c *ApplicationCache) GetFramesWithManifests(ctx context.Context) ([]*ApplicationCacheFrameWithManifest, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ApplicationCache.getFramesWithManifests"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			FrameIds []*ApplicationCacheFrameWithManifest
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.FrameIds, nil
}

type ApplicationCacheGetManifestForFrameParams struct {
	// Identifier of the frame containing document whose manifest is retrieved.
	FrameId string `json:"frameId"`
}

// GetManifestForFrameWithParams - Returns manifest URL for document in the given frame.
// Returns -  manifestURL - Manifest URL for document in the given frame.
func (c *ApplicationCache) GetManifestForFrameWithParams(ctx context.Context, v *ApplicationCacheGetManifestForFrameParams) (string, error) {
	resp, err := c.target.SendCustomReturn(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "ApplicationCache.getManifestForFrame", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			ManifestURL string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.ManifestURL, nil
}

// GetManifestForFrame - Returns manifest URL for document in the given frame.
// frameId - Identifier of the frame containing document whose manifest is retrieved.
// Returns -  manifestURL - Manifest URL for document in the given frame.
func (c *ApplicationCache) GetManifestForFrame(ctx context.Context, frameId string) (string, error) {
	var v ApplicationCacheGetManifestForFrameParams
	v.FrameId = frameId
	return c.GetManifestForFrameWithParams(ctx, &v)
}
