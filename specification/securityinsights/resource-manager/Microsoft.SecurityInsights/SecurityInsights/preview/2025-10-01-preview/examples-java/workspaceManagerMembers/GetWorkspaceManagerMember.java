
/**
 * Samples for WorkspaceManagerMembers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerMembers/GetWorkspaceManagerMember.json
     */
    /**
     * Sample code: Get a workspace manager member.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAWorkspaceManagerMember(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerMembers().getWithResponse("myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812",
            com.azure.core.util.Context.NONE);
    }
}
