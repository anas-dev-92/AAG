package models

//here will store data into database or get data from database  etc

// in this struct we put the same field we has in the uploaddata form
type UploadData struct {
	FileName       string
	PublishingDate string
	Country        string
	FileTypes      string
}
