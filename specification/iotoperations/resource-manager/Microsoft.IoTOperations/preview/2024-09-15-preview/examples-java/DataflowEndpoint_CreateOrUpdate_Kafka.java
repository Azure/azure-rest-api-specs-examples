
import com.azure.resourcemanager.iotoperations.models.CloudEventAttributeType;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationSasl;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointAuthenticationSaslType;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafka;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaAcks;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaAuthentication;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaBatching;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaCompression;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointKafkaPartitionStrategy;
import com.azure.resourcemanager.iotoperations.models.DataflowEndpointProperties;
import com.azure.resourcemanager.iotoperations.models.EndpointType;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.KafkaAuthMethod;
import com.azure.resourcemanager.iotoperations.models.OperationalMode;
import com.azure.resourcemanager.iotoperations.models.TlsProperties;

/**
 * Samples for DataflowEndpoint CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_Kafka.json
     */
    /**
     * Sample code: DataflowEndpoint_CreateOrUpdate_Kafka.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowEndpointCreateOrUpdateKafka(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().define("generic-kafka-endpoint")
            .withExistingInstance("rgiotoperations", "resource-name123")
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(new DataflowEndpointProperties().withEndpointType(EndpointType.KAFKA)
                .withKafkaSettings(new DataflowEndpointKafka()
                    .withAuthentication(new DataflowEndpointKafkaAuthentication().withMethod(KafkaAuthMethod.SASL)
                        .withSaslSettings(new DataflowEndpointAuthenticationSasl()
                            .withSaslType(DataflowEndpointAuthenticationSaslType.PLAIN)
                            .withSecretRef("fakeTokenPlaceholder")))
                    .withConsumerGroupId("dataflows").withHost("example.kafka.local:9093")
                    .withBatching(new DataflowEndpointKafkaBatching().withMode(OperationalMode.ENABLED).withLatencyMs(5)
                        .withMaxBytes(1000000).withMaxMessages(100000))
                    .withCopyMqttProperties(OperationalMode.ENABLED)
                    .withCompression(DataflowEndpointKafkaCompression.GZIP).withKafkaAcks(DataflowEndpointKafkaAcks.ALL)
                    .withPartitionStrategy(DataflowEndpointKafkaPartitionStrategy.DEFAULT)
                    .withTls(new TlsProperties().withMode(OperationalMode.ENABLED)
                        .withTrustedCaCertificateConfigMapRef("ca-certificates"))
                    .withCloudEventAttributes(CloudEventAttributeType.PROPAGATE)))
            .create();
    }
}
