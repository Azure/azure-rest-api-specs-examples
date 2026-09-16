
/**
 * Samples for WorkspaceManagerMembers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerMembers/GetAllWorkspaceManagerMembers.json
     */
    /**
     * Sample code: Get all workspace manager members.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllWorkspaceManagerMembers(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerMembers().list("myRg", "myWorkspace", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
