using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteNetworkInterface.json
// this example is just showing the usage of "NetworkInterfaces_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkInterfaceResource created on azure
// for more information of creating NetworkInterfaceResource, please refer to the document of NetworkInterfaceResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string networkInterfaceName = "test-nic";
ResourceIdentifier networkInterfaceResourceId = NetworkInterfaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkInterfaceName);
NetworkInterfaceResource networkInterface = client.GetNetworkInterfaceResource(networkInterfaceResourceId);

// invoke the operation
await networkInterface.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
