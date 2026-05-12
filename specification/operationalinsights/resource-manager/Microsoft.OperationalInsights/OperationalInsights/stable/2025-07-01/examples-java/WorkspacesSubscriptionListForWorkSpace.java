
/**
 * Samples for Workspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WorkspacesSubscriptionListForWorkSpace.json
     */
    /**
     * Sample code: WorkspacesSubscriptionList.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesSubscriptionList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.workspaces().list(com.azure.core.util.Context.NONE);
    }
}
