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

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ReplicationvCenters_Delete.json
// this example is just showing the usage of "ReplicationvCenters_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryVCenterResource created on azure
// for more information of creating SiteRecoveryVCenterResource, please refer to the document of SiteRecoveryVCenterResource
string subscriptionId = "7c943c1b-5122-4097-90c8-861411bdd574";
string resourceGroupName = "MadhaviVRG";
string resourceName = "MadhaviVault";
string fabricName = "MadhaviFabric";
string vCenterName = "esx-78";
ResourceIdentifier siteRecoveryVCenterResourceId = SiteRecoveryVCenterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, vCenterName);
SiteRecoveryVCenterResource siteRecoveryVCenter = client.GetSiteRecoveryVCenterResource(siteRecoveryVCenterResourceId);

// invoke the operation
await siteRecoveryVCenter.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
