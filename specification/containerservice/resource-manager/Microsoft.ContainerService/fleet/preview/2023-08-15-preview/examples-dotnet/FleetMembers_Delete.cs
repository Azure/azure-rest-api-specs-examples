using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerServiceFleet;
using Azure.ResourceManager.ContainerServiceFleet.Models;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-08-15-preview/examples/FleetMembers_Delete.json
// this example is just showing the usage of "FleetMembers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceFleetMemberResource created on azure
// for more information of creating ContainerServiceFleetMemberResource, please refer to the document of ContainerServiceFleetMemberResource
string subscriptionId = "subid1";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
string fleetMemberName = "member-1";
ResourceIdentifier containerServiceFleetMemberResourceId = ContainerServiceFleetMemberResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName, fleetMemberName);
ContainerServiceFleetMemberResource containerServiceFleetMember = client.GetContainerServiceFleetMemberResource(containerServiceFleetMemberResourceId);

// invoke the operation
await containerServiceFleetMember.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
