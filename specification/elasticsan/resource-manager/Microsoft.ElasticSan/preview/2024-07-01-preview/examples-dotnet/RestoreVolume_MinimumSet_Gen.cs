using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.ElasticSan;

// Generated from example definition: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-07-01-preview/examples/RestoreVolume_MinimumSet_Gen.json
// this example is just showing the usage of "RestoreVolume" operation, for the dependent resources, they will have to be created separately.

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
string volumeName = "volumename-1741526907";
ResourceIdentifier elasticSanVolumeResourceId = ElasticSanVolumeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, elasticSanName, volumeGroupName, volumeName);
ElasticSanVolumeResource elasticSanVolume = client.GetElasticSanVolumeResource(elasticSanVolumeResourceId);

// invoke the operation
ArmOperation<ElasticSanVolumeResource> lro = await elasticSanVolume.RestoreVolumeAsync(WaitUntil.Completed);
ElasticSanVolumeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticSanVolumeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
