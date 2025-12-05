using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBSqlRoleAssignmentDelete.json
// this example is just showing the usage of "SqlResources_DeleteSqlRoleAssignment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlRoleAssignmentResource created on azure
// for more information of creating CosmosDBSqlRoleAssignmentResource, please refer to the document of CosmosDBSqlRoleAssignmentResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string accountName = "myAccountName";
string roleAssignmentId = "myRoleAssignmentId";
ResourceIdentifier cosmosDBSqlRoleAssignmentResourceId = CosmosDBSqlRoleAssignmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, roleAssignmentId);
CosmosDBSqlRoleAssignmentResource cosmosDBSqlRoleAssignment = client.GetCosmosDBSqlRoleAssignmentResource(cosmosDBSqlRoleAssignmentResourceId);

// invoke the operation
await cosmosDBSqlRoleAssignment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
