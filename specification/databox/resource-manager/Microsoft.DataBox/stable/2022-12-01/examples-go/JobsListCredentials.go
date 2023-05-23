package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsListCredentials.json
func ExampleJobsClient_NewListCredentialsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatabox.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListCredentialsPager("YourResourceGroupName", "TestJobName1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.UnencryptedCredentialsList = armdatabox.UnencryptedCredentialsList{
		// 	Value: []*armdatabox.UnencryptedCredentials{
		// 		{
		// 			JobName: to.Ptr("TestJobName1"),
		// 			JobSecrets: &armdatabox.JobSecrets{
		// 				DcAccessSecurityCode: &armdatabox.DcAccessSecurityCode{
		// 				},
		// 				JobSecretsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
		// 				PodSecrets: []*armdatabox.Secret{
		// 					{
		// 						AccountCredentialDetails: []*armdatabox.AccountCredentialDetails{
		// 							{
		// 								AccountConnectionString: to.Ptr(""),
		// 								AccountName: to.Ptr("YourStorageAccountName"),
		// 								DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
		// 								ShareCredentialDetails: []*armdatabox.ShareCredentialDetails{
		// 									{
		// 										Password: to.Ptr("<password>"),
		// 										ShareName: to.Ptr("testsharename_PageBlob"),
		// 										ShareType: to.Ptr(armdatabox.ShareDestinationFormatTypePageBlob),
		// 										SupportedAccessProtocols: []*armdatabox.AccessProtocol{
		// 											to.Ptr(armdatabox.AccessProtocolSMB)},
		// 											UserName: to.Ptr("testusername"),
		// 										},
		// 										{
		// 											Password: to.Ptr("<password>"),
		// 											ShareName: to.Ptr("testsharename_BlockBlob"),
		// 											ShareType: to.Ptr(armdatabox.ShareDestinationFormatTypeBlockBlob),
		// 											SupportedAccessProtocols: []*armdatabox.AccessProtocol{
		// 												to.Ptr(armdatabox.AccessProtocolSMB)},
		// 												UserName: to.Ptr("testusername"),
		// 											},
		// 											{
		// 												Password: to.Ptr("<password>"),
		// 												ShareName: to.Ptr("testsharename_AzFile"),
		// 												ShareType: to.Ptr(armdatabox.ShareDestinationFormatTypeAzureFile),
		// 												SupportedAccessProtocols: []*armdatabox.AccessProtocol{
		// 													to.Ptr(armdatabox.AccessProtocolSMB)},
		// 													UserName: to.Ptr("testusername"),
		// 											}},
		// 									}},
		// 									DevicePassword: to.Ptr("<devicePassword>"),
		// 									DeviceSerialNumber: to.Ptr("testserialnumber"),
		// 									EncodedValidationCertPubKey: to.Ptr("xxxxxxxxxx"),
		// 									NetworkConfigurations: []*armdatabox.ApplianceNetworkConfiguration{
		// 										{
		// 											Name: to.Ptr("DataPort3"),
		// 											MacAddress: to.Ptr("XXXXXXXXXXXX"),
		// 										},
		// 										{
		// 											Name: to.Ptr("DataPort1"),
		// 											MacAddress: to.Ptr("XXXXXXXXXXXX"),
		// 										},
		// 										{
		// 											Name: to.Ptr("DataPort2"),
		// 											MacAddress: to.Ptr("XXXXXXXXXXXX"),
		// 									}},
		// 							}},
		// 						},
		// 				}},
		// 			}
	}
}
