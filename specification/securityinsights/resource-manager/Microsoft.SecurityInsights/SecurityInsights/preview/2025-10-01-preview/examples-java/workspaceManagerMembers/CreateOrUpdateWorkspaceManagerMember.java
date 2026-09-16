
/**
 * Samples for WorkspaceManagerMembers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerMembers/CreateOrUpdateWorkspaceManagerMember.json
     */
    /**
     * Sample code: Create or Update a workspace manager member.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createOrUpdateAWorkspaceManagerMember(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerMembers().define("afbd324f-6c48-459c-8710-8d1e1cd03812")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withTargetWorkspaceResourceId(
                "/subscriptions/7aef9d48-814f-45ad-b644-b0343316e174/resourceGroups/otherRg/providers/Microsoft.OperationalInsights/workspaces/Example_Workspace")
            .withTargetWorkspaceTenantId("f676d436-8d16-42db-81b7-ab578e110ccd").create();
    }
}
