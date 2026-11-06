
/**
 * Samples for NeuroShieldProtectedResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldProtectedResources_Get.json
     */
    /**
     * Sample code: Get protected resource.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void getProtectedResource(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldProtectedResources().getWithResponse("myResourceGroup", "myNeuroShield", "vm-01",
            com.azure.core.util.Context.NONE);
    }
}
