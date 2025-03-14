using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationvCenters_Create.json
// this example is just showing the usage of "ReplicationvCenters_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryFabricResource created on azure
// for more information of creating SiteRecoveryFabricResource, please refer to the document of SiteRecoveryFabricResource
string subscriptionId = "7c943c1b-5122-4097-90c8-861411bdd574";
string resourceGroupName = "MadhaviVRG";
string resourceName = "MadhaviVault";
string fabricName = "MadhaviFabric";
ResourceIdentifier siteRecoveryFabricResourceId = SiteRecoveryFabricResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName);
SiteRecoveryFabricResource siteRecoveryFabric = client.GetSiteRecoveryFabricResource(siteRecoveryFabricResourceId);

// get the collection of this SiteRecoveryVCenterResource
SiteRecoveryVCenterCollection collection = siteRecoveryFabric.GetSiteRecoveryVCenters();

// invoke the operation
string vCenterName = "esx-78";
SiteRecoveryVCenterCreateOrUpdateContent content = new SiteRecoveryVCenterCreateOrUpdateContent
{
    Properties = new SiteRecoveryAddVCenterProperties
    {
        FriendlyName = "esx-78",
        IPAddress = IPAddress.Parse("inmtest78"),
        ProcessServerId = Guid.Parse("5A720CAB-39CB-F445-BD1662B0B33164B5"),
        Port = "443",
        RunAsAccountId = "2",
    },
};
ArmOperation<SiteRecoveryVCenterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vCenterName, content);
SiteRecoveryVCenterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryVCenterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
