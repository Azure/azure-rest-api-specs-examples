using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImages_Get_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineExtensionImages_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this VirtualMachineExtensionImageResource
AzureLocation location = new AzureLocation("aaaaaaaaaaaaa");
string publisherName = "aaaaaaaaaaaaaaaaaaaa";
VirtualMachineExtensionImageCollection collection = subscriptionResource.GetVirtualMachineExtensionImages(location, publisherName);

// invoke the operation
string type = "aaaaaaaaaaaaaaaaaa";
string version = "aaaaaaaaaaaaaa";
bool result = await collection.ExistsAsync(type, version);

Console.WriteLine($"Succeeded: {result}");
