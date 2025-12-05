using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBDatabaseAccountFailoverPriorityChange.json
// this example is just showing the usage of "DatabaseAccounts_FailoverPriorityChange" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBAccountResource created on azure
// for more information of creating CosmosDBAccountResource, please refer to the document of CosmosDBAccountResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1-failover";
ResourceIdentifier cosmosDBAccountResourceId = CosmosDBAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CosmosDBAccountResource cosmosDBAccount = client.GetCosmosDBAccountResource(cosmosDBAccountResourceId);

// invoke the operation
CosmosDBFailoverPolicies failoverParameters = new CosmosDBFailoverPolicies(new CosmosDBFailoverPolicy[]
{
new CosmosDBFailoverPolicy
{
LocationName = new AzureLocation("eastus"),
FailoverPriority = 0,
},
new CosmosDBFailoverPolicy
{
LocationName = new AzureLocation("westus"),
FailoverPriority = 1,
}
});
await cosmosDBAccount.FailoverPriorityChangeAsync(WaitUntil.Completed, failoverParameters);

Console.WriteLine("Succeeded");
