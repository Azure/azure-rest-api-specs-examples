
/**
 * Samples for WorkspaceManagerConfigurations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerConfigurations/DeleteWorkspaceManagerConfiguration.json
     */
    /**
     * Sample code: Delete a workspace manager configuration.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAWorkspaceManagerConfiguration(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerConfigurations().deleteWithResponse("myRg", "myWorkspace", "default",
            com.azure.core.util.Context.NONE);
    }
}
