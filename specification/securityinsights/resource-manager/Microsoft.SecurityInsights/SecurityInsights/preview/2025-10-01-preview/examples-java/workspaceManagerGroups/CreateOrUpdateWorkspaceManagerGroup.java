
import java.util.Arrays;

/**
 * Samples for WorkspaceManagerGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/workspaceManagerGroups/CreateOrUpdateWorkspaceManagerGroup.json
     */
    /**
     * Sample code: Creates or updates a workspace manager group.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAWorkspaceManagerGroup(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.workspaceManagerGroups().define("37207a7a-3b8a-438f-a559-c7df400e1b96")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withDescription("Group of all financial and banking institutions").withDisplayName("Banks")
            .withMemberResourceNames(
                Arrays.asList("afbd324f-6c48-459c-8710-8d1e1cd03812", "f5fa104e-c0e3-4747-9182-d342dc048a9e"))
            .create();
    }
}
