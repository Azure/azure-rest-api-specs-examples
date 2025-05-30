using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Qumulo.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Qumulo;

// Generated from example definition: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_CreateOrUpdate_MinimumSet_Gen.json
// this example is just showing the usage of "FileSystems_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "aaaaaaaaaaaaaaaaaaaaaaaa";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this QumuloFileSystemResource
QumuloFileSystemResourceCollection collection = resourceGroupResource.GetQumuloFileSystemResources();

// invoke the operation
string fileSystemName = "aaaaaaaa";
QumuloFileSystemResourceData data = new QumuloFileSystemResourceData(
    new AzureLocation("aaaaaaaaaaaaaaaaaaaaaaaaa"),
    new MarketplaceDetails("aaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaa", "aa")
    {
        MarketplaceSubscriptionId = "aaaaaaaaaaaaa",
    },
    StorageSku.Standard,
    new QumuloUserDetails
    {
        Email = "viptslwulnpaupfljvnjeq",
    },
    "aaaaaaaaaa",
    "ekceujoecaashtjlsgcymnrdozk",
    9);
ArmOperation<QumuloFileSystemResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, fileSystemName, data);
QumuloFileSystemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
QumuloFileSystemResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
