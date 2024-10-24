using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerServiceFleet.Models;
using Azure.ResourceManager.ContainerServiceFleet;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/examples/UpdateRuns_Delete.json
// this example is just showing the usage of "UpdateRuns_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceFleetUpdateRunResource created on azure
// for more information of creating ContainerServiceFleetUpdateRunResource, please refer to the document of ContainerServiceFleetUpdateRunResource
string subscriptionId = "subid1";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
string updateRunName = "run1";
ResourceIdentifier containerServiceFleetUpdateRunResourceId = ContainerServiceFleetUpdateRunResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName, updateRunName);
ContainerServiceFleetUpdateRunResource containerServiceFleetUpdateRun = client.GetContainerServiceFleetUpdateRunResource(containerServiceFleetUpdateRunResourceId);

// invoke the operation
await containerServiceFleetUpdateRun.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
