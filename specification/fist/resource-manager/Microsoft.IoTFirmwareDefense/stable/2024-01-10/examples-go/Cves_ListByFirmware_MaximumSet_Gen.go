package armiotfirmwaredefense_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotfirmwaredefense/armiotfirmwaredefense"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Cves_ListByFirmware_MaximumSet_Gen.json
func ExampleCvesClient_NewListByFirmwarePager_cvesListByFirmwareMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiotfirmwaredefense.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCvesClient().NewListByFirmwarePager("FirmwareAnalysisRG", "default", "109a9886-50bf-85a8-9d75-000000000000", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CveListResult = armiotfirmwaredefense.CveListResult{
		// 	Value: []*armiotfirmwaredefense.CveResource{
		// 		{
		// 			Name: to.Ptr("7496e8a7-537e-40e4-a43b-000000000000"),
		// 			Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/cves"),
		// 			ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/cves/7496e8a7-537e-40e4-a43b-000000000000"),
		// 			Properties: &armiotfirmwaredefense.CveResult{
		// 				Name: to.Ptr("CVE-2018-1000500"),
		// 				Description: to.Ptr("Busybox contains a Missing SSL certificate validation vulnerability in The \"busybox wget\" applet that can result in arbitrary code execution. This attack appear to be exploitable via Simply download any file over HTTPS using \"busybox wget https://compromised-domain.com/important-file\"."),
		// 				Component: &armiotfirmwaredefense.CveComponent{
		// 				},
		// 				CveID: to.Ptr("7496e8a7-537e-40e4-a43b-000000000000"),
		// 				CvssScore: to.Ptr("8.1"),
		// 				CvssV2Score: to.Ptr("6.8"),
		// 				CvssV3Score: to.Ptr("8.1"),
		// 				CvssVersion: to.Ptr("3"),
		// 				Links: []*armiotfirmwaredefense.CveLink{
		// 					{
		// 						Href: to.Ptr("http://lists.busybox.net/pipermail/busybox/2018-May/086462.html"),
		// 						Label: to.Ptr("http://lists.busybox.net/pipermail/busybox/2018-May/086462.html"),
		// 					},
		// 					{
		// 						Href: to.Ptr("https://git.busybox.net/busybox/commit/?id=45fa3f18adf57ef9d743038743d9c90573aeeb91"),
		// 						Label: to.Ptr("https://git.busybox.net/busybox/commit/?id=45fa3f18adf57ef9d743038743d9c90573aeeb91"),
		// 					},
		// 					{
		// 						Href: to.Ptr("https://usn.ubuntu.com/4531-1/"),
		// 				}},
		// 				Severity: to.Ptr("High"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("fad414da-0055-4dfb-a3d4-000000000000"),
		// 			Type: to.Ptr("Microsoft.IoTFirmwareDefense/workspaces/firmwares/cves"),
		// 			ID: to.Ptr("/subscriptions/07aed47b-60ad-4d6e-a07a-000000000000/resourceGroups/FirmwareAnalysisRG/providers/Microsoft.IoTFirmwareDefense/workspaces/default/firmwares/109a9886-50bf-85a8-9d75-000000000000/cves/fad414da-0055-4dfb-a3d4-000000000000"),
		// 			Properties: &armiotfirmwaredefense.CveResult{
		// 				Name: to.Ptr("CVE-2022-28391"),
		// 				Description: to.Ptr("BusyBox through 1.35.0 allows remote attackers to execute arbitrary code if netstat is used to print a DNS PTR record's value to a VT compatible terminal. Alternatively, the attacker could choose to change the terminal's colors."),
		// 				Component: &armiotfirmwaredefense.CveComponent{
		// 				},
		// 				CveID: to.Ptr("fad414da-0055-4dfb-a3d4-000000000000"),
		// 				CvssScore: to.Ptr("8.8"),
		// 				CvssV2Score: to.Ptr("6.8"),
		// 				CvssV3Score: to.Ptr("8.8"),
		// 				CvssVersion: to.Ptr("3"),
		// 				Links: []*armiotfirmwaredefense.CveLink{
		// 					{
		// 						Href: to.Ptr("https://git.alpinelinux.org/aports/plain/main/busybox/0002-nslookup-sanitize-all-printed-strings-with-printable.patch"),
		// 						Label: to.Ptr("https://git.alpinelinux.org/aports/plain/main/busybox/0002-nslookup-sanitize-all-printed-strings-with-printable.patch"),
		// 					},
		// 					{
		// 						Href: to.Ptr("https://gitlab.alpinelinux.org/alpine/aports/-/issues/13661"),
		// 						Label: to.Ptr("https://gitlab.alpinelinux.org/alpine/aports/-/issues/13661"),
		// 					},
		// 					{
		// 						Href: to.Ptr("https://git.alpinelinux.org/aports/plain/main/busybox/0001-libbb-sockaddr2str-ensure-only-printable-characters-.patch"),
		// 						Label: to.Ptr("https://git.alpinelinux.org/aports/plain/main/busybox/0001-libbb-sockaddr2str-ensure-only-printable-characters-.patch"),
		// 				}},
		// 				Severity: to.Ptr("High"),
		// 			},
		// 	}},
		// }
	}
}
