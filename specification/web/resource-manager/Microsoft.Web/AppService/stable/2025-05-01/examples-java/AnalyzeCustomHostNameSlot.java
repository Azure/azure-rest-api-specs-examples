
/**
 * Samples for WebApps AnalyzeCustomHostnameSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AnalyzeCustomHostNameSlot.json
     */
    /**
     * Sample code: Analyze custom hostname for slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void analyzeCustomHostnameForSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().analyzeCustomHostnameSlotWithResponse("testrg123", "sitef6141", "staging",
            null, com.azure.core.util.Context.NONE);
    }
}
