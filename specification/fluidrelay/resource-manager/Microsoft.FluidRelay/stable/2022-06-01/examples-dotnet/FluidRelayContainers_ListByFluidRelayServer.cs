using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.FluidRelay;

// Generated from example definition: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayContainers_ListByFluidRelayServer.json
// this example is just showing the usage of "FluidRelayContainers_ListByFluidRelayServers" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this FluidRelayContainerResource
FluidRelayContainerCollection collection = fluidRelayServer.GetFluidRelayContainers();

// invoke the operation and iterate over the result
await foreach (FluidRelayContainerResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    FluidRelayContainerData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
