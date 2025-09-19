using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Vm.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Hci.Vm;

// Generated from example definition: 2025-06-01-preview/VirtualMachineInstances_Get.json
// this example is just showing the usage of "VirtualMachineInstance_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciVmInstanceResource created on azure
// for more information of creating HciVmInstanceResource, please refer to the document of HciVmInstanceResource
string resourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM";
ResourceIdentifier hciVmInstanceResourceId = HciVmInstanceResource.CreateResourceIdentifier(resourceUri);
HciVmInstanceResource hciVmInstance = client.GetHciVmInstanceResource(hciVmInstanceResourceId);

// invoke the operation
HciVmInstanceResource result = await hciVmInstance.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciVmInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
