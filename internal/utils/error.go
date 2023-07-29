package utils

func PrintErr(err error) {
	if err != nil {
		Log.Error(err)
	}
}
