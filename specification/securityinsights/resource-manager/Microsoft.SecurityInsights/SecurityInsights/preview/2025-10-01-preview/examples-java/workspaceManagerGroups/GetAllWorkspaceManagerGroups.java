
/**
 * Samples for WorkspaceManagerGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerGroups/GetAllWorkspaceManagerGroups.json
     */
    /**
     * Sample code: Get all workspace manager groups in the Sentinel workspace manager.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllWorkspaceManagerGroupsInTheSentinelWorkspaceManager(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerGroups().list("myRg", "myWorkspace", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
