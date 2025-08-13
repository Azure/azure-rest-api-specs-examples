using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerServiceFleet.Models;
using Azure.ResourceManager.ContainerServiceFleet;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2025-04-01-preview/examples/FleetUpdateStrategies_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "FleetUpdateStrategies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FleetUpdateStrategyResource created on azure
// for more information of creating FleetUpdateStrategyResource, please refer to the document of FleetUpdateStrategyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgfleets";
string fleetName = "fleet1";
string updateStrategyName = "fleet1";
ResourceIdentifier fleetUpdateStrategyResourceId = FleetUpdateStrategyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName, updateStrategyName);
FleetUpdateStrategyResource fleetUpdateStrategy = client.GetFleetUpdateStrategyResource(fleetUpdateStrategyResourceId);

// invoke the operation
string ifMatch = "saqprswlk";
await fleetUpdateStrategy.DeleteAsync(WaitUntil.Completed, ifMatch: ifMatch);

Console.WriteLine("Succeeded");
