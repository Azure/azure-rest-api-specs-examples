using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.ElasticSan;

// Generated from example definition: 2024-07-01-preview/Volumes_PreRestore_MaximumSet_Gen.json
// this example is just showing the usage of "VolumeGroups_PreRestore" operation, for the dependent resources, they will have to be created separately.

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
DiskSnapshotListContent content = new DiskSnapshotListContent(new ResourceIdentifier[] { new ResourceIdentifier("/subscriptions/{subscriptionid}/resourceGroups/{resourcegroupname}/providers/Microsoft.Compute/snapshots/disksnapshot1") });
ArmOperation<ElasticSanPreValidationResult> lro = await elasticSanVolumeGroup.PreRestoreVolumeAsync(WaitUntil.Completed, content);
ElasticSanPreValidationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
