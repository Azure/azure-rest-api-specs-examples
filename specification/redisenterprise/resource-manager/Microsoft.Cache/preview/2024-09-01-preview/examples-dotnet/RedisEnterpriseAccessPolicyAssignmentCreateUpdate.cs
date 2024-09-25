using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RedisEnterprise;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseAccessPolicyAssignmentCreateUpdate.json
// this example is just showing the usage of "AccessPolicyAssignment_CreateUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisEnterpriseDatabaseResource created on azure
// for more information of creating RedisEnterpriseDatabaseResource, please refer to the document of RedisEnterpriseDatabaseResource
string subscriptionId = "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
string resourceGroupName = "rg1";
string clusterName = "cache1";
string databaseName = "default";
ResourceIdentifier redisEnterpriseDatabaseResourceId = RedisEnterpriseDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName);
RedisEnterpriseDatabaseResource redisEnterpriseDatabase = client.GetRedisEnterpriseDatabaseResource(redisEnterpriseDatabaseResourceId);

// get the collection of this AccessPolicyAssignmentResource
AccessPolicyAssignmentCollection collection = redisEnterpriseDatabase.GetAccessPolicyAssignments();

// invoke the operation
string accessPolicyAssignmentName = "defaultTestEntraApp1";
AccessPolicyAssignmentData data = new AccessPolicyAssignmentData()
{
    AccessPolicyName = "default",
    UserObjectId = Guid.Parse("6497c918-11ad-41e7-1b0f-7c518a87d0b0"),
};
ArmOperation<AccessPolicyAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accessPolicyAssignmentName, data);
AccessPolicyAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AccessPolicyAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
