using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineExamples/VirtualMachineExtensions_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineExtensions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineExtensionResource created on azure
// for more information of creating VirtualMachineExtensionResource, please refer to the document of VirtualMachineExtensionResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string vmName = "aaaaaaaaaaaaa";
string vmExtensionName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier virtualMachineExtensionResourceId = VirtualMachineExtensionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName, vmExtensionName);
VirtualMachineExtensionResource virtualMachineExtension = client.GetVirtualMachineExtensionResource(virtualMachineExtensionResourceId);

// invoke the operation
await virtualMachineExtension.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
