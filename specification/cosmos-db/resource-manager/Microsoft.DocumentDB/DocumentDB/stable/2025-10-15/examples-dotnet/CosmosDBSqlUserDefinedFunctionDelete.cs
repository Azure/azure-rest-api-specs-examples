using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBSqlUserDefinedFunctionDelete.json
// this example is just showing the usage of "SqlResources_DeleteSqlUserDefinedFunction" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlUserDefinedFunctionResource created on azure
// for more information of creating CosmosDBSqlUserDefinedFunctionResource, please refer to the document of CosmosDBSqlUserDefinedFunctionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string databaseName = "databaseName";
string containerName = "containerName";
string userDefinedFunctionName = "userDefinedFunctionName";
ResourceIdentifier cosmosDBSqlUserDefinedFunctionResourceId = CosmosDBSqlUserDefinedFunctionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName, containerName, userDefinedFunctionName);
CosmosDBSqlUserDefinedFunctionResource cosmosDBSqlUserDefinedFunction = client.GetCosmosDBSqlUserDefinedFunctionResource(cosmosDBSqlUserDefinedFunctionResourceId);

// invoke the operation
await cosmosDBSqlUserDefinedFunction.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
