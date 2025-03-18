using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevCenter.Models;
using Azure.ResourceManager.DevCenter;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Catalogs_CreateGitHub.json
// this example is just showing the usage of "Catalogs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevCenterResource created on azure
// for more information of creating DevCenterResource, please refer to the document of DevCenterResource
string subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
ResourceIdentifier devCenterResourceId = DevCenterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName);
DevCenterResource devCenter = client.GetDevCenterResource(devCenterResourceId);

// get the collection of this DevCenterCatalogResource
DevCenterCatalogCollection collection = devCenter.GetDevCenterCatalogs();

// invoke the operation
string catalogName = "CentralCatalog";
DevCenterCatalogData data = new DevCenterCatalogData
{
    GitHub = new DevCenterGitCatalog
    {
        Uri = new Uri("https://github.com/Contoso/centralrepo-fake.git"),
        Branch = "main",
        SecretIdentifier = "https://contosokv.vault.azure.net/secrets/CentralRepoPat",
        Path = "/templates",
    },
};
ArmOperation<DevCenterCatalogResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, catalogName, data);
DevCenterCatalogResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevCenterCatalogData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
