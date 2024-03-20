package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01a71545e82bb98b8137d3038150c436d46a59ed/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_GetDataPlaneAccess.json
func ExampleFactoriesClient_GetDataPlaneAccess() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFactoriesClient().GetDataPlaneAccess(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.UserAccessPolicy{
		AccessResourcePath: to.Ptr(""),
		ExpireTime:         to.Ptr("2018-11-10T09:46:20.2659347Z"),
		Permissions:        to.Ptr("r"),
		ProfileName:        to.Ptr("DefaultProfile"),
		StartTime:          to.Ptr("2018-11-10T02:46:20.2659347Z"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessPolicyResponse = armdatafactory.AccessPolicyResponse{
	// 	AccessToken: to.Ptr("**********"),
	// 	DataPlaneURL: to.Ptr("https://rpeastus.svc.datafactory.azure.com:4433"),
	// 	Policy: &armdatafactory.UserAccessPolicy{
	// 		AccessResourcePath: to.Ptr(""),
	// 		ExpireTime: to.Ptr("2018-11-10T09:46:20.2659347Z"),
	// 		Permissions: to.Ptr("r"),
	// 		ProfileName: to.Ptr("DefaultProfile"),
	// 		StartTime: to.Ptr("2018-11-10T02:46:20.2659347Z"),
	// 	},
	// }
}
