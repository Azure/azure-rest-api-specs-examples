using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.WebPubSub.Models;
using Azure.ResourceManager.WebPubSub;

// Generated from example definition: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_CreateOrUpdate.json
// this example is just showing the usage of "WebPubSub_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this WebPubSubResource
WebPubSubCollection collection = resourceGroupResource.GetWebPubSubs();

// invoke the operation
string resourceName = "myWebPubSubService";
WebPubSubData data = new WebPubSubData(new AzureLocation("eastus"))
{
    Sku = new BillingInfoSku("Standard_S1")
    {
        Tier = WebPubSubSkuTier.Standard,
        Capacity = 1,
    },
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    IsClientCertEnabled = false,
    LiveTraceConfiguration = new LiveTraceConfiguration()
    {
        Categories =
        {
        new LiveTraceCategory()
        {
        Name = "ConnectivityLogs",
        }
        },
    },
    NetworkAcls = new WebPubSubNetworkAcls()
    {
        DefaultAction = AclAction.Deny,
        PublicNetwork = new PublicNetworkAcls()
        {
            Allow =
            {
            WebPubSubRequestType.ClientConnection
            },
        },
        PrivateEndpoints =
        {
        new PrivateEndpointAcl("mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e")
        {
        Allow =
        {
        WebPubSubRequestType.ServerConnection
        },
        }
        },
    },
    PublicNetworkAccess = "Enabled",
    IsLocalAuthDisabled = false,
    IsAadAuthDisabled = false,
    Tags =
    {
    ["key1"] = "value1",
    },
};
ArmOperation<WebPubSubResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
WebPubSubResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WebPubSubData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
