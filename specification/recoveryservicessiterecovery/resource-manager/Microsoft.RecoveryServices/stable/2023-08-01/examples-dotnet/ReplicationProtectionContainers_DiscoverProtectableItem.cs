using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectionContainers_DiscoverProtectableItem.json
// this example is just showing the usage of "ReplicationProtectionContainers_DiscoverProtectableItem" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryProtectionContainerResource created on azure
// for more information of creating SiteRecoveryProtectionContainerResource, please refer to the document of SiteRecoveryProtectionContainerResource
string subscriptionId = "7c943c1b-5122-4097-90c8-861411bdd574";
string resourceGroupName = "MadhaviVRG";
string resourceName = "MadhaviVault";
string fabricName = "V2A-W2K12-660";
string protectionContainerName = "cloud_7328549c-5c37-4459-a3c2-e35f9ef6893c";
ResourceIdentifier siteRecoveryProtectionContainerResourceId = SiteRecoveryProtectionContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName);
SiteRecoveryProtectionContainerResource siteRecoveryProtectionContainer = client.GetSiteRecoveryProtectionContainerResource(siteRecoveryProtectionContainerResourceId);

// invoke the operation
DiscoverProtectableItemContent content = new DiscoverProtectableItemContent()
{
    Properties = new DiscoverProtectableItemProperties()
    {
        FriendlyName = "Test",
        IPAddress = IPAddress.Parse("10.150.2.3"),
        OSType = "Windows",
    },
};
ArmOperation<SiteRecoveryProtectionContainerResource> lro = await siteRecoveryProtectionContainer.DiscoverProtectableItemAsync(WaitUntil.Completed, content);
SiteRecoveryProtectionContainerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryProtectionContainerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
