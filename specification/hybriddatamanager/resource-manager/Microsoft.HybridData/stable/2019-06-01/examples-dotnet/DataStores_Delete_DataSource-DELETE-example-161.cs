using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridData;
using Azure.ResourceManager.HybridData.Models;

// Generated from example definition: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataStores_Delete_DataSource-DELETE-example-161.json
// this example is just showing the usage of "DataStores_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridDataStoreResource created on azure
// for more information of creating HybridDataStoreResource, please refer to the document of HybridDataStoreResource
string subscriptionId = "6e0219f5-327a-4365-904f-05eed4227ad7";
string resourceGroupName = "ResourceGroupForSDKTest";
string dataManagerName = "TestAzureSDKOperations";
string dataStoreName = "TestStorSimpleSource1";
ResourceIdentifier hybridDataStoreResourceId = HybridDataStoreResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataManagerName, dataStoreName);
HybridDataStoreResource hybridDataStore = client.GetHybridDataStoreResource(hybridDataStoreResourceId);

// invoke the operation
await hybridDataStore.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
