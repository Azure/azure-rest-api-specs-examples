using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_GetStatus.json
// this example is just showing the usage of "IntegrationRuntimeStatus_Get" operation, for the dependent resources, they will have to be created separately.

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
SynapseIntegrationRuntimeStatusResult result = await synapseIntegrationRuntime.GetIntegrationRuntimeStatusAsync();

Console.WriteLine($"Succeeded: {result}");
