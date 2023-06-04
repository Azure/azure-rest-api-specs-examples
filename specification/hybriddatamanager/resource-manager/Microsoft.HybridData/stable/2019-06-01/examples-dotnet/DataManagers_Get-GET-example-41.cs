using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridData;
using Azure.ResourceManager.HybridData.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/DataManagers_Get-GET-example-41.json
// this example is just showing the usage of "DataManagers_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridDataManagerResource created on azure
// for more information of creating HybridDataManagerResource, please refer to the document of HybridDataManagerResource
string subscriptionId = "6e0219f5-327a-4365-904f-05eed4227ad7";
string resourceGroupName = "ResourceGroupForSDKTest";
string dataManagerName = "TestAzureSDKOperations";
ResourceIdentifier hybridDataManagerResourceId = HybridDataManagerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataManagerName);
HybridDataManagerResource hybridDataManager = client.GetHybridDataManagerResource(hybridDataManagerResourceId);

// invoke the operation
HybridDataManagerResource result = await hybridDataManager.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridDataManagerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
