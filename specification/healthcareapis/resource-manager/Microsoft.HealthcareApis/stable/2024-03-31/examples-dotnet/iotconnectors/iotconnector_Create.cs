using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthcareApis.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.HealthcareApis;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/iotconnectors/iotconnector_Create.json
// this example is just showing the usage of "IotConnectors_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisWorkspaceResource created on azure
// for more information of creating HealthcareApisWorkspaceResource, please refer to the document of HealthcareApisWorkspaceResource
string subscriptionId = "subid";
string resourceGroupName = "testRG";
string workspaceName = "workspace1";
ResourceIdentifier healthcareApisWorkspaceResourceId = HealthcareApisWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
HealthcareApisWorkspaceResource healthcareApisWorkspace = client.GetHealthcareApisWorkspaceResource(healthcareApisWorkspaceResourceId);

// get the collection of this HealthcareApisIotConnectorResource
HealthcareApisIotConnectorCollection collection = healthcareApisWorkspace.GetHealthcareApisIotConnectors();

// invoke the operation
string iotConnectorName = "blue";
HealthcareApisIotConnectorData data = new HealthcareApisIotConnectorData(new AzureLocation("westus"))
{
    IngestionEndpointConfiguration = new HealthcareApisIotConnectorEventHubIngestionConfiguration()
    {
        EventHubName = "MyEventHubName",
        ConsumerGroup = "ConsumerGroupA",
        FullyQualifiedEventHubNamespace = "myeventhub.servicesbus.windows.net",
    },
    DeviceMappingContent = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["template"] = new object[] { new Dictionary<string, object>()
        {
        ["template"] = new Dictionary<string, object>()
        {
        ["deviceIdExpression"] = "$.deviceid",
        ["timestampExpression"] = "$.measurementdatetime",
        ["typeMatchExpression"] = "$..[?(@heartrate)]",
        ["typeName"] = "heartrate",
        ["values"] = new object[] { new Dictionary<string, object>()
        {
        ["required"] = "true",
        ["valueExpression"] = "$.heartrate",
        ["valueName"] = "hr"} }},
        ["templateType"] = "JsonPathContent"} },
        ["templateType"] = "CollectionContent"
    }),
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Tags =
    {
    ["additionalProp1"] = "string",
    ["additionalProp2"] = "string",
    ["additionalProp3"] = "string",
    },
};
ArmOperation<HealthcareApisIotConnectorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, iotConnectorName, data);
HealthcareApisIotConnectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthcareApisIotConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
