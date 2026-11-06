
/**
 * Samples for NeuroShieldProtectedResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldProtectedResources_List.json
     */
    /**
     * Sample code: List protected resources.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        listProtectedResources(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldProtectedResources().list("myResourceGroup", "myNeuroShield", null, null,
            com.azure.core.util.Context.NONE);
    }
}
