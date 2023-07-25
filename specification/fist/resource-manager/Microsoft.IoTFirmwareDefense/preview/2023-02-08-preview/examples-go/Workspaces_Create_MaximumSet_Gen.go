package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_Create_MaximumSet_Gen.json
func ExampleWorkspacesClient_Create_workspacesCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Create(ctx, "rgworkspaces", "E___-3", armiotfirmwaredefense.Workspace{
		Location: to.Ptr("jjwbseilitjgdrhbvvkwviqj"),
		Tags: map[string]*string{
			"key450": to.Ptr("rzqqumbpfsbibnpirsm"),
		},
		Properties: &armiotfirmwaredefense.WorkspaceProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workspace = armiotfirmwaredefense.Workspace{
	// 	Name: to.Ptr("tbrqhnzpsatbrnhtj"),
	// 	Type: to.Ptr("angrpzpxmuppzqpzpljasjirao"),
	// 	ID: to.Ptr("uuuwv"),
	// 	SystemData: &armiotfirmwaredefense.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-19T05:35:25.545Z"); return t}()),
	// 		CreatedBy: to.Ptr("kmgstbzxtl"),
	// 		CreatedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-19T05:35:25.545Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("rfnhyhbyqzxnumjmjhwkhduztidk"),
	// 		LastModifiedByType: to.Ptr(armiotfirmwaredefense.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("jjwbseilitjgdrhbvvkwviqj"),
	// 	Tags: map[string]*string{
	// 		"key450": to.Ptr("rzqqumbpfsbibnpirsm"),
	// 	},
	// 	Properties: &armiotfirmwaredefense.WorkspaceProperties{
	// 		ProvisioningState: to.Ptr(armiotfirmwaredefense.ProvisioningStateSucceeded),
	// 	},
	// }
}
