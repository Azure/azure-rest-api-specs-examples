package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/stable/2024-11-01/examples/RetrieveCertificateEmailHistory.json
func ExampleCertificateOrdersClient_RetrieveCertificateEmailHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCertificateOrdersClient().RetrieveCertificateEmailHistory(ctx, "testrg123", "SampleCertOrder", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CertificateEmailArray = []*armappservice.CertificateEmail{
	// 	{
	// 		EmailID: to.Ptr("customer@email.com"),
	// 		TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-12T23:40:25.000Z"); return t}()),
	// 	},
	// 	{
	// 		EmailID: to.Ptr("customer@email.com"),
	// 		TimeStamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-12T23:40:25.000Z"); return t}()),
	// }}
}
