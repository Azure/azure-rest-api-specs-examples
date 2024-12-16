
import com.azure.resourcemanager.iotoperations.models.BatchingConfiguration;
import com.azure.resourcemanager.iotoperations.models.BrokerProtocolType;
import com.azure.resourcemanager.iotoperations.models.CloudEventAttributeType;
import com.azure.resourcemanager.iotoperations.models.DataExplorerAuthMethod;
import com.azure.resourcemanager.iotoperations.models.DataLakeStorageAuthMethod;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationAccessToken;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationSasl;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationSaslType;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationServiceAccountToken;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationSystemAssignedManagedIdentity;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationUserAssignedManagedIdentity;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationX509;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointDataExplorer;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointDataExplorerAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointDataLakeStorage;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointDataLakeStorageAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointFabricOneLake;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointFabricOneLakeAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointFabricOneLakeNames;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointFabricPathType;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafka;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaAcks;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaBatching;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaCompression;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaPartitionStrategy;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointLocalStorage;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointMqtt;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointMqttAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointProperties;
import com.azure.resourcemanager.iotoperations.models.EndpointType;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.FabricOneLakeAuthMethod;
import com.azure.resourcemanager.iotoperations.models.KafkaAuthMethod;
import com.azure.resourcemanager.iotoperations.models.MqttAuthMethod;
import com.azure.resourcemanager.iotoperations.models.MqttRetainType;
import com.azure.resourcemanager.iotoperations.models.OperationalMode;
import com.azure.resourcemanager.iotoperations.models.TlsProperties;

/**
 * Samples for DataflowEndpoint CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/DataflowEndpoint_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowEndpoint_CreateOrUpdate.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowEndpointCreateOrUpdate(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().define("resource-name123")
            .withExistingInstance("rgiotoperations", "resource-name123")
            .withExtendedLocation(new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(new DataflowEndpointProperties().withEndpointType(EndpointType.DATA_EXPLORER)
                .withDataExplorerSettings(new DataflowEndpointDataExplorer()
                    .withAuthentication(new DataflowEndpointDataExplorerAuthentication()
                        .withMethod(DataExplorerAuthMethod.SYSTEM_ASSIGNED_MANAGED_IDENTITY)
                        .withSystemAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationSystemAssignedManagedIdentity()
                                .withAudience("psxomrfbhoflycm"))
                        .withUserAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationUserAssignedManagedIdentity()
                                .withClientId("fb90f267-8872-431a-a76a-a1cec5d3c4d2").withScope("zop")
                                .withTenantId("ed060aa2-71ff-4d3f-99c4-a9138356fdec")))
                    .withDatabase("yqcdpjsifm").withHost("<cluster>.<region>.kusto.windows.net")
                    .withBatching(new BatchingConfiguration().withLatencySeconds(9312).withMaxMessages(9028)))
                .withDataLakeStorageSettings(new DataflowEndpointDataLakeStorage()
                    .withAuthentication(new DataflowEndpointDataLakeStorageAuthentication()
                        .withMethod(DataLakeStorageAuthMethod.SYSTEM_ASSIGNED_MANAGED_IDENTITY)
                        .withAccessTokenSettings(
                            new DataflowEndpointAuthenticationAccessToken().withSecretRef("fakeTokenPlaceholder"))
                        .withSystemAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationSystemAssignedManagedIdentity()
                                .withAudience("psxomrfbhoflycm"))
                        .withUserAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationUserAssignedManagedIdentity()
                                .withClientId("fb90f267-8872-431a-a76a-a1cec5d3c4d2").withScope("zop")
                                .withTenantId("ed060aa2-71ff-4d3f-99c4-a9138356fdec")))
                    .withHost("<account>.blob.core.windows.net")
                    .withBatching(new BatchingConfiguration().withLatencySeconds(9312).withMaxMessages(9028)))
                .withFabricOneLakeSettings(new DataflowEndpointFabricOneLake()
                    .withAuthentication(new DataflowEndpointFabricOneLakeAuthentication()
                        .withMethod(FabricOneLakeAuthMethod.SYSTEM_ASSIGNED_MANAGED_IDENTITY)
                        .withSystemAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationSystemAssignedManagedIdentity()
                                .withAudience("psxomrfbhoflycm"))
                        .withUserAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationUserAssignedManagedIdentity()
                                .withClientId("fb90f267-8872-431a-a76a-a1cec5d3c4d2").withScope("zop")
                                .withTenantId("ed060aa2-71ff-4d3f-99c4-a9138356fdec")))
                    .withNames(new DataflowEndpointFabricOneLakeNames().withLakehouseName("wpeathi")
                        .withWorkspaceName("nwgmitkbljztgms"))
                    .withOneLakePathType(DataflowEndpointFabricPathType.FILES)
                    .withHost("https://<host>.fabric.microsoft.com")
                    .withBatching(new BatchingConfiguration().withLatencySeconds(9312).withMaxMessages(9028)))
                .withKafkaSettings(new DataflowEndpointKafka()
                    .withAuthentication(new DataflowEndpointKafkaAuthentication()
                        .withMethod(KafkaAuthMethod.SYSTEM_ASSIGNED_MANAGED_IDENTITY)
                        .withSystemAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationSystemAssignedManagedIdentity()
                                .withAudience("psxomrfbhoflycm"))
                        .withUserAssignedManagedIdentitySettings(
                            new DataflowEndpointAuthenticationUserAssignedManagedIdentity()
                                .withClientId("fb90f267-8872-431a-a76a-a1cec5d3c4d2").withScope("zop")
                                .withTenantId("ed060aa2-71ff-4d3f-99c4-a9138356fdec"))
                        .withSaslSettings(new DataflowEndpointAuthenticationSasl()
                            .withSaslType(DataflowEndpointAuthenticationSaslType.PLAIN)
                            .withSecretRef("fakeTokenPlaceholder"))
                        .withX509CertificateSettings(
                            new DataflowEndpointAuthenticationX509().withSecretRef("fakeTokenPlaceholder")))
                    .withConsumerGroupId("ukkzcjiyenhxokat").withHost("pwcqfiqclcgneolpewnyavoulbip")
                    .withBatching(new DataflowEndpointKafkaBatching().withMode(OperationalMode.ENABLED)
                        .withLatencyMs(3679).withMaxBytes(8887).withMaxMessages(2174))
                    .withCopyMqttProperties(OperationalMode.ENABLED)
                    .withCompression(DataflowEndpointKafkaCompression.NONE)
                    .withKafkaAcks(DataflowEndpointKafkaAcks.ZERO)
                    .withPartitionStrategy(DataflowEndpointKafkaPartitionStrategy.DEFAULT)
                    .withTls(new TlsProperties().withMode(OperationalMode.ENABLED)
                        .withTrustedCaCertificateConfigMapRef("tectjjvukvelsreihwadh"))
                    .withCloudEventAttributes(CloudEventAttributeType.fromString("PassThrough")))
                .withLocalStorageSettings(new DataflowEndpointLocalStorage().withPersistentVolumeClaimRef("jjwqwvd"))
                .withMqttSettings(
                    new DataflowEndpointMqtt()
                        .withAuthentication(new DataflowEndpointMqttAuthentication()
                            .withMethod(MqttAuthMethod.SYSTEM_ASSIGNED_MANAGED_IDENTITY)
                            .withSystemAssignedManagedIdentitySettings(
                                new DataflowEndpointAuthenticationSystemAssignedManagedIdentity()
                                    .withAudience("psxomrfbhoflycm"))
                            .withUserAssignedManagedIdentitySettings(
                                new DataflowEndpointAuthenticationUserAssignedManagedIdentity()
                                    .withClientId("fb90f267-8872-431a-a76a-a1cec5d3c4d2").withScope("zop")
                                    .withTenantId("ed060aa2-71ff-4d3f-99c4-a9138356fdec"))
                            .withServiceAccountTokenSettings(new DataflowEndpointAuthenticationServiceAccountToken()
                                .withAudience("ejbklrbxgjaqleoycgpje"))
                            .withX509CertificateSettings(
                                new DataflowEndpointAuthenticationX509().withSecretRef("fakeTokenPlaceholder")))
                        .withClientIdPrefix("kkljsdxdirfhwxtkavldekeqhv").withHost("nyhnxqnbspstctl")
                        .withProtocol(BrokerProtocolType.MQTT).withKeepAliveSeconds(0).withRetain(MqttRetainType.KEEP)
                        .withMaxInflightMessages(0).withQos(1).withSessionExpirySeconds(0)
                        .withTls(new TlsProperties().withMode(OperationalMode.ENABLED)
                            .withTrustedCaCertificateConfigMapRef("tectjjvukvelsreihwadh"))
                        .withCloudEventAttributes(CloudEventAttributeType.fromString("PassThrough"))))
            .create();
    }
}
