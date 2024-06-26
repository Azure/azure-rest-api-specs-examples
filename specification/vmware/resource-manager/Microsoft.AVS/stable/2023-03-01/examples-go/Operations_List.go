package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a5e7ff51c8af3781e7f6dd3b82a3fc922e2f6f93/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationList = armavs.OperationList{
		// 	Value: []*armavs.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/operations/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Lists operations available on Microsoft.AVS resource provider."),
		// 				Operation: to.Ptr("List available Microsoft.AVS operations"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/register/action"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Register Subscription for Microsoft.AVS resource provider."),
		// 				Operation: to.Ptr("Register Subscription for Microsoft.AVS"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr(""),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/unregister/action"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Unregister Subscription for Microsoft.AVS resource provider."),
		// 				Operation: to.Ptr("Unregister Subscription for Microsoft.AVS"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr(""),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/checkNameAvailability/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Checks if the privateCloud Name is available"),
		// 				Operation: to.Ptr("Check Name Availability"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("checkNameAvailability"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/locations/checkNameAvailability/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Checks if the privateCloud Name is available"),
		// 				Operation: to.Ptr("Check Name Availability"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("locations/checkNameAvailability"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/locations/checkQuotaAvailability/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Checks if quota is available for the subscription"),
		// 				Operation: to.Ptr("Check Quota Availability"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("locations/checkQuotaAvailability"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/locations/checkTrialAvailability/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Checks if trial is available for the subscription"),
		// 				Operation: to.Ptr("Check Trial Availability"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("locations/checkTrialAvailability"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/register/action"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Registers the Microsoft Microsoft.AVS resource provider and enables creation of Private Clouds."),
		// 				Operation: to.Ptr("Register Microsoft.AVS resource provider."),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/write"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates a PrivateCloud resource."),
		// 				Operation: to.Ptr("Create or update a PrivateCloud."),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Gets the settings for the specified PrivateCloud."),
		// 				Operation: to.Ptr("Read PrivateCloud settings"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/delete"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Delete a specific PrivateCloud."),
		// 				Operation: to.Ptr("Delete a PrivateCloud."),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/operationstatuses/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Reads privateClouds operationstatuses."),
		// 				Operation: to.Ptr("Read privateClouds operationstatuses"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/operationstatuses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/clusters/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Gets the cluster settings for a PrivateCloud cluster."),
		// 				Operation: to.Ptr("Read Cluster settings."),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/clusters"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/clusters/write"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Create or update a PrivateCloud cluster resource."),
		// 				Operation: to.Ptr("Create or update a PrivateCloud cluster."),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/clusters"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/clusters/delete"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Delete a specific PrivateCloud cluster."),
		// 				Operation: to.Ptr("Delete a PriveCloud cluster."),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/clusters"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/clusters/operationstatuses/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Reads privateClouds/clusters operationstatuses."),
		// 				Operation: to.Ptr("Read privateClouds/clusters operationstatuses"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/clusters/operationstatuses"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateclouds/clusters/operationresults/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Reads privateClouds/clusters operationresults."),
		// 				Operation: to.Ptr("Read privateClouds/clusters operationresults"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateclouds/clusters/operationresults"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/operationresults/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Reads privateClouds operationresults."),
		// 				Operation: to.Ptr("Read privateClouds operationresults"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/operationresults"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/authorizations/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Gets the authorization settings for a PrivateCloud cluster."),
		// 				Operation: to.Ptr("Read Authorization settings."),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/authorizations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/authorizations/write"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Create or update a PrivateCloud authorization resource."),
		// 				Operation: to.Ptr("Create or update a PrivateCloud authorization."),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/authorizations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/authorizations/delete"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Delete a specific PrivateCloud authorization."),
		// 				Operation: to.Ptr("Delete a PriveCloud authorization."),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/authorizations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/listAdminCredentials/action"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Lists the AdminCredentials for privateClouds."),
		// 				Operation: to.Ptr("List privateClouds AdminCredentials"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds"),
		// 			},
		// 			IsDataAction: to.Ptr(true),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/hcxEnterpriseSites/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Gets the hcxEnterpriseSites for a PrivateCloud."),
		// 				Operation: to.Ptr("Read hcxEnterpriseSites"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/hcxEnterpriseSites"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/hcxEnterpriseSites/write"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Create or update a hcxEnterpriseSites."),
		// 				Operation: to.Ptr("Create or update a hcxEnterpriseSites"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/hcxEnterpriseSites"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/hcxEnterpriseSites/delete"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Delete a specific hcxEnterpriseSites."),
		// 				Operation: to.Ptr("Delete a hcxEnterpriseSites"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/hcxEnterpriseSites"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/hostInstances/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Gets the hostInstances for a PrivateCloud."),
		// 				Operation: to.Ptr("Read hostInstances"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/hostInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/hostInstances/write"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Create or update a hostInstances."),
		// 				Operation: to.Ptr("Create or update a hostInstances"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/hostInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/hostInstances/delete"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Delete a specific hostInstances."),
		// 				Operation: to.Ptr("Delete a hostInstances"),
		// 				Provider: to.Ptr("Microsoft.AVS"),
		// 				Resource: to.Ptr("privateClouds/hostInstances"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 			Properties: &armavs.OperationProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS/privateClouds/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armavs.OperationDisplay{
		// 				Description: to.Ptr("Gets the available metrics for Private Cloud"),
		// 				Operation: to.Ptr("Read Private Cloud metric definitions"),
		// 				Provider: to.Ptr("Microsoft Azure Dedicated"),
		// 				Resource: to.Ptr("privateClouds"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("system"),
		// 			Properties: &armavs.OperationProperties{
		// 				ServiceSpecification: &armavs.ServiceSpecification{
		// 					MetricSpecifications: []*armavs.MetricSpecification{
		// 						{
		// 							Name: to.Ptr("UsedLatest"),
		// 							AggregationType: to.Ptr("Average"),
		// 							DisplayDescription: to.Ptr("The total amount of disk used in SDDC"),
		// 							DisplayName: to.Ptr("Datastore disk used"),
		// 							EnableRegionalMdmAccount: to.Ptr("true"),
		// 							SourceMdmAccount: to.Ptr("AVSShoebox2"),
		// 							SourceMdmNamespace: to.Ptr("Vsphere.Datastore.Disk"),
		// 							SupportedTimeGrainTypes: []*string{
		// 								to.Ptr("PT5M"),
		// 								to.Ptr("PT15M"),
		// 								to.Ptr("PT30M"),
		// 								to.Ptr("PT1H"),
		// 								to.Ptr("PT6H"),
		// 								to.Ptr("PT12H"),
		// 								to.Ptr("P1D")},
		// 								Unit: to.Ptr("Bytes"),
		// 							},
		// 							{
		// 								Name: to.Ptr("CapacityLatest"),
		// 								AggregationType: to.Ptr("Average"),
		// 								DisplayDescription: to.Ptr("The total capacity of disk in SDDC"),
		// 								DisplayName: to.Ptr("Datastore disk total capacity"),
		// 								EnableRegionalMdmAccount: to.Ptr("true"),
		// 								SourceMdmAccount: to.Ptr("AVSShoebox2"),
		// 								SourceMdmNamespace: to.Ptr("Vsphere.Datastore.Disk"),
		// 								SupportedTimeGrainTypes: []*string{
		// 									to.Ptr("PT5M"),
		// 									to.Ptr("PT15M"),
		// 									to.Ptr("PT30M"),
		// 									to.Ptr("PT1H"),
		// 									to.Ptr("PT6H"),
		// 									to.Ptr("PT12H"),
		// 									to.Ptr("P1D")},
		// 									Unit: to.Ptr("Bytes"),
		// 								},
		// 								{
		// 									Name: to.Ptr("EffectiveMemAverage"),
		// 									AggregationType: to.Ptr("Average"),
		// 									DisplayDescription: to.Ptr("Total amount of machine memory in cluster that is available"),
		// 									DisplayName: to.Ptr("Average Effective Memory"),
		// 									EnableRegionalMdmAccount: to.Ptr("true"),
		// 									SourceMdmAccount: to.Ptr("AVSShoebox2"),
		// 									SourceMdmNamespace: to.Ptr("Vsphere.Cluster.ClusterServices"),
		// 									SupportedTimeGrainTypes: []*string{
		// 										to.Ptr("PT5M"),
		// 										to.Ptr("PT15M"),
		// 										to.Ptr("PT30M"),
		// 										to.Ptr("PT1H"),
		// 										to.Ptr("PT6H"),
		// 										to.Ptr("PT12H"),
		// 										to.Ptr("P1D")},
		// 										Unit: to.Ptr("Bytes"),
		// 									},
		// 									{
		// 										Name: to.Ptr("TotalMbAverage"),
		// 										AggregationType: to.Ptr("Average"),
		// 										DisplayDescription: to.Ptr("Total memory in cluster"),
		// 										DisplayName: to.Ptr("Average Total Memory"),
		// 										EnableRegionalMdmAccount: to.Ptr("true"),
		// 										SourceMdmAccount: to.Ptr("AVSShoebox2"),
		// 										SourceMdmNamespace: to.Ptr("Vsphere.Cluster.Mem"),
		// 										SupportedTimeGrainTypes: []*string{
		// 											to.Ptr("PT5M"),
		// 											to.Ptr("PT15M"),
		// 											to.Ptr("PT30M"),
		// 											to.Ptr("PT1H"),
		// 											to.Ptr("PT6H"),
		// 											to.Ptr("PT12H"),
		// 											to.Ptr("P1D")},
		// 											Unit: to.Ptr("Bytes"),
		// 										},
		// 										{
		// 											Name: to.Ptr("OverheadAverage"),
		// 											AggregationType: to.Ptr("Average"),
		// 											DisplayDescription: to.Ptr("Host physical memory consumed by the virtualization infrastructure"),
		// 											DisplayName: to.Ptr("Average Memory Overhead"),
		// 											EnableRegionalMdmAccount: to.Ptr("true"),
		// 											SourceMdmAccount: to.Ptr("AVSShoebox2"),
		// 											SourceMdmNamespace: to.Ptr("Vsphere.Cluster.Mem"),
		// 											SupportedTimeGrainTypes: []*string{
		// 												to.Ptr("PT5M"),
		// 												to.Ptr("PT15M"),
		// 												to.Ptr("PT30M"),
		// 												to.Ptr("PT1H"),
		// 												to.Ptr("PT6H"),
		// 												to.Ptr("PT12H"),
		// 												to.Ptr("P1D")},
		// 												Unit: to.Ptr("Bytes"),
		// 											},
		// 											{
		// 												Name: to.Ptr("UsageAverage"),
		// 												AggregationType: to.Ptr("Average"),
		// 												DisplayDescription: to.Ptr("Memory usage as percentage of total configured or available memory"),
		// 												DisplayName: to.Ptr("Average Memory Usage"),
		// 												EnableRegionalMdmAccount: to.Ptr("true"),
		// 												SourceMdmAccount: to.Ptr("AVSShoebox2"),
		// 												SourceMdmNamespace: to.Ptr("Vsphere.Cluster.Mem"),
		// 												SupportedTimeGrainTypes: []*string{
		// 													to.Ptr("PT5M"),
		// 													to.Ptr("PT15M"),
		// 													to.Ptr("PT30M"),
		// 													to.Ptr("PT1H"),
		// 													to.Ptr("PT6H"),
		// 													to.Ptr("PT12H"),
		// 													to.Ptr("P1D")},
		// 													Unit: to.Ptr("Percent"),
		// 												},
		// 												{
		// 													Name: to.Ptr("EffectiveCpuAverage"),
		// 													AggregationType: to.Ptr("Average"),
		// 													DisplayDescription: to.Ptr("Total available CPU resources in cluster"),
		// 													DisplayName: to.Ptr("Effective CPU available"),
		// 													EnableRegionalMdmAccount: to.Ptr("true"),
		// 													SourceMdmAccount: to.Ptr("AVSShoebox2"),
		// 													SourceMdmNamespace: to.Ptr("Vsphere.Cluster.ClusterServices"),
		// 													SupportedTimeGrainTypes: []*string{
		// 														to.Ptr("PT5M"),
		// 														to.Ptr("PT15M"),
		// 														to.Ptr("PT30M"),
		// 														to.Ptr("PT1H"),
		// 														to.Ptr("PT6H"),
		// 														to.Ptr("PT12H"),
		// 														to.Ptr("P1D")},
		// 														Unit: to.Ptr("Percent"),
		// 												}},
		// 											},
		// 										},
		// 									},
		// 									{
		// 										Name: to.Ptr("Microsoft.AVS/privateClouds/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 										Display: &armavs.OperationDisplay{
		// 											Description: to.Ptr("Gets the diagnostic setting for the resource"),
		// 											Operation: to.Ptr("Read diagnostic setting"),
		// 											Provider: to.Ptr("Microsoft.AVS"),
		// 											Resource: to.Ptr("privateClouds"),
		// 										},
		// 										IsDataAction: to.Ptr(false),
		// 										Origin: to.Ptr("system"),
		// 										Properties: &armavs.OperationProperties{
		// 										},
		// 									},
		// 									{
		// 										Name: to.Ptr("Microsoft.AVS/privateClouds/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 										Display: &armavs.OperationDisplay{
		// 											Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 											Operation: to.Ptr("Write diagnostic setting"),
		// 											Provider: to.Ptr("Microsoft.AVS"),
		// 											Resource: to.Ptr("privateClouds"),
		// 										},
		// 										IsDataAction: to.Ptr(false),
		// 										Origin: to.Ptr("system"),
		// 										Properties: &armavs.OperationProperties{
		// 										},
		// 								}},
		// 							}
	}
}
