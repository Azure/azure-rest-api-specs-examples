package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateSoapPassThroughApiUsingWsdlImport.json
func ExampleAPIClient_BeginCreateOrUpdate_apiManagementCreateSoapPassThroughApiUsingWsdlImport() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "soapApi", armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			Path:        to.Ptr("currency"),
			Format:      to.Ptr(armapimanagement.ContentFormatWsdlLink),
			SoapAPIType: to.Ptr(armapimanagement.SoapAPITypeSoapPassThrough),
			Value:       to.Ptr("http://www.webservicex.net/CurrencyConvertor.asmx?WSDL"),
			WsdlSelector: &armapimanagement.APICreateOrUpdatePropertiesWsdlSelector{
				WsdlEndpointName: to.Ptr("CurrencyConvertorSoap"),
				WsdlServiceName:  to.Ptr("CurrencyConvertor"),
			},
		},
	}, &armapimanagement.APIClientBeginCreateOrUpdateOptions{IfMatch: nil})
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
	// res.APIContract = armapimanagement.APIContract{
	// 	Name: to.Ptr("soapApi"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/soapApi"),
	// 	Properties: &armapimanagement.APIContractProperties{
	// 		APIType: to.Ptr(armapimanagement.APITypeSoap),
	// 		APIRevision: to.Ptr("1"),
	// 		IsCurrent: to.Ptr(true),
	// 		SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
	// 			Header: to.Ptr("Ocp-Apim-Subscription-Key"),
	// 			Query: to.Ptr("subscription-key"),
	// 		},
	// 		Path: to.Ptr("currency"),
	// 		DisplayName: to.Ptr("CurrencyConvertor"),
	// 		Protocols: []*armapimanagement.Protocol{
	// 			to.Ptr(armapimanagement.ProtocolHTTPS)},
	// 			ServiceURL: to.Ptr("http://www.webservicex.net"),
	// 		},
	// 	}
}
