using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/Channels_Update.json
// this example is just showing the usage of "Channels_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PartnerNamespaceChannelResource created on azure
// for more information of creating PartnerNamespaceChannelResource, please refer to the document of PartnerNamespaceChannelResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string partnerNamespaceName = "examplePartnerNamespaceName1";
string channelName = "exampleChannelName1";
ResourceIdentifier partnerNamespaceChannelResourceId = PartnerNamespaceChannelResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, partnerNamespaceName, channelName);
PartnerNamespaceChannelResource partnerNamespaceChannel = client.GetPartnerNamespaceChannelResource(partnerNamespaceChannelResourceId);

// invoke the operation
PartnerNamespaceChannelPatch patch = new PartnerNamespaceChannelPatch
{
    ExpireOnIfNotActivated = DateTimeOffset.Parse("2022-03-23T23:06:11.785Z"),
};
await partnerNamespaceChannel.UpdateAsync(patch);

Console.WriteLine("Succeeded");
