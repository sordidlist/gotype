package configure

func GetInitialConfigs() {
	loadUserData()
}

func loadUserData() {
	if userConfigsExist() {
		loadUserSettings()
		loadUserProgress()
	} else {

	}
}

func loadUserSettings() {

}

func loadUserProgress() {

}

func userConfigsExist() bool {
	return true
}
