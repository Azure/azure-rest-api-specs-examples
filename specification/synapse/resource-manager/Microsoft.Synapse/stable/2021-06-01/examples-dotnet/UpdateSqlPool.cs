using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateSqlPool.json
// this example is just showing the usage of "SqlPools_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseSqlPoolResource created on azure
// for more information of creating SynapseSqlPoolResource, please refer to the document of SynapseSqlPoolResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "ExampleResourceGroup";
string workspaceName = "ExampleWorkspace";
string sqlPoolName = "ExampleSqlPool";
ResourceIdentifier synapseSqlPoolResourceId = SynapseSqlPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName);
SynapseSqlPoolResource synapseSqlPool = client.GetSynapseSqlPoolResource(synapseSqlPoolResourceId);

// invoke the operation
SynapseSqlPoolPatch patch = new SynapseSqlPoolPatch()
{
    Tags =
    {
    },
    Location = new AzureLocation("West US 2"),
    Sku = new SynapseSku()
    {
        Tier = "",
        Name = "",
    },
    MaxSizeBytes = 0,
    Collation = "",
    RestorePointInTime = DateTimeOffset.Parse("1970-01-01T00:00:00.000Z"),
};
ArmOperation<SynapseSqlPoolResource> lro = await synapseSqlPool.UpdateAsync(WaitUntil.Completed, patch);
SynapseSqlPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseSqlPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
