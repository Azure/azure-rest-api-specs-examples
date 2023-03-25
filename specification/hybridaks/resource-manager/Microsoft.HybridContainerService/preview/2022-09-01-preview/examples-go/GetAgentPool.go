package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/GetAgentPool.json
func ExampleAgentPoolClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentPoolClient().Get(ctx, "test-arcappliance-resgrp", "test-hybridakscluster", "test-hybridaksnodepool", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPool = armhybridcontainerservice.AgentPool{
	// 	Name: to.Ptr("test-hybridaksnodepool"),
	// 	Type: to.Ptr("microsoft.hybridcontainerservice/provisionedclusters/agentpools"),
	// 	ID: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/provisionedClusters/test-hybridakscluster/agentPools/test-hybridaksnodepool"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armhybridcontainerservice.AgentPoolProperties{
	// 		Count: to.Ptr[int32](1),
	// 		OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
	// 		VMSize: to.Ptr("Standard_A4_v2"),
	// 		ProvisioningState: to.Ptr(armhybridcontainerservice.AgentPoolProvisioningStateSucceeded),
	// 	},
	// }
}
