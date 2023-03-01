package armvoiceservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/voiceservices/armvoiceservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a7b696c2c73218fbca91c7c3bb625ee0f0bbea0/specification/voiceservices/resource-manager/Microsoft.VoiceServices/preview/2022-12-01-preview/examples/CommunicationsGateways_Get.json
func ExampleCommunicationsGatewaysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvoiceservices.NewCommunicationsGatewaysClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "testrg", "myname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommunicationsGateway = armvoiceservices.CommunicationsGateway{
	// 	Name: to.Ptr("myname"),
	// 	Type: to.Ptr("Microsoft.VoiceService/communicationsGateways"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.VoiceService/communicationsGateways/myname"),
	// 	Location: to.Ptr("useast"),
	// 	Properties: &armvoiceservices.CommunicationsGatewayProperties{
	// 		Codecs: []*armvoiceservices.TeamsCodecs{
	// 			to.Ptr(armvoiceservices.TeamsCodecsPCMA)},
	// 			Connectivity: to.Ptr(armvoiceservices.ConnectivityPublicAddress),
	// 			E911Type: to.Ptr(armvoiceservices.E911TypeStandard),
	// 			Platforms: []*armvoiceservices.CommunicationsPlatform{
	// 				to.Ptr(armvoiceservices.CommunicationsPlatformOperatorConnect)},
	// 				ServiceLocations: []*armvoiceservices.ServiceRegionProperties{
	// 					{
	// 						Name: to.Ptr("useast"),
	// 						PrimaryRegionProperties: &armvoiceservices.PrimaryRegionProperties{
	// 							OperatorAddresses: []*string{
	// 								to.Ptr("198.51.100.1")},
	// 							},
	// 						},
	// 						{
	// 							Name: to.Ptr("useast2"),
	// 							PrimaryRegionProperties: &armvoiceservices.PrimaryRegionProperties{
	// 								OperatorAddresses: []*string{
	// 									to.Ptr("198.51.100.2")},
	// 								},
	// 						}},
	// 						Status: to.Ptr(armvoiceservices.StatusChangePending),
	// 					},
	// 				}
}
