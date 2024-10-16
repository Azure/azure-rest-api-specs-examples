using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerServiceFleet.Models;
using Azure.ResourceManager.ContainerServiceFleet;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2024-05-02-preview/examples/AutoUpgradeProfiles_Delete.json
// this example is just showing the usage of "AutoUpgradeProfiles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutoUpgradeProfileResource created on azure
// for more information of creating AutoUpgradeProfileResource, please refer to the document of AutoUpgradeProfileResource
string subscriptionId = "subid1";
string resourceGroupName = "rg1";
string fleetName = "fleet1";
string autoUpgradeProfileName = "autoupgradeprofile1";
ResourceIdentifier autoUpgradeProfileResourceId = AutoUpgradeProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, fleetName, autoUpgradeProfileName);
AutoUpgradeProfileResource autoUpgradeProfile = client.GetAutoUpgradeProfileResource(autoUpgradeProfileResourceId);

// invoke the operation
await autoUpgradeProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
