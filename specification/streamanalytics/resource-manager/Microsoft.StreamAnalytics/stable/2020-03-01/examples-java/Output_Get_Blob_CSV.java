import com.azure.core.util.Context;

/** Samples for Outputs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Get_Blob_CSV.json
     */
    /**
     * Sample code: Get a blob output with CSV serialization.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getABlobOutputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.outputs().getWithResponse("sjrg5023", "sj900", "output1623", Context.NONE);
    }
}
