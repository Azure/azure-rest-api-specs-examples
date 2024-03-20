using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Quantum;
using Azure.ResourceManager.Quantum.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/quantumWorkspacesListResourceGroup.json
// this example is just showing the usage of "Workspaces_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "quantumResourcegroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this QuantumWorkspaceResource
QuantumWorkspaceCollection collection = resourceGroupResource.GetQuantumWorkspaces();

// invoke the operation and iterate over the result
await foreach (QuantumWorkspaceResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    QuantumWorkspaceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
