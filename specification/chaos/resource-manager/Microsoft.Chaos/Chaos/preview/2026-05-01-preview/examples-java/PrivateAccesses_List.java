
/**
 * Samples for PrivateAccesses ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateAccesses_List.json
     */
    /**
     * Sample code: List all private access in a resource group.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllPrivateAccessInAResourceGroup(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.privateAccesses().listByResourceGroup("myResourceGroup", null, com.azure.core.util.Context.NONE);
    }
}
