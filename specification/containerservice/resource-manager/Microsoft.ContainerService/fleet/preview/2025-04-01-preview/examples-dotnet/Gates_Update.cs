using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerServiceFleet.Models;
using Azure.ResourceManager.ContainerServiceFleet;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2025-04-01-preview/examples/Gates_Update.json
// this example is just showing the usage of "Gates_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceFleetGateResource created on azure
// for more information of creating ContainerServiceFleetGateResource, please refer to the document of ContainerServiceFleetGateResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
string gateName = "12345678-910a-bcde-f000-000000000000";
ResourceIdentifier containerServiceFleetGateResourceId = ContainerServiceFleetGateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName, gateName);
ContainerServiceFleetGateResource containerServiceFleetGate = client.GetContainerServiceFleetGateResource(containerServiceFleetGateResourceId);

// invoke the operation
ContainerServiceFleetGatePatch patch = new ContainerServiceFleetGatePatch
{
    GatePatchState = ContainerServiceFleetGateState.Completed,
};
ArmOperation<ContainerServiceFleetGateResource> lro = await containerServiceFleetGate.UpdateAsync(WaitUntil.Completed, patch);
ContainerServiceFleetGateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerServiceFleetGateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
