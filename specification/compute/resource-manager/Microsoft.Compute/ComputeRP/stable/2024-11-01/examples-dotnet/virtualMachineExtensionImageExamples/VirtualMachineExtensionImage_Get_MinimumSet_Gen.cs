using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_Get_MinimumSet_Gen.json
// this example is just showing the usage of "VirtualMachineExtensionImages_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineExtensionImageResource created on azure
// for more information of creating VirtualMachineExtensionImageResource, please refer to the document of VirtualMachineExtensionImageResource
string subscriptionId = "{subscription-id}";
AzureLocation location = new AzureLocation("aaaaaaaaaaaaaa");
string publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
string type = "aa";
string version = "aaa";
ResourceIdentifier virtualMachineExtensionImageResourceId = VirtualMachineExtensionImageResource.CreateResourceIdentifier(subscriptionId, location, publisherName, type, version);
VirtualMachineExtensionImageResource virtualMachineExtensionImage = client.GetVirtualMachineExtensionImageResource(virtualMachineExtensionImageResourceId);

// invoke the operation
VirtualMachineExtensionImageResource result = await virtualMachineExtensionImage.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineExtensionImageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
