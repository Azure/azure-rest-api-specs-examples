using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/DataflowEndpoint_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "DataflowEndpointResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsDataflowEndpointResource created on azure
// for more information of creating IotOperationsDataflowEndpointResource, please refer to the document of IotOperationsDataflowEndpointResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string dataflowEndpointName = "resource-name123";
ResourceIdentifier iotOperationsDataflowEndpointResourceId = IotOperationsDataflowEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, dataflowEndpointName);
IotOperationsDataflowEndpointResource iotOperationsDataflowEndpoint = client.GetIotOperationsDataflowEndpointResource(iotOperationsDataflowEndpointResourceId);

// invoke the operation
IotOperationsDataflowEndpointData data = new IotOperationsDataflowEndpointData
{
    Properties = new IotOperationsDataflowEndpointProperties(DataflowEndpointType.DataExplorer)
    {
        DataExplorerSettings = new DataflowEndpointDataExplorer(new DataflowEndpointDataExplorerAuthentication(DataExplorerAuthMethod.SystemAssignedManagedIdentity)
        {
            SystemAssignedManagedIdentityAudience = "psxomrfbhoflycm",
            UserAssignedManagedIdentitySettings = new DataflowEndpointAuthenticationUserAssignedManagedIdentity("fb90f267-8872-431a-a76a-a1cec5d3c4d2", "ed060aa2-71ff-4d3f-99c4-a9138356fdec")
            {
                Scope = "zop",
            },
        }, "yqcdpjsifm", "<cluster>.<region>.kusto.windows.net")
        {
            Batching = new IotOperationsBatchingConfig
            {
                LatencySeconds = 9312,
                MaxMessages = 9028,
            },
        },
        DataLakeStorageSettings = new DataflowEndpointDataLakeStorage(new DataflowEndpointDataLakeStorageAuthentication(DataLakeStorageAuthMethod.SystemAssignedManagedIdentity)
        {
            AccessTokenSecretRef = "sevriyphcvnlrnfudqzejecwa",
            SystemAssignedManagedIdentityAudience = "psxomrfbhoflycm",
            UserAssignedManagedIdentitySettings = new DataflowEndpointAuthenticationUserAssignedManagedIdentity("fb90f267-8872-431a-a76a-a1cec5d3c4d2", "ed060aa2-71ff-4d3f-99c4-a9138356fdec")
            {
                Scope = "zop",
            },
        }, "<account>.blob.core.windows.net")
        {
            Batching = new IotOperationsBatchingConfig
            {
                LatencySeconds = 9312,
                MaxMessages = 9028,
            },
        },
        FabricOneLakeSettings = new DataflowEndpointFabricOneLake(new DataflowEndpointFabricOneLakeAuthentication(FabricOneLakeAuthMethod.SystemAssignedManagedIdentity)
        {
            SystemAssignedManagedIdentityAudience = "psxomrfbhoflycm",
            UserAssignedManagedIdentitySettings = new DataflowEndpointAuthenticationUserAssignedManagedIdentity("fb90f267-8872-431a-a76a-a1cec5d3c4d2", "ed060aa2-71ff-4d3f-99c4-a9138356fdec")
            {
                Scope = "zop",
            },
        }, new DataflowEndpointFabricOneLakeNames("wpeathi", "nwgmitkbljztgms"), DataflowEndpointFabricPathType.Files, "https://<host>.fabric.microsoft.com")
        {
            Batching = new IotOperationsBatchingConfig
            {
                LatencySeconds = 9312,
                MaxMessages = 9028,
            },
        },
        KafkaSettings = new DataflowEndpointKafka(new DataflowEndpointKafkaAuthentication(KafkaAuthMethod.SystemAssignedManagedIdentity)
        {
            SystemAssignedManagedIdentityAudience = "psxomrfbhoflycm",
            UserAssignedManagedIdentitySettings = new DataflowEndpointAuthenticationUserAssignedManagedIdentity("fb90f267-8872-431a-a76a-a1cec5d3c4d2", "ed060aa2-71ff-4d3f-99c4-a9138356fdec")
            {
                Scope = "zop",
            },
            SaslSettings = new DataflowEndpointAuthenticationSasl(DataflowEndpointAuthenticationSaslType.Plain, "visyxoztqnylvbyokhtmpdkwes"),
            X509CertificateSecretRef = "afwizrystfslkfqd",
        }, "pwcqfiqclcgneolpewnyavoulbip")
        {
            ConsumerGroupId = "ukkzcjiyenhxokat",
            Batching = new DataflowEndpointKafkaBatching
            {
                Mode = IotOperationsOperationalMode.Enabled,
                LatencyMs = 3679,
                MaxBytes = 8887,
                MaxMessages = 2174,
            },
            CopyMqttProperties = IotOperationsOperationalMode.Enabled,
            Compression = DataflowEndpointKafkaCompression.None,
            KafkaAcks = DataflowEndpointKafkaAck.Zero,
            PartitionStrategy = DataflowEndpointKafkaPartitionStrategy.Default,
            Tls = new IotOperationsTlsProperties
            {
                Mode = IotOperationsOperationalMode.Enabled,
                TrustedCaCertificateConfigMapRef = "tectjjvukvelsreihwadh",
            },
            CloudEventAttributes = new CloudEventAttributeType("PassThrough"),
        },
        LocalStoragePersistentVolumeClaimRef = "jjwqwvd",
        MqttSettings = new DataflowEndpointMqtt(new DataflowEndpointMqttAuthentication(MqttAuthMethod.SystemAssignedManagedIdentity)
        {
            SystemAssignedManagedIdentityAudience = "psxomrfbhoflycm",
            UserAssignedManagedIdentitySettings = new DataflowEndpointAuthenticationUserAssignedManagedIdentity("fb90f267-8872-431a-a76a-a1cec5d3c4d2", "ed060aa2-71ff-4d3f-99c4-a9138356fdec")
            {
                Scope = "zop",
            },
            ServiceAccountTokenAudience = "ejbklrbxgjaqleoycgpje",
            X509CertificateSecretRef = "afwizrystfslkfqd",
        })
        {
            ClientIdPrefix = "kkljsdxdirfhwxtkavldekeqhv",
            Host = "nyhnxqnbspstctl",
            Protocol = BrokerProtocolType.Mqtt,
            KeepAliveSeconds = 0,
            Retain = MqttRetainType.Keep,
            MaxInflightMessages = 0,
            Qos = 1,
            SessionExpirySeconds = 0,
            Tls = new IotOperationsTlsProperties
            {
                Mode = IotOperationsOperationalMode.Enabled,
                TrustedCaCertificateConfigMapRef = "tectjjvukvelsreihwadh",
            },
            CloudEventAttributes = new CloudEventAttributeType("PassThrough"),
        },
    },
    ExtendedLocation = new IotOperationsExtendedLocation("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123", IotOperationsExtendedLocationType.CustomLocation),
};
ArmOperation<IotOperationsDataflowEndpointResource> lro = await iotOperationsDataflowEndpoint.UpdateAsync(WaitUntil.Completed, data);
IotOperationsDataflowEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsDataflowEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
