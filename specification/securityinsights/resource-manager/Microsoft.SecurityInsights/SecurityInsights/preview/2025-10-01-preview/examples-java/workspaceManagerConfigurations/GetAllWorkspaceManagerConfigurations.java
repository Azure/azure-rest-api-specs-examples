
/**
 * Samples for WorkspaceManagerConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerConfigurations/GetAllWorkspaceManagerConfigurations.json
     */
    /**
     * Sample code: Get all workspace manager configurations for a Sentinel workspace.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllWorkspaceManagerConfigurationsForASentinelWorkspace(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerConfigurations().list("myRg", "myWorkspace", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
