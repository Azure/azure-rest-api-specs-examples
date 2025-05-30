using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateAwsAssumeRoleConnectorSubscription_example.json
// this example is just showing the usage of "Connectors_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this SecurityCloudConnectorResource
SecurityCloudConnectorCollection collection = subscriptionResource.GetSecurityCloudConnectors();

// invoke the operation
string connectorName = "aws_dev2";
SecurityCloudConnectorData data = new SecurityCloudConnectorData
{
    HybridComputeSettings = new HybridComputeSettingsProperties(AutoProvisionState.On)
    {
        ResourceGroupName = "AwsConnectorRG",
        Region = "West US 2",
        ProxyServer = new ProxyServerProperties
        {
            IP = "167.220.197.140",
            Port = "34",
        },
        ServicePrincipal = new ServicePrincipalProperties
        {
            ApplicationId = Guid.Parse("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
            Secret = "<secret>",
        },
    },
    AuthenticationDetails = new AwsAssumeRoleAuthenticationDetailsProperties("arn:aws:iam::81231569658:role/AscConnector", Guid.Parse("20ff7fc3-e762-44dd-bd96-b71116dcdc23")),
};
ArmOperation<SecurityCloudConnectorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, connectorName, data);
SecurityCloudConnectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityCloudConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
