package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/InformationProtectionPolicies/ListInformationProtectionPolicies_example.json
func ExampleInformationProtectionPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInformationProtectionPoliciesClient().NewListPager("providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e", nil)
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
		// page.InformationProtectionPolicyList = armsecurity.InformationProtectionPolicyList{
		// 	Value: []*armsecurity.InformationProtectionPolicy{
		// 		{
		// 			Name: to.Ptr("effective"),
		// 			Type: to.Ptr("Microsoft.Security/informationProtectionPolicies"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e/providers/Microsoft.Security/informationProtectionPolicies/effective"),
		// 			Properties: &armsecurity.InformationProtectionPolicyProperties{
		// 				InformationTypes: map[string]*armsecurity.InformationType{
		// 					"3bf35491-99b8-41f2-86d5-c1200a7df658": &armsecurity.InformationType{
		// 						Custom: to.Ptr(true),
		// 						DisplayName: to.Ptr("Custom"),
		// 						Enabled: to.Ptr(true),
		// 						Keywords: []*armsecurity.InformationProtectionKeyword{
		// 							{
		// 								CanBeNumeric: to.Ptr(false),
		// 								Custom: to.Ptr(true),
		// 								Pattern: to.Ptr("%networking%"),
		// 						}},
		// 						Order: to.Ptr[int32](1400),
		// 						RecommendedLabelID: to.Ptr("7aa516c7-5a53-4857-bc6e-6808c6acd542"),
		// 					},
		// 					"5856f35c-8e08-4d08-9bf7-87a146150569": &armsecurity.InformationType{
		// 						Custom: to.Ptr(false),
		// 						DisplayName: to.Ptr("Contact Info"),
		// 						Enabled: to.Ptr(true),
		// 						Keywords: []*armsecurity.InformationProtectionKeyword{
		// 							{
		// 								CanBeNumeric: to.Ptr(false),
		// 								Custom: to.Ptr(false),
		// 								Pattern: to.Ptr("%email%"),
		// 							},
		// 							{
		// 								CanBeNumeric: to.Ptr(false),
		// 								Custom: to.Ptr(false),
		// 								Pattern: to.Ptr("%e-mail%"),
		// 							},
		// 							{
		// 								CanBeNumeric: to.Ptr(false),
		// 								Custom: to.Ptr(false),
		// 								Pattern: to.Ptr("%addr%"),
		// 							},
		// 							{
		// 								CanBeNumeric: to.Ptr(true),
		// 								Custom: to.Ptr(false),
		// 								Pattern: to.Ptr("%street%"),
		// 							},
		// 							{
		// 								CanBeNumeric: to.Ptr(false),
		// 								Custom: to.Ptr(false),
		// 								Pattern: to.Ptr("%city%"),
		// 						}},
		// 						Order: to.Ptr[int32](200),
		// 						RecommendedLabelID: to.Ptr("575739d2-3d53-4df0-9042-4c7772d5c7b1"),
		// 					},
		// 					"7fb9419d-2473-4ad8-8e11-b25cc8cf6a07": &armsecurity.InformationType{
		// 						Custom: to.Ptr(false),
		// 						DisplayName: to.Ptr("Networking"),
		// 						Enabled: to.Ptr(true),
		// 						Keywords: []*armsecurity.InformationProtectionKeyword{
		// 							{
		// 								CanBeNumeric: to.Ptr(false),
		// 								Custom: to.Ptr(false),
		// 								Pattern: to.Ptr("%ip%"),
		// 							},
		// 							{
		// 								CanBeNumeric: to.Ptr(false),
		// 								Custom: to.Ptr(false),
		// 								Pattern: to.Ptr("ip%address%"),
		// 							},
		// 							{
		// 								CanBeNumeric: to.Ptr(false),
		// 								Custom: to.Ptr(false),
		// 								Pattern: to.Ptr("%mac%address%"),
		// 							},
		// 							{
		// 								CanBeNumeric: to.Ptr(true),
		// 								Custom: to.Ptr(true),
		// 								Pattern: to.Ptr("%networking%"),
		// 						}},
		// 						Order: to.Ptr[int32](100),
		// 						RecommendedLabelID: to.Ptr("575739d2-3d53-4df0-9042-4c7772d5c7b1"),
		// 					},
		// 				},
		// 				Labels: map[string]*armsecurity.SensitivityLabel{
		// 					"1345da73-bc5a-4a8f-b7dd-3820eb713da8": &armsecurity.SensitivityLabel{
		// 						DisplayName: to.Ptr("Public"),
		// 						Enabled: to.Ptr(true),
		// 						Order: to.Ptr[int32](100),
		// 					},
		// 					"575739d2-3d53-4df0-9042-4c7772d5c7b1": &armsecurity.SensitivityLabel{
		// 						DisplayName: to.Ptr("Confidential"),
		// 						Enabled: to.Ptr(true),
		// 						Order: to.Ptr[int32](300),
		// 					},
		// 					"7aa516c7-5a53-4857-bc6e-6808c6acd542": &armsecurity.SensitivityLabel{
		// 						DisplayName: to.Ptr("General"),
		// 						Enabled: to.Ptr(true),
		// 						Order: to.Ptr[int32](200),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("custom"),
		// 			Type: to.Ptr("Microsoft.Security/informationProtectionPolicies"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e/providers/Microsoft.Security/informationProtectionPolicies/custom"),
		// 			Properties: &armsecurity.InformationProtectionPolicyProperties{
		// 				InformationTypes: map[string]*armsecurity.InformationType{
		// 					"3bf35491-99b8-41f2-86d5-c1200a7df658": &armsecurity.InformationType{
		// 						Custom: to.Ptr(true),
		// 						DisplayName: to.Ptr("Custom"),
		// 						Enabled: to.Ptr(true),
		// 						Keywords: []*armsecurity.InformationProtectionKeyword{
		// 							{
		// 								CanBeNumeric: to.Ptr(true),
		// 								Custom: to.Ptr(true),
		// 								Pattern: to.Ptr("%custom%"),
		// 						}},
		// 						Order: to.Ptr[int32](1400),
		// 						RecommendedLabelID: to.Ptr("7aa516c7-5a53-4857-bc6e-6808c6acd542"),
		// 					},
		// 					"7fb9419d-2473-4ad8-8e11-b25cc8cf6a07": &armsecurity.InformationType{
		// 						Custom: to.Ptr(false),
		// 						DisplayName: to.Ptr("Networking"),
		// 						Enabled: to.Ptr(true),
		// 						Keywords: []*armsecurity.InformationProtectionKeyword{
		// 							{
		// 								CanBeNumeric: to.Ptr(false),
		// 								Custom: to.Ptr(true),
		// 								Pattern: to.Ptr("%networking%"),
		// 						}},
		// 						Order: to.Ptr[int32](100),
		// 						RecommendedLabelID: to.Ptr("575739d2-3d53-4df0-9042-4c7772d5c7b1"),
		// 					},
		// 				},
		// 				Labels: map[string]*armsecurity.SensitivityLabel{
		// 					"1345da73-bc5a-4a8f-b7dd-3820eb713da8": &armsecurity.SensitivityLabel{
		// 						DisplayName: to.Ptr("Public"),
		// 						Enabled: to.Ptr(true),
		// 						Order: to.Ptr[int32](100),
		// 					},
		// 					"575739d2-3d53-4df0-9042-4c7772d5c7b1": &armsecurity.SensitivityLabel{
		// 						DisplayName: to.Ptr("Confidential"),
		// 						Enabled: to.Ptr(true),
		// 						Order: to.Ptr[int32](300),
		// 					},
		// 					"7aa516c7-5a53-4857-bc6e-6808c6acd542": &armsecurity.SensitivityLabel{
		// 						DisplayName: to.Ptr("General"),
		// 						Enabled: to.Ptr(true),
		// 						Order: to.Ptr[int32](200),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
