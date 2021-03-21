package response

// StatusText godoc.
func StatusText(args ...string) JSONMessage {
	res := JSONMessage{
		"error pada system server",
		"internal server error",
	}

	return res
}
