using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationRecoveryPlans_Create.json
// this example is just showing the usage of "ReplicationRecoveryPlans_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
string resourceGroupName = "resourceGroupPS1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SiteRecoveryRecoveryPlanResource
string resourceName = "vault1";
SiteRecoveryRecoveryPlanCollection collection = resourceGroupResource.GetSiteRecoveryRecoveryPlans(resourceName);

// invoke the operation
string recoveryPlanName = "RPtest1";
SiteRecoveryRecoveryPlanCreateOrUpdateContent content = new SiteRecoveryRecoveryPlanCreateOrUpdateContent(new SiteRecoveryCreateRecoveryPlanProperties(new ResourceIdentifier("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"), new ResourceIdentifier("Microsoft Azure"), new SiteRecoveryPlanGroup[]
{
new SiteRecoveryPlanGroup(RecoveryPlanGroupType.Boot)
{
ReplicationProtectedItems = {new RecoveryPlanProtectedItem
{
Id = new ResourceIdentifier("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1/replicationProtectionContainers/cloud_6d224fc6-f326-5d35-96de-fbf51efb3179/replicationProtectedItems/f8491e4f-817a-40dd-a90c-af773978c75b"),
VirtualMachineId = "f8491e4f-817a-40dd-a90c-af773978c75b",
}},
StartGroupActions = {},
EndGroupActions = {},
}
})
{
    FailoverDeploymentModel = FailoverDeploymentModel.ResourceManager,
});
ArmOperation<SiteRecoveryRecoveryPlanResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, recoveryPlanName, content);
SiteRecoveryRecoveryPlanResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryRecoveryPlanData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
