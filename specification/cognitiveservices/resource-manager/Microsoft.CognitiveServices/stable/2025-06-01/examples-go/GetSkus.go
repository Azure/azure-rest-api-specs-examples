package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/44319b51c6f952fdc9543d3dc4fdd9959350d102/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/GetSkus.json
func ExampleResourceSKUsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceSKUsClient().NewListPager(nil)
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
		// page.ResourceSKUListResult = armcognitiveservices.ResourceSKUListResult{
		// 	Value: []*armcognitiveservices.ResourceSKU{
		// 		{
		// 			Name: to.Ptr("F0"),
		// 			Kind: to.Ptr("Bing.Speech"),
		// 			Locations: []*string{
		// 				to.Ptr("GLOBAL")},
		// 				ResourceType: to.Ptr("accounts"),
		// 				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 				},
		// 				Tier: to.Ptr("Free"),
		// 			},
		// 			{
		// 				Name: to.Ptr("S0"),
		// 				Kind: to.Ptr("Bing.Speech"),
		// 				Locations: []*string{
		// 					to.Ptr("GLOBAL")},
		// 					ResourceType: to.Ptr("accounts"),
		// 					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 					},
		// 					Tier: to.Ptr("Standard"),
		// 				},
		// 				{
		// 					Name: to.Ptr("F0"),
		// 					Kind: to.Ptr("SpeechTranslation"),
		// 					Locations: []*string{
		// 						to.Ptr("GLOBAL")},
		// 						ResourceType: to.Ptr("accounts"),
		// 						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 						},
		// 						Tier: to.Ptr("Free"),
		// 					},
		// 					{
		// 						Name: to.Ptr("S1"),
		// 						Kind: to.Ptr("SpeechTranslation"),
		// 						Locations: []*string{
		// 							to.Ptr("GLOBAL")},
		// 							ResourceType: to.Ptr("accounts"),
		// 							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 							},
		// 							Tier: to.Ptr("Standard"),
		// 						},
		// 						{
		// 							Name: to.Ptr("S2"),
		// 							Kind: to.Ptr("SpeechTranslation"),
		// 							Locations: []*string{
		// 								to.Ptr("GLOBAL")},
		// 								ResourceType: to.Ptr("accounts"),
		// 								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 								},
		// 								Tier: to.Ptr("Standard"),
		// 							},
		// 							{
		// 								Name: to.Ptr("S3"),
		// 								Kind: to.Ptr("SpeechTranslation"),
		// 								Locations: []*string{
		// 									to.Ptr("GLOBAL")},
		// 									ResourceType: to.Ptr("accounts"),
		// 									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 									},
		// 									Tier: to.Ptr("Standard"),
		// 								},
		// 								{
		// 									Name: to.Ptr("S4"),
		// 									Kind: to.Ptr("SpeechTranslation"),
		// 									Locations: []*string{
		// 										to.Ptr("GLOBAL")},
		// 										ResourceType: to.Ptr("accounts"),
		// 										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 										},
		// 										Tier: to.Ptr("Standard"),
		// 									},
		// 									{
		// 										Name: to.Ptr("F0"),
		// 										Kind: to.Ptr("TextTranslation"),
		// 										Locations: []*string{
		// 											to.Ptr("GLOBAL")},
		// 											ResourceType: to.Ptr("accounts"),
		// 											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 											},
		// 											Tier: to.Ptr("Free"),
		// 										},
		// 										{
		// 											Name: to.Ptr("S1"),
		// 											Kind: to.Ptr("TextTranslation"),
		// 											Locations: []*string{
		// 												to.Ptr("GLOBAL")},
		// 												ResourceType: to.Ptr("accounts"),
		// 												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 												},
		// 												Tier: to.Ptr("Standard"),
		// 											},
		// 											{
		// 												Name: to.Ptr("S2"),
		// 												Kind: to.Ptr("TextTranslation"),
		// 												Locations: []*string{
		// 													to.Ptr("GLOBAL")},
		// 													ResourceType: to.Ptr("accounts"),
		// 													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 													},
		// 													Tier: to.Ptr("Standard"),
		// 												},
		// 												{
		// 													Name: to.Ptr("S3"),
		// 													Kind: to.Ptr("TextTranslation"),
		// 													Locations: []*string{
		// 														to.Ptr("GLOBAL")},
		// 														ResourceType: to.Ptr("accounts"),
		// 														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 														},
		// 														Tier: to.Ptr("Standard"),
		// 													},
		// 													{
		// 														Name: to.Ptr("S4"),
		// 														Kind: to.Ptr("TextTranslation"),
		// 														Locations: []*string{
		// 															to.Ptr("GLOBAL")},
		// 															ResourceType: to.Ptr("accounts"),
		// 															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 															},
		// 															Tier: to.Ptr("Standard"),
		// 														},
		// 														{
		// 															Name: to.Ptr("S1"),
		// 															Kind: to.Ptr("Bing.Search.v7"),
		// 															Locations: []*string{
		// 																to.Ptr("GLOBAL")},
		// 																ResourceType: to.Ptr("accounts"),
		// 																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																},
		// 																Tier: to.Ptr("Standard"),
		// 															},
		// 															{
		// 																Name: to.Ptr("S2"),
		// 																Kind: to.Ptr("Bing.Search.v7"),
		// 																Locations: []*string{
		// 																	to.Ptr("GLOBAL")},
		// 																	ResourceType: to.Ptr("accounts"),
		// 																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																	},
		// 																	Tier: to.Ptr("Standard"),
		// 																},
		// 																{
		// 																	Name: to.Ptr("S3"),
		// 																	Kind: to.Ptr("Bing.Search.v7"),
		// 																	Locations: []*string{
		// 																		to.Ptr("GLOBAL")},
		// 																		ResourceType: to.Ptr("accounts"),
		// 																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																		},
		// 																		Tier: to.Ptr("Standard"),
		// 																	},
		// 																	{
		// 																		Name: to.Ptr("S4"),
		// 																		Kind: to.Ptr("Bing.Search.v7"),
		// 																		Locations: []*string{
		// 																			to.Ptr("GLOBAL")},
		// 																			ResourceType: to.Ptr("accounts"),
		// 																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																			},
		// 																			Tier: to.Ptr("Standard"),
		// 																		},
		// 																		{
		// 																			Name: to.Ptr("S5"),
		// 																			Kind: to.Ptr("Bing.Search.v7"),
		// 																			Locations: []*string{
		// 																				to.Ptr("GLOBAL")},
		// 																				ResourceType: to.Ptr("accounts"),
		// 																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																				},
		// 																				Tier: to.Ptr("Standard"),
		// 																			},
		// 																			{
		// 																				Name: to.Ptr("S6"),
		// 																				Kind: to.Ptr("Bing.Search.v7"),
		// 																				Locations: []*string{
		// 																					to.Ptr("GLOBAL")},
		// 																					ResourceType: to.Ptr("accounts"),
		// 																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																					},
		// 																					Tier: to.Ptr("Standard"),
		// 																				},
		// 																				{
		// 																					Name: to.Ptr("S7"),
		// 																					Kind: to.Ptr("Bing.Search.v7"),
		// 																					Locations: []*string{
		// 																						to.Ptr("GLOBAL")},
		// 																						ResourceType: to.Ptr("accounts"),
		// 																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																						},
		// 																						Tier: to.Ptr("Standard"),
		// 																					},
		// 																					{
		// 																						Name: to.Ptr("S8"),
		// 																						Kind: to.Ptr("Bing.Search.v7"),
		// 																						Locations: []*string{
		// 																							to.Ptr("GLOBAL")},
		// 																							ResourceType: to.Ptr("accounts"),
		// 																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																							},
		// 																							Tier: to.Ptr("Standard"),
		// 																						},
		// 																						{
		// 																							Name: to.Ptr("S1"),
		// 																							Kind: to.Ptr("Bing.Autosuggest.v7"),
		// 																							Locations: []*string{
		// 																								to.Ptr("GLOBAL")},
		// 																								ResourceType: to.Ptr("accounts"),
		// 																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																								},
		// 																								Tier: to.Ptr("Standard"),
		// 																							},
		// 																							{
		// 																								Name: to.Ptr("S1"),
		// 																								Kind: to.Ptr("Bing.CustomSearch"),
		// 																								Locations: []*string{
		// 																									to.Ptr("GLOBAL")},
		// 																									ResourceType: to.Ptr("accounts"),
		// 																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																									},
		// 																									Tier: to.Ptr("Standard"),
		// 																								},
		// 																								{
		// 																									Name: to.Ptr("S1"),
		// 																									Kind: to.Ptr("Bing.SpellCheck.v7"),
		// 																									Locations: []*string{
		// 																										to.Ptr("GLOBAL")},
		// 																										ResourceType: to.Ptr("accounts"),
		// 																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																										},
		// 																										Tier: to.Ptr("Standard"),
		// 																									},
		// 																									{
		// 																										Name: to.Ptr("F0"),
		// 																										Kind: to.Ptr("Bing.EntitySearch"),
		// 																										Locations: []*string{
		// 																											to.Ptr("GLOBAL")},
		// 																											ResourceType: to.Ptr("accounts"),
		// 																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																											},
		// 																											Tier: to.Ptr("Free"),
		// 																										},
		// 																										{
		// 																											Name: to.Ptr("S1"),
		// 																											Kind: to.Ptr("Bing.EntitySearch"),
		// 																											Locations: []*string{
		// 																												to.Ptr("GLOBAL")},
		// 																												ResourceType: to.Ptr("accounts"),
		// 																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																												},
		// 																												Tier: to.Ptr("Standard"),
		// 																											},
		// 																											{
		// 																												Name: to.Ptr("F0"),
		// 																												Kind: to.Ptr("Face"),
		// 																												Locations: []*string{
		// 																													to.Ptr("AUSTRALIAEAST")},
		// 																													ResourceType: to.Ptr("accounts"),
		// 																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																													},
		// 																													Tier: to.Ptr("Free"),
		// 																												},
		// 																												{
		// 																													Name: to.Ptr("S0"),
		// 																													Kind: to.Ptr("Face"),
		// 																													Locations: []*string{
		// 																														to.Ptr("AUSTRALIAEAST")},
		// 																														ResourceType: to.Ptr("accounts"),
		// 																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																														},
		// 																														Tier: to.Ptr("Standard"),
		// 																													},
		// 																													{
		// 																														Name: to.Ptr("F0"),
		// 																														Kind: to.Ptr("ComputerVision"),
		// 																														Locations: []*string{
		// 																															to.Ptr("AUSTRALIAEAST")},
		// 																															ResourceType: to.Ptr("accounts"),
		// 																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																															},
		// 																															Tier: to.Ptr("Free"),
		// 																														},
		// 																														{
		// 																															Name: to.Ptr("S1"),
		// 																															Kind: to.Ptr("ComputerVision"),
		// 																															Locations: []*string{
		// 																																to.Ptr("AUSTRALIAEAST")},
		// 																																ResourceType: to.Ptr("accounts"),
		// 																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																},
		// 																																Tier: to.Ptr("Standard"),
		// 																															},
		// 																															{
		// 																																Name: to.Ptr("F0"),
		// 																																Kind: to.Ptr("ContentModerator"),
		// 																																Locations: []*string{
		// 																																	to.Ptr("AUSTRALIAEAST")},
		// 																																	ResourceType: to.Ptr("accounts"),
		// 																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																	},
		// 																																	Tier: to.Ptr("Free"),
		// 																																},
		// 																																{
		// 																																	Name: to.Ptr("S0"),
		// 																																	Kind: to.Ptr("ContentModerator"),
		// 																																	Locations: []*string{
		// 																																		to.Ptr("AUSTRALIAEAST")},
		// 																																		ResourceType: to.Ptr("accounts"),
		// 																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																		},
		// 																																		Tier: to.Ptr("Standard"),
		// 																																	},
		// 																																	{
		// 																																		Name: to.Ptr("F0"),
		// 																																		Kind: to.Ptr("TextAnalytics"),
		// 																																		Locations: []*string{
		// 																																			to.Ptr("AUSTRALIAEAST")},
		// 																																			ResourceType: to.Ptr("accounts"),
		// 																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																			},
		// 																																			Tier: to.Ptr("Free"),
		// 																																		},
		// 																																		{
		// 																																			Name: to.Ptr("S0"),
		// 																																			Kind: to.Ptr("TextAnalytics"),
		// 																																			Locations: []*string{
		// 																																				to.Ptr("AUSTRALIAEAST")},
		// 																																				ResourceType: to.Ptr("accounts"),
		// 																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																				},
		// 																																				Tier: to.Ptr("Standard"),
		// 																																			},
		// 																																			{
		// 																																				Name: to.Ptr("S1"),
		// 																																				Kind: to.Ptr("TextAnalytics"),
		// 																																				Locations: []*string{
		// 																																					to.Ptr("AUSTRALIAEAST")},
		// 																																					ResourceType: to.Ptr("accounts"),
		// 																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																					},
		// 																																					Tier: to.Ptr("Standard"),
		// 																																				},
		// 																																				{
		// 																																					Name: to.Ptr("S2"),
		// 																																					Kind: to.Ptr("TextAnalytics"),
		// 																																					Locations: []*string{
		// 																																						to.Ptr("AUSTRALIAEAST")},
		// 																																						ResourceType: to.Ptr("accounts"),
		// 																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																						},
		// 																																						Tier: to.Ptr("Standard"),
		// 																																					},
		// 																																					{
		// 																																						Name: to.Ptr("S3"),
		// 																																						Kind: to.Ptr("TextAnalytics"),
		// 																																						Locations: []*string{
		// 																																							to.Ptr("AUSTRALIAEAST")},
		// 																																							ResourceType: to.Ptr("accounts"),
		// 																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																							},
		// 																																							Tier: to.Ptr("Standard"),
		// 																																						},
		// 																																						{
		// 																																							Name: to.Ptr("S4"),
		// 																																							Kind: to.Ptr("TextAnalytics"),
		// 																																							Locations: []*string{
		// 																																								to.Ptr("AUSTRALIAEAST")},
		// 																																								ResourceType: to.Ptr("accounts"),
		// 																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																								},
		// 																																								Tier: to.Ptr("Standard"),
		// 																																							},
		// 																																							{
		// 																																								Name: to.Ptr("F0"),
		// 																																								Kind: to.Ptr("LUIS"),
		// 																																								Locations: []*string{
		// 																																									to.Ptr("AUSTRALIAEAST")},
		// 																																									ResourceType: to.Ptr("accounts"),
		// 																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																									},
		// 																																									Tier: to.Ptr("Free"),
		// 																																								},
		// 																																								{
		// 																																									Name: to.Ptr("S0"),
		// 																																									Kind: to.Ptr("LUIS"),
		// 																																									Locations: []*string{
		// 																																										to.Ptr("AUSTRALIAEAST")},
		// 																																										ResourceType: to.Ptr("accounts"),
		// 																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																										},
		// 																																										Tier: to.Ptr("Standard"),
		// 																																									},
		// 																																									{
		// 																																										Name: to.Ptr("F0"),
		// 																																										Kind: to.Ptr("Face"),
		// 																																										Locations: []*string{
		// 																																											to.Ptr("BRAZILSOUTH")},
		// 																																											ResourceType: to.Ptr("accounts"),
		// 																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																											},
		// 																																											Tier: to.Ptr("Free"),
		// 																																										},
		// 																																										{
		// 																																											Name: to.Ptr("S0"),
		// 																																											Kind: to.Ptr("Face"),
		// 																																											Locations: []*string{
		// 																																												to.Ptr("BRAZILSOUTH")},
		// 																																												ResourceType: to.Ptr("accounts"),
		// 																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																												},
		// 																																												Tier: to.Ptr("Standard"),
		// 																																											},
		// 																																											{
		// 																																												Name: to.Ptr("F0"),
		// 																																												Kind: to.Ptr("ComputerVision"),
		// 																																												Locations: []*string{
		// 																																													to.Ptr("BRAZILSOUTH")},
		// 																																													ResourceType: to.Ptr("accounts"),
		// 																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																													},
		// 																																													Tier: to.Ptr("Free"),
		// 																																												},
		// 																																												{
		// 																																													Name: to.Ptr("S1"),
		// 																																													Kind: to.Ptr("ComputerVision"),
		// 																																													Locations: []*string{
		// 																																														to.Ptr("BRAZILSOUTH")},
		// 																																														ResourceType: to.Ptr("accounts"),
		// 																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																														},
		// 																																														Tier: to.Ptr("Standard"),
		// 																																													},
		// 																																													{
		// 																																														Name: to.Ptr("F0"),
		// 																																														Kind: to.Ptr("ContentModerator"),
		// 																																														Locations: []*string{
		// 																																															to.Ptr("BRAZILSOUTH")},
		// 																																															ResourceType: to.Ptr("accounts"),
		// 																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																															},
		// 																																															Tier: to.Ptr("Free"),
		// 																																														},
		// 																																														{
		// 																																															Name: to.Ptr("S0"),
		// 																																															Kind: to.Ptr("ContentModerator"),
		// 																																															Locations: []*string{
		// 																																																to.Ptr("BRAZILSOUTH")},
		// 																																																ResourceType: to.Ptr("accounts"),
		// 																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																},
		// 																																																Tier: to.Ptr("Standard"),
		// 																																															},
		// 																																															{
		// 																																																Name: to.Ptr("F0"),
		// 																																																Kind: to.Ptr("TextAnalytics"),
		// 																																																Locations: []*string{
		// 																																																	to.Ptr("BRAZILSOUTH")},
		// 																																																	ResourceType: to.Ptr("accounts"),
		// 																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																	},
		// 																																																	Tier: to.Ptr("Free"),
		// 																																																},
		// 																																																{
		// 																																																	Name: to.Ptr("S0"),
		// 																																																	Kind: to.Ptr("TextAnalytics"),
		// 																																																	Locations: []*string{
		// 																																																		to.Ptr("BRAZILSOUTH")},
		// 																																																		ResourceType: to.Ptr("accounts"),
		// 																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																		},
		// 																																																		Tier: to.Ptr("Standard"),
		// 																																																	},
		// 																																																	{
		// 																																																		Name: to.Ptr("S1"),
		// 																																																		Kind: to.Ptr("TextAnalytics"),
		// 																																																		Locations: []*string{
		// 																																																			to.Ptr("BRAZILSOUTH")},
		// 																																																			ResourceType: to.Ptr("accounts"),
		// 																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																			},
		// 																																																			Tier: to.Ptr("Standard"),
		// 																																																		},
		// 																																																		{
		// 																																																			Name: to.Ptr("S2"),
		// 																																																			Kind: to.Ptr("TextAnalytics"),
		// 																																																			Locations: []*string{
		// 																																																				to.Ptr("BRAZILSOUTH")},
		// 																																																				ResourceType: to.Ptr("accounts"),
		// 																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																				},
		// 																																																				Tier: to.Ptr("Standard"),
		// 																																																			},
		// 																																																			{
		// 																																																				Name: to.Ptr("S3"),
		// 																																																				Kind: to.Ptr("TextAnalytics"),
		// 																																																				Locations: []*string{
		// 																																																					to.Ptr("BRAZILSOUTH")},
		// 																																																					ResourceType: to.Ptr("accounts"),
		// 																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																					},
		// 																																																					Tier: to.Ptr("Standard"),
		// 																																																				},
		// 																																																				{
		// 																																																					Name: to.Ptr("S4"),
		// 																																																					Kind: to.Ptr("TextAnalytics"),
		// 																																																					Locations: []*string{
		// 																																																						to.Ptr("BRAZILSOUTH")},
		// 																																																						ResourceType: to.Ptr("accounts"),
		// 																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																						},
		// 																																																						Tier: to.Ptr("Standard"),
		// 																																																					},
		// 																																																					{
		// 																																																						Name: to.Ptr("F0"),
		// 																																																						Kind: to.Ptr("LUIS"),
		// 																																																						Locations: []*string{
		// 																																																							to.Ptr("BRAZILSOUTH")},
		// 																																																							ResourceType: to.Ptr("accounts"),
		// 																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																							},
		// 																																																							Tier: to.Ptr("Free"),
		// 																																																						},
		// 																																																						{
		// 																																																							Name: to.Ptr("S0"),
		// 																																																							Kind: to.Ptr("LUIS"),
		// 																																																							Locations: []*string{
		// 																																																								to.Ptr("BRAZILSOUTH")},
		// 																																																								ResourceType: to.Ptr("accounts"),
		// 																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																								},
		// 																																																								Tier: to.Ptr("Standard"),
		// 																																																							},
		// 																																																							{
		// 																																																								Name: to.Ptr("F0"),
		// 																																																								Kind: to.Ptr("Face"),
		// 																																																								Locations: []*string{
		// 																																																									to.Ptr("CENTRALUSEUAP")},
		// 																																																									ResourceType: to.Ptr("accounts"),
		// 																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																									},
		// 																																																									Tier: to.Ptr("Free"),
		// 																																																								},
		// 																																																								{
		// 																																																									Name: to.Ptr("S0"),
		// 																																																									Kind: to.Ptr("Face"),
		// 																																																									Locations: []*string{
		// 																																																										to.Ptr("CENTRALUSEUAP")},
		// 																																																										ResourceType: to.Ptr("accounts"),
		// 																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																										},
		// 																																																										Tier: to.Ptr("Standard"),
		// 																																																									},
		// 																																																									{
		// 																																																										Name: to.Ptr("F0"),
		// 																																																										Kind: to.Ptr("ComputerVision"),
		// 																																																										Locations: []*string{
		// 																																																											to.Ptr("CENTRALUSEUAP")},
		// 																																																											ResourceType: to.Ptr("accounts"),
		// 																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																											},
		// 																																																											Tier: to.Ptr("Free"),
		// 																																																										},
		// 																																																										{
		// 																																																											Name: to.Ptr("S1"),
		// 																																																											Kind: to.Ptr("ComputerVision"),
		// 																																																											Locations: []*string{
		// 																																																												to.Ptr("CENTRALUSEUAP")},
		// 																																																												ResourceType: to.Ptr("accounts"),
		// 																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																												},
		// 																																																												Tier: to.Ptr("Standard"),
		// 																																																											},
		// 																																																											{
		// 																																																												Name: to.Ptr("F0"),
		// 																																																												Kind: to.Ptr("TextAnalytics"),
		// 																																																												Locations: []*string{
		// 																																																													to.Ptr("CENTRALUSEUAP")},
		// 																																																													ResourceType: to.Ptr("accounts"),
		// 																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																													},
		// 																																																													Tier: to.Ptr("Free"),
		// 																																																												},
		// 																																																												{
		// 																																																													Name: to.Ptr("S0"),
		// 																																																													Kind: to.Ptr("TextAnalytics"),
		// 																																																													Locations: []*string{
		// 																																																														to.Ptr("CENTRALUSEUAP")},
		// 																																																														ResourceType: to.Ptr("accounts"),
		// 																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																														},
		// 																																																														Tier: to.Ptr("Standard"),
		// 																																																													},
		// 																																																													{
		// 																																																														Name: to.Ptr("S1"),
		// 																																																														Kind: to.Ptr("TextAnalytics"),
		// 																																																														Locations: []*string{
		// 																																																															to.Ptr("CENTRALUSEUAP")},
		// 																																																															ResourceType: to.Ptr("accounts"),
		// 																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																															},
		// 																																																															Tier: to.Ptr("Standard"),
		// 																																																														},
		// 																																																														{
		// 																																																															Name: to.Ptr("S2"),
		// 																																																															Kind: to.Ptr("TextAnalytics"),
		// 																																																															Locations: []*string{
		// 																																																																to.Ptr("CENTRALUSEUAP")},
		// 																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																},
		// 																																																																Tier: to.Ptr("Standard"),
		// 																																																															},
		// 																																																															{
		// 																																																																Name: to.Ptr("S3"),
		// 																																																																Kind: to.Ptr("TextAnalytics"),
		// 																																																																Locations: []*string{
		// 																																																																	to.Ptr("CENTRALUSEUAP")},
		// 																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																	},
		// 																																																																	Tier: to.Ptr("Standard"),
		// 																																																																},
		// 																																																																{
		// 																																																																	Name: to.Ptr("S4"),
		// 																																																																	Kind: to.Ptr("TextAnalytics"),
		// 																																																																	Locations: []*string{
		// 																																																																		to.Ptr("CENTRALUSEUAP")},
		// 																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																		},
		// 																																																																		Tier: to.Ptr("Standard"),
		// 																																																																	},
		// 																																																																	{
		// 																																																																		Name: to.Ptr("F0"),
		// 																																																																		Kind: to.Ptr("ContentModerator"),
		// 																																																																		Locations: []*string{
		// 																																																																			to.Ptr("CENTRALUSEUAP")},
		// 																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																			},
		// 																																																																			Tier: to.Ptr("Free"),
		// 																																																																		},
		// 																																																																		{
		// 																																																																			Name: to.Ptr("S0"),
		// 																																																																			Kind: to.Ptr("ContentModerator"),
		// 																																																																			Locations: []*string{
		// 																																																																				to.Ptr("CENTRALUSEUAP")},
		// 																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																				},
		// 																																																																				Tier: to.Ptr("Standard"),
		// 																																																																			},
		// 																																																																			{
		// 																																																																				Name: to.Ptr("F0"),
		// 																																																																				Kind: to.Ptr("LUIS"),
		// 																																																																				Locations: []*string{
		// 																																																																					to.Ptr("CENTRALUSEUAP")},
		// 																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																					},
		// 																																																																					Tier: to.Ptr("Free"),
		// 																																																																				},
		// 																																																																				{
		// 																																																																					Name: to.Ptr("S0"),
		// 																																																																					Kind: to.Ptr("LUIS"),
		// 																																																																					Locations: []*string{
		// 																																																																						to.Ptr("CENTRALUSEUAP")},
		// 																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																						},
		// 																																																																						Tier: to.Ptr("Standard"),
		// 																																																																					},
		// 																																																																					{
		// 																																																																						Name: to.Ptr("F0"),
		// 																																																																						Kind: to.Ptr("ContentModerator"),
		// 																																																																						Locations: []*string{
		// 																																																																							to.Ptr("WESTUS")},
		// 																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																							},
		// 																																																																							Tier: to.Ptr("Free"),
		// 																																																																						},
		// 																																																																						{
		// 																																																																							Name: to.Ptr("S0"),
		// 																																																																							Kind: to.Ptr("ContentModerator"),
		// 																																																																							Locations: []*string{
		// 																																																																								to.Ptr("WESTUS")},
		// 																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																								},
		// 																																																																								Tier: to.Ptr("Standard"),
		// 																																																																							},
		// 																																																																							{
		// 																																																																								Name: to.Ptr("F0"),
		// 																																																																								Kind: to.Ptr("LUIS"),
		// 																																																																								Locations: []*string{
		// 																																																																									to.Ptr("WESTUS")},
		// 																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																									},
		// 																																																																									Tier: to.Ptr("Free"),
		// 																																																																								},
		// 																																																																								{
		// 																																																																									Name: to.Ptr("S0"),
		// 																																																																									Kind: to.Ptr("LUIS"),
		// 																																																																									Locations: []*string{
		// 																																																																										to.Ptr("WESTUS")},
		// 																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																										},
		// 																																																																										Tier: to.Ptr("Standard"),
		// 																																																																									},
		// 																																																																									{
		// 																																																																										Name: to.Ptr("F0"),
		// 																																																																										Kind: to.Ptr("Face"),
		// 																																																																										Locations: []*string{
		// 																																																																											to.Ptr("WESTUS")},
		// 																																																																											ResourceType: to.Ptr("accounts"),
		// 																																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																											},
		// 																																																																											Tier: to.Ptr("Free"),
		// 																																																																										},
		// 																																																																										{
		// 																																																																											Name: to.Ptr("S0"),
		// 																																																																											Kind: to.Ptr("Face"),
		// 																																																																											Locations: []*string{
		// 																																																																												to.Ptr("WESTUS")},
		// 																																																																												ResourceType: to.Ptr("accounts"),
		// 																																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																												},
		// 																																																																												Tier: to.Ptr("Standard"),
		// 																																																																											},
		// 																																																																											{
		// 																																																																												Name: to.Ptr("F0"),
		// 																																																																												Kind: to.Ptr("TextAnalytics"),
		// 																																																																												Locations: []*string{
		// 																																																																													to.Ptr("WESTUS")},
		// 																																																																													ResourceType: to.Ptr("accounts"),
		// 																																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																													},
		// 																																																																													Tier: to.Ptr("Free"),
		// 																																																																												},
		// 																																																																												{
		// 																																																																													Name: to.Ptr("S0"),
		// 																																																																													Kind: to.Ptr("TextAnalytics"),
		// 																																																																													Locations: []*string{
		// 																																																																														to.Ptr("WESTUS")},
		// 																																																																														ResourceType: to.Ptr("accounts"),
		// 																																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																														},
		// 																																																																														Tier: to.Ptr("Standard"),
		// 																																																																													},
		// 																																																																													{
		// 																																																																														Name: to.Ptr("S1"),
		// 																																																																														Kind: to.Ptr("TextAnalytics"),
		// 																																																																														Locations: []*string{
		// 																																																																															to.Ptr("WESTUS")},
		// 																																																																															ResourceType: to.Ptr("accounts"),
		// 																																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																															},
		// 																																																																															Tier: to.Ptr("Standard"),
		// 																																																																														},
		// 																																																																														{
		// 																																																																															Name: to.Ptr("S2"),
		// 																																																																															Kind: to.Ptr("TextAnalytics"),
		// 																																																																															Locations: []*string{
		// 																																																																																to.Ptr("WESTUS")},
		// 																																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																},
		// 																																																																																Tier: to.Ptr("Standard"),
		// 																																																																															},
		// 																																																																															{
		// 																																																																																Name: to.Ptr("S3"),
		// 																																																																																Kind: to.Ptr("TextAnalytics"),
		// 																																																																																Locations: []*string{
		// 																																																																																	to.Ptr("WESTUS")},
		// 																																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																	},
		// 																																																																																	Tier: to.Ptr("Standard"),
		// 																																																																																},
		// 																																																																																{
		// 																																																																																	Name: to.Ptr("S4"),
		// 																																																																																	Kind: to.Ptr("TextAnalytics"),
		// 																																																																																	Locations: []*string{
		// 																																																																																		to.Ptr("WESTUS")},
		// 																																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																		},
		// 																																																																																		Tier: to.Ptr("Standard"),
		// 																																																																																	},
		// 																																																																																	{
		// 																																																																																		Name: to.Ptr("F0"),
		// 																																																																																		Kind: to.Ptr("SpeakerRecognition"),
		// 																																																																																		Locations: []*string{
		// 																																																																																			to.Ptr("WESTUS")},
		// 																																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																			},
		// 																																																																																			Tier: to.Ptr("Free"),
		// 																																																																																		},
		// 																																																																																		{
		// 																																																																																			Name: to.Ptr("S0"),
		// 																																																																																			Kind: to.Ptr("SpeakerRecognition"),
		// 																																																																																			Locations: []*string{
		// 																																																																																				to.Ptr("WESTUS")},
		// 																																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																				},
		// 																																																																																				Tier: to.Ptr("Standard"),
		// 																																																																																			},
		// 																																																																																			{
		// 																																																																																				Name: to.Ptr("F0"),
		// 																																																																																				Kind: to.Ptr("CustomSpeech"),
		// 																																																																																				Locations: []*string{
		// 																																																																																					to.Ptr("WESTUS")},
		// 																																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																					},
		// 																																																																																					Tier: to.Ptr("Free"),
		// 																																																																																				},
		// 																																																																																				{
		// 																																																																																					Name: to.Ptr("S2"),
		// 																																																																																					Kind: to.Ptr("CustomSpeech"),
		// 																																																																																					Locations: []*string{
		// 																																																																																						to.Ptr("WESTUS")},
		// 																																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																						},
		// 																																																																																						Tier: to.Ptr("Standard"),
		// 																																																																																					},
		// 																																																																																					{
		// 																																																																																						Name: to.Ptr("F0"),
		// 																																																																																						Kind: to.Ptr("ComputerVision"),
		// 																																																																																						Locations: []*string{
		// 																																																																																							to.Ptr("WESTUS")},
		// 																																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																							},
		// 																																																																																							Tier: to.Ptr("Free"),
		// 																																																																																						},
		// 																																																																																						{
		// 																																																																																							Name: to.Ptr("S1"),
		// 																																																																																							Kind: to.Ptr("ComputerVision"),
		// 																																																																																							Locations: []*string{
		// 																																																																																								to.Ptr("WESTUS")},
		// 																																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																								},
		// 																																																																																								Tier: to.Ptr("Standard"),
		// 																																																																																							},
		// 																																																																																							{
		// 																																																																																								Name: to.Ptr("F0"),
		// 																																																																																								Kind: to.Ptr("Face"),
		// 																																																																																								Locations: []*string{
		// 																																																																																									to.Ptr("WESTUS2")},
		// 																																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																									},
		// 																																																																																									Tier: to.Ptr("Free"),
		// 																																																																																								},
		// 																																																																																								{
		// 																																																																																									Name: to.Ptr("S0"),
		// 																																																																																									Kind: to.Ptr("Face"),
		// 																																																																																									Locations: []*string{
		// 																																																																																										to.Ptr("WESTUS2")},
		// 																																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																										},
		// 																																																																																										Tier: to.Ptr("Standard"),
		// 																																																																																									},
		// 																																																																																									{
		// 																																																																																										Name: to.Ptr("F0"),
		// 																																																																																										Kind: to.Ptr("ComputerVision"),
		// 																																																																																										Locations: []*string{
		// 																																																																																											to.Ptr("WESTUS2")},
		// 																																																																																											ResourceType: to.Ptr("accounts"),
		// 																																																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																											},
		// 																																																																																											Tier: to.Ptr("Free"),
		// 																																																																																										},
		// 																																																																																										{
		// 																																																																																											Name: to.Ptr("S1"),
		// 																																																																																											Kind: to.Ptr("ComputerVision"),
		// 																																																																																											Locations: []*string{
		// 																																																																																												to.Ptr("WESTUS2")},
		// 																																																																																												ResourceType: to.Ptr("accounts"),
		// 																																																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																												},
		// 																																																																																												Tier: to.Ptr("Standard"),
		// 																																																																																											},
		// 																																																																																											{
		// 																																																																																												Name: to.Ptr("F0"),
		// 																																																																																												Kind: to.Ptr("ContentModerator"),
		// 																																																																																												Locations: []*string{
		// 																																																																																													to.Ptr("WESTUS2")},
		// 																																																																																													ResourceType: to.Ptr("accounts"),
		// 																																																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																													},
		// 																																																																																													Tier: to.Ptr("Free"),
		// 																																																																																												},
		// 																																																																																												{
		// 																																																																																													Name: to.Ptr("S0"),
		// 																																																																																													Kind: to.Ptr("ContentModerator"),
		// 																																																																																													Locations: []*string{
		// 																																																																																														to.Ptr("WESTUS2")},
		// 																																																																																														ResourceType: to.Ptr("accounts"),
		// 																																																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																														},
		// 																																																																																														Tier: to.Ptr("Standard"),
		// 																																																																																													},
		// 																																																																																													{
		// 																																																																																														Name: to.Ptr("F0"),
		// 																																																																																														Kind: to.Ptr("TextAnalytics"),
		// 																																																																																														Locations: []*string{
		// 																																																																																															to.Ptr("WESTUS2")},
		// 																																																																																															ResourceType: to.Ptr("accounts"),
		// 																																																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																															},
		// 																																																																																															Tier: to.Ptr("Free"),
		// 																																																																																														},
		// 																																																																																														{
		// 																																																																																															Name: to.Ptr("S0"),
		// 																																																																																															Kind: to.Ptr("TextAnalytics"),
		// 																																																																																															Locations: []*string{
		// 																																																																																																to.Ptr("WESTUS2")},
		// 																																																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																},
		// 																																																																																																Tier: to.Ptr("Standard"),
		// 																																																																																															},
		// 																																																																																															{
		// 																																																																																																Name: to.Ptr("S1"),
		// 																																																																																																Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																Locations: []*string{
		// 																																																																																																	to.Ptr("WESTUS2")},
		// 																																																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																	},
		// 																																																																																																	Tier: to.Ptr("Standard"),
		// 																																																																																																},
		// 																																																																																																{
		// 																																																																																																	Name: to.Ptr("S2"),
		// 																																																																																																	Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																	Locations: []*string{
		// 																																																																																																		to.Ptr("WESTUS2")},
		// 																																																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																		},
		// 																																																																																																		Tier: to.Ptr("Standard"),
		// 																																																																																																	},
		// 																																																																																																	{
		// 																																																																																																		Name: to.Ptr("S3"),
		// 																																																																																																		Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																		Locations: []*string{
		// 																																																																																																			to.Ptr("WESTUS2")},
		// 																																																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																			},
		// 																																																																																																			Tier: to.Ptr("Standard"),
		// 																																																																																																		},
		// 																																																																																																		{
		// 																																																																																																			Name: to.Ptr("S4"),
		// 																																																																																																			Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																			Locations: []*string{
		// 																																																																																																				to.Ptr("WESTUS2")},
		// 																																																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																				},
		// 																																																																																																				Tier: to.Ptr("Standard"),
		// 																																																																																																			},
		// 																																																																																																			{
		// 																																																																																																				Name: to.Ptr("F0"),
		// 																																																																																																				Kind: to.Ptr("LUIS"),
		// 																																																																																																				Locations: []*string{
		// 																																																																																																					to.Ptr("WESTUS2")},
		// 																																																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																					},
		// 																																																																																																					Tier: to.Ptr("Free"),
		// 																																																																																																				},
		// 																																																																																																				{
		// 																																																																																																					Name: to.Ptr("S0"),
		// 																																																																																																					Kind: to.Ptr("LUIS"),
		// 																																																																																																					Locations: []*string{
		// 																																																																																																						to.Ptr("WESTUS2")},
		// 																																																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																						},
		// 																																																																																																						Tier: to.Ptr("Standard"),
		// 																																																																																																					},
		// 																																																																																																					{
		// 																																																																																																						Name: to.Ptr("F0"),
		// 																																																																																																						Kind: to.Ptr("Face"),
		// 																																																																																																						Locations: []*string{
		// 																																																																																																							to.Ptr("WESTEUROPE")},
		// 																																																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																							},
		// 																																																																																																							Tier: to.Ptr("Free"),
		// 																																																																																																						},
		// 																																																																																																						{
		// 																																																																																																							Name: to.Ptr("S0"),
		// 																																																																																																							Kind: to.Ptr("Face"),
		// 																																																																																																							Locations: []*string{
		// 																																																																																																								to.Ptr("WESTEUROPE")},
		// 																																																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																								},
		// 																																																																																																								Tier: to.Ptr("Standard"),
		// 																																																																																																							},
		// 																																																																																																							{
		// 																																																																																																								Name: to.Ptr("F0"),
		// 																																																																																																								Kind: to.Ptr("LUIS"),
		// 																																																																																																								Locations: []*string{
		// 																																																																																																									to.Ptr("WESTEUROPE")},
		// 																																																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																									},
		// 																																																																																																									Tier: to.Ptr("Free"),
		// 																																																																																																								},
		// 																																																																																																								{
		// 																																																																																																									Name: to.Ptr("S0"),
		// 																																																																																																									Kind: to.Ptr("LUIS"),
		// 																																																																																																									Locations: []*string{
		// 																																																																																																										to.Ptr("WESTEUROPE")},
		// 																																																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																										},
		// 																																																																																																										Tier: to.Ptr("Standard"),
		// 																																																																																																									},
		// 																																																																																																									{
		// 																																																																																																										Name: to.Ptr("F0"),
		// 																																																																																																										Kind: to.Ptr("ContentModerator"),
		// 																																																																																																										Locations: []*string{
		// 																																																																																																											to.Ptr("WESTEUROPE")},
		// 																																																																																																											ResourceType: to.Ptr("accounts"),
		// 																																																																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																											},
		// 																																																																																																											Tier: to.Ptr("Free"),
		// 																																																																																																										},
		// 																																																																																																										{
		// 																																																																																																											Name: to.Ptr("S0"),
		// 																																																																																																											Kind: to.Ptr("ContentModerator"),
		// 																																																																																																											Locations: []*string{
		// 																																																																																																												to.Ptr("WESTEUROPE")},
		// 																																																																																																												ResourceType: to.Ptr("accounts"),
		// 																																																																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																												},
		// 																																																																																																												Tier: to.Ptr("Standard"),
		// 																																																																																																											},
		// 																																																																																																											{
		// 																																																																																																												Name: to.Ptr("F0"),
		// 																																																																																																												Kind: to.Ptr("ComputerVision"),
		// 																																																																																																												Locations: []*string{
		// 																																																																																																													to.Ptr("WESTEUROPE")},
		// 																																																																																																													ResourceType: to.Ptr("accounts"),
		// 																																																																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																													},
		// 																																																																																																													Tier: to.Ptr("Free"),
		// 																																																																																																												},
		// 																																																																																																												{
		// 																																																																																																													Name: to.Ptr("S1"),
		// 																																																																																																													Kind: to.Ptr("ComputerVision"),
		// 																																																																																																													Locations: []*string{
		// 																																																																																																														to.Ptr("WESTEUROPE")},
		// 																																																																																																														ResourceType: to.Ptr("accounts"),
		// 																																																																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																														},
		// 																																																																																																														Tier: to.Ptr("Standard"),
		// 																																																																																																													},
		// 																																																																																																													{
		// 																																																																																																														Name: to.Ptr("F0"),
		// 																																																																																																														Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																														Locations: []*string{
		// 																																																																																																															to.Ptr("WESTEUROPE")},
		// 																																																																																																															ResourceType: to.Ptr("accounts"),
		// 																																																																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																															},
		// 																																																																																																															Tier: to.Ptr("Free"),
		// 																																																																																																														},
		// 																																																																																																														{
		// 																																																																																																															Name: to.Ptr("S0"),
		// 																																																																																																															Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																															Locations: []*string{
		// 																																																																																																																to.Ptr("WESTEUROPE")},
		// 																																																																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																},
		// 																																																																																																																Tier: to.Ptr("Standard"),
		// 																																																																																																															},
		// 																																																																																																															{
		// 																																																																																																																Name: to.Ptr("S1"),
		// 																																																																																																																Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																Locations: []*string{
		// 																																																																																																																	to.Ptr("WESTEUROPE")},
		// 																																																																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																	},
		// 																																																																																																																	Tier: to.Ptr("Standard"),
		// 																																																																																																																},
		// 																																																																																																																{
		// 																																																																																																																	Name: to.Ptr("S2"),
		// 																																																																																																																	Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																	Locations: []*string{
		// 																																																																																																																		to.Ptr("WESTEUROPE")},
		// 																																																																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																		},
		// 																																																																																																																		Tier: to.Ptr("Standard"),
		// 																																																																																																																	},
		// 																																																																																																																	{
		// 																																																																																																																		Name: to.Ptr("S3"),
		// 																																																																																																																		Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																		Locations: []*string{
		// 																																																																																																																			to.Ptr("WESTEUROPE")},
		// 																																																																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																			},
		// 																																																																																																																			Tier: to.Ptr("Standard"),
		// 																																																																																																																		},
		// 																																																																																																																		{
		// 																																																																																																																			Name: to.Ptr("S4"),
		// 																																																																																																																			Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																			Locations: []*string{
		// 																																																																																																																				to.Ptr("WESTEUROPE")},
		// 																																																																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																				},
		// 																																																																																																																				Tier: to.Ptr("Standard"),
		// 																																																																																																																			},
		// 																																																																																																																			{
		// 																																																																																																																				Name: to.Ptr("F0"),
		// 																																																																																																																				Kind: to.Ptr("Face"),
		// 																																																																																																																				Locations: []*string{
		// 																																																																																																																					to.Ptr("NORTHEUROPE")},
		// 																																																																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																					},
		// 																																																																																																																					Tier: to.Ptr("Free"),
		// 																																																																																																																				},
		// 																																																																																																																				{
		// 																																																																																																																					Name: to.Ptr("S0"),
		// 																																																																																																																					Kind: to.Ptr("Face"),
		// 																																																																																																																					Locations: []*string{
		// 																																																																																																																						to.Ptr("NORTHEUROPE")},
		// 																																																																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																						},
		// 																																																																																																																						Tier: to.Ptr("Standard"),
		// 																																																																																																																					},
		// 																																																																																																																					{
		// 																																																																																																																						Name: to.Ptr("F0"),
		// 																																																																																																																						Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																						Locations: []*string{
		// 																																																																																																																							to.Ptr("NORTHEUROPE")},
		// 																																																																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																							},
		// 																																																																																																																							Tier: to.Ptr("Free"),
		// 																																																																																																																						},
		// 																																																																																																																						{
		// 																																																																																																																							Name: to.Ptr("S1"),
		// 																																																																																																																							Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																							Locations: []*string{
		// 																																																																																																																								to.Ptr("NORTHEUROPE")},
		// 																																																																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																								},
		// 																																																																																																																								Tier: to.Ptr("Standard"),
		// 																																																																																																																							},
		// 																																																																																																																							{
		// 																																																																																																																								Name: to.Ptr("F0"),
		// 																																																																																																																								Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																								Locations: []*string{
		// 																																																																																																																									to.Ptr("NORTHEUROPE")},
		// 																																																																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																									},
		// 																																																																																																																									Tier: to.Ptr("Free"),
		// 																																																																																																																								},
		// 																																																																																																																								{
		// 																																																																																																																									Name: to.Ptr("S0"),
		// 																																																																																																																									Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																									Locations: []*string{
		// 																																																																																																																										to.Ptr("NORTHEUROPE")},
		// 																																																																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																										},
		// 																																																																																																																										Tier: to.Ptr("Standard"),
		// 																																																																																																																									},
		// 																																																																																																																									{
		// 																																																																																																																										Name: to.Ptr("F0"),
		// 																																																																																																																										Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																										Locations: []*string{
		// 																																																																																																																											to.Ptr("NORTHEUROPE")},
		// 																																																																																																																											ResourceType: to.Ptr("accounts"),
		// 																																																																																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																											},
		// 																																																																																																																											Tier: to.Ptr("Free"),
		// 																																																																																																																										},
		// 																																																																																																																										{
		// 																																																																																																																											Name: to.Ptr("S0"),
		// 																																																																																																																											Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																											Locations: []*string{
		// 																																																																																																																												to.Ptr("NORTHEUROPE")},
		// 																																																																																																																												ResourceType: to.Ptr("accounts"),
		// 																																																																																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																												},
		// 																																																																																																																												Tier: to.Ptr("Standard"),
		// 																																																																																																																											},
		// 																																																																																																																											{
		// 																																																																																																																												Name: to.Ptr("S1"),
		// 																																																																																																																												Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																												Locations: []*string{
		// 																																																																																																																													to.Ptr("NORTHEUROPE")},
		// 																																																																																																																													ResourceType: to.Ptr("accounts"),
		// 																																																																																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																													},
		// 																																																																																																																													Tier: to.Ptr("Standard"),
		// 																																																																																																																												},
		// 																																																																																																																												{
		// 																																																																																																																													Name: to.Ptr("S2"),
		// 																																																																																																																													Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																													Locations: []*string{
		// 																																																																																																																														to.Ptr("NORTHEUROPE")},
		// 																																																																																																																														ResourceType: to.Ptr("accounts"),
		// 																																																																																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																														},
		// 																																																																																																																														Tier: to.Ptr("Standard"),
		// 																																																																																																																													},
		// 																																																																																																																													{
		// 																																																																																																																														Name: to.Ptr("S3"),
		// 																																																																																																																														Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																														Locations: []*string{
		// 																																																																																																																															to.Ptr("NORTHEUROPE")},
		// 																																																																																																																															ResourceType: to.Ptr("accounts"),
		// 																																																																																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																															},
		// 																																																																																																																															Tier: to.Ptr("Standard"),
		// 																																																																																																																														},
		// 																																																																																																																														{
		// 																																																																																																																															Name: to.Ptr("S4"),
		// 																																																																																																																															Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																															Locations: []*string{
		// 																																																																																																																																to.Ptr("NORTHEUROPE")},
		// 																																																																																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																},
		// 																																																																																																																																Tier: to.Ptr("Standard"),
		// 																																																																																																																															},
		// 																																																																																																																															{
		// 																																																																																																																																Name: to.Ptr("F0"),
		// 																																																																																																																																Kind: to.Ptr("LUIS"),
		// 																																																																																																																																Locations: []*string{
		// 																																																																																																																																	to.Ptr("NORTHEUROPE")},
		// 																																																																																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																	},
		// 																																																																																																																																	Tier: to.Ptr("Free"),
		// 																																																																																																																																},
		// 																																																																																																																																{
		// 																																																																																																																																	Name: to.Ptr("S0"),
		// 																																																																																																																																	Kind: to.Ptr("LUIS"),
		// 																																																																																																																																	Locations: []*string{
		// 																																																																																																																																		to.Ptr("NORTHEUROPE")},
		// 																																																																																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																		},
		// 																																																																																																																																		Tier: to.Ptr("Standard"),
		// 																																																																																																																																	},
		// 																																																																																																																																	{
		// 																																																																																																																																		Name: to.Ptr("F0"),
		// 																																																																																																																																		Kind: to.Ptr("Face"),
		// 																																																																																																																																		Locations: []*string{
		// 																																																																																																																																			to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																			},
		// 																																																																																																																																			Tier: to.Ptr("Free"),
		// 																																																																																																																																		},
		// 																																																																																																																																		{
		// 																																																																																																																																			Name: to.Ptr("S0"),
		// 																																																																																																																																			Kind: to.Ptr("Face"),
		// 																																																																																																																																			Locations: []*string{
		// 																																																																																																																																				to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																				},
		// 																																																																																																																																				Tier: to.Ptr("Standard"),
		// 																																																																																																																																			},
		// 																																																																																																																																			{
		// 																																																																																																																																				Name: to.Ptr("F0"),
		// 																																																																																																																																				Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																				Locations: []*string{
		// 																																																																																																																																					to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																					},
		// 																																																																																																																																					Tier: to.Ptr("Free"),
		// 																																																																																																																																				},
		// 																																																																																																																																				{
		// 																																																																																																																																					Name: to.Ptr("S0"),
		// 																																																																																																																																					Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																					Locations: []*string{
		// 																																																																																																																																						to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																						},
		// 																																																																																																																																						Tier: to.Ptr("Standard"),
		// 																																																																																																																																					},
		// 																																																																																																																																					{
		// 																																																																																																																																						Name: to.Ptr("F0"),
		// 																																																																																																																																						Kind: to.Ptr("LUIS"),
		// 																																																																																																																																						Locations: []*string{
		// 																																																																																																																																							to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																							},
		// 																																																																																																																																							Tier: to.Ptr("Free"),
		// 																																																																																																																																						},
		// 																																																																																																																																						{
		// 																																																																																																																																							Name: to.Ptr("S0"),
		// 																																																																																																																																							Kind: to.Ptr("LUIS"),
		// 																																																																																																																																							Locations: []*string{
		// 																																																																																																																																								to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																								},
		// 																																																																																																																																								Tier: to.Ptr("Standard"),
		// 																																																																																																																																							},
		// 																																																																																																																																							{
		// 																																																																																																																																								Name: to.Ptr("F0"),
		// 																																																																																																																																								Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																								Locations: []*string{
		// 																																																																																																																																									to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																									},
		// 																																																																																																																																									Tier: to.Ptr("Free"),
		// 																																																																																																																																								},
		// 																																																																																																																																								{
		// 																																																																																																																																									Name: to.Ptr("S1"),
		// 																																																																																																																																									Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																									Locations: []*string{
		// 																																																																																																																																										to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																										},
		// 																																																																																																																																										Tier: to.Ptr("Standard"),
		// 																																																																																																																																									},
		// 																																																																																																																																									{
		// 																																																																																																																																										Name: to.Ptr("F0"),
		// 																																																																																																																																										Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																										Locations: []*string{
		// 																																																																																																																																											to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																											ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																											},
		// 																																																																																																																																											Tier: to.Ptr("Free"),
		// 																																																																																																																																										},
		// 																																																																																																																																										{
		// 																																																																																																																																											Name: to.Ptr("S0"),
		// 																																																																																																																																											Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																											Locations: []*string{
		// 																																																																																																																																												to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																												ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																												},
		// 																																																																																																																																												Tier: to.Ptr("Standard"),
		// 																																																																																																																																											},
		// 																																																																																																																																											{
		// 																																																																																																																																												Name: to.Ptr("S1"),
		// 																																																																																																																																												Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																												Locations: []*string{
		// 																																																																																																																																													to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																													ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																													},
		// 																																																																																																																																													Tier: to.Ptr("Standard"),
		// 																																																																																																																																												},
		// 																																																																																																																																												{
		// 																																																																																																																																													Name: to.Ptr("S2"),
		// 																																																																																																																																													Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																													Locations: []*string{
		// 																																																																																																																																														to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																														ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																														},
		// 																																																																																																																																														Tier: to.Ptr("Standard"),
		// 																																																																																																																																													},
		// 																																																																																																																																													{
		// 																																																																																																																																														Name: to.Ptr("S3"),
		// 																																																																																																																																														Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																														Locations: []*string{
		// 																																																																																																																																															to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																															ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																															},
		// 																																																																																																																																															Tier: to.Ptr("Standard"),
		// 																																																																																																																																														},
		// 																																																																																																																																														{
		// 																																																																																																																																															Name: to.Ptr("S4"),
		// 																																																																																																																																															Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																															Locations: []*string{
		// 																																																																																																																																																to.Ptr("SOUTHEASTASIA")},
		// 																																																																																																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																},
		// 																																																																																																																																																Tier: to.Ptr("Standard"),
		// 																																																																																																																																															},
		// 																																																																																																																																															{
		// 																																																																																																																																																Name: to.Ptr("F0"),
		// 																																																																																																																																																Kind: to.Ptr("Face"),
		// 																																																																																																																																																Locations: []*string{
		// 																																																																																																																																																	to.Ptr("EASTASIA")},
		// 																																																																																																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																	},
		// 																																																																																																																																																	Tier: to.Ptr("Free"),
		// 																																																																																																																																																},
		// 																																																																																																																																																{
		// 																																																																																																																																																	Name: to.Ptr("S0"),
		// 																																																																																																																																																	Kind: to.Ptr("Face"),
		// 																																																																																																																																																	Locations: []*string{
		// 																																																																																																																																																		to.Ptr("EASTASIA")},
		// 																																																																																																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																		},
		// 																																																																																																																																																		Tier: to.Ptr("Standard"),
		// 																																																																																																																																																	},
		// 																																																																																																																																																	{
		// 																																																																																																																																																		Name: to.Ptr("F0"),
		// 																																																																																																																																																		Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																		Locations: []*string{
		// 																																																																																																																																																			to.Ptr("EASTASIA")},
		// 																																																																																																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																			},
		// 																																																																																																																																																			Tier: to.Ptr("Free"),
		// 																																																																																																																																																		},
		// 																																																																																																																																																		{
		// 																																																																																																																																																			Name: to.Ptr("S1"),
		// 																																																																																																																																																			Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																			Locations: []*string{
		// 																																																																																																																																																				to.Ptr("EASTASIA")},
		// 																																																																																																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																				},
		// 																																																																																																																																																				Tier: to.Ptr("Standard"),
		// 																																																																																																																																																			},
		// 																																																																																																																																																			{
		// 																																																																																																																																																				Name: to.Ptr("F0"),
		// 																																																																																																																																																				Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																				Locations: []*string{
		// 																																																																																																																																																					to.Ptr("EASTASIA")},
		// 																																																																																																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																					},
		// 																																																																																																																																																					Tier: to.Ptr("Free"),
		// 																																																																																																																																																				},
		// 																																																																																																																																																				{
		// 																																																																																																																																																					Name: to.Ptr("S0"),
		// 																																																																																																																																																					Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																					Locations: []*string{
		// 																																																																																																																																																						to.Ptr("EASTASIA")},
		// 																																																																																																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																						},
		// 																																																																																																																																																						Tier: to.Ptr("Standard"),
		// 																																																																																																																																																					},
		// 																																																																																																																																																					{
		// 																																																																																																																																																						Name: to.Ptr("F0"),
		// 																																																																																																																																																						Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																						Locations: []*string{
		// 																																																																																																																																																							to.Ptr("EASTASIA")},
		// 																																																																																																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																							},
		// 																																																																																																																																																							Tier: to.Ptr("Free"),
		// 																																																																																																																																																						},
		// 																																																																																																																																																						{
		// 																																																																																																																																																							Name: to.Ptr("S0"),
		// 																																																																																																																																																							Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																							Locations: []*string{
		// 																																																																																																																																																								to.Ptr("EASTASIA")},
		// 																																																																																																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																								},
		// 																																																																																																																																																								Tier: to.Ptr("Standard"),
		// 																																																																																																																																																							},
		// 																																																																																																																																																							{
		// 																																																																																																																																																								Name: to.Ptr("S1"),
		// 																																																																																																																																																								Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																								Locations: []*string{
		// 																																																																																																																																																									to.Ptr("EASTASIA")},
		// 																																																																																																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																									},
		// 																																																																																																																																																									Tier: to.Ptr("Standard"),
		// 																																																																																																																																																								},
		// 																																																																																																																																																								{
		// 																																																																																																																																																									Name: to.Ptr("S2"),
		// 																																																																																																																																																									Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																									Locations: []*string{
		// 																																																																																																																																																										to.Ptr("EASTASIA")},
		// 																																																																																																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																										},
		// 																																																																																																																																																										Tier: to.Ptr("Standard"),
		// 																																																																																																																																																									},
		// 																																																																																																																																																									{
		// 																																																																																																																																																										Name: to.Ptr("S3"),
		// 																																																																																																																																																										Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																										Locations: []*string{
		// 																																																																																																																																																											to.Ptr("EASTASIA")},
		// 																																																																																																																																																											ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																											},
		// 																																																																																																																																																											Tier: to.Ptr("Standard"),
		// 																																																																																																																																																										},
		// 																																																																																																																																																										{
		// 																																																																																																																																																											Name: to.Ptr("S4"),
		// 																																																																																																																																																											Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																											Locations: []*string{
		// 																																																																																																																																																												to.Ptr("EASTASIA")},
		// 																																																																																																																																																												ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																												},
		// 																																																																																																																																																												Tier: to.Ptr("Standard"),
		// 																																																																																																																																																											},
		// 																																																																																																																																																											{
		// 																																																																																																																																																												Name: to.Ptr("F0"),
		// 																																																																																																																																																												Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																												Locations: []*string{
		// 																																																																																																																																																													to.Ptr("EASTASIA")},
		// 																																																																																																																																																													ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																													},
		// 																																																																																																																																																													Tier: to.Ptr("Free"),
		// 																																																																																																																																																												},
		// 																																																																																																																																																												{
		// 																																																																																																																																																													Name: to.Ptr("S0"),
		// 																																																																																																																																																													Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																													Locations: []*string{
		// 																																																																																																																																																														to.Ptr("EASTASIA")},
		// 																																																																																																																																																														ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																														},
		// 																																																																																																																																																														Tier: to.Ptr("Standard"),
		// 																																																																																																																																																													},
		// 																																																																																																																																																													{
		// 																																																																																																																																																														Name: to.Ptr("F0"),
		// 																																																																																																																																																														Kind: to.Ptr("Face"),
		// 																																																																																																																																																														Locations: []*string{
		// 																																																																																																																																																															to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																															ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																															},
		// 																																																																																																																																																															Tier: to.Ptr("Free"),
		// 																																																																																																																																																														},
		// 																																																																																																																																																														{
		// 																																																																																																																																																															Name: to.Ptr("S0"),
		// 																																																																																																																																																															Kind: to.Ptr("Face"),
		// 																																																																																																																																																															Locations: []*string{
		// 																																																																																																																																																																to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																},
		// 																																																																																																																																																																Tier: to.Ptr("Standard"),
		// 																																																																																																																																																															},
		// 																																																																																																																																																															{
		// 																																																																																																																																																																Name: to.Ptr("F0"),
		// 																																																																																																																																																																Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																																Locations: []*string{
		// 																																																																																																																																																																	to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																	},
		// 																																																																																																																																																																	Tier: to.Ptr("Free"),
		// 																																																																																																																																																																},
		// 																																																																																																																																																																{
		// 																																																																																																																																																																	Name: to.Ptr("S0"),
		// 																																																																																																																																																																	Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																																	Locations: []*string{
		// 																																																																																																																																																																		to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																		},
		// 																																																																																																																																																																		Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																	},
		// 																																																																																																																																																																	{
		// 																																																																																																																																																																		Name: to.Ptr("F0"),
		// 																																																																																																																																																																		Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																																		Locations: []*string{
		// 																																																																																																																																																																			to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																			},
		// 																																																																																																																																																																			Tier: to.Ptr("Free"),
		// 																																																																																																																																																																		},
		// 																																																																																																																																																																		{
		// 																																																																																																																																																																			Name: to.Ptr("S0"),
		// 																																																																																																																																																																			Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																																			Locations: []*string{
		// 																																																																																																																																																																				to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																				},
		// 																																																																																																																																																																				Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																			},
		// 																																																																																																																																																																			{
		// 																																																																																																																																																																				Name: to.Ptr("F0"),
		// 																																																																																																																																																																				Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																																				Locations: []*string{
		// 																																																																																																																																																																					to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																					},
		// 																																																																																																																																																																					Tier: to.Ptr("Free"),
		// 																																																																																																																																																																				},
		// 																																																																																																																																																																				{
		// 																																																																																																																																																																					Name: to.Ptr("S1"),
		// 																																																																																																																																																																					Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																																					Locations: []*string{
		// 																																																																																																																																																																						to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																						},
		// 																																																																																																																																																																						Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																					},
		// 																																																																																																																																																																					{
		// 																																																																																																																																																																						Name: to.Ptr("F0"),
		// 																																																																																																																																																																						Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																						Locations: []*string{
		// 																																																																																																																																																																							to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																							},
		// 																																																																																																																																																																							Tier: to.Ptr("Free"),
		// 																																																																																																																																																																						},
		// 																																																																																																																																																																						{
		// 																																																																																																																																																																							Name: to.Ptr("S0"),
		// 																																																																																																																																																																							Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																							Locations: []*string{
		// 																																																																																																																																																																								to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																								},
		// 																																																																																																																																																																								Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																							},
		// 																																																																																																																																																																							{
		// 																																																																																																																																																																								Name: to.Ptr("S1"),
		// 																																																																																																																																																																								Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																								Locations: []*string{
		// 																																																																																																																																																																									to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																									},
		// 																																																																																																																																																																									Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																								},
		// 																																																																																																																																																																								{
		// 																																																																																																																																																																									Name: to.Ptr("S2"),
		// 																																																																																																																																																																									Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																									Locations: []*string{
		// 																																																																																																																																																																										to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																										},
		// 																																																																																																																																																																										Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																									},
		// 																																																																																																																																																																									{
		// 																																																																																																																																																																										Name: to.Ptr("S3"),
		// 																																																																																																																																																																										Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																										Locations: []*string{
		// 																																																																																																																																																																											to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																											ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																											},
		// 																																																																																																																																																																											Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																										},
		// 																																																																																																																																																																										{
		// 																																																																																																																																																																											Name: to.Ptr("S4"),
		// 																																																																																																																																																																											Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																											Locations: []*string{
		// 																																																																																																																																																																												to.Ptr("WESTCENTRALUS")},
		// 																																																																																																																																																																												ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																												},
		// 																																																																																																																																																																												Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																											},
		// 																																																																																																																																																																											{
		// 																																																																																																																																																																												Name: to.Ptr("F0"),
		// 																																																																																																																																																																												Kind: to.Ptr("Face"),
		// 																																																																																																																																																																												Locations: []*string{
		// 																																																																																																																																																																													to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																													ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																													},
		// 																																																																																																																																																																													Tier: to.Ptr("Free"),
		// 																																																																																																																																																																												},
		// 																																																																																																																																																																												{
		// 																																																																																																																																																																													Name: to.Ptr("S0"),
		// 																																																																																																																																																																													Kind: to.Ptr("Face"),
		// 																																																																																																																																																																													Locations: []*string{
		// 																																																																																																																																																																														to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																														ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																														},
		// 																																																																																																																																																																														Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																													},
		// 																																																																																																																																																																													{
		// 																																																																																																																																																																														Name: to.Ptr("F0"),
		// 																																																																																																																																																																														Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																																														Locations: []*string{
		// 																																																																																																																																																																															to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																															ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																															},
		// 																																																																																																																																																																															Tier: to.Ptr("Free"),
		// 																																																																																																																																																																														},
		// 																																																																																																																																																																														{
		// 																																																																																																																																																																															Name: to.Ptr("S1"),
		// 																																																																																																																																																																															Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																																															Locations: []*string{
		// 																																																																																																																																																																																to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																},
		// 																																																																																																																																																																																Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																															},
		// 																																																																																																																																																																															{
		// 																																																																																																																																																																																Name: to.Ptr("F0"),
		// 																																																																																																																																																																																Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																																																Locations: []*string{
		// 																																																																																																																																																																																	to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																	},
		// 																																																																																																																																																																																	Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																},
		// 																																																																																																																																																																																{
		// 																																																																																																																																																																																	Name: to.Ptr("S0"),
		// 																																																																																																																																																																																	Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																																																	Locations: []*string{
		// 																																																																																																																																																																																		to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																		},
		// 																																																																																																																																																																																		Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																	},
		// 																																																																																																																																																																																	{
		// 																																																																																																																																																																																		Name: to.Ptr("F0"),
		// 																																																																																																																																																																																		Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																		Locations: []*string{
		// 																																																																																																																																																																																			to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																			},
		// 																																																																																																																																																																																			Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																		},
		// 																																																																																																																																																																																		{
		// 																																																																																																																																																																																			Name: to.Ptr("S0"),
		// 																																																																																																																																																																																			Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																			Locations: []*string{
		// 																																																																																																																																																																																				to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																				},
		// 																																																																																																																																																																																				Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																			},
		// 																																																																																																																																																																																			{
		// 																																																																																																																																																																																				Name: to.Ptr("S1"),
		// 																																																																																																																																																																																				Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																				Locations: []*string{
		// 																																																																																																																																																																																					to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																					},
		// 																																																																																																																																																																																					Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																				},
		// 																																																																																																																																																																																				{
		// 																																																																																																																																																																																					Name: to.Ptr("S2"),
		// 																																																																																																																																																																																					Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																					Locations: []*string{
		// 																																																																																																																																																																																						to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																						},
		// 																																																																																																																																																																																						Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																					},
		// 																																																																																																																																																																																					{
		// 																																																																																																																																																																																						Name: to.Ptr("S3"),
		// 																																																																																																																																																																																						Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																						Locations: []*string{
		// 																																																																																																																																																																																							to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																							},
		// 																																																																																																																																																																																							Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																						},
		// 																																																																																																																																																																																						{
		// 																																																																																																																																																																																							Name: to.Ptr("S4"),
		// 																																																																																																																																																																																							Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																							Locations: []*string{
		// 																																																																																																																																																																																								to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																								},
		// 																																																																																																																																																																																								Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																							},
		// 																																																																																																																																																																																							{
		// 																																																																																																																																																																																								Name: to.Ptr("F0"),
		// 																																																																																																																																																																																								Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																																																								Locations: []*string{
		// 																																																																																																																																																																																									to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																									},
		// 																																																																																																																																																																																									Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																								},
		// 																																																																																																																																																																																								{
		// 																																																																																																																																																																																									Name: to.Ptr("S0"),
		// 																																																																																																																																																																																									Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																																																									Locations: []*string{
		// 																																																																																																																																																																																										to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																										},
		// 																																																																																																																																																																																										Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																									},
		// 																																																																																																																																																																																									{
		// 																																																																																																																																																																																										Name: to.Ptr("F0"),
		// 																																																																																																																																																																																										Kind: to.Ptr("CustomVision.Training"),
		// 																																																																																																																																																																																										Locations: []*string{
		// 																																																																																																																																																																																											to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																											ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																											},
		// 																																																																																																																																																																																											Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																										},
		// 																																																																																																																																																																																										{
		// 																																																																																																																																																																																											Name: to.Ptr("S0"),
		// 																																																																																																																																																																																											Kind: to.Ptr("CustomVision.Training"),
		// 																																																																																																																																																																																											Locations: []*string{
		// 																																																																																																																																																																																												to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																												ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																												},
		// 																																																																																																																																																																																												Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																											},
		// 																																																																																																																																																																																											{
		// 																																																																																																																																																																																												Name: to.Ptr("F0"),
		// 																																																																																																																																																																																												Kind: to.Ptr("CustomVision.Prediction"),
		// 																																																																																																																																																																																												Locations: []*string{
		// 																																																																																																																																																																																													to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																													ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																													},
		// 																																																																																																																																																																																													Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																												},
		// 																																																																																																																																																																																												{
		// 																																																																																																																																																																																													Name: to.Ptr("S0"),
		// 																																																																																																																																																																																													Kind: to.Ptr("CustomVision.Prediction"),
		// 																																																																																																																																																																																													Locations: []*string{
		// 																																																																																																																																																																																														to.Ptr("SOUTHCENTRALUS")},
		// 																																																																																																																																																																																														ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																														},
		// 																																																																																																																																																																																														Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																													},
		// 																																																																																																																																																																																													{
		// 																																																																																																																																																																																														Name: to.Ptr("F0"),
		// 																																																																																																																																																																																														Kind: to.Ptr("Face"),
		// 																																																																																																																																																																																														Locations: []*string{
		// 																																																																																																																																																																																															to.Ptr("EASTUS")},
		// 																																																																																																																																																																																															ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																															},
		// 																																																																																																																																																																																															Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																														},
		// 																																																																																																																																																																																														{
		// 																																																																																																																																																																																															Name: to.Ptr("S0"),
		// 																																																																																																																																																																																															Kind: to.Ptr("Face"),
		// 																																																																																																																																																																																															Locations: []*string{
		// 																																																																																																																																																																																																to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																},
		// 																																																																																																																																																																																																Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																															},
		// 																																																																																																																																																																																															{
		// 																																																																																																																																																																																																Name: to.Ptr("F0"),
		// 																																																																																																																																																																																																Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																																																																Locations: []*string{
		// 																																																																																																																																																																																																	to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																	},
		// 																																																																																																																																																																																																	Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																																},
		// 																																																																																																																																																																																																{
		// 																																																																																																																																																																																																	Name: to.Ptr("S1"),
		// 																																																																																																																																																																																																	Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																																																																	Locations: []*string{
		// 																																																																																																																																																																																																		to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																		},
		// 																																																																																																																																																																																																		Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																	},
		// 																																																																																																																																																																																																	{
		// 																																																																																																																																																																																																		Name: to.Ptr("F0"),
		// 																																																																																																																																																																																																		Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																																																																		Locations: []*string{
		// 																																																																																																																																																																																																			to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																			},
		// 																																																																																																																																																																																																			Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																																		},
		// 																																																																																																																																																																																																		{
		// 																																																																																																																																																																																																			Name: to.Ptr("S0"),
		// 																																																																																																																																																																																																			Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																																																																			Locations: []*string{
		// 																																																																																																																																																																																																				to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																				},
		// 																																																																																																																																																																																																				Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																			},
		// 																																																																																																																																																																																																			{
		// 																																																																																																																																																																																																				Name: to.Ptr("F0"),
		// 																																																																																																																																																																																																				Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																				Locations: []*string{
		// 																																																																																																																																																																																																					to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																					},
		// 																																																																																																																																																																																																					Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																																				},
		// 																																																																																																																																																																																																				{
		// 																																																																																																																																																																																																					Name: to.Ptr("S0"),
		// 																																																																																																																																																																																																					Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																					Locations: []*string{
		// 																																																																																																																																																																																																						to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																						},
		// 																																																																																																																																																																																																						Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																					},
		// 																																																																																																																																																																																																					{
		// 																																																																																																																																																																																																						Name: to.Ptr("S1"),
		// 																																																																																																																																																																																																						Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																						Locations: []*string{
		// 																																																																																																																																																																																																							to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																							},
		// 																																																																																																																																																																																																							Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																						},
		// 																																																																																																																																																																																																						{
		// 																																																																																																																																																																																																							Name: to.Ptr("S2"),
		// 																																																																																																																																																																																																							Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																							Locations: []*string{
		// 																																																																																																																																																																																																								to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																								},
		// 																																																																																																																																																																																																								Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																							},
		// 																																																																																																																																																																																																							{
		// 																																																																																																																																																																																																								Name: to.Ptr("S3"),
		// 																																																																																																																																																																																																								Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																								Locations: []*string{
		// 																																																																																																																																																																																																									to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																									},
		// 																																																																																																																																																																																																									Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																								},
		// 																																																																																																																																																																																																								{
		// 																																																																																																																																																																																																									Name: to.Ptr("S4"),
		// 																																																																																																																																																																																																									Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																									Locations: []*string{
		// 																																																																																																																																																																																																										to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																										},
		// 																																																																																																																																																																																																										Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																									},
		// 																																																																																																																																																																																																									{
		// 																																																																																																																																																																																																										Name: to.Ptr("F0"),
		// 																																																																																																																																																																																																										Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																																																																										Locations: []*string{
		// 																																																																																																																																																																																																											to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																											ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																											Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																											},
		// 																																																																																																																																																																																																											Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																																										},
		// 																																																																																																																																																																																																										{
		// 																																																																																																																																																																																																											Name: to.Ptr("S0"),
		// 																																																																																																																																																																																																											Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																																																																											Locations: []*string{
		// 																																																																																																																																																																																																												to.Ptr("EASTUS")},
		// 																																																																																																																																																																																																												ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																												Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																												},
		// 																																																																																																																																																																																																												Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																											},
		// 																																																																																																																																																																																																											{
		// 																																																																																																																																																																																																												Name: to.Ptr("F0"),
		// 																																																																																																																																																																																																												Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																																																																												Locations: []*string{
		// 																																																																																																																																																																																																													to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																													ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																													Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																													},
		// 																																																																																																																																																																																																													Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																																												},
		// 																																																																																																																																																																																																												{
		// 																																																																																																																																																																																																													Name: to.Ptr("S0"),
		// 																																																																																																																																																																																																													Kind: to.Ptr("ContentModerator"),
		// 																																																																																																																																																																																																													Locations: []*string{
		// 																																																																																																																																																																																																														to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																														ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																														Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																														},
		// 																																																																																																																																																																																																														Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																													},
		// 																																																																																																																																																																																																													{
		// 																																																																																																																																																																																																														Name: to.Ptr("F0"),
		// 																																																																																																																																																																																																														Kind: to.Ptr("Face"),
		// 																																																																																																																																																																																																														Locations: []*string{
		// 																																																																																																																																																																																																															to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																															ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																															Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																															},
		// 																																																																																																																																																																																																															Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																																														},
		// 																																																																																																																																																																																																														{
		// 																																																																																																																																																																																																															Name: to.Ptr("S0"),
		// 																																																																																																																																																																																																															Kind: to.Ptr("Face"),
		// 																																																																																																																																																																																																															Locations: []*string{
		// 																																																																																																																																																																																																																to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																},
		// 																																																																																																																																																																																																																Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																															},
		// 																																																																																																																																																																																																															{
		// 																																																																																																																																																																																																																Name: to.Ptr("F0"),
		// 																																																																																																																																																																																																																Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																																																																																Locations: []*string{
		// 																																																																																																																																																																																																																	to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																	ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																	Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																	},
		// 																																																																																																																																																																																																																	Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																																																},
		// 																																																																																																																																																																																																																{
		// 																																																																																																																																																																																																																	Name: to.Ptr("S0"),
		// 																																																																																																																																																																																																																	Kind: to.Ptr("LUIS"),
		// 																																																																																																																																																																																																																	Locations: []*string{
		// 																																																																																																																																																																																																																		to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																		ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																		Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																		},
		// 																																																																																																																																																																																																																		Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																																	},
		// 																																																																																																																																																																																																																	{
		// 																																																																																																																																																																																																																		Name: to.Ptr("F0"),
		// 																																																																																																																																																																																																																		Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																																																																																		Locations: []*string{
		// 																																																																																																																																																																																																																			to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																			ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																			Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																			},
		// 																																																																																																																																																																																																																			Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																																																		},
		// 																																																																																																																																																																																																																		{
		// 																																																																																																																																																																																																																			Name: to.Ptr("S1"),
		// 																																																																																																																																																																																																																			Kind: to.Ptr("ComputerVision"),
		// 																																																																																																																																																																																																																			Locations: []*string{
		// 																																																																																																																																																																																																																				to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																				ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																				Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																				},
		// 																																																																																																																																																																																																																				Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																																			},
		// 																																																																																																																																																																																																																			{
		// 																																																																																																																																																																																																																				Name: to.Ptr("F0"),
		// 																																																																																																																																																																																																																				Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																																				Locations: []*string{
		// 																																																																																																																																																																																																																					to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																					ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																					Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																					},
		// 																																																																																																																																																																																																																					Tier: to.Ptr("Free"),
		// 																																																																																																																																																																																																																				},
		// 																																																																																																																																																																																																																				{
		// 																																																																																																																																																																																																																					Name: to.Ptr("S0"),
		// 																																																																																																																																																																																																																					Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																																					Locations: []*string{
		// 																																																																																																																																																																																																																						to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																						ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																						Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																						},
		// 																																																																																																																																																																																																																						Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																																					},
		// 																																																																																																																																																																																																																					{
		// 																																																																																																																																																																																																																						Name: to.Ptr("S1"),
		// 																																																																																																																																																																																																																						Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																																						Locations: []*string{
		// 																																																																																																																																																																																																																							to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																							ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																							Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																							},
		// 																																																																																																																																																																																																																							Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																																						},
		// 																																																																																																																																																																																																																						{
		// 																																																																																																																																																																																																																							Name: to.Ptr("S2"),
		// 																																																																																																																																																																																																																							Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																																							Locations: []*string{
		// 																																																																																																																																																																																																																								to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																								ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																								Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																								},
		// 																																																																																																																																																																																																																								Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																																							},
		// 																																																																																																																																																																																																																							{
		// 																																																																																																																																																																																																																								Name: to.Ptr("S3"),
		// 																																																																																																																																																																																																																								Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																																								Locations: []*string{
		// 																																																																																																																																																																																																																									to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																									ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																									Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																									},
		// 																																																																																																																																																																																																																									Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																																								},
		// 																																																																																																																																																																																																																								{
		// 																																																																																																																																																																																																																									Name: to.Ptr("S4"),
		// 																																																																																																																																																																																																																									Kind: to.Ptr("TextAnalytics"),
		// 																																																																																																																																																																																																																									Locations: []*string{
		// 																																																																																																																																																																																																																										to.Ptr("EASTUS2")},
		// 																																																																																																																																																																																																																										ResourceType: to.Ptr("accounts"),
		// 																																																																																																																																																																																																																										Restrictions: []*armcognitiveservices.ResourceSKURestrictions{
		// 																																																																																																																																																																																																																										},
		// 																																																																																																																																																																																																																										Tier: to.Ptr("Standard"),
		// 																																																																																																																																																																																																																								}},
		// 																																																																																																																																																																																																																							}
	}
}
