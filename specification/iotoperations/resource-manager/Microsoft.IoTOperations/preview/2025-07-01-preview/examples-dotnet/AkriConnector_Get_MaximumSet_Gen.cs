using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/AkriConnector_Get_MaximumSet_Gen.json
// this example is just showing the usage of "AkriConnectorResource_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsAkriConnectorTemplateResource created on azure
// for more information of creating IotOperationsAkriConnectorTemplateResource, please refer to the document of IotOperationsAkriConnectorTemplateResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string akriConnectorTemplateName = "resource-name123";
ResourceIdentifier iotOperationsAkriConnectorTemplateResourceId = IotOperationsAkriConnectorTemplateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, akriConnectorTemplateName);
IotOperationsAkriConnectorTemplateResource iotOperationsAkriConnectorTemplate = client.GetIotOperationsAkriConnectorTemplateResource(iotOperationsAkriConnectorTemplateResourceId);

// get the collection of this IotOperationsAkriConnectorResource
IotOperationsAkriConnectorCollection collection = iotOperationsAkriConnectorTemplate.GetIotOperationsAkriConnectors();

// invoke the operation
string connectorName = "resource-name123";
NullableResponse<IotOperationsAkriConnectorResource> response = await collection.GetIfExistsAsync(connectorName);
IotOperationsAkriConnectorResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    IotOperationsAkriConnectorData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
