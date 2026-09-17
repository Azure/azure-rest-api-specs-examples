
import com.azure.resourcemanager.securityinsights.models.Mode;

/**
 * Samples for WorkspaceManagerConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-10-01-preview/workspaceManagerConfigurations/CreateOrUpdateWorkspaceManagerConfiguration.json
     */
    /**
     * Sample code: Create or Update a workspace manager Configuration.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createOrUpdateAWorkspaceManagerConfiguration(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerConfigurations().define("default").withExistingWorkspace("myRg", "myWorkspace")
            .withMode(Mode.ENABLED).create();
    }
}
