package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2024-03-01-preview/examples/SecurityConnectors/GetSecurityConnectorSingleResource_example.json
func ExampleConnectorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectorsClient().Get(ctx, "exampleResourceGroup", "exampleSecurityConnectorName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Connector = armsecurity.Connector{
	// 	Location: to.Ptr("Central US"),
	// 	Etag: to.Ptr("etag value"),
	// 	Kind: to.Ptr(""),
	// 	Name: to.Ptr("exampleSecurityConnectorName"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors"),
	// 	ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorName"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecurity.ConnectorProperties{
	// 		EnvironmentData: &armsecurity.AwsEnvironmentData{
	// 			EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
	// 			ScanInterval: to.Ptr[int64](4),
	// 		},
	// 		EnvironmentName: to.Ptr(armsecurity.CloudNameAWS),
	// 		HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
	// 		Offerings: []armsecurity.CloudOfferingClassification{
	// 			&armsecurity.CspmMonitorAwsOffering{
	// 				OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
	// 				NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
	// 					CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
	// 				},
	// 		}},
	// 	},
	// 	SystemData: &armsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 	},
	// }
}
