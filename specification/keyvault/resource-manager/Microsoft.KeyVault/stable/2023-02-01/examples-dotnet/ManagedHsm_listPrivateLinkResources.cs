using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.KeyVault;
using Azure.ResourceManager.KeyVault.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-02-01/examples/ManagedHsm_listPrivateLinkResources.json
// this example is just showing the usage of "MHSMPrivateLinkResources_ListByMHSMResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedHsmResource created on azure
// for more information of creating ManagedHsmResource, please refer to the document of ManagedHsmResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "sample-group";
string name = "sample-mhsm";
ResourceIdentifier managedHsmResourceId = ManagedHsmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
ManagedHsmResource managedHsm = client.GetManagedHsmResource(managedHsmResourceId);

// invoke the operation and iterate over the result
await foreach (ManagedHsmPrivateLinkResourceData item in managedHsm.GetMHSMPrivateLinkResourcesByManagedHsmResourceAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
