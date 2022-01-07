Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.streamanalytics.models.AzureTableOutputDataSource;
import java.util.Arrays;

/** Samples for Outputs CreateOrReplace. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_AzureTable.json
     */
    /**
     * Sample code: Create an Azure Table output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void createAnAzureTableOutput(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .outputs()
            .define("output958")
            .withExistingStreamingjob("sjrg5176", "sj2790")
            .withDatasource(
                new AzureTableOutputDataSource()
                    .withAccountName("someAccountName")
                    .withAccountKey("accountKey==")
                    .withTable("samples")
                    .withPartitionKey("partitionKey")
                    .withRowKey("rowKey")
                    .withColumnsToRemove(Arrays.asList("column1", "column2"))
                    .withBatchSize(25))
            .create();
    }
}
```
