using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerServiceFleet.Models;
using Azure.ResourceManager.ContainerServiceFleet;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/stable/2025-03-01/examples/AutoUpgradeProfiles_CreateOrUpdate.json
// this example is just showing the usage of "AutoUpgradeProfiles_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this AutoUpgradeProfileResource
AutoUpgradeProfileCollection collection = containerServiceFleet.GetAutoUpgradeProfiles();

// invoke the operation
string autoUpgradeProfileName = "autoupgradeprofile1";
AutoUpgradeProfileData data = new AutoUpgradeProfileData
{
    Channel = ContainerServiceFleetUpgradeChannel.Stable,
};
ArmOperation<AutoUpgradeProfileResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, autoUpgradeProfileName, data);
AutoUpgradeProfileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutoUpgradeProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
