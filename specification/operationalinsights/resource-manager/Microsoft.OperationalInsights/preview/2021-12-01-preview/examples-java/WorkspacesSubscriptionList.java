import com.azure.core.util.Context;

/** Samples for DeletedWorkspaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesSubscriptionList.json
     */
    /**
     * Sample code: WorkspacesSubscriptionList.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesSubscriptionList(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.deletedWorkspaces().list(Context.NONE);
    }
}
