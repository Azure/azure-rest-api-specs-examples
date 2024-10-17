using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotHub.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.IotHub;

// Generated from example definition: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_addRoutingCosmosDBEndpoint.json
// this example is just showing the usage of "IotHubResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this IotHubDescriptionResource
IotHubDescriptionCollection collection = resourceGroupResource.GetIotHubDescriptions();

// invoke the operation
string resourceName = "testHub";
IotHubDescriptionData data = new IotHubDescriptionData(new AzureLocation("centraluseuap"), new IotHubSkuInfo(IotHubSku.S1)
{
    Capacity = 1L,
})
{
    ETag = new ETag("AAAAAAFD6M4="),
    Properties = new IotHubProperties()
    {
        IPFilterRules =
        {
        },
        NetworkRuleSets = new IotHubNetworkRuleSetProperties(true, new IotHubNetworkRuleSetIPRule[]
{
new IotHubNetworkRuleSetIPRule("rule1","131.117.159.53")
{
Action = IotHubNetworkRuleIPAction.Allow,
},new IotHubNetworkRuleSetIPRule("rule2","157.55.59.128/25")
{
Action = IotHubNetworkRuleIPAction.Allow,
}
})
        {
            DefaultAction = IotHubNetworkRuleSetDefaultAction.Deny,
        },
        MinTlsVersion = "1.2",
        EventHubEndpoints =
        {
        ["events"] = new EventHubCompatibleEndpointProperties()
        {
        RetentionTimeInDays = 1L,
        PartitionCount = 2,
        },
        },
        Routing = new IotHubRoutingProperties()
        {
            Endpoints = new RoutingEndpoints()
            {
                ServiceBusQueues =
                {
                },
                ServiceBusTopics =
                {
                },
                EventHubs =
                {
                },
                StorageContainers =
                {
                },
                CosmosDBSqlContainers =
                {
                new RoutingCosmosDBSqlApiProperties("endpointcosmos",new Uri("https://test-systemstore-test2.documents.azure.com"),"systemstore","test")
                {
                SubscriptionId = "<subscription-id>",
                ResourceGroup = "rg-test",
                AuthenticationType = IotHubAuthenticationType.KeyBased,
                PrimaryKey = "<primary-key>",
                SecondaryKey = "<secondary-key>",
                PartitionKeyName = "keystamped",
                PartitionKeyTemplate = "{deviceid}-{YYYY}-{MM}",
                }
                },
            },
            Routes =
            {
            },
            FallbackRoute = new IotHubFallbackRouteProperties(IotHubRoutingSource.DeviceMessages, new string[]
{
"events"
}, true)
            {
                Name = "$fallback",
                Condition = "true",
            },
        },
        StorageEndpoints =
        {
        ["$default"] = new IotHubStorageEndpointProperties("","")
        {
        SasTtlAsIso8601 = XmlConvert.ToTimeSpan("PT1H"),
        },
        },
        MessagingEndpoints =
        {
        ["fileNotifications"] = new MessagingEndpointProperties()
        {
        LockDurationAsIso8601 = XmlConvert.ToTimeSpan("PT1M"),
        TtlAsIso8601 = XmlConvert.ToTimeSpan("PT1H"),
        MaxDeliveryCount = 10,
        },
        },
        EnableFileUploadNotifications = false,
        CloudToDevice = new CloudToDeviceProperties()
        {
            MaxDeliveryCount = 10,
            DefaultTtlAsIso8601 = XmlConvert.ToTimeSpan("PT1H"),
            Feedback = new CloudToDeviceFeedbackQueueProperties()
            {
                LockDurationAsIso8601 = XmlConvert.ToTimeSpan("PT1M"),
                TtlAsIso8601 = XmlConvert.ToTimeSpan("PT1H"),
                MaxDeliveryCount = 10,
            },
        },
        Features = IotHubCapability.None,
        EnableDataResidency = false,
    },
    Tags =
    {
    },
};
ArmOperation<IotHubDescriptionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
IotHubDescriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotHubDescriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
