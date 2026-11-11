
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Workspaces_Get_EncryptionFailed.json
     */
    /**
     * Sample code: Get a Workspace after a customer key replacement fails.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        getAWorkspaceAfterACustomerKeyReplacementFails(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("exampleRG", "exampleWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
