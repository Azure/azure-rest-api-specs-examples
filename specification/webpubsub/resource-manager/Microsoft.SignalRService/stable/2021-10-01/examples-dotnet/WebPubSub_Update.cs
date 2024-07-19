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

// Generated from example definition: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_Update.json
// this example is just showing the usage of "WebPubSub_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebPubSubResource created on azure
// for more information of creating WebPubSubResource, please refer to the document of WebPubSubResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "myWebPubSubService";
ResourceIdentifier webPubSubResourceId = WebPubSubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
WebPubSubResource webPubSub = client.GetWebPubSubResource(webPubSubResourceId);

// invoke the operation
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
ArmOperation<WebPubSubResource> lro = await webPubSub.UpdateAsync(WaitUntil.Completed, data);
WebPubSubResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WebPubSubData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
