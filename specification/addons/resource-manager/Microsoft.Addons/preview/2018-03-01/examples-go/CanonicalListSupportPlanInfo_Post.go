package armaddons_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/addons/armaddons"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/addons/resource-manager/Microsoft.Addons/preview/2018-03-01/examples/CanonicalListSupportPlanInfo_Post.json
func ExampleSupportPlanTypesClient_ListInfo() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armaddons.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSupportPlanTypesClient().ListInfo(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CanonicalSupportPlanInfoDefinitionArray = []*armaddons.CanonicalSupportPlanInfoDefinition{
	// 	{
	// 		Enabled: to.Ptr(false),
	// 		OneTimeCharge: to.Ptr(armaddons.OneTimeChargeOnReenabled),
	// 		SupportPlanType: to.Ptr(armaddons.SupportPlanTypeStandard),
	// 	},
	// 	{
	// 		Enabled: to.Ptr(false),
	// 		OneTimeCharge: to.Ptr(armaddons.OneTimeChargeOnReenabled),
	// 		SupportPlanType: to.Ptr(armaddons.SupportPlanTypeAdvanced),
	// 	},
	// 	{
	// 		Enabled: to.Ptr(true),
	// 		OneTimeCharge: to.Ptr(armaddons.OneTimeChargeNo),
	// 		SupportPlanType: to.Ptr(armaddons.SupportPlanTypeEssential),
	// }}
}
