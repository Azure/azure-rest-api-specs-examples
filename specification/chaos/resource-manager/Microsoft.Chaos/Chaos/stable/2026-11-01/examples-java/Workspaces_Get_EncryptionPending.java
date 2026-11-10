
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Workspaces_Get_EncryptionPending.json
     */
    /**
     * Sample code: Get a CMK-created Workspace while its customer key is being replaced.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getACMKCreatedWorkspaceWhileItsCustomerKeyIsBeingReplaced(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("exampleRG", "exampleWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
