using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ConnectedVMwarevSphere;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteVirtualNetwork.json
// this example is just showing the usage of "VirtualNetworks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VMwareVirtualNetworkResource created on azure
// for more information of creating VMwareVirtualNetworkResource, please refer to the document of VMwareVirtualNetworkResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string virtualNetworkName = "ProdNetwork";
ResourceIdentifier vMwareVirtualNetworkResourceId = VMwareVirtualNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkName);
VMwareVirtualNetworkResource vMwareVirtualNetwork = client.GetVMwareVirtualNetworkResource(vMwareVirtualNetworkResourceId);

// invoke the operation
await vMwareVirtualNetwork.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
