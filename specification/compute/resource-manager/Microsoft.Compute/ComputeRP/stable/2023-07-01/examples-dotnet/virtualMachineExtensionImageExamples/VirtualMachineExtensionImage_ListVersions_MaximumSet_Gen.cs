using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_ListVersions_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineExtensionImages_ListVersions" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this VirtualMachineExtensionImageResource
AzureLocation location = new AzureLocation("aaaaaaaaaaaaaaaaaaaaaaaaaa");
string publisherName = "aaaaaaaaaaaaaaaaaaaa";
VirtualMachineExtensionImageCollection collection = subscriptionResource.GetVirtualMachineExtensionImages(location, publisherName);

// invoke the operation and iterate over the result
string type = "aaaaaaaaaaaaaaaaaa";
string filter = "aaaaaaaaaaaaaaaaaaaaaaaaa";
int? top = 22;
string orderby = "a";
await foreach (VirtualMachineExtensionImageResource item in collection.GetAllAsync(type, filter: filter, top: top, orderby: orderby))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    VirtualMachineExtensionImageData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
