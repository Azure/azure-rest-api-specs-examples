package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListSku.json
func ExampleSKUsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListPager(nil)
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
		// page.SKUsResult = armapimanagement.SKUsResult{
		// 	Value: []*armapimanagement.SKU{
		// 		{
		// 			Name: to.Ptr("Developer"),
		// 			Capacity: &armapimanagement.SKUCapacity{
		// 				Default: to.Ptr[int32](1),
		// 				Maximum: to.Ptr[int32](1),
		// 				Minimum: to.Ptr[int32](1),
		// 				ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeNone),
		// 			},
		// 			LocationInfo: []*armapimanagement.SKULocationInfo{
		// 				{
		// 					Location: to.Ptr("centralindia"),
		// 					ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 					},
		// 					Zones: []*string{
		// 					},
		// 			}},
		// 			Locations: []*string{
		// 				to.Ptr("centralindia")},
		// 				ResourceType: to.Ptr("service"),
		// 				Restrictions: []*armapimanagement.SKURestrictions{
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Standard"),
		// 				Capacity: &armapimanagement.SKUCapacity{
		// 					Default: to.Ptr[int32](1),
		// 					Maximum: to.Ptr[int32](4),
		// 					Minimum: to.Ptr[int32](1),
		// 					ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeAutomatic),
		// 				},
		// 				LocationInfo: []*armapimanagement.SKULocationInfo{
		// 					{
		// 						Location: to.Ptr("centralindia"),
		// 						ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 						},
		// 						Zones: []*string{
		// 						},
		// 				}},
		// 				Locations: []*string{
		// 					to.Ptr("centralindia")},
		// 					ResourceType: to.Ptr("service"),
		// 					Restrictions: []*armapimanagement.SKURestrictions{
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("Premium"),
		// 					Capacity: &armapimanagement.SKUCapacity{
		// 						Default: to.Ptr[int32](1),
		// 						Maximum: to.Ptr[int32](10),
		// 						Minimum: to.Ptr[int32](1),
		// 						ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeAutomatic),
		// 					},
		// 					LocationInfo: []*armapimanagement.SKULocationInfo{
		// 						{
		// 							Location: to.Ptr("centralindia"),
		// 							ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 							},
		// 							Zones: []*string{
		// 							},
		// 					}},
		// 					Locations: []*string{
		// 						to.Ptr("centralindia")},
		// 						ResourceType: to.Ptr("service"),
		// 						Restrictions: []*armapimanagement.SKURestrictions{
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("Basic"),
		// 						Capacity: &armapimanagement.SKUCapacity{
		// 							Default: to.Ptr[int32](1),
		// 							Maximum: to.Ptr[int32](2),
		// 							Minimum: to.Ptr[int32](1),
		// 							ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeManual),
		// 						},
		// 						LocationInfo: []*armapimanagement.SKULocationInfo{
		// 							{
		// 								Location: to.Ptr("centralindia"),
		// 								ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 								},
		// 								Zones: []*string{
		// 								},
		// 						}},
		// 						Locations: []*string{
		// 							to.Ptr("centralindia")},
		// 							ResourceType: to.Ptr("service"),
		// 							Restrictions: []*armapimanagement.SKURestrictions{
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("Consumption"),
		// 							Capacity: &armapimanagement.SKUCapacity{
		// 								Default: to.Ptr[int32](0),
		// 								Maximum: to.Ptr[int32](0),
		// 								Minimum: to.Ptr[int32](0),
		// 								ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeNone),
		// 							},
		// 							LocationInfo: []*armapimanagement.SKULocationInfo{
		// 								{
		// 									Location: to.Ptr("centralindia"),
		// 									ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 									},
		// 									Zones: []*string{
		// 									},
		// 							}},
		// 							Locations: []*string{
		// 								to.Ptr("centralindia")},
		// 								ResourceType: to.Ptr("service"),
		// 								Restrictions: []*armapimanagement.SKURestrictions{
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("Developer"),
		// 								Capacity: &armapimanagement.SKUCapacity{
		// 									Default: to.Ptr[int32](1),
		// 									Maximum: to.Ptr[int32](1),
		// 									Minimum: to.Ptr[int32](1),
		// 									ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeNone),
		// 								},
		// 								LocationInfo: []*armapimanagement.SKULocationInfo{
		// 									{
		// 										Location: to.Ptr("uaenorth"),
		// 										ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 										},
		// 										Zones: []*string{
		// 										},
		// 								}},
		// 								Locations: []*string{
		// 									to.Ptr("uaenorth")},
		// 									ResourceType: to.Ptr("service"),
		// 									Restrictions: []*armapimanagement.SKURestrictions{
		// 									},
		// 								},
		// 								{
		// 									Name: to.Ptr("Standard"),
		// 									Capacity: &armapimanagement.SKUCapacity{
		// 										Default: to.Ptr[int32](1),
		// 										Maximum: to.Ptr[int32](4),
		// 										Minimum: to.Ptr[int32](1),
		// 										ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeAutomatic),
		// 									},
		// 									LocationInfo: []*armapimanagement.SKULocationInfo{
		// 										{
		// 											Location: to.Ptr("uaenorth"),
		// 											ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 											},
		// 											Zones: []*string{
		// 											},
		// 									}},
		// 									Locations: []*string{
		// 										to.Ptr("uaenorth")},
		// 										ResourceType: to.Ptr("service"),
		// 										Restrictions: []*armapimanagement.SKURestrictions{
		// 										},
		// 									},
		// 									{
		// 										Name: to.Ptr("Premium"),
		// 										Capacity: &armapimanagement.SKUCapacity{
		// 											Default: to.Ptr[int32](1),
		// 											Maximum: to.Ptr[int32](10),
		// 											Minimum: to.Ptr[int32](1),
		// 											ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeAutomatic),
		// 										},
		// 										LocationInfo: []*armapimanagement.SKULocationInfo{
		// 											{
		// 												Location: to.Ptr("uaenorth"),
		// 												ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 												},
		// 												Zones: []*string{
		// 												},
		// 										}},
		// 										Locations: []*string{
		// 											to.Ptr("uaenorth")},
		// 											ResourceType: to.Ptr("service"),
		// 											Restrictions: []*armapimanagement.SKURestrictions{
		// 											},
		// 										},
		// 										{
		// 											Name: to.Ptr("Basic"),
		// 											Capacity: &armapimanagement.SKUCapacity{
		// 												Default: to.Ptr[int32](1),
		// 												Maximum: to.Ptr[int32](2),
		// 												Minimum: to.Ptr[int32](1),
		// 												ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeManual),
		// 											},
		// 											LocationInfo: []*armapimanagement.SKULocationInfo{
		// 												{
		// 													Location: to.Ptr("uaenorth"),
		// 													ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 													},
		// 													Zones: []*string{
		// 													},
		// 											}},
		// 											Locations: []*string{
		// 												to.Ptr("uaenorth")},
		// 												ResourceType: to.Ptr("service"),
		// 												Restrictions: []*armapimanagement.SKURestrictions{
		// 												},
		// 											},
		// 											{
		// 												Name: to.Ptr("Developer"),
		// 												Capacity: &armapimanagement.SKUCapacity{
		// 													Default: to.Ptr[int32](1),
		// 													Maximum: to.Ptr[int32](1),
		// 													Minimum: to.Ptr[int32](1),
		// 													ScaleType: to.Ptr(armapimanagement.APIManagementSKUCapacityScaleTypeNone),
		// 												},
		// 												LocationInfo: []*armapimanagement.SKULocationInfo{
		// 													{
		// 														Location: to.Ptr("australiacentral"),
		// 														ZoneDetails: []*armapimanagement.SKUZoneDetails{
		// 														},
		// 														Zones: []*string{
		// 														},
		// 												}},
		// 												Locations: []*string{
		// 													to.Ptr("australiacentral")},
		// 													ResourceType: to.Ptr("service"),
		// 													Restrictions: []*armapimanagement.SKURestrictions{
		// 													},
		// 											}},
		// 										}
	}
}
