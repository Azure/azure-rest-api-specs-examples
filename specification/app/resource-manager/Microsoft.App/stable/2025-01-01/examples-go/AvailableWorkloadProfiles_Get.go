package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/AvailableWorkloadProfiles_Get.json
func ExampleAvailableWorkloadProfilesClient_NewGetPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailableWorkloadProfilesClient().NewGetPager("East US", nil)
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
		// page.AvailableWorkloadProfilesCollection = armappcontainers.AvailableWorkloadProfilesCollection{
		// 	Value: []*armappcontainers.AvailableWorkloadProfile{
		// 		{
		// 			Name: to.Ptr("Dedicated-D4"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Dedicated-D4"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityLocationDefault),
		// 				Category: to.Ptr("General purpose D-series"),
		// 				Cores: to.Ptr[int32](4),
		// 				DisplayName: to.Ptr("Dedicated-D4"),
		// 				MemoryGiB: to.Ptr[int32](16),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Dedicated-D4"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Dedicated-D8"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("General purpose D-series"),
		// 				Cores: to.Ptr[int32](8),
		// 				DisplayName: to.Ptr("Dedicated-D8"),
		// 				MemoryGiB: to.Ptr[int32](32),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Dedicated-D16"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Dedicated-D16"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("General purpose D-series"),
		// 				Cores: to.Ptr[int32](16),
		// 				DisplayName: to.Ptr("Dedicated-D16"),
		// 				MemoryGiB: to.Ptr[int32](64),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Dedicated-E4"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Dedicated-E4"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("Memory optimized E-series"),
		// 				Cores: to.Ptr[int32](4),
		// 				DisplayName: to.Ptr("Dedicated-E4"),
		// 				MemoryGiB: to.Ptr[int32](32),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Dedicated-E8"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Dedicated-E8"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("Memory optimized E-series"),
		// 				Cores: to.Ptr[int32](8),
		// 				DisplayName: to.Ptr("Dedicated-E8"),
		// 				MemoryGiB: to.Ptr[int32](64),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Dedicated-E16"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Dedicated-E16"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("Memory optimized E-series"),
		// 				Cores: to.Ptr[int32](16),
		// 				DisplayName: to.Ptr("Dedicated-E16"),
		// 				MemoryGiB: to.Ptr[int32](128),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Dedicated-F4"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Dedicated-F4"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("Compute optimized F-series"),
		// 				Cores: to.Ptr[int32](4),
		// 				DisplayName: to.Ptr("Dedicated-F4"),
		// 				MemoryGiB: to.Ptr[int32](8),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Dedicated-F8"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Dedicated-F8"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("Compute optimized F-series"),
		// 				Cores: to.Ptr[int32](8),
		// 				DisplayName: to.Ptr("Dedicated-F8"),
		// 				MemoryGiB: to.Ptr[int32](16),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Dedicated-F16"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Dedicated-F16"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("Compute optimized F-series"),
		// 				Cores: to.Ptr[int32](16),
		// 				DisplayName: to.Ptr("Dedicated-F16"),
		// 				MemoryGiB: to.Ptr[int32](32),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("NC48-A100"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/NC48-A100"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("GPU-NC-A100"),
		// 				Cores: to.Ptr[int32](48),
		// 				DisplayName: to.Ptr("Dedicated-NC48-A100"),
		// 				Gpus: to.Ptr[int32](2),
		// 				MemoryGiB: to.Ptr[int32](440),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Consumption"),
		// 			Type: to.Ptr("Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.App/availableManagedEnvironmentsWorkloadProfileTypes/Consumption"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.AvailableWorkloadProfileProperties{
		// 				Applicability: to.Ptr(armappcontainers.ApplicabilityCustom),
		// 				Category: to.Ptr("Consumption"),
		// 				Cores: to.Ptr[int32](3),
		// 				DisplayName: to.Ptr("Consumption"),
		// 				MemoryGiB: to.Ptr[int32](3),
		// 			},
		// 	}},
		// }
	}
}
