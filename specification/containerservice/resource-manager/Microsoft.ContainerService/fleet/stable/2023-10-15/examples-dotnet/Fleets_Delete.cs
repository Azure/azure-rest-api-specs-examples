using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerServiceFleet;
using Azure.ResourceManager.ContainerServiceFleet.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/Fleets_Delete.json
// this example is just showing the usage of "Fleets_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceFleetResource created on azure
// for more information of creating ContainerServiceFleetResource, please refer to the document of ContainerServiceFleetResource
string subscriptionId = "subid1";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
ResourceIdentifier containerServiceFleetResourceId = ContainerServiceFleetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName);
ContainerServiceFleetResource containerServiceFleet = client.GetContainerServiceFleetResource(containerServiceFleetResourceId);

// invoke the operation
await containerServiceFleet.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
