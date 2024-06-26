package armlabservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4c2cdccf6ca3281dd50ed8788ce1de2e0d480973/specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Labs/getLab.json
func ExampleLabsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlabservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLabsClient().Get(ctx, "testrg123", "testlab", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Lab = armlabservices.Lab{
	// 	Name: to.Ptr("testlabplan"),
	// 	Type: to.Ptr("Microsoft.LabServices/Lab"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labs/testlab"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armlabservices.LabProperties{
	// 		Description: to.Ptr("This is a test lab."),
	// 		AutoShutdownProfile: &armlabservices.AutoShutdownProfile{
	// 			DisconnectDelay: to.Ptr("PT5M"),
	// 			IdleDelay: to.Ptr("PT5M"),
	// 			NoConnectDelay: to.Ptr("PT5M"),
	// 			ShutdownOnDisconnect: to.Ptr(armlabservices.EnableStateEnabled),
	// 			ShutdownOnIdle: to.Ptr(armlabservices.ShutdownOnIdleModeUserAbsence),
	// 			ShutdownWhenNotConnected: to.Ptr(armlabservices.EnableStateEnabled),
	// 		},
	// 		ConnectionProfile: &armlabservices.ConnectionProfile{
	// 			ClientRdpAccess: to.Ptr(armlabservices.ConnectionTypePublic),
	// 			ClientSSHAccess: to.Ptr(armlabservices.ConnectionTypePublic),
	// 			WebRdpAccess: to.Ptr(armlabservices.ConnectionTypeNone),
	// 			WebSSHAccess: to.Ptr(armlabservices.ConnectionTypeNone),
	// 		},
	// 		LabPlanID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labPlans/testlabplan"),
	// 		SecurityProfile: &armlabservices.SecurityProfile{
	// 			OpenAccess: to.Ptr(armlabservices.EnableStateDisabled),
	// 			RegistrationCode: to.Ptr("fAkEcodE"),
	// 		},
	// 		Title: to.Ptr("Test Lab"),
	// 		VirtualMachineProfile: &armlabservices.VirtualMachineProfile{
	// 			AdditionalCapabilities: &armlabservices.VirtualMachineAdditionalCapabilities{
	// 				InstallGpuDrivers: to.Ptr(armlabservices.EnableStateDisabled),
	// 			},
	// 			AdminUser: &armlabservices.Credentials{
	// 				Username: to.Ptr("test-user"),
	// 			},
	// 			CreateOption: to.Ptr(armlabservices.CreateOptionTemplateVM),
	// 			ImageReference: &armlabservices.ImageReference{
	// 				Offer: to.Ptr("WindowsServer"),
	// 				Publisher: to.Ptr("Microsoft"),
	// 				SKU: to.Ptr("2019-Datacenter"),
	// 				Version: to.Ptr("2019.0.20190410"),
	// 			},
	// 			OSType: to.Ptr(armlabservices.OsTypeWindows),
	// 			SKU: &armlabservices.SKU{
	// 				Name: to.Ptr("Medium"),
	// 				Capacity: to.Ptr[int32](20),
	// 			},
	// 			UsageQuota: to.Ptr("PT10H"),
	// 			UseSharedPassword: to.Ptr(armlabservices.EnableStateDisabled),
	// 		},
	// 		NetworkProfile: &armlabservices.LabNetworkProfile{
	// 			SubnetID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
	// 		},
	// 		ProvisioningState: to.Ptr(armlabservices.ProvisioningStateSucceeded),
	// 		State: to.Ptr(armlabservices.LabStateDraft),
	// 	},
	// 	SystemData: &armlabservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T10:00:00.000Z"); return t}()),
	// 		CreatedBy: to.Ptr("identity123"),
	// 		CreatedByType: to.Ptr(armlabservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-01T09:12:28.000Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identity123"),
	// 		LastModifiedByType: to.Ptr(armlabservices.CreatedByTypeUser),
	// 	},
	// }
}
