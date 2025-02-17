package dialog

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Anwert/fyne/v2/canvas"
	"github.com/Anwert/fyne/v2/storage"
	"github.com/Anwert/fyne/v2/test"
)

func TestFileItem_Name(t *testing.T) {
	f := &fileDialog{file: &FileDialog{}}
	_ = f.makeUI()

	item := f.newFileItem(storage.NewFileURI("/path/to/filename.txt"), false, false)
	assert.Equal(t, "filename", item.name)

	item = f.newFileItem(storage.NewFileURI("/path/to/MyFile.jpeg"), false, false)
	assert.Equal(t, "MyFile", item.name)

	item = f.newFileItem(storage.NewFileURI("/path/to/.maybeHidden.txt"), false, false)
	assert.Equal(t, ".maybeHidden", item.name)

	item = f.newFileItem(storage.NewFileURI("/path/to/noext"), false, false)
	assert.Equal(t, "noext", item.name)

	// Test that the extension remains for the list view.
	f.view = listView

	item = f.newFileItem(storage.NewFileURI("/path/to/filename.txt"), false, false)
	assert.Equal(t, "filename.txt", item.name)

	item = f.newFileItem(storage.NewFileURI("/path/to/MyFile.jpeg"), false, false)
	assert.Equal(t, "MyFile.jpeg", item.name)

	item = f.newFileItem(storage.NewFileURI("/path/to/.maybeHidden.txt"), false, false)
	assert.Equal(t, ".maybeHidden.txt", item.name)

	item = f.newFileItem(storage.NewFileURI("/path/to/noext"), false, false)
	assert.Equal(t, "noext", item.name)
}

func TestFileItem_FolderName(t *testing.T) {
	f := &fileDialog{file: &FileDialog{}}
	_ = f.makeUI()

	item := f.newFileItem(storage.NewFileURI("/path/to/foldername/"), true, false)
	assert.Equal(t, "foldername", item.name)

	item = f.newFileItem(storage.NewFileURI("/path/to/myapp.app/"), true, false)
	assert.Equal(t, "myapp", item.name)

	item = f.newFileItem(storage.NewFileURI("/path/to/.maybeHidden/"), true, false)
	assert.Equal(t, ".maybeHidden", item.name)

	// Test that the extension remains for the list view.
	f.view = listView
	item = f.newFileItem(storage.NewFileURI("/path/to/myapp.app/"), true, false)
	assert.Equal(t, "myapp.app", item.name)
}

func TestNewFileItem(t *testing.T) {
	f := &fileDialog{file: &FileDialog{}}
	_ = f.makeUI()
	item := f.newFileItem(storage.NewFileURI("/path/to/filename.txt"), false, false)

	assert.Equal(t, "filename", item.name)
}

func TestNewFileItem_Folder(t *testing.T) {
	f := &fileDialog{file: &FileDialog{}}
	_ = f.makeUI()
	currentDir, _ := filepath.Abs(".")
	currentLister, err := storage.ListerForURI(storage.NewFileURI(currentDir))
	if err != nil {
		t.Error(err)
	}

	parentDir := storage.NewFileURI(filepath.Dir(currentDir))
	parentLister, err := storage.ListerForURI(parentDir)
	if err != nil {
		t.Error(err)
	}
	f.setLocation(parentLister)
	item := f.newFileItem(currentLister, true, false)

	assert.Equal(t, filepath.Base(currentDir), item.name)
	assert.Equal(t, storage.NewFileURI(""+currentDir).String(), item.location.String())
}

func TestNewFileItem_ParentFolder(t *testing.T) {
	f := &fileDialog{file: &FileDialog{}}
	_ = f.makeUI()
	currentDir, _ := filepath.Abs(".")
	currentLister, err := storage.ListerForURI(storage.NewFileURI(currentDir))
	if err != nil {
		t.Error(err)
	}
	parentDir := storage.NewFileURI(filepath.Dir(currentDir))
	f.setLocation(currentLister)

	item := f.newFileItem(parentDir, true, true)
	item.ExtendBaseWidget(item)

	assert.Equal(t, "(Parent)", item.name)
	assert.Equal(t, parentDir.String()+"/", f.data[0].String())
}

func TestFileItem_Wrap(t *testing.T) {
	f := &fileDialog{file: &FileDialog{}}
	_ = f.makeUI()
	item := f.newFileItem(storage.NewFileURI("/path/to/filename.txt"), false, false)
	item.Resize(item.MinSize())
	label := test.WidgetRenderer(item).(*fileItemRenderer).text
	assert.Equal(t, "filename", label.Text)
	texts := test.WidgetRenderer(label).Objects()
	assert.Equal(t, 1, len(texts))

	item.setLocation(storage.NewFileURI("/path/to/averylongfilename.svg"), false, false)
	texts = test.WidgetRenderer(label).Objects()
	assert.Equal(t, 2, len(texts))
	assert.Equal(t, "averylon", texts[0].(*canvas.Text).Text)
}
