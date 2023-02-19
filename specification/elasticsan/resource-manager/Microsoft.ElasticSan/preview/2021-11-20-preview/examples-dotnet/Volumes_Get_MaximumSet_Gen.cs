using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ElasticSan;
using Azure.ResourceManager.ElasticSan.Models;

// Generated from example definition: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/Volumes_Get_MaximumSet_Gen.json
// this example is just showing the usage of "Volumes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticSanVolumeGroupResource created on azure
// for more information of creating ElasticSanVolumeGroupResource, please refer to the document of ElasticSanVolumeGroupResource
string subscriptionId = "aaaaaaaaaaaaaaaaaa";
string resourceGroupName = "rgelasticsan";
string elasticSanName = "ti7q-k952-1qB3J_5";
string volumeGroupName = "u_5I_1j4t3";
ResourceIdentifier elasticSanVolumeGroupResourceId = ElasticSanVolumeGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, elasticSanName, volumeGroupName);
ElasticSanVolumeGroupResource elasticSanVolumeGroup = client.GetElasticSanVolumeGroupResource(elasticSanVolumeGroupResourceId);

// get the collection of this ElasticSanVolumeResource
ElasticSanVolumeCollection collection = elasticSanVolumeGroup.GetElasticSanVolumes();

// invoke the operation
string volumeName = "9132y";
bool result = await collection.ExistsAsync(volumeName);

Console.WriteLine($"Succeeded: {result}");
