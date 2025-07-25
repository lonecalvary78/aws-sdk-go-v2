// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentcore

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentcore/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"sync"
)

// CodeInterpreterStreamOutputReader provides the interface for reading events
// from a stream.
//
// The writer's Close method must allow multiple concurrent calls.
type CodeInterpreterStreamOutputReader interface {
	Events() <-chan types.CodeInterpreterStreamOutput
	Close() error
	Err() error
}

type codeInterpreterStreamOutputReader struct {
	stream      chan types.CodeInterpreterStreamOutput
	decoder     *eventstream.Decoder
	eventStream io.ReadCloser
	err         *smithysync.OnceErr
	payloadBuf  []byte
	done        chan struct{}
	closeOnce   sync.Once
}

func newCodeInterpreterStreamOutputReader(readCloser io.ReadCloser, decoder *eventstream.Decoder) *codeInterpreterStreamOutputReader {
	w := &codeInterpreterStreamOutputReader{
		stream:      make(chan types.CodeInterpreterStreamOutput),
		decoder:     decoder,
		eventStream: readCloser,
		err:         smithysync.NewOnceErr(),
		done:        make(chan struct{}),
		payloadBuf:  make([]byte, 10*1024),
	}

	go w.readEventStream()

	return w
}

func (r *codeInterpreterStreamOutputReader) Events() <-chan types.CodeInterpreterStreamOutput {
	return r.stream
}

func (r *codeInterpreterStreamOutputReader) readEventStream() {
	defer r.Close()
	defer close(r.stream)

	for {
		r.payloadBuf = r.payloadBuf[0:0]
		decodedMessage, err := r.decoder.Decode(r.eventStream, r.payloadBuf)
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case <-r.done:
				return
			default:
				r.err.SetError(err)
				return
			}
		}

		event, err := r.deserializeEventMessage(&decodedMessage)
		if err != nil {
			r.err.SetError(err)
			return
		}

		select {
		case r.stream <- event:
		case <-r.done:
			return
		}

	}
}

func (r *codeInterpreterStreamOutputReader) deserializeEventMessage(msg *eventstream.Message) (types.CodeInterpreterStreamOutput, error) {
	messageType := msg.Headers.Get(eventstreamapi.MessageTypeHeader)
	if messageType == nil {
		return nil, fmt.Errorf("%s event header not present", eventstreamapi.MessageTypeHeader)
	}

	switch messageType.String() {
	case eventstreamapi.EventMessageType:
		var v types.CodeInterpreterStreamOutput
		if err := awsRestjson1_deserializeEventStreamCodeInterpreterStreamOutput(&v, msg); err != nil {
			return nil, err
		}
		return v, nil

	case eventstreamapi.ExceptionMessageType:
		return nil, awsRestjson1_deserializeEventStreamExceptionCodeInterpreterStreamOutput(msg)

	case eventstreamapi.ErrorMessageType:
		errorCode := "UnknownError"
		errorMessage := errorCode
		if header := msg.Headers.Get(eventstreamapi.ErrorCodeHeader); header != nil {
			errorCode = header.String()
		}
		if header := msg.Headers.Get(eventstreamapi.ErrorMessageHeader); header != nil {
			errorMessage = header.String()
		}
		return nil, &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}

	default:
		mc := msg.Clone()
		return nil, &UnknownEventMessageError{
			Type:    messageType.String(),
			Message: &mc,
		}

	}
}

func (r *codeInterpreterStreamOutputReader) ErrorSet() <-chan struct{} {
	return r.err.ErrorSet()
}

func (r *codeInterpreterStreamOutputReader) Close() error {
	r.closeOnce.Do(r.safeClose)
	return r.Err()
}

func (r *codeInterpreterStreamOutputReader) safeClose() {
	close(r.done)
	r.eventStream.Close()

}

func (r *codeInterpreterStreamOutputReader) Err() error {
	return r.err.Err()
}

func (r *codeInterpreterStreamOutputReader) Closed() <-chan struct{} {
	return r.done
}

type awsRestjson1_deserializeOpEventStreamInvokeCodeInterpreter struct {
	LogEventStreamWrites bool
	LogEventStreamReads  bool
}

func (*awsRestjson1_deserializeOpEventStreamInvokeCodeInterpreter) ID() string {
	return "OperationEventStreamDeserializer"
}

func (m *awsRestjson1_deserializeOpEventStreamInvokeCodeInterpreter) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	defer func() {
		if err == nil {
			return
		}
		m.closeResponseBody(out)
	}()

	logger := middleware.GetLogger(ctx)

	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", in.Request)
	}
	_ = request

	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	deserializeOutput, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type: %T", out.RawResponse)
	}
	_ = deserializeOutput

	output, ok := out.Result.(*InvokeCodeInterpreterOutput)
	if out.Result != nil && !ok {
		return out, metadata, fmt.Errorf("unexpected output result type: %T", out.Result)
	} else if out.Result == nil {
		output = &InvokeCodeInterpreterOutput{}
		out.Result = output
	}

	eventReader := newCodeInterpreterStreamOutputReader(
		deserializeOutput.Body,
		eventstream.NewDecoder(func(options *eventstream.DecoderOptions) {
			options.Logger = logger
			options.LogMessages = m.LogEventStreamReads

		}),
	)
	defer func() {
		if err == nil {
			return
		}
		_ = eventReader.Close()
	}()

	output.eventStream = NewInvokeCodeInterpreterEventStream(func(stream *InvokeCodeInterpreterEventStream) {
		stream.Reader = eventReader
	})

	go output.eventStream.waitStreamClose()

	return out, metadata, nil
}

func (*awsRestjson1_deserializeOpEventStreamInvokeCodeInterpreter) closeResponseBody(out middleware.DeserializeOutput) {
	if resp, ok := out.RawResponse.(*smithyhttp.Response); ok && resp != nil && resp.Body != nil {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}
}

func addEventStreamInvokeCodeInterpreterMiddleware(stack *middleware.Stack, options Options) error {
	if err := stack.Deserialize.Insert(&awsRestjson1_deserializeOpEventStreamInvokeCodeInterpreter{
		LogEventStreamWrites: options.ClientLogMode.IsRequestEventMessage(),
		LogEventStreamReads:  options.ClientLogMode.IsResponseEventMessage(),
	}, "OperationDeserializer", middleware.Before); err != nil {
		return err
	}
	return nil

}

// UnknownEventMessageError provides an error when a message is received from the stream,
// but the reader is unable to determine what kind of message it is.
type UnknownEventMessageError struct {
	Type    string
	Message *eventstream.Message
}

// Error retruns the error message string.
func (e *UnknownEventMessageError) Error() string {
	return "unknown event stream message type, " + e.Type
}

func setSafeEventStreamClientLogMode(o *Options, operation string) {
	switch operation {
	case "InvokeCodeInterpreter":
		toggleEventStreamClientLogMode(o, false, true)
		return

	default:
		return

	}
}
func toggleEventStreamClientLogMode(o *Options, request, response bool) {
	mode := o.ClientLogMode

	if request && mode.IsRequestWithBody() {
		mode.ClearRequestWithBody()
		mode |= aws.LogRequest
	}

	if response && mode.IsResponseWithBody() {
		mode.ClearResponseWithBody()
		mode |= aws.LogResponse
	}

	o.ClientLogMode = mode

}
