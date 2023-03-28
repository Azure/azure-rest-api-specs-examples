package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/BlobContainersGet.json
func ExampleBlobContainersClient_Get_getContainers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobContainersClient().Get(ctx, "res9871", "sto6217", "container1634", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BlobContainer = armstorage.BlobContainer{
	// 	Name: to.Ptr("container1634"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/blobServices/containers"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9871/providers/Microsoft.Storage/storageAccounts/sto6217/blobServices/default/containers/container1634"),
	// 	Etag: to.Ptr("\"0x8D592D74CC20EBA\""),
	// 	ContainerProperties: &armstorage.ContainerProperties{
	// 		HasImmutabilityPolicy: to.Ptr(true),
	// 		HasLegalHold: to.Ptr(true),
	// 		ImmutabilityPolicy: &armstorage.ImmutabilityPolicyProperties{
	// 			Etag: to.Ptr("\"8d592d74cb3011a\""),
	// 			Properties: &armstorage.ImmutabilityPolicyProperty{
	// 				ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](100),
	// 				State: to.Ptr(armstorage.ImmutabilityPolicyStateLocked),
	// 			},
	// 			UpdateHistory: []*armstorage.UpdateHistoryProperty{
	// 				{
	// 					ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](3),
	// 					ObjectIdentifier: to.Ptr("ce7cd28a-fc25-4bf1-8fb9-e1b9833ffd4b"),
	// 					TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-26T05:06:11.431403Z"); return t}()),
	// 					Update: to.Ptr(armstorage.ImmutabilityPolicyUpdateTypePut),
	// 				},
	// 				{
	// 					ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](3),
	// 					ObjectIdentifier: to.Ptr("ce7cd28a-fc25-4bf1-8fb9-e1b9833ffd4b"),
	// 					TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-26T05:06:13.0907641Z"); return t}()),
	// 					Update: to.Ptr(armstorage.ImmutabilityPolicyUpdateTypeLock),
	// 				},
	// 				{
	// 					ImmutabilityPeriodSinceCreationInDays: to.Ptr[int32](100),
	// 					ObjectIdentifier: to.Ptr("ce7cd28a-fc25-4bf1-8fb9-e1b9833ffd4b"),
	// 					TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-26T05:06:14.7097716Z"); return t}()),
	// 					Update: to.Ptr(armstorage.ImmutabilityPolicyUpdateTypeExtend),
	// 			}},
	// 		},
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-26T05:06:14Z"); return t}()),
	// 		LeaseState: to.Ptr(armstorage.LeaseStateAvailable),
	// 		LeaseStatus: to.Ptr(armstorage.LeaseStatusUnlocked),
	// 		LegalHold: &armstorage.LegalHoldProperties{
	// 			HasLegalHold: to.Ptr(true),
	// 			Tags: []*armstorage.TagProperty{
	// 				{
	// 					ObjectIdentifier: to.Ptr("ce7cd28a-fc25-4bf1-8fb9-e1b9833ffd4b"),
	// 					Tag: to.Ptr("tag1"),
	// 					TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-26T05:06:09.6964643Z"); return t}()),
	// 				},
	// 				{
	// 					ObjectIdentifier: to.Ptr("ce7cd28a-fc25-4bf1-8fb9-e1b9833ffd4b"),
	// 					Tag: to.Ptr("tag2"),
	// 					TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-26T05:06:09.6964643Z"); return t}()),
	// 				},
	// 				{
	// 					ObjectIdentifier: to.Ptr("ce7cd28a-fc25-4bf1-8fb9-e1b9833ffd4b"),
	// 					Tag: to.Ptr("tag3"),
	// 					TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-26T05:06:09.6964643Z"); return t}()),
	// 			}},
	// 		},
	// 		PublicAccess: to.Ptr(armstorage.PublicAccessNone),
	// 	},
	// }
}
