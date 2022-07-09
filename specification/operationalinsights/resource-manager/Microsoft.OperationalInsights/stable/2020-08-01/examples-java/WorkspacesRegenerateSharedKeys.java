import com.azure.core.util.Context;

/** Samples for SharedKeysOperation Regenerate. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesRegenerateSharedKeys.json
     */
    /**
     * Sample code: RegenerateSharedKeys.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void regenerateSharedKeys(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.sharedKeysOperations().regenerateWithResponse("rg1", "workspace1", Context.NONE);
    }
}
