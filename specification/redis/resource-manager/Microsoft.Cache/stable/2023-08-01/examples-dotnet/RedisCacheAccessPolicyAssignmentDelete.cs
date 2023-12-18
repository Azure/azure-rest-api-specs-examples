using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Redis;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheAccessPolicyAssignmentDelete.json
// this example is just showing the usage of "AccessPolicyAssignment_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisCacheAccessPolicyAssignmentResource created on azure
// for more information of creating RedisCacheAccessPolicyAssignmentResource, please refer to the document of RedisCacheAccessPolicyAssignmentResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string cacheName = "cache1";
string accessPolicyAssignmentName = "accessPolicyAssignmentName1";
ResourceIdentifier redisCacheAccessPolicyAssignmentResourceId = RedisCacheAccessPolicyAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName, accessPolicyAssignmentName);
RedisCacheAccessPolicyAssignmentResource redisCacheAccessPolicyAssignment = client.GetRedisCacheAccessPolicyAssignmentResource(redisCacheAccessPolicyAssignmentResourceId);

// invoke the operation
await redisCacheAccessPolicyAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
