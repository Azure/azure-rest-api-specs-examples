package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/ProjectsInSubscription_List.json
func ExampleProjectsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ProjectResultList = armmigrate.ProjectResultList{
		// 	Value: []*armmigrate.Project{
		// 		{
		// 			Name: to.Ptr("site1493ae9ea68project"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"0500be57-0000-0300-0000-5cb893f70000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/SMSValidations/providers/Microsoft.Migrate/assessmentprojects/site1493ae9ea68project"),
		// 			Location: to.Ptr("centralus"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/SMSValidations/providers/Microsoft.Migrate/MigrateProjects/SMSValidations-MigrateProject/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-15T09:18:02.572Z"); return t}()),
		// 				CustomerWorkspaceID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/SMSValidations/providers/Microsoft.OperationalInsights/workspaces/test-haili-01"),
		// 				CustomerWorkspaceLocation: to.Ptr("southeastasia"),
		// 				LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-18T15:12:55.238Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](12),
		// 				NumberOfGroups: to.Ptr[int32](8),
		// 				NumberOfMachines: to.Ptr[int32](26),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodcus.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-18T03:31:20.836Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("site1ad5aa6cc09project"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"8300bdec-0000-0300-0000-5cd678410000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/ppValidation/providers/Microsoft.Migrate/assessmentprojects/site1ad5aa6cc09project"),
		// 			Location: to.Ptr("centralus"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/ppValidation/providers/Microsoft.Migrate/MigrateProjects/ppValidation-MigrateProject/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-18T16:46:46.084Z"); return t}()),
		// 				LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-11T07:22:41.555Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](12),
		// 				NumberOfGroups: to.Ptr[int32](7),
		// 				NumberOfMachines: to.Ptr[int32](29),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodcus.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-18T16:46:46.084Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("migrateproject0720project"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"0d00efcf-0000-0300-0000-5d6fdac70000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/SDKUpgradeValidations/providers/Microsoft.Migrate/assessmentprojects/migrateproject0720project"),
		// 			Location: to.Ptr("centralus"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/SDKUpgradeValidations/providers/Microsoft.Migrate/MigrateProjects/SDKUpgradeValidations-MigrateProject/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-22T07:06:04.972Z"); return t}()),
		// 				LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-04T15:39:51.304Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](2),
		// 				NumberOfGroups: to.Ptr[int32](1),
		// 				NumberOfMachines: to.Ptr[int32](24),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodcus.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-22T07:06:04.972Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("migrateproject03acproject"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"00004d14-0000-0300-0000-5cb820290000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/AccessibilityTesting/providers/Microsoft.Migrate/assessmentprojects/migrateproject03acproject"),
		// 			Location: to.Ptr("centralus"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/AccessibilityTesting/providers/Microsoft.Migrate/MigrateProjects/AccessibilityTesting-MigrateProject/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-18T06:53:59.419Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](0),
		// 				NumberOfGroups: to.Ptr[int32](0),
		// 				NumberOfMachines: to.Ptr[int32](16),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodcus.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-18T06:53:59.419Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("migrateprojecta961project"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"1300df9f-0000-0300-0000-5d6e6d860000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/vmwaretesting/providers/Microsoft.Migrate/assessmentprojects/migrateprojecta961project"),
		// 			Location: to.Ptr("centralus"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/vmwaretesting/providers/Microsoft.Migrate/MigrateProjects/vmwaretesting-MigrateProject/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-20T11:59:42.735Z"); return t}()),
		// 				CustomerWorkspaceID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/vmwaretesting/providers/Microsoft.OperationalInsights/workspaces/mahpar324211"),
		// 				CustomerWorkspaceLocation: to.Ptr("westeurope"),
		// 				NumberOfAssessments: to.Ptr[int32](0),
		// 				NumberOfGroups: to.Ptr[int32](1),
		// 				NumberOfMachines: to.Ptr[int32](47),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodcus.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-03T13:41:26.677Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("PortalGAValidations43bbproject"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"02006b1f-0000-0300-0000-5d24a44a0000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/GARegressionTesting/providers/Microsoft.Migrate/assessmentprojects/PortalGAValidations43bbproject"),
		// 			Location: to.Ptr("centralus"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/GARegressionTesting/providers/Microsoft.Migrate/MigrateProjects/PortalGAValidations/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-04T03:56:50.853Z"); return t}()),
		// 				LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-09T14:27:22.885Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](4),
		// 				NumberOfGroups: to.Ptr[int32](4),
		// 				NumberOfMachines: to.Ptr[int32](49),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodcus.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-04T03:56:50.853Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("vaindana-pre-ga120dproject"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"0400d98a-0000-0d00-0000-5cd3ff790000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/vaindana-migrate-ga-rg/providers/Microsoft.Migrate/assessmentprojects/vaindana-pre-ga120dproject"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/vaindana-migrate-ga-rg/providers/Microsoft.Migrate/MigrateProjects/vaindana-pre-ga/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-08T15:09:37.856Z"); return t}()),
		// 				CustomerWorkspaceID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/vaindana-migrate-ga-rg/providers/Microsoft.OperationalInsights/workspaces/vaindana-pre-ga-oms"),
		// 				CustomerWorkspaceLocation: to.Ptr("westeurope"),
		// 				LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T10:22:48.973Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](2),
		// 				NumberOfGroups: to.Ptr[int32](2),
		// 				NumberOfMachines: to.Ptr[int32](14),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodwe.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-08T17:12:42.178Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Mahesh-V2-Europe-Bugbash181eproject"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"0b00349d-0000-0d00-0000-5d22eb5b0000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/Mahesh-RG-Eurpe-Bugbash/providers/Microsoft.Migrate/assessmentprojects/Mahesh-V2-Europe-Bugbash181eproject"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/Mahesh-RG-Eurpe-Bugbash/providers/Microsoft.Migrate/MigrateProjects/Mahesh-V2-Europe-Bugbash/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T07:38:23.034Z"); return t}()),
		// 				LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-02T06:38:34.281Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](6),
		// 				NumberOfGroups: to.Ptr[int32](3),
		// 				NumberOfMachines: to.Ptr[int32](36),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodwe.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T07:38:23.034Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("abgoyalWEselfhostb72bproject"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"0600c777-0000-0d00-0000-5cdaa4170000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.migrate/migrateprojects/abgoyalweselfhost/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T08:28:53.330Z"); return t}()),
		// 				LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T11:18:47.789Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](3),
		// 				NumberOfGroups: to.Ptr[int32](2),
		// 				NumberOfMachines: to.Ptr[int32](28),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodwe.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T10:11:16.022Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("vaindana-pre-ga-10180project"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects"),
		// 			ETag: to.Ptr("\"01003c88-0000-0d00-0000-5d41601b0000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/vaindana-migrate-ga-rg/providers/Microsoft.Migrate/assessmentprojects/vaindana-pre-ga-10180project"),
		// 			Location: to.Ptr("westeurope"),
		// 			Properties: &armmigrate.ProjectProperties{
		// 				AssessmentSolutionID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/vaindana-migrate-ga-rg/providers/Microsoft.Migrate/MigrateProjects/vaindana-pre-ga-1/Solutions/Servers-Assessment-ServerAssessment"),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T11:58:13.021Z"); return t}()),
		// 				LastAssessmentTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-31T09:32:11.496Z"); return t}()),
		// 				NumberOfAssessments: to.Ptr[int32](4),
		// 				NumberOfGroups: to.Ptr[int32](2),
		// 				NumberOfMachines: to.Ptr[int32](101),
		// 				ProjectStatus: to.Ptr(armmigrate.ProjectStatusActive),
		// 				ProvisioningState: to.Ptr(armmigrate.ProvisioningStateSucceeded),
		// 				ServiceEndpoint: to.Ptr("https://asmsrvprodwe.prod.migration.windowsazure.com/"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-15T16:52:08.918Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
