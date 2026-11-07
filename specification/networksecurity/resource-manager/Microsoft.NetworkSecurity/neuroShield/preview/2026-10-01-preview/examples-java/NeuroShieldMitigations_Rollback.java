
/**
 * Samples for NeuroShieldMitigations Rollback.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldMitigations_Rollback.json
     */
    /**
     * Sample code: Rollback mitigation.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void rollbackMitigation(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldMitigations().rollbackWithResponse("myResourceGroup", "myNeuroShield", "vm-01",
            "mitigation-01", com.azure.core.util.Context.NONE);
    }
}
