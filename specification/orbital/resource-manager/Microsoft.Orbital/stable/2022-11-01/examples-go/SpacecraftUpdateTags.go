package armorbital_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/SpacecraftUpdateTags.json
func ExampleSpacecraftsClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armorbital.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSpacecraftsClient().BeginUpdateTags(ctx, "contoso-Rgp", "CONTOSO_SAT", armorbital.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.Spacecraft = armorbital.Spacecraft{
	// 	Name: to.Ptr("CONTOSO_SAT"),
	// 	Type: to.Ptr("Microsoft.Orbital/spacecrafts"),
	// 	ID: to.Ptr("/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/spacecrafts/CONTOSO_SAT"),
	// 	Location: to.Ptr("eastus2"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armorbital.SpacecraftsProperties{
	// 		Links: []*armorbital.SpacecraftLink{
	// 			{
	// 				Name: to.Ptr("uplink_lhcp1"),
	// 				Authorizations: []*armorbital.AuthorizedGroundstation{
	// 					{
	// 						ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-06-02"); return t}()),
	// 						GroundStation: to.Ptr("EASTUS2_0"),
	// 				}},
	// 				BandwidthMHz: to.Ptr[float32](2),
	// 				CenterFrequencyMHz: to.Ptr[float32](2250),
	// 				Direction: to.Ptr(armorbital.DirectionUplink),
	// 				Polarization: to.Ptr(armorbital.PolarizationLHCP),
	// 			},
	// 			{
	// 				Name: to.Ptr("downlink_rhcp1"),
	// 				Authorizations: []*armorbital.AuthorizedGroundstation{
	// 					{
	// 						ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-06-02"); return t}()),
	// 						GroundStation: to.Ptr("EASTUS2_0"),
	// 				}},
	// 				BandwidthMHz: to.Ptr[float32](15),
	// 				CenterFrequencyMHz: to.Ptr[float32](8160),
	// 				Direction: to.Ptr(armorbital.DirectionDownlink),
	// 				Polarization: to.Ptr(armorbital.PolarizationRHCP),
	// 		}},
	// 		NoradID: to.Ptr("36411"),
	// 		ProvisioningState: to.Ptr(armorbital.SpacecraftsPropertiesProvisioningState("Succeeded")),
	// 		TitleLine: to.Ptr("CONTOSO_SAT"),
	// 		TleLine1: to.Ptr("1 27424U 02022A   22167.05119303  .00000638  00000+0  15103-3 0  9994"),
	// 		TleLine2: to.Ptr("2 27424  98.2477 108.9546 0000928  92.9194 327.0802 14.57300770 69982"),
	// 	},
	// }
}
