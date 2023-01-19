using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridData;
using Azure.ResourceManager.HybridData.Models;

// Generated from example definition: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_ListByDataService-GET-example-71.json
// this example is just showing the usage of "JobDefinitions_ListByDataService" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridDataServiceResource created on azure
// for more information of creating HybridDataServiceResource, please refer to the document of HybridDataServiceResource
string subscriptionId = "6e0219f5-327a-4365-904f-05eed4227ad7";
string resourceGroupName = "ResourceGroupForSDKTest";
string dataManagerName = "TestAzureSDKOperations";
string dataServiceName = "DataTransformation";
ResourceIdentifier hybridDataServiceResourceId = HybridDataServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataManagerName, dataServiceName);
HybridDataServiceResource hybridDataService = client.GetHybridDataServiceResource(hybridDataServiceResourceId);

// get the collection of this HybridDataJobDefinitionResource
HybridDataJobDefinitionCollection collection = hybridDataService.GetHybridDataJobDefinitions();

// invoke the operation and iterate over the result
await foreach (HybridDataJobDefinitionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HybridDataJobDefinitionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
