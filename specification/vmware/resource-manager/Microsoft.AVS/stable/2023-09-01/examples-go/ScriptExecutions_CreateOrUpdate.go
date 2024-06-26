package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f9d14b5db982b1d554651348adc9bef4b098bdb/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/ScriptExecutions_CreateOrUpdate.json
func ExampleScriptExecutionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScriptExecutionsClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", "addSsoServer", armavs.ScriptExecution{
		Properties: &armavs.ScriptExecutionProperties{
			HiddenParameters: []armavs.ScriptExecutionParameterClassification{
				&armavs.ScriptSecureStringExecutionParameter{
					Name:        to.Ptr("Password"),
					Type:        to.Ptr(armavs.ScriptExecutionParameterTypeSecureValue),
					SecureValue: to.Ptr("PlaceholderPassword"),
				}},
			Parameters: []armavs.ScriptExecutionParameterClassification{
				&armavs.ScriptStringExecutionParameter{
					Name:  to.Ptr("DomainName"),
					Type:  to.Ptr(armavs.ScriptExecutionParameterTypeValue),
					Value: to.Ptr("placeholderDomain.local"),
				},
				&armavs.ScriptStringExecutionParameter{
					Name:  to.Ptr("BaseUserDN"),
					Type:  to.Ptr(armavs.ScriptExecutionParameterTypeValue),
					Value: to.Ptr("DC=placeholder, DC=placeholder"),
				}},
			Retention:      to.Ptr("P0Y0M60DT0H60M60S"),
			ScriptCmdletID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/AVS.PowerCommands@1.0.0/scriptCmdlets/New-SsoExternalIdentitySource"),
			Timeout:        to.Ptr("P0Y0M0DT0H60M60S"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScriptExecution = armavs.ScriptExecution{
	// 	Name: to.Ptr("addSsoServer"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/scriptExecutions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptExecutions/addSsoServer"),
	// 	Properties: &armavs.ScriptExecutionProperties{
	// 		FailureReason: to.Ptr("vCenter failed to connect to the external server"),
	// 		Output: []*string{
	// 			to.Ptr("IdentitySource: placeholder.dc"),
	// 			to.Ptr("BaseDN='dc=placeholder, dc=local")},
	// 			Parameters: []armavs.ScriptExecutionParameterClassification{
	// 				&armavs.ScriptStringExecutionParameter{
	// 					Name: to.Ptr("DomainName"),
	// 					Type: to.Ptr(armavs.ScriptExecutionParameterTypeValue),
	// 					Value: to.Ptr("placeholderDomain.local"),
	// 				},
	// 				&armavs.ScriptStringExecutionParameter{
	// 					Name: to.Ptr("BaseUserDN"),
	// 					Type: to.Ptr(armavs.ScriptExecutionParameterTypeValue),
	// 					Value: to.Ptr("DC=placeholder, DC=placeholder"),
	// 			}},
	// 			ProvisioningState: to.Ptr(armavs.ScriptExecutionProvisioningStateSucceeded),
	// 			Retention: to.Ptr("P0Y0M60DT0H60M60S"),
	// 			ScriptCmdletID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/AVS.PowerCommands@1.0.0/scriptCmdlets/New-SsoExternalIdentitySource"),
	// 			Timeout: to.Ptr("P0Y0M0DT0H60M60S"),
	// 		},
	// 	}
}
