using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ElasticSan;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/ElasticSans_Update_MaximumSet_Gen.json
// this example is just showing the usage of "ElasticSans_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticSanResource created on azure
// for more information of creating ElasticSanResource, please refer to the document of ElasticSanResource
string subscriptionId = "subscriptionid";
string resourceGroupName = "resourcegroupname";
string elasticSanName = "elasticsanname";
ResourceIdentifier elasticSanResourceId = ElasticSanResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, elasticSanName);
ElasticSanResource elasticSan = client.GetElasticSanResource(elasticSanResourceId);

// invoke the operation
ElasticSanPatch patch = new ElasticSanPatch()
{
    Tags =
    {
    ["key1931"] = "yhjwkgmrrwrcoxblgwgzjqusch",
    },
    BaseSizeTiB = 13,
    ExtendedCapacitySizeTiB = 29,
    PublicNetworkAccess = ElasticSanPublicNetworkAccess.Enabled,
};
ArmOperation<ElasticSanResource> lro = await elasticSan.UpdateAsync(WaitUntil.Completed, patch);
ElasticSanResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticSanData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
