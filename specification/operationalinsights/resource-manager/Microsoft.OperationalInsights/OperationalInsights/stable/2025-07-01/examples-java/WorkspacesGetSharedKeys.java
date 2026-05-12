
/**
 * Samples for SharedKeysOperation GetSharedKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WorkspacesGetSharedKeys.json
     */
    /**
     * Sample code: SharedKeysList.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void sharedKeysList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.sharedKeysOperations().getSharedKeysWithResponse("rg1", "TestLinkWS", com.azure.core.util.Context.NONE);
    }
}
