using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerServiceFleet;
using Azure.ResourceManager.ContainerServiceFleet.Models;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2023-10-15/examples/UpdateRuns_Get.json
// this example is just showing the usage of "UpdateRuns_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceFleetResource created on azure
// for more information of creating ContainerServiceFleetResource, please refer to the document of ContainerServiceFleetResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
ResourceIdentifier containerServiceFleetResourceId = ContainerServiceFleetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName);
ContainerServiceFleetResource containerServiceFleet = client.GetContainerServiceFleetResource(containerServiceFleetResourceId);

// get the collection of this ContainerServiceFleetUpdateRunResource
ContainerServiceFleetUpdateRunCollection collection = containerServiceFleet.GetContainerServiceFleetUpdateRuns();

// invoke the operation
string updateRunName = "run1";
NullableResponse<ContainerServiceFleetUpdateRunResource> response = await collection.GetIfExistsAsync(updateRunName);
ContainerServiceFleetUpdateRunResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ContainerServiceFleetUpdateRunData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
