```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.Output;
import com.azure.resourcemanager.streamanalytics.models.PowerBIOutputDataSource;

/** Samples for Outputs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Update_PowerBI.json
     */
    /**
     * Sample code: Update a Power BI output.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void updateAPowerBIOutput(com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        Output resource =
            manager.outputs().getWithResponse("sjrg7983", "sj2331", "output3022", Context.NONE).getValue();
        resource.update().withDatasource(new PowerBIOutputDataSource().withDataset("differentDataset")).apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
