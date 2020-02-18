package gim


type Config struct {
	PagesFolder  string
	PublicFolder string
	GimFolder    string

	DistPublicFolder string
	DistBrowserFolder string
	TemplateHTMLName string

	ImbacCommand string
	ParcelCommand string
}


func NewDefaultConfig() *Config {
	return &Config{
		PagesFolder:       "pages",
		PublicFolder:      "public",
		GimFolder:         ".gim",
		DistPublicFolder:  "content",
		DistBrowserFolder: "browser",
		TemplateHTMLName:  "template.html",
		ImbacCommand:      "imbac",
		ParcelCommand:     "parcel",
	}
}

