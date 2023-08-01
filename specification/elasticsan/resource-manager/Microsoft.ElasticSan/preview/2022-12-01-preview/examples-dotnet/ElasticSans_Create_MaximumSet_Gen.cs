using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ElasticSan;
using Azure.ResourceManager.ElasticSan.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/ElasticSans_Create_MaximumSet_Gen.json
// this example is just showing the usage of "ElasticSans_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subscriptionid";
string resourceGroupName = "resourcegroupname";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ElasticSanResource
ElasticSanCollection collection = resourceGroupResource.GetElasticSans();

// invoke the operation
string elasticSanName = "elasticsanname";
ElasticSanData data = new ElasticSanData(new AzureLocation("South Central US"), new ElasticSanSku(ElasticSanSkuName.PremiumLrs)
{
    Tier = ElasticSanSkuTier.Premium,
}, 5, 25)
{
    AvailabilityZones =
    {
    "1"
    },
    Tags =
    {
    ["key9316"] = "ihndtieqibtob",
    },
};
ArmOperation<ElasticSanResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, elasticSanName, data);
ElasticSanResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ElasticSanData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
