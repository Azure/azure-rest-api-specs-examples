using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RedisEnterprise;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2025-04-01/examples/RedisEnterpriseAccessPolicyAssignmentDelete.json
// this example is just showing the usage of "AccessPolicyAssignment_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AccessPolicyAssignmentResource created on azure
// for more information of creating AccessPolicyAssignmentResource, please refer to the document of AccessPolicyAssignmentResource
string subscriptionId = "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
string resourceGroupName = "rg1";
string clusterName = "cache1";
string databaseName = "default";
string accessPolicyAssignmentName = "defaultTestEntraApp1";
ResourceIdentifier accessPolicyAssignmentResourceId = AccessPolicyAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName, accessPolicyAssignmentName);
AccessPolicyAssignmentResource accessPolicyAssignment = client.GetAccessPolicyAssignmentResource(accessPolicyAssignmentResourceId);

// invoke the operation
await accessPolicyAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
