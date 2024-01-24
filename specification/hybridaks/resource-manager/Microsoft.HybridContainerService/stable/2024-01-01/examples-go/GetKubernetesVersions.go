package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/41e4538ed7bb3ceac3c1322c9455a0812ed110ac/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/GetKubernetesVersions.json
func ExampleClient_GetKubernetesVersions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().GetKubernetesVersions(ctx, "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubernetesVersionProfile = armhybridcontainerservice.KubernetesVersionProfile{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("microsoft.hybridcontainerservice/kubernetesVersions"),
	// 	ID: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation/providers/Microsoft.HybridContainerService/kubernetesVersions/default"),
	// 	ExtendedLocation: &armhybridcontainerservice.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
	// 		Type: to.Ptr(armhybridcontainerservice.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armhybridcontainerservice.KubernetesVersionProfileProperties{
	// 		ProvisioningState: to.Ptr(armhybridcontainerservice.ResourceProvisioningStateSucceeded),
	// 		Values: []*armhybridcontainerservice.KubernetesVersionProperties{
	// 			{
	// 				IsPreview: to.Ptr(false),
	// 				PatchVersions: map[string]*armhybridcontainerservice.KubernetesPatchVersions{
	// 					"1.23.12": &armhybridcontainerservice.KubernetesPatchVersions{
	// 						Readiness: []*armhybridcontainerservice.KubernetesVersionReadiness{
	// 							{
	// 								OSSKU: to.Ptr(armhybridcontainerservice.OSSKUCBLMariner),
	// 								OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
	// 								Ready: to.Ptr(true),
	// 							},
	// 							{
	// 								OSSKU: to.Ptr(armhybridcontainerservice.OSSKU("Windows")),
	// 								OSType: to.Ptr(armhybridcontainerservice.OsTypeWindows),
	// 								Ready: to.Ptr(true),
	// 							},
	// 							{
	// 								ErrorMessage: to.Ptr("Not Ready. Reasons: Failed to find proudct stream windoes2022 in release aks-hybrid-catalog-stable-int"),
	// 								OSSKU: to.Ptr(armhybridcontainerservice.OSSKUWindows2022),
	// 								OSType: to.Ptr(armhybridcontainerservice.OsTypeWindows),
	// 								Ready: to.Ptr(false),
	// 						}},
	// 						Upgrades: []*string{
	// 							to.Ptr("1.23.13")},
	// 						},
	// 						"1.23.13": &armhybridcontainerservice.KubernetesPatchVersions{
	// 							Readiness: []*armhybridcontainerservice.KubernetesVersionReadiness{
	// 								{
	// 									OSSKU: to.Ptr(armhybridcontainerservice.OSSKUCBLMariner),
	// 									OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
	// 									Ready: to.Ptr(true),
	// 								},
	// 								{
	// 									OSSKU: to.Ptr(armhybridcontainerservice.OSSKUWindows2019),
	// 									OSType: to.Ptr(armhybridcontainerservice.OsTypeWindows),
	// 									Ready: to.Ptr(true),
	// 								},
	// 								{
	// 									OSSKU: to.Ptr(armhybridcontainerservice.OSSKUWindows2022),
	// 									OSType: to.Ptr(armhybridcontainerservice.OsTypeWindows),
	// 									Ready: to.Ptr(true),
	// 							}},
	// 							Upgrades: []*string{
	// 							},
	// 						},
	// 					},
	// 					Version: to.Ptr("1.23"),
	// 			}},
	// 		},
	// 	}
}
