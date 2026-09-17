
/**
 * Samples for WorkspaceManagerAssignments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerAssignments/GetWorkspaceManagerAssignment.json
     */
    /**
     * Sample code: Get a workspace manager assignment.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAWorkspaceManagerAssignment(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerAssignments().getWithResponse("myRg", "myWorkspace",
            "47cdc5f5-37c4-47b5-bd5f-83c84b8bdd58", com.azure.core.util.Context.NONE);
    }
}
