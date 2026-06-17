
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Workspaces_Get.json
     */
    /**
     * Sample code: Get a Workspace in a resource group.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAWorkspaceInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("exampleRG", "exampleWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
