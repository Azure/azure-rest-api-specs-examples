
/**
 * Samples for AvailableServiceTiers ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WorkspacesAvailableServiceTiers.json
     */
    /**
     * Sample code: AvailableServiceTiers.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void availableServiceTiers(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.availableServiceTiers().listByWorkspaceWithResponse("rg1", "workspace1",
            com.azure.core.util.Context.NONE);
    }
}
