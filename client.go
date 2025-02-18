//go:build !cgo
// +build !cgo

package gosseract

import (
	"errors"
	"fmt"
	"image"
	"os"
)

var ErrNotImplementWithoutCGO = errors.New("Not implement when without cgo")

// Version returns the version of Tesseract-OCR
func Version() string {
	return ""
}

// ClearPersistentCache clears any library-level memory caches. There are a variety of expensive-to-load constant data structures (mostly language dictionaries) that are cached globally – surviving the Init() and End() of individual TessBaseAPI's. This function allows the clearing of these caches.
func ClearPersistentCache() {
}

// Client is argument builder for tesseract::TessBaseAPI.
type Client struct {

	// Trim specifies characters to trim, which would be trimed from result string.
	// As results of OCR, text often contains unnecessary characters, such as newlines, on the head/foot of string.
	// If `Trim` is set, this client will remove specified characters from the result.
	Trim bool

	// TessdataPrefix can indicate directory path to `tessdata`.
	// It is set `/usr/local/share/tessdata/` or something like that, as default.
	// TODO: Implement and test
	TessdataPrefix string

	// Languages are languages to be detected. If not specified, it's gonna be "eng".
	Languages []string

	// Variables is just a pool to evaluate "tesseract::TessBaseAPI->SetVariable" in delay.
	// TODO: Think if it should be public, or private property.
	Variables map[SettableVariable]string

	// Config is a file path to the configuration for Tesseract
	// See http://www.sk-spell.sk.cx/tesseract-ocr-parameters-in-302-version
	// TODO: Fix link to official page
	ConfigFilePath string

	// internal flag to check if the instance should be initialized again
	// i.e, we should create a new gosseract client when language or config file change
	shouldInit bool
}

// NewClient construct new Client. It's due to caller to Close this client.
func NewClient() *Client {
	client := &Client{
		Variables:  map[SettableVariable]string{},
		Trim:       true,
		shouldInit: true,
		Languages:  []string{"eng"},
	}
	return client
}

// Close frees allocated API. This MUST be called for ANY client constructed by "NewClient" function.
func (client *Client) Close() (err error) {
	return ErrNotImplementWithoutCGO
}

// Version provides the version of Tesseract used by this client.
func (client *Client) Version() string {
	return ""
}

// SetImage sets path to image file to be processed OCR.
func (client *Client) SetImage(imagepath string) error {
	return nil
}

// SetImageFromBytes sets the image data to be processed OCR.
func (client *Client) SetImageFromBytes(data []byte) error {
	return ErrNotImplementWithoutCGO
}

// SetLanguage sets languages to use. English as default.
func (client *Client) SetLanguage(langs ...string) error {
	if len(langs) == 0 {
		return fmt.Errorf("languages cannot be empty")
	}

	client.Languages = langs

	client.flagForInit()

	return nil
}

// DisableOutput ...
func (client *Client) DisableOutput() error {
	err := client.SetVariable(DEBUG_FILE, os.DevNull)

	client.setVariablesToInitializedAPIIfNeeded()

	return err
}

// SetWhitelist sets whitelist chars.
// See official documentation for whitelist here https://tesseract-ocr.github.io/tessdoc/ImproveQuality#dictionaries-word-lists-and-patterns
func (client *Client) SetWhitelist(whitelist string) error {
	err := client.SetVariable(TESSEDIT_CHAR_WHITELIST, whitelist)

	client.setVariablesToInitializedAPIIfNeeded()

	return err
}

// SetBlacklist sets blacklist chars.
// See official documentation for blacklist here https://tesseract-ocr.github.io/tessdoc/ImproveQuality#dictionaries-word-lists-and-patterns
func (client *Client) SetBlacklist(blacklist string) error {
	err := client.SetVariable(TESSEDIT_CHAR_BLACKLIST, blacklist)

	client.setVariablesToInitializedAPIIfNeeded()

	return err
}

// SetVariable sets parameters, representing tesseract::TessBaseAPI->SetVariable.
// See official documentation here https://zdenop.github.io/tesseract-doc/classtesseract_1_1_tess_base_a_p_i.html#a2e09259c558c6d8e0f7e523cbaf5adf5
// Because `api->SetVariable` must be called after `api->Init`, this method cannot detect unexpected key for variables.
// Check `client.setVariablesToInitializedAPI` for more information.
func (client *Client) SetVariable(key SettableVariable, value string) error {
	return ErrNotImplementWithoutCGO
}

// SetPageSegMode sets "Page Segmentation Mode" (PSM) to detect layout of characters.
// See official documentation for PSM here https://tesseract-ocr.github.io/tessdoc/ImproveQuality#page-segmentation-method
// See https://github.com/otiai10/gosseract/issues/52 for more information.
func (client *Client) SetPageSegMode(mode PageSegMode) error {
	return ErrNotImplementWithoutCGO
}

// SetConfigFile sets the file path to config file.
func (client *Client) SetConfigFile(fpath string) error {
	return ErrNotImplementWithoutCGO
}

// SetTessdataPrefix sets path to the models directory.
// Environment variable TESSDATA_PREFIX is used as default.
func (client *Client) SetTessdataPrefix(prefix string) error {
	return ErrNotImplementWithoutCGO
}

// Initialize tesseract::TessBaseAPI
func (client *Client) init() error {
	return ErrNotImplementWithoutCGO
}

// This method flag the current instance to be initialized again on the next call to a function that
// requires a gosseract API initialized: when user change the config file or the languages
// the instance needs to init a new gosseract api
func (client *Client) flagForInit() {
	client.shouldInit = true
}

// This method sets all the sspecified variables to TessBaseAPI structure.
// Because `api->SetVariable` must be called after `api->Init()`,
// gosseract.Client.SetVariable cannot call `api->SetVariable` directly.
// See https://zdenop.github.io/tesseract-doc/classtesseract_1_1_tess_base_a_p_i.html#a2e09259c558c6d8e0f7e523cbaf5adf5
func (client *Client) setVariablesToInitializedAPI() error {

	return nil
}

// Call setVariablesToInitializedAPI only if the API is initialized
// it is useful to call when changing variables that does not requires
// to init a new tesseract instance. Otherwise it is better to just flag
// the instance for re-init (Client.flagForInit())
func (client *Client) setVariablesToInitializedAPIIfNeeded() error {
	return ErrNotImplementWithoutCGO
}

// Text finally initialize tesseract::TessBaseAPI, execute OCR and extract text detected as string.
func (client *Client) Text() (out string, err error) {
	return out, ErrNotImplementWithoutCGO

}

// HOCRText finally initialize tesseract::TessBaseAPI, execute OCR and returns hOCR text.
// See https://en.wikipedia.org/wiki/HOCR for more information of hOCR.
func (client *Client) HOCRText() (out string, err error) {
	return out, ErrNotImplementWithoutCGO
}

// BoundingBox contains the position, confidence and UTF8 text of the recognized word
type BoundingBox struct {
	Box                                image.Rectangle
	Word                               string
	Confidence                         float64
	BlockNum, ParNum, LineNum, WordNum int
}

// GetBoundingBoxes returns bounding boxes for each matched word
func (client *Client) GetBoundingBoxes(level PageIteratorLevel) (out []BoundingBox, err error) {
	return nil, ErrNotImplementWithoutCGO
}

// GetAvailableLanguages returns a list of available languages in the default tesspath
func GetAvailableLanguages() ([]string, error) {
	return nil, ErrNotImplementWithoutCGO
}

// GetBoundingBoxesVerbose returns bounding boxes at word level with block_num, par_num, line_num and word_num
// according to the c++ api that returns a formatted TSV output. Reference: `TessBaseAPI::GetTSVText`.
func (client *Client) GetBoundingBoxesVerbose() (out []BoundingBox, err error) {
	return nil, ErrNotImplementWithoutCGO
}
