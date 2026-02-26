using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.ElasticSan;

// Generated from example definition: 2024-07-01-preview/Volumes_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "Volume_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticSanVolumeResource created on azure
// for more information of creating ElasticSanVolumeResource, please refer to the document of ElasticSanVolumeResource
string subscriptionId = "subscriptionid";
string resourceGroupName = "resourcegroupname";
string elasticSanName = "elasticsanname";
string volumeGroupName = "volumegroupname";
string volumeName = "volumename";
ResourceIdentifier elasticSanVolumeResourceId = ElasticSanVolumeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, elasticSanName, volumeGroupName, volumeName);
ElasticSanVolumeResource elasticSanVolume = client.GetElasticSanVolumeResource(elasticSanVolumeResourceId);

// invoke the operation
ElasticSanDeleteType? deleteType = ElasticSanDeleteType.Permanent;
await elasticSanVolume.DeleteAsync(WaitUntil.Completed, deleteType: deleteType);

Console.WriteLine("Succeeded");
