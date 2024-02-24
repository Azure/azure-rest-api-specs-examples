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

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationProtectionContainers_SwitchProtection.json
// this example is just showing the usage of "ReplicationProtectionContainers_SwitchProtection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryProtectionContainerResource created on azure
// for more information of creating SiteRecoveryProtectionContainerResource, please refer to the document of SiteRecoveryProtectionContainerResource
string subscriptionId = "42195872-7e70-4f8a-837f-84b28ecbb78b";
string resourceGroupName = "priyanprg";
string resourceName = "priyanponeboxvault";
string fabricName = "CentralUSCanSite";
string protectionContainerName = "CentralUSCancloud";
ResourceIdentifier siteRecoveryProtectionContainerResourceId = SiteRecoveryProtectionContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName);
SiteRecoveryProtectionContainerResource siteRecoveryProtectionContainer = client.GetSiteRecoveryProtectionContainerResource(siteRecoveryProtectionContainerResourceId);

// invoke the operation
SwitchProtectionContent content = new SwitchProtectionContent()
{
    Properties = new SwitchProtectionProperties()
    {
        ReplicationProtectedItemName = "a2aSwapOsVm",
        ProviderSpecificDetails = new A2ASwitchProtectionContent(),
    },
};
ArmOperation<SiteRecoveryProtectionContainerResource> lro = await siteRecoveryProtectionContainer.SwitchProtectionAsync(WaitUntil.Completed, content);
SiteRecoveryProtectionContainerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryProtectionContainerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
