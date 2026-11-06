
/**
 * Samples for NeuroShieldProtectedResources Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldProtectedResources_Delete.json
     */
    /**
     * Sample code: Delete protected resource.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        deleteProtectedResource(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldProtectedResources().delete("myResourceGroup", "myNeuroShield", "vm-01", null,
            com.azure.core.util.Context.NONE);
    }
}
