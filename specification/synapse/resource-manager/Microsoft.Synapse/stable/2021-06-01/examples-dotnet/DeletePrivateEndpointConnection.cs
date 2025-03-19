using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapsePrivateEndpointConnectionResource created on azure
// for more information of creating SynapsePrivateEndpointConnectionResource, please refer to the document of SynapsePrivateEndpointConnectionResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "ExampleResourceGroup";
string workspaceName = "ExampleWorkspace";
string privateEndpointConnectionName = "ExamplePrivateEndpointConnection";
ResourceIdentifier synapsePrivateEndpointConnectionResourceId = SynapsePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, privateEndpointConnectionName);
SynapsePrivateEndpointConnectionResource synapsePrivateEndpointConnection = client.GetSynapsePrivateEndpointConnectionResource(synapsePrivateEndpointConnectionResourceId);

// invoke the operation
await synapsePrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
