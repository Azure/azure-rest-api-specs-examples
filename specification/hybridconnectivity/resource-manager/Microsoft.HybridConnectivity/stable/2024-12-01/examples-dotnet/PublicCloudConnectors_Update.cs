using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridConnectivity.Models;
using Azure.ResourceManager.HybridConnectivity;

// Generated from example definition: 2024-12-01/PublicCloudConnectors_Update.json
// this example is just showing the usage of "PublicCloudConnector_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PublicCloudConnectorResource created on azure
// for more information of creating PublicCloudConnectorResource, please refer to the document of PublicCloudConnectorResource
string subscriptionId = "5ACC4579-DB34-4C2F-8F8C-25061168F342";
string resourceGroupName = "rgpublicCloud";
string publicCloudConnector = "svtirlbyqpepbzyessjenlueeznhg";
ResourceIdentifier publicCloudConnectorResourceId = PublicCloudConnectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publicCloudConnector);
PublicCloudConnectorResource publicCloudConnector0 = client.GetPublicCloudConnectorResource(publicCloudConnectorResourceId);

// invoke the operation
PublicCloudConnectorPatch patch = new PublicCloudConnectorPatch
{
    AwsCloudExcludedAccounts = { "zrbtd" },
    Tags = { },
};
PublicCloudConnectorResource result = await publicCloudConnector0.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PublicCloudConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
