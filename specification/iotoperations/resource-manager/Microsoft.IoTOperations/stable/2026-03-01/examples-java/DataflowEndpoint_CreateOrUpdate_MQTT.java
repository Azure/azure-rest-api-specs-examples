
import com.azure.resourcemanager.iotoperations.models.BrokerProtocolType;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationX509;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointMqtt;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointMqttAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointProperties;
import com.azure.resourcemanager.iotoperations.models.EndpointType;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.MqttAuthMethod;
import com.azure.resourcemanager.iotoperations.models.MqttRetainType;
import com.azure.resourcemanager.iotoperations.models.OperationalMode;
import com.azure.resourcemanager.iotoperations.models.TlsProperties;

/**
 * Samples for DataflowEndpoint CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowEndpoint_CreateOrUpdate_MQTT.json
     */
    /**
     * Sample code: DataflowEndpoint_CreateOrUpdate_MQTT.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowEndpointCreateOrUpdateMQTT(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().define("generic-mqtt-broker-endpoint")
            .withExistingInstance("rgiotoperations",
                "resource-name123")
            .withProperties(
                new DataflowEndpointProperties()
                    .withEndpointType(
                        EndpointType.MQTT)
                    .withMqttSettings(new DataflowEndpointMqtt()
                        .withAuthentication(new DataflowEndpointMqttAuthentication()
                            .withMethod(MqttAuthMethod.X509CERTIFICATE).withX509CertificateSettings(
                                new DataflowEndpointAuthenticationX509().withSecretRef("fakeTokenPlaceholder")))
                        .withClientIdPrefix("factory-gateway").withHost("example.broker.local:1883")
                        .withProtocol(BrokerProtocolType.WEB_SOCKETS).withKeepAliveSeconds(60)
                        .withRetain(MqttRetainType.KEEP).withMaxInflightMessages(100).withQos(1)
                        .withSessionExpirySeconds(3600)
                        .withTls(new TlsProperties().withMode(OperationalMode.DISABLED))))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .create();
    }
}
