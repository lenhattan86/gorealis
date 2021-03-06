// Code generated by Thrift Compiler (0.14.0). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/apache/thrift/lib/go/thrift"
	"apache/aurora"
)

var _ = aurora.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  Response getRoleSummary()")
  fmt.Fprintln(os.Stderr, "  Response getJobSummary(string role)")
  fmt.Fprintln(os.Stderr, "  Response getTasksStatus(TaskQuery query)")
  fmt.Fprintln(os.Stderr, "  Response getTasksWithoutConfigs(TaskQuery query)")
  fmt.Fprintln(os.Stderr, "  Response getPendingReason(TaskQuery query)")
  fmt.Fprintln(os.Stderr, "  Response getConfigSummary(JobKey job)")
  fmt.Fprintln(os.Stderr, "  Response getJobs(string ownerRole)")
  fmt.Fprintln(os.Stderr, "  Response getQuota(string ownerRole)")
  fmt.Fprintln(os.Stderr, "  Response populateJobConfig(JobConfiguration description)")
  fmt.Fprintln(os.Stderr, "  Response getJobUpdateSummaries(JobUpdateQuery jobUpdateQuery)")
  fmt.Fprintln(os.Stderr, "  Response getJobUpdateDetails(JobUpdateQuery query)")
  fmt.Fprintln(os.Stderr, "  Response getJobUpdateDiff(JobUpdateRequest request)")
  fmt.Fprintln(os.Stderr, "  Response getTierConfigs()")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := aurora.NewReadOnlySchedulerClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "getRoleSummary":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetRoleSummary requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetRoleSummary(context.Background()))
    fmt.Print("\n")
    break
  case "getJobSummary":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetJobSummary requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetJobSummary(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getTasksStatus":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTasksStatus requires 1 args")
      flag.Usage()
    }
    arg132 := flag.Arg(1)
    mbTrans133 := thrift.NewTMemoryBufferLen(len(arg132))
    defer mbTrans133.Close()
    _, err134 := mbTrans133.WriteString(arg132)
    if err134 != nil {
      Usage()
      return
    }
    factory135 := thrift.NewTJSONProtocolFactory()
    jsProt136 := factory135.GetProtocol(mbTrans133)
    argvalue0 := aurora.NewTaskQuery()
    err137 := argvalue0.Read(context.Background(), jsProt136)
    if err137 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTasksStatus(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getTasksWithoutConfigs":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTasksWithoutConfigs requires 1 args")
      flag.Usage()
    }
    arg138 := flag.Arg(1)
    mbTrans139 := thrift.NewTMemoryBufferLen(len(arg138))
    defer mbTrans139.Close()
    _, err140 := mbTrans139.WriteString(arg138)
    if err140 != nil {
      Usage()
      return
    }
    factory141 := thrift.NewTJSONProtocolFactory()
    jsProt142 := factory141.GetProtocol(mbTrans139)
    argvalue0 := aurora.NewTaskQuery()
    err143 := argvalue0.Read(context.Background(), jsProt142)
    if err143 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTasksWithoutConfigs(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getPendingReason":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetPendingReason requires 1 args")
      flag.Usage()
    }
    arg144 := flag.Arg(1)
    mbTrans145 := thrift.NewTMemoryBufferLen(len(arg144))
    defer mbTrans145.Close()
    _, err146 := mbTrans145.WriteString(arg144)
    if err146 != nil {
      Usage()
      return
    }
    factory147 := thrift.NewTJSONProtocolFactory()
    jsProt148 := factory147.GetProtocol(mbTrans145)
    argvalue0 := aurora.NewTaskQuery()
    err149 := argvalue0.Read(context.Background(), jsProt148)
    if err149 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetPendingReason(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getConfigSummary":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetConfigSummary requires 1 args")
      flag.Usage()
    }
    arg150 := flag.Arg(1)
    mbTrans151 := thrift.NewTMemoryBufferLen(len(arg150))
    defer mbTrans151.Close()
    _, err152 := mbTrans151.WriteString(arg150)
    if err152 != nil {
      Usage()
      return
    }
    factory153 := thrift.NewTJSONProtocolFactory()
    jsProt154 := factory153.GetProtocol(mbTrans151)
    argvalue0 := aurora.NewJobKey()
    err155 := argvalue0.Read(context.Background(), jsProt154)
    if err155 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetConfigSummary(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getJobs":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetJobs requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetJobs(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getQuota":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetQuota requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetQuota(context.Background(), value0))
    fmt.Print("\n")
    break
  case "populateJobConfig":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "PopulateJobConfig requires 1 args")
      flag.Usage()
    }
    arg158 := flag.Arg(1)
    mbTrans159 := thrift.NewTMemoryBufferLen(len(arg158))
    defer mbTrans159.Close()
    _, err160 := mbTrans159.WriteString(arg158)
    if err160 != nil {
      Usage()
      return
    }
    factory161 := thrift.NewTJSONProtocolFactory()
    jsProt162 := factory161.GetProtocol(mbTrans159)
    argvalue0 := aurora.NewJobConfiguration()
    err163 := argvalue0.Read(context.Background(), jsProt162)
    if err163 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.PopulateJobConfig(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getJobUpdateSummaries":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetJobUpdateSummaries requires 1 args")
      flag.Usage()
    }
    arg164 := flag.Arg(1)
    mbTrans165 := thrift.NewTMemoryBufferLen(len(arg164))
    defer mbTrans165.Close()
    _, err166 := mbTrans165.WriteString(arg164)
    if err166 != nil {
      Usage()
      return
    }
    factory167 := thrift.NewTJSONProtocolFactory()
    jsProt168 := factory167.GetProtocol(mbTrans165)
    argvalue0 := aurora.NewJobUpdateQuery()
    err169 := argvalue0.Read(context.Background(), jsProt168)
    if err169 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetJobUpdateSummaries(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getJobUpdateDetails":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetJobUpdateDetails requires 1 args")
      flag.Usage()
    }
    arg170 := flag.Arg(1)
    mbTrans171 := thrift.NewTMemoryBufferLen(len(arg170))
    defer mbTrans171.Close()
    _, err172 := mbTrans171.WriteString(arg170)
    if err172 != nil {
      Usage()
      return
    }
    factory173 := thrift.NewTJSONProtocolFactory()
    jsProt174 := factory173.GetProtocol(mbTrans171)
    argvalue0 := aurora.NewJobUpdateQuery()
    err175 := argvalue0.Read(context.Background(), jsProt174)
    if err175 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetJobUpdateDetails(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getJobUpdateDiff":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetJobUpdateDiff requires 1 args")
      flag.Usage()
    }
    arg176 := flag.Arg(1)
    mbTrans177 := thrift.NewTMemoryBufferLen(len(arg176))
    defer mbTrans177.Close()
    _, err178 := mbTrans177.WriteString(arg176)
    if err178 != nil {
      Usage()
      return
    }
    factory179 := thrift.NewTJSONProtocolFactory()
    jsProt180 := factory179.GetProtocol(mbTrans177)
    argvalue0 := aurora.NewJobUpdateRequest()
    err181 := argvalue0.Read(context.Background(), jsProt180)
    if err181 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetJobUpdateDiff(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getTierConfigs":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetTierConfigs requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetTierConfigs(context.Background()))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
