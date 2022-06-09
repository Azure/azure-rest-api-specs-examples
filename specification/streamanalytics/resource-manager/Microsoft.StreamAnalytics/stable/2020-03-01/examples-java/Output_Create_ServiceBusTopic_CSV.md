```java
import com.azure.resourcemanager.streamanalytics.models.CsvSerialization;
import com.azure.resourcemanager.streamanalytics.models.Encoding;
import com.azure.resourcemanager.streamanalytics.models.ServiceBusTopicOutputDataSource;
import java.util.Arrays;

/** Samples for Outputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_ServiceBusTopic_CSV.json
     */
    /**
     * Sample code: Create a Service Bus Topic output with CSV serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAServiceBusTopicOutputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .outputs()
            .define("output7886")
            .withExistingStreamingjob("sjrg6450", "sj7094")
            .withDatasource(
                new ServiceBusTopicOutputDataSource()
                    .withTopicName("sdktopic")
                    .withPropertyColumns(Arrays.asList("column1", "column2"))
                    .withServiceBusNamespace("sdktest")
                    .withSharedAccessPolicyName("RootManageSharedAccessKey")
                    .withSharedAccessPolicyKey("sharedAccessPolicyKey="))
            .withSerialization(new CsvSerialization().withFieldDelimiter(",").withEncoding(Encoding.UTF8))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
