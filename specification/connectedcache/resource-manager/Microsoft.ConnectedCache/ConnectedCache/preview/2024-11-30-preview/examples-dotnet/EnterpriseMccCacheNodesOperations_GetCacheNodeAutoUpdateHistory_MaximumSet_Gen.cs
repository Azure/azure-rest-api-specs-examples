using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConnectedCache.Models;
using Azure.ResourceManager.ConnectedCache;

// Generated from example definition: 2024-11-30-preview/EnterpriseMccCacheNodesOperations_GetCacheNodeAutoUpdateHistory_MaximumSet_Gen.json
// this example is just showing the usage of "EnterpriseMccCacheNodesOperations_GetCacheNodeAutoUpdateHistory" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EnterpriseMccCacheNodeResource created on azure
// for more information of creating EnterpriseMccCacheNodeResource, please refer to the document of EnterpriseMccCacheNodeResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "rgConnectedCache";
string customerResourceName = "MccRPTest1";
string cacheNodeResourceName = "MCCCachenode1";
ResourceIdentifier enterpriseMccCacheNodeResourceId = EnterpriseMccCacheNodeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, customerResourceName, cacheNodeResourceName);
EnterpriseMccCacheNodeResource enterpriseMccCacheNode = client.GetEnterpriseMccCacheNodeResource(enterpriseMccCacheNodeResourceId);

// invoke the operation
MccCacheNodeAutoUpdateHistoryData result = await enterpriseMccCacheNode.GetCacheNodeAutoUpdateHistoryAsync();

Console.WriteLine($"Succeeded: {result}");
