using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.ElasticSan;

// Generated from example definition: 2024-07-01-preview/Volumes_ListByVolumeGroup_MinimumSet_Gen.json
// this example is just showing the usage of "Volume_ListByVolumeGroup" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ElasticSanVolumeResource
ElasticSanVolumeCollection collection = elasticSanVolumeGroup.GetElasticSanVolumes();

// invoke the operation and iterate over the result
await foreach (ElasticSanVolumeResource item in collection.GetAllAsync(cancellationToken: default))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ElasticSanVolumeData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
