package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/Validations_ValidateOrganizationV2_MaximumSet_Gen.json
func ExampleValidationsClient_ValidateOrganizationV2() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewValidationsClient().ValidateOrganizationV2(ctx, "rgconfluent", "qhipfdfhxjzvwlergbvldnwhttfb", armconfluent.OrganizationResource{
		Properties: &armconfluent.OrganizationResourceProperties{
			ProvisioningState: to.Ptr(armconfluent.ProvisionStateAccepted),
			OfferDetail: &armconfluent.OfferDetail{
				PublisherID:    to.Ptr("jvmchwpbqvavlgmuwquhqrnacgpvlobkkavwppwvhjfqcy"),
				ID:             to.Ptr("ufewkfngssvswmxfurnchnvgmnjuzzsoys"),
				PlanID:         to.Ptr("l"),
				PlanName:       to.Ptr("ycpeesrtyybhvmkdenugbkffjwistugfertrprgevcczlsnbcinotsdtsmealomyzsinypzimyyubepkuewirtcxhvxhsmwhwptvzuhirckvrgogahfwchvxnfkgfwqxqy"),
				TermUnit:       to.Ptr("ipefrkgclpjaswyxpyjkppo"),
				TermID:         to.Ptr("vujdve"),
				PrivateOfferID: to.Ptr("goshpcnjukfzfhubmynjxiulurrwplzcjpjstebtsiigbkovchcrlfmgoymqfuayhihnxruthwjywtedlcsqqsgaelqthvfzvafyjhsfzfjwotsiajpcogwrwylgcphxfhvvwemynoyovnvqcetftiofkthgdzfvybvhpviqlwlslaupndcxlvjssdap"),
				PrivateOfferIDs: []*string{
					to.Ptr("nrbzkbcpvsakewlgubfmej"),
				},
				Status: to.Ptr(armconfluent.SaaSOfferStatusStarted),
			},
			UserDetail: &armconfluent.UserDetail{
				FirstName:         to.Ptr("gqxqhtniapwvnsliaifhvmbtvvrciebktpeadanapfcqzflomz"),
				LastName:          to.Ptr("vhdbyshxnnxivxbgzxscscdsvlwbsukqmcw"),
				EmailAddress:      to.Ptr("user@example.com"),
				UserPrincipalName: to.Ptr("g"),
				AADEmail:          to.Ptr("swugcwecfnkp"),
			},
			LinkOrganization: &armconfluent.LinkOrganization{
				Token: to.Ptr("yvynviychdid"),
			},
		},
		Tags: map[string]*string{
			"key2047": to.Ptr("maumzdwvkwd"),
		},
		Location: to.Ptr("ogifpylahax"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconfluent.ValidationsClientValidateOrganizationV2Response{
	// 	ValidationResponse: &armconfluent.ValidationResponse{
	// 		Info: map[string]*string{
	// 			"key7115": to.Ptr("owvfyhravpcrkzc"),
	// 		},
	// 	},
	// }
}
