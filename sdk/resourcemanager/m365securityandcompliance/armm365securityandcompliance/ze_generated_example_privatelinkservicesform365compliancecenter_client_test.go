//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceGet.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"rg1",
		"service1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceCreate.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"service1",
		armm365securityandcompliance.PrivateLinkServicesForM365ComplianceCenterDescription{
			Identity: &armm365securityandcompliance.ServicesResourceIdentity{
				Type: to.Ptr(armm365securityandcompliance.ManagedServiceIdentityTypeSystemAssigned),
			},
			Kind:     to.Ptr(armm365securityandcompliance.KindFhirR4),
			Location: to.Ptr("westus2"),
			Tags:     map[string]*string{},
			Properties: &armm365securityandcompliance.ServicesProperties{
				AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
					{
						ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
					},
					{
						ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
					}},
				AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
					Audience:          to.Ptr("https://azurehealthcareapis.com"),
					Authority:         to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
					SmartProxyEnabled: to.Ptr(true),
				},
				CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
					AllowCredentials: to.Ptr(false),
					Headers: []*string{
						to.Ptr("*")},
					MaxAge: to.Ptr[int64](1440),
					Methods: []*string{
						to.Ptr("DELETE"),
						to.Ptr("GET"),
						to.Ptr("OPTIONS"),
						to.Ptr("PATCH"),
						to.Ptr("POST"),
						to.Ptr("PUT")},
					Origins: []*string{
						to.Ptr("*")},
				},
				CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
					KeyVaultKeyURI:  to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
					OfferThroughput: to.Ptr[int64](1000),
				},
				ExportConfiguration: &armm365securityandcompliance.ServiceExportConfigurationInfo{
					StorageAccountName: to.Ptr("existingStorageAccount"),
				},
				PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{},
				PublicNetworkAccess:        to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServicePatch.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"rg1",
		"service1",
		armm365securityandcompliance.ServicesPatchDescription{
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceDelete.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"rg1",
		"service1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceList.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/ComplianceCenterServiceListByResourceGroup.json
func ExamplePrivateLinkServicesForM365ComplianceCenterClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armm365securityandcompliance.NewPrivateLinkServicesForM365ComplianceCenterClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("rgname",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
