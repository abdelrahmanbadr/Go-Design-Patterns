package proxy_example2

type Printer interface {
	SetPrinterName(name string)
	GetPrinterName() string
	Print(str string) string
}

// @implements Printer
// will be private because we want to hide it and use the proxy
type realPrinter struct {
	name string
}

func (self *realPrinter) SetPrinterName(name string) {
	self.name = name
}

func (self *realPrinter) GetPrinterName() string {
	return self.name
}

func (self *realPrinter) Print(str string) string {
	return self.name + ":" + str
}

// ----------------------------------------------------------------------
// @implements Printer
type PrinterProxy struct {
	Name    string
	IsAdmin bool
	real    *realPrinter
}

func (self *PrinterProxy) SetPrinterName(name string) {
	if self.real != nil {
		self.real.SetPrinterName(name)
	}
	self.Name = name
}

func (self *PrinterProxy) GetPrinterName() string {
	return self.Name
}

func (self *PrinterProxy) Print(str string) string {
	if self.real == nil {
		self.real = &realPrinter{self.Name}
	}
	if self.IsAdmin {
		return self.real.Print(str)
	}
	return "You don't have the permission to use the printer"

}
