using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBTableDelete.json
// this example is just showing the usage of "TableResources_DeleteTable" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBTableResource created on azure
// for more information of creating CosmosDBTableResource, please refer to the document of CosmosDBTableResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string tableName = "tableName";
ResourceIdentifier cosmosDBTableResourceId = CosmosDBTableResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, tableName);
CosmosDBTableResource cosmosDBTable = client.GetCosmosDBTableResource(cosmosDBTableResourceId);

// invoke the operation
await cosmosDBTable.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
