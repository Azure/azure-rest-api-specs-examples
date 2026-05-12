
/**
 * Samples for ManagementGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WorkspacesListManagementGroups.json
     */
    /**
     * Sample code: WorkspacesListManagementGroups.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void
        workspacesListManagementGroups(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.managementGroups().list("rg1", "TestLinkWS", com.azure.core.util.Context.NONE);
    }
}
