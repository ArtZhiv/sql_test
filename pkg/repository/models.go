package repository

const (
	PathCommercial = "//192.168.37.222/24x7/Макросы и шаблоны/Комерция для шаблонов и макросов/"
	ConnDB         = "beclouderp:becloud$erp@tcp(192.168.37.65:3306)/beCloud_database?parseTime=true"
)

type eNodeb struct {
	id          int
	number      int    //номер
	dismantling string //демонтаж
	area        string //область
	district    string //район
	city        string //город
	address     string //адрес
	vendor      string //вендор
	location    string //место расположения
	mts         bool   //мтс
	life        bool   //лайф
	a1          bool   //а1
}

type sectors struct {
	id           int
	basestantion int    //номер
	cell_number  string //номер сектора
	bandwidth    int    //полоса
	mts          bool   //мтс
	life         bool   //лайф
	a1           bool   //а1
	beCloud      bool   //бот
}

type commercialUPD struct {
	id         int
	dateCreate string
}

type counter struct {
	id    int
	count string
}
