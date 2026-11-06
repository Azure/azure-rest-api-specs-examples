
/**
 * Samples for NeuroShieldFindings ListByNeuroShield.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldFindings_ListByNeuroShield.json
     */
    /**
     * Sample code: List findings by NeuroShield.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        listFindingsByNeuroShield(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldFindings().listByNeuroShield("myResourceGroup", "myNeuroShield", 50, null,
            com.azure.core.util.Context.NONE);
    }
}
