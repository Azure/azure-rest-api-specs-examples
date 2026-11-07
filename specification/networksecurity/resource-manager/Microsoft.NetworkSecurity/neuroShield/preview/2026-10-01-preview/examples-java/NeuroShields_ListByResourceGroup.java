
/**
 * Samples for NeuroShields ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShields_ListByResourceGroup.json
     */
    /**
     * Sample code: List NeuroShield resources by resource group.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void listNeuroShieldResourcesByResourceGroup(
        com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShields().listByResourceGroup("myResourceGroup", null, null, com.azure.core.util.Context.NONE);
    }
}
