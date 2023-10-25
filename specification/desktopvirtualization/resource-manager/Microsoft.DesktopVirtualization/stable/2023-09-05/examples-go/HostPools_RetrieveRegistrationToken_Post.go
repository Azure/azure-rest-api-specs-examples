package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/HostPools_RetrieveRegistrationToken_Post.json
func ExampleHostPoolsClient_RetrieveRegistrationToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHostPoolsClient().RetrieveRegistrationToken(ctx, "resourceGroup1", "hostPool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegistrationInfo = armdesktopvirtualization.RegistrationInfo{
	// 	ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.9571247Z"); return t}()),
	// 	RegistrationTokenOperation: to.Ptr(armdesktopvirtualization.RegistrationTokenOperationUpdate),
	// 	Token: to.Ptr("token"),
	// }
}
