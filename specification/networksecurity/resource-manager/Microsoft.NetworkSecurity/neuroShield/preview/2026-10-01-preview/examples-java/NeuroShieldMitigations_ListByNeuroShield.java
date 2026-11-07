
/**
 * Samples for NeuroShieldMitigations ListByNeuroShield.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldMitigations_ListByNeuroShield.json
     */
    /**
     * Sample code: List mitigations by NeuroShield.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        listMitigationsByNeuroShield(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldMitigations().listByNeuroShield("myResourceGroup", "myNeuroShield", 50, null,
            com.azure.core.util.Context.NONE);
    }
}
