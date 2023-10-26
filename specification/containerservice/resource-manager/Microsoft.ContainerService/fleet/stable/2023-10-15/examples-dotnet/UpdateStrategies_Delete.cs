using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerServiceFleet;
using Azure.ResourceManager.ContainerServiceFleet.Models;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/UpdateStrategies_Delete.json
// this example is just showing the usage of "FleetUpdateStrategies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FleetUpdateStrategyResource created on azure
// for more information of creating FleetUpdateStrategyResource, please refer to the document of FleetUpdateStrategyResource
string subscriptionId = "subid1";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
string updateStrategyName = "strategy1";
ResourceIdentifier fleetUpdateStrategyResourceId = FleetUpdateStrategyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName, updateStrategyName);
FleetUpdateStrategyResource fleetUpdateStrategy = client.GetFleetUpdateStrategyResource(fleetUpdateStrategyResourceId);

// invoke the operation
await fleetUpdateStrategy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
