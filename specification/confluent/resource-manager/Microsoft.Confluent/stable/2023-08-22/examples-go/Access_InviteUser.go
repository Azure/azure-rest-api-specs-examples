package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_InviteUser.json
func ExampleAccessClient_InviteUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().InviteUser(ctx, "myResourceGroup", "myOrganization", armconfluent.AccessInviteUserAccountModel{
		InvitedUserDetails: &armconfluent.AccessInvitedUserDetails{
			AuthType:     to.Ptr("AUTH_TYPE_SSO"),
			InvitedEmail: to.Ptr("user2@onmicrosoft.com"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InvitationRecord = armconfluent.InvitationRecord{
	// 	AcceptedAt: to.Ptr("2022-07-06T17:21:33Z"),
	// 	AuthType: to.Ptr("AUTH_TYPE_SSO"),
	// 	Email: to.Ptr("johndoe@confluent.io"),
	// 	ExpiresAt: to.Ptr("2022-07-07T17:22:39Z"),
	// 	ID: to.Ptr("dlz-f3a90de"),
	// 	Kind: to.Ptr("Invitation"),
	// 	Metadata: &armconfluent.MetadataEntity{
	// 		CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 		DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 		ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/invitation=i-12345"),
	// 		Self: to.Ptr("https://api.confluent.cloud/iam/v2/invitations/i-12345"),
	// 		UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 	},
	// 	Status: to.Ptr("INVITE_STATUS_SENT"),
	// }
}
