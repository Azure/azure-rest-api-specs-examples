package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementCreateSoapPassThroughApiUsingWsdlImport.json
func ExampleAPIClient_BeginCreateOrUpdate_apiManagementCreateSoapPassThroughApiUsingWsdlImport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "soapApi", armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			Format:      to.Ptr(armapimanagement.ContentFormatWsdlLink),
			Path:        to.Ptr("currency"),
			SoapAPIType: to.Ptr(armapimanagement.SoapAPITypeSoapPassThrough),
			Value:       to.Ptr("http://www.webservicex.net/CurrencyConvertor.asmx?WSDL"),
			WsdlSelector: &armapimanagement.APICreateOrUpdatePropertiesWsdlSelector{
				WsdlEndpointName: to.Ptr("CurrencyConvertorSoap"),
				WsdlServiceName:  to.Ptr("CurrencyConvertor"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.APIClientCreateOrUpdateResponse{
	// 	APIContract: armapimanagement.APIContract{
	// 		Name: to.Ptr("soapApi"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/apis"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/soapApi"),
	// 		Properties: &armapimanagement.APIContractProperties{
	// 			APIType: to.Ptr(armapimanagement.APITypeSoap),
	// 			Path: to.Ptr("currency"),
	// 			APIRevision: to.Ptr("1"),
	// 			DisplayName: to.Ptr("CurrencyConvertor"),
	// 			IsCurrent: to.Ptr(true),
	// 			Protocols: []*armapimanagement.Protocol{
	// 				to.Ptr(armapimanagement.ProtocolHTTPS),
	// 			},
	// 			ProvisioningState: to.Ptr("InProgress"),
	// 			ServiceURL: to.Ptr("http://www.webservicex.net"),
	// 			SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 				Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 				Query: to.Ptr("subscription-key"),
	// 			},
	// 		},
	// 	},
	// }
}
