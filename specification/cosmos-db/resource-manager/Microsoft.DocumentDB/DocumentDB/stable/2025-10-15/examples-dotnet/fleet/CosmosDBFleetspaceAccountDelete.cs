using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/fleet/CosmosDBFleetspaceAccountDelete.json
// this example is just showing the usage of "FleetspaceAccount_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBFleetspaceAccountResource created on azure
// for more information of creating CosmosDBFleetspaceAccountResource, please refer to the document of CosmosDBFleetspaceAccountResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
string fleetspaceName = "fleetspace1";
string fleetspaceAccountName = "db1";
ResourceIdentifier cosmosDBFleetspaceAccountResourceId = CosmosDBFleetspaceAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName, fleetspaceName, fleetspaceAccountName);
CosmosDBFleetspaceAccountResource cosmosDBFleetspaceAccount = client.GetCosmosDBFleetspaceAccountResource(cosmosDBFleetspaceAccountResourceId);

// invoke the operation
await cosmosDBFleetspaceAccount.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
