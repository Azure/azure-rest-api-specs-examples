package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/SecurityConnectors/PutSecurityConnector_example.json
func ExampleConnectorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewConnectorsClient("a5caac9c-5c04-49af-b3d0-e204f40345d5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"exampleResourceGroup",
		"exampleSecurityConnectorName",
		armsecurity.Connector{
			Location: to.Ptr("Central US"),
			Etag:     to.Ptr("etag value (must be supplied for update)"),
			Tags:     map[string]*string{},
			Properties: &armsecurity.ConnectorProperties{
				CloudName:           to.Ptr(armsecurity.CloudNameAWS),
				HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
				Offerings: []armsecurity.CloudOfferingClassification{
					&armsecurity.CspmMonitorAwsOffering{
						OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
						NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
							CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
