
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WorkspacesGet.json
     */
    /**
     * Sample code: WorkspaceGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspaceGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("oiautorest6685", "oiautorest6685",
            com.azure.core.util.Context.NONE);
    }
}
