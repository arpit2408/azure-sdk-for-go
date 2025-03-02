//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinkScopeOperationStatusClient contains the methods for the PrivateLinkScopeOperationStatus group.
// Don't use this type directly, use NewPrivateLinkScopeOperationStatusClient() instead.
type PrivateLinkScopeOperationStatusClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkScopeOperationStatusClient creates a new instance of PrivateLinkScopeOperationStatusClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkScopeOperationStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkScopeOperationStatusClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkScopeOperationStatusClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get the status of an azure asynchronous operation associated with a private link scope operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-07-01-preview
// asyncOperationID - The operation Id.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - PrivateLinkScopeOperationStatusClientGetOptions contains the optional parameters for the PrivateLinkScopeOperationStatusClient.Get
// method.
func (client *PrivateLinkScopeOperationStatusClient) Get(ctx context.Context, asyncOperationID string, resourceGroupName string, options *PrivateLinkScopeOperationStatusClientGetOptions) (PrivateLinkScopeOperationStatusClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, asyncOperationID, resourceGroupName, options)
	if err != nil {
		return PrivateLinkScopeOperationStatusClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkScopeOperationStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkScopeOperationStatusClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkScopeOperationStatusClient) getCreateRequest(ctx context.Context, asyncOperationID string, resourceGroupName string, options *PrivateLinkScopeOperationStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/privateLinkScopeOperationStatuses/{asyncOperationId}"
	if asyncOperationID == "" {
		return nil, errors.New("parameter asyncOperationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{asyncOperationId}", url.PathEscape(asyncOperationID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkScopeOperationStatusClient) getHandleResponse(resp *http.Response) (PrivateLinkScopeOperationStatusClientGetResponse, error) {
	result := PrivateLinkScopeOperationStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatus); err != nil {
		return PrivateLinkScopeOperationStatusClientGetResponse{}, err
	}
	return result, nil
}
