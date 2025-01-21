package playlist

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i PlaylistRepository -o ./mocks/ -s "_minimock.go"
