/**
 * Deinterlace modes:
 *
 * - `OBS_DEINTERLACE_MODE_DISABLE`
 * - `OBS_DEINTERLACE_MODE_DISCARD`
 * - `OBS_DEINTERLACE_MODE_RETRO`
 * - `OBS_DEINTERLACE_MODE_BLEND`
 * - `OBS_DEINTERLACE_MODE_BLEND_2X`
 * - `OBS_DEINTERLACE_MODE_LINEAR`
 * - `OBS_DEINTERLACE_MODE_LINEAR_2X`
 * - `OBS_DEINTERLACE_MODE_YADIF`
 * - `OBS_DEINTERLACE_MODE_YADIF_2X`
 *
 */

package inputs

// Represents the request body for the SetInputDeinterlaceMode request.
type SetInputDeinterlaceModeParams struct {
	// Name of the input to take a DeinterlaceMode of
	InputName string `json:"inputName,omitempty"`

	InputDeinterlaceMode string `json:"inputDeinterlaceMode"`
}

// Returns the associated request.
func (o *SetInputDeinterlaceModeParams) GetRequestName() string {
	return "SetInputDeinterlaceMode"
}

// Represents the response body for the SetInputDeinterlaceMode request.
type SetInputDeinterlaceModeResponse struct {
	_response
}

/*
Gets a Base64-encoded DeinterlaceMode of a input.

The `imageWidth` and `imageHeight` parameters are treated as "scale to inner", meaning the smallest ratio will be used and the aspect ratio of the original resolution is kept.
If `imageWidth` and `imageHeight` are not specified, the compressed image will use the full resolution of the input.

**Compatible with inputs and scenes.**
*/
func (c *Client) SetInputDeinterlaceMode(params *SetInputDeinterlaceModeParams) (*SetInputDeinterlaceModeResponse, error) {
	data := &SetInputDeinterlaceModeResponse{}
	return data, c.client.SendRequest(params, data)
}
