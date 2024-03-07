package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/InformationProtectionPolicies/CreateOrUpdateInformationProtectionPolicy_example.json
func ExampleInformationProtectionPoliciesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInformationProtectionPoliciesClient().CreateOrUpdate(ctx, "providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e", armsecurity.InformationProtectionPolicyNameCustom, armsecurity.InformationProtectionPolicy{
		Properties: &armsecurity.InformationProtectionPolicyProperties{
			InformationTypes: map[string]*armsecurity.InformationType{
				"3bf35491-99b8-41f2-86d5-c1200a7df658": {
					Custom:      to.Ptr(true),
					DisplayName: to.Ptr("Custom"),
					Enabled:     to.Ptr(true),
					Keywords: []*armsecurity.InformationProtectionKeyword{
						{
							CanBeNumeric: to.Ptr(true),
							Custom:       to.Ptr(true),
							Pattern:      to.Ptr("%custom%"),
						}},
					Order:              to.Ptr[int32](1400),
					RecommendedLabelID: to.Ptr("7aa516c7-5a53-4857-bc6e-6808c6acd542"),
				},
				"7fb9419d-2473-4ad8-8e11-b25cc8cf6a07": {
					Custom:      to.Ptr(false),
					DisplayName: to.Ptr("Networking"),
					Enabled:     to.Ptr(true),
					Keywords: []*armsecurity.InformationProtectionKeyword{
						{
							CanBeNumeric: to.Ptr(false),
							Custom:       to.Ptr(true),
							Pattern:      to.Ptr("%networking%"),
						}},
					Order:              to.Ptr[int32](100),
					RecommendedLabelID: to.Ptr("575739d2-3d53-4df0-9042-4c7772d5c7b1"),
				},
			},
			Labels: map[string]*armsecurity.SensitivityLabel{
				"1345da73-bc5a-4a8f-b7dd-3820eb713da8": {
					DisplayName: to.Ptr("Public"),
					Enabled:     to.Ptr(true),
					Order:       to.Ptr[int32](100),
				},
				"575739d2-3d53-4df0-9042-4c7772d5c7b1": {
					DisplayName: to.Ptr("Confidential"),
					Enabled:     to.Ptr(true),
					Order:       to.Ptr[int32](300),
				},
				"7aa516c7-5a53-4857-bc6e-6808c6acd542": {
					DisplayName: to.Ptr("General"),
					Enabled:     to.Ptr(true),
					Order:       to.Ptr[int32](200),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InformationProtectionPolicy = armsecurity.InformationProtectionPolicy{
	// 	Name: to.Ptr("custom"),
	// 	Type: to.Ptr("Microsoft.Security/informationProtectionPolicies"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/148059f7-faf3-49a6-ba35-85122112291e/providers/Microsoft.Security/informationProtectionPolicies/custom"),
	// 	Properties: &armsecurity.InformationProtectionPolicyProperties{
	// 		InformationTypes: map[string]*armsecurity.InformationType{
	// 			"3bf35491-99b8-41f2-86d5-c1200a7df658": &armsecurity.InformationType{
	// 				Custom: to.Ptr(true),
	// 				DisplayName: to.Ptr("Custom"),
	// 				Enabled: to.Ptr(true),
	// 				Keywords: []*armsecurity.InformationProtectionKeyword{
	// 					{
	// 						CanBeNumeric: to.Ptr(true),
	// 						Custom: to.Ptr(true),
	// 						Pattern: to.Ptr("%custom%"),
	// 				}},
	// 				Order: to.Ptr[int32](1400),
	// 				RecommendedLabelID: to.Ptr("7aa516c7-5a53-4857-bc6e-6808c6acd542"),
	// 			},
	// 			"7fb9419d-2473-4ad8-8e11-b25cc8cf6a07": &armsecurity.InformationType{
	// 				Custom: to.Ptr(false),
	// 				DisplayName: to.Ptr("Networking"),
	// 				Enabled: to.Ptr(true),
	// 				Keywords: []*armsecurity.InformationProtectionKeyword{
	// 					{
	// 						CanBeNumeric: to.Ptr(false),
	// 						Custom: to.Ptr(true),
	// 						Pattern: to.Ptr("%networking%"),
	// 				}},
	// 				Order: to.Ptr[int32](100),
	// 				RecommendedLabelID: to.Ptr("575739d2-3d53-4df0-9042-4c7772d5c7b1"),
	// 			},
	// 		},
	// 		Labels: map[string]*armsecurity.SensitivityLabel{
	// 			"1345da73-bc5a-4a8f-b7dd-3820eb713da8": &armsecurity.SensitivityLabel{
	// 				DisplayName: to.Ptr("Public"),
	// 				Enabled: to.Ptr(true),
	// 				Order: to.Ptr[int32](100),
	// 			},
	// 			"575739d2-3d53-4df0-9042-4c7772d5c7b1": &armsecurity.SensitivityLabel{
	// 				DisplayName: to.Ptr("Confidential"),
	// 				Enabled: to.Ptr(true),
	// 				Order: to.Ptr[int32](300),
	// 			},
	// 			"7aa516c7-5a53-4857-bc6e-6808c6acd542": &armsecurity.SensitivityLabel{
	// 				DisplayName: to.Ptr("General"),
	// 				Enabled: to.Ptr(true),
	// 				Order: to.Ptr[int32](200),
	// 			},
	// 		},
	// 	},
	// }
}
