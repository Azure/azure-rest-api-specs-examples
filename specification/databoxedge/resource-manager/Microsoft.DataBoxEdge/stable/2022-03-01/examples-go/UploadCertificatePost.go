package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/UploadCertificatePost.json
func ExampleDevicesClient_UploadCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDevicesClient().UploadCertificate(ctx, "testedgedevice", "GroupForEdgeAutomation", armdataboxedge.UploadCertificateRequest{
		Properties: &armdataboxedge.RawCertificateData{
			Certificate: to.Ptr("MIIC9DCCAdygAwIBAgIQWJae7GNjiI9Mcv/gJyrOPTANBgkqhkiG9w0BAQUFADASMRAwDgYDVQQDDAdXaW5kb3dzMB4XDTE4MTEyNzAwMTA0NVoXDTIxMTEyODAwMTA0NVowEjEQMA4GA1UEAwwHV2luZG93czCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKxkRExqxf0qH1avnyORptIbRC2yQwqe3EIbJ2FPKr5jtAppGeX/dGKrFSnX+7/0HFr77aJHafdpEAtOiLyJ4zCAVs0obZCCIq4qJdmjYUTU0UXH/w/YzXfQA0d9Zh9AN+NJBX9xj05NzgsT24fkgsK2v6mWJQXT7YcWAsl5sEYPnx1e+MrupNyVSL/RUJmrS+etJSysHtFeWRhsUhVAs1DD5ExJvBLU3WH0IsojEvpXcjrutB5/MDQNrd/StGI6WovoSSPH7FyT9tgERx+q+Yg3YUGzfaIPCctlrRGehcdtzdNoKd0rsX62yCq0U6POoSfwe22NJu41oAUMd7e6R8cCAwEAAaNGMEQwEwYDVR0lBAwwCgYIKwYBBQUHAwIwHQYDVR0OBBYEFDd0VxnS3LnMIfwc7xW4b4IZWG5GMA4GA1UdDwEB/wQEAwIFIDANBgkqhkiG9w0BAQUFAAOCAQEAPQRby2u9celvtvL/DLEb5Vt3/tPStRQC5MyTD62L5RT/q8E6EMCXVZNkXF5WlWucLJi/18tY+9PNgP9xWLJh7kpSWlWdi9KPtwMqKDlEH8L2TnQdjimt9XuiCrTnoFy/1X2BGLY/rCaUJNSd15QCkz2xeW+Z+YSk2GwAc/A/4YfNpqSIMfNuPrT76o02VdD9WmJUA3fS/HY0sU9qgQRS/3F5/0EPS+HYQ0SvXCK9tggcCd4O050ytNBMJC9qMOJ7yE0iOrFfOJSCfDAuPhn/rHFh79Kn1moF+/CE+nc0/2RPiLC8r54/rt5dYyyxJDfXg0a3VrrX39W69WZGW5OXiw=="),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UploadCertificateResponse = armdataboxedge.UploadCertificateResponse{
	// 	AADAuthority: to.Ptr("https://login.windows.net"),
	// 	AADTenantID: to.Ptr("100b019d-4626-4a9e-a83c-9cff8fe41909"),
	// 	AuthType: to.Ptr(armdataboxedge.AuthenticationTypeAzureActiveDirectory),
	// 	AzureManagementEndpointAudience: to.Ptr("https://wus-bvtgateway.ext.trafficmanager.net/"),
	// 	ResourceID: to.Ptr("392799901267771"),
	// 	ServicePrincipalClientID: to.Ptr("ffd8a688-82b1-4e5a-a4c1-7ede8c928e68"),
	// 	ServicePrincipalObjectID: to.Ptr("cef6d40c-24e6-4a6f-bb9e-3ec60b4adec4"),
	// }
}
