package armplatformvalidation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/platformvalidation/armplatformvalidation"
)

// Generated from example definition: 2026-08-01-preview/ValidationExecutionPlans_Update_MaximumSet_Gen.json
func ExampleValidationExecutionPlansClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armplatformvalidation.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewValidationExecutionPlansClient().BeginUpdate(ctx, "rgvalidate", "cvtest01", "veptest01", armplatformvalidation.ValidationExecutionPlanUpdate{
		Properties: &armplatformvalidation.ValidationExecutionPlanUpdateProperties{
			Description:           to.Ptr("Validation execution plan that runs the network connectivity checks against the target image."),
			PlanConfigurationJSON: to.Ptr("{\"apiVersion\":\"microsoft.PlatformValidation/validationExecutionPlan.v0\",\"kind\":\"ValidationExecutionPlan\",\"metadata\":{\"name\":\"contoso-linux-cert\"},\"parameters\":{\"certificationPackageReference\":{\"osType\":\"Linux\",\"vmGenerationType\":\"V1\",\"architectureType\":\"X64\",\"recommendedVMSizes\":[\"Standard_D4s_v3\"],\"storageProfile\":{\"osDiskImage\":{\"sourceVhdUri\":\"https://contoso.blob.core.windows.net/vhds/img.vhd?<sas>\"},\"dataDiskImages\":[]},\"additionalProperties\":{}}},\"authoring\":{\"steps\":[{\"name\":\"os-disk-size\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/os-disk-size/versions/1.0.0\"},{\"name\":\"data-disk-size\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/data-disk-size/versions/1.0.0\"},{\"name\":\"malware-defender\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/malware-defender/versions/1.0.0\"},{\"name\":\"malware-esrp\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/malware-esrp/versions/1.0.0\"},{\"name\":\"linux-quality-validation\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/linux-quality-validation/versions/1.0.0\",\"inputs\":{\"concurrency\":1,\"testSuite\":[{\"testNames\":[\"smoke_test\",\"validate_netvsc_reload\"]}]}}]}}"),
		},
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armplatformvalidation.ValidationExecutionPlansClientUpdateResponse{
	// 	ValidationExecutionPlan: armplatformvalidation.ValidationExecutionPlan{
	// 		Properties: &armplatformvalidation.ValidationExecutionPlanProperties{
	// 			Description: to.Ptr("Validation execution plan that runs the network connectivity checks against the target image."),
	// 			PlanConfigurationJSON: to.Ptr("{\"apiVersion\":\"microsoft.PlatformValidation/validationExecutionPlan.v0\",\"kind\":\"ValidationExecutionPlan\",\"metadata\":{\"name\":\"contoso-linux-cert\"},\"parameters\":{\"certificationPackageReference\":{\"osType\":\"Linux\",\"vmGenerationType\":\"V1\",\"architectureType\":\"X64\",\"recommendedVMSizes\":[\"Standard_D4s_v3\"],\"storageProfile\":{\"osDiskImage\":{\"sourceVhdUri\":\"https://contoso.blob.core.windows.net/vhds/img.vhd?<sas>\"},\"dataDiskImages\":[]},\"additionalProperties\":{}}},\"authoring\":{\"steps\":[{\"name\":\"os-disk-size\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/os-disk-size/versions/1.0.0\"},{\"name\":\"data-disk-size\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/data-disk-size/versions/1.0.0\"},{\"name\":\"malware-defender\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/malware-defender/versions/1.0.0\"},{\"name\":\"malware-esrp\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/malware-esrp/versions/1.0.0\"},{\"name\":\"linux-quality-validation\",\"type\":\"test\",\"testRef\":\"/providers/Microsoft.PlatformValidation/validationTests/linux-quality-validation/versions/1.0.0\",\"inputs\":{\"concurrency\":1,\"testSuite\":[{\"testNames\":[\"smoke_test\",\"validate_netvsc_reload\"]}]}}]}}"),
	// 			ProvisioningState: to.Ptr(armplatformvalidation.ValidationExecutionPlanProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"owner-team": to.Ptr("azure-platform-validation"),
	// 		},
	// 		Location: to.Ptr("southcentralus"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgvalidate/providers/Microsoft.PlatformValidation/cloudValidations/cvtest01/validationExecutionPlans/veptest01"),
	// 		Name: to.Ptr("veptest01"),
	// 		Type: to.Ptr("Microsoft.PlatformValidation/cloudValidations/validationExecutionPlans"),
	// 		SystemData: &armplatformvalidation.SystemData{
	// 			CreatedBy: to.Ptr("user@example.com"),
	// 			CreatedByType: to.Ptr(armplatformvalidation.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.June, 1, 11, 52, 22, 926000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@example.com"),
	// 			LastModifiedByType: to.Ptr(armplatformvalidation.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.June, 1, 11, 52, 22, 926000000, time.UTC)),
	// 		},
	// 	},
	// }
}
