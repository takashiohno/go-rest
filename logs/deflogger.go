package logs

import (
    "io"
    "io/ioutil"
    "log"
    "math"
    "os"
    "time"

    "github.com/takashiohno/go-rest/config"
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var (
    Fatal *log.Logger
    Error *log.Logger
    Warn *log.Logger
    Info *log.Logger
    Debug *log.Logger
    Trace *log.Logger
)

var logLevel int

const (
    fatal int = 1 + iota
    err
    warn
    info
    debug
    trace
)

func init() {

    // set log level
    defConf := config.DefaultConf()
    logLevel = int(math.Min(float64(trace), math.Max(float64(fatal), float64(defConf.LogLevel))))

    // set log rotate
    var logf io.Writer
    var rlerr error
    logf, rlerr = rotatelogs.New(
        getFilePath(),
        rotatelogs.WithRotationTime(time.Hour),
    )
    if rlerr != nil {
        logf = os.Stdout
    }

    handle := ioutil.Discard
    if logLevel >= fatal {
        handle = os.Stderr
    }
    Fatal = log.New(handle, "FATAL: ", defConf.LogFlag)
    Fatal.SetOutput(logf)

    handle = ioutil.Discard
    if logLevel >= err {
        handle = os.Stderr
    }
    Error = log.New(handle, "ERROR: ", defConf.LogFlag)
    Error.SetOutput(logf)

    handle = ioutil.Discard
    if logLevel >= warn {
        handle = os.Stdout
    }
    Warn = log.New(handle, "WARN: ", defConf.LogFlag)
    Warn.SetOutput(logf)

    handle = ioutil.Discard
    if logLevel >= info {
        handle = os.Stdout
    }
    Info = log.New(handle, "INFO: ", defConf.LogFlag)
    Info.SetOutput(logf)

    handle = ioutil.Discard
    if logLevel >= debug {
        handle = os.Stdout
    }
    Debug = log.New(handle, "DEBUG: ", defConf.LogFlag)
    Debug.SetOutput(logf)

    handle = ioutil.Discard
    if logLevel >= trace {
        handle = os.Stdout
    }
    Trace = log.New(handle, "TRACE: ", defConf.LogFlag)
    Trace.SetOutput(logf)
}

func getFilePath() (string) {
/*
    currentUser, _ := user.Current()
    var buf bytes.Buffer
    buf.Write([]byte(currentUser.HomeDir))
    buf.Write([]byte("/deflogger.log.%Y%m%d"))
    return buf.String()
*/
    return "./deflogger.%Y%m%d.log"
}
