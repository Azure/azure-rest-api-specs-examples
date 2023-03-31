package armdatabox_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databox/resource-manager/Microsoft.DataBox/stable/2022-02-01/examples/JobsListCredentials.json
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
	pager := clientFactory.NewJobsClient().NewListCredentialsPager("bvttoolrg6", "TJ-636646322037905056", nil)
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
		// 			JobName: to.Ptr("TJ-636646322037905056"),
		// 			JobSecrets: &armdatabox.JobSecrets{
		// 				DcAccessSecurityCode: &armdatabox.DcAccessSecurityCode{
		// 				},
		// 				JobSecretsType: to.Ptr(armdatabox.ClassDiscriminatorDataBox),
		// 				PodSecrets: []*armdatabox.Secret{
		// 					{
		// 						AccountCredentialDetails: []*armdatabox.AccountCredentialDetails{
		// 							{
		// 								AccountConnectionString: to.Ptr(""),
		// 								AccountName: to.Ptr("databoxbvttestaccount"),
		// 								DataAccountType: to.Ptr(armdatabox.DataAccountTypeStorageAccount),
		// 								ShareCredentialDetails: []*armdatabox.ShareCredentialDetails{
		// 									{
		// 										Password: to.Ptr("<password>"),
		// 										ShareName: to.Ptr("databoxbvttestaccount_PageBlob"),
		// 										ShareType: to.Ptr(armdatabox.ShareDestinationFormatTypePageBlob),
		// 										SupportedAccessProtocols: []*armdatabox.AccessProtocol{
		// 											to.Ptr(armdatabox.AccessProtocolSMB)},
		// 											UserName: to.Ptr("databoxbvttestac_903"),
		// 										},
		// 										{
		// 											Password: to.Ptr("<password>"),
		// 											ShareName: to.Ptr("databoxbvttestaccount_BlockBlob"),
		// 											ShareType: to.Ptr(armdatabox.ShareDestinationFormatTypeBlockBlob),
		// 											SupportedAccessProtocols: []*armdatabox.AccessProtocol{
		// 												to.Ptr(armdatabox.AccessProtocolSMB)},
		// 												UserName: to.Ptr("databoxbvttestac_903"),
		// 											},
		// 											{
		// 												Password: to.Ptr("<password>"),
		// 												ShareName: to.Ptr("databoxbvttestaccount_AzFile"),
		// 												ShareType: to.Ptr(armdatabox.ShareDestinationFormatTypeAzureFile),
		// 												SupportedAccessProtocols: []*armdatabox.AccessProtocol{
		// 													to.Ptr(armdatabox.AccessProtocolSMB)},
		// 													UserName: to.Ptr("databoxbvttestac_903"),
		// 											}},
		// 									}},
		// 									DevicePassword: to.Ptr("<devicePassword>"),
		// 									DeviceSerialNumber: to.Ptr("testimolapod-3ecc44ce"),
		// 									EncodedValidationCertPubKey: to.Ptr("5CYoAoVKEBa4WgPVis8keX94w30pon4jGMADSqcdE/NlHLChj6Cmhbl4q9QOFKSB/US4AwhS7zY1QS3YMDrkAPfOy7Hi6kWMBpJWZidTq3oXX8FAQjg+IqQESti/2jvAlcDpO2453rgd7Yb6XZ43P8MMTpTjcarI0ImCf//eITQWnFa3AzfIJ9C+hxCCaA7HTYhwQEPUBMwyQJsI6v6WuQysROtlBgx1YtbWFhDVbcqYRSLIbaj+RdzlvxvDJSo70kv+8em5upuDTpVE7xP+WePLlARdSPNwwfRzHnvCUqC2UqXHpRUhQlYnMqAJEcjjroRnyIGumPmmQ8O155X8aw=="),
		// 									NetworkConfigurations: []*armdatabox.ApplianceNetworkConfiguration{
		// 										{
		// 											Name: to.Ptr("DataPort3"),
		// 											MacAddress: to.Ptr("D05099C1F439"),
		// 										},
		// 										{
		// 											Name: to.Ptr("DataPort1"),
		// 											MacAddress: to.Ptr("EC0D9A21A6C0"),
		// 										},
		// 										{
		// 											Name: to.Ptr("DataPort2"),
		// 											MacAddress: to.Ptr("EC0D9A21A6C1"),
		// 									}},
		// 							}},
		// 						},
		// 				}},
		// 			}
	}
}
