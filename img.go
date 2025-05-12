package dtest

func RunImage() {
	lgr.Debug("No command given, defaulting to dtest img ls")
	RunImageLs()

}

func RunImageLs() {
	lgr.Info("Called image list operation")
}
