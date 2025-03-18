using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasesUpdate.json
// this example is just showing the usage of "KustoPoolDatabases_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseDatabaseResource created on azure
// for more information of creating SynapseDatabaseResource, please refer to the document of SynapseDatabaseResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string workspaceName = "synapseWorkspaceName";
string kustoPoolName = "kustoclusterrptest4";
string databaseName = "KustoDatabase8";
ResourceIdentifier synapseDatabaseResourceId = SynapseDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, kustoPoolName, databaseName);
SynapseDatabaseResource synapseDatabase = client.GetSynapseDatabaseResource(synapseDatabaseResourceId);

// invoke the operation
SynapseDatabaseData data = new SynapseReadWriteDatabase
{
    SoftDeletePeriod = XmlConvert.ToTimeSpan("P1D"),
};
ArmOperation<SynapseDatabaseResource> lro = await synapseDatabase.UpdateAsync(WaitUntil.Completed, data);
SynapseDatabaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseDatabaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
