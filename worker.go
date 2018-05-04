package Kopie

//import (
//	"fmt"
//	"github.com/jinzhu/gorm"
//)
//
//type Worker struct {
//	Config        Config
//	Local         Database
//	Remote        Database
//	Procedure     Procedure
//	Specification Specification
//	LinkName      string
//}
//
//func (w *Worker) Init(spec Specification) {
//	var err error
//	// Set the specification
//	w.Specification = spec
//	w.Local.Con, err = gorm.Open(w.Local.Type, w.Local.Url)
//	if err != nil {
//		panic(err)
//	}
//	w.Remote.Con, err = gorm.Open(w.Remote.Type, w.Remote.Url)
//	if err != nil {
//		panic(err)
//	}
//	w.LinkName = w.Procedure.Name
//}
//
//func (w *Worker) ExecuteProcedure() error {
//	fmt.Println(w.Procedure.Type)
//	fmt.Println(w.Specification)
//	return nil
//}
//
//func (w *Worker) Start() {
//	for {
//		err := w.ExecuteProcedure()
//		if err != nil {
//			w.BackOffProcedure(err)
//		}
//	}
//}
//
//func (w *Worker) BackOffProcedure(err error) {
//	fmt.Println(err)
//}
