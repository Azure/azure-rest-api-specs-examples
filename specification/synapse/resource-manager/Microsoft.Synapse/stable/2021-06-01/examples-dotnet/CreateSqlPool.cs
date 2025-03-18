using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPool.json
// this example is just showing the usage of "SqlPools_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseWorkspaceResource created on azure
// for more information of creating SynapseWorkspaceResource, please refer to the document of SynapseWorkspaceResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "ExampleResourceGroup";
string workspaceName = "ExampleWorkspace";
ResourceIdentifier synapseWorkspaceResourceId = SynapseWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
SynapseWorkspaceResource synapseWorkspace = client.GetSynapseWorkspaceResource(synapseWorkspaceResourceId);

// get the collection of this SynapseSqlPoolResource
SynapseSqlPoolCollection collection = synapseWorkspace.GetSynapseSqlPools();

// invoke the operation
string sqlPoolName = "ExampleSqlPool";
SynapseSqlPoolData data = new SynapseSqlPoolData(new AzureLocation("Southeast Asia"))
{
    Sku = new SynapseSku
    {
        Tier = "",
        Name = "",
    },
    MaxSizeBytes = 0L,
    Collation = "",
    SourceDatabaseId = "",
    RecoverableDatabaseId = "",
    CreateMode = new SqlPoolCreateMode(""),
    StorageAccountType = SqlPoolStorageAccountType.Lrs,
    Tags = { },
};
ArmOperation<SynapseSqlPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sqlPoolName, data);
SynapseSqlPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseSqlPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
