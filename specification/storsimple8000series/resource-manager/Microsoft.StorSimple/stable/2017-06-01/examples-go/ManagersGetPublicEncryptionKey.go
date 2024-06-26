package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersGetPublicEncryptionKey.json
func ExampleManagersClient_GetPublicEncryptionKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagersClient().GetPublicEncryptionKey(ctx, "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SymmetricEncryptedSecret = armstorsimple8000series.SymmetricEncryptedSecret{
	// 	EncryptionAlgorithm: to.Ptr(armstorsimple8000series.EncryptionAlgorithmAES256),
	// 	Value: to.Ptr("EAAAACcRUjoLMPhaI/fH4EFvdYv355vUfR6L+21trxRIl3toWpu91W804sPvAFQbCjwPAdG2epj8GYgXO44DZVHvf8x/X0f5cHhvLlcdx4yo7hdcEqoSpKwKytik8GfDTgsiqcEG4Fod05JF7E9CSAvFGOFf4hFfdQWjeiU3kTP5wCM4lJCCUWqI2uPSc5Izmg+XmumlXACavVKVtRzDEki4cbNCFc0tR1RAdMMDgziQmJFVlKNIvOOUSmurE1a5rXyUP8n3VvuCokIjlSSQSyySOSk/cWicAXjiBVHbJ6pFjJkDFLV/hTi/xdnrywVHJ7fDrdyBtMIRTZU6085lp/kWBkjVt5RNiPy9p/xHjOlGxa6DUSxYmUPZKzQytNQ49uLCVTH/JvUV0tWM6LU9bLvZxOYTJfD8hiHNaHk+XbrsX+AE4XrOj2cN/Ih1tULcX12cS21OLVYadWm9fkvU0U/csDu+r7PuwH/rBO1Bw/vmv2Kochj6h1akedro6R4+0txxu2t3HVh1a9Xc1/2xYNnSv/06/ekxEzxYwe7Sb6yAHZYZ4nCdS1FiGsBJ+2U1iSYgzDUDYcn9vwmFSJMdLlujwp5gNEWR7hZFrnmbPMtq+kBmuWA1mjJDTMZhefsJHLL3sEB1qqHTOaj8yAxM2oWYOvbc6/E/osy+qve0f1HAFWAsmgK2wOZ/yM9myAQDNFmVnxGibjNVb/svTIgfvYfSpDqWOKISDW1T4A8eHUk2Gz/8xnx+dSArjuVuLFk7u96tMbw4glEc74HNahtZbY8cvyFIStLugxrwlYmFlgN86tly9SZicd0x2cbsrIYySX4gvuSsgOwcQigoyVxFojE9u0rJa0gJYcTS4YbBjnd0+HGQJG604TKPb9hMmH7V1FbmH3bJc03PAOfPdNDgjA61uK7PgTEaamdIyoDg/FjBWpFiVJhdIu+m7Mm1gNf9GwXyv9XjQ9P42aAc/vUKzkvJESQ4uqkYFSlYTgyvMUyCpeRI9U2OD3iziNj5/Y1wr4eJ5BLePP6lVTR+UnWDjyu863d7N5grH7sWof1b+TZFb9Il6vMHA3BA9ujiFedQwqUdIdkdpKZuOBhkq+Ij9nMTQtIG1ODEpPetEXxaAoCKCS9DcBdmsN3MLvN46IIWN//vBmkABP3oUwnPoCgyx7P5V4NUpAJbm70mO1Mo6a1JEcFiqXJNZa9BO6v8e+9mRZokpDHOJCIJ0UwLpevpbKvBnLx9mcSNMN9JvcgGQOKvxt2/vBZdOv0DcO0aSYkV+bv+4ZAtLsBKOWeQTJ5CtSPs17m/y5a8+vxirvWxiGuMsNFBLHrD1aSAGBPd/5llsEde/TfIgWs+j8GW9Ep6E64LMjqALF0L5BZOSShleMbb2seTNi+d+0xOg5a2nuSRvPhcxFhRqPykmOQ/wAAMm1fE5K2l+jJfE3gZ5J3CdLXulmZHM1hV0oAGlqutqM6SCW7Czg813162G99tQb2l90MD1G01KkO7ppzn1Hy6MOazanOCzwRH6CMqZw09mEU9JozEKw=="),
	// 	ValueCertificateThumbprint: to.Ptr("D3411C80B443A01C51B8A8295A21EE82535972D0"),
	// }
}
