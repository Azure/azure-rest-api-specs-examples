package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ApplicationGatewayAvailableSslOptionsPredefinedPolicyGet.json
func ExampleApplicationGatewaysClient_GetSSLPredefinedPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationGatewaysClient().GetSSLPredefinedPolicy(ctx, "AppGwSslPolicy20150501", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationGatewaySSLPredefinedPolicy = armnetwork.ApplicationGatewaySSLPredefinedPolicy{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups//providers/Microsoft.Network/ApplicationGatewayAvailableSslOptions/default/ApplicationGatewaySslPredefinedPolicy/AppGwSslPolicy20150501"),
	// 	Name: to.Ptr("AppGwSslPolicy20150501"),
	// 	Properties: &armnetwork.ApplicationGatewaySSLPredefinedPolicyPropertiesFormat{
	// 		CipherSuites: []*armnetwork.ApplicationGatewaySSLCipherSuite{
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHERSAWITHAES256GCMSHA384),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHERSAWITHAES128GCMSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHERSAWITHAES256CBCSHA384),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHERSAWITHAES128CBCSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHERSAWITHAES256CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHERSAWITHAES128CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSDHERSAWITHAES256GCMSHA384),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSDHERSAWITHAES128GCMSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSDHERSAWITHAES256CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSDHERSAWITHAES128CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSRSAWITHAES256GCMSHA384),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSRSAWITHAES128GCMSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSRSAWITHAES256CBCSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSRSAWITHAES128CBCSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSRSAWITHAES256CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSRSAWITHAES128CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHEECDSAWITHAES256GCMSHA384),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHEECDSAWITHAES128GCMSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHEECDSAWITHAES256CBCSHA384),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHEECDSAWITHAES128CBCSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHEECDSAWITHAES256CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSECDHEECDSAWITHAES128CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSDHEDSSWITHAES256CBCSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSDHEDSSWITHAES128CBCSHA256),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSDHEDSSWITHAES256CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSDHEDSSWITHAES128CBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSRSAWITH3DESEDECBCSHA),
	// 			to.Ptr(armnetwork.ApplicationGatewaySSLCipherSuiteTLSDHEDSSWITH3DESEDECBCSHA)},
	// 			MinProtocolVersion: to.Ptr(armnetwork.ApplicationGatewaySSLProtocolTLSv10),
	// 		},
	// 	}
}
