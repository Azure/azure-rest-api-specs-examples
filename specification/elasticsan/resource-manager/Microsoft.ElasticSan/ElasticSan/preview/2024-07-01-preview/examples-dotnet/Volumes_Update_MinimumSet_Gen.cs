using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.ElasticSan;

// Generated from example definition: 2024-07-01-preview/Volumes_Update_MinimumSet_Gen.json
// this example is just showing the usage of "Volume_Update" operation, for the dependent resources, they will have to be created separately.

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
ElasticSanVolumePatch patch = new ElasticSanVolumePatch();
ArmOperation<ElasticSanVolumeResource> lro = await elasticSanVolume.UpdateAsync(WaitUntil.Completed, patch);
ElasticSanVolumeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticSanVolumeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
