package mkfile

import (
	"os"
	"testing"
)

// Testing file creating

func TestCreateFile(t *testing.T){
	// setup
	tests := []struct {
		name string
		filename string
		wantErr bool
	}{
		{
            name:     "Create New File",
            filename: "file.txt",
            wantErr:  false,
        },
        {
            name:     "Empty Filename",
            filename: "",  // Actually empty now
            wantErr:  true,
        },
        // Could add more cases like:
        {
            name:     "Invalid Characters",
            filename: "file*.txt",
            wantErr:  true,
        },
        {
            name:     "Long Path",
            filename: "dir/subdir/file.txt",
            wantErr:  true,
        },
	}

	// run tests
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			defer func(){
				os.Remove(tt.filename)
			}()

			err := CreateFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if _, err := os.Stat(tt.filename); os.IsNotExist(err) {
					t.Errorf("CreateFile() file not created")
				}
			}
		})
	}
}