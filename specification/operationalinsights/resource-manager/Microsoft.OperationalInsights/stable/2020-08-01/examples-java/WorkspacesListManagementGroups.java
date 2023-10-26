/** Samples for ManagementGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesListManagementGroups.json
     */
    /**
     * Sample code: WorkspacesListManagementGroups.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void workspacesListManagementGroups(
        com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.managementGroups().list("rg1", "TestLinkWS", com.azure.core.util.Context.NONE);
    }
}
