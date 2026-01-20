package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryScriptExamples/GalleryScriptVersion_Create.json
func ExampleGalleryScriptVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryScriptVersionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryScriptName", "1.0.0", armcompute.GalleryScriptVersion{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryScriptVersionProperties{
			PublishingProfile: &armcompute.GalleryScriptVersionPublishingProfile{
				EndOfLifeDate:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2027-07-01T07:00:00.000Z"); return t }()),
				ReplicaCount:       to.Ptr[int32](2),
				StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name:                 to.Ptr("West US"),
						ExcludeFromLatest:    to.Ptr(false),
						RegionalReplicaCount: to.Ptr[int32](2),
						StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardLRS),
					}},
				Source: &armcompute.ScriptSource{
					Parameters: []*armcompute.GalleryScriptParameter{
						{
							Name:         to.Ptr("location"),
							DefaultValue: to.Ptr("westus"),
							Required:     to.Ptr(true),
							Type:         to.Ptr(armcompute.GalleryScriptParameterTypeString),
						},
						{
							Name:         to.Ptr("myGalleryScriptParameter1"),
							Description:  to.Ptr("description of the parameter"),
							DefaultValue: to.Ptr("default value of parameter"),
							Required:     to.Ptr(true),
							Type:         to.Ptr(armcompute.GalleryScriptParameterTypeString),
						},
						{
							Name:         to.Ptr("myGalleryScriptParameter2"),
							Description:  to.Ptr("description of the parameter"),
							DefaultValue: to.Ptr("default value of parameter"),
							Required:     to.Ptr(false),
							Type:         to.Ptr(armcompute.GalleryScriptParameterTypeString),
						},
						{
							Name:         to.Ptr("numberOfUnits"),
							Description:  to.Ptr("description of the parameter"),
							DefaultValue: to.Ptr("3"),
							Required:     to.Ptr(true),
							Type:         to.Ptr(armcompute.GalleryScriptParameterTypeInt),
							MaxValue:     to.Ptr("5"),
							MinValue:     to.Ptr("1"),
						},
						{
							Name:         to.Ptr("weightOfUnit"),
							Description:  to.Ptr("description of the parameter"),
							DefaultValue: to.Ptr("0.6"),
							Required:     to.Ptr(true),
							Type:         to.Ptr(armcompute.GalleryScriptParameterTypeDouble),
							MaxValue:     to.Ptr("2"),
							MinValue:     to.Ptr("0.1"),
						},
						{
							Name:         to.Ptr("typeOfProduct"),
							Description:  to.Ptr("description of the parameter"),
							DefaultValue: to.Ptr("Fruit"),
							Required:     to.Ptr(false),
							Type:         to.Ptr(armcompute.GalleryScriptParameterTypeEnum),
							EnumValues: []*string{
								to.Ptr("Fruit"),
								to.Ptr("Vegetable"),
								to.Ptr("Greens"),
								to.Ptr("Nuts")},
						}},
					ScriptLink: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/myScript.ps1?{sasKey}"),
				},
			},
			SafetyProfile: &armcompute.GalleryScriptVersionSafetyProfile{
				AllowDeletionOfReplicatedLocations: to.Ptr(false),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GalleryScriptVersion = armcompute.GalleryScriptVersion{
	// 	Name: to.Ptr("1.0.0"),
	// 	Type: to.Ptr("Microsoft.Compute/galleries/scripts/versions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/scripts/myGalleryScript/versions/1.0.0"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryScriptVersionProperties{
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		PublishingProfile: &armcompute.GalleryScriptVersionPublishingProfile{
	// 			EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2027-07-01T07:00:00.000Z"); return t}()),
	// 			ExcludeFromLatest: to.Ptr(false),
	// 			PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-21T17:13:57.597Z"); return t}()),
	// 			ReplicaCount: to.Ptr[int32](2),
	// 			StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			TargetRegions: []*armcompute.TargetRegion{
	// 				{
	// 					Name: to.Ptr("West US"),
	// 					ExcludeFromLatest: to.Ptr(false),
	// 					RegionalReplicaCount: to.Ptr[int32](2),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			}},
	// 			Source: &armcompute.ScriptSource{
	// 				Parameters: []*armcompute.GalleryScriptParameter{
	// 					{
	// 						Name: to.Ptr("location"),
	// 						DefaultValue: to.Ptr("westus"),
	// 						Required: to.Ptr(true),
	// 						Type: to.Ptr(armcompute.GalleryScriptParameterTypeString),
	// 					},
	// 					{
	// 						Name: to.Ptr("myGalleryScriptParameter1"),
	// 						Description: to.Ptr("description of the parameter"),
	// 						DefaultValue: to.Ptr("default value of parameter"),
	// 						Required: to.Ptr(true),
	// 						Type: to.Ptr(armcompute.GalleryScriptParameterTypeString),
	// 					},
	// 					{
	// 						Name: to.Ptr("myGalleryScriptParameter2"),
	// 						Description: to.Ptr("description of the parameter"),
	// 						DefaultValue: to.Ptr("default value of parameter"),
	// 						Required: to.Ptr(false),
	// 						Type: to.Ptr(armcompute.GalleryScriptParameterTypeString),
	// 					},
	// 					{
	// 						Name: to.Ptr("numberOfUnits"),
	// 						Description: to.Ptr("description of the parameter"),
	// 						DefaultValue: to.Ptr("3"),
	// 						Required: to.Ptr(true),
	// 						Type: to.Ptr(armcompute.GalleryScriptParameterTypeInt),
	// 						MaxValue: to.Ptr("5"),
	// 						MinValue: to.Ptr("1"),
	// 					},
	// 					{
	// 						Name: to.Ptr("weightOfUnit"),
	// 						Description: to.Ptr("description of the parameter"),
	// 						DefaultValue: to.Ptr("0.6"),
	// 						Required: to.Ptr(true),
	// 						Type: to.Ptr(armcompute.GalleryScriptParameterTypeDouble),
	// 						MaxValue: to.Ptr("2"),
	// 						MinValue: to.Ptr("0.1"),
	// 					},
	// 					{
	// 						Name: to.Ptr("typeOfProduct"),
	// 						Description: to.Ptr("description of the parameter"),
	// 						DefaultValue: to.Ptr("Fruit"),
	// 						Required: to.Ptr(false),
	// 						Type: to.Ptr(armcompute.GalleryScriptParameterTypeEnum),
	// 						EnumValues: []*string{
	// 							to.Ptr("Fruit"),
	// 							to.Ptr("Vegetable"),
	// 							to.Ptr("Greens"),
	// 							to.Ptr("Nuts")},
	// 					}},
	// 					ScriptLink: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/myScript.ps1"),
	// 				},
	// 			},
	// 			SafetyProfile: &armcompute.GalleryScriptVersionSafetyProfile{
	// 				AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 			},
	// 		},
	// 	}
}
