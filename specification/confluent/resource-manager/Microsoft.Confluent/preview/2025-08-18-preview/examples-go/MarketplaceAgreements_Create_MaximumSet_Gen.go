package armconfluent_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent/v2"
)

// Generated from example definition: 2025-08-18-preview/MarketplaceAgreements_Create_MaximumSet_Gen.json
func ExampleMarketplaceAgreementsClient_Create_createConfluentMarketplaceAgreementInTheSubscriptionMaximumset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("DC34558A-05D3-4370-AED8-75E60B381F94", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMarketplaceAgreementsClient().Create(ctx, armconfluent.AgreementResource{
		Properties: &armconfluent.AgreementProperties{
			Publisher:         to.Ptr("cxcrrfggvdmvcchohkyatlvbpyy"),
			Product:           to.Ptr("ogusipjbwihlwbfivdbjfuvoqwija"),
			Plan:              to.Ptr("vgphlikczel"),
			LicenseTextLink:   to.Ptr("ztckliskduxmcluia"),
			PrivacyPolicyLink: to.Ptr("wwvlrlfhzmvfjgimkhkqcaxn"),
			RetrieveDatetime:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-18T11:10:31.028Z"); return t }()),
			Signature:         to.Ptr("cfdxpybzzsrgcdtebmqzzskxfiool"),
			Accepted:          to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armconfluent.MarketplaceAgreementsClientCreateResponse{
	// 	AgreementResource: &armconfluent.AgreementResource{
	// 		ID: to.Ptr("hqrct"),
	// 		Name: to.Ptr("nazsgocpqpboswffunhuxjytrya"),
	// 		Type: to.Ptr("mshlvpqsrhkba"),
	// 		SystemData: &armconfluent.SystemData{
	// 			CreatedBy: to.Ptr("lfskmafvssxoohhokqsa"),
	// 			CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-18T11:10:31.028Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("txubvkbhgirdizxd"),
	// 			LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-18T11:10:31.028Z"); return t}()),
	// 		},
	// 		Properties: &armconfluent.AgreementProperties{
	// 			Publisher: to.Ptr("cxcrrfggvdmvcchohkyatlvbpyy"),
	// 			Product: to.Ptr("ogusipjbwihlwbfivdbjfuvoqwija"),
	// 			Plan: to.Ptr("vgphlikczel"),
	// 			LicenseTextLink: to.Ptr("ztckliskduxmcluia"),
	// 			PrivacyPolicyLink: to.Ptr("wwvlrlfhzmvfjgimkhkqcaxn"),
	// 			RetrieveDatetime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-08-18T11:10:31.028Z"); return t}()),
	// 			Signature: to.Ptr("cfdxpybzzsrgcdtebmqzzskxfiool"),
	// 			Accepted: to.Ptr(true),
	// 		},
	// 	},
	// }
}
