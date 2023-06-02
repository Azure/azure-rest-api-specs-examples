using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetAzureADOnlyAuthentication.json
// this example is just showing the usage of "AzureADOnlyAuthentications_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseWorkspaceResource created on azure
// for more information of creating SynapseWorkspaceResource, please refer to the document of SynapseWorkspaceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "workspace-6852";
string workspaceName = "workspace-2080";
ResourceIdentifier synapseWorkspaceResourceId = SynapseWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
SynapseWorkspaceResource synapseWorkspace = client.GetSynapseWorkspaceResource(synapseWorkspaceResourceId);

// get the collection of this SynapseAadOnlyAuthenticationResource
SynapseAadOnlyAuthenticationCollection collection = synapseWorkspace.GetSynapseAadOnlyAuthentications();

// invoke the operation
SynapseAadOnlyAuthenticationName azureADOnlyAuthenticationName = SynapseAadOnlyAuthenticationName.Default;
bool result = await collection.ExistsAsync(azureADOnlyAuthenticationName);

Console.WriteLine($"Succeeded: {result}");
