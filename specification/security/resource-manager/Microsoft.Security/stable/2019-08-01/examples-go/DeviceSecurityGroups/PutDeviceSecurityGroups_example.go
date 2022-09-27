package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/PutDeviceSecurityGroups_example.json
func ExampleDeviceSecurityGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewDeviceSecurityGroupsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub", "samplesecuritygroup", armsecurity.DeviceSecurityGroup{
		Properties: &armsecurity.DeviceSecurityGroupProperties{
			TimeWindowRules: []armsecurity.TimeWindowCustomAlertRuleClassification{
				&armsecurity.ActiveConnectionsNotInAllowedRange{
					IsEnabled:      to.Ptr(true),
					RuleType:       to.Ptr("ActiveConnectionsNotInAllowedRange"),
					MaxThreshold:   to.Ptr[int32](30),
					MinThreshold:   to.Ptr[int32](0),
					TimeWindowSize: to.Ptr("PT05M"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
