
/**
 * Samples for WorkspaceManagerConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerConfigurations/GetWorkspaceManagerConfiguration.json
     */
    /**
     * Sample code: Get a workspace manager configuration.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAWorkspaceManagerConfiguration(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerConfigurations().getWithResponse("myRg", "myWorkspace", "default",
            com.azure.core.util.Context.NONE);
    }
}
