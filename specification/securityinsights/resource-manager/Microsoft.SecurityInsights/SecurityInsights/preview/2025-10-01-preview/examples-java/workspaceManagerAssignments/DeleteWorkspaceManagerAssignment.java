
/**
 * Samples for WorkspaceManagerAssignments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerAssignments/DeleteWorkspaceManagerAssignment.json
     */
    /**
     * Sample code: Delete a workspace manager assignment.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteAWorkspaceManagerAssignment(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerAssignments().deleteWithResponse("myRg", "myWorkspace",
            "47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58", com.azure.core.util.Context.NONE);
    }
}
