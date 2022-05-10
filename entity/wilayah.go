package entity

type Provinces struct {
	Prov_id    int
	Prov_name  string
	Locationid int
	Status     int
}

type Cities struct {
	City_id   int16
	City_name string
	Prov_id   int8
}

type Districts struct {
	Dis_id   int16
	Dis_name string
	City_id  int16
}

type Subdistricts struct {
	Subdis_id   int
	Subdis_name string
	Dis_id      int16
}