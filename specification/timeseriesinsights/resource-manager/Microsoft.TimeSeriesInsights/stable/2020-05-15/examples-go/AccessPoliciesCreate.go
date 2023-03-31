package armtimeseriesinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/stable/2020-05-15/examples/AccessPoliciesCreate.json
func ExampleAccessPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armtimeseriesinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessPoliciesClient().CreateOrUpdate(ctx, "rg1", "env1", "ap1", armtimeseriesinsights.AccessPolicyCreateOrUpdateParameters{
		Properties: &armtimeseriesinsights.AccessPolicyResourceProperties{
			Description:       to.Ptr("some description"),
			PrincipalObjectID: to.Ptr("aGuid"),
			Roles: []*armtimeseriesinsights.AccessPolicyRole{
				to.Ptr(armtimeseriesinsights.AccessPolicyRoleReader)},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessPolicyResource = armtimeseriesinsights.AccessPolicyResource{
	// 	Name: to.Ptr("ap1"),
	// 	Type: to.Ptr("Microsoft.TimeSeriesInsights/Environments/AccessPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.TimeSeriesInsights/Environments/env1/accessPolicies/ap1"),
	// 	Properties: &armtimeseriesinsights.AccessPolicyResourceProperties{
	// 		Description: to.Ptr("some description"),
	// 		PrincipalObjectID: to.Ptr("aGuid"),
	// 		Roles: []*armtimeseriesinsights.AccessPolicyRole{
	// 			to.Ptr(armtimeseriesinsights.AccessPolicyRoleReader)},
	// 		},
	// 	}
}
