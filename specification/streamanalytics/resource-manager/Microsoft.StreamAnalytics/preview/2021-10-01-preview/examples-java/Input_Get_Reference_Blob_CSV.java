
/**
 * Samples for Inputs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/
     * Input_Get_Reference_Blob_CSV.json
     */
    /**
     * Sample code: Get a reference blob input with CSV serialization.
     * 
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void getAReferenceBlobInputWithCSVSerialization(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().getWithResponse("sjrg8440", "sj9597", "input7225", com.azure.core.util.Context.NONE);
    }
}
