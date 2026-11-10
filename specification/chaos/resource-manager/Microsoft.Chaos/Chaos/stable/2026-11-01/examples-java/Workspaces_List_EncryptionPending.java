
/**
 * Samples for Workspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Workspaces_List_EncryptionPending.json
     */
    /**
     * Sample code: List CMK-created Workspaces with customer key replacement pending.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listCMKCreatedWorkspacesWithCustomerKeyReplacementPending(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaces().listByResourceGroup("exampleRG", null, com.azure.core.util.Context.NONE);
    }
}
