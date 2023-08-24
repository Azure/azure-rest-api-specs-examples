package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateCertificate.json
func ExampleCertificateClient_CreateOrUpdate_apiManagementCreateCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateClient().CreateOrUpdate(ctx, "rg1", "apimService1", "tempcert", armapimanagement.CertificateCreateOrUpdateParameters{
		Properties: &armapimanagement.CertificateCreateOrUpdateProperties{
			Data:     to.Ptr("****************Base 64 Encoded Certificate *******************************"),
			Password: to.Ptr("****Certificate Password******"),
		},
	}, &armapimanagement.CertificateClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateContract = armapimanagement.CertificateContract{
	// 	Name: to.Ptr("tempcert"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/certificates"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/tempcert"),
	// 	Properties: &armapimanagement.CertificateContractProperties{
	// 		ExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-17T21:55:07+00:00"); return t}()),
	// 		Subject: to.Ptr("CN=contoso.com"),
	// 		Thumbprint: to.Ptr("*******************3"),
	// 	},
	// }
}
