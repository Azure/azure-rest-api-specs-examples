using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/Domains_Update.json
// this example is just showing the usage of "Domains_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventGridDomainResource created on azure
// for more information of creating EventGridDomainResource, please refer to the document of EventGridDomainResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string domainName = "exampledomain1";
ResourceIdentifier eventGridDomainResourceId = EventGridDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, domainName);
EventGridDomainResource eventGridDomain = client.GetEventGridDomainResource(eventGridDomainResourceId);

// invoke the operation
EventGridDomainPatch patch = new EventGridDomainPatch()
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
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
};
await eventGridDomain.UpdateAsync(WaitUntil.Completed, patch);

Console.WriteLine($"Succeeded");
