package armconfidentialledger_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7e295a19c5382a4df2f8101e545fed34186d83bf/specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2023-01-26-preview/examples/ManagedCCF_Update.json
func ExampleManagedCCFClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfidentialledger.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedCCFClient().BeginUpdate(ctx, "DummyResourceGroupName", "DummyMccfAppName", armconfidentialledger.ManagedCCF{
		Location: to.Ptr("EastUS"),
		Tags: map[string]*string{
			"additionalProps1": to.Ptr("additional properties"),
		},
		Properties: &armconfidentialledger.ManagedCCFProperties{
			DeploymentType: &armconfidentialledger.DeploymentType{
				AppSourceURI:    to.Ptr("https://myaccount.blob.core.windows.net/storage/mccfsource?sv=2022-02-11%st=2022-03-11"),
				LanguageRuntime: to.Ptr(armconfidentialledger.LanguageRuntimeCPP),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
