using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/virtualMachineExamples/VirtualMachine_Capture_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachines_Capture" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineResource created on azure
// for more information of creating VirtualMachineResource, please refer to the document of VirtualMachineResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string vmName = "aaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier virtualMachineResourceId = VirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName);
VirtualMachineResource virtualMachine = client.GetVirtualMachineResource(virtualMachineResourceId);

// invoke the operation
VirtualMachineCaptureContent content = new VirtualMachineCaptureContent("aaaaaaaaa", "aaaaaaa", true);
ArmOperation<VirtualMachineCaptureResult> lro = await virtualMachine.CaptureAsync(WaitUntil.Completed, content);
VirtualMachineCaptureResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
