using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrivateLinkResourcesList.json
// this example is just showing the usage of "KustoPoolPrivateLinkResources_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseKustoPoolResource created on azure
// for more information of creating SynapseKustoPoolResource, please refer to the document of SynapseKustoPoolResource
string subscriptionId = "7a587823-959d-4ad0-85bd-cf2a7cef436a";
string resourceGroupName = "DP-900";
string workspaceName = "synapse-ws-ebi-data";
string kustoPoolName = "dataexplorerpool900";
ResourceIdentifier synapseKustoPoolResourceId = SynapseKustoPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, kustoPoolName);
SynapseKustoPoolResource synapseKustoPool = client.GetSynapseKustoPoolResource(synapseKustoPoolResourceId);

// invoke the operation and iterate over the result
await foreach (SynapseKustoPoolPrivateLinkData item in synapseKustoPool.GetAllKustoPoolPrivateLinkDataAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
