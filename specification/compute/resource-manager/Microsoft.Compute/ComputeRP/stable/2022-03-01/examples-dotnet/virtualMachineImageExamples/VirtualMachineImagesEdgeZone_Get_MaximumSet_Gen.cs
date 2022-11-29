using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MaximumSet_Gen.json
// this example is just showing the usage of "VirtualMachineImagesEdgeZone_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("aaaaaaaaaaaaaaaaaaaaaaaa");
string edgeZone = "aaaaaaaa";
string publisherName = "aaaaaaaaaaaaaaaaaaaaaaa";
string offer = "aaaaaaaaaaaaaaaaaaaaaaaaaaa";
string skus = "aaaaaaaaaa";
string version = "aaaaaaaaaaaaaaaaaaaaaaaaaaa";
VirtualMachineImage result = await subscriptionResource.GetVirtualMachineImagesEdgeZoneAsync(location, edgeZone, publisherName, offer, skus, version);

Console.WriteLine($"Succeeded: {result}");
