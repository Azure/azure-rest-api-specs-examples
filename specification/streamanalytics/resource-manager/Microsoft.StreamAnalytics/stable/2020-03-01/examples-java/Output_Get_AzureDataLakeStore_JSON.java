import com.azure.core.util.Context;

/** Samples for Outputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Get_AzureDataLakeStore_JSON.json
     */
    /**
     * Sample code: Get an Azure Data Lake Store output with JSON serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAnAzureDataLakeStoreOutputWithJSONSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg6912", "sj3310", "output5195", Context.NONE);
    }
}
