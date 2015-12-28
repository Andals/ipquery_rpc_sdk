// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package ipquery

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type Ipquery interface {
	// Parameters:
	//  - IP
	//  - Tag
	Ipquery(ip string, tag string) (r string, err error)
	// Parameters:
	//  - IpList
	//  - Tag
	IpqueryEx(ipList []string, tag string) (r []string, err error)
	IpqueryVersion() (r string, err error)
}

type IpqueryClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewIpqueryClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *IpqueryClient {
	return &IpqueryClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewIpqueryClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *IpqueryClient {
	return &IpqueryClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - IP
//  - Tag
func (p *IpqueryClient) Ipquery(ip string, tag string) (r string, err error) {
	if err = p.sendIpquery(ip, tag); err != nil {
		return
	}
	return p.recvIpquery()
}

func (p *IpqueryClient) sendIpquery(ip string, tag string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("ipquery", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := IpqueryIpqueryArgs{
		IP:  ip,
		Tag: tag,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *IpqueryClient) recvIpquery() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "ipquery" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "ipquery failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ipquery failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "ipquery failed: invalid message type")
		return
	}
	result := IpqueryIpqueryResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - IpList
//  - Tag
func (p *IpqueryClient) IpqueryEx(ipList []string, tag string) (r []string, err error) {
	if err = p.sendIpqueryEx(ipList, tag); err != nil {
		return
	}
	return p.recvIpqueryEx()
}

func (p *IpqueryClient) sendIpqueryEx(ipList []string, tag string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("ipqueryEx", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := IpqueryIpqueryExArgs{
		IpList: ipList,
		Tag:    tag,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *IpqueryClient) recvIpqueryEx() (value []string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "ipqueryEx" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "ipqueryEx failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ipqueryEx failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "ipqueryEx failed: invalid message type")
		return
	}
	result := IpqueryIpqueryExResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

func (p *IpqueryClient) IpqueryVersion() (r string, err error) {
	if err = p.sendIpqueryVersion(); err != nil {
		return
	}
	return p.recvIpqueryVersion()
}

func (p *IpqueryClient) sendIpqueryVersion() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("ipqueryVersion", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := IpqueryIpqueryVersionArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *IpqueryClient) recvIpqueryVersion() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "ipqueryVersion" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "ipqueryVersion failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ipqueryVersion failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "ipqueryVersion failed: invalid message type")
		return
	}
	result := IpqueryIpqueryVersionResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type IpqueryProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Ipquery
}

func (p *IpqueryProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *IpqueryProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *IpqueryProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewIpqueryProcessor(handler Ipquery) *IpqueryProcessor {

	self6 := &IpqueryProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self6.processorMap["ipquery"] = &ipqueryProcessorIpquery{handler: handler}
	self6.processorMap["ipqueryEx"] = &ipqueryProcessorIpqueryEx{handler: handler}
	self6.processorMap["ipqueryVersion"] = &ipqueryProcessorIpqueryVersion{handler: handler}
	return self6
}

func (p *IpqueryProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x7.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x7

}

type ipqueryProcessorIpquery struct {
	handler Ipquery
}

func (p *ipqueryProcessorIpquery) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IpqueryIpqueryArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ipquery", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := IpqueryIpqueryResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.Ipquery(args.IP, args.Tag); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ipquery: "+err2.Error())
		oprot.WriteMessageBegin("ipquery", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("ipquery", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ipqueryProcessorIpqueryEx struct {
	handler Ipquery
}

func (p *ipqueryProcessorIpqueryEx) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IpqueryIpqueryExArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ipqueryEx", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := IpqueryIpqueryExResult{}
	var retval []string
	var err2 error
	if retval, err2 = p.handler.IpqueryEx(args.IpList, args.Tag); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ipqueryEx: "+err2.Error())
		oprot.WriteMessageBegin("ipqueryEx", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("ipqueryEx", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type ipqueryProcessorIpqueryVersion struct {
	handler Ipquery
}

func (p *ipqueryProcessorIpqueryVersion) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := IpqueryIpqueryVersionArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ipqueryVersion", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := IpqueryIpqueryVersionResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.IpqueryVersion(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ipqueryVersion: "+err2.Error())
		oprot.WriteMessageBegin("ipqueryVersion", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("ipqueryVersion", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - IP
//  - Tag
type IpqueryIpqueryArgs struct {
	IP  string `thrift:"ip,1" json:"ip"`
	Tag string `thrift:"tag,2" json:"tag"`
}

func NewIpqueryIpqueryArgs() *IpqueryIpqueryArgs {
	return &IpqueryIpqueryArgs{}
}

func (p *IpqueryIpqueryArgs) GetIP() string {
	return p.IP
}

func (p *IpqueryIpqueryArgs) GetTag() string {
	return p.Tag
}
func (p *IpqueryIpqueryArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IpqueryIpqueryArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.IP = v
	}
	return nil
}

func (p *IpqueryIpqueryArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Tag = v
	}
	return nil
}

func (p *IpqueryIpqueryArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ipquery_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IpqueryIpqueryArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ip", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ip: ", p), err)
	}
	if err := oprot.WriteString(string(p.IP)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ip (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ip: ", p), err)
	}
	return err
}

func (p *IpqueryIpqueryArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tag", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tag: ", p), err)
	}
	if err := oprot.WriteString(string(p.Tag)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.tag (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tag: ", p), err)
	}
	return err
}

func (p *IpqueryIpqueryArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IpqueryIpqueryArgs(%+v)", *p)
}

// Attributes:
//  - Success
type IpqueryIpqueryResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewIpqueryIpqueryResult() *IpqueryIpqueryResult {
	return &IpqueryIpqueryResult{}
}

var IpqueryIpqueryResult_Success_DEFAULT string

func (p *IpqueryIpqueryResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return IpqueryIpqueryResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *IpqueryIpqueryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IpqueryIpqueryResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IpqueryIpqueryResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *IpqueryIpqueryResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ipquery_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IpqueryIpqueryResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *IpqueryIpqueryResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IpqueryIpqueryResult(%+v)", *p)
}

// Attributes:
//  - IpList
//  - Tag
type IpqueryIpqueryExArgs struct {
	IpList []string `thrift:"ipList,1" json:"ipList"`
	Tag    string   `thrift:"tag,2" json:"tag"`
}

func NewIpqueryIpqueryExArgs() *IpqueryIpqueryExArgs {
	return &IpqueryIpqueryExArgs{}
}

func (p *IpqueryIpqueryExArgs) GetIpList() []string {
	return p.IpList
}

func (p *IpqueryIpqueryExArgs) GetTag() string {
	return p.Tag
}
func (p *IpqueryIpqueryExArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IpqueryIpqueryExArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.IpList = tSlice
	for i := 0; i < size; i++ {
		var _elem8 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem8 = v
		}
		p.IpList = append(p.IpList, _elem8)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *IpqueryIpqueryExArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Tag = v
	}
	return nil
}

func (p *IpqueryIpqueryExArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ipqueryEx_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IpqueryIpqueryExArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ipList", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ipList: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.IpList)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.IpList {
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ipList: ", p), err)
	}
	return err
}

func (p *IpqueryIpqueryExArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("tag", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tag: ", p), err)
	}
	if err := oprot.WriteString(string(p.Tag)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.tag (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tag: ", p), err)
	}
	return err
}

func (p *IpqueryIpqueryExArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IpqueryIpqueryExArgs(%+v)", *p)
}

// Attributes:
//  - Success
type IpqueryIpqueryExResult struct {
	Success []string `thrift:"success,0" json:"success,omitempty"`
}

func NewIpqueryIpqueryExResult() *IpqueryIpqueryExResult {
	return &IpqueryIpqueryExResult{}
}

var IpqueryIpqueryExResult_Success_DEFAULT []string

func (p *IpqueryIpqueryExResult) GetSuccess() []string {
	return p.Success
}
func (p *IpqueryIpqueryExResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IpqueryIpqueryExResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IpqueryIpqueryExResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		var _elem9 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem9 = v
		}
		p.Success = append(p.Success, _elem9)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *IpqueryIpqueryExResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ipqueryEx_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IpqueryIpqueryExResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := oprot.WriteString(string(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *IpqueryIpqueryExResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IpqueryIpqueryExResult(%+v)", *p)
}

type IpqueryIpqueryVersionArgs struct {
}

func NewIpqueryIpqueryVersionArgs() *IpqueryIpqueryVersionArgs {
	return &IpqueryIpqueryVersionArgs{}
}

func (p *IpqueryIpqueryVersionArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IpqueryIpqueryVersionArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ipqueryVersion_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IpqueryIpqueryVersionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IpqueryIpqueryVersionArgs(%+v)", *p)
}

// Attributes:
//  - Success
type IpqueryIpqueryVersionResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewIpqueryIpqueryVersionResult() *IpqueryIpqueryVersionResult {
	return &IpqueryIpqueryVersionResult{}
}

var IpqueryIpqueryVersionResult_Success_DEFAULT string

func (p *IpqueryIpqueryVersionResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return IpqueryIpqueryVersionResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *IpqueryIpqueryVersionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *IpqueryIpqueryVersionResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IpqueryIpqueryVersionResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *IpqueryIpqueryVersionResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ipqueryVersion_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IpqueryIpqueryVersionResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *IpqueryIpqueryVersionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IpqueryIpqueryVersionResult(%+v)", *p)
}