/*
 * Tesera API
 *
 * API for Tesera
 *
 * API version: v1
 * Contact:
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type ThoughtsApiService service

/*
ThoughtsApiService List of thoughts
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ThoughtsApiGetOpts - Optional Parameters:
     * @param "Offset" (optional.Int32) -
     * @param "Limit" (optional.Int32) -
     * @param "IsCancellationRequested" (optional.Bool) -
     * @param "CanBeCanceled" (optional.Bool) -
     * @param "WaitHandleHandle" (optional.Interface of map[string]string) -
     * @param "WaitHandleSafeWaitHandleIsInvalid" (optional.Bool) -
     * @param "WaitHandleSafeWaitHandleIsClosed" (optional.Bool) -

@return []ThoughtInfoShort
*/

type ThoughtsApiGetOpts struct {
	Offset                            optional.Int32
	Limit                             optional.Int32
	IsCancellationRequested           optional.Bool
	CanBeCanceled                     optional.Bool
	WaitHandleHandle                  optional.Interface
	WaitHandleSafeWaitHandleIsInvalid optional.Bool
	WaitHandleSafeWaitHandleIsClosed  optional.Bool
}

func (a *ThoughtsApiService) Get(ctx context.Context, localVarOptionals *ThoughtsApiGetOpts) ([]ThoughtInfoShort, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []ThoughtInfoShort
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/thoughts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("Offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("Limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsCancellationRequested.IsSet() {
		localVarQueryParams.Add("IsCancellationRequested", parameterToString(localVarOptionals.IsCancellationRequested.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CanBeCanceled.IsSet() {
		localVarQueryParams.Add("CanBeCanceled", parameterToString(localVarOptionals.CanBeCanceled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleHandle.IsSet() {
		localVarQueryParams.Add("WaitHandle.Handle", parameterToString(localVarOptionals.WaitHandleHandle.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsInvalid", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsClosed.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsClosed", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsClosed.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v []ThoughtInfoShort
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ThoughtsApiService List of thoughts
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param version
 * @param optional nil or *ThoughtsApiGet_1Opts - Optional Parameters:
     * @param "Offset" (optional.Int32) -
     * @param "Limit" (optional.Int32) -
     * @param "IsCancellationRequested" (optional.Bool) -
     * @param "CanBeCanceled" (optional.Bool) -
     * @param "WaitHandleHandle" (optional.Interface of map[string]string) -
     * @param "WaitHandleSafeWaitHandleIsInvalid" (optional.Bool) -
     * @param "WaitHandleSafeWaitHandleIsClosed" (optional.Bool) -

@return []ThoughtInfoShort
*/

type ThoughtsApiGet_1Opts struct {
	Offset                            optional.Int32
	Limit                             optional.Int32
	IsCancellationRequested           optional.Bool
	CanBeCanceled                     optional.Bool
	WaitHandleHandle                  optional.Interface
	WaitHandleSafeWaitHandleIsInvalid optional.Bool
	WaitHandleSafeWaitHandleIsClosed  optional.Bool
}

func (a *ThoughtsApiService) Get_1(ctx context.Context, version string, localVarOptionals *ThoughtsApiGet_1Opts) ([]ThoughtInfoShort, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []ThoughtInfoShort
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v{version}/thoughts"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("Offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("Limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsCancellationRequested.IsSet() {
		localVarQueryParams.Add("IsCancellationRequested", parameterToString(localVarOptionals.IsCancellationRequested.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CanBeCanceled.IsSet() {
		localVarQueryParams.Add("CanBeCanceled", parameterToString(localVarOptionals.CanBeCanceled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleHandle.IsSet() {
		localVarQueryParams.Add("WaitHandle.Handle", parameterToString(localVarOptionals.WaitHandleHandle.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsInvalid", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsClosed.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsClosed", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsClosed.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v []ThoughtInfoShort
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ThoughtsApiService Specific thought by alias
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param alias Thought alias
 * @param optional nil or *ThoughtsApiGet_2Opts - Optional Parameters:
     * @param "IsCancellationRequested" (optional.Bool) -
     * @param "CanBeCanceled" (optional.Bool) -
     * @param "WaitHandleHandle" (optional.Interface of map[string]string) -
     * @param "WaitHandleSafeWaitHandleIsInvalid" (optional.Bool) -
     * @param "WaitHandleSafeWaitHandleIsClosed" (optional.Bool) -

@return ThoughtInfoResponse
*/

type ThoughtsApiGet_2Opts struct {
	IsCancellationRequested           optional.Bool
	CanBeCanceled                     optional.Bool
	WaitHandleHandle                  optional.Interface
	WaitHandleSafeWaitHandleIsInvalid optional.Bool
	WaitHandleSafeWaitHandleIsClosed  optional.Bool
}

func (a *ThoughtsApiService) Get_2(ctx context.Context, alias string, localVarOptionals *ThoughtsApiGet_2Opts) (ThoughtInfoResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ThoughtInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/thoughts/{alias}"
	localVarPath = strings.Replace(localVarPath, "{"+"alias"+"}", fmt.Sprintf("%v", alias), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IsCancellationRequested.IsSet() {
		localVarQueryParams.Add("IsCancellationRequested", parameterToString(localVarOptionals.IsCancellationRequested.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CanBeCanceled.IsSet() {
		localVarQueryParams.Add("CanBeCanceled", parameterToString(localVarOptionals.CanBeCanceled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleHandle.IsSet() {
		localVarQueryParams.Add("WaitHandle.Handle", parameterToString(localVarOptionals.WaitHandleHandle.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsInvalid", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsClosed.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsClosed", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsClosed.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v ThoughtInfoResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ThoughtsApiService Specific thought by alias
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param alias Thought alias
 * @param version
 * @param optional nil or *ThoughtsApiGet_3Opts - Optional Parameters:
     * @param "IsCancellationRequested" (optional.Bool) -
     * @param "CanBeCanceled" (optional.Bool) -
     * @param "WaitHandleHandle" (optional.Interface of map[string]string) -
     * @param "WaitHandleSafeWaitHandleIsInvalid" (optional.Bool) -
     * @param "WaitHandleSafeWaitHandleIsClosed" (optional.Bool) -

@return ThoughtInfoResponse
*/

type ThoughtsApiGet_3Opts struct {
	IsCancellationRequested           optional.Bool
	CanBeCanceled                     optional.Bool
	WaitHandleHandle                  optional.Interface
	WaitHandleSafeWaitHandleIsInvalid optional.Bool
	WaitHandleSafeWaitHandleIsClosed  optional.Bool
}

func (a *ThoughtsApiService) Get_3(ctx context.Context, alias string, version string, localVarOptionals *ThoughtsApiGet_3Opts) (ThoughtInfoResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ThoughtInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v{version}/thoughts/{alias}"
	localVarPath = strings.Replace(localVarPath, "{"+"alias"+"}", fmt.Sprintf("%v", alias), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IsCancellationRequested.IsSet() {
		localVarQueryParams.Add("IsCancellationRequested", parameterToString(localVarOptionals.IsCancellationRequested.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CanBeCanceled.IsSet() {
		localVarQueryParams.Add("CanBeCanceled", parameterToString(localVarOptionals.CanBeCanceled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleHandle.IsSet() {
		localVarQueryParams.Add("WaitHandle.Handle", parameterToString(localVarOptionals.WaitHandleHandle.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsInvalid", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsClosed.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsClosed", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsClosed.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v ThoughtInfoResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
