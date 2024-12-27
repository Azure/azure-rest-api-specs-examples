using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBTableBackupInformation.json
// this example is just showing the usage of "TableResources_RetrieveContinuousBackupInformation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBTableResource created on azure
// for more information of creating CosmosDBTableResource, please refer to the document of CosmosDBTableResource
string subscriptionId = "subid";
string resourceGroupName = "rgName";
string accountName = "ddb1";
string tableName = "tableName1";
ResourceIdentifier cosmosDBTableResourceId = CosmosDBTableResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, tableName);
CosmosDBTableResource cosmosDBTable = client.GetCosmosDBTableResource(cosmosDBTableResourceId);

// invoke the operation
ContinuousBackupRestoreLocation location = new ContinuousBackupRestoreLocation
{
    Location = new AzureLocation("North Europe"),
};
ArmOperation<CosmosDBBackupInformation> lro = await cosmosDBTable.RetrieveContinuousBackupInformationAsync(WaitUntil.Completed, location);
CosmosDBBackupInformation result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
