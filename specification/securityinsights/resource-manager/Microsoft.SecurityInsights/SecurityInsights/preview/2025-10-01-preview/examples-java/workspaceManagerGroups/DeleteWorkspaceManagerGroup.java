
/**
 * Samples for WorkspaceManagerGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerGroups/DeleteWorkspaceManagerGroup.json
     */
    /**
     * Sample code: Delete a workspace manager group.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteAWorkspaceManagerGroup(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerGroups().deleteWithResponse("myRg", "myWorkspace",
            "37207a7a-3b8a-438f-a559-c7df400e1b96", com.azure.core.util.Context.NONE);
    }
}
