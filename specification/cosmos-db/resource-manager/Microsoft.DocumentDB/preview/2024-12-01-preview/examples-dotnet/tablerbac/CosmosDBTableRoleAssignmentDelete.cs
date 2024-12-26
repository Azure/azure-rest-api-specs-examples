using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/tablerbac/CosmosDBTableRoleAssignmentDelete.json
// this example is just showing the usage of "TableResources_DeleteTableRoleAssignment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBTableRoleAssignmentResource created on azure
// for more information of creating CosmosDBTableRoleAssignmentResource, please refer to the document of CosmosDBTableRoleAssignmentResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "myResourceGroupName";
string accountName = "myAccountName";
string roleAssignmentId = "myRoleAssignmentId";
ResourceIdentifier cosmosDBTableRoleAssignmentResourceId = CosmosDBTableRoleAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, roleAssignmentId);
CosmosDBTableRoleAssignmentResource cosmosDBTableRoleAssignment = client.GetCosmosDBTableRoleAssignmentResource(cosmosDBTableRoleAssignmentResourceId);

// invoke the operation
await cosmosDBTableRoleAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
