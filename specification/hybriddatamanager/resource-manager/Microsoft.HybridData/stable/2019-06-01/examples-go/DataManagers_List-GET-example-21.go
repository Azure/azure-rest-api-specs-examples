package armhybriddatamanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataManagers_List-GET-example-21.json
func ExampleDataManagersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybriddatamanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataManagersClient().NewListPager(nil)
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
		// page.DataManagerList = armhybriddatamanager.DataManagerList{
		// 	Value: []*armhybriddatamanager.DataManager{
		// 		{
		// 			Name: to.Ptr("batchcertneafterdep1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertneafterdep1"),
		// 			Location: to.Ptr("northeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-23T09%3A58%3A15.3299896Z'\"_W/\"datetime'2018-07-23T09%3A58%3A15.3500041Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("batchcertneold"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertneold"),
		// 			Location: to.Ptr("northeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-16T09%3A46%3A06.3239385Z'\"_W/\"datetime'2018-07-16T09%3A46%3A06.3289422Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("dmsnesmoketest"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/dmsnesmoketest"),
		// 			Location: to.Ptr("northeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-03-15T10%3A53%3A10.0933461Z'\"_W/\"datetime'2018-03-15T10%3A53%3A10.1213654Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ne-07-10"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/ne-07-10"),
		// 			Location: to.Ptr("northeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-10T06%3A36%3A18.0878861Z'\"_W/\"datetime'2018-07-10T06%3A36%3A18.1139046Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("dms-04-10"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/dms-04-10"),
		// 			Location: to.Ptr("eastus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-04-10T06%3A09%3A32.3093315Z'\"_W/\"datetime'2018-04-10T06%3A09%3A32.3243425Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("eus-07-08"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/eus-07-08"),
		// 			Location: to.Ptr("eastus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-07T19%3A05%3A00.7594128Z'\"_W/\"datetime'2018-07-07T19%3A05%3A00.7664174Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("eus-07-10"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/eus-07-10"),
		// 			Location: to.Ptr("eastus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-10T06%3A36%3A46.642745Z'\"_W/\"datetime'2018-07-10T06%3A36%3A46.7658314Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("batchcertwus2afterdep1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertwus2afterdep1"),
		// 			Location: to.Ptr("westus2"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-23T07%3A44%3A12.1141909Z'\"_W/\"datetime'2018-07-23T07%3A44%3A12.1432118Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("batchcertwus2old"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertwus2old"),
		// 			Location: to.Ptr("westus2"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-16T09%3A44%3A49.261222Z'\"_W/\"datetime'2018-07-16T09%3A44%3A49.2702259Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("smoketestwus2"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/smoketestwus2"),
		// 			Location: to.Ptr("westus2"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-03-20T08%3A08%3A52.7007451Z'\"_W/\"datetime'2018-03-20T08%3A08%3A52.7207592Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("AzureSDKOperations"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/AzureSDKOperations"),
		// 			Location: to.Ptr("westus2"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-02-17T14%3A50%3A37.866739Z'\"_W/\"datetime'2019-02-17T14%3A50%3A38.038859Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("batchcertwcusafterdep1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertwcusafterdep1"),
		// 			Location: to.Ptr("westcentralus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-18T04%3A48%3A50.1962283Z'\"_W/\"datetime'2018-07-18T04%3A48%3A50.3433306Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("batchcertwcusold"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertwcusold"),
		// 			Location: to.Ptr("westcentralus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-16T09%3A43%3A43.569014Z'\"_W/\"datetime'2018-07-16T09%3A43%3A43.5740171Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("smoketestwcus"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/smoketestwcus"),
		// 			Location: to.Ptr("westcentralus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-03-20T08%3A08%3A11.5901685Z'\"_W/\"datetime'2018-03-20T08%3A08%3A11.6161871Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("wcus04-13"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/wcus04-13"),
		// 			Location: to.Ptr("westcentralus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-04-13T09%3A17%3A33.9031753Z'\"_W/\"datetime'2018-04-13T09%3A17%3A33.9301949Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("wcus07-05"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/wcus07-05"),
		// 			Location: to.Ptr("westcentralus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-05T14%3A40%3A04.7506699Z'\"_W/\"datetime'2018-07-05T14%3A40%3A04.7776888Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("wcussmoketest"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/wcussmoketest"),
		// 			Location: to.Ptr("westcentralus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-06-05T11%3A17%3A08.7276428Z'\"_W/\"datetime'2018-06-05T11%3A17%3A08.7336469Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("smoketest"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ForDMS/providers/Microsoft.HybridData/dataManagers/smoketest"),
		// 			Location: to.Ptr("westcentralus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-05-17T09%3A53%3A29.1283879Z'\"_W/\"datetime'2019-05-17T09%3A53%3A29.1844278Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("seasmoketestresource"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/amemigration/providers/Microsoft.HybridData/dataManagers/seasmoketestresource"),
		// 			Location: to.Ptr("southeastasia"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-07-06T16%3A47%3A17.0360354Z'\"_W/\"datetime'2019-07-06T16%3A47%3A17.2111588Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("batchcertseaafterdep1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertseaafterdep1"),
		// 			Location: to.Ptr("southeastasia"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-23T08%3A24%3A48.4779951Z'\"_W/\"datetime'2018-07-23T08%3A24%3A48.5120189Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("batchcertseaold"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertseaold"),
		// 			Location: to.Ptr("southeastasia"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-16T10%3A00%3A28.3513242Z'\"_W/\"datetime'2018-07-16T10%3A00%3A28.3583291Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("testresourceon613"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/testresourceon613"),
		// 			Location: to.Ptr("southeastasia"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-06-13T18%3A04%3A02.4340032Z'\"_W/\"datetime'2018-06-13T18%3A04%3A02.4430097Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("copyexp-wus-int"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/copyspeedexp/providers/Microsoft.HybridData/dataManagers/copyexp-wus-int"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-05-24T10%3A28%3A41.92206Z'\"_W/\"datetime'2019-05-24T10%3A28%3A42.0641646Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("testServiceIncident"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/DmsBvtRG3/providers/Microsoft.HybridData/dataManagers/testServiceIncident"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-02-22T11%3A13%3A22.4906285Z'\"_W/\"datetime'2018-02-22T11%3A13%3A22.6667555Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("prtankWusTestRes"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/DmsBvtRG/providers/Microsoft.HybridData/dataManagers/prtankWusTestRes"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-11-29T04%3A12%3A49.3583436Z'\"_W/\"datetime'2019-11-29T04%3A12%3A49.5344699Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("storagemanagerresource"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/DmsBvtRG/providers/Microsoft.HybridData/dataManagers/storagemanagerresource"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"doNotDelete": to.Ptr("Yes"),
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-11-19T08%3A06%3A46.7263604Z'\"_W/\"datetime'2019-11-19T08%3A06%3A46.7824012Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("sdf"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsdatasource/providers/Microsoft.HybridData/dataManagers/sdf"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-05-27T05%3A27%3A35.150482Z'\"_W/\"datetime'2019-05-27T05%3A27%3A35.3105951Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("int-wus-11-29"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsintrg/providers/Microsoft.HybridData/dataManagers/int-wus-11-29"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-11-29T07%3A55%3A30.6749871Z'\"_W/\"datetime'2018-11-29T07%3A55%3A30.6809932Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("intnewresource"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsintrg/providers/Microsoft.HybridData/dataManagers/intnewresource"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-03-17T15%3A03%3A36.5993632Z'\"_W/\"datetime'2018-03-17T15%3A03%3A36.7594788Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("intresource"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsintrg/providers/Microsoft.HybridData/dataManagers/intresource"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-03-13T08%3A11%3A39.9025375Z'\"_W/\"datetime'2018-03-13T08%3A11%3A40.1036797Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("intresource-10-30"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsintrg/providers/Microsoft.HybridData/dataManagers/intresource-10-30"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-10-30T09%3A56%3A32.9053638Z'\"_W/\"datetime'2018-10-30T09%3A56%3A33.0704765Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("testresourceint"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsintrg/providers/Microsoft.HybridData/dataManagers/testresourceint"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-03-13T08%3A11%3A40.1389003Z'\"_W/\"datetime'2018-03-13T08%3A11%3A40.2479789Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("batchcertoldresource"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertoldresource"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-12T09%3A46%3A50.1774789Z'\"_W/\"datetime'2018-07-12T09%3A46%3A50.3215821Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("dmswusresource"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/dmswusresource"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-03-01T11%3A52%3A02.7715263Z'\"_W/\"datetime'2018-03-01T11%3A52%3A02.7765292Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("dsmsnodebinarytest"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/dsmsnodebinarytest"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-06-07T13%3A08%3A29.4575101Z'\"_W/\"datetime'2018-06-07T13%3A08%3A29.629632Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("dsmsresource"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/dsmsresource"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-05-23T08%3A45%3A40.8795252Z'\"_W/\"datetime'2018-05-23T08%3A45%3A41.0666584Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("dsmsresource1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/dsmsresource1"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-06-05T05%3A02%3A54.5475181Z'\"_W/\"datetime'2018-06-05T05%3A02%3A54.73265Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("dsmsresource2"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/dsmsresource2"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-06-05T05%3A34%3A59.227556Z'\"_W/\"datetime'2018-06-05T05%3A34%3A59.4357029Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("dsmsresource3"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/dsmsresource3"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-06-05T11%3A02%3A52.3744002Z'\"_W/\"datetime'2018-06-05T11%3A02%3A52.5365141Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("dsmsresource4"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/dsmsresource4"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-06-08T15%3A58%3A07.1761866Z'\"_W/\"datetime'2018-06-08T15%3A58%3A07.3703309Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("govtcloudtest"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/govtcloudtest"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-10-03T08%3A03%3A21.5040402Z'\"_W/\"datetime'2018-10-03T08%3A03%3A21.6421388Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("intresource2"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/intresource2"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-02-01T10%3A38%3A07.9487894Z'\"_W/\"datetime'2018-02-01T10%3A38%3A08.1499556Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("intresource3rdjuly"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/intresource3rdjuly"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-03T15%3A33%3A45.7585904Z'\"_W/\"datetime'2018-07-03T15%3A33%3A45.932713Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("intsmallfiles"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/intsmallfiles"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-09-24T12%3A01%3A29.4312765Z'\"_W/\"datetime'2018-09-24T12%3A01%3A29.489317Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("longrunningjob"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/longrunningjob"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-05-31T09%3A41%3A28.6739766Z'\"_W/\"datetime'2018-05-31T09%3A41%3A28.8551054Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleclientdatamanager"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/sampleclientdatamanager"),
		// 			Location: to.Ptr("WestUS"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-02-21T10%3A06%3A55.2175183Z'\"_W/\"datetime'2018-02-21T10%3A06%3A55.2235223Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ssdmbcdrresource"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/ssdmbcdrresource"),
		// 			Location: to.Ptr("westus"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-04-23T10%3A16%3A00.5751127Z'\"_W/\"datetime'2018-04-23T10%3A16%3A00.5841193Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("copyexp-we-int1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/copyspeedexp/providers/Microsoft.HybridData/dataManagers/copyexp-we-int1"),
		// 			Location: to.Ptr("westeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-12-06T05%3A48%3A21.5818324Z'\"_W/\"datetime'2018-12-06T05%3A48%3A21.7748343Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("prtankbvttest"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/DmsBvtRG/providers/Microsoft.HybridData/dataManagers/prtankbvttest"),
		// 			Location: to.Ptr("westeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"doNotDelete": to.Ptr("yes"),
		// 				"test": to.Ptr("true"),
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-11-19T08%3A08%3A48.3033932Z'\"_W/\"datetime'2019-11-19T08%3A08%3A48.319407Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("prtankBvtWeTest"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/DmsBvtRG/providers/Microsoft.HybridData/dataManagers/prtankBvtWeTest"),
		// 			Location: to.Ptr("westeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2020-01-30T10%3A18%3A35.3300821Z'\"_W/\"datetime'2020-01-30T10%3A18%3A35.3460951Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("we-int-1"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/we-int-1"),
		// 			Location: to.Ptr("westeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-11-28T10%3A55%3A58.530785Z'\"_W/\"datetime'2018-11-28T10%3A55%3A58.8107858Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("mrinsaha-dms-we"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ForDMS/providers/Microsoft.HybridData/dataManagers/mrinsaha-dms-we"),
		// 			Location: to.Ptr("westeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-08-06T11%3A48%3A28.3406561Z'\"_W/\"datetime'2019-08-06T11%3A48%3A28.5107985Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("testdurga2"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ForDMS/providers/Microsoft.HybridData/dataManagers/testdurga2"),
		// 			Location: to.Ptr("westeurope"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-05-14T09%3A07%3A12.3469862Z'\"_W/\"datetime'2019-05-14T09%3A07%3A12.3589889Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ecy-aftermigration"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/amemigration/providers/Microsoft.HybridData/dataManagers/ecy-aftermigration"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-11-14T02%3A42%3A38.6041538Z'\"_W/\"datetime'2018-11-14T02%3A42%3A38.6081571Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ecy-ame"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/amemigration/providers/Microsoft.HybridData/dataManagers/ecy-ame"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-11-11T03%3A11%3A40.5411975Z'\"_W/\"datetime'2018-11-11T03%3A11%3A40.6752923Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ecysmoketest"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/cleanupservice/providers/Microsoft.HybridData/dataManagers/ecysmoketest"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-04-15T09%3A05%3A36.2264018Z'\"_W/\"datetime'2019-04-15T09%3A05%3A36.2854433Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("copyspeed-ecy"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/copyspeedexp/providers/Microsoft.HybridData/dataManagers/copyspeed-ecy"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-05-09T09%3A51%3A11.720705Z'\"_W/\"datetime'2019-05-09T09%3A51%3A11.8728117Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ecyresource-04-10"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/ecyresource-04-10"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-04-10T09%3A41%3A03.749493Z'\"_W/\"datetime'2018-04-10T09%3A41%3A03.8805855Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("feb20ecyresource"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/feb20ecyresource"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-02-20T10%3A01%3A55.2712681Z'\"_W/\"datetime'2018-02-20T10%3A01%3A55.441388Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("hari-dms-analysis-ecy"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/hari-dms-analysis-ecy"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-02-06T08%3A37%3A23.6486177Z'\"_W/\"datetime'2018-02-06T08%3A37%3A23.7857152Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("EcySmoke"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ForDMS/providers/Microsoft.HybridData/dataManagers/EcySmoke"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-01-14T05%3A16%3A38.9573714Z'\"_W/\"datetime'2019-01-14T05%3A16%3A39.0774564Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("EcyTestDMSRes"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/EcyTestDMSRes"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-08-12T10%3A20%3A40.4679832Z'\"_W/\"datetime'2019-08-12T10%3A20%3A40.6030796Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ccy-ame"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/amemigration/providers/Microsoft.HybridData/dataManagers/ccy-ame"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-11-11T03%3A13%3A19.0340832Z'\"_W/\"datetime'2018-11-11T03%3A13%3A19.038086Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ccytest-26-march"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/amemigration/providers/Microsoft.HybridData/dataManagers/ccytest-26-march"),
		// 			Location: to.Ptr("Central US EUAP"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-03-26T04%3A54%3A23.06065Z'\"_W/\"datetime'2019-03-26T04%3A54%3A23.1216929Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("batchcertoldccy"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/batchcertoldccy"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-07-12T09%3A47%3A44.2155827Z'\"_W/\"datetime'2018-07-12T09%3A47%3A44.2235883Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ccylargenumberoffiles"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/ccylargenumberoffiles"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-09-04T06%3A47%3A31.8676921Z'\"_W/\"datetime'2018-09-04T06%3A47%3A32.0268062Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("feb20ccyresoirce"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/feb20ccyresoirce"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-02-20T10%3A02%3A19.8430119Z'\"_W/\"datetime'2018-02-20T10%3A02%3A20.0291444Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("smallfilesccyjob"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/dmsResourceGroup/providers/Microsoft.HybridData/dataManagers/smallfilesccyjob"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-09-24T05%3A24%3A44.2753634Z'\"_W/\"datetime'2018-09-24T05%3A24%3A44.4114594Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("MSCCY"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ForDMS/providers/Microsoft.HybridData/dataManagers/MSCCY"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-04-22T07%3A45%3A22.6635575Z'\"_W/\"datetime'2019-04-22T07%3A45%3A22.8346781Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("res4ccyBCDR"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ForDMS/providers/Microsoft.HybridData/dataManagers/res4ccyBCDR"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-02-07T09%3A38%3A22.6976591Z'\"_W/\"datetime'2019-02-07T09%3A38%3A22.8767865Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("prpare50lakhsmallfiles"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/prpare/providers/Microsoft.HybridData/dataManagers/prpare50lakhsmallfiles"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-09-25T07%3A11%3A55.7616108Z'\"_W/\"datetime'2018-09-25T07%3A11%3A55.8957061Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("CcyTestDMSRes2"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/CcyTestDMSRes2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2019-08-13T04%3A47%3A07.5063631Z'\"_W/\"datetime'2019-08-13T04%3A47%3A07.5113667Z'\""),
		// 		},
		// 		{
		// 			Name: to.Ptr("ccytestingpav2"),
		// 			Type: to.Ptr("Microsoft.HybridData/dataManagers"),
		// 			ID: to.Ptr("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourcegroups/test-varavi/providers/Microsoft.HybridData/dataManagers/ccytestingpav2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			SKU: &armhybriddatamanager.SKU{
		// 				Name: to.Ptr("DS0"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 			Tags: map[string]*string{
		// 			},
		// 			Etag: to.Ptr("W/\"datetime'2018-06-15T07%3A04%3A34.2882012Z'\"_W/\"datetime'2018-06-15T07%3A04%3A34.3472441Z'\""),
		// 	}},
		// }
	}
}
