using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-05-15-preview/examples/CosmosDBCassandraViewCreateUpdate.json
// this example is just showing the usage of "CassandraResources_CreateUpdateCassandraView" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CassandraKeyspaceResource created on azure
// for more information of creating CassandraKeyspaceResource, please refer to the document of CassandraKeyspaceResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string accountName = "ddb1";
string keyspaceName = "keyspacename";
ResourceIdentifier cassandraKeyspaceResourceId = CassandraKeyspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, keyspaceName);
CassandraKeyspaceResource cassandraKeyspace = client.GetCassandraKeyspaceResource(cassandraKeyspaceResourceId);

// get the collection of this CassandraViewGetResultResource
CassandraViewGetResultCollection collection = cassandraKeyspace.GetCassandraViewGetResults();

// invoke the operation
string viewName = "viewname";
CassandraViewGetResultCreateOrUpdateContent content = new CassandraViewGetResultCreateOrUpdateContent(new AzureLocation("placeholder"), new CassandraViewResource("viewname", "SELECT columna, columnb, columnc FROM keyspacename.srctablename WHERE columna IS NOT NULL AND columnc IS NOT NULL PRIMARY (columnc, columna)"))
{
    Options = new CosmosDBCreateUpdateConfig(),
    Tags =
    {
    },
};
ArmOperation<CassandraViewGetResultResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, viewName, content);
CassandraViewGetResultResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CassandraViewGetResultData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
