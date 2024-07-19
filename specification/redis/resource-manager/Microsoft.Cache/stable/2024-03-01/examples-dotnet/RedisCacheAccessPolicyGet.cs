using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Redis;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheAccessPolicyGet.json
// this example is just showing the usage of "AccessPolicy_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisCacheAccessPolicyResource created on azure
// for more information of creating RedisCacheAccessPolicyResource, please refer to the document of RedisCacheAccessPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string cacheName = "cache1";
string accessPolicyName = "accessPolicy1";
ResourceIdentifier redisCacheAccessPolicyResourceId = RedisCacheAccessPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName, accessPolicyName);
RedisCacheAccessPolicyResource redisCacheAccessPolicy = client.GetRedisCacheAccessPolicyResource(redisCacheAccessPolicyResourceId);

// invoke the operation
RedisCacheAccessPolicyResource result = await redisCacheAccessPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisCacheAccessPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
