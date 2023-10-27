/** Samples for DeletedWorkspaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2022-10-01/examples/WorkspacesSubscriptionList.json
     */
    /**
     * Sample code: WorkspacesSubscriptionList.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesSubscriptionList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.deletedWorkspaces().list(com.azure.core.util.Context.NONE);
    }
}
