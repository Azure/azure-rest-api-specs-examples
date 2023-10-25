package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Dra_Get.json
func ExampleDraClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDraClient().Get(ctx, "rgrecoveryservicesdatareplication", "wPR", "M", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DraModel = armrecoveryservicesdatareplication.DraModel{
	// 	Name: to.Ptr("ioxmwhzrzdilxivkvhpvzexl"),
	// 	Type: to.Ptr("ptgmahzsyv"),
	// 	ID: to.Ptr("anp"),
	// 	Properties: &armrecoveryservicesdatareplication.DraModelProperties{
	// 		AuthenticationIdentity: &armrecoveryservicesdatareplication.IdentityModel{
	// 			AADAuthority: to.Ptr("bubwwbowfhdmujrt"),
	// 			ApplicationID: to.Ptr("cwktzrwajuvfyyymfstpey"),
	// 			Audience: to.Ptr("dkjobanyqgzenivyxhvavottpc"),
	// 			ObjectID: to.Ptr("khsiaqfbpuhp"),
	// 			TenantID: to.Ptr("joclkkdovixwapephhxaqtefubhhmq"),
	// 		},
	// 		CorrelationID: to.Ptr("t"),
	// 		CustomProperties: &armrecoveryservicesdatareplication.DraModelCustomProperties{
	// 			InstanceType: to.Ptr("DraModelCustomProperties"),
	// 		},
	// 		HealthErrors: []*armrecoveryservicesdatareplication.HealthErrorModel{
	// 			{
	// 				AffectedResourceCorrelationIDs: []*string{
	// 					to.Ptr("fope")},
	// 					AffectedResourceType: to.Ptr("scfniv"),
	// 					Category: to.Ptr("leigw"),
	// 					Causes: to.Ptr("xznphqrrmsdzm"),
	// 					ChildErrors: []*armrecoveryservicesdatareplication.InnerHealthErrorModel{
	// 						{
	// 							Category: to.Ptr("lcsdxrqxquke"),
	// 							Causes: to.Ptr("kefaugkpxjkpulimjthjnl"),
	// 							Code: to.Ptr("yuxxpblihirpedwkigywgwjjrlzq"),
	// 							CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
	// 							HealthCategory: to.Ptr("mhdgfjqwbikhxmhtomkl"),
	// 							IsCustomerResolvable: to.Ptr(true),
	// 							Message: to.Ptr("sskcei"),
	// 							Recommendation: to.Ptr("kqybwaesqumywtjepi"),
	// 							Severity: to.Ptr("wqxxiuaqjyagq"),
	// 							Source: to.Ptr("wevvftugwydzzw"),
	// 							Summary: to.Ptr("djsmgrltruljo"),
	// 					}},
	// 					Code: to.Ptr("dgxkefzmeukd"),
	// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
	// 					HealthCategory: to.Ptr("itc"),
	// 					IsCustomerResolvable: to.Ptr(true),
	// 					Message: to.Ptr("lbywtdprdqdekl"),
	// 					Recommendation: to.Ptr("gmssteizlhjtclyeoo"),
	// 					Severity: to.Ptr("vvdajssdcypewdyechilxjmuijvdd"),
	// 					Source: to.Ptr("iy"),
	// 					Summary: to.Ptr("jtooblbvaxxrvcwgscbobq"),
	// 			}},
	// 			IsResponsive: to.Ptr(true),
	// 			LastHeartbeat: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.127Z"); return t}()),
	// 			MachineID: to.Ptr("envzcoijbqhtrpncbjbhk"),
	// 			MachineName: to.Ptr("y"),
	// 			ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateSucceeded),
	// 			ResourceAccessIdentity: &armrecoveryservicesdatareplication.IdentityModel{
	// 				AADAuthority: to.Ptr("bubwwbowfhdmujrt"),
	// 				ApplicationID: to.Ptr("cwktzrwajuvfyyymfstpey"),
	// 				Audience: to.Ptr("dkjobanyqgzenivyxhvavottpc"),
	// 				ObjectID: to.Ptr("khsiaqfbpuhp"),
	// 				TenantID: to.Ptr("joclkkdovixwapephhxaqtefubhhmq"),
	// 			},
	// 			VersionNumber: to.Ptr("wnksfnisrhs"),
	// 		},
	// 		SystemData: &armrecoveryservicesdatareplication.DraModelSystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
	// 			CreatedBy: to.Ptr("fazidmklka"),
	// 			CreatedByType: to.Ptr("obpndgkaeyklqzmpjh"),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:52.128Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("cfoopkrisaroztncgss"),
	// 			LastModifiedByType: to.Ptr("dysxbvohxhrpl"),
	// 		},
	// 	}
}
