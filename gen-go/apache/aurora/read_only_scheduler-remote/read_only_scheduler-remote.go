// Code generated by Thrift Compiler (0.14.1). DO NOT EDIT.

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
	"aurora"
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
    arg145 := flag.Arg(1)
    mbTrans146 := thrift.NewTMemoryBufferLen(len(arg145))
    defer mbTrans146.Close()
    _, err147 := mbTrans146.WriteString(arg145)
    if err147 != nil {
      Usage()
      return
    }
    factory148 := thrift.NewTJSONProtocolFactory()
    jsProt149 := factory148.GetProtocol(mbTrans146)
    argvalue0 := aurora.NewTaskQuery()
    err150 := argvalue0.Read(context.Background(), jsProt149)
    if err150 != nil {
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
    arg151 := flag.Arg(1)
    mbTrans152 := thrift.NewTMemoryBufferLen(len(arg151))
    defer mbTrans152.Close()
    _, err153 := mbTrans152.WriteString(arg151)
    if err153 != nil {
      Usage()
      return
    }
    factory154 := thrift.NewTJSONProtocolFactory()
    jsProt155 := factory154.GetProtocol(mbTrans152)
    argvalue0 := aurora.NewTaskQuery()
    err156 := argvalue0.Read(context.Background(), jsProt155)
    if err156 != nil {
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
    arg157 := flag.Arg(1)
    mbTrans158 := thrift.NewTMemoryBufferLen(len(arg157))
    defer mbTrans158.Close()
    _, err159 := mbTrans158.WriteString(arg157)
    if err159 != nil {
      Usage()
      return
    }
    factory160 := thrift.NewTJSONProtocolFactory()
    jsProt161 := factory160.GetProtocol(mbTrans158)
    argvalue0 := aurora.NewTaskQuery()
    err162 := argvalue0.Read(context.Background(), jsProt161)
    if err162 != nil {
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
    arg163 := flag.Arg(1)
    mbTrans164 := thrift.NewTMemoryBufferLen(len(arg163))
    defer mbTrans164.Close()
    _, err165 := mbTrans164.WriteString(arg163)
    if err165 != nil {
      Usage()
      return
    }
    factory166 := thrift.NewTJSONProtocolFactory()
    jsProt167 := factory166.GetProtocol(mbTrans164)
    argvalue0 := aurora.NewJobKey()
    err168 := argvalue0.Read(context.Background(), jsProt167)
    if err168 != nil {
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
    arg171 := flag.Arg(1)
    mbTrans172 := thrift.NewTMemoryBufferLen(len(arg171))
    defer mbTrans172.Close()
    _, err173 := mbTrans172.WriteString(arg171)
    if err173 != nil {
      Usage()
      return
    }
    factory174 := thrift.NewTJSONProtocolFactory()
    jsProt175 := factory174.GetProtocol(mbTrans172)
    argvalue0 := aurora.NewJobConfiguration()
    err176 := argvalue0.Read(context.Background(), jsProt175)
    if err176 != nil {
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
    arg177 := flag.Arg(1)
    mbTrans178 := thrift.NewTMemoryBufferLen(len(arg177))
    defer mbTrans178.Close()
    _, err179 := mbTrans178.WriteString(arg177)
    if err179 != nil {
      Usage()
      return
    }
    factory180 := thrift.NewTJSONProtocolFactory()
    jsProt181 := factory180.GetProtocol(mbTrans178)
    argvalue0 := aurora.NewJobUpdateQuery()
    err182 := argvalue0.Read(context.Background(), jsProt181)
    if err182 != nil {
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
    arg183 := flag.Arg(1)
    mbTrans184 := thrift.NewTMemoryBufferLen(len(arg183))
    defer mbTrans184.Close()
    _, err185 := mbTrans184.WriteString(arg183)
    if err185 != nil {
      Usage()
      return
    }
    factory186 := thrift.NewTJSONProtocolFactory()
    jsProt187 := factory186.GetProtocol(mbTrans184)
    argvalue0 := aurora.NewJobUpdateQuery()
    err188 := argvalue0.Read(context.Background(), jsProt187)
    if err188 != nil {
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
    arg189 := flag.Arg(1)
    mbTrans190 := thrift.NewTMemoryBufferLen(len(arg189))
    defer mbTrans190.Close()
    _, err191 := mbTrans190.WriteString(arg189)
    if err191 != nil {
      Usage()
      return
    }
    factory192 := thrift.NewTJSONProtocolFactory()
    jsProt193 := factory192.GetProtocol(mbTrans190)
    argvalue0 := aurora.NewJobUpdateRequest()
    err194 := argvalue0.Read(context.Background(), jsProt193)
    if err194 != nil {
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
