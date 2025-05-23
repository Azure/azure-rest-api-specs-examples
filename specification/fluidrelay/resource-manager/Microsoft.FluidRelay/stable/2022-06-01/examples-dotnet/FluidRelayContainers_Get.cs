using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FluidRelay;

// Generated from example definition: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayContainers_Get.json
// this example is just showing the usage of "FluidRelayContainers_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FluidRelayContainerResource created on azure
// for more information of creating FluidRelayContainerResource, please refer to the document of FluidRelayContainerResource
string subscriptionId = "xxxx-xxxx-xxxx-xxxx";
string resourceGroup = "myResourceGroup";
string fluidRelayServerName = "myFluidRelayServer";
string fluidRelayContainerName = "myFluidRelayContainer";
ResourceIdentifier fluidRelayContainerResourceId = FluidRelayContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroup, fluidRelayServerName, fluidRelayContainerName);
FluidRelayContainerResource fluidRelayContainer = client.GetFluidRelayContainerResource(fluidRelayContainerResourceId);

// invoke the operation
FluidRelayContainerResource result = await fluidRelayContainer.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FluidRelayContainerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
