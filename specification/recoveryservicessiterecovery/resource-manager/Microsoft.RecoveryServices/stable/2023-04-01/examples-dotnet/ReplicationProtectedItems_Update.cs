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

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ReplicationProtectedItems_Update.json
// this example is just showing the usage of "ReplicationProtectedItems_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReplicationProtectedItemResource created on azure
// for more information of creating ReplicationProtectedItemResource, please refer to the document of ReplicationProtectedItemResource
string subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
string resourceGroupName = "resourceGroupPS1";
string resourceName = "vault1";
string fabricName = "cloud1";
string protectionContainerName = "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179";
string replicatedProtectedItemName = "f8491e4f-817a-40dd-a90c-af773978c75b";
ResourceIdentifier replicationProtectedItemResourceId = ReplicationProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, fabricName, protectionContainerName, replicatedProtectedItemName);
ReplicationProtectedItemResource replicationProtectedItem = client.GetReplicationProtectedItemResource(replicationProtectedItemResourceId);

// invoke the operation
ReplicationProtectedItemPatch patch = new ReplicationProtectedItemPatch()
{
    Properties = new UpdateReplicationProtectedItemProperties()
    {
        RecoveryAzureVmName = "vm1",
        RecoveryAzureVmSize = "Basic_A0",
        SelectedRecoveryAzureNetworkId = new ResourceIdentifier("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai"),
        VmNics =
        {
        new VmNicContentDetails()
        {
        NicId = "TWljcm9zb2Z0OkY4NDkxRTRGLTgxN0EtNDBERC1BOTBDLUFGNzczOTc4Qzc1Qlw3NjAwMzMxRS03NDk4LTQ0QTQtQjdDNy0xQjY1NkJDREQ1MkQ=",
        IPConfigs =
        {
        new HyperVFailoverIPConfigDetails()
        {
        IPConfigName = "ipconfig1",
        IsPrimary = true,
        RecoverySubnetName = "subnet1",
        RecoveryStaticIPAddress = IPAddress.Parse("10.0.2.46"),
        }
        },
        SelectionType = "SelectedByUser",
        }
        },
        LicenseType = SiteRecoveryLicenseType.WindowsServer,
        ProviderSpecificDetails = new HyperVReplicaAzureUpdateReplicationProtectedItemContent(),
    },
};
ArmOperation<ReplicationProtectedItemResource> lro = await replicationProtectedItem.UpdateAsync(WaitUntil.Completed, patch);
ReplicationProtectedItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ReplicationProtectedItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
