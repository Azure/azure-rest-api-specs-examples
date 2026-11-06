
/**
 * Samples for NeuroShields GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShields_Get.json
     */
    /**
     * Sample code: Get NeuroShield resource.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        getNeuroShieldResource(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShields().getByResourceGroupWithResponse("myResourceGroup", "myNeuroShield",
            com.azure.core.util.Context.NONE);
    }
}
