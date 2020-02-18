package gim


type config struct {
	PagesFolder string
	PublicFolder string
	MintFolder string

	DistPublicFolder string
	DistBrowserFolder string
	TemplateHTMLName string

	ImbacCommand string
	ParcelCommand string
}


func newDefaultConfig() *config {
	return &config{
		PagesFolder:       "pages",
		PublicFolder:      "public",
		MintFolder:        ".mint",
		DistPublicFolder:  "content",
		DistBrowserFolder: "browser",
		TemplateHTMLName:  "template.html",
		ImbacCommand:      "imbac",
		ParcelCommand:     "parcel",
	}
}