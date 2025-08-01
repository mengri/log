package init

import (
	"github.com/mengri/log"
	"github.com/mengri/log/filelog"
	"github.com/mengri/utils/autowire-v2"
	"github.com/mengri/utils/cftool"
	"os"
)

func init() {
	cftool.Register[ErrorLogConfig]("error_log")
	autowire.Autowired(new(logInit))

}

type logInit struct {
	config *ErrorLogConfig `autowired:""`
}

func (m *logInit) OnComplete() {
	formatter := &log.LineFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		CallerPrettyfier: nil,
	}

	fileWriter := filelog.NewFileWriteByPeriod(filelog.Config{
		Dir:    m.config.GetLogDir(),
		File:   m.config.GetLogName(),
		Expire: m.config.GetLogExpire(),
		Period: filelog.ParsePeriod(m.config.GetLogPeriod()),
	})

	writer := ToCopyToIoWriter(os.Stderr, fileWriter)

	transport := log.NewTransport(writer, m.config.GetLogLevel())
	//plugin_client.SetLog(m.config.GetLogLevel().String(), writer)
	transport.SetFormatter(formatter)
	log.Reset(transport)
}
