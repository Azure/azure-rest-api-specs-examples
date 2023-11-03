using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ConfigurationGroupSchemaCreate.json
// this example is just showing the usage of "ConfigurationGroupSchemas_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PublisherResource created on azure
// for more information of creating PublisherResource, please refer to the document of PublisherResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string publisherName = "testPublisher";
ResourceIdentifier publisherResourceId = PublisherResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, publisherName);
PublisherResource publisher = client.GetPublisherResource(publisherResourceId);

// get the collection of this ConfigurationGroupSchemaResource
ConfigurationGroupSchemaCollection collection = publisher.GetConfigurationGroupSchemas();

// invoke the operation
string configurationGroupSchemaName = "testConfigurationGroupSchema";
ConfigurationGroupSchemaData data = new ConfigurationGroupSchemaData(new AzureLocation("westUs2"))
{
    Properties = new ConfigurationGroupSchemaPropertiesFormat()
    {
        Description = "Schema with no secrets",
        SchemaDefinition = "{\"type\":\"object\",\"properties\":{\"interconnect-groups\":{\"type\":\"object\",\"properties\":{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"international-interconnects\":{\"type\":\"array\",\"item\":{\"type\":\"string\"}},\"domestic-interconnects\":{\"type\":\"array\",\"item\":{\"type\":\"string\"}}}}},\"interconnect-group-assignments\":{\"type\":\"object\",\"properties\":{\"type\":\"object\",\"properties\":{\"ssc\":{\"type\":\"string\"},\"interconnects-interconnects\":{\"type\":\"string\"}}}}},\"required\":[\"interconnect-groups\",\"interconnect-group-assignments\"]}",
    },
};
ArmOperation<ConfigurationGroupSchemaResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, configurationGroupSchemaName, data);
ConfigurationGroupSchemaResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConfigurationGroupSchemaData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
