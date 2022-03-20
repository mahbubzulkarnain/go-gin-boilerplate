package response

// StatusText ...
func StatusText(args ...string) JSONMessage {
	res := JSONMessage{
		"berhasil",
		"success",
	}

	return res
}
