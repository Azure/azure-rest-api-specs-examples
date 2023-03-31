package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/GenerateCertificate.json
func ExampleDevicesClient_GenerateCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().GenerateCertificate(ctx, "testedgedevice", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GenerateCertResponse = armdataboxedge.GenerateCertResponse{
	// 	ExpiryTimeInUTC: to.Ptr("2020-11-22T05:06:20.000Z"),
	// 	PublicKey: to.Ptr("MIIEDjCCAnagAwIBAgIQEW4wrxj9+JdA4kFLDgegTTANBgkqhkiG9w0BAQUFADBDMUEwPwYDVQQDHjgAQwBCAF8AUABvAHIAdABhAGwAXwA2ADMANwA0ADEAMwA1ADkAMQA4ADAAMAA3ADAAOQAxADcANTAeFw0yMDExMTkwNDU2MjBaFw0yMDExMjIwNTA2MjBaMEMxQTA/BgNVBAMeOABDAEIAXwBQAG8AcgB0AGEAbABfADYAMwA3ADQAMQAzADUAOQAxADgAMAAwADcAMAA5ADEANwA1MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEAqG617KAUc4gOoht527o6FGE/c97W/wZIXqKbx4G8vifljXWw3NGDDuSQfidE3MqSMMJ/AAlDhBH/Deo/PAgUWuWB2WgM+KFWrDS0oDBvLVeIIeiMx4LOsbz4J8IbcKrGWf+ulQyum7dE/yBLzPXcPTjhJP0oMfhGOMlVVWbiKaP78/WZk2PmBQaVv7PxAQnJAPqETH2qAyuc9bd0TaMlIER56WhX9+nzsoI0u7RmBEs8+BaudVroiiK/GpEczy73TCRlpVXGChdBHg+NvYRYz91ltCV9Ijo1bdvcUjdgDJZ4Tz7G8XVCQmxbd9ml6OBLgxQUpd9HSGTcIfFiv+rTUwHR7YkY+1UGHsDPOYnBVDTXs3RFm+c5uNnCKvtyTExDJgabT+FTNb3eyD/BSYsy6b/YZBRqPiO8BRFUsIhWhXtPaUjWOEx6XMeZRnGeF4Hi8iH8TR79kEJhfYy3piSPAc3JrhgoglpyyfZOgyPOje+8tAWS9FHUNyVXKKeksDoZAgMBAAEwDQYJKoZIhvcNAQEFBQADggGBACpfQt34mgr+WlFtY0tqaPAoF6XnVzyKB6XsVAoj1uKFTSCQnr0mWvUMWKSRFTu1bV1VdDokV5I+h92kBcO4ddtHhWtkt16j9XFDsGdDvoUWJeqecZVr9fBPC69wF/9R/2E6iJporR46Pui7gcilYGqb3IjJlh8RT/XBPamQrfDbS7G8eTe91ST5cHr89R6EtDwfyhyTZr2sRrvQrxY2AjEmdH5Zcl9q2Xf8DholZhZ9vwnuYBFfl6VCXWaDAzQWiPdxo1wgXxEpNNDstGFWJKnsnz9UCohNAiPf47B+csPqxjcTLVMhcGLdU/lwpzzkEOciijpYILscvPoU3IM6DmFROgU2z6CForpgwHPDJktyWU1MC7lFdfrSzwu/5b3ZAEQhnJf7rvrybULoYBQWJEuj26Is1S6kXpNUaHAc5HIyIoc2V2eRjasbPIqsU0smjPLS/1egjHxtw9u+v4qDDvDuOjhFJtC3NkPl06TfNgO9+5LgiB1JppXnkLyCQ9RQgw=="),
	// }
}
