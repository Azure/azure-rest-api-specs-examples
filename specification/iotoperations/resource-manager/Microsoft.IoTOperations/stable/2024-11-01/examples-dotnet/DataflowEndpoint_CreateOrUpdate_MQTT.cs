using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2024-11-01/DataflowEndpoint_CreateOrUpdate_MQTT.json
// this example is just showing the usage of "DataflowEndpointResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsInstanceResource created on azure
// for more information of creating IotOperationsInstanceResource, please refer to the document of IotOperationsInstanceResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
ResourceIdentifier iotOperationsInstanceResourceId = IotOperationsInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName);
IotOperationsInstanceResource iotOperationsInstance = client.GetIotOperationsInstanceResource(iotOperationsInstanceResourceId);

// get the collection of this IotOperationsDataflowEndpointResource
IotOperationsDataflowEndpointCollection collection = iotOperationsInstance.GetIotOperationsDataflowEndpoints();

// invoke the operation
string dataflowEndpointName = "generic-mqtt-broker-endpoint";
IotOperationsDataflowEndpointData data = new IotOperationsDataflowEndpointData(new IotOperationsExtendedLocation("qmbrfwcpwwhggszhrdjv", IotOperationsExtendedLocationType.CustomLocation))
{
    Properties = new IotOperationsDataflowEndpointProperties(DataflowEndpointType.Mqtt)
    {
        MqttSettings = new DataflowEndpointMqtt(new DataflowEndpointMqttAuthentication(MqttAuthMethod.X509Certificate)
        {
            X509CertificateSecretRef = "example-secret",
        })
        {
            ClientIdPrefix = "factory-gateway",
            Host = "example.broker.local:1883",
            Protocol = BrokerProtocolType.WebSockets,
            KeepAliveSeconds = 60,
            Retain = MqttRetainType.Keep,
            MaxInflightMessages = 100,
            Qos = 1,
            SessionExpirySeconds = 3600,
            Tls = new IotOperationsTlsProperties
            {
                Mode = IotOperationsOperationalMode.Disabled,
            },
        },
    },
};
ArmOperation<IotOperationsDataflowEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataflowEndpointName, data);
IotOperationsDataflowEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsDataflowEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
