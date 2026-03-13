package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2026-02-01/PrivateLinkResources_Get_MaximumSet_Gen.json
func ExampleSchedulersClient_GetPrivateLink() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("851A7597-D699-45CC-899B-7487A5B3B775", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSchedulersClient().GetPrivateLink(ctx, "rgdurabletask", "testscheduler", "ulbdiqhrmwnkejje", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdurabletask.SchedulersClientGetPrivateLinkResponse{
	// 	SchedulerPrivateLinkResource: &armdurabletask.SchedulerPrivateLinkResource{
	// 		Properties: &armdurabletask.PrivateLinkResourceProperties{
	// 			GroupID: to.Ptr("mexetpneryldlrtmuxzxhwezfjkcvr"),
	// 			RequiredMembers: []*string{
	// 				to.Ptr("ftzshharzmwhcemnbdwlmyhtxkpa"),
	// 			},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("lkwwgycaduib"),
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/851A7597-D699-45CC-899B-7487A5B3B775/resourceGroups/rgdurabletask/providers/Microsoft.DurableTask/schedulers/testscheduler/privateLinkResources/ulbdiqhrmwnkejje"),
	// 		Name: to.Ptr("ulbdiqhrmwnkejje"),
	// 		Type: to.Ptr("Microsoft.DurableTask/schedulers/privateLinkResources"),
	// 		SystemData: &armdurabletask.SystemData{
	// 			CreatedBy: to.Ptr("tenmbevaunjzikxowqexrsx"),
	// 			CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.365Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xfvdcegtj"),
	// 			LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.366Z"); return t}()),
	// 		},
	// 	},
	// }
}
