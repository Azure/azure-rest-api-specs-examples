
/**
 * Samples for SharedKeysOperation Regenerate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesRegenerateSharedKeys.json
     */
    /**
     * Sample code: RegenerateSharedKeys.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void regenerateSharedKeys(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.sharedKeysOperations().regenerateWithResponse("rg1", "workspace1", com.azure.core.util.Context.NONE);
    }
}
