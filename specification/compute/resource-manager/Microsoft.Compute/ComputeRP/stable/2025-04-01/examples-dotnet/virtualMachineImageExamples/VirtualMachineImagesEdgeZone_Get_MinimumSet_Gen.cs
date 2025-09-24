using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MinimumSet_Gen.json
// this example is just showing the usage of "VirtualMachineImagesEdgeZone_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "{subscription-id}";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("aaaaaaaaaaaaaaaaaaaaaaa");
string edgeZone = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
string publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
string offer = "aaaaaaaaaaa";
string skus = "aaaaaaaaaaaaaaaaaa";
string version = "aa";
SubscriptionResourceGetVirtualMachineImagesEdgeZoneOptions options = new SubscriptionResourceGetVirtualMachineImagesEdgeZoneOptions(
    location,
    edgeZone,
    publisherName,
    offer,
    skus,
    version);
VirtualMachineImage result = await subscriptionResource.GetVirtualMachineImagesEdgeZoneAsync(options);

Console.WriteLine($"Succeeded: {result}");
