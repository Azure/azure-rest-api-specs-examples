using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CosmosDB;
using Azure.ResourceManager.CosmosDB.Models;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-09-15-preview/examples/CosmosDBSqlStoredProcedureDelete.json
// this example is just showing the usage of "SqlResources_DeleteSqlStoredProcedure" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlStoredProcedureResource created on azure
// for more information of creating CosmosDBSqlStoredProcedureResource, please refer to the document of CosmosDBSqlStoredProcedureResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
string containerName = "containerName";
string storedProcedureName = "storedProcedureName";
ResourceIdentifier cosmosDBSqlStoredProcedureResourceId = CosmosDBSqlStoredProcedureResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName, containerName, storedProcedureName);
CosmosDBSqlStoredProcedureResource cosmosDBSqlStoredProcedure = client.GetCosmosDBSqlStoredProcedureResource(cosmosDBSqlStoredProcedureResourceId);

// invoke the operation
await cosmosDBSqlStoredProcedure.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
