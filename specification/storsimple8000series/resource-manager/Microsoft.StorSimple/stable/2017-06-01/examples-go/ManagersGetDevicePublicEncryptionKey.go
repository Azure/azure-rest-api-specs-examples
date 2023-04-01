package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersGetDevicePublicEncryptionKey.json
func ExampleManagersClient_GetDevicePublicEncryptionKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagersClient().GetDevicePublicEncryptionKey(ctx, "sca01forsdktest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PublicKey = armstorsimple8000series.PublicKey{
	// 	Key: to.Ptr("MIIDhjCCAm6gAwIBAgIQenen3LlBm4pFTfxko4pMMjANBgkqhkiG9w0BAQUFADB/MX0wewYDVQQDHnQAQwBCAF8AZgBhAGMAOAA2ADEAOAA3AC0AZgAzADYANgAtADQAYgA5ADUALQBiADQAMgBmAC0AZgBmADEAMAA3ADEANQAzAGQAOAA1AGUAXwA2ADMANgAzADMANwAzADcAMgAyADQANQAyADYANwAxADgAMDAeFw0xNzA2MTcwMDAwMDBaFw0yMDA2MjIxNDA3MDRaMH8xfTB7BgNVBAMedABDAEIAXwBmAGEAYwA4ADYAMQA4ADcALQBmADMANgA2AC0ANABiADkANQAtAGIANAAyAGYALQBmAGYAMQAwADcAMQA1ADMAZAA4ADUAZQBfADYAMwA2ADMAMwA3ADMANwAyADIANAA1ADIANgA3ADEAOAAwMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAz6r59tE8cxAf5g4cCs/JyaLKKpLf+5OEpSAsPnxHR5FBOaebsOV1raGDOmqIrc1dSOXArAyY6NcqizKkdslA0aNftKjya8Kks618Jf61mQVNHhq5fEyqTuVNSe3FZn1I2Xs2yS2NEt3YQl+GwisLgoT5ecC27NFZMEgjw2EOJwD9fjZG21CoCBCErEf+5j5HJkXk1jrn/o4wI7OnCGswPhlbGgDzl8eKcfSvMKpib50faCWdsvglC097CGtSmgu0a73IwGnac1eZCAnoVQmjaIFsdMeVNOqFot23tj1HT+Dwb9cECN8wKT7bkHCGSS84+veGiR0XEh3Op1iZ46zsIQIDAQABMA0GCSqGSIb3DQEBBQUAA4IBAQAfgsz7cH8NsiR5Ub0Pgx4T0cm8saMgJIfXZdMkCPTysdxU60EgmoDC/3KypFHAhhsBaFAvFrlXEybM4n4WLiDPeNM1Dr4xq56zzBhO0gO84Vp2PXk2VdybLfiobDyMb8cgAPXXcxzzb3jdRb1/bobiUuWBJvWSRHv5fON00c+BzloUzWJHY2BezNHBCof5awoMd+hzSfxX5WRFNW29rKgWW0h8A1cYNhaWGRxSA+U7YciYRyMDew8Nae1WCxVj9v/qsEXGd+jFDOckdJd5ySgXOQCPyG8j9/QM+Lsrsx29JnGsiiddaRfCtI4K8QdjkOjh0+RXMO3DeWhkFEmoOlNK"),
	// }
}
