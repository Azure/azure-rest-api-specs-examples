using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/CosmosDBCassandraKeyspaceDelete.json
// this example is just showing the usage of "CassandraResources_DeleteCassandraKeyspace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CassandraKeyspaceResource created on azure
// for more information of creating CassandraKeyspaceResource, please refer to the document of CassandraKeyspaceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string keyspaceName = "keyspaceName";
ResourceIdentifier cassandraKeyspaceResourceId = CassandraKeyspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, keyspaceName);
CassandraKeyspaceResource cassandraKeyspace = client.GetCassandraKeyspaceResource(cassandraKeyspaceResourceId);

// invoke the operation
await cassandraKeyspace.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
