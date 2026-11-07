
/**
 * Samples for NeuroShields Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShields_Delete.json
     */
    /**
     * Sample code: Delete NeuroShield resource.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        deleteNeuroShieldResource(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShields().deleteByResourceGroupWithResponse("myResourceGroup", "myNeuroShield",
            com.azure.core.util.Context.NONE);
    }
}
