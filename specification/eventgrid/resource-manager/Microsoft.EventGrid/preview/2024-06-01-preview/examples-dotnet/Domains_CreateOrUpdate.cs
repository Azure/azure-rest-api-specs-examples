using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/Domains_CreateOrUpdate.json
// this example is just showing the usage of "Domains_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this EventGridDomainResource
EventGridDomainCollection collection = resourceGroupResource.GetEventGridDomains();

// invoke the operation
string domainName = "exampledomain1";
EventGridDomainData data = new EventGridDomainData(new AzureLocation("westus2"))
{
    PublicNetworkAccess = EventGridPublicNetworkAccess.Enabled,
    InboundIPRules =
    {
    new EventGridInboundIPRule()
    {
    IPMask = "12.18.30.15",
    Action = EventGridIPActionType.Allow,
    },new EventGridInboundIPRule()
    {
    IPMask = "12.18.176.1",
    Action = EventGridIPActionType.Allow,
    }
    },
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
};
ArmOperation<EventGridDomainResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, domainName, data);
EventGridDomainResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventGridDomainData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
