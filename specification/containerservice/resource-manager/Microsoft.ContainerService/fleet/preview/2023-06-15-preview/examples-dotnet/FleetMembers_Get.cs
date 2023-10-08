using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerServiceFleet;
using Azure.ResourceManager.ContainerServiceFleet.Models;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-06-15-preview/examples/FleetMembers_Get.json
// this example is just showing the usage of "FleetMembers_Get" operation, for the dependent resources, they will have to be created separately.

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
ContainerServiceFleetMemberResource result = await containerServiceFleetMember.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerServiceFleetMemberData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
