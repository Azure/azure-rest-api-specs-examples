package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/GetAmazonWebServicesS3ById.json
func ExampleDataConnectorsClient_Get_getAnAwsS3DataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "afef3743-0c88-469c-84ff-ca2e87dc1e48", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.AwsS3DataConnector{
	// 		Name: to.Ptr("afef3743-0c88-469c-84ff-ca2e87dc1e48"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/afef3743-0c88-469c-84ff-ca2e87dc1e48"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindAmazonWebServicesS3),
	// 		Properties: &armsecurityinsights.AwsS3DataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.AwsS3DataConnectorDataTypes{
	// 				Logs: &armsecurityinsights.AwsS3DataConnectorDataTypesLogs{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			DestinationTable: to.Ptr("AWSVPCFlow"),
	// 			RoleArn: to.Ptr("arn:aws:iam::072643944673:role/RoleName"),
	// 			SqsUrls: []*string{
	// 				to.Ptr("https://sqs.us-west-1.amazonaws.com/111111111111/sqsTestName")},
	// 			},
	// 		},
	// 		                        }
}
