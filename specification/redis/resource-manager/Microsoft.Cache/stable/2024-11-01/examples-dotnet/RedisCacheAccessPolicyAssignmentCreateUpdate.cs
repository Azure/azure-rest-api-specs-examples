using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Redis;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2024-11-01/examples/RedisCacheAccessPolicyAssignmentCreateUpdate.json
// this example is just showing the usage of "AccessPolicyAssignment_CreateUpdate" operation, for the dependent resources, they will have to be created separately.

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
RedisCacheAccessPolicyAssignmentData data = new RedisCacheAccessPolicyAssignmentData
{
    ObjectId = Guid.Parse("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
    ObjectIdAlias = "TestAADAppRedis",
    AccessPolicyName = "accessPolicy1",
};
ArmOperation<RedisCacheAccessPolicyAssignmentResource> lro = await redisCacheAccessPolicyAssignment.UpdateAsync(WaitUntil.Completed, data);
RedisCacheAccessPolicyAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisCacheAccessPolicyAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
