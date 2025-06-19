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

// Represents the request body for the GetInputDeinterlaceFieldOrder request.
type GetInputDeinterlaceFieldOrderParams struct {
	// Name of the input to take a DeinterlaceFieldOrder of
	InputName string `json:"inputName,omitempty"`
}

// Returns the associated request.
func (o *GetInputDeinterlaceFieldOrderParams) GetRequestName() string {
	return "GetInputDeinterlaceFieldOrder"
}

// Represents the response body for the GetInputDeinterlaceFieldOrder request.
type GetInputDeinterlaceFieldOrderResponse struct {
	_response

	// Base64-encoded DeinterlaceFieldOrder
	InputDeinterlaceFieldOrder string `json:"inputDeinterlaceFieldOrder,omitempty"`
}

/*
Gets a Base64-encoded DeinterlaceFieldOrder of a input.

The `imageWidth` and `imageHeight` parameters are treated as "scale to inner", meaning the smallest ratio will be used and the aspect ratio of the original resolution is kept.
If `imageWidth` and `imageHeight` are not specified, the compressed image will use the full resolution of the input.

**Compatible with inputs and scenes.**
*/
func (c *Client) GetInputDeinterlaceFieldOrder(params *GetInputDeinterlaceFieldOrderParams) (*GetInputDeinterlaceFieldOrderResponse, error) {
	data := &GetInputDeinterlaceFieldOrderResponse{}
	return data, c.client.SendRequest(params, data)
}
