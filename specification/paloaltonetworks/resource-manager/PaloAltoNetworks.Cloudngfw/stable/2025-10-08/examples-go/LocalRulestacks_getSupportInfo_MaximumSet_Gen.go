package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/LocalRulestacks_getSupportInfo_MaximumSet_Gen.json
func ExampleLocalRulestacksClient_GetSupportInfo_localRulestacksGetSupportInfoMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("2bf4a339-294d-4c25-b0b2-ef649e9f5c27", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalRulestacksClient().GetSupportInfo(ctx, "rgopenapi", "lrs1", &armpanngfw.LocalRulestacksClientGetSupportInfoOptions{
		Email: to.Ptr("user1@domain.com")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.LocalRulestacksClientGetSupportInfoResponse{
	// 	SupportInfo: &armpanngfw.SupportInfo{
	// 		AccountID: to.Ptr("3cg5b439-294d-4c25-b0b2-ef649e0g6d38"),
	// 		AccountRegistered: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 		FreeTrial: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 		FreeTrialCreditLeft: to.Ptr[int32](10),
	// 		FreeTrialDaysLeft: to.Ptr[int32](1),
	// 		HelpURL: to.Ptr("https://ssopreview.paloaltonetworks.com/home/bookmark/0oa4ao61shG4rd3Ub1d7/2557"),
	// 		ProductSerial: to.Ptr("e22715cb-7e4e-4814-ad4f-ccd1417755d7"),
	// 		ProductSKU: to.Ptr("62f63e3c-bc5a-4d68-a8a1-fcba9f526c90"),
	// 		RegisterURL: to.Ptr("https://ssopreview.paloaltonetworks.com/home/bookmark/0oa4ao61shG4rd3Ub1d7/2557"),
	// 		SupportURL: to.Ptr("https://ssopreview.paloaltonetworks.com/home/bookmark/0oa4ao61shG4rd3Ub1d7/2557"),
	// 		UserDomainSupported: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 		UserRegistered: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 	},
	// }
}
