package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveevent-operation-location.json
func ExampleLiveEventsClient_OperationLocation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLiveEventsClient().OperationLocation(ctx, "mediaresources", "slitestmedia10", "myLiveEvent1", "62e4d893-d233-4005-988e-a428d9f77076", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LiveEvent = armmediaservices.LiveEvent{
	// 	Name: to.Ptr("myLiveEvent1"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/liveevents"),
	// 	ID: to.Ptr("/subscriptions/0a6ec948-5a62-437d-b9df-934dc7c1b722/resourceGroups/mediaresources/providers/Microsoft.Media/mediaservices/slitestmedia10/liveevents/myLiveEvent1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armmediaservices.LiveEventProperties{
	// 		Description: to.Ptr(""),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 		CrossSiteAccessPolicies: &armmediaservices.CrossSiteAccessPolicies{
	// 			ClientAccessPolicy: to.Ptr("<access-policy><cross-domain-access><policy><allow-from http-methods=\"*\"><domain uri=\"http://*\"/></allow-from><grant-to><resource path=\"/\" include-subpaths=\"true\"/></grant-to></policy></cross-domain-access></access-policy>"),
	// 			CrossDomainPolicy: to.Ptr("<cross-domain-policy><allow-access-from domain=\"*\" secure=\"false\" /></cross-domain-policy>"),
	// 		},
	// 		Encoding: &armmediaservices.LiveEventEncoding{
	// 			EncodingType: to.Ptr(armmediaservices.LiveEventEncodingTypeNone),
	// 		},
	// 		Input: &armmediaservices.LiveEventInput{
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 				{
	// 					URL: to.Ptr("http://clouddeployment.media-test.net/ingest.isml"),
	// 					Protocol: to.Ptr("FragmentedMP4"),
	// 			}},
	// 			KeyFrameIntervalDuration: to.Ptr("PT6S"),
	// 			StreamingProtocol: to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
	// 		},
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 		Preview: &armmediaservices.LiveEventPreview{
	// 			AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
	// 				IP: &armmediaservices.IPAccessControl{
	// 					Allow: []*armmediaservices.IPRange{
	// 						{
	// 							Name: to.Ptr("AllowAll"),
	// 							Address: to.Ptr("0.0.0.0"),
	// 							SubnetPrefixLength: to.Ptr[int32](0),
	// 					}},
	// 				},
	// 			},
	// 			Endpoints: []*armmediaservices.LiveEventEndpoint{
	// 				{
	// 					URL: to.Ptr("https://testeventopito4idh2r-weibzmedia05.preview-ts051.channel.media-test.windows-int.net/763f3ea4-d94f-441c-a634-c833f61a4e04/preview.ism/manifest"),
	// 					Protocol: to.Ptr("FragmentedMP4"),
	// 			}},
	// 			PreviewLocator: to.Ptr("763f3ea4-d94f-441c-a634-c833f61a4e04"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceState: to.Ptr(armmediaservices.LiveEventResourceStateStopped),
	// 		StreamOptions: []*armmediaservices.StreamOptionsFlag{
	// 			to.Ptr(armmediaservices.StreamOptionsFlagDefault)},
	// 			UseStaticHostname: to.Ptr(false),
	// 		},
	// 		SystemData: &armmediaservices.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 			CreatedBy: to.Ptr("example@microsoft.com"),
	// 			CreatedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-03T02:25:08.347Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("example@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 		},
	// 	}
}
