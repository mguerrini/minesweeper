package logger

type dummyLog struct{

}

func NewDummyLog() *dummyLog {
	return &dummyLog{}
}

func CreateDummyLog (configurationName string) (interface{}, error) {
	return NewDummyLog(), nil
}


func (this *dummyLog) Info(msg string) {
}

func (this *dummyLog) Error(msg string) {
}

