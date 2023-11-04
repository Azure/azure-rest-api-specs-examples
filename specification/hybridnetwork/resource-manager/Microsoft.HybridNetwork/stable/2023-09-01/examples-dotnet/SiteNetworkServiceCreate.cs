using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/SiteNetworkServiceCreate.json
// this example is just showing the usage of "SiteNetworkServices_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this SiteNetworkServiceResource
SiteNetworkServiceCollection collection = resourceGroupResource.GetSiteNetworkServices();

// invoke the operation
string siteNetworkServiceName = "testSiteNetworkServiceName";
SiteNetworkServiceData data = new SiteNetworkServiceData(new AzureLocation("westUs2"))
{
    Properties = new SiteNetworkServicePropertiesFormat()
    {
        SiteReferenceId = new ResourceIdentifier("/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/sites/testSite"),
        NetworkServiceDesignVersionResourceReference = new OpenDeploymentResourceReference()
        {
            Id = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/TestPublisher/networkServiceDesignGroups/TestNetworkServiceDesignGroupName/networkServiceDesignVersions/1.0.0"),
        },
        DesiredStateConfigurationGroupValueReferences =
        {
        ["MyVM_Configuration"] = new WritableSubResource()
        {
        Id = new ResourceIdentifier("/subscriptions/subid/resourcegroups/contosorg1/providers/microsoft.hybridnetwork/configurationgroupvalues/MyVM_Configuration1"),
        },
        },
    },
    Sku = new HybridNetworkSku(HybridNetworkSkuName.Standard),
};
ArmOperation<SiteNetworkServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, siteNetworkServiceName, data);
SiteNetworkServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteNetworkServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
