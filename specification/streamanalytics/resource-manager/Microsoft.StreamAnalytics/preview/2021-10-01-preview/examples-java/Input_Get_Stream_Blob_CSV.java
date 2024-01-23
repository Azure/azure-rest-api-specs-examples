
/**
 * Samples for Inputs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Get_Stream_Blob_CSV.json
     */
    /**
     * Sample code: Get a stream blob input with CSV serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAStreamBlobInputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().getWithResponse("sjrg8161", "sj6695", "input8899", com.azure.core.util.Context.NONE);
    }
}
