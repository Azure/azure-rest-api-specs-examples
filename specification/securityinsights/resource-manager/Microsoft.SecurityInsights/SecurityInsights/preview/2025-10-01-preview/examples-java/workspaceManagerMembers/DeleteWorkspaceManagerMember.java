
/**
 * Samples for WorkspaceManagerMembers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerMembers/DeleteWorkspaceManagerMember.json
     */
    /**
     * Sample code: Delete a workspace manager member.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        deleteAWorkspaceManagerMember(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerMembers().deleteWithResponse("myRg", "myWorkspace",
            "afbd324f-6c48-459c-8710-8d1e1cd03812", com.azure.core.util.Context.NONE);
    }
}
