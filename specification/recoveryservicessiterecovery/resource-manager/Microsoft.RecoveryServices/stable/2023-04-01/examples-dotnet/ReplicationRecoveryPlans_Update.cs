using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ReplicationRecoveryPlans_Update.json
// this example is just showing the usage of "ReplicationRecoveryPlans_Update" operation, for the dependent resources, they will have to be created separately.

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
SiteRecoveryRecoveryPlanPatch patch = new SiteRecoveryRecoveryPlanPatch()
{
    UpdateRecoveryPlanContentGroups =
    {
    new SiteRecoveryPlanGroup(RecoveryPlanGroupType.Shutdown)
    {
    ReplicationProtectedItems =
    {
    },
    StartGroupActions =
    {
    },
    EndGroupActions =
    {
    },
    },new SiteRecoveryPlanGroup(RecoveryPlanGroupType.Failover)
    {
    ReplicationProtectedItems =
    {
    },
    StartGroupActions =
    {
    },
    EndGroupActions =
    {
    },
    },new SiteRecoveryPlanGroup(RecoveryPlanGroupType.Boot)
    {
    ReplicationProtectedItems =
    {
    new RecoveryPlanProtectedItem()
    {
    Id = new ResourceIdentifier("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
    VirtualMachineId = "f8491e4f-817a-40dd-a90c-af773978c75b",
    }
    },
    StartGroupActions =
    {
    },
    EndGroupActions =
    {
    },
    },new SiteRecoveryPlanGroup(RecoveryPlanGroupType.Boot)
    {
    ReplicationProtectedItems =
    {
    new RecoveryPlanProtectedItem()
    {
    Id = new ResourceIdentifier("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/c0c14913-3d7a-48ea-9531-cc99e0e686e6"),
    VirtualMachineId = "c0c14913-3d7a-48ea-9531-cc99e0e686e6",
    }
    },
    StartGroupActions =
    {
    },
    EndGroupActions =
    {
    },
    }
    },
};
ArmOperation<SiteRecoveryRecoveryPlanResource> lro = await siteRecoveryRecoveryPlan.UpdateAsync(WaitUntil.Completed, patch);
SiteRecoveryRecoveryPlanResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryRecoveryPlanData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
