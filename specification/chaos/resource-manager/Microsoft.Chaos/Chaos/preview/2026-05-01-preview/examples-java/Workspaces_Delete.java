
/**
 * Samples for Workspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Workspaces_Delete.json
     */
    /**
     * Sample code: Delete a Workspace in a resource group.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void deleteAWorkspaceInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().delete("exampleRG", "exampleWorkspace", com.azure.core.util.Context.NONE);
    }
}
