using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolTableGet.json
// this example is just showing the usage of "SqlPoolTables_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseSqlPoolSchemaResource created on azure
// for more information of creating SynapseSqlPoolSchemaResource, please refer to the document of SynapseSqlPoolSchemaResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myRG";
string workspaceName = "serverName";
string sqlPoolName = "myDatabase";
string schemaName = "dbo";
ResourceIdentifier synapseSqlPoolSchemaResourceId = SynapseSqlPoolSchemaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName, schemaName);
SynapseSqlPoolSchemaResource synapseSqlPoolSchema = client.GetSynapseSqlPoolSchemaResource(synapseSqlPoolSchemaResourceId);

// get the collection of this SynapseSqlPoolTableResource
SynapseSqlPoolTableCollection collection = synapseSqlPoolSchema.GetSynapseSqlPoolTables();

// invoke the operation
string tableName = "table1";
bool result = await collection.ExistsAsync(tableName);

Console.WriteLine($"Succeeded: {result}");
