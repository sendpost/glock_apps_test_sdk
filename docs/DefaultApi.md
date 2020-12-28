# \DefaultApi

All URIs are relative to *https://spamtest.glockapps.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTestPost**](DefaultApi.md#CreateTestPost) | **Post** /CreateTest | 
[**GetProvidersGet**](DefaultApi.md#GetProvidersGet) | **Get** /GetProviders | 
[**GetTestResultGet**](DefaultApi.md#GetTestResultGet) | **Get** /GetTestResult | 


# **CreateTestPost**
> ModelsCreatetest CreateTestPost(ctx, apikey, groups, providers, v)


Create a test

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| the api key for you | 
  **groups** | **int32**| User groups | 
  **providers** | **string**| providers for api | 
  **v** | **int32**| version | 

### Return type

[**ModelsCreatetest**](models.createtest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvidersGet**
> ModelsGetprovider GetProvidersGet(ctx, apikey, v)


Get test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| the api key for you | 
  **v** | **int32**| version | 

### Return type

[**ModelsGetprovider**](models.getprovider.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestResultGet**
> ModelsTestresult GetTestResultGet(ctx, apikey, testId)


Get test result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apikey** | **string**| the api key for you | 
  **testId** | **string**| the testid for you | 

### Return type

[**ModelsTestresult**](models.testresult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

