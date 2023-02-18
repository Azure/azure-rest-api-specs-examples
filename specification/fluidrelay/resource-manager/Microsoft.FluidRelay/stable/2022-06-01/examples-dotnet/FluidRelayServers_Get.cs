using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.FluidRelay;
using Azure.ResourceManager.FluidRelay.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_Get.json
// this example is just showing the usage of "FluidRelayServers_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "xxxx-xxxx-xxxx-xxxx";
string resourceGroup = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroup);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this FluidRelayServerResource
FluidRelayServerCollection collection = resourceGroupResource.GetFluidRelayServers();

// invoke the operation
string fluidRelayServerName = "myFluidRelayServer";
bool result = await collection.ExistsAsync(fluidRelayServerName);

Console.WriteLine($"Succeeded: {result}");
