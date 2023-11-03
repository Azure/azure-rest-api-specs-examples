using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ConfigurationGroupValueFirstPartyCreate.json
// this example is just showing the usage of "ConfigurationGroupValues_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ConfigurationGroupValueResource
ConfigurationGroupValueCollection collection = resourceGroupResource.GetConfigurationGroupValues();

// invoke the operation
string configurationGroupValueName = "testConfigurationGroupValue";
ConfigurationGroupValueData data = new ConfigurationGroupValueData(new AzureLocation("eastus"))
{
    Properties = new ConfigurationValueWithoutSecrets()
    {
        ConfigurationValue = "{\"interconnect-groups\":{\"stripe-one\":{\"name\":\"Stripe one\",\"international-interconnects\":[\"france\",\"germany\"],\"domestic-interconnects\":[\"birmingham\",\"edinburgh\"]},\"stripe-two\":{\"name\":\"Stripe two\",\"international-interconnects\":[\"germany\",\"italy\"],\"domestic-interconnects\":[\"edinburgh\",\"london\"]}},\"interconnect-group-assignments\":{\"ssc-one\":{\"ssc\":\"SSC 1\",\"interconnects\":\"stripe-one\"},\"ssc-two\":{\"ssc\":\"SSC 2\",\"interconnects\":\"stripe-two\"}}}",
        ConfigurationGroupSchemaResourceReference = new SecretDeploymentResourceReference()
        {
            Id = new ResourceIdentifier("/subscriptions/subid/resourcegroups/testRG/providers/microsoft.hybridnetwork/publishers/testPublisher/configurationGroupSchemas/testConfigurationGroupSchemaName"),
        },
    },
};
ArmOperation<ConfigurationGroupValueResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, configurationGroupValueName, data);
ConfigurationGroupValueResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConfigurationGroupValueData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
