using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBSqlTriggerCreateUpdate.json
// this example is just showing the usage of "SqlResources_CreateUpdateSqlTrigger" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlTriggerResource created on azure
// for more information of creating CosmosDBSqlTriggerResource, please refer to the document of CosmosDBSqlTriggerResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
string containerName = "containerName";
string triggerName = "triggerName";
ResourceIdentifier cosmosDBSqlTriggerResourceId = CosmosDBSqlTriggerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName, containerName, triggerName);
CosmosDBSqlTriggerResource cosmosDBSqlTrigger = client.GetCosmosDBSqlTriggerResource(cosmosDBSqlTriggerResourceId);

// invoke the operation
CosmosDBSqlTriggerCreateOrUpdateContent content = new CosmosDBSqlTriggerCreateOrUpdateContent(default, new CosmosDBSqlTriggerResourceInfo("triggerName")
{
    Body = "body",
    TriggerType = new CosmosDBSqlTriggerType("triggerType"),
    TriggerOperation = new CosmosDBSqlTriggerOperation("triggerOperation"),
})
{
    Options = new CosmosDBCreateUpdateConfig(),
};
ArmOperation<CosmosDBSqlTriggerResource> lro = await cosmosDBSqlTrigger.UpdateAsync(WaitUntil.Completed, content);
CosmosDBSqlTriggerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBSqlTriggerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
