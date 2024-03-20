using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Quantum;
using Azure.ResourceManager.Quantum.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/quantumWorkspacesPut.json
// this example is just showing the usage of "Workspaces_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
string workspaceName = "quantumworkspace1";
QuantumWorkspaceData data = new QuantumWorkspaceData(new AzureLocation("West US"))
{
    Providers =
    {
    new QuantumProvider()
    {
    ProviderId = "Honeywell",
    ProviderSku = "Basic",
    },new QuantumProvider()
    {
    ProviderId = "IonQ",
    ProviderSku = "Basic",
    },new QuantumProvider()
    {
    ProviderId = "OneQBit",
    ProviderSku = "Basic",
    }
    },
    StorageAccount = "/subscriptions/1C4B2828-7D49-494F-933D-061373BE28C2/resourceGroups/quantumResourcegroup/providers/Microsoft.Storage/storageAccounts/testStorageAccount",
};
ArmOperation<QuantumWorkspaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, workspaceName, data);
QuantumWorkspaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
QuantumWorkspaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
