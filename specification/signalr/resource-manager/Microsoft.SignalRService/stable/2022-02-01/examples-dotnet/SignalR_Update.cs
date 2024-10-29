using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SignalR;
using Azure.ResourceManager.SignalR.Models;

// Generated from example definition: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalR_Update.json
// this example is just showing the usage of "SignalR_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SignalRResource created on azure
// for more information of creating SignalRResource, please refer to the document of SignalRResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "mySignalRService";
ResourceIdentifier signalRResourceId = SignalRResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
SignalRResource signalR = client.GetSignalRResource(signalRResourceId);

// invoke the operation
SignalRData data = new SignalRData(new AzureLocation("eastus"))
{
    Sku = new SignalRResourceSku("Standard_S1")
    {
        Tier = SignalRSkuTier.Standard,
        Capacity = 1,
    },
    Kind = SignalRServiceKind.SignalR,
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    IsClientCertEnabled = false,
    Features =
    {
    new SignalRFeature(SignalRFeatureFlag.ServiceMode,"Serverless")
    {
    Properties =
    {
    },
    },new SignalRFeature(SignalRFeatureFlag.EnableConnectivityLogs,"True")
    {
    Properties =
    {
    },
    },new SignalRFeature(SignalRFeatureFlag.EnableMessagingLogs,"False")
    {
    Properties =
    {
    },
    },new SignalRFeature(SignalRFeatureFlag.EnableLiveTrace,"False")
    {
    Properties =
    {
    },
    }
    },
    LiveTraceConfiguration = new SignalRLiveTraceConfiguration()
    {
        Enabled = "false",
        Categories =
        {
        new SignalRLiveTraceCategory()
        {
        Name = "ConnectivityLogs",
        Enabled = "true",
        }
        },
    },
    CorsAllowedOrigins =
    {
    "https://foo.com","https://bar.com"
    },
    UpstreamTemplates =
    {
    new SignalRUpstreamTemplate("https://example.com/chat/api/connect")
    {
    HubPattern = "*",
    EventPattern = "connect,disconnect",
    CategoryPattern = "*",
    Auth = new SignalRUpstreamAuthSettings()
    {
    AuthType = SignalRUpstreamAuthType.ManagedIdentity,
    ManagedIdentityResource = "api://example",
    },
    }
    },
    NetworkACLs = new SignalRNetworkAcls()
    {
        DefaultAction = SignalRNetworkAclAction.Deny,
        PublicNetwork = new SignalRNetworkAcl()
        {
            Allow =
            {
            SignalRRequestType.ClientConnection
            },
        },
        PrivateEndpoints =
        {
        new SignalRPrivateEndpointAcl("mysignalrservice.1fa229cd-bf3f-47f0-8c49-afb36723997e")
        {
        Allow =
        {
        SignalRRequestType.ServerConnection
        },
        }
        },
    },
    PublicNetworkAccess = "Enabled",
    DisableLocalAuth = false,
    DisableAadAuth = false,
    Tags =
    {
    ["key1"] = "value1",
    },
};
ArmOperation<SignalRResource> lro = await signalR.UpdateAsync(WaitUntil.Completed, data);
SignalRResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SignalRData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
