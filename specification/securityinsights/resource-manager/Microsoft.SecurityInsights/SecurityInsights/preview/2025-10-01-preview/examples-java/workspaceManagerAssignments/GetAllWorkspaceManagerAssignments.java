
/**
 * Samples for WorkspaceManagerAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerAssignments/GetAllWorkspaceManagerAssignments.json
     */
    /**
     * Sample code: Get all workspace manager assignments for the Sentinel workspace manager.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllWorkspaceManagerAssignmentsForTheSentinelWorkspaceManager(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerAssignments().list("myRg", "myWorkspace", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
