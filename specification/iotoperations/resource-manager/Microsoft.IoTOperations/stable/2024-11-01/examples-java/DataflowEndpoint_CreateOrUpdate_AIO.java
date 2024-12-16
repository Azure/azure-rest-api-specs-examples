
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationServiceAccountToken;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointMqtt;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointMqttAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointProperties;
import com.azure.resourcemanager.iotoperations.models.EndpointType;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.MqttAuthMethod;
import com.azure.resourcemanager.iotoperations.models.OperationalMode;
import com.azure.resourcemanager.iotoperations.models.TlsProperties;

/**
 * Samples for DataflowEndpoint CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/DataflowEndpoint_CreateOrUpdate_AIO.json
     */
    /**
     * Sample code: DataflowEndpoint_CreateOrUpdate_AIO.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowEndpointCreateOrUpdateAIO(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().define("aio-builtin-broker-endpoint")
            .withExistingInstance("rgiotoperations", "resource-name123")
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(
                new DataflowEndpointProperties().withEndpointType(EndpointType.MQTT)
                    .withMqttSettings(new DataflowEndpointMqtt()
                        .withAuthentication(
                            new DataflowEndpointMqttAuthentication().withMethod(MqttAuthMethod.fromString("Kubernetes"))
                                .withServiceAccountTokenSettings(new DataflowEndpointAuthenticationServiceAccountToken()
                                    .withAudience("aio-internal")))
                        .withHost("aio-broker:18883").withTls(new TlsProperties().withMode(OperationalMode.ENABLED)
                            .withTrustedCaCertificateConfigMapRef("aio-ca-trust-bundle-test-only"))))
            .create();
    }
}
