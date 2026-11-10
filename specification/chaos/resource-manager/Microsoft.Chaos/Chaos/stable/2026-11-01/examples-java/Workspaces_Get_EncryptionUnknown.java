
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Workspaces_Get_EncryptionUnknown.json
     */
    /**
     * Sample code: Get a Workspace when active encryption protection is unknown.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        getAWorkspaceWhenActiveEncryptionProtectionIsUnknown(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("exampleRG", "exampleWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
