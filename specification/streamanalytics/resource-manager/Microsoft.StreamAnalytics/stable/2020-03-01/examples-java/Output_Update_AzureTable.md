Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.AzureTableOutputDataSource;
import com.azure.resourcemanager.streamanalytics.models.Output;

/** Samples for Outputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_AzureTable.json
     */
    /**
     * Sample code: Update an Azure Table output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAnAzureTableOutput(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource = manager.outputs().getWithResponse("sjrg5176", "sj2790", "output958", Context.NONE).getValue();
        resource
            .update()
            .withDatasource(new AzureTableOutputDataSource().withPartitionKey("differentPartitionKey"))
            .apply();
    }
}
```
