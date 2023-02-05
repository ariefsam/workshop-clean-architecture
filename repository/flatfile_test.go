package repository_test

type flatFileRepository struct{}

func NewFlatFile(folderPath string) (repo *flatFileRepository) {
	repo = &flatFileRepository{}
	return
}
