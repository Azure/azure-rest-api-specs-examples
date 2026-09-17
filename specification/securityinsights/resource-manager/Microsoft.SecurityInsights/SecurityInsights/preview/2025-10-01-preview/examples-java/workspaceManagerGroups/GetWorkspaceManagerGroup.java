
/**
 * Samples for WorkspaceManagerGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerGroups/GetWorkspaceManagerGroup.json
     */
    /**
     * Sample code: Get a workspace manager group.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAWorkspaceManagerGroup(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerGroups().getWithResponse("myRg", "myWorkspace", "37207a7a-3b8a-438f-a559-c7df400e1b96",
            com.azure.core.util.Context.NONE);
    }
}
