using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateAwsCredConnectorSubscription_example.json
// this example is just showing the usage of "Connectors_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityCloudConnectorResource created on azure
// for more information of creating SecurityCloudConnectorResource, please refer to the document of SecurityCloudConnectorResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string connectorName = "aws_dev1";
ResourceIdentifier securityCloudConnectorResourceId = SecurityCloudConnectorResource.CreateResourceIdentifier(subscriptionId, connectorName);
SecurityCloudConnectorResource securityCloudConnector = client.GetSecurityCloudConnectorResource(securityCloudConnectorResourceId);

// invoke the operation
SecurityCloudConnectorData data = new SecurityCloudConnectorData()
{
    HybridComputeSettings = new HybridComputeSettingsProperties(AutoProvisionState.On)
    {
        ResourceGroupName = "AwsConnectorRG",
        Region = "West US 2",
        ProxyServer = new ProxyServerProperties()
        {
            IP = "167.220.197.140",
            Port = "34",
        },
        ServicePrincipal = new ServicePrincipalProperties()
        {
            ApplicationId = Guid.Parse("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
            Secret = "<secret>",
        },
    },
    AuthenticationDetails = new AwsCredsAuthenticationDetailsProperties("AKIARPZCNODDNAEQFSOE", "<awsSecretAccessKey>"),
};
ArmOperation<SecurityCloudConnectorResource> lro = await securityCloudConnector.UpdateAsync(WaitUntil.Completed, data);
SecurityCloudConnectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityCloudConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
