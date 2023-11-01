using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBSqlContainerCreateUpdate.json
// this example is just showing the usage of "SqlResources_CreateUpdateSqlContainer" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlDatabaseResource created on azure
// for more information of creating CosmosDBSqlDatabaseResource, please refer to the document of CosmosDBSqlDatabaseResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
ResourceIdentifier cosmosDBSqlDatabaseResourceId = CosmosDBSqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName);
CosmosDBSqlDatabaseResource cosmosDBSqlDatabase = client.GetCosmosDBSqlDatabaseResource(cosmosDBSqlDatabaseResourceId);

// get the collection of this CosmosDBSqlContainerResource
CosmosDBSqlContainerCollection collection = cosmosDBSqlDatabase.GetCosmosDBSqlContainers();

// invoke the operation
string containerName = "containerName";
CosmosDBSqlContainerCreateOrUpdateContent content = new CosmosDBSqlContainerCreateOrUpdateContent(new AzureLocation("West US"), new CosmosDBSqlContainerResourceInfo("containerName")
{
    IndexingPolicy = new CosmosDBIndexingPolicy()
    {
        IsAutomatic = true,
        IndexingMode = CosmosDBIndexingMode.Consistent,
        IncludedPaths =
        {
        new CosmosDBIncludedPath()
        {
        Path = "/*",
        Indexes =
        {
        new CosmosDBPathIndexes()
        {
        DataType = CosmosDBDataType.String,
        Precision = -1,
        Kind = CosmosDBIndexKind.Range,
        },new CosmosDBPathIndexes()
        {
        DataType = CosmosDBDataType.Number,
        Precision = -1,
        Kind = CosmosDBIndexKind.Range,
        }
        },
        }
        },
        ExcludedPaths =
        {
        },
    },
    PartitionKey = new CosmosDBContainerPartitionKey()
    {
        Paths =
        {
        "/AccountNumber"
        },
        Kind = CosmosDBPartitionKind.Hash,
    },
    DefaultTtl = 100,
    UniqueKeys =
    {
    new CosmosDBUniqueKey()
    {
    Paths =
    {
    "/testPath"
    },
    }
    },
    ConflictResolutionPolicy = new ConflictResolutionPolicy()
    {
        Mode = ConflictResolutionMode.LastWriterWins,
        ConflictResolutionPath = "/path",
    },
    ClientEncryptionPolicy = new CosmosDBClientEncryptionPolicy(new CosmosDBClientEncryptionIncludedPath[]
{
new CosmosDBClientEncryptionIncludedPath("/path","keyId","Deterministic","AEAD_AES_256_CBC_HMAC_SHA256")
}, 2),
})
{
    Options = new CosmosDBCreateUpdateConfig(),
    Tags =
    {
    },
};
ArmOperation<CosmosDBSqlContainerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, containerName, content);
CosmosDBSqlContainerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBSqlContainerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
