using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/SiteCreate.json
// this example is just showing the usage of "Sites_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SiteResource
SiteCollection collection = resourceGroupResource.GetSites();

// invoke the operation
string siteName = "testSite";
SiteData data = new SiteData(new AzureLocation("westUs2"))
{
    Properties = new SitePropertiesFormat()
    {
        Nfvis =
        {
        new AzureCoreNfviDetails()
        {
        Location = new AzureLocation("westUs2"),
        Name = "nfvi1",
        },new AzureArcK8SClusterNfviDetails()
        {
        CustomLocationReferenceId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation1"),
        Name = "nfvi2",
        },new AzureOperatorNexusClusterNfviDetails()
        {
        CustomLocationReferenceId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation2"),
        Name = "nfvi3",
        }
        },
    },
};
ArmOperation<SiteResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, siteName, data);
SiteResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
