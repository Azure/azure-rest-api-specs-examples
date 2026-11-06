
/**
 * Samples for NeuroShieldMitigations Execute.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShieldMitigations_Execute.json
     */
    /**
     * Sample code: Execute mitigation.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void executeMitigation(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShieldMitigations().executeWithResponse("myResourceGroup", "myNeuroShield", "vm-01",
            "mitigation-01", com.azure.core.util.Context.NONE);
    }
}
