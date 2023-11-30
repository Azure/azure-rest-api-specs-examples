using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.FluidRelay;
using Azure.ResourceManager.FluidRelay.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_Get.json
// this example is just showing the usage of "FluidRelayServers_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FluidRelayServerResource created on azure
// for more information of creating FluidRelayServerResource, please refer to the document of FluidRelayServerResource
string subscriptionId = "xxxx-xxxx-xxxx-xxxx";
string resourceGroup = "myResourceGroup";
string fluidRelayServerName = "myFluidRelayServer";
ResourceIdentifier fluidRelayServerResourceId = FluidRelayServerResource.CreateResourceIdentifier(subscriptionId, resourceGroup, fluidRelayServerName);
FluidRelayServerResource fluidRelayServer = client.GetFluidRelayServerResource(fluidRelayServerResourceId);

// invoke the operation
FluidRelayServerResource result = await fluidRelayServer.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FluidRelayServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
