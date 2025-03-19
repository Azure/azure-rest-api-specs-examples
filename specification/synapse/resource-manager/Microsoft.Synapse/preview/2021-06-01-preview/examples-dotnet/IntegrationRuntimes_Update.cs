using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Update.json
// this example is just showing the usage of "IntegrationRuntimes_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseIntegrationRuntimeResource created on azure
// for more information of creating SynapseIntegrationRuntimeResource, please refer to the document of SynapseIntegrationRuntimeResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string workspaceName = "exampleWorkspace";
string integrationRuntimeName = "exampleIntegrationRuntime";
ResourceIdentifier synapseIntegrationRuntimeResourceId = SynapseIntegrationRuntimeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, integrationRuntimeName);
SynapseIntegrationRuntimeResource synapseIntegrationRuntime = client.GetSynapseIntegrationRuntimeResource(synapseIntegrationRuntimeResourceId);

// invoke the operation
SynapseIntegrationRuntimePatch patch = new SynapseIntegrationRuntimePatch
{
    AutoUpdate = SynapseIntegrationRuntimeAutoUpdate.Off,
    UpdateDelayOffset = "\"PT3H\"",
};
SynapseIntegrationRuntimeResource result = await synapseIntegrationRuntime.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseIntegrationRuntimeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
