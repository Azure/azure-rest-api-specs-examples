using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBMongoDBRoleDefinitionCreateUpdate.json
// this example is just showing the usage of "MongoDBResources_CreateUpdateMongoRoleDefinition" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoDBRoleDefinitionResource created on azure
// for more information of creating MongoDBRoleDefinitionResource, please refer to the document of MongoDBRoleDefinitionResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string accountName = "myAccountName";
string mongoRoleDefinitionId = "myMongoRoleDefinitionId";
ResourceIdentifier mongoDBRoleDefinitionResourceId = MongoDBRoleDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, mongoRoleDefinitionId);
MongoDBRoleDefinitionResource mongoDBRoleDefinition = client.GetMongoDBRoleDefinitionResource(mongoDBRoleDefinitionResourceId);

// invoke the operation
MongoDBRoleDefinitionCreateOrUpdateContent content = new MongoDBRoleDefinitionCreateOrUpdateContent
{
    RoleName = "myRoleName",
    DatabaseName = "sales",
    Privileges = {new MongoDBPrivilege
    {
    Resource = new MongoDBPrivilegeResourceInfo
    {
    DBName = "sales",
    Collection = "sales",
    },
    Actions = {"insert", "find"},
    }},
    Roles = {new MongoDBRole
    {
    DBName = "sales",
    Role = "myInheritedRole",
    }},
};
ArmOperation<MongoDBRoleDefinitionResource> lro = await mongoDBRoleDefinition.UpdateAsync(WaitUntil.Completed, content);
MongoDBRoleDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoDBRoleDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
