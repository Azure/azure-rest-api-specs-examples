using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/CosmosDBSqlContainerRestore.json
// this example is just showing the usage of "SqlResources_CreateUpdateSqlContainer" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlContainerResource created on azure
// for more information of creating CosmosDBSqlContainerResource, please refer to the document of CosmosDBSqlContainerResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
string containerName = "containerName";
ResourceIdentifier cosmosDBSqlContainerResourceId = CosmosDBSqlContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName, containerName);
CosmosDBSqlContainerResource cosmosDBSqlContainer = client.GetCosmosDBSqlContainerResource(cosmosDBSqlContainerResourceId);

// invoke the operation
CosmosDBSqlContainerCreateOrUpdateContent content = new CosmosDBSqlContainerCreateOrUpdateContent(new AzureLocation("West US"), new CosmosDBSqlContainerResourceInfo("containerName")
{
    RestoreParameters = new ResourceRestoreParameters()
    {
        RestoreSource = "/subscriptions/subid/providers/Microsoft.DocumentDB/locations/WestUS/restorableDatabaseAccounts/restorableDatabaseAccountId",
        RestoreTimestampInUtc = DateTimeOffset.Parse("2022-07-20T18:28:00Z"),
        IsRestoreWithTtlDisabled = true,
    },
    CreateMode = CosmosDBAccountCreateMode.Restore,
})
{
    Options = new CosmosDBCreateUpdateConfig(),
    Tags =
    {
    },
};
ArmOperation<CosmosDBSqlContainerResource> lro = await cosmosDBSqlContainer.UpdateAsync(WaitUntil.Completed, content);
CosmosDBSqlContainerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBSqlContainerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
