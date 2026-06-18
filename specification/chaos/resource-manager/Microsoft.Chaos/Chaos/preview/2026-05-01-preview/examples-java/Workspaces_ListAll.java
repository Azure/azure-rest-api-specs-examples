
/**
 * Samples for Workspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Workspaces_ListAll.json
     */
    /**
     * Sample code: List all Workspaces in a subscription.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllWorkspacesInASubscription(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().list(null, com.azure.core.util.Context.NONE);
    }
}
