using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerServiceFleet.Models;
using Azure.ResourceManager.ContainerServiceFleet;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2025-04-01-preview/examples/FleetMembers_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "FleetMembers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceFleetMemberResource created on azure
// for more information of creating ContainerServiceFleetMemberResource, please refer to the document of ContainerServiceFleetMemberResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgfleets";
string fleetName = "fleet1";
string fleetMemberName = "fleet1";
ResourceIdentifier containerServiceFleetMemberResourceId = ContainerServiceFleetMemberResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName, fleetMemberName);
ContainerServiceFleetMemberResource containerServiceFleetMember = client.GetContainerServiceFleetMemberResource(containerServiceFleetMemberResourceId);

// invoke the operation
string ifMatch = "klroqfozx";
await containerServiceFleetMember.DeleteAsync(WaitUntil.Completed, ifMatch: ifMatch);

Console.WriteLine("Succeeded");
