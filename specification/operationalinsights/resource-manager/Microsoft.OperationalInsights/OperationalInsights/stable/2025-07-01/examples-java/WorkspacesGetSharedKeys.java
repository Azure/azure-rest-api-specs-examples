
/**
 * Samples for SharedKeysOperation GetSharedKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesGetSharedKeys.json
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
