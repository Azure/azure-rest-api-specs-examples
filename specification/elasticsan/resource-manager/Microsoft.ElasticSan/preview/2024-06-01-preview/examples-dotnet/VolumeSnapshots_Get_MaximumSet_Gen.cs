using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.ElasticSan;

// Generated from example definition: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2024-06-01-preview/examples/VolumeSnapshots_Get_MaximumSet_Gen.json
// this example is just showing the usage of "VolumeSnapshots_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticSanSnapshotResource created on azure
// for more information of creating ElasticSanSnapshotResource, please refer to the document of ElasticSanSnapshotResource
string subscriptionId = "subscriptionid";
string resourceGroupName = "resourcegroupname";
string elasticSanName = "elasticsanname";
string volumeGroupName = "volumegroupname";
string snapshotName = "snapshotname";
ResourceIdentifier elasticSanSnapshotResourceId = ElasticSanSnapshotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, elasticSanName, volumeGroupName, snapshotName);
ElasticSanSnapshotResource elasticSanSnapshot = client.GetElasticSanSnapshotResource(elasticSanSnapshotResourceId);

// invoke the operation
ElasticSanSnapshotResource result = await elasticSanSnapshot.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticSanSnapshotData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
