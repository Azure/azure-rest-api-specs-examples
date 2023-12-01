using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteVirtualMachineInstance.json
// this example is just showing the usage of "VirtualMachineInstances_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineInstanceResource created on azure
// for more information of creating VirtualMachineInstanceResource, please refer to the document of VirtualMachineInstanceResource
string resourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM";
ResourceIdentifier virtualMachineInstanceResourceId = VirtualMachineInstanceResource.CreateResourceIdentifier(resourceUri);
VirtualMachineInstanceResource virtualMachineInstance = client.GetVirtualMachineInstanceResource(virtualMachineInstanceResourceId);

// invoke the operation
await virtualMachineInstance.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
