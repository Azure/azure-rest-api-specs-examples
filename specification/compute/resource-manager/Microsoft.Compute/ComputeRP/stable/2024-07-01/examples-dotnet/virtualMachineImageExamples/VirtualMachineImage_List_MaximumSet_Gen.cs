using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/virtualMachineImageExamples/VirtualMachineImage_List_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineImages_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
AzureLocation location = new AzureLocation("aaaaaaaaaaaaaaa");
string publisherName = "aaaaaa";
string offer = "aaaaaaaaaaaaaaaa";
string skus = "aaaaaaaaaaaaaaaaaaaaaaa";
SubscriptionResourceGetVirtualMachineImagesOptions options = new SubscriptionResourceGetVirtualMachineImagesOptions(location, publisherName, offer, skus) { Expand = "aaaaaaaaaaaaaaaaaaaaaaaa", Top = 18, Orderby = "aa" };
await foreach (VirtualMachineImageBase item in subscriptionResource.GetVirtualMachineImagesAsync(options))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
