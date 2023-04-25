package armappplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-01-01-preview/examples/Gateways_CreateOrUpdate.json
func ExampleGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappplatform.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGatewaysClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myservice", "default", armappplatform.GatewayResource{
		Properties: &armappplatform.GatewayProperties{
			Public: to.Ptr(true),
			ResourceRequests: &armappplatform.GatewayResourceRequests{
				CPU:    to.Ptr("1"),
				Memory: to.Ptr("1G"),
			},
		},
		SKU: &armappplatform.SKU{
			Name:     to.Ptr("E0"),
			Capacity: to.Ptr[int32](2),
			Tier:     to.Ptr("Enterprise"),
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
	// res.GatewayResource = armappplatform.GatewayResource{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.AppPlatform/Spring/gateways"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/gateways/default"),
	// 	SystemData: &armappplatform.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:16:03.944Z"); return t}()),
	// 		CreatedBy: to.Ptr("sample-user"),
	// 		CreatedByType: to.Ptr(armappplatform.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-11T03:17:03.944Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sample-user"),
	// 		LastModifiedByType: to.Ptr(armappplatform.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armappplatform.GatewayProperties{
	// 		Instances: []*armappplatform.GatewayInstance{
	// 			{
	// 				Name: to.Ptr("instance1"),
	// 				Status: to.Ptr("Running"),
	// 			},
	// 			{
	// 				Name: to.Ptr("instance2"),
	// 				Status: to.Ptr("Running"),
	// 		}},
	// 		OperatorProperties: &armappplatform.GatewayOperatorProperties{
	// 			Instances: []*armappplatform.GatewayInstance{
	// 				{
	// 					Name: to.Ptr("instance1"),
	// 					Status: to.Ptr("Running"),
	// 				},
	// 				{
	// 					Name: to.Ptr("instance2"),
	// 					Status: to.Ptr("Running"),
	// 			}},
	// 			ResourceRequests: &armappplatform.GatewayOperatorResourceRequests{
	// 				CPU: to.Ptr("1"),
	// 				InstanceCount: to.Ptr[int32](2),
	// 				Memory: to.Ptr("1G"),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armappplatform.GatewayProvisioningStateSucceeded),
	// 		Public: to.Ptr(true),
	// 		ResourceRequests: &armappplatform.GatewayResourceRequests{
	// 			CPU: to.Ptr("1"),
	// 			Memory: to.Ptr("1G"),
	// 		},
	// 		URL: to.Ptr("test-url"),
	// 	},
	// 	SKU: &armappplatform.SKU{
	// 		Name: to.Ptr("E0"),
	// 		Capacity: to.Ptr[int32](2),
	// 		Tier: to.Ptr("Enterprise"),
	// 	},
	// }
}
