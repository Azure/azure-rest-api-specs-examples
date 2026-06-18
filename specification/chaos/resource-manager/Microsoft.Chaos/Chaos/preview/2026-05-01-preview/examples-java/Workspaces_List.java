
/**
 * Samples for Workspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Workspaces_List.json
     */
    /**
     * Sample code: List all Workspaces in a resource group.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllWorkspacesInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().listByResourceGroup("exampleRG", null, com.azure.core.util.Context.NONE);
    }
}
