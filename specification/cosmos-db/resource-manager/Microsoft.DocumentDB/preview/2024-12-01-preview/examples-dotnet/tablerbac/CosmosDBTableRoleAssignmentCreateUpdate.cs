using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/tablerbac/CosmosDBTableRoleAssignmentCreateUpdate.json
// this example is just showing the usage of "TableResources_CreateUpdateTableRoleAssignment" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBAccountResource created on azure
// for more information of creating CosmosDBAccountResource, please refer to the document of CosmosDBAccountResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "myResourceGroupName";
string accountName = "myAccountName";
ResourceIdentifier cosmosDBAccountResourceId = CosmosDBAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CosmosDBAccountResource cosmosDBAccount = client.GetCosmosDBAccountResource(cosmosDBAccountResourceId);

// get the collection of this CosmosDBTableRoleAssignmentResource
CosmosDBTableRoleAssignmentCollection collection = cosmosDBAccount.GetCosmosDBTableRoleAssignments();

// invoke the operation
string roleAssignmentId = "myRoleAssignmentId";
CosmosDBTableRoleAssignmentData data = new CosmosDBTableRoleAssignmentData
{
    RoleDefinitionId = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/tableRoleDefinitions/myRoleDefinitionId"),
    Scope = new ResourceIdentifier("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroupName/providers/Microsoft.DocumentDB/databaseAccounts/myAccountName/dbs/purchases/colls/redmond-purchases"),
    PrincipalId = Guid.Parse("myPrincipalId"),
};
ArmOperation<CosmosDBTableRoleAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, roleAssignmentId, data);
CosmosDBTableRoleAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBTableRoleAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
