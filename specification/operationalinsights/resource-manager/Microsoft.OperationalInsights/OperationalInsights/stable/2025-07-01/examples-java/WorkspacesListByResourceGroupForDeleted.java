
/**
 * Samples for DeletedWorkspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WorkspacesListByResourceGroupForDeleted.json
     */
    /**
     * Sample code: WorkspacesGet.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesGet(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.deletedWorkspaces().listByResourceGroup("oiautorest6685", com.azure.core.util.Context.NONE);
    }
}
