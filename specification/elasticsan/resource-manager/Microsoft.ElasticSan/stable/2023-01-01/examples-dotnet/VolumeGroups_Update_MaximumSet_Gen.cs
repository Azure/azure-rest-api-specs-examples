using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ElasticSan;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeGroups_Update_MaximumSet_Gen.json
// this example is just showing the usage of "VolumeGroups_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticSanVolumeGroupResource created on azure
// for more information of creating ElasticSanVolumeGroupResource, please refer to the document of ElasticSanVolumeGroupResource
string subscriptionId = "subscriptionid";
string resourceGroupName = "resourcegroupname";
string elasticSanName = "elasticsanname";
string volumeGroupName = "volumegroupname";
ResourceIdentifier elasticSanVolumeGroupResourceId = ElasticSanVolumeGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, elasticSanName, volumeGroupName);
ElasticSanVolumeGroupResource elasticSanVolumeGroup = client.GetElasticSanVolumeGroupResource(elasticSanVolumeGroupResourceId);

// invoke the operation
ElasticSanVolumeGroupPatch patch = new ElasticSanVolumeGroupPatch()
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key7482")] = new UserAssignedIdentity(),
        },
    },
    ProtocolType = StorageTargetType.Iscsi,
    Encryption = ElasticSanEncryptionType.EncryptionAtRestWithPlatformKey,
    EncryptionProperties = new Models.EncryptionProperties()
    {
        KeyVaultProperties = new Models.KeyVaultProperties()
        {
            KeyName = "sftaiernmrzypnrkpakrrawxcbsqzc",
            KeyVersion = "c",
            KeyVaultUri = new Uri("https://microsoft.com/axmblwp"),
        },
        EncryptionUserAssignedIdentity = "im",
    },
    VirtualNetworkRules =
    {
    new ElasticSanVirtualNetworkRule(new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}"))
    {
    Action = ElasticSanVirtualNetworkRuleAction.Allow,
    }
    },
};
ArmOperation<ElasticSanVolumeGroupResource> lro = await elasticSanVolumeGroup.UpdateAsync(WaitUntil.Completed, patch);
ElasticSanVolumeGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticSanVolumeGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
