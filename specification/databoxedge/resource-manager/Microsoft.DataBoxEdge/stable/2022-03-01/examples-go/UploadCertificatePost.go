package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/UploadCertificatePost.json
func ExampleDevicesClient_UploadCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataboxedge.NewDevicesClient("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UploadCertificate(ctx,
		"testedgedevice",
		"GroupForEdgeAutomation",
		armdataboxedge.UploadCertificateRequest{
			Properties: &armdataboxedge.RawCertificateData{
				Certificate: to.Ptr("MIIC9DCCAdygAwIBAgIQWJae7GNjiI9Mcv/gJyrOPTANBgkqhkiG9w0BAQUFADASMRAwDgYDVQQDDAdXaW5kb3dzMB4XDTE4MTEyNzAwMTA0NVoXDTIxMTEyODAwMTA0NVowEjEQMA4GA1UEAwwHV2luZG93czCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKxkRExqxf0qH1avnyORptIbRC2yQwqe3EIbJ2FPKr5jtAppGeX/dGKrFSnX+7/0HFr77aJHafdpEAtOiLyJ4zCAVs0obZCCIq4qJdmjYUTU0UXH/w/YzXfQA0d9Zh9AN+NJBX9xj05NzgsT24fkgsK2v6mWJQXT7YcWAsl5sEYPnx1e+MrupNyVSL/RUJmrS+etJSysHtFeWRhsUhVAs1DD5ExJvBLU3WH0IsojEvpXcjrutB5/MDQNrd/StGI6WovoSSPH7FyT9tgERx+q+Yg3YUGzfaIPCctlrRGehcdtzdNoKd0rsX62yCq0U6POoSfwe22NJu41oAUMd7e6R8cCAwEAAaNGMEQwEwYDVR0lBAwwCgYIKwYBBQUHAwIwHQYDVR0OBBYEFDd0VxnS3LnMIfwc7xW4b4IZWG5GMA4GA1UdDwEB/wQEAwIFIDANBgkqhkiG9w0BAQUFAAOCAQEAPQRby2u9celvtvL/DLEb5Vt3/tPStRQC5MyTD62L5RT/q8E6EMCXVZNkXF5WlWucLJi/18tY+9PNgP9xWLJh7kpSWlWdi9KPtwMqKDlEH8L2TnQdjimt9XuiCrTnoFy/1X2BGLY/rCaUJNSd15QCkz2xeW+Z+YSk2GwAc/A/4YfNpqSIMfNuPrT76o02VdD9WmJUA3fS/HY0sU9qgQRS/3F5/0EPS+HYQ0SvXCK9tggcCd4O050ytNBMJC9qMOJ7yE0iOrFfOJSCfDAuPhn/rHFh79Kn1moF+/CE+nc0/2RPiLC8r54/rt5dYyyxJDfXg0a3VrrX39W69WZGW5OXiw=="),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
