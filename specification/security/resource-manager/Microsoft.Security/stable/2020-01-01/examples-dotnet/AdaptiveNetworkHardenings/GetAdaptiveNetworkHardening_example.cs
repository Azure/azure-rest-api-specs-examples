using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/GetAdaptiveNetworkHardening_example.json
// this example is just showing the usage of "AdaptiveNetworkHardenings_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AdaptiveNetworkHardeningResource
string resourceNamespace = "Microsoft.Compute";
string resourceType = "virtualMachines";
string resourceName = "vm1";
AdaptiveNetworkHardeningCollection collection = resourceGroupResource.GetAdaptiveNetworkHardenings(resourceNamespace, resourceType, resourceName);

// invoke the operation
string adaptiveNetworkHardeningResourceName = "default";
NullableResponse<AdaptiveNetworkHardeningResource> response = await collection.GetIfExistsAsync(adaptiveNetworkHardeningResourceName);
AdaptiveNetworkHardeningResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AdaptiveNetworkHardeningData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
