using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsListByKustoPool.json
// this example is just showing the usage of "KustoPoolAttachedDatabaseConfigurations_ListByKustoPool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseKustoPoolResource created on azure
// for more information of creating SynapseKustoPoolResource, please refer to the document of SynapseKustoPoolResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string workspaceName = "kustorptest";
string kustoPoolName = "kustoclusterrptest4";
ResourceIdentifier synapseKustoPoolResourceId = SynapseKustoPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, kustoPoolName);
SynapseKustoPoolResource synapseKustoPool = client.GetSynapseKustoPoolResource(synapseKustoPoolResourceId);

// get the collection of this SynapseAttachedDatabaseConfigurationResource
SynapseAttachedDatabaseConfigurationCollection collection = synapseKustoPool.GetSynapseAttachedDatabaseConfigurations();

// invoke the operation and iterate over the result
await foreach (SynapseAttachedDatabaseConfigurationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SynapseAttachedDatabaseConfigurationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
