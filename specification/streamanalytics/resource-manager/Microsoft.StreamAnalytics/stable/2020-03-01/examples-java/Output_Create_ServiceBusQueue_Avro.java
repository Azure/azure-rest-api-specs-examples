import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.streamanalytics.models.AvroSerialization;
import com.azure.resourcemanager.streamanalytics.models.ServiceBusQueueOutputDataSource;
import java.io.IOException;
import java.util.Arrays;

/** Samples for Outputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_ServiceBusQueue_Avro.json
     */
    /**
     * Sample code: Create a Service Bus Queue output with Avro serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAServiceBusQueueOutputWithAvroSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) throws IOException {
        manager
            .outputs()
            .define("output3456")
            .withExistingStreamingjob("sjrg3410", "sj5095")
            .withDatasource(
                new ServiceBusQueueOutputDataSource()
                    .withQueueName("sdkqueue")
                    .withPropertyColumns(Arrays.asList("column1", "column2"))
                    .withSystemPropertyColumns(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize(
                                "{\"MessageId\":\"col3\",\"PartitionKey\":\"col4\"}",
                                Object.class,
                                SerializerEncoding.JSON))
                    .withServiceBusNamespace("sdktest")
                    .withSharedAccessPolicyName("RootManageSharedAccessKey")
                    .withSharedAccessPolicyKey("sharedAccessPolicyKey="))
            .withSerialization(new AvroSerialization())
            .create();
    }
}
