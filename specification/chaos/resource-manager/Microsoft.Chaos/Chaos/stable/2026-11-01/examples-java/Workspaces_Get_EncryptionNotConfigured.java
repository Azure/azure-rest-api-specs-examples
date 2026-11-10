
/**
 * Samples for Workspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Workspaces_Get_EncryptionNotConfigured.json
     */
    /**
     * Sample code: Get a Workspace service application before customer key enrollment.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void getAWorkspaceServiceApplicationBeforeCustomerKeyEnrollment(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("exampleRG", "exampleWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
