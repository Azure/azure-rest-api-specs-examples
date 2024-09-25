package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/SessionHost_Get.json
func ExampleSessionHostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSessionHostsClient().Get(ctx, "resourceGroup1", "hostPool1", "sessionHost1.microsoft.com", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SessionHost = armdesktopvirtualization.SessionHost{
	// 	Name: to.Ptr("sessionHost1.microsoft.com"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/hostPools/sessionHosts"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/hostPools/hostPool1/sessionHosts/sessionHost1.microsoft.com"),
	// 	SystemData: &armdesktopvirtualization.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 	},
	// 	Properties: &armdesktopvirtualization.SessionHostProperties{
	// 		AgentVersion: to.Ptr("1.0.0.1391"),
	// 		AllowNewSession: to.Ptr(true),
	// 		AssignedUser: to.Ptr("user1@microsoft.com"),
	// 		FriendlyName: to.Ptr("friendly"),
	// 		LastHeartBeat: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.957Z"); return t}()),
	// 		LastUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.957Z"); return t}()),
	// 		ObjectID: to.Ptr("7877fb31-4bde-49fd-9df3-c046e0ec5325"),
	// 		OSVersion: to.Ptr("10.0.17763"),
	// 		ResourceID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.Compute/virtualMachines/sessionHost1"),
	// 		SessionHostHealthCheckResults: []*armdesktopvirtualization.SessionHostHealthCheckReport{
	// 			{
	// 				AdditionalFailureDetails: &armdesktopvirtualization.SessionHostHealthCheckFailureDetails{
	// 					ErrorCode: to.Ptr[int32](0),
	// 					LastHealthCheckDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-14T02:09:37.623Z"); return t}()),
	// 					Message: to.Ptr("SessionHost healthy: is joined to domain ≤wvdarmtest1.net≥"),
	// 				},
	// 				HealthCheckName: to.Ptr(armdesktopvirtualization.HealthCheckNameDomainJoinedCheck),
	// 				HealthCheckResult: to.Ptr(armdesktopvirtualization.HealthCheckResultHealthCheckSucceeded),
	// 		}},
	// 		Sessions: to.Ptr[int32](1),
	// 		Status: to.Ptr(armdesktopvirtualization.StatusAvailable),
	// 		StatusTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2008-09-22T14:01:54.957Z"); return t}()),
	// 		SxSStackVersion: to.Ptr("rdp-sxs190816002"),
	// 		UpdateErrorMessage: to.Ptr(""),
	// 		UpdateState: to.Ptr(armdesktopvirtualization.UpdateStateSucceeded),
	// 		VirtualMachineID: to.Ptr("29491b54-c033-4dec-b09a-18bf0ebafaef"),
	// 	},
	// }
}
