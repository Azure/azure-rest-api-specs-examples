
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.ServiceBusTopicOutputDataSource;
import java.util.Arrays;

/**
 * Samples for Outputs CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Output_Create_ServiceBusTopic_CSV.json
     */
    /**
     * Sample code: Create a Service Bus Topic output with CSV serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAServiceBusTopicOutputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().define("output7886").withExistingStreamingjob("sjrg6450", "sj7094")
            .withDatasource(new ServiceBusTopicOutputDataSource().withTopicName("sdktopic")
                .withPropertyColumns(Arrays.asList("column1", "column2")).withServiceBusNamespace("sdktest")
                .withSharedAccessPolicyName("RootManageSharedAccessKey")
                .withSharedAccessPolicyKey("fakeTokenPlaceholder"))
            .withSerialization(new CsvSerialization().withFieldDelimiter(",").withEncoding(Encoding.UTF8)).create();
    }
}
