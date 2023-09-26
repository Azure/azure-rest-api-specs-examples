using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationRecoveryPlans_Reprotect.json
// this example is just showing the usage of "ReplicationRecoveryPlans_Reprotect" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryRecoveryPlanResource created on azure
// for more information of creating SiteRecoveryRecoveryPlanResource, please refer to the document of SiteRecoveryRecoveryPlanResource
string subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
string resourceGroupName = "resourceGroupPS1";
string resourceName = "vault1";
string recoveryPlanName = "RPtest1";
ResourceIdentifier siteRecoveryRecoveryPlanResourceId = SiteRecoveryRecoveryPlanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, recoveryPlanName);
SiteRecoveryRecoveryPlanResource siteRecoveryRecoveryPlan = client.GetSiteRecoveryRecoveryPlanResource(siteRecoveryRecoveryPlanResourceId);

// invoke the operation
ArmOperation<SiteRecoveryRecoveryPlanResource> lro = await siteRecoveryRecoveryPlan.ReprotectAsync(WaitUntil.Completed);
SiteRecoveryRecoveryPlanResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryRecoveryPlanData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
