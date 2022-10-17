//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/stretchr/testify/suite"
)

type ActivitylogalertsTestSuite struct {
	suite.Suite

	ctx                  context.Context
	cred                 azcore.TokenCredential
	options              *arm.ClientOptions
	activityLogAlertName string
	location             string
	resourceGroupName    string
	subscriptionId       string
}

func (testsuite *ActivitylogalertsTestSuite) SetupSuite() {
	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/monitor/armmonitor/testdata")
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.activityLogAlertName = testutil.GenerateAlphaNumericID(testsuite.T(), "activitylogalertna", 6)
	testsuite.location = testutil.GetEnv("LOCATION", "eastus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
}

func (testsuite *ActivitylogalertsTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestActivitylogalertsTestSuite(t *testing.T) {
	suite.Run(t, new(ActivitylogalertsTestSuite))
}

// Microsoft.Insights/activityLogAlerts
func (testsuite *ActivitylogalertsTestSuite) TestActivitylogalert() {
	var err error
	// From step ActivityLogAlerts_Create
	activityLogAlertsClient, err := armmonitor.NewActivityLogAlertsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	_, err = activityLogAlertsClient.CreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.activityLogAlertName, armmonitor.ActivityLogAlertResource{
		Location: to.Ptr("Global"),
		Properties: &armmonitor.AlertRuleProperties{
			Condition: &armmonitor.AlertRuleAllOfCondition{
				AllOf: []*armmonitor.AlertRuleAnyOfOrLeafCondition{
					{
						Equals: to.Ptr("Administrative"),
						Field:  to.Ptr("category"),
					},
					{
						Equals: to.Ptr("Error"),
						Field:  to.Ptr("level"),
					}},
			},
			Scopes: []*string{
				to.Ptr("/subscriptions/" + testsuite.subscriptionId)},
		},
	}, nil)
	testsuite.Require().NoError(err)

	// From step ActivityLogAlerts_ListBySubscriptionId
	activityLogAlertsClientNewListBySubscriptionIDPager := activityLogAlertsClient.NewListBySubscriptionIDPager(nil)
	for activityLogAlertsClientNewListBySubscriptionIDPager.More() {
		_, err := activityLogAlertsClientNewListBySubscriptionIDPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step ActivityLogAlerts_ListByResourceGroup
	activityLogAlertsClientNewListByResourceGroupPager := activityLogAlertsClient.NewListByResourceGroupPager(testsuite.resourceGroupName, nil)
	for activityLogAlertsClientNewListByResourceGroupPager.More() {
		_, err := activityLogAlertsClientNewListByResourceGroupPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step ActivityLogAlerts_Get
	_, err = activityLogAlertsClient.Get(testsuite.ctx, testsuite.resourceGroupName, testsuite.activityLogAlertName, nil)
	testsuite.Require().NoError(err)

	// From step ActivityLogAlerts_Update
	_, err = activityLogAlertsClient.Update(testsuite.ctx, testsuite.resourceGroupName, testsuite.activityLogAlertName, armmonitor.AlertRulePatchObject{
		Properties: &armmonitor.AlertRulePatchProperties{
			Enabled: to.Ptr(false),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	testsuite.Require().NoError(err)

	// From step ActivityLogAlerts_Delete
	_, err = activityLogAlertsClient.Delete(testsuite.ctx, testsuite.resourceGroupName, testsuite.activityLogAlertName, nil)
	testsuite.Require().NoError(err)
}
